package influxdb

import (
	"bytes"
	"fmt"
	"strings"

	"github.com/influxdata/influxql"
	log "github.com/sirupsen/logrus"
)

// ExprReturn represents the type returned by the expr.
type ExprReturn int

const (
	// Unknown type.
	Unknown ExprReturn = 0
	// Scalar type.
	Scalar ExprReturn = 1
	// SeriesSet type.
	SeriesSet ExprReturn = 2
)

// BinaryOperationType represents the operation type
type BinaryOperationType int

const (
	// UnknownBinary type.
	UnknownBinary BinaryOperationType = 0
	// ScalarToScalar Operation.
	ScalarToScalar BinaryOperationType = 1
	// ScalarToSeries and series type.
	ScalarToSeries BinaryOperationType = 2
	// SeriesToScalar and series type.
	SeriesToScalar BinaryOperationType = 3
	// SeriesToSeries type.
	SeriesToSeries BinaryOperationType = 4
)

type binaryOperationType struct {
}

// parseExpr Parse native Influx Queries expression
func (p *InfluxParser) parseExpr(getExpr influxql.Expr, level int, selectors []string, where [][]*WhereCond) (string, ExprReturn, error) {
	exprType := Unknown
	mc2 := ""
	// Return the function name or variable name, if available.
	switch expr := getExpr.(type) {
	case *influxql.Call:
		call, exprTypeCall, err := p.parseCall(expr, selectors, where)
		if err != nil {
			return "", Unknown, err
		}
		exprType = exprTypeCall
		mc2 += call
	case *influxql.BinaryExpr:
		leftString, leftType, leftError := p.parseExpr(expr.LHS, level, selectors, where)
		if leftError != nil {
			return "", Unknown, leftError
		}
		rightString, rightType, rightError := p.parseExpr(expr.RHS, level, selectors, where)
		if rightError != nil {
			return "", Unknown, rightError
		}

		mc2 += leftString + fmt.Sprintf(" 'left-%d' STORE\n", level)
		mc2 += rightString + fmt.Sprintf(" 'right-%d' STORE\n", level)

		binaryType := getBinaryOperationType(leftType, rightType)

		opMc2, opErr := p.parseMathOperation(binaryType, expr.Op, level)
		if opErr != nil {
			return "", Unknown, rightError
		}
		mc2 += opMc2
		if leftType == Scalar && rightType == Scalar {
			exprType = Scalar
		} else {
			exprType = SeriesSet
		}
	case *influxql.ParenExpr:
		return p.parseExpr(expr.Expr, level+1, selectors, where)
	case *influxql.VarRef:
		p.HasWildCard = false
		fetch, exprTypeCall, err := p.parseFetch(expr.Val, expr.Type, selectors, where)
		if err != nil {
			return "", Unknown, err
		}
		exprType = exprTypeCall
		mc2 += fetch
	case *influxql.IntegerLiteral:
		return fmt.Sprintf(" %d ", expr.Val), Scalar, nil
	default:
		log.Warnf("Parse expr %s : %T", expr.String(), expr)
	}

	return mc2, exprType, nil

}

// Check a Binary Operation to get it's type
func getBinaryOperationType(leftType, rightType ExprReturn) BinaryOperationType {
	if leftType == Scalar && rightType == Scalar {
		return ScalarToScalar
	}
	if leftType == Scalar && rightType == SeriesSet {
		return ScalarToSeries
	}
	if leftType == SeriesSet && rightType == Scalar {
		return SeriesToScalar
	}
	if leftType == SeriesSet && rightType == SeriesSet {
		return SeriesToSeries
	}
	return UnknownBinary
}

// Parse Influx Operator
// https://docs.influxdata.com/influxdb/v1.7/query_language/math_operators/
func (p *InfluxParser) parseMathOperation(binaryOperation BinaryOperationType, op influxql.Token, level int) (string, error) {

	switch op {
	case influxql.ADD:
		switch binaryOperation {
		case ScalarToScalar:
			return fmt.Sprintf("$left-%d $right-%d + \n", level, level), nil
		case ScalarToSeries:
			return fmt.Sprintf("[ $right-%d $left-%d TODOUBLE mapper.add 0 0 0 ] MAP\n", level, level), nil
		case SeriesToScalar:
			return fmt.Sprintf("[ $left-%d $right-%d TODOUBLE mapper.add 0 0 0 ] MAP\n", level, level), nil
		case SeriesToSeries:
			return fmt.Sprintf("[ $right-%d $left-%d NULL op.add ] APPLY NONEMPTY\n", level, level), nil
		default:
			return "", fmt.Errorf("Unvalid operation types")
		}

	case influxql.SUB:
		switch binaryOperation {
		case ScalarToScalar:
			return fmt.Sprintf("$left-%d $right-%d - \n", level, level), nil
		case ScalarToSeries:
			return fmt.Sprintf("[ [ $right-%d -1.0 mapper.mul 0 0 0 ] MAP $left-%d TODOUBLE mapper.add 0 0 0 ] MAP\n", level, level), nil
		case SeriesToScalar:
			return fmt.Sprintf("[ $left-%d $right-%d TODOUBLE -1 * mapper.add 0 0 0 ] MAP\n", level, level), nil
		case SeriesToSeries:
			return fmt.Sprintf("[ $right-%d ", level) + "'%2B.tosub' RENAME " + fmt.Sprintf("$left-%d ", level) + " NULL op.sub ] APPLY NONEMPTY\n", nil
		default:
			return "", fmt.Errorf("Unvalid operation types")
		}

	case influxql.MUL:
		switch binaryOperation {
		case ScalarToScalar:
			return fmt.Sprintf("$left-%d $right-%d * \n", level, level), nil
		case ScalarToSeries:
			return fmt.Sprintf("[ $right-%d $left-%d TODOUBLE mapper.mul 0 0 0 ] MAP\n", level, level), nil
		case SeriesToScalar:
			return fmt.Sprintf("[ $left-%d $right-%d TODOUBLE mapper.mul 0 0 0 ] MAP\n", level, level), nil
		case SeriesToSeries:
			return fmt.Sprintf("[ $right-%d $left-%d NULL op.mul ] APPLY NONEMPTY\n", level, level), nil
		default:
			return "", fmt.Errorf("Unvalid operation types")
		}

	case influxql.DIV:
		switch binaryOperation {
		case ScalarToScalar:
			return fmt.Sprintf("$left-%d $right-%d + \n", level, level), nil
		case ScalarToSeries:
			return fmt.Sprintf("[ $right-%d ", level) + " <% DROP " + fmt.Sprintf("$left-%d", level) + " TODOUBLE / %> LMAP\n", nil
		case SeriesToScalar:
			return fmt.Sprintf("[ $left-%d 1 $right-%d TODOUBLE / mapper.mul 0 0 0 ] MAP\n", level, level), nil
		case SeriesToSeries:
			return fmt.Sprintf("[ $right-%d ", level) + "'%2B.todiv' RENAME " + fmt.Sprintf("$left-%d ", level) + " NULL op.div ] APPLY NONEMPTY\n", nil
		default:
			return "", fmt.Errorf("Unvalid operation types")
		}

	case influxql.MOD:
		switch binaryOperation {
		case ScalarToScalar:
			return fmt.Sprintf("$left-%d $right-%d ", level, level) + "% \n", nil
		case ScalarToSeries:
			return fmt.Sprintf("[ $right-%d $left-%d mapper.mod 0 0 0 ] MAP\n", level, level), nil
		case SeriesToScalar:
			return fmt.Sprintf("[ $left-%d $right-%d mapper.mod 0 0 0 ] MAP\n", level, level), nil
		case SeriesToSeries:
			return "", fmt.Errorf("Modulo is currently not supported between two GTS sets") // FIXME
		default:
			return "", fmt.Errorf("Unvalid operation types")
		}

	case influxql.BITWISE_AND:
		switch binaryOperation {
		case ScalarToScalar:
			return fmt.Sprintf("$left-%d $right-%d & \n", level, level), nil
		case ScalarToSeries:
			return fmt.Sprintf("[ $right-%d NONEMPTY mapper.tolong 0 0 0 ] MAP ", level) + " <% DROP " + fmt.Sprintf("$left-%d", level) + " TOLONG & %> LMAP\n", nil
		case SeriesToScalar:
			return fmt.Sprintf("[ $left-%d NONEMPTY mapper.tolong 0 0 0 ] MAP ", level) + " <% DROP " + fmt.Sprintf("$right-%d", level) + " TOLONG & %> LMAP\n", nil
		case SeriesToSeries:
			mc2 := fmt.Sprintf("[ $right-%d  NONEMPTY mapper.tolong 0 0 0 ] MAP NULL PARTITION 'leftKeys' STORE\n", level)
			mc2 += fmt.Sprintf("[ $left-%d  NONEMPTY mapper.tolong 0 0 0 ] MAP NULL PARTITION \n", level)
			mc2 += getBitwiseApply("&")
			return mc2, nil
		default:
			return "", fmt.Errorf("Unvalid operation types")
		}

	case influxql.BITWISE_OR:
		switch binaryOperation {
		case ScalarToScalar:
			return fmt.Sprintf("$left-%d $right-%d | \n", level, level), nil
		case ScalarToSeries:
			return fmt.Sprintf("[ $right-%d NONEMPTY mapper.tolong 0 0 0 ] MAP", level) + " <% DROP " + fmt.Sprintf("$left-%d", level) + " TOLONG | %> LMAP\n", nil
		case SeriesToScalar:
			return fmt.Sprintf("[ $left-%d NONEMPTY mapper.tolong 0 0 0 ] MAP ", level) + " <% DROP " + fmt.Sprintf("$right-%d", level) + " TOLONG | %> LMAP\n", nil
		case SeriesToSeries:
			mc2 := fmt.Sprintf("[ $right-%d  NONEMPTY mapper.tolong 0 0 0 ] MAP NULL PARTITION 'leftKeys' STORE\n", level)
			mc2 += fmt.Sprintf("[ $left-%d  NONEMPTY mapper.tolong 0 0 0 ] MAP NULL PARTITION \n", level)
			mc2 += getBitwiseApply("|")
			return mc2, nil
		default:
			return "", fmt.Errorf("Unvalid operation types")
		}

	case influxql.BITWISE_XOR:
		switch binaryOperation {
		case ScalarToScalar:
			return fmt.Sprintf("$left-%d $right-%d ^ \n", level, level), nil
		case ScalarToSeries:
			return fmt.Sprintf("[ $right-%d NONEMPTY mapper.tolong 0 0 0 ] MAP ", level) + " <% DROP " + fmt.Sprintf("$left-%d", level) + " TOLONG ^ %> LMAP\n", nil
		case SeriesToScalar:
			return fmt.Sprintf("[ $left-%d NONEMPTY mapper.tolong 0 0 0 ] MAP ", level) + " <% DROP " + fmt.Sprintf("$right-%d", level) + " TOLONG ^ %> LMAP\n", nil
		case SeriesToSeries:
			mc2 := fmt.Sprintf("[ $right-%d  NONEMPTY mapper.tolong 0 0 0 ] MAP NULL PARTITION 'leftKeys' STORE\n", level)
			mc2 += fmt.Sprintf("[ $left-%d  NONEMPTY mapper.tolong 0 0 0 ] MAP NULL PARTITION \n", level)
			mc2 += getBitwiseApply("^")
			return mc2, nil
		default:
			return "", fmt.Errorf("Unvalid operation types")
		}

	case influxql.AND:
		switch binaryOperation {
		case ScalarToScalar:
			return fmt.Sprintf("$left-%d $right-%d AND \n", level, level), nil
		case ScalarToSeries:
			return fmt.Sprintf("[ $right-%d $left-%d TOBOOLEAN mapper.and 0 0 0 ] MAP\n", level, level), nil
		case SeriesToScalar:
			return fmt.Sprintf("[ $left-%d $right-%d TOBOOLEAN mapper.and 0 0 0 ] MAP\n", level, level), nil
		case SeriesToSeries:
			return fmt.Sprintf("[ $right-%d $left-%d NULL op.and ] APPLY NONEMPTY\n", level, level), nil
		default:
			return "", fmt.Errorf("Unvalid operation types")
		}

	case influxql.OR:
		switch binaryOperation {
		case ScalarToScalar:
			return fmt.Sprintf("$left-%d $right-%d OR \n", level, level), nil
		case ScalarToSeries:
			return fmt.Sprintf("[ $right-%d $left-%d TOBOOLEAN mapper.or 0 0 0 ] MAP\n", level, level), nil
		case SeriesToScalar:
			return fmt.Sprintf("[ $left-%d $right-%d TOBOOLEAN mapper.or 0 0 0 ] MAP\n", level, level), nil
		case SeriesToSeries:
			return fmt.Sprintf("[ $right-%d $left-%d NULL op.or ] APPLY NONEMPTY\n", level, level), nil
		default:
			return "", fmt.Errorf("Unvalid operation types")
		}

	case influxql.EQ:
		switch binaryOperation {
		case ScalarToScalar:
			return fmt.Sprintf("$left-%d $right-%d == \n", level, level), nil
		case ScalarToSeries:
			return fmt.Sprintf("[ $right-%d $left-%d mapper.eq 0 0 0 ] MAP\n", level, level), nil
		case SeriesToScalar:
			return fmt.Sprintf("[ $left-%d $right-%d mapper.eq 0 0 0 ] MAP\n", level, level), nil
		case SeriesToSeries:
			return fmt.Sprintf("[ $right-%d $left-%d NULL op.eq ] APPLY NONEMPTY\n", level, level), nil
		default:
			return "", fmt.Errorf("Unvalid operation types")
		}

	case influxql.NEQ:
		switch binaryOperation {
		case ScalarToScalar:
			return fmt.Sprintf("$left-%d $right-%d ! \n", level, level), nil
		case ScalarToSeries:
			return fmt.Sprintf("[ $right-%d $left-%d TOBOOLEAN mapper.ne 0 0 0 ] MAP\n", level, level), nil
		case SeriesToScalar:
			return fmt.Sprintf("[ $left-%d $right-%d TOBOOLEAN mapper.ne 0 0 0 ] MAP\n", level, level), nil
		case SeriesToSeries:
			return fmt.Sprintf("[ $right-%d $left-%d NULL op.ne ] APPLY NONEMPTY\n", level, level), nil
		default:
			return "", fmt.Errorf("Unvalid operation types")
		}

	case influxql.LT:
		switch binaryOperation {
		case ScalarToScalar:
			return fmt.Sprintf("$left-%d $right-%d < \n", level, level), nil
		case ScalarToSeries:
			return fmt.Sprintf("[ $right-%d $left-%d TOBOOLEAN mapper.lt 0 0 0 ] MAP\n", level, level), nil
		case SeriesToScalar:
			return fmt.Sprintf("[ $left-%d $right-%d TOBOOLEAN mapper.lt 0 0 0 ] MAP\n", level, level), nil
		case SeriesToSeries:
			return fmt.Sprintf("[ $right-%d $left-%d NULL op.lt ] APPLY NONEMPTY\n", level, level), nil
		default:
			return "", fmt.Errorf("Unvalid operation types")
		}

	case influxql.LTE:
		switch binaryOperation {
		case ScalarToScalar:
			return fmt.Sprintf("$left-%d $right-%d <= \n", level, level), nil
		case ScalarToSeries:
			return fmt.Sprintf("[ $right-%d $left-%d TOBOOLEAN mapper.le 0 0 0 ] MAP\n", level, level), nil
		case SeriesToScalar:
			return fmt.Sprintf("[ $left-%d $right-%d TOBOOLEAN mapper.le 0 0 0 ] MAP\n", level, level), nil
		case SeriesToSeries:
			return fmt.Sprintf("[ $right-%d $left-%d NULL op.le ] APPLY NONEMPTY\n", level, level), nil
		default:
			return "", fmt.Errorf("Unvalid operation types")
		}

	case influxql.GT:
		switch binaryOperation {
		case ScalarToScalar:
			return fmt.Sprintf("$left-%d $right-%d > \n", level, level), nil
		case ScalarToSeries:
			return fmt.Sprintf("[ $right-%d $left-%d TOBOOLEAN mapper.gt 0 0 0 ] MAP\n", level, level), nil
		case SeriesToScalar:
			return fmt.Sprintf("[ $left-%d $right-%d TOBOOLEAN mapper.gt 0 0 0 ] MAP\n", level, level), nil
		case SeriesToSeries:
			return fmt.Sprintf("[ $right-%d $left-%d NULL op.gt ] APPLY NONEMPTY\n", level, level), nil
		default:
			return "", fmt.Errorf("Unvalid operation types")
		}

	case influxql.GTE:
		switch binaryOperation {
		case ScalarToScalar:
			return fmt.Sprintf("$left-%d $right-%d >= \n", level, level), nil
		case ScalarToSeries:
			return fmt.Sprintf("[ $right-%d $left-%d TOBOOLEAN mapper.ge 0 0 0 ] MAP\n", level, level), nil
		case SeriesToScalar:
			return fmt.Sprintf("[ $left-%d $right-%d TOBOOLEAN mapper.ge 0 0 0 ] MAP\n", level, level), nil
		case SeriesToSeries:
			return fmt.Sprintf("[ $right-%d $left-%d NULL op.ge ] APPLY NONEMPTY\n", level, level), nil
		default:
			return "", fmt.Errorf("Unvalid operation types")
		}
	default:
		log.Warnf("parseMathOperation %s %T", op.String(), op)
		return "", fmt.Errorf("Unsupported %s operator", op.String())
	}
}

func getBitwiseApply(operator string) string {
	mc2 :=
		`
			[] SWAP <%
			'seriesSet' STORE
			'key' STORE
			
			$leftKeys $key GET 
			<%
			  DUP ISNULL
			%>
			<%
			  DROP []
			%>
			IFT
			'secondSet' STORE
			$seriesSet
			$secondSet
			APPEND 
			<% DUP SIZE 1 > %>
			<%
				0 REMOVE 
				DUP
				CLONEEMPTY DUP NAME 'cName' STORE
				DUP LABELS 'cLabels' STORE
				ATTRIBUTES 'cAttributes' STORE
				SWAP
				<%
				`
	mc2 += operator
	mc2 += `
				%>
				FOREACH
				$cName RENAME
				$cLabels RELABEL
				$cAttributes SETATTRIBUTES
				+
			%>
			<%
			   DROP CONTINUE
			%>
			IFTE
		 %>
		 FOREACH
	`

	return mc2

}

// Generate WarpScript Fetch
func (p *InfluxParser) parseFetch(name string, fetchType influxql.DataType, selectors []string, where [][]*WhereCond) (string, ExprReturn, error) {

	fetchString := "FETCH"

	switch fetchType {
	case influxql.Float:
		fetchString = "FETCHDOUBLE"
	case influxql.Integer:
		fetchString = "FETCHLONG"
	case influxql.String:
		fetchString = "FETCHSTRING"
	case influxql.Boolean:
		fetchString = "FETCHBOOLEAN"
	}
	valueName := name
	if name == "*" {
		valueName = ".*"
	}

	mc2 := ""
	classname := fmt.Sprintf("~(%s)", strings.Join(p.Classnames, "|"))

	hasFilter := false

	selectFields := make([]string, 0)

	var computeTags bytes.Buffer

	for i, wheres := range where {

		if len(wheres) > 0 {
			computeTags.WriteString(" $set \n")

			filterTags := make([]string, 0)
			for _, whereCond := range wheres {
				if whereCond.isTag {
					op := ""
					value := whereCond.value.Val
					if strings.HasPrefix(whereCond.value.Val, "'") {
						value = strings.Trim(whereCond.value.Val, "'")
					} else if strings.HasPrefix(whereCond.value.Val, "\"") {
						value = strings.Trim(whereCond.value.Val, "\"")
					}

					if whereCond.op == influxql.REGEX || whereCond.op == influxql.EQREGEX {
						op = "~"
					} else if whereCond.op == influxql.NEQ {
						op = "~"
						value = fmt.Sprintf("(?!%s).*?", whereCond.value)
					} else if whereCond.op == influxql.NEQREGEX {
						op = "~"
						value = fmt.Sprintf("(?!%s)?", whereCond.value)
					}
					filterTags = append(filterTags, fmt.Sprintf(" '%s' '%s%s' ", whereCond.key, op, value))
				}
			}

			computeTags.WriteString(fmt.Sprintf("[ SWAP [] { %s } filter.bylabels ] FILTER \n", strings.Join(filterTags, " ")))
			for _, whereCond := range wheres {
				if !whereCond.isTag {
					hasFilter = true
					selectFields = append(selectFields, whereCond.key)
					value := whereCond.value.Val

					if whereCond.value.Type == influxql.String || whereCond.value.Type == influxql.Unknown {
						if !strings.HasPrefix(value, "'") && !strings.HasPrefix(value, "\"") {
							value = "'" + value + "'"
						}
					}
					computeTags.WriteString(fmt.Sprintf(" [ SWAP [] '%s' filter.byclass ] FILTER \n", classname+p.Separator+whereCond.key))
					switch whereCond.op {
					case influxql.SUB:
						computeTags.WriteString(fmt.Sprintf("[ SWAP %s TODOUBLE -1 * mapper.add 0 0 0 ] MAP\n", value))
					case influxql.DIV:
						computeTags.WriteString(fmt.Sprintf("[ SWAP 1 %s TODOUBLE / mapper.mul 0 0 0 ] MAP\n", value))
					case influxql.BITWISE_AND, influxql.BITWISE_OR, influxql.BITWISE_XOR:
						return "", Unknown, fmt.Errorf("Bitwise %s is currently only supported between numbers (int and float)", whereCond.op.String()) // FIXME)
					default:
						computeTags.WriteString(fmt.Sprintf(" [ SWAP %s mapper.%s 0 0 0 ] MAP \n", value, getWarpScriptOpString(whereCond.op)))

					}
				}
			}

			computeTags.WriteString(" [ SWAP true mapper.replace 0 0 0 ] MAP \n")

			computeTags.WriteString("[ SWAP NULL reducer.and ] REDUCE \n")
		}
		if i > 1 {
			computeTags.WriteString("APPEND")
		}
	}
	separator := "\\."
	if p.Separator != "." {
		separator = p.Separator
	}

	if hasFilter {
		classname = classname + fmt.Sprintf(separator+"(%s|%s)", valueName, strings.Join(selectFields, "|"))
	} else {
		classname = classname + separator + "(" + valueName + ")"
	}

	fetchSelector := fmt.Sprintf("'%s{}'", classname)

	if len(selectors) > 0 {
		fetchSelector = " "

		for _, selector := range selectors {
			fetchSelector += fmt.Sprintf("'%s{%s}' ", classname, selector)
		}
	}

	mc2 += fmt.Sprintf("{ 'token' '%s' 'selectors' [ %s ] 'end' $end 'timespan' $interval } %s\n", p.Token, fetchSelector, fetchString)
	mc2 += "<% DROP DUP NAME 'name' STORE { '.app' NULL '.InfluxDBName' $name '"
	mc2 += p.Separator
	mc2 += "' <% DUP '' == %> <% DROP 1 ->LIST %> <% SPLIT %> IFTE\n"
	mc2 += " 0 GET } RELABEL %> LMAP  \n"

	if p.HasSubQuery {
		mc2 += "<% " + fmt.Sprintf(" 'subqueries-%d' ", p.SubQueryLevel) + " DEFINED %> <% " + fmt.Sprintf("$subqueries-%d ", p.SubQueryLevel)
		mc2 += " <% DROP DUP ATTRIBUTES RELABEL [ [ $start $end ] ] CLIP %> LMAP APPEND %> IFT FLATTEN \n"

	}
	if hasFilter {
		filter := `
		'set' STORE
		$set
		NULL
		PARTITION 
		'partitionSet' STORE
		`

		filter += computeTags.String()
		filter += `
		//APPEND
		[]
		SWAP
		NULL
		PARTITION
		<%  
		[ SWAP [] reducer.or.exclude-nulls ] REDUCE
		'mask' STORE 
		$partitionSet SWAP GET 
		<%
			DROP 
			DUP NAME 'namePSeries' STORE
			'pSeries' STORE
			[
				$mask 
				$pSeries 1 ->LIST
				[]
				op.mask
			] APPLY
			0 GET $namePSeries RENAME
		%>
		LMAP
		APPEND
		%>
		FOREACH
		NONEMPTY
		`

		separator := "\\."

		if p.Separator != "." {
			separator = p.Separator
		}

		filter += fmt.Sprintf("[ SWAP [] '%s' filter.byclass ] FILTER ", fmt.Sprintf("~(%s)", strings.Join(p.Classnames, "|"))+separator+"("+name+")")
		mc2 += filter
	}

	return mc2, SeriesSet, nil
}

// getWarpScriptOpString: Convert an Influx Operation to a single value WarpScript mapper
func getWarpScriptOpString(op influxql.Token) string {
	switch op {
	case influxql.ADD:
		return "add"
	case influxql.MUL:
		return "mul"
	case influxql.DIV:
		return "mul"
	case influxql.MOD:
		return "mod"
	case influxql.EQ, influxql.EQREGEX:
		return "eq"
	case influxql.NEQREGEX, influxql.NEQ:
		return "ne"
	case influxql.LT:
		return "lt"
	case influxql.LTE:
		return "le"
	case influxql.GT:
		return "gt"
	case influxql.GTE:
		return "ge"
	default:
		return ""
	}
}

// getParameterErrorString: Parameters generic error handling
func (p *InfluxParser) getParameterErrorString(lenArgs int, expArgsLen int, expected int, name string) string {
	if expArgsLen == expected {
		return fmt.Sprintf("invalid number of arguments for %s, expected %d, got %d", name, expected, lenArgs)
	}
	return fmt.Sprintf("invalid number of arguments for %s, expected at least %d but no more than %d, got %d", name, expected, expArgsLen, lenArgs)
}

// Parse function call
func (p *InfluxParser) parseCall(expr *influxql.Call, selectors []string, where [][]*WhereCond) (string, ExprReturn, error) {

	isTop := false
	mc2 := ""
	args := make([]string, 0)

	switch expr.Name {
	// Parse an Aggregation method
	case "bottom", "count", "distinct", "first", "integral", "last", "max", "mean", "median", "min", "mode", "percentile", "sample", "spread", "stddev", "sum", "top":
		expArgsLen := 1
		expected := 1

		switch expr.Name {
		case "distinct":
			if len(expr.Args) > 1 {
				errorString := "distinct function can only have one argument"
				return "", Unknown, fmt.Errorf(errorString)
			}
		case "bottom", "top":
			expArgsLen = len(expr.Args) + 1
			expected = 2
			isTop = true
		case "percentile", "sample":
			expArgsLen = 2
			expected = 2
		case "integral":
			expArgsLen = 2
		}

		if len(expr.Args) > expArgsLen {
			errorString := p.getParameterErrorString(len(expr.Args), expArgsLen, expected, expr.Name)
			return "", Unknown, fmt.Errorf(errorString)
		}

		for i, arg := range expr.Args {

			if i == 0 {

				fetch, _, err := p.parseArgumentFetch(expr.Name, arg, selectors, where)
				if err != nil {
					return "", Unknown, err
				}
				mc2 += fetch
			} else if i > 0 {
				compArgs := ""
				var err error
				compArgs, args, err = p.parseComplementaryArgument(expr.Name, arg, args, true, selectors, where, isTop)
				if err != nil {
					return "", Unknown, err
				}
				mc2 += compArgs
			}
		}

		if p.HasGroupBy {
			p.GroupByTags = append(p.GroupByTags, "'.InfluxDBName'")
			mc2 += p.startPartition()
		}

		// Parse method WarpScript
		curFun, err := p.parseAggregationName(expr.Name, args)
		if err != nil {
			return "", Unknown, err
		}
		mc2 += curFun

		if len(p.GroupByTags) > 0 {
			mc2 += `
				+
			%>
			FOREACH
			FLATTEN
			`
		}

	case "atan2":
		if len(expr.Args) != 2 {
			return "", Unknown, fmt.Errorf(invalidNumberOfArgs(expr.Name, 2, len(args)+1))
		}
		for _, arg := range expr.Args {
			compArgs := ""
			var err error
			compArgs, args, err = p.parseComplementaryArgument(expr.Name, arg, args, false, selectors, where, isTop)
			if err != nil {
				return "", Unknown, err
			}
			mc2 += compArgs
		}

		// atan2 expects 2 arguments on the stack
		if p.HasGroupBy {
			p.GroupByTags = append(p.GroupByTags, "'.InfluxDBName'")
			mc2 += `
			<% DROP { '.atan2' 'py' } RELABEL %> LMAP
			SWAP
			<% DROP { '.atan2' 'px' } RELABEL %> LMAP APPEND 
			`
			mc2 += p.startPartition()
			mc2 += `
				'set' STORE
				[ $set [] { '.atan2' 'py' } filter.bylabels ] FILTER 
    			[ $set [] { '.atan2' 'px' } filter.bylabels ] FILTER 
			`
		}

		// Parse method WarpScript
		curFun, err := p.parseTransformationsName(expr.Name, args)
		if err != nil {
			return "", Unknown, err
		}
		mc2 += curFun

		if len(p.GroupByTags) > 0 {
			mc2 += `
				+
			%>
			FOREACH
			FLATTEN
			`
		}
	case "histogram":
		errorString := "undefined function histogram()"
		return errorString, Unknown, fmt.Errorf(errorString)

	case "abs", "acos", "asin", "atan", "ceil", "cos", "cumulative_sum", "derivative", "difference", "elapsed", "exp", "floor",
		"ln", "log", "log2", "log10", "moving_average", "non_negative_derivative", "non_negative_difference", "pow", "round", "sin", "sqrt", "tan":

		expArgsLen := 1
		expected := 1

		switch expr.Name {
		case "derivative", "elapsed", "non_negative_derivative":
			expArgsLen = 2
		case "log", "moving_average", "pow":
			expArgsLen = 2
			expected = 2
		}

		if len(expr.Args) > expArgsLen {
			errorString := p.getParameterErrorString(len(expr.Args), expArgsLen, expected, expr.Name)
			return errorString, Unknown, fmt.Errorf(errorString)
		}

		for _, arg := range expr.Args {
			compArgs := ""
			var err error
			compArgs, args, err = p.parseComplementaryArgument(expr.Name, arg, args, false, selectors, where, isTop)
			if err != nil {
				return "", Unknown, err
			}
			mc2 += compArgs
		}

		if p.HasGroupBy {
			p.GroupByTags = append(p.GroupByTags, "'.InfluxDBName'")
			mc2 += p.startPartition()
		}

		// Parse method WarpScript
		curFun, err := p.parseTransformationsName(expr.Name, args)
		if err != nil {
			return "", Unknown, err
		}
		mc2 += curFun

		if len(p.GroupByTags) > 0 {
			mc2 += `
				+
			%>
			FOREACH
			FLATTEN
			`
		}

	default:
		return "", Unknown, fmt.Errorf("Unimplemented method: %s()", expr.Name)
	}
	return mc2, SeriesSet, nil
}

// parseArgumentFetch: parse Fetch based on influx Expression argument
func (p *InfluxParser) parseArgumentFetch(name string, arg influxql.Expr, selectors []string, where [][]*WhereCond) (string, ExprReturn, error) {
	mc2 := ""
	exprType := Unknown

	// Argument have to be the selected field
	switch argExpr := arg.(type) {
	case *influxql.Wildcard:
		p.HasWildCard = true
		fetch, fetchType, err := p.parseFetch(".*", influxql.Unknown, selectors, where)
		if err != nil {
			return "", Unknown, err
		}
		mc2 += fetch
		p.HasGroupBy = true
		mc2 += p.renameAndSetInfluxStarLabels()
		p.GroupByTags = append(p.GroupByTags, "'.INFLUXQL_COLUMN_NAME'")
		exprType = fetchType
	case *influxql.VarRef:
		p.HasWildCard = false
		fetch, fetchType, err := p.parseFetch(argExpr.Val, argExpr.Type, selectors, where)
		if err != nil {
			return "", Unknown, err
		}
		mc2 += fetch
		exprType = fetchType
	default:
		return "", Unknown, fmt.Errorf("expected field argument in %s()", name)
	}
	return mc2, exprType, nil
}

// parseComplementaryArgument Parse a secondary argument
func (p *InfluxParser) parseComplementaryArgument(name string, arg influxql.Expr, args []string, aggregation bool, selectors []string, where [][]*WhereCond, isTop bool) (string, []string, error) {
	mc2 := ""

	// Store complementary method argument
	switch argExpr := arg.(type) {
	case *influxql.BinaryExpr:
		log.Warnf("parseComplementaryArgument - *influxql.BinaryExpr arg %s", influxql.BinaryExprName(argExpr))
	case *influxql.ParenExpr:
		log.Warnf("parseComplementaryArgument - *influxql.ParenExpr arg %s", influxql.Field{Expr: argExpr.Expr})
	case *influxql.Call:
		call, _, err := p.parseCall(argExpr, selectors, where)
		if err != nil {
			return "", args, err
		}
		mc2 += call
	case *influxql.Wildcard:
		p.HasWildCard = true
		fetch, _, err := p.parseFetch(".*", influxql.Unknown, selectors, where)
		if err != nil {
			return "", args, err
		}
		mc2 += fetch
		p.HasGroupBy = true
		mc2 += p.renameAndSetInfluxStarLabels()
		p.GroupByTags = append(p.GroupByTags, "'.INFLUXQL_COLUMN_NAME'")
	case *influxql.VarRef:
		if isTop {
			args = append(args, fmt.Sprintf("%s", argExpr.Val))
		} else if !aggregation {
			p.HasWildCard = false
			fetch, _, err := p.parseFetch(argExpr.Val, argExpr.Type, selectors, where)
			if err != nil {
				return "", args, err
			}
			mc2 += fetch
		}
	case *influxql.DurationLiteral:
		args = append(args, fmt.Sprintf("%d", argExpr.Val.Nanoseconds()/1000))
	case *influxql.StringLiteral:
		args = append(args, fmt.Sprintf("'%s'", argExpr.Val))
	case *influxql.BooleanLiteral, *influxql.UnsignedLiteral, *influxql.IntegerLiteral, *influxql.NumberLiteral:
		args = append(args, fmt.Sprintf("%s", argExpr.String()))
	default:
		log.Warnf("parseComplementaryArgument - Default %s %T", argExpr, argExpr)
	}
	return mc2, args, nil
}
