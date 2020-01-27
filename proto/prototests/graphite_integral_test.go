package prototests

import (
	"testing"

	"github.com/ovh/erlenmeyer/proto/graphite"
)

//
// @code auto-generated
// Test can be started with
// go test proto/prototests/graphite_integral_test.go proto/prototests/exec_test.go -v
//

// TestGraphiteIntegral process Invert Graphite Unit tests
func TestGraphiteIntegral(t *testing.T) {
	RunTest(t, graphiteGraphiteIntegral, "")
}

var graphiteGraphiteIntegral = []unitTests{

	{
		Plan: []graphite.Function{
			{
				Name:       "integral",
				Arguments:  []string{"SWAP"},
				Parameters: make(map[string]string),
			},
		},
		Samples: []OperatorGTSTest{
			{
				SamplePrefix: BucketizeGtsList(BuildGtsList([]FloatGeoTimeSeries{emptyGTS, gts1, gts1})),
				GTSResult: []FloatGeoTimeSeries{
					{
						Class:  "empty",
						Labels: map[string]string{"label": "42"},
						Attrs:  map[string]string{},
						Values: [][]float64{{-9223372036854775808.000000, 0.000000}},
					},
					{
						Class:  "sample",
						Labels: map[string]string{"label": "42"},
						Attrs:  map[string]string{},
						Values: [][]float64{{25000000.000000, 0.000000}, {85000000.000000, 60.000000}, {145000000.000000, 40.000000}, {205000000.000000, 640.000000}},
					},
					{
						Class:  "sample",
						Labels: map[string]string{"label": "42"},
						Attrs:  map[string]string{},
						Values: [][]float64{{25000000.000000, 0.000000}, {85000000.000000, 60.000000}, {145000000.000000, 40.000000}, {205000000.000000, 640.000000}},
					},
				},

				SeriesTests: seriesEqualityTestMap,
			},
			{
				SamplePrefix: BucketizeGtsList(BuildGtsList([]FloatGeoTimeSeries{emptyGTS, gts1})),
				GTSResult: []FloatGeoTimeSeries{
					{
						Class:  "empty",
						Labels: map[string]string{"label": "42"},
						Attrs:  map[string]string{},
						Values: [][]float64{{-9223372036854775808.000000, 0.000000}},
					},
					{
						Class:  "sample",
						Labels: map[string]string{"label": "42"},
						Attrs:  map[string]string{},
						Values: [][]float64{{25000000.000000, 0.000000}, {85000000.000000, 60.000000}, {145000000.000000, 40.000000}, {205000000.000000, 640.000000}},
					},
				},

				SeriesTests: seriesEqualityTestMap,
			},
			{
				SamplePrefix: BucketizeGtsList(BuildGtsList([]FloatGeoTimeSeries{emptyGTS})),
				GTSResult: []FloatGeoTimeSeries{
					{
						Class:  "empty",
						Labels: map[string]string{"label": "42"},
						Attrs:  map[string]string{},
						Values: [][]float64{{-9223372036854775808.000000, 0.000000}},
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
						Values: [][]float64{{25000000.000000, 0.000000}, {85000000.000000, 60.000000}, {145000000.000000, 40.000000}, {205000000.000000, 640.000000}},
					},
					{
						Class:  "sample",
						Labels: map[string]string{"label": "42"},
						Attrs:  map[string]string{},
						Values: [][]float64{{25000000.000000, 0.000000}, {85000000.000000, 60.000000}, {145000000.000000, 40.000000}, {205000000.000000, 640.000000}},
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
						Values: [][]float64{{25000000.000000, 0.000000}, {85000000.000000, 60.000000}, {145000000.000000, 40.000000}, {205000000.000000, 640.000000}},
					},
					{
						Class:  "sample",
						Labels: map[string]string{"other": "test", "label": "41"},
						Attrs:  map[string]string{},
						Values: [][]float64{{25000000.000000, 0.000000}, {85000000.000000, 60.000000}, {145000000.000000, 180.000000}, {205000000.000000, 510.000000}},
					},
					{
						Class:  "sample",
						Labels: map[string]string{"label": "41", "other": "test"},
						Attrs:  map[string]string{},
						Values: [][]float64{{25000000.000000, 0.000000}, {85000000.000000, 60.000000}, {145000000.000000, 150.000000}, {205000000.000000, 180.000000}},
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
						Values: [][]float64{{25000000.000000, 0.000000}, {85000000.000000, 60.000000}, {145000000.000000, 40.000000}, {205000000.000000, 640.000000}},
					},
					{
						Class:  "sample",
						Labels: map[string]string{"other": "test", "label": "41"},
						Attrs:  map[string]string{},
						Values: [][]float64{{25000000.000000, 0.000000}, {85000000.000000, 60.000000}, {145000000.000000, 180.000000}, {205000000.000000, 510.000000}},
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
						Values: [][]float64{{25000000.000000, 0.000000}, {85000000.000000, 60.000000}, {145000000.000000, 40.000000}, {205000000.000000, 640.000000}},
					},
					{
						Class:  "sample",
						Labels: map[string]string{"other": "test", "label": "41"},
						Attrs:  map[string]string{},
						Values: [][]float64{{25000000.000000, 0.000000}, {85000000.000000, 60.000000}, {145000000.000000, 150.000000}, {205000000.000000, 180.000000}},
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
						Values: [][]float64{{25000000.000000, 0.000000}, {85000000.000000, 60.000000}, {145000000.000000, 40.000000}, {205000000.000000, 640.000000}},
					},
				},

				SeriesTests: seriesEqualityTestMap,
			},
			{
				SamplePrefix: BuildGtsList([]FloatGeoTimeSeries{emptyGTS, gts1, gts1}),
				GTSResult: []FloatGeoTimeSeries{
					{
						Class:  "empty",
						Labels: map[string]string{"label": "42"},
						Attrs:  map[string]string{},
						Values: [][]float64{{-9223372036854775808.000000, 0.000000}},
					},
					{
						Class:  "sample",
						Labels: map[string]string{"label": "42"},
						Attrs:  map[string]string{},
						Values: [][]float64{{0.000000, 0.000000}, {35000000.000000, 35.000000}, {60000000.000000, 10.000000}, {72000000.000000, 34.000000}, {88000000.000000, 2.000000}, {112000000.000000, 98.000000}, {122000000.000000, 178.000000}, {132000000.000000, 298.000000}, {162000000.000000, 778.000000}, {182000000.000000, 1138.000000}, {202000000.000000, 738.000000}},
					},
					{
						Class:  "sample",
						Labels: map[string]string{"label": "42"},
						Attrs:  map[string]string{},
						Values: [][]float64{{0.000000, 0.000000}, {35000000.000000, 35.000000}, {60000000.000000, 10.000000}, {72000000.000000, 34.000000}, {88000000.000000, 2.000000}, {112000000.000000, 98.000000}, {122000000.000000, 178.000000}, {132000000.000000, 298.000000}, {162000000.000000, 778.000000}, {182000000.000000, 1138.000000}, {202000000.000000, 738.000000}},
					},
				},

				SeriesTests: seriesEqualityTestMap,
			},
			{
				SamplePrefix: BuildGtsList([]FloatGeoTimeSeries{emptyGTS, gts1}),
				GTSResult: []FloatGeoTimeSeries{
					{
						Class:  "empty",
						Labels: map[string]string{"label": "42"},
						Attrs:  map[string]string{},
						Values: [][]float64{{-9223372036854775808.000000, 0.000000}},
					},
					{
						Class:  "sample",
						Labels: map[string]string{"label": "42"},
						Attrs:  map[string]string{},
						Values: [][]float64{{0.000000, 0.000000}, {35000000.000000, 35.000000}, {60000000.000000, 10.000000}, {72000000.000000, 34.000000}, {88000000.000000, 2.000000}, {112000000.000000, 98.000000}, {122000000.000000, 178.000000}, {132000000.000000, 298.000000}, {162000000.000000, 778.000000}, {182000000.000000, 1138.000000}, {202000000.000000, 738.000000}},
					},
				},

				SeriesTests: seriesEqualityTestMap,
			},
			{
				SamplePrefix: BuildGtsList([]FloatGeoTimeSeries{emptyGTS}),
				GTSResult: []FloatGeoTimeSeries{
					{
						Class:  "empty",
						Labels: map[string]string{"label": "42"},
						Attrs:  map[string]string{},
						Values: [][]float64{{-9223372036854775808.000000, 0.000000}},
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
						Values: [][]float64{{0.000000, 0.000000}, {35000000.000000, 35.000000}, {60000000.000000, 10.000000}, {72000000.000000, 34.000000}, {88000000.000000, 2.000000}, {112000000.000000, 98.000000}, {122000000.000000, 178.000000}, {132000000.000000, 298.000000}, {162000000.000000, 778.000000}, {182000000.000000, 1138.000000}, {202000000.000000, 738.000000}},
					},
					{
						Class:  "sample",
						Labels: map[string]string{"label": "42"},
						Attrs:  map[string]string{},
						Values: [][]float64{{0.000000, 0.000000}, {35000000.000000, 35.000000}, {60000000.000000, 10.000000}, {72000000.000000, 34.000000}, {88000000.000000, 2.000000}, {112000000.000000, 98.000000}, {122000000.000000, 178.000000}, {132000000.000000, 298.000000}, {162000000.000000, 778.000000}, {182000000.000000, 1138.000000}, {202000000.000000, 738.000000}},
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
						Values: [][]float64{{0.000000, 0.000000}, {35000000.000000, 35.000000}, {60000000.000000, 10.000000}, {72000000.000000, 34.000000}, {88000000.000000, 2.000000}, {112000000.000000, 98.000000}, {122000000.000000, 178.000000}, {132000000.000000, 298.000000}, {162000000.000000, 778.000000}, {182000000.000000, 1138.000000}, {202000000.000000, 738.000000}},
					},
					{
						Class:  "sample",
						Labels: map[string]string{"label": "41", "other": "test"},
						Attrs:  map[string]string{},
						Values: [][]float64{{0.000000, 0.000000}, {35000000.000000, 35.000000}, {60000000.000000, 60.000000}, {72000000.000000, 84.000000}, {88000000.000000, 132.000000}, {112000000.000000, 228.000000}, {122000000.000000, 278.000000}, {132000000.000000, 338.000000}, {162000000.000000, 548.000000}, {182000000.000000, 708.000000}, {202000000.000000, 888.000000}},
					},
					{
						Class:  "sample",
						Labels: map[string]string{"other": "test", "label": "41"},
						Attrs:  map[string]string{},
						Values: [][]float64{{0.000000, 0.000000}, {40000000.000000, 40.000000}, {82000000.000000, 124.000000}, {110000000.000000, 152.000000}, {129000000.000000, 152.000000}, {159000000.000000, 182.000000}, {192000000.000000, 248.000000}, {205000000.000000, 248.000000}},
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
						Values: [][]float64{{0.000000, 0.000000}, {35000000.000000, 35.000000}, {60000000.000000, 10.000000}, {72000000.000000, 34.000000}, {88000000.000000, 2.000000}, {112000000.000000, 98.000000}, {122000000.000000, 178.000000}, {132000000.000000, 298.000000}, {162000000.000000, 778.000000}, {182000000.000000, 1138.000000}, {202000000.000000, 738.000000}},
					},
					{
						Class:  "sample",
						Labels: map[string]string{"other": "test", "label": "41"},
						Attrs:  map[string]string{},
						Values: [][]float64{{0.000000, 0.000000}, {35000000.000000, 35.000000}, {60000000.000000, 60.000000}, {72000000.000000, 84.000000}, {88000000.000000, 132.000000}, {112000000.000000, 228.000000}, {122000000.000000, 278.000000}, {132000000.000000, 338.000000}, {162000000.000000, 548.000000}, {182000000.000000, 708.000000}, {202000000.000000, 888.000000}},
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
						Values: [][]float64{{0.000000, 0.000000}, {35000000.000000, 35.000000}, {60000000.000000, 10.000000}, {72000000.000000, 34.000000}, {88000000.000000, 2.000000}, {112000000.000000, 98.000000}, {122000000.000000, 178.000000}, {132000000.000000, 298.000000}, {162000000.000000, 778.000000}, {182000000.000000, 1138.000000}, {202000000.000000, 738.000000}},
					},
					{
						Class:  "sample",
						Labels: map[string]string{"other": "test", "label": "41"},
						Attrs:  map[string]string{},
						Values: [][]float64{{0.000000, 0.000000}, {40000000.000000, 40.000000}, {82000000.000000, 124.000000}, {110000000.000000, 152.000000}, {129000000.000000, 152.000000}, {159000000.000000, 182.000000}, {192000000.000000, 248.000000}, {205000000.000000, 248.000000}},
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
						Values: [][]float64{{0.000000, 0.000000}, {35000000.000000, 35.000000}, {60000000.000000, 10.000000}, {72000000.000000, 34.000000}, {88000000.000000, 2.000000}, {112000000.000000, 98.000000}, {122000000.000000, 178.000000}, {132000000.000000, 298.000000}, {162000000.000000, 778.000000}, {182000000.000000, 1138.000000}, {202000000.000000, 738.000000}},
					},
				},

				SeriesTests: seriesEqualityTestMap,
			},
		},
	},
}
