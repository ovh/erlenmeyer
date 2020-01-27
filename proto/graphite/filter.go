package graphite

import (
	"errors"
	"fmt"

	"github.com/ovh/erlenmeyer/core"
)

const (
	maximumAboveFilter = `
	NONEMPTY
[] SWAP
<%% DUP 
	[ SWAP bucketizer.max 0 0 1 ] BUCKETIZE
	[ SWAP %s mapper.gt 0 0 0 ] MAP VALUES FLATTEN SIZE
	<%% 0 == %%>
			<%% DROP %%>
			<%% + %%>
	IFTE
%%> FOREACH`

	maximumBelowFilter = `
	NONEMPTY
[] SWAP
<%% DUP 
	[ SWAP bucketizer.max 0 0 1 ] BUCKETIZE
	[ SWAP %s mapper.lt 0 0 0 ] MAP VALUES FLATTEN SIZE
	<%% 0 == %%>
			<%% DROP %%>
			<%% + %%>
	IFTE
%%> FOREACH`

	minimumAboveFilter = `
	NONEMPTY
[] SWAP
<%% DUP 
	[ SWAP bucketizer.min 0 0 1 ] BUCKETIZE
	[ SWAP %s mapper.gt 0 0 0 ] MAP VALUES FLATTEN SIZE
	<%% 0 == %%>
			<%% DROP %%>
			<%% + %%>
	IFTE
%%> FOREACH`

	minimumBelowFilter = `
	NONEMPTY
[] SWAP
<%% DUP 
	[ SWAP bucketizer.min 0 0 1 ] BUCKETIZE
	[ SWAP %s mapper.lt 0 0 0 ] MAP VALUES FLATTEN SIZE
	<%% 0 == %%>
			<%% DROP %%>
			<%% + %%>
	IFTE
%%> FOREACH`

	highestAverageFilter = `
	NONEMPTY
<%% [ SWAP bucketizer.mean 0 0 1 ] BUCKETIZE VALUES FLATTEN 0 GET %%> SORTBY REVERSE
<%% DUP SIZE 0 != %%> <%% [ 0 %s 1 - ] SUBLIST %%> IFT
`

	highestCurrentFilter = `
	NONEMPTY
<%% VALUES FLATTEN 0 GET %%> SORTBY REVERSE
<%% DUP SIZE 0 != %%> <%% [ 0 %s 1 - ] SUBLIST %%> IFT
`

	highestMaxFilter = `
	NONEMPTY
<%% VALUES FLATTEN LSORT REVERSE 0 GET %%> SORTBY REVERSE
<%% DUP SIZE 0 != %%> <%% [ 0 %s 1 - ] SUBLIST %%> IFT
`

	lowestAverageFilter = `
	NONEMPTY
<%% [ SWAP bucketizer.mean 0 0 1 ] BUCKETIZE VALUES FLATTEN 0 GET %%> SORTBY
<%% DUP SIZE 0 != %%> <%% [ 0 %s 1 - ] SUBLIST %%> IFT
`

	lowestCurrentFilter = `
	NONEMPTY
<%% VALUES FLATTEN 0 GET %%> SORTBY
<%% DUP SIZE 0 != %%> <%% [ 0 %s 1 - ] SUBLIST %%> IFT
`

	averageAboveFilter = `
	NONEMPTY
    []
    SWAP
<%% DUP
	[ SWAP bucketizer.mean 0 0 1 ] BUCKETIZE
	<%% VALUES FLATTEN 0 GET %s < %%>
		<%% DROP %%>
		<%% + %%> 
	IFTE
	%%>  FOREACH
`

	averageBelowFilter = `
	NONEMPTY
    []
    SWAP
<%% DUP
	[ SWAP bucketizer.mean 0 0 1 ] BUCKETIZE
	<%% VALUES FLATTEN 0 GET %s > %%>
		<%% DROP %%>
		<%% + %%> 
	IFTE
%%>  FOREACH
`
)

// ----------------------------------------------------------------------------
// graphite functions implementations

func averageAbove(node *core.Node, args []string, kwargs map[string]string) (*core.Node, error) {
	if len(args) < 2 {
		return nil, errors.New("The averageAbove function take two parameters which are series and a number")
	}

	node.Left = core.NewNode(core.WarpScriptPayload{
		WarpScript: fmt.Sprintf(averageAboveFilter, args[1]),
	})

	if args[0] != swap {
		return fetch(node.Left, []string{args[0], kwargs["from"], kwargs["until"]}, kwargs)
	}

	return node.Left, nil
}

func averageBelow(node *core.Node, args []string, kwargs map[string]string) (*core.Node, error) {
	if len(args) < 2 {
		return nil, errors.New("The averageBelow function take two parameters which are series and a number")
	}

	node.Left = core.NewNode(core.WarpScriptPayload{
		WarpScript: fmt.Sprintf(averageBelowFilter, args[1]),
	})

	if args[0] != swap {
		return fetch(node.Left, []string{args[0], kwargs["from"], kwargs["until"]}, kwargs)
	}

	return node.Left, nil
}

func exclude(node *core.Node, args []string, kwargs map[string]string) (*core.Node, error) {
	if len(args) < 2 {
		return nil, errors.New("The exclude function take two parameters which  a list of series and a regexp pattern")
	}

	node.Left = core.NewNode(core.WarpScriptPayload{
		WarpScript: fmt.Sprintf("[ SWAP [] '~(?!%s).*' filter.byclass ] FILTER", args[1]),
	})

	if args[0] != swap {
		return fetch(node.Left, []string{args[0], kwargs["from"], kwargs["until"]}, kwargs)
	}

	return node.Left, nil
}

func grep(node *core.Node, args []string, kwargs map[string]string) (*core.Node, error) {
	if len(args) < 2 {
		return nil, errors.New("The grep function take two parameters which a list of series and a regexp pattern")
	}

	node.Left = core.NewNode(core.WarpScriptPayload{
		WarpScript: fmt.Sprintf("[ SWAP [] '~%s' filter.byclass ] FILTER", args[1]),
	})

	if args[0] != swap {
		return fetch(node.Left, []string{args[0], kwargs["from"], kwargs["until"]}, kwargs)
	}

	return node.Left, nil
}

func maximumAbove(node *core.Node, args []string, kwargs map[string]string) (*core.Node, error) {
	if len(args) < 2 {
		return nil, errors.New("The minimumAbove function take two parameters which is a list of series and a number")
	}

	node.Left = core.NewNode(core.WarpScriptPayload{
		WarpScript: fmt.Sprintf(maximumAboveFilter, args[1]),
	})

	if args[0] != swap {
		return fetch(node.Left, []string{args[0], kwargs["from"], kwargs["until"]}, kwargs)
	}

	return node.Left, nil
}

func maximumBelow(node *core.Node, args []string, kwargs map[string]string) (*core.Node, error) {
	if len(args) < 2 {
		return nil, errors.New("The minimumAbove function take two parameters which is a list of series and a number")
	}

	node.Left = core.NewNode(core.WarpScriptPayload{
		WarpScript: fmt.Sprintf(maximumBelowFilter, args[1]),
	})

	if args[0] != swap {
		return fetch(node.Left, []string{args[0], kwargs["from"], kwargs["until"]}, kwargs)
	}

	return node.Left, nil
}

func minimumAbove(node *core.Node, args []string, kwargs map[string]string) (*core.Node, error) {
	if len(args) < 2 {
		return nil, errors.New("The minimumAbove function take two parameters which is a list of series and a number")
	}

	node.Left = core.NewNode(core.WarpScriptPayload{
		WarpScript: fmt.Sprintf(minimumAboveFilter, args[1]),
	})

	if args[0] != swap {
		return fetch(node.Left, []string{args[0], kwargs["from"], kwargs["until"]}, kwargs)
	}

	return node.Left, nil
}

func minimumBelow(node *core.Node, args []string, kwargs map[string]string) (*core.Node, error) {
	if len(args) < 2 {
		return nil, errors.New("The minimumAbove function take two parameters which is a list of series and a number")
	}

	node.Left = core.NewNode(core.WarpScriptPayload{
		WarpScript: fmt.Sprintf(minimumBelowFilter, args[1]),
	})

	if args[0] != swap {
		return fetch(node.Left, []string{args[0], kwargs["from"], kwargs["until"]}, kwargs)
	}

	return node.Left, nil
}

func removeAboveValue(node *core.Node, args []string, kwargs map[string]string) (*core.Node, error) {
	if len(args) < 2 {
		return nil, errors.New("The removeAboveValue function take two parameters which are a list of series and a number")
	}

	node.Left = core.NewNode(core.MapperPayload{
		Constant:    args[1],
		Mapper:      "lt",
		Occurrences: "0",
		PostWindow:  "0",
		PreWindow:   "0",
	})

	if args[0] != swap {
		return fetch(node.Left, []string{args[0], kwargs["from"], kwargs["until"]}, kwargs)
	}

	return node.Left, nil
}

func removeBelowValue(node *core.Node, args []string, kwargs map[string]string) (*core.Node, error) {
	if len(args) < 2 {
		return nil, errors.New("The removeAboveValue function take two parameters which are a list of series and a number")
	}

	node.Left = core.NewNode(core.MapperPayload{
		Constant:    args[1],
		Mapper:      "gt",
		Occurrences: "0",
		PostWindow:  "0",
		PreWindow:   "0",
	})

	if args[0] != swap {
		return fetch(node.Left, []string{args[0], kwargs["from"], kwargs["until"]}, kwargs)
	}

	return node.Left, nil
}

func removeEmptySeries(node *core.Node, args []string, kwargs map[string]string) (*core.Node, error) {
	if len(args) < 1 {
		return nil, errors.New("The removeEmptySeries function take one parameter which is a list of series")
	}

	node.Left = core.NewNode(core.WarpScriptPayload{
		WarpScript: "NONEMPTY",
	})

	if args[0] != swap {
		return fetch(node.Left, []string{args[0], kwargs["from"], kwargs["until"]}, kwargs)
	}

	return node.Left, nil
}

func limit(node *core.Node, args []string, kwargs map[string]string) (*core.Node, error) {
	if len(args) < 2 {
		return nil, errors.New("The limit function take two parameters which are a list of series and a number")
	}

	node.Left = core.NewNode(core.WarpScriptPayload{
		WarpScript: fmt.Sprintf(`<%% DUP SIZE 0 > %%> <%% [ 0 %s 1 - ] SUBLIST %%> IFT`, args[1]),
	})

	if args[0] != swap {
		return fetch(node.Left, []string{args[0], kwargs["from"], kwargs["until"]}, kwargs)
	}

	return node.Left, nil
}

func currentAbove(node *core.Node, args []string, kwargs map[string]string) (*core.Node, error) {
	if len(args) < 2 {
		return nil, errors.New("The currentAbove take two parameters which are a series and number")
	}

	node.Left = core.NewNode(core.WarpScriptPayload{
		WarpScript: fmt.Sprintf("[ SWAP [] %s filter.last.gt ] FILTER", args[1]),
	})

	if args[0] != swap {
		return fetch(node.Left, []string{args[0], kwargs["from"], kwargs["until"]}, kwargs)
	}

	return node.Left, nil
}

func currentBelow(node *core.Node, args []string, kwargs map[string]string) (*core.Node, error) {
	if len(args) < 2 {
		return nil, errors.New("The currentBelow take two parameters which are a series and a number")
	}

	node.Left = core.NewNode(core.WarpScriptPayload{
		WarpScript: fmt.Sprintf("[ SWAP [] %s filter.last.lt ] FILTER", args[1]),
	})

	if args[0] != swap {
		return fetch(node.Left, []string{args[0], kwargs["from"], kwargs["until"]}, kwargs)
	}

	return node.Left, nil
}

func highestAverage(node *core.Node, args []string, kwargs map[string]string) (*core.Node, error) {
	if len(args) < 2 {
		return nil, errors.New("The highestAverage take two parameters which are a list of series and a number")
	}

	node.Left = core.NewNode(core.WarpScriptPayload{
		WarpScript: fmt.Sprintf(highestAverageFilter, args[1]),
	})

	if args[0] != swap {
		return fetch(node.Left, []string{args[0], kwargs["from"], kwargs["until"]}, kwargs)
	}

	return node.Left, nil
}

func highestCurrent(node *core.Node, args []string, kwargs map[string]string) (*core.Node, error) {
	if len(args) < 2 {
		return nil, errors.New("The highestCurrent take two parameters which are a list of series and a number")
	}

	node.Left = core.NewNode(core.WarpScriptPayload{
		WarpScript: fmt.Sprintf(highestCurrentFilter, args[1]),
	})

	if args[0] != swap {
		return fetch(node.Left, []string{args[0], kwargs["from"], kwargs["until"]}, kwargs)
	}

	return node.Left, nil
}

func highestMax(node *core.Node, args []string, kwargs map[string]string) (*core.Node, error) {
	if len(args) < 2 {
		return nil, errors.New("The highestMax take two parameters which are a list of series and a number")
	}

	node.Left = core.NewNode(core.WarpScriptPayload{
		WarpScript: fmt.Sprintf(highestMaxFilter, args[1]),
	})

	if args[0] != swap {
		return fetch(node.Left, []string{args[0], kwargs["from"], kwargs["until"]}, kwargs)
	}

	return node.Left, nil
}

func lowestAverage(node *core.Node, args []string, kwargs map[string]string) (*core.Node, error) {
	if len(args) < 2 {
		return nil, errors.New("The lowestAverage take two parameters which are a list of series and a number")
	}

	node.Left = core.NewNode(core.WarpScriptPayload{
		WarpScript: fmt.Sprintf(lowestAverageFilter, args[1]),
	})

	if args[0] != swap {
		return fetch(node.Left, []string{args[0], kwargs["from"], kwargs["until"]}, kwargs)
	}

	return node.Left, nil
}

func lowestCurrent(node *core.Node, args []string, kwargs map[string]string) (*core.Node, error) {
	if len(args) < 2 {
		return nil, errors.New("The lowestCurrent take two parameters which are a list of series and a number")
	}

	node.Left = core.NewNode(core.WarpScriptPayload{
		WarpScript: fmt.Sprintf(lowestCurrentFilter, args[1]),
	})

	if args[0] != swap {
		return fetch(node.Left, []string{args[0], kwargs["from"], kwargs["until"]}, kwargs)
	}

	return node.Left, nil
}
