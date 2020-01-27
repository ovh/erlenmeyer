package graphite

import (
	"errors"

	"github.com/ovh/erlenmeyer/core"
)

// ----------------------------------------------------------------------------
// graphite functions implementations

func unique(node *core.Node, args []string, kwargs map[string]string) (*core.Node, error) {
	if len(args) < 1 {
		return nil, errors.New("The unique function take a list of series in parameters")
	}

	node.Left = core.NewNode(core.WarpScriptPayload{
		WarpScript: "<% NAME %> SORTBY UNIQUE",
	})

	if args[0] != swap {
		return fetch(node.Left, []string{args[0], kwargs["from"], kwargs["until"]}, kwargs)
	}

	return node.Left, nil
}

func sortByMaxima(node *core.Node, args []string, kwargs map[string]string) (*core.Node, error) {
	if len(args) < 1 {
		return nil, errors.New("The sortByMaxima function take one parameter which is a list of series")
	}

	node.Left = core.NewNode(core.WarpScriptPayload{
		WarpScript: "NONEMPTY <% [ [ ROT ] bucketizer.max 0 0 1 ] BUCKETIZE VALUES FLATTEN 0 GET %> SORTBY",
	})

	if args[0] != swap {
		return fetch(node.Left, []string{args[0], kwargs["from"], kwargs["until"]}, kwargs)
	}

	return node.Left, nil
}

func sortByMinima(node *core.Node, args []string, kwargs map[string]string) (*core.Node, error) {
	if len(args) < 1 {
		return nil, errors.New("The sortByMinima function take one parameter which is a list of series")
	}

	node.Left = core.NewNode(core.WarpScriptPayload{
		WarpScript: "NONEMPTY <% [ [ ROT ] bucketizer.max 0 0 1 ] BUCKETIZE VALUES FLATTEN 0 GET %> SORTBY REVERSE",
	})

	if args[0] != swap {
		return fetch(node.Left, []string{args[0], kwargs["from"], kwargs["until"]}, kwargs)
	}

	return node.Left, nil
}

func sortByName(node *core.Node, args []string, kwargs map[string]string) (*core.Node, error) {
	if len(args) < 1 {
		return nil, errors.New("The sortByName function take one parameter which is a list of series")
	}

	node.Left = core.NewNode(core.WarpScriptPayload{
		WarpScript: "<% NAME %> SORTBY",
	})

	if args[0] != swap {
		return fetch(node.Left, []string{args[0], kwargs["from"], kwargs["until"]}, kwargs)
	}

	return node.Left, nil
}

func sortByTotal(node *core.Node, args []string, kwargs map[string]string) (*core.Node, error) {
	if len(args) < 1 {
		return nil, errors.New("The sortByTotal function take one parameter which is a list of series")
	}

	node.Left = core.NewNode(core.WarpScriptPayload{
		WarpScript: "NONEMPTY <% [ [ ROT ] bucketizer.sum 0 0 1 ] BUCKETIZE VALUES FLATTEN 0 GET %> SORTBY",
	})

	if args[0] != swap {
		return fetch(node.Left, []string{args[0], kwargs["from"], kwargs["until"]}, kwargs)
	}

	return node.Left, nil
}
