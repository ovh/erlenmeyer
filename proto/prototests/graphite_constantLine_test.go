package prototests

import (
	"testing"

	"github.com/ovh/erlenmeyer/proto/graphite"
)

//
// @code auto-generated
// Test can be started with
// go test proto/prototests/graphite_constantLine_test.go proto/prototests/exec_test.go -v
//

// TestGraphiteConstantLine process Invert Graphite Unit tests
func TestGraphiteConstantLine(t *testing.T) {
	RunTest(t, graphiteGraphiteConstantLine, "")
}

var graphiteGraphiteConstantLine = []unitTests{

	{
		Plan: []graphite.Function{
			{
				Name:       "constantLine",
				Arguments:  []string{"42"},
				Parameters: map[string]string{"until": "1548168564507231", "from": "1548168564507231 2 m -"},
			},
		},
		Samples: []OperatorGTSTest{
			{
				SamplePrefix: "CLEAR",
				GTSResult: []FloatGeoTimeSeries{
					{
						Class:  "42",
						Labels: map[string]string{},
						Attrs:  map[string]string{},
						Values: [][]float64{{1548168444507231.000000, 42.000000}, {1548168504507231.000000, 42.000000}, {1548168564507231.000000, 42.000000}},
					},
				},

				SeriesTests: seriesEqualityTestMap,
			},
		},
	},
}
