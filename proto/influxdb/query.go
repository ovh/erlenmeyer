package influxdb

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"math"
	"net/http"
	"net/http/httputil"
	"regexp"
	"sort"
	"strings"
	"time"

	"github.com/influxdata/influxql"
	"github.com/ovh/erlenmeyer/core"
	"github.com/ovh/erlenmeyer/middlewares"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

const (
	secondTonanoSecond = float64(1000000000)
)

type (
	// Error structure
	Error struct {
		Message string `json:"error,omitempty"`
	}
)

// InfluxTimeColumn data
type InfluxTimeColumn struct {
	ReverseSort    bool
	TimeAlias      string
	OmitColumnTime bool
}

// InfluxMeasurement data
type InfluxMeasurement struct {
	StripMeasurement bool
	EmitName         string
}

// InfluxResponse data
type InfluxResponse struct {
	Dedup   bool
	Limit   int
	Offset  int
	SLimit  int
	SOffset int
}

type influxValues struct {
	Timestamp float64
	Columns   map[string]interface{}
}

// InfluxShowStatement in the main structure when parsing a show statement
type InfluxShowStatement struct {
	QueryType  ShowType
	Name       string
	Columns    []string
	TagKeyExpr influxql.Literal
	Separator  string
}

// ShowType represents the FIND operation type
type ShowType int

const (
	// ShowMeasurements type.
	ShowMeasurements ShowType = 0
	// ShowFieldKeys type.
	ShowFieldKeys ShowType = 1
	// ShowTagKeys type.
	ShowTagKeys ShowType = 2
	// ShowSeries type.
	ShowSeries ShowType = 3
	// ShowTagValues type.
	ShowTagValues ShowType = 4
	// ShowTagValuesCardinality type.
	ShowTagValuesCardinality ShowType = 5
)

// Query is handling Query
// Taken from https://github.com/influxdata/influxdb/blob/f19588360e601ffe5a0bc11c1af1e2fa7fff1112/services/httpd/handler.go#L307
func (i *InfluxDB) Query(w http.ResponseWriter, r *http.Request) {

	if r.Method != http.MethodPost && r.Method != http.MethodGet {
		i.WarnCounter.Inc()
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}

	token := core.RetrieveToken(r)
	if len(token) == 0 {
		i.WarnCounter.Inc()
		http.Error(w, http.StatusText(http.StatusUnauthorized), http.StatusUnauthorized)
		return
	}

	var qr io.Reader
	// Attempt to read the form value from the "q" form value.
	if qp := strings.TrimSpace(r.FormValue("q")); qp != "" {
		qr = strings.NewReader(qp)
	} else if r.MultipartForm != nil && r.MultipartForm.File != nil {
		// If we have a multipart/form-data, try to retrieve a file from 'q'.
		if fhs := r.MultipartForm.File["q"]; len(fhs) > 0 {
			f, err := fhs[0].Open()
			if err != nil {
				http.Error(w, err.Error(), http.StatusBadRequest)
				return
			}
			defer f.Close()
			qr = f
		}
	}

	if qr == nil {
		http.Error(w, `missing required parameter "q"`, http.StatusBadRequest)
		return
	}

	epoch := strings.TrimSpace(r.FormValue("epoch"))

	p := influxql.NewParser(qr)
	db := r.FormValue("db")
	timePrecision := r.FormValue("precision")

	if timePrecision == "" {
		if epoch == "" {
			timePrecision = "rfc3339"
		} else {
			timePrecision = epoch
		}
	}

	// Parse query from query string.
	q, err := p.ParseQuery()

	if err != nil {
		influxParsing := &SimpleErrorResult{Err: "error parsing query: " + err.Error()}
		errorString, errJSON := json.Marshal(influxParsing)
		if errJSON != nil {
			log.WithFields(log.Fields{
				"type":  "query",
				"proto": "influxQL",
				"Error": err.Error(),
			}).Error("error when parsing an Influx Error")
			http.Error(w, "error parsing query: "+err.Error(), http.StatusInternalServerError)
		}
		http.Error(w, string(errorString), http.StatusBadRequest)
		return
	}

	body, err := handleQuery(q, epoch, db, w.Header().Get(middlewares.TxnHeader), token, timePrecision)
	if err != nil {

		influxParsing := &SimpleErrorResult{Err: "error executing query: " + err.Error()}
		errorString, errJSON := json.Marshal(influxParsing)
		if errJSON != nil {
			log.WithFields(log.Fields{
				"type":  "query",
				"proto": "influxQL",
				"Error": err.Error(),
			}).Error("error when parsing an Influx Error")
			http.Error(w, "error parsing query: "+err.Error(), http.StatusInternalServerError)
		}
		http.Error(w, string(errorString), http.StatusBadRequest)
		return
	}

	i.ReqCounter.Inc()
	w.WriteHeader(http.StatusOK)
	if body != nil {
		if _, err = w.Write(body); err != nil {
			log.WithError(err).Error("Could not awnser to the influx request")
		}
	}
}

func handleQuery(q *influxql.Query, epoch string, db string, txn string, token string, timePrecision string) ([]byte, error) {

	// Use the copied Response of InfluxQL:
	// https://github.com/influxdata/influxdb/blob/master/query/influxql/go
	influx := &Response{}

	// Support for Select ans some needed statements
	for i, statement := range q.Statements {
		switch stmt := statement.(type) {
		case *influxql.SelectStatement:
			result, err := parseInfluxSelect(stmt, i, txn, token, timePrecision)
			if err != nil {
				return nil, err
			}
			influx.Results = append(influx.Results, *result)
		case *influxql.CreateDatabaseStatement:
			influx.Results = append(influx.Results, defaultCreateDB(i))
		case *influxql.ShowDatabasesStatement:
			influx.Results = append(influx.Results, defaultShowDatabasesStatement(i))
		case *influxql.ShowSeriesStatement:
			separator, _ := parseSeparatorCondition(stmt.Condition)
			showStatement := &InfluxShowStatement{QueryType: ShowSeries, Name: "", Columns: []string{"key"}, Separator: separator}
			result, err := showStatement.parseInfluxSeries(i, txn, token, stmt.Limit, stmt.Offset, stmt.Sources, stmt.Condition)
			if err != nil {
				return nil, err
			}
			influx.Results = append(influx.Results, *result)
		case *influxql.ShowFieldKeysStatement:
			showStatement := &InfluxShowStatement{QueryType: ShowFieldKeys, Name: "none", Columns: []string{"fieldKey", "fieldType"}, Separator: "."}
			result, err := showStatement.parseInfluxSeries(i, txn, token, stmt.Limit, stmt.Offset, stmt.Sources, nil)
			if err != nil {
				return nil, err
			}
			influx.Results = append(influx.Results, *result)
		case *influxql.ShowTagKeysStatement:
			separator, _ := parseSeparatorCondition(stmt.Condition)
			showStatement := &InfluxShowStatement{QueryType: ShowTagKeys, Name: "none", Columns: []string{"tagKey"}, Separator: separator}
			result, err := showStatement.parseInfluxSeries(i, txn, token, stmt.Limit, stmt.Offset, stmt.Sources, stmt.Condition)
			if err != nil {
				return nil, err
			}
			influx.Results = append(influx.Results, *result)
		case *influxql.ShowTagValuesStatement:
			separator, _ := parseSeparatorCondition(stmt.Condition)
			showStatement := &InfluxShowStatement{QueryType: ShowTagValues, Name: "none", Columns: []string{"key", "value"}, TagKeyExpr: stmt.TagKeyExpr, Separator: separator}

			result, err := showStatement.parseInfluxSeries(i, txn, token, stmt.Limit, stmt.Offset, stmt.Sources, stmt.Condition)
			if err != nil {
				return nil, err
			}
			influx.Results = append(influx.Results, *result)
		case *influxql.DropDatabaseStatement:
			return nil, errors.New("Please delete series to drop your account")
		case *influxql.ShowRetentionPoliciesStatement:
			influx.Results = append(influx.Results, defaultShowRetentionPoliciesStatement(i))
		case *influxql.ShowMeasurementsStatement:
			separator, _ := parseSeparatorCondition(stmt.Condition)
			showStatement := &InfluxShowStatement{QueryType: ShowMeasurements, Name: "measurements", Columns: []string{"name"}, Separator: separator}
			tmpSources := make([]influxql.Source, 1)
			tmpSources[0] = stmt.Source
			result, err := showStatement.parseInfluxSeries(i, txn, token, stmt.Limit, stmt.Offset, tmpSources, stmt.Condition)
			if err != nil {
				return nil, err
			}
			influx.Results = append(influx.Results, *result)
		case *influxql.ShowTagValuesCardinalityStatement:
			separator, _ := parseSeparatorCondition(stmt.Condition)
			showStatement := &InfluxShowStatement{QueryType: ShowTagValuesCardinality, Name: "none", Columns: []string{"count"}, TagKeyExpr: stmt.TagKeyExpr, Separator: separator}

			result, err := showStatement.parseInfluxSeries(i, txn, token, stmt.Limit, stmt.Offset, stmt.Sources, stmt.Condition)
			if err != nil {
				return nil, err
			}
			influx.Results = append(influx.Results, *result)
		default:
			log.Warnf("handleQuery - Default %s %T", stmt.String(), stmt)
			return nil, fmt.Errorf("Statement not implemented yet: %T", stmt)
		}
	}
	return json.Marshal(influx)
}

// parseInfluxSeries parse ALL SHOW META Influx statements kind
func (showStatement *InfluxShowStatement) parseInfluxSeries(statementid int, txn string, token string, paramLimit, offset int, sources influxql.Sources, condition influxql.Expr) (*Result, error) {

	warpServer := core.NewWarpServer(viper.GetString("warp_endpoint"), "influxql")

	limit := paramLimit
	if paramLimit <= 0 {
		limit = math.MaxInt64
	}

	classnames := make([]string, 0)

	for _, source := range sources {
		switch source := source.(type) {
		case *influxql.Measurement:
			if source.Regex != nil {
				classnames = append(classnames, source.Regex.Val.String()+".*")
			} else {
				classnames = append(classnames, regexp.QuoteMeta(source.Name))
			}
		}
	}

	if len(classnames) == 0 {
		classnames = append(classnames, ".*")
	}

	classname := fmt.Sprintf("(%s", strings.Join(classnames, "|"))
	where := make([][]*WhereCond, 0)

	keepTopLabels := make([]string, 0)

	p := &InfluxParser{Classnames: classnames, Token: token, End: "$end", BucketTime: "0", BucketCount: "1", Separator: showStatement.Separator, KeepTopLabels: keepTopLabels}

	if condition != nil {

		whereCond, err := p.parseWhere(condition)
		if err != nil {
			return &Result{}, nil
		}
		where = whereCond
	}

	labelsFilter := "{}"

	if len(where) > 0 {
		labelsFilter = "{"

		for _, wheres := range where {
			if len(wheres) == 0 {
				continue
			}
			for _, whereItem := range wheres {

				op := whereItem.op.String()

				rhsValue := whereItem.value.Val
				if strings.HasPrefix(whereItem.value.Val, "'") {
					rhsValue = strings.Trim(whereItem.value.Val, "'")
				} else if strings.HasPrefix(whereItem.value.Val, "\"") {
					rhsValue = strings.Trim(whereItem.value.Val, "\"")
				}

				value := fmt.Sprintf("%s", rhsValue)

				switch whereItem.op {
				case influxql.EQREGEX:
					op = "~"
				case influxql.EQ:
					op = "="
				case influxql.NEQREGEX, influxql.NEQ:
					op = "~"
					value = fmt.Sprintf("(?!%s).*?", rhsValue)
				default:
					return nil, fmt.Errorf("Unsupported %s operator on tag key: %s", op, whereItem.key)
				}

				labelsFilter += fmt.Sprintf(" '%s' '%s%s'", whereItem.key, op, value)
			}
		}

		labelsFilter += " }"
	}

	findClass := classname

	separator := "\\."
	if p.Separator != "." {
		separator = p.Separator
	}

	findSelector := fmt.Sprintf("~%s)"+separator+".*", findClass)

	mc2 := fmt.Sprintf("'%s' AUTHENTICATE\n", token)
	mc2 += `
	'stack.maxops.hard' STACKATTRIBUTE DUP <% ISNULL ! %> <% MAXOPS %> <% DROP %> IFTE
	'fetch.limit.hard' STACKATTRIBUTE DUP <% ISNULL ! %> <% LIMIT %> <% DROP %> IFTE
	'gts.limit.hard' STACKATTRIBUTE DUP <% ISNULL ! %> <% MAXGTS %> <% DROP %> IFTE
	`

	switch showStatement.QueryType {
	case ShowSeries:
		mc2 += getShowSeriesWarpScript(token, findSelector, labelsFilter, offset, limit, showStatement.Separator)
		mc2 += `
		<% DROP <% DROP NEWGTS SWAP RENAME NOW NaN NaN NaN 1 ADDVALUE %> LMAP %> LMAP
		`
	case ShowMeasurements:
		mc2 += getShowMeasurementsWarpScript(token, findSelector, labelsFilter, offset, limit, showStatement.Separator)
		mc2 += `
		<% DROP <% DROP NEWGTS SWAP RENAME NOW NaN NaN NaN 1 ADDVALUE %> LMAP %> LMAP
		`
	case ShowFieldKeys:
		mc2 += getShowFieldKeysWarpScript(token, findSelector, labelsFilter, offset, limit, showStatement.Separator)
	case ShowTagKeys:
		mc2 += getShowTagKeysWarpScript(token, findSelector, labelsFilter, offset, limit, showStatement.Separator)
	case ShowTagValues, ShowTagValuesCardinality:
		switch showStatement.TagKeyExpr.(type) {
		case *influxql.RegexLiteral:
			regExp := showStatement.TagKeyExpr.String()
			regExp = strings.TrimLeft(regExp, "/")
			regExp = strings.TrimRight(regExp, "/")
			mc2 += "[] 'labelsKeys' STORE\n"
			mc2 += fmt.Sprintf("[ '%s' ] 'regExp' STORE\n", regExp)
		case *influxql.StringLiteral:
			labelsKey := showStatement.TagKeyExpr.String()
			labelsKey = strings.TrimLeft(labelsKey, "'")
			labelsKey = strings.TrimRight(labelsKey, "'")
			mc2 += fmt.Sprintf("[ '%s' ] 'labelsKeys' STORE\n", labelsKey)
			mc2 += "[] 'regExp' STORE\n"
		case *influxql.ListLiteral:
			listString := strings.TrimLeft(showStatement.TagKeyExpr.String(), "(")
			listString = strings.TrimRight(listString, ")")
			listItems := strings.Split(listString, ",")
			mc2 += fmt.Sprintf("[ '%s' ] 'labelsKeys' STORE\n", strings.Join(listItems, "' '"))
			mc2 += "[] 'regExp' STORE\n"
		}
		if showStatement.QueryType == ShowTagValues {
			mc2 += getShowTagValuesWarpScript(token, findSelector, labelsFilter, offset, limit, showStatement.Separator)
		} else {
			mc2 += getShowTagValuesCardinalityWarpScript(token, classnames, labelsFilter, offset, limit, showStatement.Separator)
		}
	default:
		return nil, fmt.Errorf("unvalid type")
	}

	// Show Measurement represents a FIND on series

	log.Debug(mc2)

	queryRes, err := warpServer.Query(mc2, txn)

	if err != nil {
		log.WithFields(log.Fields{
			"error": err.Error(),
			"proto": "influxQL",
		}).Error("Bad response from Egress")
		return nil, err
	}

	log.Debug(queryRes.Body)
	buffer, err := ioutil.ReadAll(queryRes.Body)
	if err != nil {
		wErr := queryRes.Header.Get("X-Warp10-Error-Message")
		if wErr == "" {
			dump, err := httputil.DumpResponse(queryRes, true)
			if err == nil {
				wErr = string(dump)
			} else {
				wErr = "Unparsable error"
			}
		}

		log.WithFields(log.Fields{
			"error": fmt.Errorf(wErr),
			"proto": "influxQL",
		}).Error("can't fully read Egress response: " + err.Error())
		return nil, fmt.Errorf(wErr)
	}
	// HACK : replace NaN values from Warp to 0
	s := strings.Replace(string(buffer), "NaN", "0", -1)
	buffer = []byte(s)

	responses := [][][]core.GeoTimeSeries{}

	err = json.Unmarshal(buffer, &responses)
	if err != nil {
		wErr := queryRes.Header.Get("X-Warp10-Error-Message")
		if wErr == "" {
			dump, err := httputil.DumpResponse(queryRes, true)
			if err == nil {
				wErr = string(dump)
			} else {
				wErr = "Unparsable error"
			}
		}

		log.WithFields(log.Fields{
			"error": fmt.Errorf(wErr),
			"proto": "influxQL",
		}).Error("Cannot unmarshal egress response: " + err.Error())
		return nil, fmt.Errorf(wErr)
	}

	if len(responses) == 0 {
		return &Result{StatementID: statementid, Series: make([]*Row, 0)}, nil
	}

	splitByNames := make(map[string][][]core.GeoTimeSeries)

	if showStatement.QueryType == ShowSeries || showStatement.QueryType == ShowMeasurements {
		splitByNames[showStatement.Name] = responses[0]
	} else {
		for _, seriesSet := range responses[0] {
			if _, ok := seriesSet[0].Labels[".name"]; ok {
				classname := seriesSet[0].Labels[".name"]
				existingResponse, contains := splitByNames[classname]

				if contains {
					splitByNames[classname] = append(existingResponse, seriesSet)
				} else {
					splitByNames[classname] = append(make([][]core.GeoTimeSeries, 0), seriesSet)
				}
			}
		}
	}

	series := make([]*Row, 0)

	for columnName, resps := range splitByNames {

		findAnswer := &Row{Name: columnName, Columns: showStatement.Columns, Values: make([][]interface{}, len(resps))}

		for index, seriesSet := range resps {

			findAnswer.Values[index] = make([]interface{}, 0)
			for _, rawSeries := range seriesSet {
				// Case of non series selector but a value, returns the value: useful for Find cardinality queries
				if rawSeries.Class == "" && len(rawSeries.Values) > 0 {
					findAnswer.Values[index] = append(findAnswer.Values[index], rawSeries.Values[0][1])
				} else {
					seriesShow := strings.Replace(rawSeries.Class, "{", ",", 1)
					seriesShow = strings.TrimSuffix(seriesShow, "}")
					findAnswer.Values[index] = append(findAnswer.Values[index], seriesShow)
				}
			}
		}
		series = append(series, findAnswer)
	}
	return &Result{StatementID: statementid, Series: series}, nil
}

// getShowSeriesWarpScript specific SHOW SERIES STATEMENT to extract series from WarpScript FIND
func getShowSeriesWarpScript(token, findString, labelsFilter string, offset, limit int, separator string) string {
	mc2 := fmt.Sprintf("[ '%s' '%s' %s ", token, findString, labelsFilter)
	// FIXME: controle FIND Limits here
	mc2 += ` 
		] FIND 
		<% 
			DROP 
	    `
	mc2 += fmt.Sprintf("DUP NAME '%s' ", separator)
	mc2 += "<% DUP '' == %> <% DROP 1 ->LIST %> <% SPLIT %> IFTE\n"
	mc2 += `
			<%
				DUP SIZE 1 >= 
			%>
			<%
				0 GET RENAME
			%>
			<%
				DROP
			%>
			IFTE
			TOSELECTOR 
		%> LMAP UNIQUE
		<%
			DROP
			1 ->LIST 
		%>
		LMAP
		` + fmt.Sprintf("[ %d %d ] SUBLIST", offset, limit)
	return mc2
}

// getShowMeasurementsWarpScript specific SHOW MEASUREMENTS STATEMENT to extract measurements from WarpScript FIND
func getShowMeasurementsWarpScript(token, findString, labelsFilter string, offset, limit int, separator string) string {
	// FIXME: controle FIND Limits here
	mc2 := fmt.Sprintf("[ '%s' '%s' %s ", token, findString, labelsFilter)

	mc2 += ` ] FIND 
		<% DROP NAME
		`
	mc2 += fmt.Sprintf(" '%s' ", separator)
	mc2 += "<% DUP '' == %> <% DROP 1 ->LIST %> <% SPLIT %> IFTE\n"
	mc2 += ` 0 GET %> LMAP UNIQUE
		<%
			DROP
			1 ->LIST 
		%>
		LMAP
		` + fmt.Sprintf("[ %d %d ] SUBLIST", offset, limit)

	return mc2
}

// getShowFieldKeysWarpScript specific SHOW FIELD KEYS STATEMENT to extract field keys from WarpScript FIND
func getShowFieldKeysWarpScript(token, findString, labelsFilter string, offset, limit int, separator string) string {
	// FIXME: controle FIND Limits here
	mc2 := fmt.Sprintf("[ '%s' '%s' %s ", token, findString, labelsFilter)
	mc2 += ` ] FIND 
		{} 'res' STORE
		<% `
	mc2 += fmt.Sprintf(" NAME '%s' ", separator)
	mc2 += "<% DUP '' == %> <% DROP 1 ->LIST %> <% SPLIT %> IFTE\n"
	mc2 += `<% 
						DUP SIZE 0 >
				%>
				<%
						0 REMOVE 'key' STORE '.' JOIN
						
				%>
				<%
						0 GET DUP 'key' STORE
				%>
				IFTE
				
				$res 
				<% 
					$key CONTAINSKEY
				%> 
				<%
					$key GET SWAP +
				%>
				<%
					DROP 1 ->LIST
				%>
				IFTE
				$res SWAP $key PUT
				'res' STORE
		%> 
		FOREACH
		[]
		$res
		<%
			UNIQUE SWAP 'name' STORE 
			<% 
				NEWGTS SWAP RENAME { '.name' $name } RELABEL NOW NaN NaN NaN 1 ADDVALUE 
				NEWGTS 'float' RENAME { '.name' $name } RELABEL NOW NaN NaN NaN 1 ADDVALUE 
				2 ->LIST + 
			%> FOREACH
		%>
		FOREACH 
		` + fmt.Sprintf("[ %d %d ] SUBLIST", offset, limit)

	return mc2
}

// getShowTagKeysWarpScript specific SHOW TAG KEYS STATEMENT to extract tag keys from WarpScript FIND
func getShowTagKeysWarpScript(token, findString, labelsFilter string, offset, limit int, selector string) string {
	// FIXME: controle FIND Limits here
	mc2 := fmt.Sprintf("[ '%s' '%s' %s ", token, findString, labelsFilter) +
		` ] FIND 
		{} 'res' STORE
                <%
`

	mc2 += fmt.Sprintf("DUP NAME '%s' ", selector)

	mc2 += "<% DUP '' == %> <% DROP 1 ->LIST %> <% SPLIT %> IFTE\n"
	mc2 += ` 0 GET 'key' STORE
					LABELS KEYLIST

					$res 
					<% 
							$key CONTAINSKEY
					%> 
					<%
							$key GET SWAP +
					%>
					<%
							DROP 1 ->LIST
					%>
					IFTE
					$res SWAP $key PUT
					'res' STORE
                %> 
                FOREACH
                []
                $res
                <%
					SWAP 'name' STORE 		
					FLATTEN UNIQUE 
					<% 
							NEWGTS SWAP RENAME { '.name' $name } RELABEL NOW NaN NaN NaN 1 ADDVALUE 
							1 ->LIST + 
					%> FOREACH
                %>
                FOREACH 
		` + fmt.Sprintf("[ %d %d ] SUBLIST", offset, limit)

	return mc2
}

func getShowTagValuesCardinalityWarpScript(token string, selectors []string, labelsFilter string, offset, limit int, separator string) string {
	// TODO: controle FIND Limits here
	mc2 := "[\n"
	for _, selector := range selectors {
		mc2 += fmt.Sprintf("'%s' 'selector' STORE\n", selector)
		mc2 += fmt.Sprintf("[ '%s' '~%s.*' %s ", token, selector, labelsFilter) +
			` ] FINDSTATS 
		'per.label.value.estimate' GET
		<% 
			$regExp SIZE 0 >
		%>
		<% 
			{} 'regExpLabels' STORE
			<%
					'valueRegExp' STORE
				DUP 'keyRegExp' STORE
				$regExp 0 GET
				MATCH
				<%
					SIZE 0 >
				%>
				<%
					$regExpLabels
					$valueRegExp
					$keyRegExp
					PUT
					'regExpLabels' STORE
				%>
				IFT
			%>
			FOREACH
			$regExpLabels
		%>
		<%
			$labelsKeys SUBMAP
		%>
		IFTE
		[]
		SWAP
		<% 
			SWAP DROP
			'value' STORE
			NEWGTS { '.name' $selector } RELABEL NOW NaN NaN NaN $value ADDVALUE 
			+
		%>
		FOREACH
		[ SWAP bucketizer.last NOW 0 1 ] BUCKETIZE
		[ SWAP [ '.name' ] reducer.sum ] REDUCE 
		NONEMPTY
		`
	}
	mc2 += "]"
	return mc2
}

// getShowTagValuesWarpScript specific SHOW TAG VALUES STATEMENT to extract tag values from WarpScript FIND
func getShowTagValuesWarpScript(token, findString, labelsFilter string, offset, limit int, selector string) string {
	// FIXME: controle FIND Limits here
	mc2 := fmt.Sprintf("[ '%s' '%s' %s ", token, findString, labelsFilter) +
		` ] FIND 
		{} 'res' STORE
                <%
		`

	mc2 += fmt.Sprintf("DUP NAME '%s' ", selector)
	mc2 += "<% DUP '' == %> <% DROP 1 ->LIST %> <% SPLIT %> IFTE\n"
	mc2 += `0 GET 'key' STORE
                                LABELS
                                <% 
                                    $regExp SIZE 0 >
                                %>
                                <% 
                                    {} 'regExpLabels' STORE
                                    <%
                                        'valueRegExp' STORE
                                        DUP 'keyRegExp' STORE
                                        $regExp 0 GET 
                                        MATCH
                                        <%
                                           SIZE 0 >
                                        %>
                                        <%
                                            $regExpLabels
                                            $valueRegExp
                                            $keyRegExp
                                            PUT
                                            'regExpLabels' STORE
                                        %>
                                        IFT
                                    %>
                                    FOREACH
                                    $regExpLabels
                                %>
                                <%
                                    $labelsKeys SUBMAP
                                %>
                                IFTE
                                1 ->LIST
                                $res 
                                <% 
                                        $key CONTAINSKEY 
                                %> 
                                <%
                                        $key GET SWAP  +
                                %>
                                <%
                                        DROP 1 ->LIST
                                %>
                                IFTE
                                $res SWAP $key PUT
                                'res' STORE
                %> 
                FOREACH
            
                []
                $res
                <%
                    SWAP 'name' STORE 
                    <% DROP ->JSON %> LMAP FLATTEN UNIQUE
                    <% 
                        JSON-> 0 GET
                        <%
                            DUP KEYLIST SIZE 0 ==
                        %>
                        <%
                            DROP
                            CONTINUE
                        %>
                        <%
                            DUP
                            KEYLIST 0 GET 
                            NEWGTS SWAP RENAME { '.name' $name } RELABEL NOW NaN NaN NaN 1 ADDVALUE 
                            SWAP
                            VALUELIST 0 GET 
                            NEWGTS SWAP RENAME { '.name' $name } RELABEL NOW NaN NaN NaN 1 ADDVALUE 
                            2 ->LIST +
                        %>
                        IFTE
                    %> FOREACH
                %>
                FOREACH 
		` + fmt.Sprintf("[ %d %d ] SUBLIST", offset, limit)

	return mc2
}

// parseInfluxSelect function used to Create Warp 10 query based on an Influx SELECT statement
func parseInfluxSelect(statement *influxql.SelectStatement, statementid int, txn string, token string, timePrecision string) (*Result, error) {

	warpServer := core.NewWarpServer(viper.GetString("warp_endpoint"), "influxql")

	keepTopLabels := make([]string, 0)

	p := &InfluxParser{Classnames: make([]string, 0), Token: token, End: "$end", BucketTime: "0", BucketCount: "1", Separator: ".", KeepTopLabels: keepTopLabels}

	mc2, respResult, _, _, err := p.getSelectStatementScript(statement, warpServer, txn, 0, statementid)

	if respResult != nil {
		return respResult, nil
	}

	if err != nil {
		return nil, err
	}

	mc2 = "'" + token + "' " + `
	AUTHENTICATE
	// keys of STACKATTRIBUTE can be found here: https://github.com/cityzendata/warp10-platform/blob/master/warp10/src/main/java/io/warp10/script/WarpScriptStack.java
	'stack.maxops.hard' STACKATTRIBUTE DUP <% ISNULL ! %> <% MAXOPS %> <% DROP %> IFTE
	'fetch.limit.hard' STACKATTRIBUTE DUP <% ISNULL ! %> <% LIMIT %> <% DROP %> IFTE
	'gts.limit.hard' STACKATTRIBUTE DUP <% ISNULL ! %> <% MAXGTS %> <% DROP %> IFTE
	` + mc2

	queryRes, err := warpServer.Query(mc2, txn)

	if err != nil {
		log.WithFields(log.Fields{
			"error": err.Error(),
			"proto": "influxQL",
		}).Error("Bad response from Egress")
		return nil, err
	}
	buffer, err := ioutil.ReadAll(queryRes.Body)
	if err != nil {
		log.WithFields(log.Fields{
			"error": err.Error(),
			"proto": "influxQL",
		}).Error("can't fully read Egress response")
		return nil, err
	}

	// HACK : replace NaN values from Warp to 0
	s := strings.Replace(string(buffer), "NaN", "0", -1)
	buffer = []byte(s)

	if queryRes.StatusCode != http.StatusOK {
		return &Result{Err: "Invalid generated query: returned error is - " + string(buffer)}, nil
	}

	responses := [][]core.GeoTimeSeries{}
	err = json.Unmarshal(buffer, &responses)
	if err != nil {
		log.WithFields(log.Fields{
			"error": err.Error(),
			"proto": "influxQL",
		}).Error("Cannot unmarshal egress response")
		return nil, err
	}
	// Since it's a range_query, we can enforce the matrix resultType
	influxResponse := &Result{StatementID: statementid, Series: make([]*Row, 0)}

	for _, series := range responses {
		if len(series) == 0 {
			continue
		}

		if p.HasGroupBy {
			_, tags := statement.Dimensions.Normalize()
			resp, err := splitPerTagsResult(series, statementid, p.TimeColumn, p.Measurement, timePrecision, statement.Location, tags, statement.Fields.AliasNames())
			if err != nil {
				return nil, err
			}
			influxResponse.Series = append(influxResponse.Series, resp...)
		} else {
			resp, err := warpToInfluxResponse(series, statementid, p.TimeColumn, p.Measurement, timePrecision, statement.Location)
			if err != nil {
				return nil, err
			}
			influxResponse.Series = append(influxResponse.Series, resp)
		}
	}

	return influxResponse, nil
}

// invalidNumberOfArgs: generate generic invalidNumberOfArgs strings
func invalidNumberOfArgs(function string, expected int, got int) string {
	return fmt.Sprintf("invalid number of arguments for %s, expected %d, got %d", function, expected, got)
}

// containsString does an array contains a single string
func containsString(array []string, val string) bool {
	for _, item := range array {
		if item == val {
			return true
		}
	}
	return false
}

// splitPerTagsResult Split series result per tags
func splitPerTagsResult(seriesSet []core.GeoTimeSeries, statementid int, timeColumn *InfluxTimeColumn, influxMeasurement *InfluxMeasurement, timePrecision string, location *time.Location, tags []string, selectedField []string) ([]*Row, error) {
	resp := make([]*Row, 0)
	splitSeries := make(map[string][]core.GeoTimeSeries)

	for _, series := range seriesSet {
		stringSet := fmt.Sprintf("%s.%s", series.Class, series.Attrs)

		seriesSet, containsSeries := splitSeries[stringSet]

		for tag := range series.Attrs {
			delete(series.Labels, tag)
		}

		if containsSeries {
			splitSeries[stringSet] = append(seriesSet, series)
		} else {
			groupBySeries := make([]core.GeoTimeSeries, 1)
			groupBySeries[0] = series
			splitSeries[stringSet] = groupBySeries
		}
	}

	for _, series := range splitSeries {

		if len(series) > 0 {
			currrenRow, err := warpToInfluxResponse(series, statementid, timeColumn, influxMeasurement, timePrecision, location)

			if err != nil {
				return nil, err
			}

			currrenRow.Tags = series[0].Attrs

			resp = append(resp, currrenRow)
		}
	}

	return resp, nil
}

// warpToInfluxResponse Translate Warp10 responses to a native Influx one
func warpToInfluxResponse(seriesSet []core.GeoTimeSeries, statementid int, timeColumn *InfluxTimeColumn, influxMeasurement *InfluxMeasurement, timePrecision string, location *time.Location) (*Row, error) {

	seriesValues := make(map[float64]map[string]map[string]map[int]interface{})

	columns := make(map[string]int, 0)
	columns[timeColumn.TimeAlias] = 0
	columnsIndex := 1

	resp := &Row{}

	// Parse all result series set
	for index, series := range seriesSet {

		for key, value := range series.Labels {
			_, ok := columns[key]

			// Load native Influx Name
			if key == ".INFLUXQL_COLUMN_NAME" {
				_, ok := columns[value]
				if !ok {
					columns[value] = columnsIndex
					columnsIndex++
				}
				continue
			}
			if !ok {
				columns[key] = columnsIndex
				columnsIndex++
			}
		}

		if index == 0 {
			if !influxMeasurement.StripMeasurement {
				resp.Name = series.Class
			}
			if influxMeasurement.EmitName != "" {
				resp.Name = influxMeasurement.EmitName
			}
		}

		fieldName, ok := series.Labels[".INFLUXQL_COLUMN_NAME"]
		if !ok {
			fieldName = "value"
		} else {
			delete(series.Labels, ".INFLUXQL_COLUMN_NAME")
		}
		id := fmt.Sprint(series.Labels)

		// Parse series values
		for _, seriesInternalValues := range series.Values {
			val := seriesInternalValues[1]
			tick := seriesInternalValues[0].(float64)

			seriesMap, containsTick := seriesValues[tick]
			if !containsTick {
				seriesMap = make(map[string]map[string]map[int]interface{})
			}

			valueInterface, containsSeries := seriesMap[id]

			if !containsSeries {
				valueInterface = make(map[string]map[int]interface{}, len(series.Labels))

				for key, value := range series.Labels {
					setValues := make(map[int]interface{})
					setValues[0] = value
					valueInterface[key] = setValues
				}
				setValuesField := make(map[int]interface{})
				setValuesField[0] = val
				valueInterface[fieldName] = setValuesField
			}

			if containsSeries {
				valueSeries, containsKey := valueInterface[fieldName]

				nextKey := 0
				if containsKey {
					nextKey := 1
					for key := range valueSeries {
						if key > nextKey {
							nextKey = key
						}
					}

					nextKey++
				}

				setValuesField := make(map[int]interface{})
				setValuesField[nextKey] = val
				valueInterface[fieldName] = setValuesField

				for key, value := range series.Labels {
					setValues := make(map[int]interface{})
					setValues[nextKey] = value
					valueInterface[key] = setValues
				}
			}
			seriesMap[id] = valueInterface
			seriesValues[tick] = seriesMap
		}
	}

	sortedColumns := make(map[string]int, len(columns))

	// Sort Influx Columns by series meta
	keys := make([]string, 0, len(columns))
	for key := range columns {
		keys = append(keys, key)
	}
	sort.Strings(keys)

	columnsIndex = 0

	if !timeColumn.OmitColumnTime {
		sortedColumns[timeColumn.TimeAlias] = 0
		columnsIndex = 1
	}
	for _, key := range keys {
		if key == timeColumn.TimeAlias {
			continue
		}
		sortedColumns[key] = columnsIndex
		columnsIndex++
	}

	columnsStringArray := make([]string, len(sortedColumns))

	for key, columnIndex := range sortedColumns {
		columnsStringArray[columnIndex] = key
	}
	resp.Columns = columnsStringArray

	sortedTicks := make([]float64, len(seriesValues))

	tickIndex := 0
	for tick := range seriesValues {
		sortedTicks[tickIndex] = tick
		tickIndex++
	}

	if timeColumn.ReverseSort {
		sort.Sort(sort.Reverse(sort.Float64Slice(sortedTicks)))
	} else {
		sort.Float64s(sortedTicks)
	}

	valuesResult := make([][]interface{}, 0)
	for _, tick := range sortedTicks {

		mapSeries := seriesValues[tick]
		for _, mapValue := range mapSeries {

			createdInterface := make(map[int][]interface{})
			valueInterfaceToSet := make([]interface{}, len(sortedColumns))
			if !timeColumn.OmitColumnTime {
				valueInterfaceToSet[0] = getFormatedTime(tick/1000, timePrecision, location)
			}
			createdInterface[0] = valueInterfaceToSet
			for key, columnIndex := range sortedColumns {

				if valueLabels, ok := mapValue[key]; ok {
					switch key {
					case timeColumn.TimeAlias:
						continue
					default:
					}

					for key, value := range valueLabels {
						valueInterfaceToSet, containsValueInterface := createdInterface[key]
						if !containsValueInterface {
							valueInterfaceToSet = make([]interface{}, len(sortedColumns))
							if !timeColumn.OmitColumnTime {
								valueInterfaceToSet[0] = getFormatedTime(tick/1000, timePrecision, location)
							}
						}
						valueInterfaceToSet[columnIndex] = value
						createdInterface[key] = valueInterfaceToSet
					}
				}
			}

			for _, valueInterface := range createdInterface {
				valuesResult = append(valuesResult, valueInterface)
			}
		}
	}

	resp.Values = valuesResult
	return resp, nil
}

// getFormatedTime Format Time results based on timePrecision field
func getFormatedTime(tick float64, timePrecision string, location *time.Location) interface{} {
	if timePrecision != "rfc3339" && location.String() != "UTC" {
		return time.Unix(0, int64(tick)*int64(time.Millisecond)).UTC().In(location)
	}

	switch timePrecision {
	case "rfc3339":
		unixTimeUTC := time.Unix(0, int64(tick)*int64(time.Millisecond)).UTC()
		return unixTimeUTC.Format(time.RFC3339)
	case "ms":
		return tick
	case "s":
		return tick / 1000
	case "us", "timestamp":
		return tick * 1000
	case "ns":
		return tick * 1000000
	default:
		if location.String() != "UTC" {
			return time.Unix(0, int64(tick)*int64(time.Millisecond)).UTC().In(location)
		}
		return tick
	}
}

func defaultCreateDB(statementid int) Result {
	return Result{StatementID: statementid}
}

// Default Influx SHOW DATABASES return
// https://github.com/influxdata/influxdb-java/blob/90395a8609798bb6fc7f149716ab12d38fc0c96f/src/main/java/org/influxdb/impl/InfluxDBImpl.java#L718
/* Returned result is:
databases = `{
	"results": [
		{
			"series": [
				{
					"name": "databases",
					"columns": [
						"name"
					],
					"values": [
						[
							"metrics"
						]
					]
				}
			]
		}
	]
}`
*/
func defaultShowDatabasesStatement(statementid int) Result {
	series := make([]*Row, 1)
	defaultValue := make([][]interface{}, 1)
	interfaceValues := make([]interface{}, 1)
	interfaceValues[0] = "metrics"
	defaultValue[0] = interfaceValues

	defaultRow := &Row{Name: "databases", Columns: []string{"name"}, Values: defaultValue}
	series[0] = defaultRow
	return Result{StatementID: statementid, Series: series}
}

// Default Influx Retention return
/* Returned result is:
retention = `{
	"results": [
		{
			"statement_id": 0,
			"series": [
				{
					"columns": [
						"name",
						"duration",
						"shardGroupDuration",
						"replicaN",
						"default"
					],
					"values": [
						[
							"autogen",
							"0s",
							"0s",
							1,
							true
						]
					]
				}
			]
		}
	]
}`
*/
func defaultShowRetentionPoliciesStatement(statementid int) Result {
	series := make([]*Row, 1)
	defaultValue := make([][]interface{}, 1)
	interfaceValues := make([]interface{}, 5)
	interfaceValues[0] = "autogen"
	interfaceValues[1] = "0s"
	interfaceValues[2] = "0s"
	interfaceValues[3] = 1
	interfaceValues[4] = true
	defaultValue[0] = interfaceValues

	defaultRow := &Row{Columns: []string{"name", "duration", "shardGroupDuration", "replicaN", "default"}, Values: defaultValue}
	series[0] = defaultRow
	return Result{StatementID: statementid, Series: series}
}
