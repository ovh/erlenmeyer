package prototests

import (
	"testing"

	"github.com/ovh/erlenmeyer/proto/graphite"
)

//
// @code auto-generated
// Test can be started with
// go test proto/prototests/graphite_countSeries_test.go proto/prototests/exec_test.go -v
//

// TestGraphiteCountSeries process Invert Graphite Unit tests
func TestGraphiteCountSeries(t *testing.T) {
	RunTest(t, graphiteGraphiteCountSeries, "")
}

var graphiteGraphiteCountSeries = []unitTests{

	{
		Plan: []graphite.Function{
			{
				Name:       "countSeries",
				Arguments:  []string{"SWAP", "test"},
				Parameters: map[string]string{"until": "1548168564507231", "from": "1548168564507231 2 m -"},
			},
		},
		Samples: []OperatorGTSTest{
			{
				SamplePrefix: BucketizeGtsList(BuildGtsList([]FloatGeoTimeSeries{emptyGTS, gts1, gts1})),
				GTSResult: []FloatGeoTimeSeries{
					{
						Class:  "$countSeries",
						Labels: map[string]string{},
						Attrs:  map[string]string{},
						Values: [][]float64{{1548168444507231.000000, 3.000000}, {1548168504507231.000000, 3.000000}, {1548168564507231.000000, 3.000000}},
					},
				},

				SeriesTests: seriesEqualityTestMap,
			},
			{
				SamplePrefix: BucketizeGtsList(BuildGtsList([]FloatGeoTimeSeries{emptyGTS, gts1})),
				GTSResult: []FloatGeoTimeSeries{
					{
						Class:  "$countSeries",
						Labels: map[string]string{},
						Attrs:  map[string]string{},
						Values: [][]float64{{1548168444507231.000000, 2.000000}, {1548168504507231.000000, 2.000000}, {1548168564507231.000000, 2.000000}},
					},
				},

				SeriesTests: seriesEqualityTestMap,
			},
			{
				SamplePrefix: BucketizeGtsList(BuildGtsList([]FloatGeoTimeSeries{emptyGTS})),
				GTSResult: []FloatGeoTimeSeries{
					{
						Class:  "$countSeries",
						Labels: map[string]string{},
						Attrs:  map[string]string{},
						Values: [][]float64{{1548168444507231.000000, 1.000000}, {1548168504507231.000000, 1.000000}, {1548168564507231.000000, 1.000000}},
					},
				},

				SeriesTests: seriesEqualityTestMap,
			},
			{
				SamplePrefix: BucketizeGtsList(BuildGtsList([]FloatGeoTimeSeries{gts1, gts1})),
				GTSResult: []FloatGeoTimeSeries{
					{
						Class:  "$countSeries",
						Labels: map[string]string{},
						Attrs:  map[string]string{},
						Values: [][]float64{{1548168444507231.000000, 2.000000}, {1548168504507231.000000, 2.000000}, {1548168564507231.000000, 2.000000}},
					},
				},

				SeriesTests: seriesEqualityTestMap,
			},
			{
				SamplePrefix: BucketizeGtsList(BuildGtsList([]FloatGeoTimeSeries{gts1, gts2, gts3})),
				GTSResult: []FloatGeoTimeSeries{
					{
						Class:  "$countSeries",
						Labels: map[string]string{},
						Attrs:  map[string]string{},
						Values: [][]float64{{1548168444507231.000000, 3.000000}, {1548168504507231.000000, 3.000000}, {1548168564507231.000000, 3.000000}},
					},
				},

				SeriesTests: seriesEqualityTestMap,
			},
			{
				SamplePrefix: BucketizeGtsList(BuildGtsList([]FloatGeoTimeSeries{gts1, gts2})),
				GTSResult: []FloatGeoTimeSeries{
					{
						Class:  "$countSeries",
						Labels: map[string]string{},
						Attrs:  map[string]string{},
						Values: [][]float64{{1548168444507231.000000, 2.000000}, {1548168504507231.000000, 2.000000}, {1548168564507231.000000, 2.000000}},
					},
				},

				SeriesTests: seriesEqualityTestMap,
			},
			{
				SamplePrefix: BucketizeGtsList(BuildGtsList([]FloatGeoTimeSeries{gts1, gts3})),
				GTSResult: []FloatGeoTimeSeries{
					{
						Class:  "$countSeries",
						Labels: map[string]string{},
						Attrs:  map[string]string{},
						Values: [][]float64{{1548168444507231.000000, 2.000000}, {1548168504507231.000000, 2.000000}, {1548168564507231.000000, 2.000000}},
					},
				},

				SeriesTests: seriesEqualityTestMap,
			},
			{
				SamplePrefix: BucketizeGtsList(BuildGtsList([]FloatGeoTimeSeries{gts1})),
				GTSResult: []FloatGeoTimeSeries{
					{
						Class:  "$countSeries",
						Labels: map[string]string{},
						Attrs:  map[string]string{},
						Values: [][]float64{{1548168444507231.000000, 1.000000}, {1548168504507231.000000, 1.000000}, {1548168564507231.000000, 1.000000}},
					},
				},

				SeriesTests: seriesEqualityTestMap,
			},
			{
				SamplePrefix: BuildGtsList([]FloatGeoTimeSeries{emptyGTS, gts1, gts1}),
				GTSResult: []FloatGeoTimeSeries{
					{
						Class:  "$countSeries",
						Labels: map[string]string{},
						Attrs:  map[string]string{},
						Values: [][]float64{{1548168444507231.000000, 3.000000}, {1548168504507231.000000, 3.000000}, {1548168564507231.000000, 3.000000}},
					},
				},

				SeriesTests: seriesEqualityTestMap,
			},
			{
				SamplePrefix: BuildGtsList([]FloatGeoTimeSeries{emptyGTS, gts1}),
				GTSResult: []FloatGeoTimeSeries{
					{
						Class:  "$countSeries",
						Labels: map[string]string{},
						Attrs:  map[string]string{},
						Values: [][]float64{{1548168444507231.000000, 2.000000}, {1548168504507231.000000, 2.000000}, {1548168564507231.000000, 2.000000}},
					},
				},

				SeriesTests: seriesEqualityTestMap,
			},
			{
				SamplePrefix: BuildGtsList([]FloatGeoTimeSeries{emptyGTS}),
				GTSResult: []FloatGeoTimeSeries{
					{
						Class:  "$countSeries",
						Labels: map[string]string{},
						Attrs:  map[string]string{},
						Values: [][]float64{{1548168444507231.000000, 1.000000}, {1548168504507231.000000, 1.000000}, {1548168564507231.000000, 1.000000}},
					},
				},

				SeriesTests: seriesEqualityTestMap,
			},
			{
				SamplePrefix: BuildGtsList([]FloatGeoTimeSeries{gts1, gts1}),
				GTSResult: []FloatGeoTimeSeries{
					{
						Class:  "$countSeries",
						Labels: map[string]string{},
						Attrs:  map[string]string{},
						Values: [][]float64{{1548168444507231.000000, 2.000000}, {1548168504507231.000000, 2.000000}, {1548168564507231.000000, 2.000000}},
					},
				},

				SeriesTests: seriesEqualityTestMap,
			},
			{
				SamplePrefix: BuildGtsList([]FloatGeoTimeSeries{gts1, gts2, gts3}),
				GTSResult: []FloatGeoTimeSeries{
					{
						Class:  "$countSeries",
						Labels: map[string]string{},
						Attrs:  map[string]string{},
						Values: [][]float64{{1548168444507231.000000, 3.000000}, {1548168504507231.000000, 3.000000}, {1548168564507231.000000, 3.000000}},
					},
				},

				SeriesTests: seriesEqualityTestMap,
			},
			{
				SamplePrefix: BuildGtsList([]FloatGeoTimeSeries{gts1, gts2}),
				GTSResult: []FloatGeoTimeSeries{
					{
						Class:  "$countSeries",
						Labels: map[string]string{},
						Attrs:  map[string]string{},
						Values: [][]float64{{1548168444507231.000000, 2.000000}, {1548168504507231.000000, 2.000000}, {1548168564507231.000000, 2.000000}},
					},
				},

				SeriesTests: seriesEqualityTestMap,
			},
			{
				SamplePrefix: BuildGtsList([]FloatGeoTimeSeries{gts1, gts3}),
				GTSResult: []FloatGeoTimeSeries{
					{
						Class:  "$countSeries",
						Labels: map[string]string{},
						Attrs:  map[string]string{},
						Values: [][]float64{{1548168444507231.000000, 2.000000}, {1548168504507231.000000, 2.000000}, {1548168564507231.000000, 2.000000}},
					},
				},

				SeriesTests: seriesEqualityTestMap,
			},
			{
				SamplePrefix: BuildGtsList([]FloatGeoTimeSeries{gts1}),
				GTSResult: []FloatGeoTimeSeries{
					{
						Class:  "$countSeries",
						Labels: map[string]string{},
						Attrs:  map[string]string{},
						Values: [][]float64{{1548168444507231.000000, 1.000000}, {1548168504507231.000000, 1.000000}, {1548168564507231.000000, 1.000000}},
					},
				},

				SeriesTests: seriesEqualityTestMap,
			},
		},
	},
}
