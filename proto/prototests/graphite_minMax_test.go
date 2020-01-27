package prototests

import (
	"testing"

	"github.com/ovh/erlenmeyer/proto/graphite"
)

//
// @code auto-generated
// Test can be started with
// go test proto/prototests/graphite_minMax_test.go proto/prototests/exec_test.go -v
//

// TestGraphiteMinMax process Invert Graphite Unit tests
func TestGraphiteMinMax(t *testing.T) {
	RunTest(t, graphiteGraphiteMinMax, "")
}

var graphiteGraphiteMinMax = []unitTests{

	{
		Plan: []graphite.Function{
			{
				Name:       "minMax",
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
						Values: [][]float64{},
					},
					{
						Class:  "sample",
						Labels: map[string]string{"label": "42"},
						Attrs:  map[string]string{},
						Values: [][]float64{{205000000.000000, 1.000000}, {145000000.000000, 0.313131}, {85000000.000000, 0.000000}, {25000000.000000, 0.040404}},
					},
					{
						Class:  "sample",
						Labels: map[string]string{"label": "42"},
						Attrs:  map[string]string{},
						Values: [][]float64{{205000000.000000, 1.000000}, {145000000.000000, 0.313131}, {85000000.000000, 0.000000}, {25000000.000000, 0.040404}},
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
						Values: [][]float64{{205000000.000000, 1.000000}, {145000000.000000, 0.313131}, {85000000.000000, 0.000000}, {25000000.000000, 0.040404}},
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
						Values: [][]float64{{205000000.000000, 1.000000}, {145000000.000000, 0.313131}, {85000000.000000, 0.000000}, {25000000.000000, 0.040404}},
					},
					{
						Class:  "sample",
						Labels: map[string]string{"label": "42"},
						Attrs:  map[string]string{},
						Values: [][]float64{{205000000.000000, 1.000000}, {145000000.000000, 0.313131}, {85000000.000000, 0.000000}, {25000000.000000, 0.040404}},
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
						Values: [][]float64{{205000000.000000, 1.000000}, {145000000.000000, 0.313131}, {85000000.000000, 0.000000}, {25000000.000000, 0.040404}},
					},
					{
						Class:  "sample",
						Labels: map[string]string{"other": "test", "label": "41"},
						Attrs:  map[string]string{},
						Values: [][]float64{{205000000.000000, 1.000000}, {145000000.000000, 0.562500}, {85000000.000000, 0.125000}, {25000000.000000, 0.000000}},
					},
					{
						Class:  "sample",
						Labels: map[string]string{"other": "test", "label": "41"},
						Attrs:  map[string]string{},
						Values: [][]float64{{205000000.000000, 0.500000}, {145000000.000000, 0.000000}, {85000000.000000, 1.000000}, {25000000.000000, 0.500000}},
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
						Values: [][]float64{{205000000.000000, 1.000000}, {145000000.000000, 0.313131}, {85000000.000000, 0.000000}, {25000000.000000, 0.040404}},
					},
					{
						Class:  "sample",
						Labels: map[string]string{"other": "test", "label": "41"},
						Attrs:  map[string]string{},
						Values: [][]float64{{205000000.000000, 1.000000}, {145000000.000000, 0.562500}, {85000000.000000, 0.125000}, {25000000.000000, 0.000000}},
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
						Values: [][]float64{{205000000.000000, 1.000000}, {145000000.000000, 0.313131}, {85000000.000000, 0.000000}, {25000000.000000, 0.040404}},
					},
					{
						Class:  "sample",
						Labels: map[string]string{"other": "test", "label": "41"},
						Attrs:  map[string]string{},
						Values: [][]float64{{205000000.000000, 0.500000}, {145000000.000000, 0.000000}, {85000000.000000, 1.000000}, {25000000.000000, 0.500000}},
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
						Values: [][]float64{{205000000.000000, 1.000000}, {145000000.000000, 0.313131}, {85000000.000000, 0.000000}, {25000000.000000, 0.040404}},
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
						Values: [][]float64{{0.000000, 0.175000}, {35000000.000000, 0.158333}, {60000000.000000, 0.183333}, {72000000.000000, 0.150000}, {88000000.000000, 0.200000}, {112000000.000000, 0.233333}, {122000000.000000, 0.266667}, {132000000.000000, 0.300000}, {162000000.000000, 0.316667}, {182000000.000000, 0.000000}, {202000000.000000, 1.000000}},
					},
					{
						Class:  "sample",
						Labels: map[string]string{"label": "42"},
						Attrs:  map[string]string{},
						Values: [][]float64{{0.000000, 0.175000}, {35000000.000000, 0.158333}, {60000000.000000, 0.183333}, {72000000.000000, 0.150000}, {88000000.000000, 0.200000}, {112000000.000000, 0.233333}, {122000000.000000, 0.266667}, {132000000.000000, 0.300000}, {162000000.000000, 0.316667}, {182000000.000000, 0.000000}, {202000000.000000, 1.000000}},
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
						Values: [][]float64{{0.000000, 0.175000}, {35000000.000000, 0.158333}, {60000000.000000, 0.183333}, {72000000.000000, 0.150000}, {88000000.000000, 0.200000}, {112000000.000000, 0.233333}, {122000000.000000, 0.266667}, {132000000.000000, 0.300000}, {162000000.000000, 0.316667}, {182000000.000000, 0.000000}, {202000000.000000, 1.000000}},
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
						Values: [][]float64{{0.000000, 0.175000}, {35000000.000000, 0.158333}, {60000000.000000, 0.183333}, {72000000.000000, 0.150000}, {88000000.000000, 0.200000}, {112000000.000000, 0.233333}, {122000000.000000, 0.266667}, {132000000.000000, 0.300000}, {162000000.000000, 0.316667}, {182000000.000000, 0.000000}, {202000000.000000, 1.000000}},
					},
					{
						Class:  "sample",
						Labels: map[string]string{"label": "42"},
						Attrs:  map[string]string{},
						Values: [][]float64{{0.000000, 0.175000}, {35000000.000000, 0.158333}, {60000000.000000, 0.183333}, {72000000.000000, 0.150000}, {88000000.000000, 0.200000}, {112000000.000000, 0.233333}, {122000000.000000, 0.266667}, {132000000.000000, 0.300000}, {162000000.000000, 0.316667}, {182000000.000000, 0.000000}, {202000000.000000, 1.000000}},
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
						Values: [][]float64{{0.000000, 0.175000}, {35000000.000000, 0.158333}, {60000000.000000, 0.183333}, {72000000.000000, 0.150000}, {88000000.000000, 0.200000}, {112000000.000000, 0.233333}, {122000000.000000, 0.266667}, {132000000.000000, 0.300000}, {162000000.000000, 0.316667}, {182000000.000000, 0.000000}, {202000000.000000, 1.000000}},
					},
					{
						Class:  "sample",
						Labels: map[string]string{"other": "test", "label": "41"},
						Attrs:  map[string]string{},
						Values: [][]float64{{0.000000, 0.000000}, {35000000.000000, 0.000000}, {60000000.000000, 0.111111}, {72000000.000000, 0.222222}, {88000000.000000, 0.333333}, {112000000.000000, 0.444444}, {122000000.000000, 0.555556}, {132000000.000000, 0.666667}, {162000000.000000, 0.777778}, {182000000.000000, 0.888889}, {202000000.000000, 1.000000}},
					},
					{
						Class:  "sample",
						Labels: map[string]string{"label": "41", "other": "test"},
						Attrs:  map[string]string{},
						Values: [][]float64{{0.000000, 0.500000}, {40000000.000000, 1.000000}, {82000000.000000, 0.500000}, {110000000.000000, 0.000000}, {129000000.000000, 0.500000}, {159000000.000000, 1.000000}, {192000000.000000, 0.000000}, {205000000.000000, 0.500000}},
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
						Values: [][]float64{{0.000000, 0.175000}, {35000000.000000, 0.158333}, {60000000.000000, 0.183333}, {72000000.000000, 0.150000}, {88000000.000000, 0.200000}, {112000000.000000, 0.233333}, {122000000.000000, 0.266667}, {132000000.000000, 0.300000}, {162000000.000000, 0.316667}, {182000000.000000, 0.000000}, {202000000.000000, 1.000000}},
					},
					{
						Class:  "sample",
						Labels: map[string]string{"label": "41", "other": "test"},
						Attrs:  map[string]string{},
						Values: [][]float64{{0.000000, 0.000000}, {35000000.000000, 0.000000}, {60000000.000000, 0.111111}, {72000000.000000, 0.222222}, {88000000.000000, 0.333333}, {112000000.000000, 0.444444}, {122000000.000000, 0.555556}, {132000000.000000, 0.666667}, {162000000.000000, 0.777778}, {182000000.000000, 0.888889}, {202000000.000000, 1.000000}},
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
						Values: [][]float64{{0.000000, 0.175000}, {35000000.000000, 0.158333}, {60000000.000000, 0.183333}, {72000000.000000, 0.150000}, {88000000.000000, 0.200000}, {112000000.000000, 0.233333}, {122000000.000000, 0.266667}, {132000000.000000, 0.300000}, {162000000.000000, 0.316667}, {182000000.000000, 0.000000}, {202000000.000000, 1.000000}},
					},
					{
						Class:  "sample",
						Labels: map[string]string{"other": "test", "label": "41"},
						Attrs:  map[string]string{},
						Values: [][]float64{{0.000000, 0.500000}, {40000000.000000, 1.000000}, {82000000.000000, 0.500000}, {110000000.000000, 0.000000}, {129000000.000000, 0.500000}, {159000000.000000, 1.000000}, {192000000.000000, 0.000000}, {205000000.000000, 0.500000}},
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
						Values: [][]float64{{0.000000, 0.175000}, {35000000.000000, 0.158333}, {60000000.000000, 0.183333}, {72000000.000000, 0.150000}, {88000000.000000, 0.200000}, {112000000.000000, 0.233333}, {122000000.000000, 0.266667}, {132000000.000000, 0.300000}, {162000000.000000, 0.316667}, {182000000.000000, 0.000000}, {202000000.000000, 1.000000}},
					},
				},

				SeriesTests: seriesEqualityTestMap,
			},
		},
	},
}
