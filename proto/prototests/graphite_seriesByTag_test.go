package prototests

import (
	"testing"

	"github.com/ovh/erlenmeyer/proto/graphite"
)

//
// @code auto-generated
// Test can be started with
// go test proto/prototests/graphite_seriesByTag_test.go proto/prototests/exec_test.go -v
//

// TestGraphiteSeriesByTag process Invert Graphite Unit tests
func TestGraphiteSeriesByTag(t *testing.T) {
	RunTest(t, graphiteGraphiteSeriesByTag, "")
}

var graphiteGraphiteSeriesByTag = []unitTests{

	{
		Plan: []graphite.Function{
			{
				Name:       "seriesByTag",
				Arguments:  []string{"tag1=value1"},
				Parameters: map[string]string{"end": "202000000", "until": "202000000", "from": "202000000 2 m -", "token": "test", "span": "60 s", "count": "0"},
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
%> 'FETCH' DEF
	`,
				GTSResult: []FloatGeoTimeSeries{
					{
						Class:  "sample",
						Labels: map[string]string{"tag1": "value1"},
						Attrs:  map[string]string{},
						Values: [][]float64{{202000000.000000, 32.666667}, {142000000.000000, 10.000000}, {82000000.000000, -0.333333}, {22000000.000000, 1.000000}},
					},
				},

				SeriesTests: seriesEqualityTestMap,
			},
		},
	},
}
