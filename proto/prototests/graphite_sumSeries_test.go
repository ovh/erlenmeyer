package prototests

import (
	"testing"

	"github.com/ovh/erlenmeyer/proto/graphite"
)

//
// @code auto-generated
// Test can be started with
// go test proto/prototests/graphite_sumSeries_test.go proto/prototests/exec_test.go -v
//

// TestGraphiteSumSeries process Invert Graphite Unit tests
func TestGraphiteSumSeries(t *testing.T) {
	RunTest(t, graphiteGraphiteSumSeries, "")
}

var graphiteGraphiteSumSeries = []unitTests{

	{
		Plan: []graphite.Function{
			{
				Name:       "sumSeries",
				Arguments:  []string{"SWAP"},
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
						Values: [][]float64{{25000000.000000, 2.000000}, {85000000.000000, -0.666667}, {145000000.000000, 20.000000}, {205000000.000000, 65.333333}},
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
						Values: [][]float64{{25000000.000000, 1.000000}, {85000000.000000, -0.333333}, {145000000.000000, 10.000000}, {205000000.000000, 32.666667}},
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
						Values: [][]float64{{25000000.000000, 2.000000}, {85000000.000000, -0.666667}, {145000000.000000, 20.000000}, {205000000.000000, 65.333333}},
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
						Values: [][]float64{{25000000.000000, 3.000000}, {85000000.000000, 3.166667}, {145000000.000000, 16.000000}, {205000000.000000, 42.666667}},
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
						Values: [][]float64{{25000000.000000, 2.000000}, {85000000.000000, 1.666667}, {145000000.000000, 15.500000}, {205000000.000000, 41.666667}},
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
						Values: [][]float64{{25000000.000000, 2.000000}, {85000000.000000, 1.166667}, {145000000.000000, 10.500000}, {205000000.000000, 33.666667}},
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
						Values: [][]float64{{0.000000, 2.000000}, {35000000.000000, -2.000000}, {60000000.000000, 4.000000}, {72000000.000000, -4.000000}, {88000000.000000, 8.000000}, {112000000.000000, 16.000000}, {122000000.000000, 24.000000}, {132000000.000000, 32.000000}, {162000000.000000, 36.000000}, {182000000.000000, -40.000000}, {202000000.000000, 200.000000}},
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
						Values: [][]float64{{0.000000, 1.000000}, {35000000.000000, -1.000000}, {60000000.000000, 2.000000}, {72000000.000000, -2.000000}, {88000000.000000, 4.000000}, {112000000.000000, 8.000000}, {122000000.000000, 12.000000}, {132000000.000000, 16.000000}, {162000000.000000, 18.000000}, {182000000.000000, -20.000000}, {202000000.000000, 100.000000}},
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
						Values: [][]float64{{0.000000, 2.000000}, {35000000.000000, -2.000000}, {60000000.000000, 4.000000}, {72000000.000000, -4.000000}, {88000000.000000, 8.000000}, {112000000.000000, 16.000000}, {122000000.000000, 24.000000}, {132000000.000000, 32.000000}, {162000000.000000, 36.000000}, {182000000.000000, -40.000000}, {202000000.000000, 200.000000}},
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
						Values: [][]float64{{0.000000, 3.000000}, {35000000.000000, 0.000000}, {40000000.000000, 2.000000}, {60000000.000000, 4.000000}, {72000000.000000, 1.000000}, {82000000.000000, 1.000000}, {88000000.000000, 8.000000}, {110000000.000000, 0.000000}, {112000000.000000, 13.000000}, {122000000.000000, 18.000000}, {129000000.000000, 1.000000}, {132000000.000000, 23.000000}, {159000000.000000, 2.000000}, {162000000.000000, 26.000000}, {182000000.000000, -11.000000}, {192000000.000000, 0.000000}, {202000000.000000, 110.000000}, {205000000.000000, 1.000000}},
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
						Values: [][]float64{{0.000000, 2.000000}, {35000000.000000, 0.000000}, {60000000.000000, 4.000000}, {72000000.000000, 1.000000}, {88000000.000000, 8.000000}, {112000000.000000, 13.000000}, {122000000.000000, 18.000000}, {132000000.000000, 23.000000}, {162000000.000000, 26.000000}, {182000000.000000, -11.000000}, {202000000.000000, 110.000000}},
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
						Values: [][]float64{{0.000000, 2.000000}, {35000000.000000, -1.000000}, {40000000.000000, 2.000000}, {60000000.000000, 2.000000}, {72000000.000000, -2.000000}, {82000000.000000, 1.000000}, {88000000.000000, 4.000000}, {110000000.000000, 0.000000}, {112000000.000000, 8.000000}, {122000000.000000, 12.000000}, {129000000.000000, 1.000000}, {132000000.000000, 16.000000}, {159000000.000000, 2.000000}, {162000000.000000, 18.000000}, {182000000.000000, -20.000000}, {192000000.000000, 0.000000}, {202000000.000000, 100.000000}, {205000000.000000, 1.000000}},
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
