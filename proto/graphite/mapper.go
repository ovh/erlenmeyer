package graphite

import (
	"errors"

	"github.com/ovh/erlenmeyer/core"
)

const (
	mapperLabelToSeries = `
<%% 
		DROP
		'series' STORE
		
		[] 'newName' STORE

		// alias params list
		[ %s ] 
		<%% 
		
			'item' STORE
			<%% $item 'name' == %%>
			<%%
				$newName $series NAME + 
				'newName' STORE
				CONTINUE
			%%>
			IFT
			<%% $item '^[0-9]+$' MATCH SIZE 0 > %%>
			<%%
				
				// Add splitted name by dot item number
				$newName $series NAME '.' SPLIT 
				
				// Manage case where items is superior to series name size
				<%% DUP SIZE $item TOLONG <=  %%>
				<%% CLEAR CONTINUE %%>
				IFT
				
				$item TOLONG GET + 
				'newName' STORE
				CONTINUE
			%%>
			IFT
			<%% 
				$series LABELS $item CONTAINSKEY SWAP DROP
			%%>
			<%%
				// Add series label value
				$newName $series LABELS $item GET + 
				'newName' STORE
				CONTINUE
			%%>
			IFT
		%%> 
		FOREACH
		
		// Set by default an empty name
		<%% 
			$newName SIZE 0 == 
		%%>
		<%%
			$newName '' +
			'newName' STORE
		%%>
		IFT
		$series $newName '.' JOIN RENAME
		{ NULL NULL } RELABEL
%%> LMAP`
)

// ----------------------------------------------------------------------------
// graphite functions implementations

func derivative(node *core.Node, args []string, kwargs map[string]string) (*core.Node, error) {
	if len(args) < 1 {
		return nil, errors.New("The derivative function take one parameter which a list of series")
	}

	node.Left = core.NewNode(core.MapperPayload{
		Mapper:      "delta",
		PostWindow:  "1",
		PreWindow:   "0",
		Occurrences: "0",
	})

	if args[0] != swap {
		return fetch(node.Left, []string{args[0], kwargs["from"], kwargs["until"]}, kwargs)
	}

	return node.Left, nil
}

func invert(node *core.Node, args []string, kwargs map[string]string) (*core.Node, error) {
	if len(args) < 1 {
		return nil, errors.New("The invert function take one parameter which is a series or a list of series")
	}

	warpScript := `
	// As only replacing series values, use list index as HASH
	<% { 'hash_945fa9bc3027d7025e3' ROT TOSTRING } RELABEL  %> LMAP
	DUP
	[ SWAP 0 TODOUBLE mapper.eq 0 0 0 ] MAP 'zero' STORE
	[ SWAP 0 TODOUBLE mapper.ne 0 0 0 ] MAP
	'toDiv' STORE
	[
			[ $toDiv 1 TODOUBLE mapper.replace 0 0 0 ] MAP 'dividende' RENAME
			$toDiv
			[ 'hash_945fa9bc3027d7025e3' ]
			op.div
	]
	APPLY
	[ SWAP $zero APPEND  [ 'hash_945fa9bc3027d7025e3' ] reducer.sum ] REDUCE 
	NONEMPTY { 'hash_945fa9bc3027d7025e3' '' } RELABEL
	`
	node.Left = core.NewNode(core.WarpScriptPayload{
		WarpScript: warpScript,
	})

	if args[0] != swap {
		return fetch(node.Left, []string{args[0], kwargs["from"], kwargs["until"]}, kwargs)
	}

	return node.Left, nil
}

func offset(node *core.Node, args []string, kwargs map[string]string) (*core.Node, error) {
	if len(args) < 2 {
		return nil, errors.New("The offset function take two parameters which is a list of series and a number")
	}

	node.Left = core.NewNode(core.MapperPayload{
		Mapper:      "add",
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

func perSecond(node *core.Node, args []string, kwargs map[string]string) (*core.Node, error) {
	if len(args) < 1 {
		return nil, errors.New("The perSecond function take one parameter which is a list of series")
	}

	node.Left = core.NewNode(core.MapperPayload{
		Mapper:      "rate",
		Occurrences: "1",
		PostWindow:  "0",
		PreWindow:   "0",
	})

	if args[0] != swap {
		return fetch(node.Left, []string{args[0], kwargs["from"], kwargs["until"]}, kwargs)
	}

	return node.Left, nil
}

func scale(node *core.Node, args []string, kwargs map[string]string) (*core.Node, error) {
	if len(args) < 2 {
		return nil, errors.New("The scale function take two parameters which is a list of series and a number")
	}

	node.Left = core.NewNode(core.MapperPayload{
		Mapper:      "mul",
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

func scaleToSeconds(node *core.Node, args []string, kwargs map[string]string) (*core.Node, error) {
	if len(args) < 1 {
		return nil, errors.New("The scaleToSeconds function take at least one parameter which is a list of series and optionally a number")
	}

	seconds := "1"
	if len(args) >= 2 {
		seconds = args[1]
	}

	node.Left = core.NewNode(core.MapperPayload{
		Mapper:      "mul",
		Constant:    "1.0 " + seconds + " TODOUBLE 60.0 * /",
		Occurrences: "0",
		PostWindow:  "0",
		PreWindow:   "0",
	})

	node = node.Left

	node.Left = core.NewNode(core.MapperPayload{
		Mapper:      "rate",
		Occurrences: "0",
		PostWindow:  "1",
		PreWindow:   "1",
	})

	if args[0] != swap {
		return fetch(node.Left, []string{args[0], kwargs["from"], kwargs["until"]}, kwargs)
	}

	return node.Left, nil
}

func divideSeries(node *core.Node, args []string, kwargs map[string]string) (*core.Node, error) {
	if len(args) < 2 {
		return nil, errors.New("The divideSeries function take two parameters which are a list of series and a number")
	}

	warpScript := `
		// Verify Right series 
		<%
			<% $right TYPEOF  'LIST' == %> 
			<% 
				// In case of a list expect it's size to be equal to 1
				$right SIZE 1 == 
				<% 
					DUP
				%>
				<%
					$right 0 GET '%2B.divisorSeries' RENAME 'right' STORE
				%>
				IFT

				// Let's equality inverse result on top of the stack to propagate an error
				!
			%> 
			IFT
		%>
		<%
			'divideSeries can be applied only on a singleton series as divisorSeries' MSGFAIL
		%>
		IFT

		// Then divide each "left" series per the right singleton series
		$left
		<%
			DROP
			'series' STORE
			[ $series CLONEEMPTY ] 
			[ 
				[ $series ]
				[ $right ]
				[]
				op.div
			] 
			APPLY
			APPEND
			MERGE
		%>
		LMAP
		`

	return divideCore(node, args, kwargs, warpScript)
}

func divideSeriesLists(node *core.Node, args []string, kwargs map[string]string) (*core.Node, error) {
	if len(args) < 2 {
		return nil, errors.New("The divideSeries function take two parameters which are a list of series and a number")
	}

	warpScript := `
			// Verify Left and Right series length
			<%
				$right SIZE
				$left SIZE !=
			%>
			<%
				'divideSeriesList need dividendSeriesList and divisorSeriesList to have the same length' MSGFAIL
			%>
			IFT

			// Divide all Left series per the Right one at the same index
			$left
			<%
				'index' STORE
				'series' STORE
				[ $series CLONEEMPTY ]
				[
					[ $series ]
					[ $right $index GET '%2B.divisorSeries' RENAME ]
					[]
					op.div
				]
				APPLY
				APPEND
				MERGE
			%>
			LMAP`

	return divideCore(node, args, kwargs, warpScript)
}

// divideCore manages all possible cases for args 0 (resp 1) being a fetch or intermediary query result
func divideCore(node *core.Node, args []string, kwargs map[string]string, warpScript string) (*core.Node, error) {
	var err error

	if args[1] != swap && args[0] != swap {
		kwargs["node"] = "true"

		prefix := "DROP"
		node.Left = core.NewNode(core.WarpScriptPayload{
			WarpScript: prefix + warpScript,
		})
		node.Left.Left, err = fetch(node.Left, []string{args[0], kwargs["from"], kwargs["until"]}, kwargs)
		if err != nil {
			return nil, err
		}

		node.Left.Right = core.NewEmptyNode()
		node.Left.Right.Left, err = fetch(node.Left.Right, []string{args[1], kwargs["from"], kwargs["until"]}, kwargs)

		if err != nil {
			return nil, err
		}
	} else if args[1] == swap && args[0] != swap {
		kwargs["node"] = "true"

		prefix := `
			DUP 0 GET 'right' STORE
			1 GET 'left' STORE`
		node.Left = core.NewNode(core.WarpScriptPayload{
			WarpScript: prefix + warpScript,
		})
		node.Left.Right, err = fetch(node.Left, []string{args[0], kwargs["from"], kwargs["until"]}, kwargs)

		if err != nil {
			return nil, err
		}
	} else if args[1] != swap && args[0] == swap {
		kwargs["node"] = "true"

		prefix := `
			DROP`
		node.Left = core.NewNode(core.WarpScriptPayload{
			WarpScript: prefix + warpScript,
		})
		node.Left.Right, err = fetch(node.Left, []string{args[1], kwargs["from"], kwargs["until"]}, kwargs)

		if err != nil {
			return nil, err
		}
	} else {
		prefix := `
		'right' STORE
		'left' STORE
		`
		node.Left = core.NewNode(core.WarpScriptPayload{
			WarpScript: prefix + warpScript,
		})
	}

	return node.Left, nil
}

func drawAsInfinite(node *core.Node, args []string, kwargs map[string]string) (*core.Node, error) {
	if len(args) < 1 {
		return nil, errors.New("The drawAsInfinite function take one parameter which is a list of series")
	}

	node.Left = core.NewNode(core.WarpScriptPayload{
		WarpScript: `	
			// As only replacing series values, use list index as HASH
			<% { 'hash_945fa9bc3027d7025e3' ROT TOSTRING } RELABEL  %> LMAP
			DUP
			[ SWAP 0 TODOUBLE mapper.eq 0 0 0 ] MAP 'zero' STORE
			[ SWAP 0 TODOUBLE mapper.ne 0 0 0 ] MAP
			[ SWAP MAXLONG mapper.replace 0 0 0 ] MAP
			[ SWAP $zero APPEND  [ 'hash_945fa9bc3027d7025e3' ] reducer.sum ] REDUCE 
			NONEMPTY { 'hash_945fa9bc3027d7025e3' '' } RELABEL 
		`,
	})

	if args[0] != swap {
		return fetch(node.Left, []string{args[0], kwargs["from"], kwargs["until"]}, kwargs)
	}

	return node.Left, nil
}
