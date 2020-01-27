package prototests

import (
	"testing"

	"github.com/ovh/erlenmeyer/proto/graphite"
)

//
// @code auto-generated
// Test can be started with
// go test proto/prototests/graphite_scale_test.go proto/prototests/exec_test.go -v
//

// TestGraphiteScale process Invert Graphite Unit tests
func TestGraphiteScale(t *testing.T) {
	RunTest(t, graphiteGraphiteScale, "")
}

var graphiteGraphiteScale = []unitTests{

	{
		Plan: []graphite.Function{
			{
				Name:       "scale",
				Arguments:  []string{"SWAP", "3"},
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
						Values: [][]float64{},
					},
					{
						Class:  "sample",
						Labels: map[string]string{"label": "42"},
						Attrs:  map[string]string{},
						Values: [][]float64{{25000000.000000, 3.000000}, {85000000.000000, 0.000000}, {145000000.000000, 30.000000}, {205000000.000000, 96.000000}},
					},
					{
						Class:  "sample",
						Labels: map[string]string{"label": "42"},
						Attrs:  map[string]string{},
						Values: [][]float64{{25000000.000000, 3.000000}, {85000000.000000, 0.000000}, {145000000.000000, 30.000000}, {205000000.000000, 96.000000}},
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
						Values: [][]float64{},
					},
					{
						Class:  "sample",
						Labels: map[string]string{"label": "42"},
						Attrs:  map[string]string{},
						Values: [][]float64{{25000000.000000, 3.000000}, {85000000.000000, 0.000000}, {145000000.000000, 30.000000}, {205000000.000000, 96.000000}},
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
						Values: [][]float64{{25000000.000000, 3.000000}, {85000000.000000, 0.000000}, {145000000.000000, 30.000000}, {205000000.000000, 96.000000}},
					},
					{
						Class:  "sample",
						Labels: map[string]string{"label": "42"},
						Attrs:  map[string]string{},
						Values: [][]float64{{25000000.000000, 3.000000}, {85000000.000000, 0.000000}, {145000000.000000, 30.000000}, {205000000.000000, 96.000000}},
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
						Values: [][]float64{{25000000.000000, 3.000000}, {85000000.000000, 0.000000}, {145000000.000000, 30.000000}, {205000000.000000, 96.000000}},
					},
					{
						Class:  "sample",
						Labels: map[string]string{"other": "test", "label": "41"},
						Attrs:  map[string]string{},
						Values: [][]float64{{25000000.000000, 3.000000}, {85000000.000000, 6.000000}, {145000000.000000, 15.000000}, {205000000.000000, 27.000000}},
					},
					{
						Class:  "sample",
						Labels: map[string]string{"label": "41", "other": "test"},
						Attrs:  map[string]string{},
						Values: [][]float64{{25000000.000000, 3.000000}, {85000000.000000, 3.000000}, {145000000.000000, 0.000000}, {205000000.000000, 3.000000}},
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
						Values: [][]float64{{25000000.000000, 3.000000}, {85000000.000000, 0.000000}, {145000000.000000, 30.000000}, {205000000.000000, 96.000000}},
					},
					{
						Class:  "sample",
						Labels: map[string]string{"other": "test", "label": "41"},
						Attrs:  map[string]string{},
						Values: [][]float64{{25000000.000000, 3.000000}, {85000000.000000, 6.000000}, {145000000.000000, 15.000000}, {205000000.000000, 27.000000}},
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
						Values: [][]float64{{25000000.000000, 3.000000}, {85000000.000000, 0.000000}, {145000000.000000, 30.000000}, {205000000.000000, 96.000000}},
					},
					{
						Class:  "sample",
						Labels: map[string]string{"other": "test", "label": "41"},
						Attrs:  map[string]string{},
						Values: [][]float64{{25000000.000000, 3.000000}, {85000000.000000, 3.000000}, {145000000.000000, 0.000000}, {205000000.000000, 3.000000}},
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
						Values: [][]float64{{25000000.000000, 3.000000}, {85000000.000000, 0.000000}, {145000000.000000, 30.000000}, {205000000.000000, 96.000000}},
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
						Values: [][]float64{},
					},
					{
						Class:  "sample",
						Labels: map[string]string{"label": "42"},
						Attrs:  map[string]string{},
						Values: [][]float64{{0.000000, 3.000000}, {35000000.000000, -3.000000}, {60000000.000000, 6.000000}, {72000000.000000, -6.000000}, {88000000.000000, 12.000000}, {112000000.000000, 24.000000}, {122000000.000000, 36.000000}, {132000000.000000, 48.000000}, {162000000.000000, 54.000000}, {182000000.000000, -60.000000}, {202000000.000000, 300.000000}},
					},
					{
						Class:  "sample",
						Labels: map[string]string{"label": "42"},
						Attrs:  map[string]string{},
						Values: [][]float64{{0.000000, 3.000000}, {35000000.000000, -3.000000}, {60000000.000000, 6.000000}, {72000000.000000, -6.000000}, {88000000.000000, 12.000000}, {112000000.000000, 24.000000}, {122000000.000000, 36.000000}, {132000000.000000, 48.000000}, {162000000.000000, 54.000000}, {182000000.000000, -60.000000}, {202000000.000000, 300.000000}},
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
						Values: [][]float64{},
					},
					{
						Class:  "sample",
						Labels: map[string]string{"label": "42"},
						Attrs:  map[string]string{},
						Values: [][]float64{{0.000000, 3.000000}, {35000000.000000, -3.000000}, {60000000.000000, 6.000000}, {72000000.000000, -6.000000}, {88000000.000000, 12.000000}, {112000000.000000, 24.000000}, {122000000.000000, 36.000000}, {132000000.000000, 48.000000}, {162000000.000000, 54.000000}, {182000000.000000, -60.000000}, {202000000.000000, 300.000000}},
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
						Values: [][]float64{{0.000000, 3.000000}, {35000000.000000, -3.000000}, {60000000.000000, 6.000000}, {72000000.000000, -6.000000}, {88000000.000000, 12.000000}, {112000000.000000, 24.000000}, {122000000.000000, 36.000000}, {132000000.000000, 48.000000}, {162000000.000000, 54.000000}, {182000000.000000, -60.000000}, {202000000.000000, 300.000000}},
					},
					{
						Class:  "sample",
						Labels: map[string]string{"label": "42"},
						Attrs:  map[string]string{},
						Values: [][]float64{{0.000000, 3.000000}, {35000000.000000, -3.000000}, {60000000.000000, 6.000000}, {72000000.000000, -6.000000}, {88000000.000000, 12.000000}, {112000000.000000, 24.000000}, {122000000.000000, 36.000000}, {132000000.000000, 48.000000}, {162000000.000000, 54.000000}, {182000000.000000, -60.000000}, {202000000.000000, 300.000000}},
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
						Values: [][]float64{{0.000000, 3.000000}, {35000000.000000, -3.000000}, {60000000.000000, 6.000000}, {72000000.000000, -6.000000}, {88000000.000000, 12.000000}, {112000000.000000, 24.000000}, {122000000.000000, 36.000000}, {132000000.000000, 48.000000}, {162000000.000000, 54.000000}, {182000000.000000, -60.000000}, {202000000.000000, 300.000000}},
					},
					{
						Class:  "sample",
						Labels: map[string]string{"other": "test", "label": "41"},
						Attrs:  map[string]string{},
						Values: [][]float64{{0.000000, 3.000000}, {35000000.000000, 3.000000}, {60000000.000000, 6.000000}, {72000000.000000, 9.000000}, {88000000.000000, 12.000000}, {112000000.000000, 15.000000}, {122000000.000000, 18.000000}, {132000000.000000, 21.000000}, {162000000.000000, 24.000000}, {182000000.000000, 27.000000}, {202000000.000000, 30.000000}},
					},
					{
						Class:  "sample",
						Labels: map[string]string{"other": "test", "label": "41"},
						Attrs:  map[string]string{},
						Values: [][]float64{{0.000000, 3.000000}, {40000000.000000, 6.000000}, {82000000.000000, 3.000000}, {110000000.000000, 0.000000}, {129000000.000000, 3.000000}, {159000000.000000, 6.000000}, {192000000.000000, 0.000000}, {205000000.000000, 3.000000}},
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
						Values: [][]float64{{0.000000, 3.000000}, {35000000.000000, -3.000000}, {60000000.000000, 6.000000}, {72000000.000000, -6.000000}, {88000000.000000, 12.000000}, {112000000.000000, 24.000000}, {122000000.000000, 36.000000}, {132000000.000000, 48.000000}, {162000000.000000, 54.000000}, {182000000.000000, -60.000000}, {202000000.000000, 300.000000}},
					},
					{
						Class:  "sample",
						Labels: map[string]string{"other": "test", "label": "41"},
						Attrs:  map[string]string{},
						Values: [][]float64{{0.000000, 3.000000}, {35000000.000000, 3.000000}, {60000000.000000, 6.000000}, {72000000.000000, 9.000000}, {88000000.000000, 12.000000}, {112000000.000000, 15.000000}, {122000000.000000, 18.000000}, {132000000.000000, 21.000000}, {162000000.000000, 24.000000}, {182000000.000000, 27.000000}, {202000000.000000, 30.000000}},
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
						Values: [][]float64{{0.000000, 3.000000}, {35000000.000000, -3.000000}, {60000000.000000, 6.000000}, {72000000.000000, -6.000000}, {88000000.000000, 12.000000}, {112000000.000000, 24.000000}, {122000000.000000, 36.000000}, {132000000.000000, 48.000000}, {162000000.000000, 54.000000}, {182000000.000000, -60.000000}, {202000000.000000, 300.000000}},
					},
					{
						Class:  "sample",
						Labels: map[string]string{"other": "test", "label": "41"},
						Attrs:  map[string]string{},
						Values: [][]float64{{0.000000, 3.000000}, {40000000.000000, 6.000000}, {82000000.000000, 3.000000}, {110000000.000000, 0.000000}, {129000000.000000, 3.000000}, {159000000.000000, 6.000000}, {192000000.000000, 0.000000}, {205000000.000000, 3.000000}},
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
						Values: [][]float64{{0.000000, 3.000000}, {35000000.000000, -3.000000}, {60000000.000000, 6.000000}, {72000000.000000, -6.000000}, {88000000.000000, 12.000000}, {112000000.000000, 24.000000}, {122000000.000000, 36.000000}, {132000000.000000, 48.000000}, {162000000.000000, 54.000000}, {182000000.000000, -60.000000}, {202000000.000000, 300.000000}},
					},
				},

				SeriesTests: seriesEqualityTestMap,
			},
		},
	},
	{
		Plan: []graphite.Function{
			{
				Name:       "scale",
				Arguments:  []string{"SWAP", "0.01"},
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
						Values: [][]float64{},
					},
					{
						Class:  "sample",
						Labels: map[string]string{"label": "42"},
						Attrs:  map[string]string{},
						Values: [][]float64{{25000000.000000, 0.010000}, {85000000.000000, -0.003333}, {145000000.000000, 0.100000}, {205000000.000000, 0.326667}},
					},
					{
						Class:  "sample",
						Labels: map[string]string{"label": "42"},
						Attrs:  map[string]string{},
						Values: [][]float64{{25000000.000000, 0.010000}, {85000000.000000, -0.003333}, {145000000.000000, 0.100000}, {205000000.000000, 0.326667}},
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
						Values: [][]float64{},
					},
					{
						Class:  "sample",
						Labels: map[string]string{"label": "42"},
						Attrs:  map[string]string{},
						Values: [][]float64{{25000000.000000, 0.010000}, {85000000.000000, -0.003333}, {145000000.000000, 0.100000}, {205000000.000000, 0.326667}},
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
						Values: [][]float64{{25000000.000000, 0.010000}, {85000000.000000, -0.003333}, {145000000.000000, 0.100000}, {205000000.000000, 0.326667}},
					},
					{
						Class:  "sample",
						Labels: map[string]string{"label": "42"},
						Attrs:  map[string]string{},
						Values: [][]float64{{25000000.000000, 0.010000}, {85000000.000000, -0.003333}, {145000000.000000, 0.100000}, {205000000.000000, 0.326667}},
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
						Values: [][]float64{{25000000.000000, 0.010000}, {85000000.000000, -0.003333}, {145000000.000000, 0.100000}, {205000000.000000, 0.326667}},
					},
					{
						Class:  "sample",
						Labels: map[string]string{"other": "test", "label": "41"},
						Attrs:  map[string]string{},
						Values: [][]float64{{25000000.000000, 0.010000}, {85000000.000000, 0.020000}, {145000000.000000, 0.055000}, {205000000.000000, 0.090000}},
					},
					{
						Class:  "sample",
						Labels: map[string]string{"other": "test", "label": "41"},
						Attrs:  map[string]string{},
						Values: [][]float64{{25000000.000000, 0.010000}, {85000000.000000, 0.015000}, {145000000.000000, 0.005000}, {205000000.000000, 0.010000}},
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
						Values: [][]float64{{25000000.000000, 0.010000}, {85000000.000000, -0.003333}, {145000000.000000, 0.100000}, {205000000.000000, 0.326667}},
					},
					{
						Class:  "sample",
						Labels: map[string]string{"other": "test", "label": "41"},
						Attrs:  map[string]string{},
						Values: [][]float64{{25000000.000000, 0.010000}, {85000000.000000, 0.020000}, {145000000.000000, 0.055000}, {205000000.000000, 0.090000}},
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
						Values: [][]float64{{25000000.000000, 0.010000}, {85000000.000000, -0.003333}, {145000000.000000, 0.100000}, {205000000.000000, 0.326667}},
					},
					{
						Class:  "sample",
						Labels: map[string]string{"other": "test", "label": "41"},
						Attrs:  map[string]string{},
						Values: [][]float64{{25000000.000000, 0.010000}, {85000000.000000, 0.015000}, {145000000.000000, 0.005000}, {205000000.000000, 0.010000}},
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
						Values: [][]float64{{25000000.000000, 0.010000}, {85000000.000000, -0.003333}, {145000000.000000, 0.100000}, {205000000.000000, 0.326667}},
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
						Values: [][]float64{},
					},
					{
						Class:  "sample",
						Labels: map[string]string{"label": "42"},
						Attrs:  map[string]string{},
						Values: [][]float64{{0.000000, 0.010000}, {35000000.000000, -0.010000}, {60000000.000000, 0.020000}, {72000000.000000, -0.020000}, {88000000.000000, 0.040000}, {112000000.000000, 0.080000}, {122000000.000000, 0.120000}, {132000000.000000, 0.160000}, {162000000.000000, 0.180000}, {182000000.000000, -0.200000}, {202000000.000000, 1.000000}},
					},
					{
						Class:  "sample",
						Labels: map[string]string{"label": "42"},
						Attrs:  map[string]string{},
						Values: [][]float64{{0.000000, 0.010000}, {35000000.000000, -0.010000}, {60000000.000000, 0.020000}, {72000000.000000, -0.020000}, {88000000.000000, 0.040000}, {112000000.000000, 0.080000}, {122000000.000000, 0.120000}, {132000000.000000, 0.160000}, {162000000.000000, 0.180000}, {182000000.000000, -0.200000}, {202000000.000000, 1.000000}},
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
						Values: [][]float64{},
					},
					{
						Class:  "sample",
						Labels: map[string]string{"label": "42"},
						Attrs:  map[string]string{},
						Values: [][]float64{{0.000000, 0.010000}, {35000000.000000, -0.010000}, {60000000.000000, 0.020000}, {72000000.000000, -0.020000}, {88000000.000000, 0.040000}, {112000000.000000, 0.080000}, {122000000.000000, 0.120000}, {132000000.000000, 0.160000}, {162000000.000000, 0.180000}, {182000000.000000, -0.200000}, {202000000.000000, 1.000000}},
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
						Values: [][]float64{{0.000000, 0.010000}, {35000000.000000, -0.010000}, {60000000.000000, 0.020000}, {72000000.000000, -0.020000}, {88000000.000000, 0.040000}, {112000000.000000, 0.080000}, {122000000.000000, 0.120000}, {132000000.000000, 0.160000}, {162000000.000000, 0.180000}, {182000000.000000, -0.200000}, {202000000.000000, 1.000000}},
					},
					{
						Class:  "sample",
						Labels: map[string]string{"label": "42"},
						Attrs:  map[string]string{},
						Values: [][]float64{{0.000000, 0.010000}, {35000000.000000, -0.010000}, {60000000.000000, 0.020000}, {72000000.000000, -0.020000}, {88000000.000000, 0.040000}, {112000000.000000, 0.080000}, {122000000.000000, 0.120000}, {132000000.000000, 0.160000}, {162000000.000000, 0.180000}, {182000000.000000, -0.200000}, {202000000.000000, 1.000000}},
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
						Values: [][]float64{{0.000000, 0.010000}, {35000000.000000, -0.010000}, {60000000.000000, 0.020000}, {72000000.000000, -0.020000}, {88000000.000000, 0.040000}, {112000000.000000, 0.080000}, {122000000.000000, 0.120000}, {132000000.000000, 0.160000}, {162000000.000000, 0.180000}, {182000000.000000, -0.200000}, {202000000.000000, 1.000000}},
					},
					{
						Class:  "sample",
						Labels: map[string]string{"other": "test", "label": "41"},
						Attrs:  map[string]string{},
						Values: [][]float64{{0.000000, 0.010000}, {35000000.000000, 0.010000}, {60000000.000000, 0.020000}, {72000000.000000, 0.030000}, {88000000.000000, 0.040000}, {112000000.000000, 0.050000}, {122000000.000000, 0.060000}, {132000000.000000, 0.070000}, {162000000.000000, 0.080000}, {182000000.000000, 0.090000}, {202000000.000000, 0.100000}},
					},
					{
						Class:  "sample",
						Labels: map[string]string{"other": "test", "label": "41"},
						Attrs:  map[string]string{},
						Values: [][]float64{{0.000000, 0.010000}, {40000000.000000, 0.020000}, {82000000.000000, 0.010000}, {110000000.000000, 0.000000}, {129000000.000000, 0.010000}, {159000000.000000, 0.020000}, {192000000.000000, 0.000000}, {205000000.000000, 0.010000}},
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
						Values: [][]float64{{0.000000, 0.010000}, {35000000.000000, -0.010000}, {60000000.000000, 0.020000}, {72000000.000000, -0.020000}, {88000000.000000, 0.040000}, {112000000.000000, 0.080000}, {122000000.000000, 0.120000}, {132000000.000000, 0.160000}, {162000000.000000, 0.180000}, {182000000.000000, -0.200000}, {202000000.000000, 1.000000}},
					},
					{
						Class:  "sample",
						Labels: map[string]string{"other": "test", "label": "41"},
						Attrs:  map[string]string{},
						Values: [][]float64{{0.000000, 0.010000}, {35000000.000000, 0.010000}, {60000000.000000, 0.020000}, {72000000.000000, 0.030000}, {88000000.000000, 0.040000}, {112000000.000000, 0.050000}, {122000000.000000, 0.060000}, {132000000.000000, 0.070000}, {162000000.000000, 0.080000}, {182000000.000000, 0.090000}, {202000000.000000, 0.100000}},
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
						Values: [][]float64{{0.000000, 0.010000}, {35000000.000000, -0.010000}, {60000000.000000, 0.020000}, {72000000.000000, -0.020000}, {88000000.000000, 0.040000}, {112000000.000000, 0.080000}, {122000000.000000, 0.120000}, {132000000.000000, 0.160000}, {162000000.000000, 0.180000}, {182000000.000000, -0.200000}, {202000000.000000, 1.000000}},
					},
					{
						Class:  "sample",
						Labels: map[string]string{"other": "test", "label": "41"},
						Attrs:  map[string]string{},
						Values: [][]float64{{0.000000, 0.010000}, {40000000.000000, 0.020000}, {82000000.000000, 0.010000}, {110000000.000000, 0.000000}, {129000000.000000, 0.010000}, {159000000.000000, 0.020000}, {192000000.000000, 0.000000}, {205000000.000000, 0.010000}},
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
						Values: [][]float64{{0.000000, 0.010000}, {35000000.000000, -0.010000}, {60000000.000000, 0.020000}, {72000000.000000, -0.020000}, {88000000.000000, 0.040000}, {112000000.000000, 0.080000}, {122000000.000000, 0.120000}, {132000000.000000, 0.160000}, {162000000.000000, 0.180000}, {182000000.000000, -0.200000}, {202000000.000000, 1.000000}},
					},
				},

				SeriesTests: seriesEqualityTestMap,
			},
		},
	},
}
