package prototests

import (
	"testing"

	"github.com/ovh/erlenmeyer/proto/graphite"
)

//
// @code auto-generated
// Test can be started with
// go test proto/prototests/graphite_minimumAbove_test.go proto/prototests/exec_test.go -v
//

// TestGraphiteMinimumAbove process Invert Graphite Unit tests
func TestGraphiteMinimumAbove(t *testing.T) {
	RunTest(t, graphiteGraphiteMinimumAbove, "")
}

var graphiteGraphiteMinimumAbove = []unitTests{

	{
		Plan: []graphite.Function{
			{
				Name:       "minimumAbove",
				Arguments:  []string{"SWAP", "-42"},
				Parameters: make(map[string]string),
			},
		},
		Samples: []OperatorGTSTest{
			{
				SamplePrefix: BucketizeGtsList(BuildGtsList([]FloatGeoTimeSeries{emptyGTS, gts1, gts1})),
				GTSResult: []FloatGeoTimeSeries{
					{
						Class:  "sample",
						Labels: map[string]string{"label": "42"},
						Attrs:  map[string]string{},
						Values: [][]float64{{25000000.000000, 1.000000}, {85000000.000000, -0.333333}, {145000000.000000, 10.000000}, {205000000.000000, 32.666667}},
					},
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
				SamplePrefix: BucketizeGtsList(BuildGtsList([]FloatGeoTimeSeries{emptyGTS, gts1})),
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
				SamplePrefix: BucketizeGtsList(BuildGtsList([]FloatGeoTimeSeries{gts1, gts1})),
				GTSResult: []FloatGeoTimeSeries{
					{
						Class:  "sample",
						Labels: map[string]string{"label": "42"},
						Attrs:  map[string]string{},
						Values: [][]float64{{25000000.000000, 1.000000}, {85000000.000000, -0.333333}, {145000000.000000, 10.000000}, {205000000.000000, 32.666667}},
					},
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
						Labels: map[string]string{"label": "42"},
						Attrs:  map[string]string{},
						Values: [][]float64{{25000000.000000, 1.000000}, {85000000.000000, -0.333333}, {145000000.000000, 10.000000}, {205000000.000000, 32.666667}},
					},
					{
						Class:  "sample",
						Labels: map[string]string{"other": "test", "label": "41"},
						Attrs:  map[string]string{},
						Values: [][]float64{{25000000.000000, 1.000000}, {85000000.000000, 2.000000}, {145000000.000000, 5.500000}, {205000000.000000, 9.000000}},
					},
					{
						Class:  "sample",
						Labels: map[string]string{"other": "test", "label": "41"},
						Attrs:  map[string]string{},
						Values: [][]float64{{25000000.000000, 1.000000}, {85000000.000000, 1.500000}, {145000000.000000, 0.500000}, {205000000.000000, 1.000000}},
					},
				},

				SeriesTests: seriesEqualityTestMap,
			},
			{
				SamplePrefix: BucketizeGtsList(BuildGtsList([]FloatGeoTimeSeries{gts1, gts2})),
				GTSResult: []FloatGeoTimeSeries{
					{
						Class:  "sample",
						Labels: map[string]string{"label": "42"},
						Attrs:  map[string]string{},
						Values: [][]float64{{25000000.000000, 1.000000}, {85000000.000000, -0.333333}, {145000000.000000, 10.000000}, {205000000.000000, 32.666667}},
					},
					{
						Class:  "sample",
						Labels: map[string]string{"other": "test", "label": "41"},
						Attrs:  map[string]string{},
						Values: [][]float64{{25000000.000000, 1.000000}, {85000000.000000, 2.000000}, {145000000.000000, 5.500000}, {205000000.000000, 9.000000}},
					},
				},

				SeriesTests: seriesEqualityTestMap,
			},
			{
				SamplePrefix: BucketizeGtsList(BuildGtsList([]FloatGeoTimeSeries{gts1, gts3})),
				GTSResult: []FloatGeoTimeSeries{
					{
						Class:  "sample",
						Labels: map[string]string{"label": "42"},
						Attrs:  map[string]string{},
						Values: [][]float64{{25000000.000000, 1.000000}, {85000000.000000, -0.333333}, {145000000.000000, 10.000000}, {205000000.000000, 32.666667}},
					},
					{
						Class:  "sample",
						Labels: map[string]string{"other": "test", "label": "41"},
						Attrs:  map[string]string{},
						Values: [][]float64{{25000000.000000, 1.000000}, {85000000.000000, 1.500000}, {145000000.000000, 0.500000}, {205000000.000000, 1.000000}},
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
						Class:  "sample",
						Labels: map[string]string{"label": "42"},
						Attrs:  map[string]string{},
						Values: [][]float64{{0.000000, 1.000000}, {35000000.000000, -1.000000}, {60000000.000000, 2.000000}, {72000000.000000, -2.000000}, {88000000.000000, 4.000000}, {112000000.000000, 8.000000}, {122000000.000000, 12.000000}, {132000000.000000, 16.000000}, {162000000.000000, 18.000000}, {182000000.000000, -20.000000}, {202000000.000000, 100.000000}},
					},
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
				SamplePrefix: BuildGtsList([]FloatGeoTimeSeries{emptyGTS, gts1}),
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
				SamplePrefix: BuildGtsList([]FloatGeoTimeSeries{gts1, gts1}),
				GTSResult: []FloatGeoTimeSeries{
					{
						Class:  "sample",
						Labels: map[string]string{"label": "42"},
						Attrs:  map[string]string{},
						Values: [][]float64{{0.000000, 1.000000}, {35000000.000000, -1.000000}, {60000000.000000, 2.000000}, {72000000.000000, -2.000000}, {88000000.000000, 4.000000}, {112000000.000000, 8.000000}, {122000000.000000, 12.000000}, {132000000.000000, 16.000000}, {162000000.000000, 18.000000}, {182000000.000000, -20.000000}, {202000000.000000, 100.000000}},
					},
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
						Labels: map[string]string{"label": "42"},
						Attrs:  map[string]string{},
						Values: [][]float64{{0.000000, 1.000000}, {35000000.000000, -1.000000}, {60000000.000000, 2.000000}, {72000000.000000, -2.000000}, {88000000.000000, 4.000000}, {112000000.000000, 8.000000}, {122000000.000000, 12.000000}, {132000000.000000, 16.000000}, {162000000.000000, 18.000000}, {182000000.000000, -20.000000}, {202000000.000000, 100.000000}},
					},
					{
						Class:  "sample",
						Labels: map[string]string{"other": "test", "label": "41"},
						Attrs:  map[string]string{},
						Values: [][]float64{{0.000000, 1.000000}, {35000000.000000, 1.000000}, {60000000.000000, 2.000000}, {72000000.000000, 3.000000}, {88000000.000000, 4.000000}, {112000000.000000, 5.000000}, {122000000.000000, 6.000000}, {132000000.000000, 7.000000}, {162000000.000000, 8.000000}, {182000000.000000, 9.000000}, {202000000.000000, 10.000000}},
					},
					{
						Class:  "sample",
						Labels: map[string]string{"label": "41", "other": "test"},
						Attrs:  map[string]string{},
						Values: [][]float64{{0.000000, 1.000000}, {40000000.000000, 2.000000}, {82000000.000000, 1.000000}, {110000000.000000, 0.000000}, {129000000.000000, 1.000000}, {159000000.000000, 2.000000}, {192000000.000000, 0.000000}, {205000000.000000, 1.000000}},
					},
				},

				SeriesTests: seriesEqualityTestMap,
			},
			{
				SamplePrefix: BuildGtsList([]FloatGeoTimeSeries{gts1, gts2}),
				GTSResult: []FloatGeoTimeSeries{
					{
						Class:  "sample",
						Labels: map[string]string{"label": "42"},
						Attrs:  map[string]string{},
						Values: [][]float64{{0.000000, 1.000000}, {35000000.000000, -1.000000}, {60000000.000000, 2.000000}, {72000000.000000, -2.000000}, {88000000.000000, 4.000000}, {112000000.000000, 8.000000}, {122000000.000000, 12.000000}, {132000000.000000, 16.000000}, {162000000.000000, 18.000000}, {182000000.000000, -20.000000}, {202000000.000000, 100.000000}},
					},
					{
						Class:  "sample",
						Labels: map[string]string{"other": "test", "label": "41"},
						Attrs:  map[string]string{},
						Values: [][]float64{{0.000000, 1.000000}, {35000000.000000, 1.000000}, {60000000.000000, 2.000000}, {72000000.000000, 3.000000}, {88000000.000000, 4.000000}, {112000000.000000, 5.000000}, {122000000.000000, 6.000000}, {132000000.000000, 7.000000}, {162000000.000000, 8.000000}, {182000000.000000, 9.000000}, {202000000.000000, 10.000000}},
					},
				},

				SeriesTests: seriesEqualityTestMap,
			},
			{
				SamplePrefix: BuildGtsList([]FloatGeoTimeSeries{gts1, gts3}),
				GTSResult: []FloatGeoTimeSeries{
					{
						Class:  "sample",
						Labels: map[string]string{"label": "42"},
						Attrs:  map[string]string{},
						Values: [][]float64{{0.000000, 1.000000}, {35000000.000000, -1.000000}, {60000000.000000, 2.000000}, {72000000.000000, -2.000000}, {88000000.000000, 4.000000}, {112000000.000000, 8.000000}, {122000000.000000, 12.000000}, {132000000.000000, 16.000000}, {162000000.000000, 18.000000}, {182000000.000000, -20.000000}, {202000000.000000, 100.000000}},
					},
					{
						Class:  "sample",
						Labels: map[string]string{"other": "test", "label": "41"},
						Attrs:  map[string]string{},
						Values: [][]float64{{0.000000, 1.000000}, {40000000.000000, 2.000000}, {82000000.000000, 1.000000}, {110000000.000000, 0.000000}, {129000000.000000, 1.000000}, {159000000.000000, 2.000000}, {192000000.000000, 0.000000}, {205000000.000000, 1.000000}},
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
