package prototests

import (
	"testing"

	"github.com/ovh/erlenmeyer/proto/graphite"
)

//
// @code auto-generated
// Test can be started with
// go test proto/prototests/graphite_randomWalk_test.go proto/prototests/exec_test.go -v
//

// TestGraphiteRandomWalk process Invert Graphite Unit tests
func TestGraphiteRandomWalk(t *testing.T) {
	RunTest(t, graphiteGraphiteRandomWalk, "")
}

var graphiteGraphiteRandomWalk = []unitTests{

	{
		Plan: []graphite.Function{
			{
				Name:       "randomWalk",
				Arguments:  []string{"test"},
				Parameters: map[string]string{"until": "1548168564507231", "from": "1548168564507231 2 m -"},
			},
		},
		Samples: []OperatorGTSTest{
			{
				SamplePrefix: "CLEAR",
				GTSResult: []FloatGeoTimeSeries{
					{
						Class:  "test",
						Labels: map[string]string{},
						Attrs:  map[string]string{},
						Values: [][]float64{{1548168384507231.000000, 0.000000}, {1548168444507231.000000, 0.443996}, {1548168504507231.000000, 0.569946}, {1548168564507231.000000, 0.128644}},
					},
				},

				SeriesTests: seriesEqualitySkipValuesTestMap,
			},
		},
	},
	{
		Plan: []graphite.Function{
			{
				Name:       "randomWalk",
				Arguments:  []string{"test", "42"},
				Parameters: map[string]string{"until": "1548168564507231", "from": "1548168564507231 2 m -"},
			},
		},
		Samples: []OperatorGTSTest{
			{
				SamplePrefix: "CLEAR",
				GTSResult: []FloatGeoTimeSeries{
					{
						Class:  "test",
						Labels: map[string]string{},
						Attrs:  map[string]string{},
						Values: [][]float64{{1548168402507231.000000, 0.000000}, {1548168480507231.000000, 0.328745}, {1548168522507231.000000, 0.155467}, {1548168564507231.000000, 0.072316}},
					},
				},

				SeriesTests: seriesEqualitySkipValuesTestMap,
			},
		},
	},
}
