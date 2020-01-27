package prototests

import (
	"testing"

	"github.com/ovh/erlenmeyer/proto/graphite"
)

//
// @code auto-generated
// Test can be started with
// go test proto/prototests/graphite_pow_test.go proto/prototests/exec_test.go -v
//

// TestGraphitePow process Invert Graphite Unit tests
func TestGraphitePow(t *testing.T) {
	RunTest(t, graphiteGraphitePow, "")
}

var graphiteGraphitePow = []unitTests{

	{
		Plan: []graphite.Function{
			{
				Name:       "pow",
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
						Values: [][]float64{{25000000.000000, 1.000000}, {85000000.000000, -0.037037}, {145000000.000000, 1000.000000}, {205000000.000000, 34858.962963}},
					},
					{
						Class:  "sample",
						Labels: map[string]string{"label": "42"},
						Attrs:  map[string]string{},
						Values: [][]float64{{25000000.000000, 1.000000}, {85000000.000000, -0.037037}, {145000000.000000, 1000.000000}, {205000000.000000, 34858.962963}},
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
						Values: [][]float64{{25000000.000000, 1.000000}, {85000000.000000, -0.037037}, {145000000.000000, 1000.000000}, {205000000.000000, 34858.962963}},
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
						Values: [][]float64{{25000000.000000, 1.000000}, {85000000.000000, -0.037037}, {145000000.000000, 1000.000000}, {205000000.000000, 34858.962963}},
					},
					{
						Class:  "sample",
						Labels: map[string]string{"label": "42"},
						Attrs:  map[string]string{},
						Values: [][]float64{{25000000.000000, 1.000000}, {85000000.000000, -0.037037}, {145000000.000000, 1000.000000}, {205000000.000000, 34858.962963}},
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
						Values: [][]float64{{25000000.000000, 1.000000}, {85000000.000000, -0.037037}, {145000000.000000, 1000.000000}, {205000000.000000, 34858.962963}},
					},
					{
						Class:  "sample",
						Labels: map[string]string{"other": "test", "label": "41"},
						Attrs:  map[string]string{},
						Values: [][]float64{{25000000.000000, 1.000000}, {85000000.000000, 8.000000}, {145000000.000000, 166.375000}, {205000000.000000, 729.000000}},
					},
					{
						Class:  "sample",
						Labels: map[string]string{"other": "test", "label": "41"},
						Attrs:  map[string]string{},
						Values: [][]float64{{25000000.000000, 1.000000}, {85000000.000000, 3.375000}, {145000000.000000, 0.125000}, {205000000.000000, 1.000000}},
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
						Values: [][]float64{{25000000.000000, 1.000000}, {85000000.000000, -0.037037}, {145000000.000000, 1000.000000}, {205000000.000000, 34858.962963}},
					},
					{
						Class:  "sample",
						Labels: map[string]string{"other": "test", "label": "41"},
						Attrs:  map[string]string{},
						Values: [][]float64{{25000000.000000, 1.000000}, {85000000.000000, 8.000000}, {145000000.000000, 166.375000}, {205000000.000000, 729.000000}},
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
						Values: [][]float64{{25000000.000000, 1.000000}, {85000000.000000, -0.037037}, {145000000.000000, 1000.000000}, {205000000.000000, 34858.962963}},
					},
					{
						Class:  "sample",
						Labels: map[string]string{"other": "test", "label": "41"},
						Attrs:  map[string]string{},
						Values: [][]float64{{25000000.000000, 1.000000}, {85000000.000000, 3.375000}, {145000000.000000, 0.125000}, {205000000.000000, 1.000000}},
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
						Values: [][]float64{{25000000.000000, 1.000000}, {85000000.000000, -0.037037}, {145000000.000000, 1000.000000}, {205000000.000000, 34858.962963}},
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
						Values: [][]float64{{0.000000, 1.000000}, {35000000.000000, -1.000000}, {60000000.000000, 8.000000}, {72000000.000000, -8.000000}, {88000000.000000, 64.000000}, {112000000.000000, 512.000000}, {122000000.000000, 1728.000000}, {132000000.000000, 4096.000000}, {162000000.000000, 5832.000000}, {182000000.000000, -8000.000000}, {202000000.000000, 1000000.000000}},
					},
					{
						Class:  "sample",
						Labels: map[string]string{"label": "42"},
						Attrs:  map[string]string{},
						Values: [][]float64{{0.000000, 1.000000}, {35000000.000000, -1.000000}, {60000000.000000, 8.000000}, {72000000.000000, -8.000000}, {88000000.000000, 64.000000}, {112000000.000000, 512.000000}, {122000000.000000, 1728.000000}, {132000000.000000, 4096.000000}, {162000000.000000, 5832.000000}, {182000000.000000, -8000.000000}, {202000000.000000, 1000000.000000}},
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
						Values: [][]float64{{0.000000, 1.000000}, {35000000.000000, -1.000000}, {60000000.000000, 8.000000}, {72000000.000000, -8.000000}, {88000000.000000, 64.000000}, {112000000.000000, 512.000000}, {122000000.000000, 1728.000000}, {132000000.000000, 4096.000000}, {162000000.000000, 5832.000000}, {182000000.000000, -8000.000000}, {202000000.000000, 1000000.000000}},
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
						Values: [][]float64{{0.000000, 1.000000}, {35000000.000000, -1.000000}, {60000000.000000, 8.000000}, {72000000.000000, -8.000000}, {88000000.000000, 64.000000}, {112000000.000000, 512.000000}, {122000000.000000, 1728.000000}, {132000000.000000, 4096.000000}, {162000000.000000, 5832.000000}, {182000000.000000, -8000.000000}, {202000000.000000, 1000000.000000}},
					},
					{
						Class:  "sample",
						Labels: map[string]string{"label": "42"},
						Attrs:  map[string]string{},
						Values: [][]float64{{0.000000, 1.000000}, {35000000.000000, -1.000000}, {60000000.000000, 8.000000}, {72000000.000000, -8.000000}, {88000000.000000, 64.000000}, {112000000.000000, 512.000000}, {122000000.000000, 1728.000000}, {132000000.000000, 4096.000000}, {162000000.000000, 5832.000000}, {182000000.000000, -8000.000000}, {202000000.000000, 1000000.000000}},
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
						Values: [][]float64{{0.000000, 1.000000}, {35000000.000000, -1.000000}, {60000000.000000, 8.000000}, {72000000.000000, -8.000000}, {88000000.000000, 64.000000}, {112000000.000000, 512.000000}, {122000000.000000, 1728.000000}, {132000000.000000, 4096.000000}, {162000000.000000, 5832.000000}, {182000000.000000, -8000.000000}, {202000000.000000, 1000000.000000}},
					},
					{
						Class:  "sample",
						Labels: map[string]string{"other": "test", "label": "41"},
						Attrs:  map[string]string{},
						Values: [][]float64{{0.000000, 1.000000}, {35000000.000000, 1.000000}, {60000000.000000, 8.000000}, {72000000.000000, 27.000000}, {88000000.000000, 64.000000}, {112000000.000000, 125.000000}, {122000000.000000, 216.000000}, {132000000.000000, 343.000000}, {162000000.000000, 512.000000}, {182000000.000000, 729.000000}, {202000000.000000, 1000.000000}},
					},
					{
						Class:  "sample",
						Labels: map[string]string{"other": "test", "label": "41"},
						Attrs:  map[string]string{},
						Values: [][]float64{{0.000000, 1.000000}, {40000000.000000, 8.000000}, {82000000.000000, 1.000000}, {110000000.000000, 0.000000}, {129000000.000000, 1.000000}, {159000000.000000, 8.000000}, {192000000.000000, 0.000000}, {205000000.000000, 1.000000}},
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
						Values: [][]float64{{0.000000, 1.000000}, {35000000.000000, -1.000000}, {60000000.000000, 8.000000}, {72000000.000000, -8.000000}, {88000000.000000, 64.000000}, {112000000.000000, 512.000000}, {122000000.000000, 1728.000000}, {132000000.000000, 4096.000000}, {162000000.000000, 5832.000000}, {182000000.000000, -8000.000000}, {202000000.000000, 1000000.000000}},
					},
					{
						Class:  "sample",
						Labels: map[string]string{"other": "test", "label": "41"},
						Attrs:  map[string]string{},
						Values: [][]float64{{0.000000, 1.000000}, {35000000.000000, 1.000000}, {60000000.000000, 8.000000}, {72000000.000000, 27.000000}, {88000000.000000, 64.000000}, {112000000.000000, 125.000000}, {122000000.000000, 216.000000}, {132000000.000000, 343.000000}, {162000000.000000, 512.000000}, {182000000.000000, 729.000000}, {202000000.000000, 1000.000000}},
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
						Values: [][]float64{{0.000000, 1.000000}, {35000000.000000, -1.000000}, {60000000.000000, 8.000000}, {72000000.000000, -8.000000}, {88000000.000000, 64.000000}, {112000000.000000, 512.000000}, {122000000.000000, 1728.000000}, {132000000.000000, 4096.000000}, {162000000.000000, 5832.000000}, {182000000.000000, -8000.000000}, {202000000.000000, 1000000.000000}},
					},
					{
						Class:  "sample",
						Labels: map[string]string{"other": "test", "label": "41"},
						Attrs:  map[string]string{},
						Values: [][]float64{{0.000000, 1.000000}, {40000000.000000, 8.000000}, {82000000.000000, 1.000000}, {110000000.000000, 0.000000}, {129000000.000000, 1.000000}, {159000000.000000, 8.000000}, {192000000.000000, 0.000000}, {205000000.000000, 1.000000}},
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
						Values: [][]float64{{0.000000, 1.000000}, {35000000.000000, -1.000000}, {60000000.000000, 8.000000}, {72000000.000000, -8.000000}, {88000000.000000, 64.000000}, {112000000.000000, 512.000000}, {122000000.000000, 1728.000000}, {132000000.000000, 4096.000000}, {162000000.000000, 5832.000000}, {182000000.000000, -8000.000000}, {202000000.000000, 1000000.000000}},
					},
				},

				SeriesTests: seriesEqualityTestMap,
			},
		},
	},
}
