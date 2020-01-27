package graphite

import (
	"bytes"
	"errors"
	"fmt"
	"strconv"
	"strings"

	"github.com/ovh/erlenmeyer/core"
)

var (
	aggregateOperator = map[string]string{
		"average":  "mean",
		"avg":      "mean",
		"count":    "count",
		"last":     "last",
		"median":   "median",
		"sum":      "sum",
		"min":      "min",
		"max":      "max",
		"diff":     "diff",
		"stddev":   "sd",
		"range":    "range",
		"multiply": "product",
		"first":    "first",
	}
)

// ----------------------------------------------------------------------------
// helper functions

func insertWildcardAt(target string, position int) string {
	if target == swap {
		return target
	}

	pieces := strings.SplitN(target, ";", 1)

	split := strings.Split(pieces[0], ".")
	start := split[:position-1]
	end := split[position-1:]

	serie := append(start, "*")
	serie = append(serie, end...)

	if len(pieces) == 1 {
		return strings.Join(serie, ".")
	}

	return fmt.Sprintf("%s;%s", strings.Join(serie, "."), pieces[1])
}

// ----------------------------------------------------------------------------
// graphite functions implementations

func aggregate(node *core.Node, args []string, kwargs map[string]string) (*core.Node, error) {
	if len(args) < 2 {
		return nil, errors.New("The aggregate function take two parameters which are series and the aggregator operator")
	}

	op, ok := aggregateOperator[args[1]]
	if !ok {
		return nil, fmt.Errorf("The aggregator operator %s is not supported", args[1])
	}

	switch op {
	case "mean", "median", "count", "sum", "min", "max", "product":
		node.Left = core.NewNode(core.ReducerPayload{
			Reducer: op,
		})

		node = node.Left
	case "diff":
		node.Left = core.NewNode(core.MapperPayload{
			Constant:    "-1.0",
			Mapper:      "mul",
			Occurrences: "0",
			PostWindow:  "0",
			PreWindow:   "0",
		})

		node.Left.Left = core.NewNode(core.ReducerPayload{
			Reducer: "sum",
		})

		node = node.Left.Left
	case "range", "rangeOf":
		node.Left = core.NewNode(core.WarpScriptPayload{
			WarpScript: `
			DUP [ SWAP [ ] reducer.max ] REDUCE
			SWAP [ SWAP [ ] reducer.min ] REDUCE
			APPEND
			[ SWAP -1.0 mapper.mul 0 0 0 ] MAP
			[ SWAP [ ] reducer.sum ] REDUCE`,
		})

		node = node.Left

	case "sd":
		node.Left = core.NewNode(core.ReducerPayload{
			Reducer: op,
			Value:   "true",
		})

		node = node.Left
	case "last":
		node.Left = core.NewNode(core.WarpScriptPayload{
			WarpScript: `
			DUP SIZE 1 - GET 1 ->LIST`,
		})
	case "first":
		node.Left = core.NewNode(core.WarpScriptPayload{
			WarpScript: `
			0 GET 1 ->LIST`,
		})
	}

	if args[0] != swap {
		return fetch(node, []string{args[0], kwargs["from"], kwargs["until"]}, kwargs)
	}

	return node, nil
}

func aggregateLine(node *core.Node, args []string, kwargs map[string]string) (*core.Node, error) {
	if len(args) < 1 {
		return nil, errors.New("The aggregateLine function take at least one parameter which is a list of series")
	}

	aggregator := "avg"
	if len(args) >= 2 {
		aggregator = args[1]
	}

	keepStep := false
	keepStepString := ""
	if len(args) >= 3 {
		var err error
		keepStep, err = strconv.ParseBool(args[2])
		if err != nil {
			return nil, fmt.Errorf("Expect the keepStep %s parameter to be a boolean", args[1])
		}
	}

	if !keepStep {
		keepStepString = " FILLPREVIOUS FILLNEXT "
	}

	op, ok := aggregateOperator[aggregator]
	if !ok {
		return nil, fmt.Errorf("The aggregator operator %s is not supported", args[1])
	}

	switch op {
	case "mean", "median", "count", "sum", "min", "max", "product", "first", "last":

		node.Left = core.NewNode(core.WarpScriptPayload{
			WarpScript: `
			NONEMPTY
			<%
				DROP DUP 
				[ SWAP bucketizer.` + op + ` 0 0 1 ] BUCKETIZE 0 GET VALUES 0 GET 
				[ ROT ROT mapper.replace 0 0 0 ] MAP 0 GET
			%>
			LMAP ` + keepStepString,
		})

		node = node.Left
	case "sd":
		node.Left = core.NewNode(core.WarpScriptPayload{
			WarpScript: `
			NONEMPTY
			<%
				DROP DUP 
				[ SWAP true mapper.` + op + ` MAXLONG 0 -1 ] MAP 0 GET VALUES 0 GET 
				[ ROT ROT mapper.replace 0 0 0 ] MAP 0 GET
			%>
			LMAP`,
		})

		node = node.Left

	case "diff", "range", "rangeOf":
		return nil, fmt.Errorf("The aggregator operator %s is not supported", args[1])
	}

	if args[0] != swap {
		return fetch(node, []string{args[0], kwargs["from"], kwargs["until"]}, kwargs)
	}

	return node, nil
}

func groupByNode(node *core.Node, args []string, kwargs map[string]string) (*core.Node, error) {
	// Conditional reverse of args
	// Handle both :
	//   From graphite documentation : groupByNode(series, nodeNum, callback='average')
	//   Previous behavior : groupByNode => aggregateWithWildcards(series, func, *position)
	if len(args) == 2 {
		args = append(args, "average")
	}
	if len(args) == 3 {
		_, err := strconv.Atoi(args[1])
		if err == nil {
			tmp := args[2]
			args[2] = args[1]
			args[1] = tmp
		}
	}
	return aggregateWithWildcards(node, args, kwargs)
}

func aggregateWithWildcards(node *core.Node, args []string, kwargs map[string]string) (*core.Node, error) {
	var err error
	var equivalenceClass bytes.Buffer
	if len(args) < 3 {
		return nil, errors.New("The aggregateWithWildcards function takes at least three parameters which are a series, an aggregator operator and number(s)")
	}

	// Fill equivalence class based wildcards positions arguments
	equivalenceClass.WriteString("[ ")
	for i, arg := range args {
		if i >= 2 {
			_, err := strconv.Atoi(arg)
			if err != nil {
				return nil, err
			}
			equivalenceClass.WriteString("'.class-' '" + arg + "' + ")
		}
	}
	equivalenceClass.WriteString("] 'equivalenceClass' STORE")

	// Post Treatment: Remove generated labels
	node.Left = core.NewNode(core.WarpScriptPayload{
		WarpScript: `
		<%
			DROP DUP LABELS DUP 'labels' STORE KEYLIST LSORT REVERSE {}
			'' 'name' STORE
			SWAP
			<%
					<%
							DUP '.class-.*'
							MATCH
							SIZE 0 >
					%>
					<%
							DUP $labels SWAP GET  '.' SWAP + $name SWAP + 'name' STORE
							'' SWAP PUT
					%>
					<%
							DROP
					%>
					IFTE
			%>
			FOREACH
			RELABEL
			<% 
				$name SIZE 1 >
			%>
			<%
				$name 
				1 $name SIZE SUBSTRING
				RENAME
			%>
			IFT
		%>
		LMAP
		`,
	})

	node = node.Left

	// Apply the aggregation reduce
	node, err = aggregate(node, []string{swap, args[1]}, kwargs)
	if err != nil {
		return nil, err
	}

	// Pre Treatment: generate a '.class' label per point
	node.Left = core.NewNode(core.WarpScriptPayload{
		WarpScript: `
		<% 
			TOSTRING 'index' STORE DUP NAME '.' SPLIT REVERSE
			{} SWAP
			<%  
				TOSTRING '.class-' SWAP +
				PUT 
				0
			%>
			LMAP
			DROP
			RELABEL
		%> LMAP 
		` + equivalenceClass.String(),
	})

	node = node.Left

	if args[0] != swap {
		return fetch(node, []string{args[0], kwargs["from"], kwargs["until"]}, kwargs)
	}

	return node, nil
}

func averageSeries(node *core.Node, args []string, kwargs map[string]string) (*core.Node, error) {
	if len(args) < 1 {
		return nil, errors.New("The averageSeries take one parameter which is the series or series")
	}

	return aggregate(node, []string{args[0], "average"}, kwargs)
}

func averageSeriesWithWildcards(node *core.Node, args []string, kwargs map[string]string) (*core.Node, error) {
	if len(args) < 2 {
		return nil, errors.New("The averageSeriesWithWildcards function take two parameters which are a list of series and a number")
	}

	aggArgs := make([]string, len(args)+1)

	for i, arg := range args {
		if i == 0 {
			aggArgs[i] = arg
			aggArgs[1] = "average"
		} else {
			aggArgs[i+1] = arg
		}
	}

	return aggregateWithWildcards(node, aggArgs, kwargs)
}

func sumSeries(node *core.Node, args []string, kwargs map[string]string) (*core.Node, error) {
	if len(args) < 1 {
		return nil, errors.New("The sumSeries function take one parameter which is a list of series")
	}

	return aggregate(node, []string{args[0], "sum"}, kwargs)
}

func sumSeriesWithWildcards(node *core.Node, args []string, kwargs map[string]string) (*core.Node, error) {
	if len(args) < 2 {
		return nil, errors.New("The sumSeriesWithWildcards function take two parameters which are a list of series and a number")
	}

	aggArgs := make([]string, len(args)+1)

	for i, arg := range args {
		if i == 0 {
			aggArgs[i] = arg
			aggArgs[1] = "sum"
		} else {
			aggArgs[i+1] = arg
		}
	}

	return aggregateWithWildcards(node, aggArgs, kwargs)
}

func diffSeries(node *core.Node, args []string, kwargs map[string]string) (*core.Node, error) {
	if len(args) < 1 {
		return nil, errors.New("The diffSeries function take one parameter which is a list of series")
	}

	return aggregate(node, []string{args[0], "diff"}, kwargs)
}

func minSeries(node *core.Node, args []string, kwargs map[string]string) (*core.Node, error) {
	if len(args) < 1 {
		return nil, errors.New("The minSeries function take one parameter which is a list of series")
	}

	return aggregate(node, []string{args[0], "min"}, kwargs)
}

func maxSeries(node *core.Node, args []string, kwargs map[string]string) (*core.Node, error) {
	if len(args) < 1 {
		return nil, errors.New("The maxSeries function take one parameter which is a list of series")
	}

	return aggregate(node, []string{args[0], "max"}, kwargs)
}

func multiplySeries(node *core.Node, args []string, kwargs map[string]string) (*core.Node, error) {
	if len(args) < 1 {
		return nil, errors.New("The multiplySeries function take one parameter which is a list of series")
	}

	return aggregate(node, []string{args[0], "multiply"}, kwargs)
}

func stddevSeries(node *core.Node, args []string, kwargs map[string]string) (*core.Node, error) {
	if len(args) < 1 {
		return nil, errors.New("The stddevSeries function take one parameter which is a list of series")
	}

	return aggregate(node, []string{args[0], "stddev"}, kwargs)
}

func stdev(node *core.Node, args []string, kwargs map[string]string) (*core.Node, error) {
	var err error

	if len(args) < 2 {
		return nil, errors.New("The stdev function take at least two parameters which are a list of series and a threshold")
	}

	node, err = maximumAbove(node, []string{args[0], args[1]}, kwargs)
	if err != nil {
		return nil, err
	}

	return aggregate(node, []string{args[0], "stddev"}, kwargs)
}

func multiplySeriesWithWildcards(node *core.Node, args []string, kwargs map[string]string) (*core.Node, error) {
	if len(args) < 2 {
		return nil, errors.New("The multiplySeriesWithWildcards function take two parameters which are a list of series and a number")
	}

	aggArgs := make([]string, len(args)+1)

	for i, arg := range args {
		if i == 0 {
			aggArgs[i] = arg
			aggArgs[1] = "multiply"
		} else {
			aggArgs[i+1] = arg
		}
	}

	return aggregateWithWildcards(node, aggArgs, kwargs)
}

func rangeOfSeries(node *core.Node, args []string, kwargs map[string]string) (*core.Node, error) {
	if len(args) < 1 {
		return nil, errors.New("The rangeOfSeries function take one parameter which is a list of series")
	}

	return aggregate(node, []string{args[0], "range"}, kwargs)
}
