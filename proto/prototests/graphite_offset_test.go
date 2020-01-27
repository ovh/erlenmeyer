package prototests

import (
	"testing"

	"github.com/ovh/erlenmeyer/proto/graphite"
)

//
// @code auto-generated
// Test can be started with
// go test proto/prototests/graphite_offset_test.go proto/prototests/exec_test.go -v
//

// TestGraphiteOffset process Invert Graphite Unit tests
func TestGraphiteOffset(t *testing.T) {
	RunTest(t, graphiteGraphiteOffset, "")
}

var graphiteGraphiteOffset = []unitTests{

	{
		Plan: []graphite.Function{
			{
				Name:       "offset",
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
						Values: [][]float64{{25000000.000000, 4.000000}, {85000000.000000, 3.000000}, {145000000.000000, 13.000000}, {205000000.000000, 35.000000}},
					},
					{
						Class:  "sample",
						Labels: map[string]string{"label": "42"},
						Attrs:  map[string]string{},
						Values: [][]float64{{25000000.000000, 4.000000}, {85000000.000000, 3.000000}, {145000000.000000, 13.000000}, {205000000.000000, 35.000000}},
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
						Values: [][]float64{{25000000.000000, 4.000000}, {85000000.000000, 3.000000}, {145000000.000000, 13.000000}, {205000000.000000, 35.000000}},
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
						Values: [][]float64{{25000000.000000, 4.000000}, {85000000.000000, 3.000000}, {145000000.000000, 13.000000}, {205000000.000000, 35.000000}},
					},
					{
						Class:  "sample",
						Labels: map[string]string{"label": "42"},
						Attrs:  map[string]string{},
						Values: [][]float64{{25000000.000000, 4.000000}, {85000000.000000, 3.000000}, {145000000.000000, 13.000000}, {205000000.000000, 35.000000}},
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
						Values: [][]float64{{25000000.000000, 4.000000}, {85000000.000000, 3.000000}, {145000000.000000, 13.000000}, {205000000.000000, 35.000000}},
					},
					{
						Class:  "sample",
						Labels: map[string]string{"other": "test", "label": "41"},
						Attrs:  map[string]string{},
						Values: [][]float64{{25000000.000000, 4.000000}, {85000000.000000, 5.000000}, {145000000.000000, 8.000000}, {205000000.000000, 12.000000}},
					},
					{
						Class:  "sample",
						Labels: map[string]string{"label": "41", "other": "test"},
						Attrs:  map[string]string{},
						Values: [][]float64{{25000000.000000, 4.000000}, {85000000.000000, 4.000000}, {145000000.000000, 3.000000}, {205000000.000000, 4.000000}},
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
						Values: [][]float64{{25000000.000000, 4.000000}, {85000000.000000, 3.000000}, {145000000.000000, 13.000000}, {205000000.000000, 35.000000}},
					},
					{
						Class:  "sample",
						Labels: map[string]string{"label": "41", "other": "test"},
						Attrs:  map[string]string{},
						Values: [][]float64{{25000000.000000, 4.000000}, {85000000.000000, 5.000000}, {145000000.000000, 8.000000}, {205000000.000000, 12.000000}},
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
						Values: [][]float64{{25000000.000000, 4.000000}, {85000000.000000, 3.000000}, {145000000.000000, 13.000000}, {205000000.000000, 35.000000}},
					},
					{
						Class:  "sample",
						Labels: map[string]string{"other": "test", "label": "41"},
						Attrs:  map[string]string{},
						Values: [][]float64{{25000000.000000, 4.000000}, {85000000.000000, 4.000000}, {145000000.000000, 3.000000}, {205000000.000000, 4.000000}},
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
						Values: [][]float64{{25000000.000000, 4.000000}, {85000000.000000, 3.000000}, {145000000.000000, 13.000000}, {205000000.000000, 35.000000}},
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
						Values: [][]float64{{0.000000, 4.000000}, {35000000.000000, 2.000000}, {60000000.000000, 5.000000}, {72000000.000000, 1.000000}, {88000000.000000, 7.000000}, {112000000.000000, 11.000000}, {122000000.000000, 15.000000}, {132000000.000000, 19.000000}, {162000000.000000, 21.000000}, {182000000.000000, -17.000000}, {202000000.000000, 103.000000}},
					},
					{
						Class:  "sample",
						Labels: map[string]string{"label": "42"},
						Attrs:  map[string]string{},
						Values: [][]float64{{0.000000, 4.000000}, {35000000.000000, 2.000000}, {60000000.000000, 5.000000}, {72000000.000000, 1.000000}, {88000000.000000, 7.000000}, {112000000.000000, 11.000000}, {122000000.000000, 15.000000}, {132000000.000000, 19.000000}, {162000000.000000, 21.000000}, {182000000.000000, -17.000000}, {202000000.000000, 103.000000}},
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
						Values: [][]float64{{0.000000, 4.000000}, {35000000.000000, 2.000000}, {60000000.000000, 5.000000}, {72000000.000000, 1.000000}, {88000000.000000, 7.000000}, {112000000.000000, 11.000000}, {122000000.000000, 15.000000}, {132000000.000000, 19.000000}, {162000000.000000, 21.000000}, {182000000.000000, -17.000000}, {202000000.000000, 103.000000}},
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
						Values: [][]float64{{0.000000, 4.000000}, {35000000.000000, 2.000000}, {60000000.000000, 5.000000}, {72000000.000000, 1.000000}, {88000000.000000, 7.000000}, {112000000.000000, 11.000000}, {122000000.000000, 15.000000}, {132000000.000000, 19.000000}, {162000000.000000, 21.000000}, {182000000.000000, -17.000000}, {202000000.000000, 103.000000}},
					},
					{
						Class:  "sample",
						Labels: map[string]string{"label": "42"},
						Attrs:  map[string]string{},
						Values: [][]float64{{0.000000, 4.000000}, {35000000.000000, 2.000000}, {60000000.000000, 5.000000}, {72000000.000000, 1.000000}, {88000000.000000, 7.000000}, {112000000.000000, 11.000000}, {122000000.000000, 15.000000}, {132000000.000000, 19.000000}, {162000000.000000, 21.000000}, {182000000.000000, -17.000000}, {202000000.000000, 103.000000}},
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
						Values: [][]float64{{0.000000, 4.000000}, {35000000.000000, 2.000000}, {60000000.000000, 5.000000}, {72000000.000000, 1.000000}, {88000000.000000, 7.000000}, {112000000.000000, 11.000000}, {122000000.000000, 15.000000}, {132000000.000000, 19.000000}, {162000000.000000, 21.000000}, {182000000.000000, -17.000000}, {202000000.000000, 103.000000}},
					},
					{
						Class:  "sample",
						Labels: map[string]string{"other": "test", "label": "41"},
						Attrs:  map[string]string{},
						Values: [][]float64{{0.000000, 4.000000}, {35000000.000000, 4.000000}, {60000000.000000, 5.000000}, {72000000.000000, 6.000000}, {88000000.000000, 7.000000}, {112000000.000000, 8.000000}, {122000000.000000, 9.000000}, {132000000.000000, 10.000000}, {162000000.000000, 11.000000}, {182000000.000000, 12.000000}, {202000000.000000, 13.000000}},
					},
					{
						Class:  "sample",
						Labels: map[string]string{"other": "test", "label": "41"},
						Attrs:  map[string]string{},
						Values: [][]float64{{0.000000, 4.000000}, {40000000.000000, 5.000000}, {82000000.000000, 4.000000}, {110000000.000000, 3.000000}, {129000000.000000, 4.000000}, {159000000.000000, 5.000000}, {192000000.000000, 3.000000}, {205000000.000000, 4.000000}},
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
						Values: [][]float64{{0.000000, 4.000000}, {35000000.000000, 2.000000}, {60000000.000000, 5.000000}, {72000000.000000, 1.000000}, {88000000.000000, 7.000000}, {112000000.000000, 11.000000}, {122000000.000000, 15.000000}, {132000000.000000, 19.000000}, {162000000.000000, 21.000000}, {182000000.000000, -17.000000}, {202000000.000000, 103.000000}},
					},
					{
						Class:  "sample",
						Labels: map[string]string{"other": "test", "label": "41"},
						Attrs:  map[string]string{},
						Values: [][]float64{{0.000000, 4.000000}, {35000000.000000, 4.000000}, {60000000.000000, 5.000000}, {72000000.000000, 6.000000}, {88000000.000000, 7.000000}, {112000000.000000, 8.000000}, {122000000.000000, 9.000000}, {132000000.000000, 10.000000}, {162000000.000000, 11.000000}, {182000000.000000, 12.000000}, {202000000.000000, 13.000000}},
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
						Values: [][]float64{{0.000000, 4.000000}, {35000000.000000, 2.000000}, {60000000.000000, 5.000000}, {72000000.000000, 1.000000}, {88000000.000000, 7.000000}, {112000000.000000, 11.000000}, {122000000.000000, 15.000000}, {132000000.000000, 19.000000}, {162000000.000000, 21.000000}, {182000000.000000, -17.000000}, {202000000.000000, 103.000000}},
					},
					{
						Class:  "sample",
						Labels: map[string]string{"other": "test", "label": "41"},
						Attrs:  map[string]string{},
						Values: [][]float64{{0.000000, 4.000000}, {40000000.000000, 5.000000}, {82000000.000000, 4.000000}, {110000000.000000, 3.000000}, {129000000.000000, 4.000000}, {159000000.000000, 5.000000}, {192000000.000000, 3.000000}, {205000000.000000, 4.000000}},
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
						Values: [][]float64{{0.000000, 4.000000}, {35000000.000000, 2.000000}, {60000000.000000, 5.000000}, {72000000.000000, 1.000000}, {88000000.000000, 7.000000}, {112000000.000000, 11.000000}, {122000000.000000, 15.000000}, {132000000.000000, 19.000000}, {162000000.000000, 21.000000}, {182000000.000000, -17.000000}, {202000000.000000, 103.000000}},
					},
				},

				SeriesTests: seriesEqualityTestMap,
			},
		},
	},
}
