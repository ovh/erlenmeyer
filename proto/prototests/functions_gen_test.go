package prototests

import (
	"testing"

	"github.com/ovh/erlenmeyer/proto/graphite"
)

// Exec:
// go test proto/prototests/functions_gen_test.go proto/prototests/exec_test.go -v
//
// Will execute the test on a warp 10 instancte started at WARP_TEST_ENDPOINT or "http://127.0.0.1:8090/api/v0/exec"

//
// Generate fail for following functions:
// "aggregate" with median argument: keeping only bucketized series due to
// https://github.com/senx/warp10-platform/issues/393
//	"aggregateLine": miss "multiply" param
//
// Cannot generate yet the tests for functions

// testGenerateFunctions is used to generate "Functions" Graphite Unit tests based on the []unitTests given as parameter RunTest
// The generated GO test code is written in function_test.go file
// To re-generate a file replace testGenerateFunctions by TestGenerateFunctions

func testGenerateFunctions(t *testing.T) {
	RunTest(t, generateFailTests, generateFile)
}

var generateFailTests = []unitTests{}

var generateFunctionsTests = []unitTests{
	{
		Plan: []graphite.Function{{
			Name:       "alias",
			Arguments:  []string{testSwap, "test"},
			Parameters: make(map[string]string),
		},
		},
		generate: true,
	},
	{
		Plan: []graphite.Function{{
			Name:       "aliasByMetric",
			Arguments:  []string{testSwap},
			Parameters: make(map[string]string),
		},
		},
		generate: true,
	},
	{
		Plan: []graphite.Function{{
			Name:       "aliasByNode",
			Arguments:  []string{testSwap, "1"},
			Parameters: make(map[string]string),
		},
		},
		generate: true,
	},
	{
		Plan: []graphite.Function{{
			Name:       "aliasByTags",
			Arguments:  []string{testSwap, "label"},
			Parameters: make(map[string]string),
		},
		},
		generate: true,
	},
	{
		Plan: []graphite.Function{{
			Name:       "aliasSub",
			Arguments:  []string{testSwap, "empty", "sample"},
			Parameters: make(map[string]string),
		},
		},
		generate: true,
	},
	{
		Plan: []graphite.Function{{
			Name:       "drawAsInfinite",
			Arguments:  []string{testSwap},
			Parameters: make(map[string]string),
		},
		},
		generate: true,
		generateSuffix: `
		<%
			DUP TOSELECTOR SWAP true MUSIGMA DROP TOSTRING + 
		%>
		SORTBY
		`,
	},
	{
		Plan: []graphite.Function{{
			Name:       "invert",
			Arguments:  []string{testSwap},
			Parameters: make(map[string]string),
		},
		},
		generate: true,
		generateSuffix: `
		<%
			DUP TOSELECTOR SWAP true MUSIGMA DROP TOSTRING + 
		%>
		SORTBY
		`,
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
			Arguments:  []string{testSwap, "median"},
			Parameters: make(map[string]string),
		},
		},
		generate: true,
		generateSample: map[string]string{
			"BucketizeGtsList(BuildGtsList([]FloatGeoTimeSeries{gts1}))":     BucketizeGtsList(BuildGtsList([]FloatGeoTimeSeries{gts1})),
			"BucketizeGtsList(BuildGtsList([]FloatGeoTimeSeries{emptyGTS}))": BucketizeGtsList(BuildGtsList([]FloatGeoTimeSeries{emptyGTS})),
			"BucketizeGtsList(BuildGtsList([]FloatGeoTimeSeries{emptyGTS, gts1}))": BucketizeGtsList(BuildGtsList([]FloatGeoTimeSeries{
				emptyGTS, gts1})),
			"BucketizeGtsList(BuildGtsList([]FloatGeoTimeSeries{emptyGTS, gts1, gts1}))": BucketizeGtsList(BuildGtsList([]FloatGeoTimeSeries{
				emptyGTS, gts1, gts1})),
			"BucketizeGtsList(BuildGtsList([]FloatGeoTimeSeries{gts1, gts1}))": BucketizeGtsList(BuildGtsList([]FloatGeoTimeSeries{
				gts1, gts1})),
			"BucketizeGtsList(BuildGtsList([]FloatGeoTimeSeries{gts1, gts2}))": BucketizeGtsList(BuildGtsList([]FloatGeoTimeSeries{
				gts1, gts2})),
			"BucketizeGtsList(BuildGtsList([]FloatGeoTimeSeries{gts1, gts3}))": BucketizeGtsList(BuildGtsList([]FloatGeoTimeSeries{
				gts1, gts3})),
			"BucketizeGtsList(BuildGtsList([]FloatGeoTimeSeries{gts1, gts2, gts3}))": BucketizeGtsList(BuildGtsList([]FloatGeoTimeSeries{
				gts1, gts2, gts3})),
		},
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
			Name:       "aggregate",
			Arguments:  []string{testSwap, "min"},
			Parameters: make(map[string]string),
		},
		},
		generate: true,
	},
	{
		Plan: []graphite.Function{{
			Name:       "aggregate",
			Arguments:  []string{testSwap, "max"},
			Parameters: make(map[string]string),
		},
		},
		generate: true,
	},
	{
		Plan: []graphite.Function{{
			Name:       "aggregate",
			Arguments:  []string{testSwap, "diff"},
			Parameters: make(map[string]string),
		},
		},
		generate: true,
	},
	{
		Plan: []graphite.Function{{
			Name:       "aggregate",
			Arguments:  []string{testSwap, "stddev"},
			Parameters: make(map[string]string),
		},
		},
		generate: true,
	},
	{
		Plan: []graphite.Function{{
			Name:       "aggregate",
			Arguments:  []string{testSwap, "count"},
			Parameters: make(map[string]string),
		},
		},
		generate: true,
	},
	{
		Plan: []graphite.Function{{
			Name:       "aggregate",
			Arguments:  []string{testSwap, "range"},
			Parameters: make(map[string]string),
		},
		},
		generate: true,
	},
	{
		Plan: []graphite.Function{{
			Name:       "aggregate",
			Arguments:  []string{testSwap, "multiply"},
			Parameters: make(map[string]string),
		},
		},
		generate: true,
	},
	{
		Plan: []graphite.Function{{
			Name:       "aggregate",
			Arguments:  []string{testSwap, "last"},
			Parameters: make(map[string]string),
		},
		},
		generate: true,
	},
	{
		Plan: []graphite.Function{{
			Name:       "aggregateWithWildcards",
			Arguments:  []string{testSwap, "average", "2"},
			Parameters: make(map[string]string),
		},
		},
		generate: true,
	},
	{
		Plan: []graphite.Function{{
			Name:       "aggregateWithWildcards",
			Arguments:  []string{testSwap, "median", "2"},
			Parameters: make(map[string]string),
		},
		},
		generate: true,
		generateSample: map[string]string{
			"BucketizeGtsList(BuildGtsList([]FloatGeoTimeSeries{gts1}))":     BucketizeGtsList(BuildGtsList([]FloatGeoTimeSeries{gts1})),
			"BucketizeGtsList(BuildGtsList([]FloatGeoTimeSeries{emptyGTS}))": BucketizeGtsList(BuildGtsList([]FloatGeoTimeSeries{emptyGTS})),
			"BucketizeGtsList(BuildGtsList([]FloatGeoTimeSeries{emptyGTS, gts1}))": BucketizeGtsList(BuildGtsList([]FloatGeoTimeSeries{
				emptyGTS, gts1})),
			"BucketizeGtsList(BuildGtsList([]FloatGeoTimeSeries{emptyGTS, gts1, gts1}))": BucketizeGtsList(BuildGtsList([]FloatGeoTimeSeries{
				emptyGTS, gts1, gts1})),
			"BucketizeGtsList(BuildGtsList([]FloatGeoTimeSeries{gts1, gts1}))": BucketizeGtsList(BuildGtsList([]FloatGeoTimeSeries{
				gts1, gts1})),
			"BucketizeGtsList(BuildGtsList([]FloatGeoTimeSeries{gts1, gts2}))": BucketizeGtsList(BuildGtsList([]FloatGeoTimeSeries{
				gts1, gts2})),
			"BucketizeGtsList(BuildGtsList([]FloatGeoTimeSeries{gts1, gts3}))": BucketizeGtsList(BuildGtsList([]FloatGeoTimeSeries{
				gts1, gts3})),
			"BucketizeGtsList(BuildGtsList([]FloatGeoTimeSeries{gts1, gts2, gts3}))": BucketizeGtsList(BuildGtsList([]FloatGeoTimeSeries{
				gts1, gts2, gts3})),
		},
	},
	{
		Plan: []graphite.Function{{
			Name:       "aggregateWithWildcards",
			Arguments:  []string{testSwap, "sum", "2"},
			Parameters: make(map[string]string),
		},
		},
		generate: true,
	},
	{
		Plan: []graphite.Function{{
			Name:       "aggregateWithWildcards",
			Arguments:  []string{testSwap, "min", "2"},
			Parameters: make(map[string]string),
		},
		},
		generate: true,
	},
	{
		Plan: []graphite.Function{{
			Name:       "aggregateWithWildcards",
			Arguments:  []string{testSwap, "max", "2"},
			Parameters: make(map[string]string),
		},
		},
		generate: true,
	},
	{
		Plan: []graphite.Function{{
			Name:       "aggregateWithWildcards",
			Arguments:  []string{testSwap, "diff", "2"},
			Parameters: make(map[string]string),
		},
		},
		generate: true,
	},
	{
		Plan: []graphite.Function{{
			Name:       "aggregateWithWildcards",
			Arguments:  []string{testSwap, "stddev", "2"},
			Parameters: make(map[string]string),
		},
		},
		generate: true,
	},
	{
		Plan: []graphite.Function{{
			Name:       "aggregateWithWildcards",
			Arguments:  []string{testSwap, "count", "2"},
			Parameters: make(map[string]string),
		},
		},
		generate: true,
	},
	{
		Plan: []graphite.Function{{
			Name:       "aggregateWithWildcards",
			Arguments:  []string{testSwap, "range", "2"},
			Parameters: make(map[string]string),
		},
		},
		generate: true,
	},
	{
		Plan: []graphite.Function{{
			Name:       "aggregateWithWildcards",
			Arguments:  []string{testSwap, "multiply", "2"},
			Parameters: make(map[string]string),
		},
		},
		generate: true,
	},
	{
		Plan: []graphite.Function{{
			Name:       "aggregateWithWildcards",
			Arguments:  []string{testSwap, "last", "2"},
			Parameters: make(map[string]string),
		},
		},
		generate: true,
	},
	{
		Plan: []graphite.Function{{
			Name:       "aggregateLine",
			Arguments:  []string{testSwap, "average"},
			Parameters: make(map[string]string),
		},
		},
		generate: true,
		generateSample: map[string]string{
			"BucketizeGtsList(BuildGtsList([]FloatGeoTimeSeries{gts1}))":     BucketizeGtsList(BuildGtsList([]FloatGeoTimeSeries{gts1})),
			"BucketizeGtsList(BuildGtsList([]FloatGeoTimeSeries{emptyGTS}))": BucketizeGtsList(BuildGtsList([]FloatGeoTimeSeries{emptyGTS})),
			"BucketizeGtsList(BuildGtsList([]FloatGeoTimeSeries{emptyGTS, gts1}))": BucketizeGtsList(BuildGtsList([]FloatGeoTimeSeries{
				emptyGTS, gts1})),
			"BucketizeGtsList(BuildGtsList([]FloatGeoTimeSeries{emptyGTS, gts1, gts1}))": BucketizeGtsList(BuildGtsList([]FloatGeoTimeSeries{
				emptyGTS, gts1, gts1})),
			"BucketizeGtsList(BuildGtsList([]FloatGeoTimeSeries{gts1, gts1}))": BucketizeGtsList(BuildGtsList([]FloatGeoTimeSeries{
				gts1, gts1})),
			"BucketizeGtsList(BuildGtsList([]FloatGeoTimeSeries{gts1, gts2}))": BucketizeGtsList(BuildGtsList([]FloatGeoTimeSeries{
				gts1, gts2})),
			"BucketizeGtsList(BuildGtsList([]FloatGeoTimeSeries{gts1, gts3}))": BucketizeGtsList(BuildGtsList([]FloatGeoTimeSeries{
				gts1, gts3})),
			"BucketizeGtsList(BuildGtsList([]FloatGeoTimeSeries{gts1, gts2, gts3}))": BucketizeGtsList(BuildGtsList([]FloatGeoTimeSeries{
				gts1, gts2, gts3})),
		},
	},
	{
		Plan: []graphite.Function{{
			Name:       "aggregateLine",
			Arguments:  []string{testSwap, "median"},
			Parameters: make(map[string]string),
		},
		},
		generate: true,
		generateSample: map[string]string{
			"BucketizeGtsList(BuildGtsList([]FloatGeoTimeSeries{gts1}))":     BucketizeGtsList(BuildGtsList([]FloatGeoTimeSeries{gts1})),
			"BucketizeGtsList(BuildGtsList([]FloatGeoTimeSeries{emptyGTS}))": BucketizeGtsList(BuildGtsList([]FloatGeoTimeSeries{emptyGTS})),
			"BucketizeGtsList(BuildGtsList([]FloatGeoTimeSeries{emptyGTS, gts1}))": BucketizeGtsList(BuildGtsList([]FloatGeoTimeSeries{
				emptyGTS, gts1})),
			"BucketizeGtsList(BuildGtsList([]FloatGeoTimeSeries{emptyGTS, gts1, gts1}))": BucketizeGtsList(BuildGtsList([]FloatGeoTimeSeries{
				emptyGTS, gts1, gts1})),
			"BucketizeGtsList(BuildGtsList([]FloatGeoTimeSeries{gts1, gts1}))": BucketizeGtsList(BuildGtsList([]FloatGeoTimeSeries{
				gts1, gts1})),
			"BucketizeGtsList(BuildGtsList([]FloatGeoTimeSeries{gts1, gts2}))": BucketizeGtsList(BuildGtsList([]FloatGeoTimeSeries{
				gts1, gts2})),
			"BucketizeGtsList(BuildGtsList([]FloatGeoTimeSeries{gts1, gts3}))": BucketizeGtsList(BuildGtsList([]FloatGeoTimeSeries{
				gts1, gts3})),
			"BucketizeGtsList(BuildGtsList([]FloatGeoTimeSeries{gts1, gts2, gts3}))": BucketizeGtsList(BuildGtsList([]FloatGeoTimeSeries{
				gts1, gts2, gts3})),
		},
		generateSuffix: `
		<%
			DUP TOSELECTOR SWAP true MUSIGMA DROP TOSTRING + 
		%>
		SORTBY
		`,
	},
	{
		Plan: []graphite.Function{{
			Name:       "aggregateLine",
			Arguments:  []string{testSwap, "sum"},
			Parameters: make(map[string]string),
		},
		},
		generate: true,
		generateSample: map[string]string{
			"BucketizeGtsList(BuildGtsList([]FloatGeoTimeSeries{gts1}))":     BucketizeGtsList(BuildGtsList([]FloatGeoTimeSeries{gts1})),
			"BucketizeGtsList(BuildGtsList([]FloatGeoTimeSeries{emptyGTS}))": BucketizeGtsList(BuildGtsList([]FloatGeoTimeSeries{emptyGTS})),
			"BucketizeGtsList(BuildGtsList([]FloatGeoTimeSeries{emptyGTS, gts1}))": BucketizeGtsList(BuildGtsList([]FloatGeoTimeSeries{
				emptyGTS, gts1})),
			"BucketizeGtsList(BuildGtsList([]FloatGeoTimeSeries{emptyGTS, gts1, gts1}))": BucketizeGtsList(BuildGtsList([]FloatGeoTimeSeries{
				emptyGTS, gts1, gts1})),
			"BucketizeGtsList(BuildGtsList([]FloatGeoTimeSeries{gts1, gts1}))": BucketizeGtsList(BuildGtsList([]FloatGeoTimeSeries{
				gts1, gts1})),
			"BucketizeGtsList(BuildGtsList([]FloatGeoTimeSeries{gts1, gts2}))": BucketizeGtsList(BuildGtsList([]FloatGeoTimeSeries{
				gts1, gts2})),
			"BucketizeGtsList(BuildGtsList([]FloatGeoTimeSeries{gts1, gts3}))": BucketizeGtsList(BuildGtsList([]FloatGeoTimeSeries{
				gts1, gts3})),
			"BucketizeGtsList(BuildGtsList([]FloatGeoTimeSeries{gts1, gts2, gts3}))": BucketizeGtsList(BuildGtsList([]FloatGeoTimeSeries{
				gts1, gts2, gts3})),
		},
		generateSuffix: `
		<%
			DUP TOSELECTOR SWAP true MUSIGMA DROP TOSTRING + 
		%>
		SORTBY
		`,
	},
	{
		Plan: []graphite.Function{{
			Name:       "aggregateLine",
			Arguments:  []string{testSwap, "min"},
			Parameters: make(map[string]string),
		},
		},
		generate: true,
		generateSample: map[string]string{
			"BucketizeGtsList(BuildGtsList([]FloatGeoTimeSeries{gts1}))":     BucketizeGtsList(BuildGtsList([]FloatGeoTimeSeries{gts1})),
			"BucketizeGtsList(BuildGtsList([]FloatGeoTimeSeries{emptyGTS}))": BucketizeGtsList(BuildGtsList([]FloatGeoTimeSeries{emptyGTS})),
			"BucketizeGtsList(BuildGtsList([]FloatGeoTimeSeries{emptyGTS, gts1}))": BucketizeGtsList(BuildGtsList([]FloatGeoTimeSeries{
				emptyGTS, gts1})),
			"BucketizeGtsList(BuildGtsList([]FloatGeoTimeSeries{emptyGTS, gts1, gts1}))": BucketizeGtsList(BuildGtsList([]FloatGeoTimeSeries{
				emptyGTS, gts1, gts1})),
			"BucketizeGtsList(BuildGtsList([]FloatGeoTimeSeries{gts1, gts1}))": BucketizeGtsList(BuildGtsList([]FloatGeoTimeSeries{
				gts1, gts1})),
			"BucketizeGtsList(BuildGtsList([]FloatGeoTimeSeries{gts1, gts2}))": BucketizeGtsList(BuildGtsList([]FloatGeoTimeSeries{
				gts1, gts2})),
			"BucketizeGtsList(BuildGtsList([]FloatGeoTimeSeries{gts1, gts3}))": BucketizeGtsList(BuildGtsList([]FloatGeoTimeSeries{
				gts1, gts3})),
			"BucketizeGtsList(BuildGtsList([]FloatGeoTimeSeries{gts1, gts2, gts3}))": BucketizeGtsList(BuildGtsList([]FloatGeoTimeSeries{
				gts1, gts2, gts3})),
		},
		generateSuffix: `
		<%
			DUP TOSELECTOR SWAP true MUSIGMA DROP TOSTRING + 
		%>
		SORTBY
		`,
	},
	{
		Plan: []graphite.Function{{
			Name:       "aggregateLine",
			Arguments:  []string{testSwap, "max"},
			Parameters: make(map[string]string),
		},
		},
		generate: true,
		generateSample: map[string]string{
			"BucketizeGtsList(BuildGtsList([]FloatGeoTimeSeries{gts1}))":     BucketizeGtsList(BuildGtsList([]FloatGeoTimeSeries{gts1})),
			"BucketizeGtsList(BuildGtsList([]FloatGeoTimeSeries{emptyGTS}))": BucketizeGtsList(BuildGtsList([]FloatGeoTimeSeries{emptyGTS})),
			"BucketizeGtsList(BuildGtsList([]FloatGeoTimeSeries{emptyGTS, gts1}))": BucketizeGtsList(BuildGtsList([]FloatGeoTimeSeries{
				emptyGTS, gts1})),
			"BucketizeGtsList(BuildGtsList([]FloatGeoTimeSeries{emptyGTS, gts1, gts1}))": BucketizeGtsList(BuildGtsList([]FloatGeoTimeSeries{
				emptyGTS, gts1, gts1})),
			"BucketizeGtsList(BuildGtsList([]FloatGeoTimeSeries{gts1, gts1}))": BucketizeGtsList(BuildGtsList([]FloatGeoTimeSeries{
				gts1, gts1})),
			"BucketizeGtsList(BuildGtsList([]FloatGeoTimeSeries{gts1, gts2}))": BucketizeGtsList(BuildGtsList([]FloatGeoTimeSeries{
				gts1, gts2})),
			"BucketizeGtsList(BuildGtsList([]FloatGeoTimeSeries{gts1, gts3}))": BucketizeGtsList(BuildGtsList([]FloatGeoTimeSeries{
				gts1, gts3})),
			"BucketizeGtsList(BuildGtsList([]FloatGeoTimeSeries{gts1, gts2, gts3}))": BucketizeGtsList(BuildGtsList([]FloatGeoTimeSeries{
				gts1, gts2, gts3})),
		},
		generateSuffix: `
		<%
			DUP TOSELECTOR SWAP true MUSIGMA DROP TOSTRING + 
		%>
		SORTBY
		`,
	},
	{
		Plan: []graphite.Function{{
			Name:       "aggregateLine",
			Arguments:  []string{testSwap, "stddev"},
			Parameters: make(map[string]string),
		},
		},
		generate: true,
		generateSample: map[string]string{
			"BucketizeGtsList(BuildGtsList([]FloatGeoTimeSeries{gts1}))":     BucketizeGtsList(BuildGtsList([]FloatGeoTimeSeries{gts1})),
			"BucketizeGtsList(BuildGtsList([]FloatGeoTimeSeries{emptyGTS}))": BucketizeGtsList(BuildGtsList([]FloatGeoTimeSeries{emptyGTS})),
			"BucketizeGtsList(BuildGtsList([]FloatGeoTimeSeries{emptyGTS, gts1}))": BucketizeGtsList(BuildGtsList([]FloatGeoTimeSeries{
				emptyGTS, gts1})),
			"BucketizeGtsList(BuildGtsList([]FloatGeoTimeSeries{emptyGTS, gts1, gts1}))": BucketizeGtsList(BuildGtsList([]FloatGeoTimeSeries{
				emptyGTS, gts1, gts1})),
			"BucketizeGtsList(BuildGtsList([]FloatGeoTimeSeries{gts1, gts1}))": BucketizeGtsList(BuildGtsList([]FloatGeoTimeSeries{
				gts1, gts1})),
			"BucketizeGtsList(BuildGtsList([]FloatGeoTimeSeries{gts1, gts2}))": BucketizeGtsList(BuildGtsList([]FloatGeoTimeSeries{
				gts1, gts2})),
			"BucketizeGtsList(BuildGtsList([]FloatGeoTimeSeries{gts1, gts3}))": BucketizeGtsList(BuildGtsList([]FloatGeoTimeSeries{
				gts1, gts3})),
			"BucketizeGtsList(BuildGtsList([]FloatGeoTimeSeries{gts1, gts2, gts3}))": BucketizeGtsList(BuildGtsList([]FloatGeoTimeSeries{
				gts1, gts2, gts3})),
		},
		generateSuffix: `
		<%
			DUP TOSELECTOR SWAP true MUSIGMA DROP TOSTRING + 
		%>
		SORTBY
		`,
	},
	{
		Plan: []graphite.Function{{
			Name:       "aggregateLine",
			Arguments:  []string{testSwap, "last"},
			Parameters: make(map[string]string),
		},
		},
		generate: true,
		generateSample: map[string]string{
			"BucketizeGtsList(BuildGtsList([]FloatGeoTimeSeries{gts1}))":     BucketizeGtsList(BuildGtsList([]FloatGeoTimeSeries{gts1})),
			"BucketizeGtsList(BuildGtsList([]FloatGeoTimeSeries{emptyGTS}))": BucketizeGtsList(BuildGtsList([]FloatGeoTimeSeries{emptyGTS})),
			"BucketizeGtsList(BuildGtsList([]FloatGeoTimeSeries{emptyGTS, gts1}))": BucketizeGtsList(BuildGtsList([]FloatGeoTimeSeries{
				emptyGTS, gts1})),
			"BucketizeGtsList(BuildGtsList([]FloatGeoTimeSeries{emptyGTS, gts1, gts1}))": BucketizeGtsList(BuildGtsList([]FloatGeoTimeSeries{
				emptyGTS, gts1, gts1})),
			"BucketizeGtsList(BuildGtsList([]FloatGeoTimeSeries{gts1, gts1}))": BucketizeGtsList(BuildGtsList([]FloatGeoTimeSeries{
				gts1, gts1})),
			"BucketizeGtsList(BuildGtsList([]FloatGeoTimeSeries{gts1, gts2}))": BucketizeGtsList(BuildGtsList([]FloatGeoTimeSeries{
				gts1, gts2})),
			"BucketizeGtsList(BuildGtsList([]FloatGeoTimeSeries{gts1, gts3}))": BucketizeGtsList(BuildGtsList([]FloatGeoTimeSeries{
				gts1, gts3})),
			"BucketizeGtsList(BuildGtsList([]FloatGeoTimeSeries{gts1, gts2, gts3}))": BucketizeGtsList(BuildGtsList([]FloatGeoTimeSeries{
				gts1, gts2, gts3})),
		},
	},
	{
		Plan: []graphite.Function{{
			Name:       "averageAbove",
			Arguments:  []string{testSwap, "2"},
			Parameters: make(map[string]string),
		},
		},
		generate: true,
	},
	{
		Plan: []graphite.Function{{
			Name:       "averageBelow",
			Arguments:  []string{testSwap, "20"},
			Parameters: make(map[string]string),
		},
		},
		generate: true,
	},
	{
		Plan: []graphite.Function{{
			Name:       "averageSeries",
			Arguments:  []string{testSwap},
			Parameters: make(map[string]string),
		},
		},
		generate: true,
	},
	{
		Plan: []graphite.Function{{
			Name:       "averageSeriesWithWildcards",
			Arguments:  []string{testSwap, "2"},
			Parameters: make(map[string]string),
		},
		},
		generate: true,
	},
	{
		Plan: []graphite.Function{{
			Name:       "consolidateBy",
			Arguments:  []string{testSwap, "avg"},
			Parameters: make(map[string]string),
		},
		},
		generate: true,
	},
	{
		Plan: []graphite.Function{{
			Name:       "consolidateBy",
			Arguments:  []string{testSwap, "average"},
			Parameters: make(map[string]string),
		},
		},
		generate: true,
	},
	{
		Plan: []graphite.Function{{
			Name:       "consolidateBy",
			Arguments:  []string{testSwap, "median"},
			Parameters: make(map[string]string),
		},
		},
		generate: true,
	},
	{
		Plan: []graphite.Function{{
			Name:       "consolidateBy",
			Arguments:  []string{testSwap, "median"},
			Parameters: make(map[string]string),
		},
		},
		generate: true,
	}, {
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
			Name:       "consolidateBy",
			Arguments:  []string{testSwap, "min"},
			Parameters: make(map[string]string),
		},
		},
		generate: true,
	},
	{
		Plan: []graphite.Function{{
			Name:       "consolidateBy",
			Arguments:  []string{testSwap, "max"},
			Parameters: make(map[string]string),
		},
		},
		generate: true,
	},
	{
		Plan: []graphite.Function{{
			Name:       "consolidateBy",
			Arguments:  []string{testSwap, "count"},
			Parameters: make(map[string]string),
		},
		},
		generate: true,
	},
	{
		Plan: []graphite.Function{{
			Name:       "consolidateBy",
			Arguments:  []string{testSwap, "multiply"},
			Parameters: make(map[string]string),
		},
		},
		generate: true,
	},
	{
		Plan: []graphite.Function{{
			Name:       "consolidateBy",
			Arguments:  []string{testSwap, "last"},
			Parameters: make(map[string]string),
		},
		},
		generate: true,
	},
	{
		Plan: []graphite.Function{{
			Name:       "constantLine",
			Arguments:  []string{"42"},
			Parameters: map[string]string{"until": "1548168564507231", "from": "1548168564507231 2 m -"},
		},
		},
		generate:       true,
		generateSample: emptySample,
	},
	{
		Plan: []graphite.Function{{
			Name:       "countSeries",
			Arguments:  []string{testSwap, "test"},
			Parameters: map[string]string{"until": "1548168564507231", "from": "1548168564507231 2 m -"},
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
			Arguments:  []string{testSwap, "2"},
			Parameters: make(map[string]string),
		},
		},
		generate: true,
	},
	{
		Plan: []graphite.Function{{
			Name:       "currentBelow",
			Arguments:  []string{testSwap, "102"},
			Parameters: make(map[string]string),
		},
		},
		generate: true,
	},
	{
		Plan: []graphite.Function{{
			Name:       "delay",
			Arguments:  []string{testSwap, "2"},
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
			Name:       "diffSeries",
			Arguments:  []string{testSwap},
			Parameters: make(map[string]string),
		},
		},
		generate: true,
	},
	{
		Plan: []graphite.Function{{
			Name:       "exclude",
			Arguments:  []string{testSwap, "empty"},
			Parameters: make(map[string]string),
		},
		},
		generate: true,
	},
	{
		Plan: []graphite.Function{{
			Name:       "grep",
			Arguments:  []string{testSwap, "sample"},
			Parameters: make(map[string]string),
		},
		},
		generate: true,
	},
	{
		Plan: []graphite.Function{{
			Name:       "hitcount",
			Arguments:  []string{testSwap, "1h"},
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
			Name:       "keepLastValue",
			Arguments:  []string{testSwap},
			Parameters: make(map[string]string),
		},
		},
		generate: true,
	},
	{
		Plan: []graphite.Function{{
			Name:       "limit",
			Arguments:  []string{testSwap, "3"},
			Parameters: make(map[string]string),
		},
		},
		generate: true,
	},
	{
		Plan: []graphite.Function{{
			Name:       "logarithm",
			Arguments:  []string{testSwap, "2"},
			Parameters: make(map[string]string),
		},
		},
		generate: true,
		generateSuffix: `[ SWAP mapper.finite 0 0 0 ] MAP
		`,
	},
	{
		Plan: []graphite.Function{{
			Name:       "maxSeries",
			Arguments:  []string{testSwap},
			Parameters: make(map[string]string),
		},
		},
		generate: true,
	},
	{
		Plan: []graphite.Function{{
			Name:       "minMax",
			Arguments:  []string{testSwap},
			Parameters: make(map[string]string),
		},
		},
		generate: true,
	},
	{
		Plan: []graphite.Function{{
			Name:       "minSeries",
			Arguments:  []string{testSwap},
			Parameters: make(map[string]string),
		},
		},
		generate: true,
	},
	{
		Plan: []graphite.Function{{
			Name:       "offset",
			Arguments:  []string{testSwap, "3"},
			Parameters: make(map[string]string),
		},
		},
		generate: true,
	},
	{
		Plan: []graphite.Function{{
			Name:       "perSecond",
			Arguments:  []string{testSwap, "3"},
			Parameters: make(map[string]string),
		},
		},
		generate: true,
	},
	{
		Plan: []graphite.Function{{
			Name:       "perSecond",
			Arguments:  []string{testSwap},
			Parameters: make(map[string]string),
		},
		},
		generate: true,
	},
	{
		Plan: []graphite.Function{{
			Name:       "pow",
			Arguments:  []string{testSwap, "3"},
			Parameters: make(map[string]string),
		},
		},
		generate: true,
	},
	{
		Plan: []graphite.Function{{
			Name:       "rangeOfSeries",
			Arguments:  []string{testSwap},
			Parameters: make(map[string]string),
		},
		},
		generate: true,
	},
	{
		Plan: []graphite.Function{{
			Name:       "removeAboveValue",
			Arguments:  []string{testSwap, "3"},
			Parameters: make(map[string]string),
		},
		},
		generate: true,
	},
	{
		Plan: []graphite.Function{{
			Name:       "removeBelowValue",
			Arguments:  []string{testSwap, "3"},
			Parameters: make(map[string]string),
		},
		},
		generate: true,
	},
	{
		Plan: []graphite.Function{{
			Name:       "scale",
			Arguments:  []string{testSwap, "3"},
			Parameters: make(map[string]string),
		},
		},
		generate: true,
	},
	{
		Plan: []graphite.Function{{
			Name:       "scale",
			Arguments:  []string{testSwap, "0.01"},
			Parameters: make(map[string]string),
		},
		},
		generate: true,
	},
	{
		Plan: []graphite.Function{{
			Name:       "scaleToSeconds",
			Arguments:  []string{testSwap, "60"},
			Parameters: make(map[string]string),
		},
		},
		generate: true,
	},
	{
		Plan: []graphite.Function{{
			Name:       "sortByName",
			Arguments:  []string{testSwap},
			Parameters: make(map[string]string),
		},
		},
		generate: true,
	},
	{
		Plan: []graphite.Function{{
			Name:       "stddevSeries",
			Arguments:  []string{testSwap},
			Parameters: make(map[string]string),
		},
		},
		generate: true,
	},
	{
		Plan: []graphite.Function{{
			Name:       "substr",
			Arguments:  []string{testSwap, "0", "2"},
			Parameters: make(map[string]string),
		},
		},
		generate: true,
	},
	{
		Plan: []graphite.Function{{
			Name:       "substr",
			Arguments:  []string{testSwap, "5", "6"},
			Parameters: make(map[string]string),
		},
		},
		generate: true,
	},
	{
		Plan: []graphite.Function{{
			Name:       "substr",
			Arguments:  []string{testSwap, "4", "3"},
			Parameters: make(map[string]string),
		},
		},
		generate: true,
	},
	{
		Plan: []graphite.Function{{
			Name:       "sumSeries",
			Arguments:  []string{testSwap},
			Parameters: make(map[string]string),
		},
		},
		generate: true,
	},
	{
		Plan: []graphite.Function{{
			Name:       "timeShift",
			Arguments:  []string{testSwap, "1h"},
			Parameters: make(map[string]string),
		},
		},
		generate: true,
	},
	{
		Plan: []graphite.Function{{
			Name:       "transformNull",
			Arguments:  []string{testSwap},
			Parameters: make(map[string]string),
		},
		},
		generate: true,
	},
	{
		Plan: []graphite.Function{{
			Name:       "unique",
			Arguments:  []string{testSwap},
			Parameters: make(map[string]string),
		},
		},
		generate: true,
		generateSuffix: `
		NONEMPTY
		<%
			DUP TOSELECTOR SWAP true MUSIGMA DROP TOSTRING + 
		%>
		SORTBY
		`,
	},
	{
		Plan: []graphite.Function{{
			Name:       "highestAverage",
			Arguments:  []string{testSwap, "3"},
			Parameters: make(map[string]string),
		},
		},
		generate: true,
	},
	{
		Plan: []graphite.Function{{
			Name:       "highestCurrent",
			Arguments:  []string{testSwap, "3"},
			Parameters: make(map[string]string),
		},
		},
		generate: true,
	},
	{
		Plan: []graphite.Function{{
			Name:       "highestMax",
			Arguments:  []string{testSwap, "3"},
			Parameters: make(map[string]string),
		},
		},
		generate: true,
	},
	{
		Plan: []graphite.Function{{
			Name:       "lowestAverage",
			Arguments:  []string{testSwap, "3"},
			Parameters: make(map[string]string),
		},
		},
		generate: true,
	},
	{
		Plan: []graphite.Function{{
			Name:       "lowestCurrent",
			Arguments:  []string{testSwap, "3"},
			Parameters: make(map[string]string),
		},
		},
		generate: true,
	},
	{
		Plan: []graphite.Function{{
			Name:       "maximumAbove",
			Arguments:  []string{testSwap, "3"},
			Parameters: make(map[string]string),
		},
		},
		generate: true,
	},
	{
		Plan: []graphite.Function{{
			Name:       "maximumBelow",
			Arguments:  []string{testSwap, "102"},
			Parameters: make(map[string]string),
		},
		},
		generate: true,
	},
	{
		Plan: []graphite.Function{{
			Name:       "minimumAbove",
			Arguments:  []string{testSwap, "-42"},
			Parameters: make(map[string]string),
		},
		},
		generate: true,
	},
	{
		Plan: []graphite.Function{{
			Name:       "minimumBelow",
			Arguments:  []string{testSwap, "3"},
			Parameters: make(map[string]string),
		},
		},
		generate: true,
	},
	// Random Walk test use series equality checks skipping values equality (random values at each execution)
	{
		Plan: []graphite.Function{{
			Name:       "randomWalk",
			Arguments:  []string{"test"},
			Parameters: map[string]string{"until": "1548168564507231", "from": "1548168564507231 2 m -"},
		},
		},
		generate:       true,
		generateSample: emptySample,
		generateChecks: "seriesEqualitySkipValuesTestMap",
	},
	{
		Plan: []graphite.Function{{
			Name:       "randomWalk",
			Arguments:  []string{"test", "42"},
			Parameters: map[string]string{"until": "1548168564507231", "from": "1548168564507231 2 m -"},
		},
		},
		generate:       true,
		generateSample: emptySample,
		generateChecks: "seriesEqualitySkipValuesTestMap",
	},
	{
		Plan: []graphite.Function{{
			Name:       "seriesByTag",
			Arguments:  []string{"tag1=value1"},
			Parameters: map[string]string{"until": "202000000", "from": "202000000 2 m -", "token": "test"},
		},
		},
		generate:       true,
		generateSample: redefinedFetch,
	},
	{
		Plan: []graphite.Function{{
			Name:       "sin",
			Arguments:  []string{"The.time.series", "2"},
			Parameters: map[string]string{"until": "1548168564507231", "from": "1548168564507231 2 m -"},
		},
		},
		generate:       true,
		generateSample: emptySample,
	},
	{
		Plan: []graphite.Function{{
			Name:       "sortByMaxima",
			Arguments:  []string{testSwap},
			Parameters: make(map[string]string),
		},
		},
		generate: true,
	},
	{
		Plan: []graphite.Function{{
			Name:       "sortByMinima",
			Arguments:  []string{testSwap},
			Parameters: make(map[string]string),
		},
		},
		generate: true,
	},
	{
		Plan: []graphite.Function{{
			Name:       "sortByTotal",
			Arguments:  []string{testSwap},
			Parameters: make(map[string]string),
		},
		},
		generate: true,
	},
	{
		Plan: []graphite.Function{{
			Name:       "sumSeriesWithWildcards",
			Arguments:  []string{testSwap, "2"},
			Parameters: make(map[string]string),
		},
		},
		generate: true,
	},
	{
		Plan: []graphite.Function{{
			Name:       "summarize",
			Arguments:  []string{testSwap, "1h"},
			Parameters: make(map[string]string),
		},
		},
		generate: true,
	},
	{
		Plan: []graphite.Function{{
			Name:       "threshold",
			Arguments:  []string{"42", "omgwtfbbq"},
			Parameters: map[string]string{"until": "1548168564507231", "from": "1548168564507231 2 m -"},
		},
		},
		generate:       true,
		generateSample: emptySample,
	},
	{
		Plan: []graphite.Function{{
			Name:       "threshold",
			Arguments:  []string{"42", "omgwtfbbq", "red"},
			Parameters: map[string]string{"until": "1548168564507231", "from": "1548168564507231 2 m -"},
		},
		},
		generate:       true,
		generateSample: emptySample,
	},
	{
		Plan: []graphite.Function{{
			Name:       "time",
			Arguments:  []string{"The.time.series"},
			Parameters: map[string]string{"until": "1548168564507231", "from": "1548168564507231 2 m -"},
		},
		},
		generate:       true,
		generateSample: emptySample,
	},
	{
		Plan: []graphite.Function{{
			Name:       "timeSlice",
			Arguments:  []string{testSwap, "182000000"},
			Parameters: map[string]string{"until": "1548168564507231", "from": "1548168564507231 2 m -"},
		},
		},
		generate: true,
	},
	{
		Plan: []graphite.Function{{
			Name:       "multiplySeries",
			Arguments:  []string{testSwap},
			Parameters: make(map[string]string),
		},
		},
		generate: true,
	},
	{
		Plan: []graphite.Function{{
			Name:       "multiplySeriesWithWildcards",
			Arguments:  []string{testSwap, "2"},
			Parameters: make(map[string]string),
		},
		},
		generate: true,
	},
	{
		Plan: []graphite.Function{{
			Name:       "divideSeries",
			Arguments:  []string{"sample", "sample"},
			Parameters: map[string]string{"until": "182000000", "from": "182000000 2 m -"},
		},
		},
		generate:       true,
		generateSample: redefinedFetch,
		generateSuffix: `
		<%
			DUP TOSELECTOR SWAP true MUSIGMA DROP TOSTRING +
		%>
		SORTBY
		`,
	},
	{
		Plan: []graphite.Function{{
			Name:       "divideSeriesLists",
			Arguments:  []string{"sample", "sample"},
			Parameters: map[string]string{"until": "182000000", "from": "182000000 2 m -"},
		},
		},
		generate:       true,
		generateSample: redefinedFetch,
		generateSuffix: `
		<%
			DUP TOSELECTOR SWAP true MUSIGMA DROP TOSTRING +
		%>
		SORTBY
		`,
	},
}
