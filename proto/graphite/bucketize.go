package graphite

import (
	"errors"
	"fmt"

	"github.com/ovh/erlenmeyer/core"
)

var (
	consolidateOperator = map[string]string{
		"sum":     "sum",
		"average": "mean",
		"min":     "min",
		"max":     "max",
		"first":   "first",
		"last":    "last",
	}
)

// ----------------------------------------------------------------------------
// helper functions

func bucketize(node *core.Node, args []string, kwargs map[string]string) (*core.Node, error) {
	count := "0"
	span := "60 s" // nolint: goconst
	end := "0"     // nolint: goconst

	if val, ok := kwargs["span"]; ok && len(val) > 0 {
		span = val
	}

	if val, ok := kwargs["count"]; ok && len(val) > 0 {
		count = val
	}
	if val, ok := kwargs["end"]; ok && len(val) > 0 {
		end = val
	}

	node.Left = core.NewNode(core.BucketizePayload{
		Op:          fmt.Sprintf("bucketizer.%s", args[1]),
		BucketCount: count,
		BucketSpan:  span,
		LastBucket:  end,
	})

	if args[0] != swap {
		return fetch(node.Left, []string{args[0], kwargs["from"], kwargs["until"]}, kwargs)
	}

	return node.Left, nil
}

// ----------------------------------------------------------------------------
// graphite functions implementations

// cumulative is an alias for consolidateBy(series, 'sum')
func cumulative(node *core.Node, args []string, kwargs map[string]string) (*core.Node, error) {
	if len(args) < 1 {
		return nil, errors.New("The cumulative function take one parameter which is a series or a list of series")
	}

	return consolidateBy(node, []string{args[0], "sum"}, kwargs)
}

func hitcount(node *core.Node, args []string, kwargs map[string]string) (*core.Node, error) {
	if len(args) < 2 {
		return nil, errors.New("The hitcount function take three parameters which are a series or a list of series, an interval string and optionally a boolean")
	}

	node.Left = core.NewNode(core.MapperPayload{
		Mapper:      "delta",
		Occurrences: "0",
		PostWindow:  "1",
		PreWindow:   "0",
	})

	t, err := convertTimeToWarpScript(args[1])
	if err != nil {
		return nil, err
	}

	node.Left.Left = core.NewNode(core.BucketizePayload{
		Op:          "bucketizer.max",
		BucketCount: "0",
		BucketSpan:  *t,
		LastBucket:  "0",
	})

	if args[0] != swap {
		return fetch(node.Left.Left, []string{args[0], kwargs["from"], kwargs["until"]}, kwargs)
	}

	return node.Left.Left, nil
}

func consolidateBy(node *core.Node, args []string, kwargs map[string]string) (*core.Node, error) {
	if len(args) < 2 {
		return nil, errors.New("The consolidateBy function take two parameters which are a list of series and the consolidate function")
	}
	aggregator := "avg"
	if len(args) >= 2 {
		aggregator = args[1]
	}

	op, ok := aggregateOperator[aggregator]
	if !ok {
		return nil, fmt.Errorf("The aggregator operator %s is not supported", args[1])
	}

	switch op {
	case "mean", "median", "count", "sum", "min", "max", "product", "first", "last":
		kwargs["consolidate"] = op

	case "diff", "range", "rangeOf", "sd":
		return nil, fmt.Errorf("The aggregator operator %s is not supported", args[1])
	}

	if args[0] != swap {
		return fetch(node, []string{args[0], kwargs["from"], kwargs["until"]}, kwargs)
	}

	return node, nil
}

func summarize(node *core.Node, args []string, kwargs map[string]string) (*core.Node, error) {
	var err error

	if len(args) < 2 {
		return nil, errors.New("The summarize function take at least two parameters which a list of series and an interval string")
	}

	t, err := convertTimeToWarpScript(args[1])
	if err != nil {
		return nil, err
	}

	op := "sum"
	if len(args) >= 3 {
		op = args[2]
	}

	kwargs["span"] = *t
	node, err = bucketize(node, []string{args[0], op}, kwargs)
	if err != nil {
		return nil, err
	}

	if args[0] != swap {
		return fetch(node, []string{args[0], kwargs["from"], kwargs["until"]}, kwargs)
	}

	return node, nil
}
