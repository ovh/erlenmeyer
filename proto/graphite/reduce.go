package graphite

import (
	"errors"

	"github.com/ovh/erlenmeyer/core"
)

// ----------------------------------------------------------------------------
// graphite functions implementations

func countSeries(node *core.Node, args []string, kwargs map[string]string) (*core.Node, error) {
	if len(args) < 1 {
		return nil, errors.New("The countSeries function take one parameter which is a series or a list of series")
	}

	node, err := constantLine(node, []string{"$countSeries"}, kwargs)
	if err != nil {
		return nil, err
	}

	node.Left = core.NewNode(core.WarpScriptPayload{
		WarpScript: "SIZE 'countSeries' STORE",
	})

	if args[0] != swap {
		return find(node.Left, args, kwargs)
	}

	return node.Left, nil
}
