package prototests

import (
	"testing"

	"github.com/ovh/erlenmeyer/proto/graphite"
)

//
// @code auto-generated
// Test can be started with
// go test proto/prototests/graphite_sin_test.go proto/prototests/exec_test.go -v
//

// TestGraphiteSin process Invert Graphite Unit tests
func TestGraphiteSin(t *testing.T) {
	RunTest(t, graphiteGraphiteSin, "")
}

var graphiteGraphiteSin = []unitTests{

	{
		Plan: []graphite.Function{
			{
				Name:       "sin",
				Arguments:  []string{"The.time.series", "2"},
				Parameters: map[string]string{"until": "1548168564507231", "from": "1548168564507231 2 m -"},
			},
		},
		Samples: []OperatorGTSTest{
			{
				SamplePrefix: "CLEAR",
				GTSResult: []FloatGeoTimeSeries{
					{
						Class:  "The.time.series",
						Labels: map[string]string{},
						Attrs:  map[string]string{},
						Values: [][]float64{{1548168444507231.000000, 0.350942}, {1548168504507231.000000, 0.706260}, {1548168564507231.000000, -1.564430}, {1548168624507231.000000, 1.981726}},
					},
				},

				SeriesTests: seriesEqualityTestMap,
			},
		},
	},
}
