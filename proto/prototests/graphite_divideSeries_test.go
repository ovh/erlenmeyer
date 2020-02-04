package prototests

import (
	"testing"

	"github.com/ovh/erlenmeyer/proto/graphite"
)

//
// @code auto-generated
// Test can be started with
// go test proto/prototests/graphite_divideSeries_test.go proto/prototests/exec_test.go -v
//

// TestGraphiteDivideSeries process Invert Graphite Unit tests
func TestGraphiteDivideSeries(t *testing.T) {
	RunTest(t, graphiteGraphiteDivideSeries, "")
}

var graphiteGraphiteDivideSeries = []unitTests{

	{
		Plan: []graphite.Function{
			{
				Name:       "divideSeries",
				Arguments:  []string{"sample", "sample"},
				Parameters: map[string]string{"until": "182000000", "from": "182000000 2 m -", "node": "true", "span": "60 s", "count": "0", "end": "182000000"},
			},
		},
		Samples: []OperatorGTSTest{
			{
				SamplePrefix: `
	'' 'token' STORE
<% 
2 GET
'tags' STORE
[
    NEWGTS 'sample'  RENAME
    $tags RELABEL
    0.000000 NaN NaN NaN 1.000000 ADDVALUE
    35000000.000000 NaN NaN NaN -1.000000 ADDVALUE
    60000000.000000 NaN NaN NaN 2.000000 ADDVALUE
    72000000.000000 NaN NaN NaN -2.000000 ADDVALUE
    88000000.000000 NaN NaN NaN 4.000000 ADDVALUE
    112000000.000000 NaN NaN NaN 8.000000 ADDVALUE
    122000000.000000 NaN NaN NaN 12.000000 ADDVALUE
    132000000.000000 NaN NaN NaN 16.000000 ADDVALUE
    162000000.000000 NaN NaN NaN 18.000000 ADDVALUE
    182000000.000000 NaN NaN NaN -20.000000 ADDVALUE
    202000000.000000 NaN NaN NaN 100.000000 ADDVALUE
    
]
%> 'FETCHDOUBLE' DEF
	`,
				SampleSuffix: `
		<%
			DUP TOSELECTOR SWAP true MUSIGMA DROP TOSTRING +
		%>
		SORTBY
		`,
				GTSResult: []FloatGeoTimeSeries{
					{
						Class:  "sample",
						Labels: map[string]string{},
						Attrs:  map[string]string{},
						Values: [][]float64{{2000000.000000, 1.000000}, {62000000.000000, 1.000000}, {122000000.000000, 1.000000}, {182000000.000000, 1.000000}},
					},
				},

				SeriesTests: seriesEqualityTestMap,
			},
		},
	},
}
