package prototests

import (
	"testing"

	"github.com/ovh/erlenmeyer/proto/graphite"
)

//
// @code auto-generated
// Test can be started with
// go test proto/prototests/graphite_averageSeriesWithWildcards_test.go proto/prototests/exec_test.go -v
//

// TestGraphiteAverageSeriesWithWildcards process Invert Graphite Unit tests
func TestGraphiteAverageSeriesWithWildcards(t *testing.T) {
	RunTest(t, graphiteGraphiteAverageSeriesWithWildcards, "")
}

var graphiteGraphiteAverageSeriesWithWildcards = []unitTests{

	{
		Plan: []graphite.Function{
			{
				Name:       "averageSeriesWithWildcards",
				Arguments:  []string{"SWAP", "2"},
				Parameters: make(map[string]string),
			},
		},
		Samples: []OperatorGTSTest{
			{
				SamplePrefix: BucketizeGtsList(BuildGtsList([]FloatGeoTimeSeries{emptyGTS, gts1, gts1})),
				GTSResult: []FloatGeoTimeSeries{
					{
						Class:  "",
						Labels: map[string]string{"label": "42"},
						Attrs:  map[string]string{},
						Values: [][]float64{},
					},
				},

				SeriesTests: seriesEqualityTestMap,
			},
			{
				SamplePrefix: BucketizeGtsList(BuildGtsList([]FloatGeoTimeSeries{emptyGTS, gts1})),
				GTSResult: []FloatGeoTimeSeries{
					{
						Class:  "",
						Labels: map[string]string{"label": "42"},
						Attrs:  map[string]string{},
						Values: [][]float64{},
					},
				},

				SeriesTests: seriesEqualityTestMap,
			},
			{
				SamplePrefix: BucketizeGtsList(BuildGtsList([]FloatGeoTimeSeries{gts1, gts1})),
				GTSResult: []FloatGeoTimeSeries{
					{
						Class:  "sample",
						Labels: map[string]string{"label": "42"},
						Attrs:  map[string]string{},
						Values: [][]float64{{25000000.000000, 1.000000}, {85000000.000000, -0.333333}, {145000000.000000, 10.000000}, {205000000.000000, 32.666667}},
					},
				},

				SeriesTests: seriesEqualityTestMap,
			},
			{
				SamplePrefix: BucketizeGtsList(BuildGtsList([]FloatGeoTimeSeries{gts1, gts2, gts3})),
				GTSResult: []FloatGeoTimeSeries{
					{
						Class:  "sample",
						Labels: map[string]string{},
						Attrs:  map[string]string{},
						Values: [][]float64{{25000000.000000, 1.000000}, {85000000.000000, 1.055556}, {145000000.000000, 5.333333}, {205000000.000000, 14.222222}},
					},
				},

				SeriesTests: seriesEqualityTestMap,
			},
			{
				SamplePrefix: BucketizeGtsList(BuildGtsList([]FloatGeoTimeSeries{gts1, gts2})),
				GTSResult: []FloatGeoTimeSeries{
					{
						Class:  "sample",
						Labels: map[string]string{},
						Attrs:  map[string]string{},
						Values: [][]float64{{25000000.000000, 1.000000}, {85000000.000000, 0.833333}, {145000000.000000, 7.750000}, {205000000.000000, 20.833333}},
					},
				},

				SeriesTests: seriesEqualityTestMap,
			},
			{
				SamplePrefix: BucketizeGtsList(BuildGtsList([]FloatGeoTimeSeries{gts1, gts3})),
				GTSResult: []FloatGeoTimeSeries{
					{
						Class:  "sample",
						Labels: map[string]string{},
						Attrs:  map[string]string{},
						Values: [][]float64{{25000000.000000, 1.000000}, {85000000.000000, 0.583333}, {145000000.000000, 5.250000}, {205000000.000000, 16.833333}},
					},
				},

				SeriesTests: seriesEqualityTestMap,
			},
			{
				SamplePrefix: BucketizeGtsList(BuildGtsList([]FloatGeoTimeSeries{gts1})),
				GTSResult: []FloatGeoTimeSeries{
					{
						Class:  "sample",
						Labels: map[string]string{"label": "42"},
						Attrs:  map[string]string{},
						Values: [][]float64{{25000000.000000, 1.000000}, {85000000.000000, -0.333333}, {145000000.000000, 10.000000}, {205000000.000000, 32.666667}},
					},
				},

				SeriesTests: seriesEqualityTestMap,
			},
			{
				SamplePrefix: BuildGtsList([]FloatGeoTimeSeries{emptyGTS, gts1, gts1}),
				GTSResult: []FloatGeoTimeSeries{
					{
						Class:  "",
						Labels: map[string]string{"label": "42"},
						Attrs:  map[string]string{},
						Values: [][]float64{},
					},
				},

				SeriesTests: seriesEqualityTestMap,
			},
			{
				SamplePrefix: BuildGtsList([]FloatGeoTimeSeries{emptyGTS, gts1}),
				GTSResult: []FloatGeoTimeSeries{
					{
						Class:  "",
						Labels: map[string]string{"label": "42"},
						Attrs:  map[string]string{},
						Values: [][]float64{},
					},
				},

				SeriesTests: seriesEqualityTestMap,
			},
			{
				SamplePrefix: BuildGtsList([]FloatGeoTimeSeries{gts1, gts1}),
				GTSResult: []FloatGeoTimeSeries{
					{
						Class:  "sample",
						Labels: map[string]string{"label": "42"},
						Attrs:  map[string]string{},
						Values: [][]float64{{0.000000, 1.000000}, {35000000.000000, -1.000000}, {60000000.000000, 2.000000}, {72000000.000000, -2.000000}, {88000000.000000, 4.000000}, {112000000.000000, 8.000000}, {122000000.000000, 12.000000}, {132000000.000000, 16.000000}, {162000000.000000, 18.000000}, {182000000.000000, -20.000000}, {202000000.000000, 100.000000}},
					},
				},

				SeriesTests: seriesEqualityTestMap,
			},
			{
				SamplePrefix: BuildGtsList([]FloatGeoTimeSeries{gts1, gts2, gts3}),
				GTSResult: []FloatGeoTimeSeries{
					{
						Class:  "sample",
						Labels: map[string]string{},
						Attrs:  map[string]string{},
						Values: [][]float64{{0.000000, 1.000000}},
					},
				},

				SeriesTests: seriesEqualityTestMap,
			},
			{
				SamplePrefix: BuildGtsList([]FloatGeoTimeSeries{gts1, gts2}),
				GTSResult: []FloatGeoTimeSeries{
					{
						Class:  "sample",
						Labels: map[string]string{},
						Attrs:  map[string]string{},
						Values: [][]float64{{0.000000, 1.000000}, {35000000.000000, 0.000000}, {60000000.000000, 2.000000}, {72000000.000000, 0.500000}, {88000000.000000, 4.000000}, {112000000.000000, 6.500000}, {122000000.000000, 9.000000}, {132000000.000000, 11.500000}, {162000000.000000, 13.000000}, {182000000.000000, -5.500000}, {202000000.000000, 55.000000}},
					},
				},

				SeriesTests: seriesEqualityTestMap,
			},
			{
				SamplePrefix: BuildGtsList([]FloatGeoTimeSeries{gts1, gts3}),
				GTSResult: []FloatGeoTimeSeries{
					{
						Class:  "sample",
						Labels: map[string]string{},
						Attrs:  map[string]string{},
						Values: [][]float64{{0.000000, 1.000000}},
					},
				},

				SeriesTests: seriesEqualityTestMap,
			},
			{
				SamplePrefix: BuildGtsList([]FloatGeoTimeSeries{gts1}),
				GTSResult: []FloatGeoTimeSeries{
					{
						Class:  "sample",
						Labels: map[string]string{"label": "42"},
						Attrs:  map[string]string{},
						Values: [][]float64{{0.000000, 1.000000}, {35000000.000000, -1.000000}, {60000000.000000, 2.000000}, {72000000.000000, -2.000000}, {88000000.000000, 4.000000}, {112000000.000000, 8.000000}, {122000000.000000, 12.000000}, {132000000.000000, 16.000000}, {162000000.000000, 18.000000}, {182000000.000000, -20.000000}, {202000000.000000, 100.000000}},
					},
				},

				SeriesTests: seriesEqualityTestMap,
			},
		},
	},
}
