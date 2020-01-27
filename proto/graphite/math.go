package graphite

import (
	"errors"

	"github.com/ovh/erlenmeyer/core"
)

// ----------------------------------------------------------------------------
// graphite functions implementations

func absolute(node *core.Node, args []string, kwargs map[string]string) (*core.Node, error) {
	if len(args) < 1 {
		return nil, errors.New("The absolute function take one parameter which is a series or a list of series")
	}

	node.Left = core.NewNode(core.MapperPayload{
		Mapper:      "abs",
		PostWindow:  "0",
		PreWindow:   "0",
		Occurrences: "0",
	})

	if args[0] != swap {
		return fetch(node.Left, []string{args[0], kwargs["from"], kwargs["until"]}, kwargs)
	}

	return node.Left, nil
}

func integral(node *core.Node, args []string, kwargs map[string]string) (*core.Node, error) {
	if len(args) < 1 {
		return nil, errors.New("The integral function take one parameter which are a series or a list of series")
	}

	node.Left = core.NewNode(core.WarpScriptPayload{
		WarpScript: "0 INTEGRATE",
	})

	if args[0] != swap {
		return fetch(node.Left, []string{args[0], kwargs["from"], kwargs["until"]}, kwargs)
	}

	return node.Left, nil
}

func interpolate(node *core.Node, args []string, kwargs map[string]string) (*core.Node, error) {
	if len(args) < 1 {
		return nil, errors.New("The interpolate function take two parameters which are a series or a list of series and optionally a number")
	}

	node.Left = core.NewNode(core.WarpScriptPayload{
		WarpScript: "INTERPOLATE SORT",
	})

	if args[0] != swap {
		return fetch(node.Left, []string{args[0], kwargs["from"], kwargs["until"]}, kwargs)
	}

	return node.Left, nil
}

func logarithm(node *core.Node, args []string, kwargs map[string]string) (*core.Node, error) {
	if len(args) < 1 {
		return nil, errors.New("The logarithm function take at list one parameter which is a list of series and optionally a number")
	}

	constant := "10"
	if len(args) < 2 {
		constant = args[1]
	}

	node.Left = core.NewNode(core.MapperPayload{
		Constant:    constant,
		Mapper:      "log",
		Occurrences: "0",
		PostWindow:  "0",
		PreWindow:   "0",
	})

	if args[0] != swap {
		return fetch(node.Left, []string{args[0], kwargs["from"], kwargs["until"]}, kwargs)
	}

	return node.Left, nil
}

func minMax(node *core.Node, args []string, kwargs map[string]string) (*core.Node, error) {
	if len(args) < 1 {
		return nil, errors.New("The minMax function take two parameters which is a list of series")
	}

	node.Left = core.NewNode(core.WarpScriptPayload{
		WarpScript: "<% DROP NORMALIZE %> LMAP",
	})

	if args[0] != swap {
		return fetch(node.Left, []string{args[0], kwargs["from"], kwargs["until"]}, kwargs)
	}

	return node.Left, nil
}

func pow(node *core.Node, args []string, kwargs map[string]string) (*core.Node, error) {
	if len(args) < 2 {
		return nil, errors.New("The pow function take two parameters which is a list of series and a number")
	}

	node.Left = core.NewNode(core.MapperPayload{
		Mapper:      "pow",
		Constant:    args[1],
		Occurrences: "0",
		PostWindow:  "0",
		PreWindow:   "0",
	})

	if args[0] != swap {
		return fetch(node.Left, []string{args[0], kwargs["from"], kwargs["until"]}, kwargs)
	}

	return node.Left, nil
}

func squareRoot(node *core.Node, args []string, kwargs map[string]string) (*core.Node, error) {
	if len(args) < 1 {
		return nil, errors.New("The squareRoot function take one parameter which is a list of series")
	}

	node.Left = core.NewNode(core.MapperPayload{
		Constant:    "0.5",
		Mapper:      "pow",
		Occurrences: "0",
		PostWindow:  "0",
		PreWindow:   "0",
	})

	if args[0] != swap {
		return fetch(node.Left, []string{args[0], kwargs["from"], kwargs["until"]}, kwargs)
	}

	return node.Left, nil
}
