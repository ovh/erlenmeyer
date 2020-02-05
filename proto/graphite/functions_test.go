package graphite_test

import (
	"strings"
	"testing"

	"github.com/ovh/erlenmeyer/core"
	"github.com/ovh/erlenmeyer/proto/graphite"
)

const (
	swap = "SWAP"
)

type functionTest struct {
	Function       graphite.Function
	ShouldContains []string
}

// removed functions
// "cumulative": Fail should contains test as updates native bucketize
// "consolidateBy": Fail should contains test as updates native bucketize

func TestFunction(t *testing.T) {
	for _, test := range functionsTest {
		fn, err := graphite.GetFunction(test.Function.Name)
		if err != nil {
			t.Error(err)

			continue
		}

		root := core.NewEmptyNode()
		_, err = fn(root, test.Function.Arguments, test.Function.Parameters)
		if err != nil {
			t.Errorf("%s: %s", test.Function.Name, err.Error())

			continue
		}

		ws := root.ToWarpScript("token", "", "")
		for _, shouldContain := range test.ShouldContains {
			if !strings.Contains(ws, shouldContain) {
				t.Errorf("%s does not contain this warpscript: %s but\n %s ", test.Function.Name, shouldContain, ws)
			}
		}
	}
}

var functionsTest = []functionTest{
	{
		Function: graphite.Function{
			Name:       "noOp",
			Arguments:  make([]string, 0),
			Parameters: make(map[string]string),
		},
		ShouldContains: make([]string, 0),
	},
	{
		Function: graphite.Function{
			Name:       "absolute",
			Arguments:  []string{swap},
			Parameters: make(map[string]string),
		},
		ShouldContains: []string{
			"[ SWAP  mapper.abs 0 0 0 ] MAP",
		},
	},
	{
		Function: graphite.Function{
			Name:       "aggregate",
			Arguments:  []string{swap, "sum"},
			Parameters: make(map[string]string),
		},
		ShouldContains: []string{
			"MARK",
			"COUNTTOMARK ->LIST SWAP DROP 'equivalenceClass' CSTORE",
			"[ SWAP $equivalenceClass reducer.sum ] REDUCE",
		},
	},
	{
		Function: graphite.Function{
			Name:       "alias",
			Arguments:  []string{swap, "nina"},
			Parameters: make(map[string]string),
		},
		ShouldContains: []string{
			"'nina' RENAME",
		},
	},
	{
		Function: graphite.Function{
			Name:       "aliasSub",
			Arguments:  []string{swap, ".*", "b"},
			Parameters: make(map[string]string),
		},
		ShouldContains: []string{
			"<% DROP DUP NAME '.*' 'b' REPLACE RENAME %> LMAP",
		},
	},
	{
		Function: graphite.Function{
			Name:       "aggregate",
			Arguments:  []string{swap, "average"},
			Parameters: make(map[string]string),
		},
		ShouldContains: []string{
			"MARK",
			"COUNTTOMARK ->LIST SWAP DROP 'equivalenceClass' CSTORE",
			"[ SWAP $equivalenceClass reducer.mean ] REDUCE",
		},
	},
	{
		Function: graphite.Function{
			Name:       "aggregate",
			Arguments:  []string{swap, "avg"},
			Parameters: make(map[string]string),
		},
		ShouldContains: []string{
			"MARK",
			"COUNTTOMARK ->LIST SWAP DROP 'equivalenceClass' CSTORE",
			"[ SWAP $equivalenceClass reducer.mean ] REDUCE",
		},
	},
	{
		Function: graphite.Function{
			Name:       "aggregate",
			Arguments:  []string{swap, "min"},
			Parameters: make(map[string]string),
		},
		ShouldContains: []string{
			"MARK",
			"COUNTTOMARK ->LIST SWAP DROP 'equivalenceClass' CSTORE",
			"[ SWAP $equivalenceClass reducer.min ] REDUCE",
		},
	},
	{
		Function: graphite.Function{
			Name:       "averageSeries",
			Arguments:  []string{swap, "25"},
			Parameters: make(map[string]string),
		},
		ShouldContains: []string{
			"MARK",
			"COUNTTOMARK ->LIST SWAP DROP 'equivalenceClass' CSTORE",
			"[ SWAP $equivalenceClass reducer.mean ] REDUCE",
		},
	},
	{
		Function: graphite.Function{
			Name:       "avg",
			Arguments:  []string{swap, "25"},
			Parameters: make(map[string]string),
		},
		ShouldContains: []string{
			"MARK",
			"COUNTTOMARK ->LIST SWAP DROP 'equivalenceClass' CSTORE",
			"[ SWAP $equivalenceClass reducer.mean ] REDUCE",
		},
	},
	{
		Function: graphite.Function{
			Name:       "countSeries",
			Arguments:  []string{swap},
			Parameters: make(map[string]string),
		},
		ShouldContains: []string{
			"SIZE",
		},
	},
	{
		Function: graphite.Function{
			Name:       "currentAbove",
			Arguments:  []string{swap, "100000"},
			Parameters: make(map[string]string),
		},
		ShouldContains: []string{
			"[ SWAP [] 100000 filter.last.gt ] FILTER",
		},
	},
	{
		Function: graphite.Function{
			Name:       "currentBelow",
			Arguments:  []string{swap, "100000"},
			Parameters: make(map[string]string),
		},
		ShouldContains: []string{
			"[ SWAP [] 100000 filter.last.lt ] FILTER",
		},
	},
	{
		Function: graphite.Function{
			Name:       "delay",
			Arguments:  []string{swap, "100000"},
			Parameters: make(map[string]string),
		},
		ShouldContains: []string{
			"100000 TIMESHIFT",
		},
	},
	{
		Function: graphite.Function{
			Name:       "derivative",
			Arguments:  []string{swap},
			Parameters: make(map[string]string),
		},
		ShouldContains: []string{
			"[ SWAP  mapper.delta 0 1 0 ] MAP",
		},
	},
	{
		Function: graphite.Function{
			Name:       "exclude",
			Arguments:  []string{swap, "pattern"},
			Parameters: make(map[string]string),
		},
		ShouldContains: []string{
			"[ SWAP [] '~(?!pattern).*' filter.byclass ] FILTER",
		},
	},
	{
		Function: graphite.Function{
			Name:       "grep",
			Arguments:  []string{swap, "pattern"},
			Parameters: make(map[string]string),
		},
		ShouldContains: []string{
			"[ SWAP [] '~pattern' filter.byclass ] FILTER",
		},
	},
	{
		Function: graphite.Function{
			Name:       "integral",
			Arguments:  []string{swap},
			Parameters: make(map[string]string),
		},
		ShouldContains: []string{
			"0 INTEGRATE",
		},
	},
	{
		Function: graphite.Function{
			Name:       "interpolate",
			Arguments:  []string{swap},
			Parameters: make(map[string]string),
		},
		ShouldContains: []string{
			"INTERPOLATE",
		},
	},
	{
		Function: graphite.Function{
			Name:       "invert",
			Arguments:  []string{swap},
			Parameters: make(map[string]string),
		},
		ShouldContains: []string{
			"<% { 'hash_945fa9bc3027d7025e3' ROT TOSTRING } RELABEL  %> LMAP",
			"[ SWAP 0 TODOUBLE mapper.eq 0 0 0 ] MAP 'zero' STORE",
			"[ SWAP 0 TODOUBLE mapper.ne 0 0 0 ] MAP",
			"'toDiv' STORE",
			"[ $toDiv 1 TODOUBLE mapper.replace 0 0 0 ] MAP 'dividende' RENAME",
			"op.div",
			"[ SWAP $zero APPEND  [ 'hash_945fa9bc3027d7025e3' ] reducer.sum ] REDUCE",
			"NONEMPTY { 'hash_945fa9bc3027d7025e3' '' } RELABEL",
		},
	},
	{
		Function: graphite.Function{
			Name:       "substr",
			Arguments:  []string{swap, "0"},
			Parameters: make(map[string]string),
		},
		ShouldContains: []string{},
	},
	{
		Function: graphite.Function{
			Name:       "hitcount",
			Arguments:  []string{swap, "60s"},
			Parameters: make(map[string]string),
		},
		ShouldContains: []string{
			"[ SWAP bucketizer.max 0 60 s 0 ] BUCKETIZE",
			"[ SWAP  mapper.delta 0 1 0 ] MAP",
		},
	},
	{
		Function: graphite.Function{
			Name:       "hitcount",
			Arguments:  []string{swap, "60w"},
			Parameters: make(map[string]string),
		},
		ShouldContains: []string{
			"[ SWAP bucketizer.max 0 60 w 0 ] BUCKETIZE",
			"[ SWAP  mapper.delta 0 1 0 ] MAP",
		},
	},
	{
		Function: graphite.Function{
			Name:       "aliasByTags",
			Arguments:  []string{swap, "host"},
			Parameters: make(map[string]string),
		},
		ShouldContains: []string{
			"$item '^[0-9]+$' MATCH SIZE 0 >",
			"$series LABELS $item CONTAINSKEY SWAP DROP",
			"$series LABELS $item CONTAINSKEY SWAP DROP",
			"$newName $series LABELS $item GET +",
		},
	},
	{
		Function: graphite.Function{
			Name:       "timeShift",
			Arguments:  []string{swap, "1d"},
			Parameters: make(map[string]string),
		},
		ShouldContains: []string{
			"1 d -1 * TIMESHIFT",
		},
	},
	{
		Function: graphite.Function{
			Name:       "seriesByTag",
			Arguments:  []string{"name=os.cpu"},
			Parameters: make(map[string]string),
		},
		ShouldContains: []string{
			"[ $token '~os\\.cpu' {}   ISO8601  ISO8601 ] FETCH",
		},
	},
	{
		Function: graphite.Function{
			Name:       "seriesByTag",
			Arguments:  []string{"name=os.cpu", "host=dn1"},
			Parameters: make(map[string]string),
		},
		ShouldContains: []string{
			"[ $token '~os\\.cpu' { 'host'  'dn1' }   ISO8601  ISO8601 ] FETCH",
		},
	},
	{
		Function: graphite.Function{
			Name:       "seriesByTag",
			Arguments:  []string{"name!=os.cpu"},
			Parameters: make(map[string]string),
		},
		ShouldContains: []string{
			"[ $token '~(?!os\\.cpu)' {}   ISO8601  ISO8601 ] FETCH",
		},
	},
	{
		Function: graphite.Function{
			Name:       "seriesByTag",
			Arguments:  []string{"name!=~os.cpu"},
			Parameters: make(map[string]string),
		},
		ShouldContains: []string{
			"[ $token '~(?!os\\.cpu)' {}   ISO8601  ISO8601 ] FETCH",
		},
	},
	{
		Function: graphite.Function{
			Name:       "seriesByTag",
			Arguments:  []string{"name=~os.cpu"},
			Parameters: make(map[string]string),
		},
		ShouldContains: []string{
			"[ $token '~os\\.cpu' {}   ISO8601  ISO8601 ] FETCH",
		},
	},
}
