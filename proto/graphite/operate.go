package graphite

import (
	"errors"

	"github.com/ovh/erlenmeyer/core"
)

// ----------------------------------------------------------------------------
// graphite functions implementations

func transformNull(node *core.Node, args []string, kwargs map[string]string) (*core.Node, error) {
	if len(args) < 1 {
		return nil, errors.New("The transformNull function take at least one parameter which is a list of series")
	}

	val := "0"
	if len(args) >= 2 {
		val = args[1]
	}

	node.Left = core.NewNode(core.FillValuePayload{
		Elevation: "NaN",
		Latitude:  "NaN",
		Longitude: "NaN",
		Value:     val,
	})

	if args[0] != swap {
		return fetch(node.Left, []string{args[0], kwargs["from"], kwargs["until"]}, kwargs)
	}

	return node.Left, nil
}

func keepLastValue(node *core.Node, args []string, kwargs map[string]string) (*core.Node, error) {
	if len(args) < 1 {
		return nil, errors.New("The keepLastValue function take at least one parameter which is a list of series")
	}

	node.Left = core.NewNode(core.WarpScriptPayload{
		WarpScript: "FILLPREVIOUS SORT",
	})

	if args[0] != swap {
		return fetch(node.Left, []string{args[0], kwargs["from"], kwargs["until"]}, kwargs)
	}

	return node.Left, nil
}
