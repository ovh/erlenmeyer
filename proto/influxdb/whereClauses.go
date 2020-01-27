package influxdb

// Parse a Where statement
import (
	"fmt"
	"strings"

	"github.com/influxdata/influxql"
	log "github.com/sirupsen/logrus"
)

//WhereCond Where statement data
type WhereCond struct {
	key   string
	op    influxql.Token
	value *influxql.VarRef
	isTag bool
}

// String: WhereCond to string
func (w *WhereCond) String() string {
	return fmt.Sprintf("WhereCond{%s %s %s %t}", w.key, w.op.String(), w.value, w.isTag)
}

// parseWhere Parse an Influx Where statement
// https://docs.influxdata.com/influxdb/v1.7/query_language/data_exploration/#the-where-clause
func (p *InfluxParser) parseWhere(where influxql.Expr) ([][]*WhereCond, error) {
	selectTags := make([][]*WhereCond, 0)
	// Return the function name or variable name, if available.
	switch expr := where.(type) {
	case *influxql.BinaryExpr:
		return p.parseOperator(expr.LHS, expr.Op, expr.RHS)
	case *influxql.ParenExpr:
		return p.parseWhere(expr.Expr)
	default:
		log.Warnf("parseWhere Default - where: %s, %T", expr.String(), expr)
	}
	return selectTags, nil
}

// Parse a Whereclause To get Separator
func parseSeparatorCondition(getExpr influxql.Expr) (string, error) {
	separator := "."
	hasSeparator := false

	// Return the function name or variable name, if available.
	switch expr := getExpr.(type) {
	case *influxql.BinaryExpr:
		if expr.LHS.String() == "_separator" {
			switch expr.Op {
			case influxql.EQ:
				separator = getInfluxValSelector(expr.RHS)
				hasSeparator = true
			default:
			}
		} else if expr.RHS.String() == "_separator" {
			switch expr.Op {
			case influxql.EQ:
				separator = getInfluxValSelector(expr.LHS)
				hasSeparator = true
			default:
			}
		} else {
			if expr.Op == influxql.AND || expr.Op == influxql.OR {
				separatorA, errA := parseSeparatorCondition(expr.LHS)
				separatorB, errB := parseSeparatorCondition(expr.RHS)

				if errA == nil && errB != nil {
					separator = separatorA
					hasSeparator = true
				} else if errB == nil && errA != nil {
					separator = separatorB
					hasSeparator = true
				}
			}
		}
	default:
		log.Warnf("parseSeparatorCondition - Default: %s - %T", expr, expr)
	}

	if !hasSeparator {
		return separator, fmt.Errorf("Nil separator")
	}

	return separator, nil
}

// Parse a Where Time
func parseTimeCondition(getExpr influxql.Expr) (string, string) {
	end := "$now"
	start := "0"

	// Return the function name or variable name, if available.
	switch expr := getExpr.(type) {
	case *influxql.BinaryExpr:
		if expr.LHS.String() == "time" {
			switch expr.Op {
			case influxql.EQ:
				end = getInfluxValue(expr.RHS) + " 1 +"
				start = "$end 10 -"
			case influxql.LT:
				end = getInfluxValue(expr.RHS)
			case influxql.LTE:
				end = getInfluxValue(expr.RHS) + " 1 +"
			case influxql.GT:
				start = getInfluxValue(expr.RHS)
			case influxql.GTE:
				start = getInfluxValue(expr.RHS) + " 1 +"
			default:
			}
		} else if expr.RHS.String() == "time" {
			switch expr.Op {
			case influxql.EQ:
				end = getInfluxValue(expr.LHS) + " 1 +"
				start = "10"
			case influxql.LT:
				end = getInfluxValue(expr.LHS)
			case influxql.LTE:
				end = getInfluxValue(expr.LHS) + " 1 +"
			case influxql.GT:
				start = getInfluxValue(expr.LHS)
			case influxql.GTE:
				start = getInfluxValue(expr.LHS) + " 1 +"
			default:
			}
		} else {
			if expr.Op == influxql.AND || expr.Op == influxql.OR {
				lhsEnd, lhsStart := parseTimeCondition(expr.LHS)
				rhsEnd, rhsStart := parseTimeCondition(expr.RHS)

				if lhsEnd != end && rhsEnd == end {
					end = lhsEnd
				} else if rhsEnd != end && lhsEnd == end {
					end = rhsEnd
				} else if rhsEnd != end && lhsEnd != end {
					end = rhsEnd + " " + lhsEnd + " MAX"
				}

				if lhsStart != start && rhsStart == start {
					start = lhsStart
				} else if rhsStart != start && lhsStart == start {
					start = rhsStart
				} else if rhsStart != start && lhsStart != start {
					start = rhsStart + " " + lhsStart + " MIN"
				}
			}
		}
	default:
		log.Warnf("parseTimeCondition - Default: %s - %T", expr, expr)
	}
	return end, start
}

// Parse a specific Influx value used in Where clauses
func getInfluxValSelector(getExpr influxql.Expr) string {
	value := ""
	// Return the function name or variable name, if available.
	switch expr := getExpr.(type) {
	case *influxql.BooleanLiteral:
		return fmt.Sprintf("%T", expr.Val)
	case *influxql.DurationLiteral:
		return fmt.Sprintf("%d", expr.Val.Nanoseconds()/1000)
	case *influxql.IntegerLiteral:
		return fmt.Sprintf("%d", expr.Val)
	case *influxql.UnsignedLiteral:
		return fmt.Sprintf("%d", expr.Val)
	case *influxql.NilLiteral:
		return "nil"
	case *influxql.NumberLiteral:
		return fmt.Sprintf("%f", expr.Val)
	case *influxql.StringLiteral:
		return expr.Val
	case *influxql.VarRef:
		return expr.Val
	default:
		log.Warnf("getInfluxValSelector - Default: %s - %T", expr, expr)
	}
	return value
}

// Parse a specific Influx value used in Where clauses
func getInfluxValue(getExpr influxql.Expr) string {
	value := ""
	// Return the function name or variable name, if available.
	switch expr := getExpr.(type) {
	case *influxql.BinaryExpr:
		return getInfluxValue(expr.LHS) + " " + getInfluxValue(expr.RHS) + " " + expr.Op.String()
	case *influxql.BooleanLiteral:
		return fmt.Sprintf("%T", expr.Val)
	case *influxql.BoundParameter:
		log.Warnf("getInfluxValue - BoundParameter: %s", expr)
	case *influxql.Call:
		switch expr.Name {
		case "now":
			return "$now"
		}
	case *influxql.Distinct:
		log.Warnf("getInfluxValue - Distinct: %s", expr)
	case *influxql.DurationLiteral:
		return fmt.Sprintf("%d", expr.Val.Nanoseconds()/1000)
	case *influxql.IntegerLiteral:
		return fmt.Sprintf("%d", expr.Val)
	case *influxql.UnsignedLiteral:
		return fmt.Sprintf("%d", expr.Val)
	case *influxql.NilLiteral:
		return "NULL"
	case *influxql.NumberLiteral:
		return fmt.Sprintf("%f", expr.Val)
	case *influxql.ParenExpr:
		log.Warnf("getInfluxValue - ParenExpr: %s", expr)
	case *influxql.RegexLiteral:
		log.Warnf("getInfluxValue - RegexLiteral: %s", expr)
	case *influxql.ListLiteral:
		log.Warnf("getInfluxValue - ListLiteral: %s", expr)
	case *influxql.StringLiteral:
		return "'" + expr.Val + "'"
	case *influxql.TimeLiteral:
		log.Warnf("getInfluxValue - TimeLiteral: %s", expr)
	case *influxql.VarRef:
		switch expr.Type {
		case influxql.Unknown:
			return "'" + expr.Val + "' TOTIMESTAMP"
		case influxql.String:
			return "'" + expr.Val + "'"
		case influxql.Duration:
			return parseShift(expr.Val)
		default:
			return expr.Val
		}
	case *influxql.Wildcard:
		log.Warnf("getInfluxValue - Wildcard: %s", expr)
	default:
		log.Warnf("getInfluxValue - Default: %s - %T", expr, expr)
	}
	return value
}

// Parse an Influx Operator used in Where clauses
func (p *InfluxParser) parseOperator(lhs influxql.Expr, op influxql.Token, rhs influxql.Expr) ([][]*WhereCond, error) {
	if lhs.String() == "time" || rhs.String() == "time" {
		return make([][]*WhereCond, 0), nil
	} else if lhs.String() == "_separator" || rhs.String() == "_separator" {
		return make([][]*WhereCond, 0), nil
	}

	switch op {
	case influxql.AND:
		left, errLeft := p.parseWhere(lhs)
		if errLeft != nil {
			return nil, errLeft
		}
		right, errRight := p.parseWhere(rhs)
		if errRight != nil {
			return nil, errRight
		}

		if len(left) == 0 {
			return right, nil
		}
		if len(right) == 0 {
			return left, nil
		}
		resAND := make([][]*WhereCond, 0)
		for _, leftList := range left {
			resList := append(leftList[:0:0], leftList...)
			for _, rightList := range right {
				resList = append(resList, rightList...)
			}
			resAND = append(resAND, resList)
		}
		return resAND, nil

	case influxql.OR:
		left, errLeft := p.parseWhere(lhs)
		if errLeft != nil {
			return nil, errLeft
		}
		right, errRight := p.parseWhere(rhs)
		if errRight != nil {
			return nil, errRight
		}
		return append(left, right...), nil
	case influxql.EQ, influxql.EQREGEX, influxql.NEQREGEX, influxql.NEQ, influxql.LT, influxql.LTE, influxql.GT, influxql.GTE, influxql.ADD, influxql.SUB, influxql.MUL, influxql.DIV:
		eqTags := make([][]*WhereCond, 0)
		switch rhsExpr := rhs.(type) {
		case *influxql.VarRef:
			eqTag := make([]*WhereCond, 1)
			eqTag[0] = &WhereCond{key: lhs.String(), op: op, value: rhsExpr}
			eqTags = append(eqTags, eqTag)
		case *influxql.StringLiteral:
			rhsValue := &influxql.VarRef{Type: influxql.String, Val: rhsExpr.String()}
			eqTag := make([]*WhereCond, 1)
			eqTag[0] = &WhereCond{key: lhs.String(), op: op, value: rhsValue}
			eqTags = append(eqTags, eqTag)
		case *influxql.BooleanLiteral:
			rhsValue := &influxql.VarRef{Type: influxql.Boolean, Val: rhsExpr.String()}
			eqTag := make([]*WhereCond, 1)
			eqTag[0] = &WhereCond{key: lhs.String(), op: op, value: rhsValue}
			eqTags = append(eqTags, eqTag)
		case *influxql.NumberLiteral:
			rhsValue := &influxql.VarRef{Type: influxql.Float, Val: rhsExpr.String()}
			eqTag := make([]*WhereCond, 1)
			eqTag[0] = &WhereCond{key: lhs.String(), op: op, value: rhsValue}
			eqTags = append(eqTags, eqTag)
		case *influxql.IntegerLiteral:
			rhsValue := &influxql.VarRef{Type: influxql.Integer, Val: rhsExpr.String()}
			eqTag := make([]*WhereCond, 1)
			eqTag[0] = &WhereCond{key: lhs.String(), op: op, value: rhsValue}
			eqTags = append(eqTags, eqTag)
		case *influxql.RegexLiteral:
			rhsValue := &influxql.VarRef{Type: influxql.String, Val: strings.Trim(rhsExpr.String(), "/")}
			eqTag := make([]*WhereCond, 1)
			eqTag[0] = &WhereCond{key: lhs.String(), op: op, value: rhsValue}
			eqTags = append(eqTags, eqTag)
		default:
			log.Warnf("parseOperator - Default: %s %T", rhsExpr.String(), rhsExpr)
		}
		return eqTags, nil
	default:
		return nil, fmt.Errorf("Operator %s is not supported", op.String())
	}
}

// Generate duration string value based on a String value
func parseShift(val string) string {
	twoEnd := val[len(val)-2:]
	end := val[len(val)-1:]

	switch twoEnd {
	case "ms", "us", "ns", "ps":
		return getUnit(val, twoEnd)
	}

	switch end {
	case "w", "d", "h", "m", "s":
		return getUnit(val, end)
	}

	return ""
}

// Trim a time unit
func getUnit(duration string, unit string) string {
	return strings.TrimRight(duration, unit) + " " + unit
}
