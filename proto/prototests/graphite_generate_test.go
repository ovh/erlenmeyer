package prototests

import (
	"testing"

	"github.com/ovh/erlenmeyer/proto/graphite"
)

//
// Generates tests can be started with
// go test proto/prototests/graphite_generate_test.go proto/prototests/exec_test.go -v
// Will execute the test on a warp 10 instancte started at WARP_TEST_ENDPOINT or "http://127.0.0.1:8090/api/v0/exec"
//
// Generate fail for following functions:
// "countSeries"
// "aliasByNode"
// "seriesByTag"
//
// Empty generation for cases:
// "aggregate": single empty series in List
// "averageSeries": single empty series in List
// "avg": single empty series in List
// "currentAbove": single empty series in List
// "currentBelow": single empty series in List
// "grep": empty series for ALL test cases
// "invert": Half the test generate empty return

// testGenerateGraphite is used to generate Graphite Unit tests
func testGenerateGraphite(t *testing.T) {
	RunTest(t, generateGraphiteTests, "generateGraphite.txt")
}

var generateGraphiteTests = []unitTests{
	{
		Plan: []graphite.Function{{
			Name:       "absolute",
			Arguments:  []string{testSwap},
			Parameters: make(map[string]string),
		},
		},
		generate: true,
	},
	{
		Plan: []graphite.Function{{
			Name:       "aggregate",
			Arguments:  []string{testSwap, "sum"},
			Parameters: make(map[string]string),
		},
		},
		generate: true,
	},
	{
		Plan: []graphite.Function{{
			Name:       "alias",
			Arguments:  []string{testSwap, "nina"},
			Parameters: make(map[string]string),
		},
		},
		generate: true,
	},
	{
		Plan: []graphite.Function{{
			Name:       "aliasSub",
			Arguments:  []string{testSwap, ".*", "b"},
			Parameters: make(map[string]string),
		},
		},
		generate: true,
	},
	{
		Plan: []graphite.Function{{
			Name:       "aggregate",
			Arguments:  []string{testSwap, "average"},
			Parameters: make(map[string]string),
		},
		},
		generate: true,
	},
	{
		Plan: []graphite.Function{{
			Name:       "aggregate",
			Arguments:  []string{testSwap, "avg"},
			Parameters: make(map[string]string),
		},
		},
		generate: true,
	},
	{
		Plan: []graphite.Function{{
			Name:       "aggregate",
			Arguments:  []string{testSwap, "min"},
			Parameters: make(map[string]string),
		},
		},
		generate: true,
	},
	{
		Plan: []graphite.Function{{
			Name:       "averageSeries",
			Arguments:  []string{testSwap, "25"},
			Parameters: make(map[string]string),
		},
		},
		generate: true,
	},
	{
		Plan: []graphite.Function{{
			Name:       "avg",
			Arguments:  []string{testSwap, "25"},
			Parameters: make(map[string]string),
		},
		},
		generate: true,
	},
	{
		Plan: []graphite.Function{{
			Name:       "consolidateBy",
			Arguments:  []string{testSwap, "sum"},
			Parameters: make(map[string]string),
		},
		},
		generate: true,
	},
	{
		Plan: []graphite.Function{{
			Name:       "cumulative",
			Arguments:  []string{testSwap},
			Parameters: make(map[string]string),
		},
		},
		generate: true,
	},
	{
		Plan: []graphite.Function{{
			Name:       "currentAbove",
			Arguments:  []string{testSwap, "10"},
			Parameters: make(map[string]string),
		},
		},
		generate: true,
	},
	{
		Plan: []graphite.Function{{
			Name:       "currentBelow",
			Arguments:  []string{testSwap, "100000"},
			Parameters: make(map[string]string),
		},
		},
		generate: true,
	},
	{
		Plan: []graphite.Function{{
			Name:       "delay",
			Arguments:  []string{testSwap, "100000"},
			Parameters: make(map[string]string),
		},
		},
		generate: true,
	},
	{
		Plan: []graphite.Function{{
			Name:       "derivative",
			Arguments:  []string{testSwap},
			Parameters: make(map[string]string),
		},
		},
		generate: true,
	},
	{
		Plan: []graphite.Function{{
			Name:       "exclude",
			Arguments:  []string{testSwap, "pattern"},
			Parameters: make(map[string]string),
		},
		},
		generate: true,
	},
	{
		Plan: []graphite.Function{{
			Name:       "grep",
			Arguments:  []string{testSwap, "pattern"},
			Parameters: make(map[string]string),
		},
		},
		generate: true,
	},
	{
		Plan: []graphite.Function{{
			Name:       "integral",
			Arguments:  []string{testSwap},
			Parameters: make(map[string]string),
		},
		},
		generate: true,
	},
	{
		Plan: []graphite.Function{{
			Name:       "interpolate",
			Arguments:  []string{testSwap},
			Parameters: make(map[string]string),
		},
		},
		generate: true,
	},
	{
		Plan: []graphite.Function{{
			Name:       "invert",
			Arguments:  []string{testSwap},
			Parameters: make(map[string]string),
		},
		},
		generate: true,
	},
	{
		Plan: []graphite.Function{{
			Name:       "substr",
			Arguments:  []string{testSwap, "0"},
			Parameters: make(map[string]string),
		},
		},
		generate: true,
	},
	{
		Plan: []graphite.Function{{
			Name:       "hitcount",
			Arguments:  []string{testSwap, "60s"},
			Parameters: make(map[string]string),
		},
		},
		generate: true,
	},
	{
		Plan: []graphite.Function{{
			Name:       "hitcount",
			Arguments:  []string{testSwap, "60w"},
			Parameters: make(map[string]string),
		},
		},
		generate: true,
	},
	{
		Plan: []graphite.Function{{
			Name:       "timeShift",
			Arguments:  []string{testSwap, "1d"},
			Parameters: make(map[string]string),
		},
		},
		generate: true,
	},
}
