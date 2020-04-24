package influxdb

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"

	"github.com/influxdata/influxql"
	"github.com/ovh/erlenmeyer/core"
	log "github.com/sirupsen/logrus"
)

// InfluxParser in the main structure when parsing a select statement
type InfluxParser struct {
	Classnames    []string
	Field         string
	BucketTime    string
	BucketCount   string
	End           string
	Start         string
	Token         string
	Shift         string
	GroupByTags   []string
	Fill          influxql.FillOption
	FillValue     interface{}
	TimeColumn    *InfluxTimeColumn
	HasGroupBy    bool
	Measurement   *InfluxMeasurement
	Response      *InfluxResponse
	TimePrecision string
	HasSubQuery   bool
	SubQueryLevel int
	StarQuery     bool
	HasWildCard   bool
	Separator     string
	KeepTopLabels []string
}

// Parse a select Statement
func (p *InfluxParser) getSelectStatementScript(statement *influxql.SelectStatement, warpServer *core.HTTPWarp10Server, txn string, subqueryLevel int, statementid int) (string, *Result, map[string]bool, map[string]bool, error) {
	selectValidField := make(map[string]bool)
	selectTagsField := make(map[string]bool)

	subQueries := make([]*influxql.SelectStatement, 0)

	p.Classnames = make([]string, 0)

	for _, source := range statement.Sources {
		switch source := source.(type) {
		case *influxql.Measurement:
			if source.Regex != nil {
				p.Classnames = append(p.Classnames, source.Regex.Val.String())
			} else {
				p.Classnames = append(p.Classnames, source.Name)
			}
		case *influxql.SubQuery:
			subQueries = append(subQueries, source.Statement)
		}
	}

	selectSubMc2 := make([]string, len(subQueries))
	duration, tags := statement.Dimensions.Normalize()
	hasSubQueries := len(subQueries) > 0
	for index, subQuery := range subQueries {
		subQueryParser := &InfluxParser{Classnames: make([]string, 0), Token: p.Token, End: "$end", BucketTime: "0", BucketCount: "1", Separator: p.Separator, KeepTopLabels: p.KeepTopLabels}

		prefix, resResp, subselectValidField, subselectTagsField, err := subQueryParser.getSelectStatementScript(subQuery, warpServer, txn, subqueryLevel+1, statementid)

		if resResp != nil {
			if resResp.Err != "" {
				return "", resResp, selectValidField, selectTagsField, nil
			}
		}
		if err != nil {
			return "", nil, selectValidField, selectTagsField, err
		}
		for key, value := range subselectValidField {
			selectValidField[key] = value
		}
		for key, value := range subselectTagsField {
			selectTagsField[key] = value
		}
		selectSubMc2[index] = prefix
	}

	mc2 := ""
	hasSub := false
	if len(selectSubMc2) > 0 {
		mc2 += strings.Join(selectSubMc2, " APPEND \n")
		mc2 += fmt.Sprintf(" 'subqueries-%d' STORE\n", subqueryLevel)
		hasSub = true
	}
	p.HasSubQuery = hasSub
	p.SubQueryLevel = subqueryLevel
	mc2 += fmt.Sprintf("/*\n %s \n*/\n", statement.String())
	mc2 += "NOW 'now' CSTORE\n"

	// Parse query mandatory info as start, end
	end, start := parseTimeCondition(statement.Condition)

	// Parse query to get separator when custom
	p.Separator, _ = parseSeparatorCondition(statement.Condition)
	mc2 += fmt.Sprintf("%s 'end' STORE \n", end)

	p.Fill = statement.Fill
	p.FillValue = statement.FillValue

	statement.RewriteTimeFields()

	p.TimeColumn = &InfluxTimeColumn{TimeAlias: statement.TimeFieldName(), OmitColumnTime: statement.OmitTime}

	p.Measurement = &InfluxMeasurement{EmitName: statement.EmitName, StripMeasurement: statement.StripName}

	p.Response = &InfluxResponse{Dedup: statement.Dedupe, Limit: statement.Limit, Offset: statement.Offset, SLimit: statement.SLimit, SOffset: statement.SOffset}

	for _, sort := range statement.SortFields {
		if sort.Name == "time" {
			p.TimeColumn.ReverseSort = !sort.Ascending
		} else {
			return "", nil, selectValidField, selectTagsField, fmt.Errorf("error parsing query: only ORDER BY time supported at this time")
		}
	}
	for _, tag := range tags {
		p.HasGroupBy = true
		p.GroupByTags = append(p.GroupByTags, "'"+tag+"'")
	}

	if duration != 0 {
		p.HasGroupBy = true
		groupByDuration := fmt.Sprintf("%d", duration.Nanoseconds()/1000)
		p.Shift = groupByDuration
		p.BucketTime = groupByDuration
		p.BucketCount = fmt.Sprintf("$interval %s / 1 +", groupByDuration)
	}

	if p.HasSubQuery && start == "0" {

	} else {
		if p.BucketTime != "0" {
			mc2 += fmt.Sprintf("%s %s 2 * - 'start'  STORE \n", start, p.BucketTime)
		} else {
			mc2 += fmt.Sprintf("%s 'start' STORE \n", start)
		}
		mc2 += "$end $start - 'interval' STORE \n"
	}
	selectedField := statement.Fields.AliasNames()

	findClass := "~"
	classname := fmt.Sprintf("(%s", strings.Join(p.Classnames, "|"))
	findResponses := [][]core.GeoTimeSeries{}
	where := make([][]*WhereCond, 0)
	if statement.Condition != nil {

		whereCond, err := p.parseWhere(statement.Condition)
		if err != nil {
			return "", nil, selectValidField, selectTagsField, err
		}
		where = whereCond
	}

	for _, field := range statement.Fields {
		separator := "\\."
		if p.Separator != "." {
			separator = p.Separator
		}

		if strings.Contains(field.String(), "*") {
			if !p.StarQuery {
				findClass += classname + ")" + separator + "(.*)|"
				p.StarQuery = true
			}
			continue
		}
		varRefNames := influxql.ExprNames(field.Expr)
		for _, ref := range varRefNames {
			findClass += classname + ")" + separator + "(" + ref.Val + ")|"
		}

		// Handle case of misformed varref returned by influx library
		if len(varRefNames) == 0 {
			findClass += classname + ")" + separator + "(.*)|"
		}
	}
	findClass += ""

	// Check if an error exists in ParseExpr prior to run any queries
	selectors := make([]string, 0)
	for _, field := range statement.Fields {
		_, _, err := p.parseExpr(field.Expr, 0, selectors, where)
		if err != nil {
			return "", &Result{StatementID: statementid, Err: err.Error()}, selectValidField, selectTagsField, nil
		}
	}

	if len(where) > 0 {
		findmc2 := fmt.Sprintf("[ '%s' '%s' {} ] \n", p.Token, findClass)
		findmc2 += `
		FINDSETS KEYLIST SWAP KEYLIST APPEND UNIQUE 'tags' STORE 
		<%
			DROP
			NEWGTS SWAP RENAME
			$tags
			<%
				{ 
				SWAP
				'true'
				}
				RELABEL
			%>
			FOREACH
		%>
		LMAP
		`
		findQuery, err := warpServer.Query(findmc2, txn)

		if err != nil {
			log.WithFields(log.Fields{
				"error": err.Error(),
				"proto": "influxQL",
			}).Error("Bad response from Egress")
			return "", nil, selectValidField, selectTagsField, err
		}
		bufferFind, err := ioutil.ReadAll(findQuery.Body)
		if err != nil {
			log.WithFields(log.Fields{
				"error": err.Error(),
				"proto": "influxQL",
			}).Error("can't fully read Egress response")
			return "", nil, selectValidField, selectTagsField, err
		}

		if findQuery.StatusCode != http.StatusOK {
			return "", &Result{Err: "Invalid generated query: returned error is - " + string(bufferFind)}, selectValidField, selectTagsField, nil
		}

		err = json.Unmarshal(bufferFind, &findResponses)
		if err != nil {
			log.WithFields(log.Fields{
				"error": err.Error(),
				"proto": "influxQL",
			}).Error("Cannot unmarshal egress response")
			return "", nil, selectValidField, selectTagsField, err
		}
	}

	for _, gtsSet := range findResponses {
		if len(gtsSet) == 0 && !hasSub {
			return "", &Result{}, selectValidField, selectTagsField, nil
		}

		for _, series := range gtsSet {
			validName := ""
			seriesName := strings.SplitAfterN(series.Class, p.Separator, 2)

			if len(seriesName) > 1 {
				validName = seriesName[1]
			} else {
				validName = series.Class
			}
			selectValidField[validName] = true

			for tagKey := range series.Labels {
				selectTagsField[tagKey] = true
			}
		}
	}

	selectors = make([]string, 0)

	for _, wheres := range where {
		if len(wheres) == 0 {
			continue
		}
		selector := make([]string, 0)

		for _, whereItem := range wheres {
			if _, containsTag := selectTagsField[whereItem.key]; containsTag {
				op := whereItem.op.String()

				rhsValue := whereItem.value.Val
				if strings.HasPrefix(whereItem.value.Val, "'") {
					rhsValue = strings.Trim(whereItem.value.Val, "'")
				} else if strings.HasPrefix(whereItem.value.Val, "\"") {
					rhsValue = strings.Trim(whereItem.value.Val, "\"")
				}
				if whereItem.op == influxql.EQREGEX {
					op = "~"
				}
				if whereItem.op == influxql.EQREGEX || whereItem.op == influxql.EQ {
					selector = append(selector, whereItem.key+op+rhsValue)
					whereItem.isTag = true
				} else if whereItem.op == influxql.NEQREGEX || whereItem.op == influxql.NEQ {

					value := fmt.Sprintf("(?!%s).*?", rhsValue)
					if whereItem.op == influxql.NEQREGEX {
						value = fmt.Sprintf("(?!%s).*?", rhsValue)
					}

					op = "~"
					selector = append(selector, whereItem.key+op+value)
					whereItem.isTag = true
				} else {
					return "", nil, selectValidField, selectTagsField, fmt.Errorf("Unsupported %s operator on tag key: %s", op, whereItem.key)
				}
			}
		}
		if len(selector) > 0 {
			selectors = append(selectors, strings.Join(selector, ","))
		}
	}

	renameField := make(map[string]int)

	for i, field := range statement.Fields {
		hasASeries := false
		starQuery := false

		if field.String() == "*" {
			field.Expr = &influxql.VarRef{Val: ".*"}
			starQuery = true
		} else if strings.Contains(field.String(), "*") {
			hasASeries = true
		}

		varRefNames := influxql.ExprNames(field.Expr)

		for _, ref := range varRefNames {
			_, hasASeries = selectValidField[ref.Val]
			if hasASeries || starQuery {
				break
			}
		}

		if !starQuery && (!hasASeries && len(where) > 0) && len(varRefNames) > 0 {
			continue
		}

		var expr string

		expr, _, err := p.parseExpr(field.Expr, 0, selectors, where)
		if err != nil {
			return "", &Result{StatementID: statementid, Err: err.Error()}, selectValidField, selectTagsField, nil
		}

		if p.HasGroupBy {
			expr += p.startPartition()
			expr += `
				+
			%>
			FOREACH
			FLATTEN
			`
		}

		mc2 += expr

		if p.HasWildCard {
			wildCards := append(selectedField, ".INFLUXQL_COLUMN_NAME")
			mc2 += removeAllLabels(wildCards)
		} else if !starQuery {
			mc2 += removeAllLabels(append(selectedField, p.KeepTopLabels...))
			mc2 += p.renameAndSetInfluxLabels()
		} else {
			mc2 += p.renameAndSetInfluxStarLabels()
		}
		mc2 += p.shiftSeries()

		fieldName := field.Name()
		idField, containsField := renameField[fieldName]

		if !containsField {
			renameField[fieldName] = 1
		} else {
			renameField[fieldName] = idField + 1
			fieldName = fmt.Sprintf("%s_%d", fieldName, idField)
		}

		if p.HasWildCard {
		} else if !starQuery {
			if !hasSubQueries {
				mc2 += fmt.Sprintf(" { '.INFLUXQL_COLUMN_NAME' '%s' } RELABEL \n", fieldName)
			}
		} else {
			mc2 += " { '.InfluxDBName' NULL } RELABEL \n"
		}

		if i >= 1 {
			mc2 += " APPEND \n"
		}
		selectValidField[fieldName] = true
		// Apply timeshift $end $start - only in case agg without groupBy
		// When groupBy is apply shift by $duration
	}

	mc2 += formatResultSeries(p.Response)

	return mc2, nil, selectValidField, selectTagsField, nil
}
