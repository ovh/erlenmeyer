package graphite

import (
	"errors"
	"fmt"
	"strconv"
	"strings"

	"github.com/ovh/erlenmeyer/core"
)

const (
	substrLmap = `
	<%%
		DROP DUP 
	
		NAME '.' SPLIT 'name' STORE
		%d 'start' STORE
		%d 'end' STORE
		<%% $end 0 <= %%> <%% $name SIZE $end + %%> <%% $end %%> IFTE 'end' STORE 
	
		$name [ $start $end ] SUBLIST LIST-> '.' SWAP JOIN
	
		RENAME      
	%%> LMAP`

	labelToSerie = "SWAP DUP LABELS '%s' GET ROT SWAP +"
	replaceLmap  = "<%% DROP DUP NAME '%s' '%s' REPLACE RENAME %%> LMAP"
)

// ----------------------------------------------------------------------------
// helper functions

func name(node *core.Node, args []string, kwargs map[string]string) (*core.Node, error) {
	if len(args) < 1 {
		return nil, errors.New("The name function take a list of series in parameters")
	}

	node.Left = core.NewNode(core.WarpScriptPayload{
		WarpScript: "<% DROP NAME %> LMAP UNIQUE",
	})

	if args[0] != swap {
		return fetch(node.Left, []string{args[0], kwargs["from"], kwargs["until"]}, kwargs)
	}

	return node.Left, nil
}

// ----------------------------------------------------------------------------
// graphite functions implementations

func alias(node *core.Node, args []string, kwargs map[string]string) (*core.Node, error) {
	if len(args) < 2 {
		return nil, errors.New("The alias function take two parameters which are a series and the new name")
	}

	node.Left = core.NewNode(core.WarpScriptPayload{
		WarpScript: fmt.Sprintf("'%s' RENAME", args[1]),
	})

	if args[0] != swap {
		return fetch(node.Left, []string{args[0], kwargs["from"], kwargs["until"]}, kwargs)
	}

	return node.Left, nil
}

func aliasSub(node *core.Node, args []string, kwargs map[string]string) (*core.Node, error) {
	if len(args) < 3 {
		return nil, errors.New("The aliasSub function take three parameters which are a list of series, a pattern, and a replace string")
	}

	node.Left = core.NewNode(core.WarpScriptPayload{
		WarpScript: fmt.Sprintf(replaceLmap, args[1], args[2]),
	})

	if args[0] != swap {
		return fetch(node.Left, []string{args[0], kwargs["from"], kwargs["until"]}, kwargs)
	}

	return node.Left, nil
}

func aliasByNode(node *core.Node, args []string, kwargs map[string]string) (*core.Node, error) {
	if len(args) < 2 {
		return nil, errors.New("The aliasByNode function take at least two parameters which are a list of series and parameters to rewrite metric")
	}

	last := len(args) - 1
	index, err := strconv.Atoi(args[last])
	if err != nil {
		index = 0
		last = len(args)
	}

	labels := make([]string, 0)
	for _, label := range args[1:last] {
		labels = append(labels, fmt.Sprintf(labelToSerie, label))
	}

	node.Left = core.NewNode(core.WarpScriptPayload{
		WarpScript: fmt.Sprintf(mapperLabelToSeriesIndex, index, strings.Join(labels, "\n")),
	})

	if args[0] != swap {
		return fetch(node.Left, []string{args[0], kwargs["from"], kwargs["until"]}, kwargs)
	}

	return node.Left, nil
}

func aliasByTags(node *core.Node, args []string, kwargs map[string]string) (*core.Node, error) {
	if len(args) < 2 {
		return nil, errors.New("The aliasByTags function take at least two parameters which are a list of series and parameters to rewrite metric")
	}

	labels := make([]string, 0)
	for _, label := range args[1:] {
		labels = append(labels, fmt.Sprintf("'%s'", label))
	}

	node.Left = core.NewNode(core.WarpScriptPayload{
		WarpScript: fmt.Sprintf(mapperLabelToSeries, strings.Join(labels, " ")),
	})

	if args[0] != swap {
		return fetch(node.Left, []string{args[0], kwargs["from"], kwargs["until"]}, kwargs)
	}

	return node.Left, nil
}

func aliasByMetric(node *core.Node, args []string, kwargs map[string]string) (*core.Node, error) {
	if len(args) < 1 {
		return nil, errors.New("The aliasByMetric function take one parameter which is a series or a list of series")
	}

	node, err := substr(node, []string{swap, "-1"}, kwargs)
	if err != nil {
		return nil, err
	}

	if args[0] != swap {
		return fetch(node, []string{args[0], kwargs["from"], kwargs["until"]}, kwargs)
	}

	return node, nil
}

func substr(node *core.Node, args []string, kwargs map[string]string) (*core.Node, error) {
	var err error

	if len(args) < 2 {
		return nil, errors.New("The substr function take three parameters which is a series, a start number and optionally a stop number")
	}

	start := 0
	end := -1
	if len(args) > 2 {
		end, err = strconv.Atoi(args[2])
		if err != nil {
			return nil, err
		}

		start, err = strconv.Atoi(args[1])
		if err != nil {
			return nil, err
		}
	} else if len(args) > 1 {
		end, err = strconv.Atoi(args[1])
		if err != nil {
			return nil, err
		}
	}

	node.Left = core.NewNode(core.WarpScriptPayload{
		WarpScript: fmt.Sprintf(substrLmap, start, end),
	})

	if args[0] != swap {
		return fetch(node.Left, []string{args[0], kwargs["from"], kwargs["until"]}, kwargs)
	}

	return node.Left, nil
}
