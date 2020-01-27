package graphite

import (
	"errors"
	"fmt"
	"strconv"
	"strings"

	"github.com/ovh/erlenmeyer/core"
)

const (
	constantYield = `
[
	%s 'end' STORE
	60 s 'span' STORE
	%s $span - 'timestamp' STORE
	%s 'val' STORE

	NEWGTS '%s' RENAME
	<%% $timestamp $span + $end <= %%>
		<%%
				$timestamp $span + 'timestamp' STORE
				$timestamp NaN NaN NaN $val ADDVALUE
		%%>
	WHILE
]`

	timeYield = `
[
	%s 'end' STORE
	%s 'span' STORE
	%s $span - 'timestamp' STORE

	NEWGTS '%s' RENAME
	<%% $timestamp $span + $end <= %%>
		<%%
				$timestamp $span + 'timestamp' STORE
				$timestamp NaN NaN NaN $timestamp 1000 / ADDVALUE
		%%>
	WHILE
]`

	randomYield = `
[
	%s 'end' STORE
	%s 'span' STORE
	%s $span - 'timestamp' STORE
	0.0 'last' STORE

	NEWGTS '%s' RENAME
	$timestamp NaN NaN NaN $last ADDVALUE
	[ $timestamp $end ] [] [] [] [ 0.0 DUP ] MAKEGTS
	[ SWAP bucketizer.last $end $span 0 ] BUCKETIZE INTERPOLATE SORT 0 GET 
	TICKLIST 0 REMOVE DROP
	<%% 
		'timestamp' STORE
		$last RAND 0.5 - + 'val' STORE
		$val 'last' STORE
		$timestamp NaN NaN NaN $val  ADDVALUE
	%%>
	FOREACH
]`

	sinYield = `
[
	%s 'end' STORE
	%s 'span' STORE
	%s $span - 'timestamp' STORE
	%s 'amplitude' STORE

	NEWGTS '%s' RENAME
	[ $timestamp $end ] [] [] [] [ 0.0 DUP ] MAKEGTS
	[ SWAP bucketizer.last $end $span 0 ] BUCKETIZE INTERPOLATE SORT 0 GET 
	TICKLIST
	<%% 
		'timestamp' STORE
		$timestamp $span + 'timestamp' STORE
		$timestamp NaN NaN NaN $timestamp SIN $amplitude *  ADDVALUE
	%%>
	FOREACH
]`
)

// ----------------------------------------------------------------------------
// graphite functions implementations

func constantLine(node *core.Node, args []string, kwargs map[string]string) (*core.Node, error) {
	if len(args) != 1 {
		return nil, fmt.Errorf("Error with constantLine function, expecting only one number parameter")
	}

	if !isNumeric(args[0]) && !strings.HasPrefix(args[0], "$") {
		return nil, fmt.Errorf("Error with constantLine function, expecting a number as only parameter and got '%s' ", args[0])
	}

	node.Left = core.NewNode(core.WarpScriptPayload{
		WarpScript: fmt.Sprintf(constantYield, kwargs["until"], kwargs["from"], args[0], args[0]),
	})

	return node.Left, nil
}

// isNumeric Internal check to test if a Graphite string parameter is a numeric value
func isNumeric(s string) bool {
	_, err := strconv.ParseFloat(s, 64)
	return err == nil
}

func threshold(node *core.Node, args []string, kwargs map[string]string) (*core.Node, error) {
	var err error

	if len(args) < 1 {
		return nil, errors.New("the threshold function take at least one parameter which is a threshold number")
	}

	if len(args) >= 2 {
		node.Left = core.NewNode(core.WarpScriptPayload{
			WarpScript: fmt.Sprintf("'%s' RENAME", args[1]),
		})

		node = node.Left
	}

	node, err = constantLine(node, []string{args[0]}, kwargs)
	if err != nil {
		return nil, err
	}

	return node, nil
}

func sinFunction(node *core.Node, args []string, kwargs map[string]string) (*core.Node, error) {
	if len(args) < 1 {
		return nil, errors.New("The sin function take one parameter which is the name of the series")
	}

	span := "60"
	if len(args) >= 3 {
		span = args[2]
	}

	span = parseDuration(span)

	amplitude := "1"
	if len(args) >= 2 {
		amplitude = args[1]
	}

	node.Left = core.NewNode(core.WarpScriptPayload{
		WarpScript: fmt.Sprintf(sinYield, kwargs["until"], span, kwargs["from"], amplitude, args[0]),
	})

	return node.Left, nil
}

func parseDuration(duration string) string {
	twoEnd := duration[len(duration)-2:]
	end := duration[len(duration)-1:]

	switch twoEnd {
	case "ms", "us", "ns", "ps":
		return getUnit(duration, twoEnd)
	}

	switch end {
	case "w", "d", "h", "m", "s":
		return getUnit(duration, end)
	}

	return duration + " s"
}

func getUnit(duration string, unit string) string {
	return strings.TrimRight(duration, unit) + " " + unit
}

func randomWalkFunction(node *core.Node, args []string, kwargs map[string]string) (*core.Node, error) {
	if len(args) < 1 {
		return nil, errors.New("The time function take one parameter which is the name of the series")
	}

	span := "60" // nolint: goconst
	if len(args) >= 2 {
		span = args[1]
	}

	span = parseDuration(span)

	node.Left = core.NewNode(core.WarpScriptPayload{
		WarpScript: fmt.Sprintf(randomYield, kwargs["until"], span, kwargs["from"], args[0]),
	})

	return node.Left, nil
}

func timeFunction(node *core.Node, args []string, kwargs map[string]string) (*core.Node, error) {
	if len(args) < 1 {
		return nil, errors.New("The time function take one parameter which is the name of the series")
	}

	span := "60" // nolint: goconst
	if len(args) >= 2 {
		span = args[1]
	}
	span = parseDuration(span)

	warpScript := `
	[ ` + kwargs["from"] + ` ` + kwargs["until"] + ` DUP 'timeFunctionEnd' STORE ] [] [] [] [ 1 DUP ] MAKEGTS '` + args[0] + `' RENAME
 	[ SWAP bucketizer.last $timeFunctionEnd ` + span + ` 0 ] BUCKETIZE INTERPOLATE SORT
	[ SWAP mapper.tick 0 0 0 ] MAP [ SWAP 0.000001 mapper.mul 0 0 0 ] MAP [ SWAP mapper.floor 0 0 0 ] MAP`

	node.Left = core.NewNode(core.WarpScriptPayload{
		WarpScript: warpScript,
	})

	//node.Left = core.NewNode(core.WarpScriptPayload{
	//	WarpScript: fmt.Sprintf(timeYield, kwargs["until"], span, kwargs["from"], args[0]),
	//})

	return node.Left, nil
}
