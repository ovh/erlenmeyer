package prom

import (
	"fmt"
	"strings"

	"github.com/ovh/erlenmeyer/core"
	"github.com/ovh/erlenmeyer/proto/prom/promql"
	"github.com/spf13/viper"

	log "github.com/sirupsen/logrus"

	"github.com/prometheus/common/model"
	"github.com/prometheus/prometheus/pkg/labels"
)

// An evaluator evaluates given expressions at a fixed timestamp. It is attached to an
// engine through which it connects to a querier and reports errors. On timeout or
// cancellation of its context it terminates.
type evaluator struct {
}

// GenerateInstantQueryTree is creating a new Node for instant queries
func (ev *evaluator) GenerateInstantQueryTree(ctx Context) *core.Node {
	node := core.NewEmptyNode()

	ctx.IsInstant = true

	ctx.Bucketizer = "bucketizer.last"
	ev.eval(ctx.Expr, node, ctx)
	return node
}

func (ev *evaluator) GenerateQueryTree(ctx Context) *core.Node {
	node := core.NewEmptyNode()
	node.Level = 0

	ctx.IsInstant = false

	ctx.Bucketizer = "bucketizer.last"
	ev.eval(ctx.Expr, node, ctx)
	return node
}

// eval evaluates the given expression as the given AST expression node requires.
func (ev *evaluator) eval(expr promql.Expr, node *core.Node, ctx Context) {
	log.WithFields(log.Fields{
		"expr":  expr,
		"proto": "promql",
		"type":  fmt.Sprintf("%T", expr),
		"level": node.Level,
	}).Info("Starting eval")

	switch e := expr.(type) {

	case *promql.BinaryExpr:
		vm := e.VectorMatching
		lhs := core.NewEmptyNode()
		rhs := core.NewEmptyNode()

		node.Left = lhs
		node.Left.Level = node.Level + 1
		if vm != nil {
			node.Left.ChildLabels = vm.Include
		}
		node.Right = rhs
		node.Right.Level = node.Level + 1
		if vm != nil {
			node.Left.ChildLabels = vm.Include
		}

		ev.eval(e.LHS, lhs, ctx)
		ev.eval(e.RHS, rhs, ctx)

		node.Payload = core.BinaryExprPayload{
			Op:             fmt.Sprintf("%+v", e.Op),
			FilteredLabels: make([]string, 0),
			IncludeLabels:  make([]string, 0),
			ReturnBool:     e.ReturnBool,
		}

		if vm != nil {
			node.Payload = core.BinaryExprPayload{
				Op:             fmt.Sprintf("%+v", e.Op),
				IsOn:           vm.On,
				IsIgnoring:     len(vm.MatchingLabels) > 0 && vm.On == false,
				FilteredLabels: vm.MatchingLabels,
				IncludeLabels:  vm.Include,
				Card:           vm.Card.String(),
				ReturnBool:     e.ReturnBool,
			}
		}

	case *promql.Call:
		ev.evalCall(e, node, ctx)

	// http_requests_total{job="prometheus",group="canary"}
	case *promql.VectorSelector:
		ev.vectorSelector(e, node, ctx)

		// instance_cpu_time_ns[5m]
	case *promql.MatrixSelector:
		ev.matrixSelector(e, node, ctx)

	// sum(http_requests_total) without (instance)"
	case *promql.AggregateExpr:
		ev.aggregation(e, node, ctx)

		//(instance_memory_limit_bytes - instance_memory_usage_bytes)
	case *promql.ParenExpr:
		ev.eval(e.Expr, node, ctx)

		// 12
	case *promql.NumberLiteral:
		node.Payload = core.NumberLiteralPayload{
			Value: e.String(),
		}

	case *promql.StringLiteral:
		node.Payload = core.NumberLiteralPayload{
			Value: e.String(),
		}

	case *promql.UnaryExpr:
		node.Payload = core.UnaryExprPayload{
			Op: e.Op.String(),
		}
		lhs := core.NewEmptyNode()
		node.Left = lhs

		ev.eval(e.Expr, lhs, ctx)

	default:
		// FIXME:Flush into a GTS
		log.Errorf(fmt.Sprintf("Type %T is not handled", expr))
	}
}

func (ev *evaluator) evalCall(e *promql.Call, node *core.Node, ctx Context) {
	cfp := core.FunctionPayload{
		Name: e.Func.Name,
	}

	if strings.Contains(e.Func.Name, "over_time") {
		switch e.Func.Name {
		case "quantile_over_time":
			// Verify in WarpScript, if prom param is valid, otherwise return an error message
			ctx.HasMapper = true
			ctx.Mapper = `mean DROP
			TODOUBLE 'quantile' STORE
			<% $quantile 0.0 < %> <% [ SWAP -1.0 0.0 / mapper.replace 0 0 0 ] MAP 0.5 'quantile' STORE %> IFT 
			<% $quantile 1.0 > %> <% [ SWAP 1.0 0.0 / mapper.replace 0 0 0 ] MAP 0.5 'quantile' STORE %> IFT
			$quantile 100.0 * mapper.percentile
			`
		case "avg_over_time":
			ctx.HasMapper = true
			ctx.Mapper = "mean"
		case "min_over_time":
			ctx.HasMapper = true
			ctx.Mapper = "min"
		case "max_over_time":
			ctx.HasMapper = true
			ctx.Mapper = "max"
		case "count_over_time":
			ctx.HasMapper = true
			ctx.Mapper = "count"
		case "sum_over_time":
			ctx.HasMapper = true
			ctx.Mapper = "sum"
		case "stddev_over_time":
			ctx.HasMapper = true
			ctx.Mapper = "sd"
			ctx.MapperValue = "false"
		case "stdvar_over_time":
			ctx.HasMapper = true
			ctx.Mapper = "var"
			ctx.MapperValue = "false"
		}
	}

	if e.Func.Name == "quantile_over_time" {
		// Compute args expression first
		cfp.Args = []string{fmt.Sprintf("'%v' ", e.Args[0])}
		ctx.Args = cfp.Args

		// Prepare left node with quantile_over_time bucketize
		node.Left = core.NewEmptyNode()
		ev.eval(e.Args[1], node.Left, ctx)

	} else if e.Func.Name == "histogram_quantile" {
		cfp.Args = []string{fmt.Sprintf("%v ", e.Args[0])}
		node.Payload = cfp
		node.Left = core.NewEmptyNode()
		ev.eval(e.Args[1], node.Left, ctx)
	} else {
		if len(e.Args) > 1 {

			a := make([]string, len(e.Args)-1)
			for i, arg := range e.Args[1:] {
				evaluator := evaluator{}
				node := core.NewEmptyNode()
				node.Level = 0
				ctx.IsInstant = false
				ctx.Bucketizer = "bucketizer.last"
				evaluator.eval(arg, node, ctx)
				a[i] = node.InternalToWarpScript(fmt.Sprintf("%v ", arg))
			}
			cfp.Args = a

		}

		switch cfp.Name {
		case "changes", "resets", "delta", "rate", "increase", "idelta", "irate", "predict_linear":
			ctx.HasFunction = true
			ctx.FunctionName = cfp.Name
			ctx.Args = cfp.Args
		default:
			node.Payload = cfp
		}
		if len(e.Args) > 0 {
			node.Left = core.NewEmptyNode()

			if strings.Compare(e.Func.Name, "absent") == 0 {
				ctx.hasAbsent = true
			}

			ev.eval(e.Args[0], node.Left, ctx)
		}
	}
}

func (ev *evaluator) matrixSelector(selector *promql.MatrixSelector, node *core.Node, ctx Context) {

	var bucketizePayload core.BucketizePayload
	selRange := fmt.Sprint(selector.Range.Nanoseconds() / 1000)
	bucketizePayload.Op = ctx.Bucketizer
	bucketizePayload.LastBucket = fmt.Sprintf("%v000 ", ctx.End)
	if ctx.IsInstant {
		bucketizePayload.BucketCount = "1"
		bucketizePayload.BucketSpan = "0"
	} else {
		bucketizePayload.BucketCount = fmt.Sprintf("%v000 %v000 %v 2 * - - %v / TOLONG 1 + 2 - ABS", ctx.End, ctx.Start, ctx.Step, ctx.Step)
		bucketizePayload.BucketSpan = fmt.Sprintf("%v ", ctx.Step)
	}
	bucketizePayload.PreBucketize = `
<%
	DROP 
	` + viper.GetString("prometheus.fillprevious.period") + ` DUP 'FILL_PREVIOUS_PERIOD' STORE
    1 'splits_945fa9bc3027d7025e3' TIMESPLIT 
    <% 
        DROP
		DUP LASTTICK 'lt' STORE
		<%
            $lt $end $FILL_PREVIOUS_PERIOD - <=
        %>
        <%
        	DUP FIRSTTICK 'ft' STORE
			[ SWAP bucketizer.last $lt $step $lt $ft - $step / TOLONG 1 + ] BUCKETIZE FILLPREVIOUS 0 GET
		%>
		IFT
        { 'splits_945fa9bc3027d7025e3' '' } RELABEL
    %>
    LMAP
    MERGE
%> 
LMAP 
UNBUCKETIZE
	`

	var fetchPayload core.FetchPayload
	if ctx.hasAbsent {
		bucketizePayload.Absent = true
		fetchPayload.Absent = true
	}

	if ctx.IsInstant {
		fetchPayload.Instant = true
	}

	node.Left = core.NewEmptyNode()
	node.Left.Level = node.Level + 1

	var setName string
	var hasName bool
	setName, hasName, fetchPayload.Labels = labelMatchersToMapLabels(selector.LabelMatchers...)

	if hasName {
		selector.Name = setName
	}
	fetchPayload.ClassName = string(selector.Name)

	if ctx.IsInstant {
		fetchPayload.End = fmt.Sprintf("%v000 ", ctx.End) + fmt.Sprintf("%v", selector.Offset.Nanoseconds()/1000) + " - ISO8601"
		fetchPayload.Start = fmt.Sprintf("%v000 ", ctx.End) + fmt.Sprintf("%v", selector.Offset.Nanoseconds()/1000) + " -  $range - ISO8601"
	} else {
		fetchPayload.Start = fmt.Sprintf("%v000 %v 2 * - ", ctx.Start, ctx.Step) + fmt.Sprintf("%v", selector.Offset.Nanoseconds()/1000) + " - "
		fetchPayload.End = fmt.Sprintf("%v000 ", ctx.End)
		fetchPayload.Step = ctx.Step
	}
	if selector.Offset.String() != "0s" {
		fetchPayload.Offset = fmt.Sprintf("%v", selector.Offset.Nanoseconds()/1000)
	}
	fetchPayload.BucketRange = fmt.Sprintf("%v 'range' STORE\n", selRange)

	node.Left.Left = core.NewEmptyNode()
	node.Left.Left.Level = node.Level + 2
	node.Left.Left.Payload = fetchPayload

	if ctx.HasMapper {
		var mapperPayload core.MapperPayload
		if ctx.MapperValue != "" {
			mapperPayload.Constant = ctx.MapperValue
		}
		mapperPayload.Mapper = ctx.Mapper
		mapperPayload.PreWindow = "1 s $range 1 s - MAX -1 *"
		mapperPayload.PostWindow = "0"
		mapperPayload.Occurrences = "0"
		mapperPayload.Suffix = " { '" + core.ShouldRemoveNameLabel + "' 'true' } SETATTRIBUTES \n"
		if len(ctx.Args) > 0 {
			mapperPayload.Constant = mapperPayload.Constant + ctx.Args[0]
		}
		node.Left.Payload = mapperPayload
	} else if ctx.HasFunction {
		var functionPayload core.FunctionPayload
		functionPayload.Name = ctx.FunctionName
		functionPayload.Args = ctx.Args
		functionPayload.Prefix = "<% $range $step < %> <% DROP [] %> IFT\n"
		node.Left.Payload = functionPayload
	}
	node.Payload = bucketizePayload
}

func labelMatchersToMapLabels(matrixSelector ...*labels.Matcher) (string, bool, map[string]string) {
	returnLabels := make(map[string]string)

	name := ""
	hasName := false
	for _, label := range matrixSelector {
		var val string
		switch label.Type {
		case labels.MatchEqual:
			val = string(label.Value)
		case labels.MatchNotEqual:
			if label.Value == "" {
				val = "~(?=\\s*\\S).*"
			} else {
				val = "~(?!" + string(label.Value) + ").*"
			}
		case labels.MatchRegexp:
			val = "~" + string(label.Value)
		case labels.MatchNotRegexp:
			val = "~(?!" + string(label.Value) + ").*"
		}
		if label.Name == "__name__" {
			name = val
			hasName = true
		} else {
			returnLabels[string(label.Name)] = val
		}
	}
	return name, hasName, returnLabels
}

// vectorSelector evaluates a *VectorSelector expression.
func (ev *evaluator) vectorSelector(selector *promql.VectorSelector, node *core.Node, ctx Context) {

	if ctx.IsInstant {
		var bucketizePayload core.BucketizePayload
		bucketizePayload.Op = "bucketizer.last"
		bucketizePayload.LastBucket = fmt.Sprintf("%v000 ", ctx.End)
		bucketizePayload.BucketSpan = "0"
		bucketizePayload.BucketCount = "1"

		node.Payload = bucketizePayload

		var fetchPayload core.FetchPayload
		var setName string
		var hasName bool
		fetchPayload.Instant = true

		setName, hasName, fetchPayload.Labels = labelMatchersToMapLabels(selector.LabelMatchers...)

		if hasName {
			selector.Name = setName
		}
		fetchPayload.ClassName = string(selector.Name)

		fetchPayload.Step = ctx.Step

		fetchPayload.End = " -1 "
		fetchPayload.Start = fmt.Sprintf("%v000 ", ctx.End) + fmt.Sprintf("%v", selector.Offset.Nanoseconds()/1000) + " - "

		if selector.Offset.String() != "0s" {
			fetchPayload.Offset = fmt.Sprintf("%v", selector.Offset.Nanoseconds()/1000)
		}

		node.Left = core.NewEmptyNode()
		node.Left.Level = node.Level + 1
		node.Left.Payload = fetchPayload

	} else {

		var bucketizePayload core.BucketizePayload
		bucketizePayload.Op = "bucketizer.last"
		bucketizePayload.LastBucket = fmt.Sprintf("%v000 ", ctx.End)
		bucketizePayload.BucketSpan = fmt.Sprintf("%v ", ctx.Step)

		bucketizePayload.BucketCount = fmt.Sprintf("%v000 %v000 %v 2 * -  - %v / TOLONG 1 + 2 - ABS", ctx.End, ctx.Start, ctx.Step, ctx.Step)
		bucketizePayload.PreBucketize = `
<%
	DROP 
	` + viper.GetString("prometheus.fillprevious.period") + ` DUP 'FILL_PREVIOUS_PERIOD' STORE
    1 'splits_945fa9bc3027d7025e3' TIMESPLIT 
    <% 
        DROP
		DUP LASTTICK 'lt' STORE
		<%
            $lt $end $FILL_PREVIOUS_PERIOD - <=
        %>
        <%
        	DUP FIRSTTICK 'ft' STORE
			[ SWAP bucketizer.last $lt $step $lt $ft - $step / TOLONG 1 + ] BUCKETIZE FILLPREVIOUS 0 GET
		%>
		IFT
        { 'splits_945fa9bc3027d7025e3' '' } RELABEL
    %>
    LMAP
    MERGE
%> 
LMAP 
UNBUCKETIZE
	`
		var fetchPayload core.FetchPayload
		if ctx.hasAbsent {
			bucketizePayload.Absent = true
			fetchPayload.Absent = true
		}
		node.Payload = bucketizePayload

		node.Left = core.NewEmptyNode()
		node.Left.Level = node.Level + 1

		var setName string
		var hasName bool

		setName, hasName, fetchPayload.Labels = labelMatchersToMapLabels(selector.LabelMatchers...)

		if hasName {
			selector.Name = setName
		}

		fetchPayload.ClassName = string(selector.Name)
		fetchPayload.Step = ctx.Step

		fetchPayload.End = fmt.Sprintf("%v000 ", ctx.End)
		fetchPayload.Start = fmt.Sprintf("%v000 %v 2 * - ", ctx.Start, ctx.Step) + fmt.Sprintf("%v", selector.Offset.Nanoseconds()/1000) + " - "

		if selector.Offset.String() != "0s" {
			fetchPayload.Offset = fmt.Sprintf("%v", selector.Offset.Nanoseconds()/1000)
		}

		node.Left.Payload = fetchPayload
	}
}

func labelNamesToStringSlice(grouping model.LabelNames) []string {
	var labels []string
	for _, label := range grouping {
		labels = append(labels, string(label))
	}
	return labels
}

// aggregation evaluates an aggregation operation on a vector.
func (ev *evaluator) aggregation(e *promql.AggregateExpr, node *core.Node, ctx Context) {

	var payload core.AggregatePayload

	payload.Op = fmt.Sprintf("%+v", e.Op)
	payload.Grouping = e.Grouping
	payload.Without = e.Without

	// Parameter used by some aggregators like topk(3, ...)
	if e.Param != nil {

		if payload.Op == "quantile" {
			payload.Param = "$right"
			node.Right = core.NewEmptyNode()
			node.Right.Level = node.Level + 1
			ev.eval(e.Param, node.Right, ctx)
		} else {
			payload.Param = string(e.Param.String())
		}
	}

	node.Payload = payload
	node.Left = core.NewEmptyNode()
	node.Left.Level = node.Level + 1

	ev.eval(e.Expr, node.Left, ctx)
}

// Copy paste from prometheus source code
type itemType int

const (
	itemError itemType = iota // Error occurred, value is error message
	itemEOF
	itemComment
	itemIdentifier
	itemMetricIdentifier
	itemLeftParen
	itemRightParen
	itemLeftBrace
	itemRightBrace
	itemLeftBracket
	itemRightBracket
	itemComma
	itemAssign
	itemSemicolon
	itemString
	itemNumber
	itemDuration
	itemBlank
	itemTimes

	operatorsStart
	// Operators.
	itemSUB
	itemADD
	itemMUL
	itemMOD
	itemDIV
	itemLAND
	itemLOR
	itemLUnless
	itemEQL
	itemNEQ
	itemLTE
	itemLSS
	itemGTE
	itemGTR
	itemEQLRegex
	itemNEQRegex
	itemPOW
	operatorsEnd

	aggregatorsStart
	// Aggregators.
	itemAvg
	itemCount
	itemSum
	itemMin
	itemMax
	itemStddev
	itemStdvar
	itemTopK
	itemBottomK
	itemCountValues
	itemQuantile
	aggregatorsEnd

	keywordsStart
	// Keywords.
	itemAlert
	itemIf
	itemFor
	itemLabels
	itemAnnotations
	itemKeepCommon
	itemOffset
	itemBy
	itemWithout
	itemOn
	itemIgnoring
	itemGroupLeft
	itemGroupRight
	itemBool
	keywordsEnd
)
