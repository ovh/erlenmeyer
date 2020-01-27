package prototests

import (
	"testing"

	"github.com/ovh/erlenmeyer/proto/graphite"
)

//
// @code auto-generated
// Test can be started with
// go test proto/prototests/graphite_threshold_test.go proto/prototests/exec_test.go -v
//

// TestGraphiteThreshold process Invert Graphite Unit tests
func TestGraphiteThreshold(t *testing.T) {
	RunTest(t, graphiteGraphiteThreshold, "")
}

var graphiteGraphiteThreshold = []unitTests{

	{
		Plan: []graphite.Function{
			{
				Name:       "threshold",
				Arguments:  []string{"42", "omgwtfbbq"},
				Parameters: map[string]string{"until": "1548168564507231", "from": "1548168564507231 2 m -"},
			},
		},
		Samples: []OperatorGTSTest{
			{
				SamplePrefix: "CLEAR",
				GTSResult: []FloatGeoTimeSeries{
					{
						Class:  "omgwtfbbq",
						Labels: map[string]string{},
						Attrs:  map[string]string{},
						Values: [][]float64{{1548168444507231.000000, 42.000000}, {1548168504507231.000000, 42.000000}, {1548168564507231.000000, 42.000000}},
					},
				},

				SeriesTests: seriesEqualityTestMap,
			},
		},
	},
	{
		Plan: []graphite.Function{
			{
				Name:       "threshold",
				Arguments:  []string{"42", "omgwtfbbq", "red"},
				Parameters: map[string]string{"until": "1548168564507231", "from": "1548168564507231 2 m -"},
			},
		},
		Samples: []OperatorGTSTest{
			{
				SamplePrefix: "CLEAR",
				GTSResult: []FloatGeoTimeSeries{
					{
						Class:  "omgwtfbbq",
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
