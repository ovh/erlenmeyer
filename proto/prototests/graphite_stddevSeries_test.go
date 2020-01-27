package prototests

import (
	"testing"

	"github.com/ovh/erlenmeyer/proto/graphite"
)

//
// @code auto-generated
// Test can be started with
// go test proto/prototests/graphite_stddevSeries_test.go proto/prototests/exec_test.go -v
//

// TestGraphiteStddevSeries process Invert Graphite Unit tests
func TestGraphiteStddevSeries(t *testing.T) {
	RunTest(t, graphiteGraphiteStddevSeries, "")
}

var graphiteGraphiteStddevSeries = []unitTests{

	{
		Plan: []graphite.Function{
			{
				Name:       "stddevSeries",
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
						Values: [][]float64{{25000000.000000, 0.000000}, {85000000.000000, 0.000000}, {145000000.000000, 0.000000}, {205000000.000000, 0.000000}},
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
						Values: [][]float64{{25000000.000000, 0.000000}, {85000000.000000, 0.000000}, {145000000.000000, 0.000000}, {205000000.000000, 0.000000}},
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
						Values: [][]float64{{25000000.000000, 0.000000}, {85000000.000000, 0.000000}, {145000000.000000, 0.000000}, {205000000.000000, 0.000000}},
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
						Values: [][]float64{{25000000.000000, 0.000000}, {85000000.000000, 1.228519}, {145000000.000000, 4.752192}, {205000000.000000, 16.466577}},
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
						Values: [][]float64{{25000000.000000, 0.000000}, {85000000.000000, 1.649916}, {145000000.000000, 3.181981}, {205000000.000000, 16.734860}},
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
						Values: [][]float64{{25000000.000000, 0.000000}, {85000000.000000, 1.296362}, {145000000.000000, 6.717514}, {205000000.000000, 22.391715}},
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
						Values: [][]float64{{25000000.000000, 0.000000}, {85000000.000000, 0.000000}, {145000000.000000, 0.000000}, {205000000.000000, 0.000000}},
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
						Values: [][]float64{{0.000000, 0.000000}, {35000000.000000, 0.000000}, {60000000.000000, 0.000000}, {72000000.000000, 0.000000}, {88000000.000000, 0.000000}, {112000000.000000, 0.000000}, {122000000.000000, 0.000000}, {132000000.000000, 0.000000}, {162000000.000000, 0.000000}, {182000000.000000, 0.000000}, {202000000.000000, 0.000000}},
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
						Values: [][]float64{{0.000000, 0.000000}, {35000000.000000, 0.000000}, {60000000.000000, 0.000000}, {72000000.000000, 0.000000}, {88000000.000000, 0.000000}, {112000000.000000, 0.000000}, {122000000.000000, 0.000000}, {132000000.000000, 0.000000}, {162000000.000000, 0.000000}, {182000000.000000, 0.000000}, {202000000.000000, 0.000000}},
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
						Values: [][]float64{{0.000000, 0.000000}, {35000000.000000, 0.000000}, {60000000.000000, 0.000000}, {72000000.000000, 0.000000}, {88000000.000000, 0.000000}, {112000000.000000, 0.000000}, {122000000.000000, 0.000000}, {132000000.000000, 0.000000}, {162000000.000000, 0.000000}, {182000000.000000, 0.000000}, {202000000.000000, 0.000000}},
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
						Values: [][]float64{{0.000000, 0.000000}, {35000000.000000, 1.414214}, {40000000.000000, 0.000000}, {60000000.000000, 0.000000}, {72000000.000000, 3.535534}, {82000000.000000, 0.000000}, {88000000.000000, 0.000000}, {110000000.000000, 0.000000}, {112000000.000000, 2.121320}, {122000000.000000, 4.242641}, {129000000.000000, 0.000000}, {132000000.000000, 6.363961}, {159000000.000000, 0.000000}, {162000000.000000, 7.071068}, {182000000.000000, 20.506097}, {192000000.000000, 0.000000}, {202000000.000000, 63.639610}, {205000000.000000, 0.000000}},
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
						Values: [][]float64{{0.000000, 0.000000}, {35000000.000000, 1.414214}, {60000000.000000, 0.000000}, {72000000.000000, 3.535534}, {88000000.000000, 0.000000}, {112000000.000000, 2.121320}, {122000000.000000, 4.242641}, {132000000.000000, 6.363961}, {162000000.000000, 7.071068}, {182000000.000000, 20.506097}, {202000000.000000, 63.639610}},
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
						Values: [][]float64{{0.000000, 0.000000}, {35000000.000000, 0.000000}, {40000000.000000, 0.000000}, {60000000.000000, 0.000000}, {72000000.000000, 0.000000}, {82000000.000000, 0.000000}, {88000000.000000, 0.000000}, {110000000.000000, 0.000000}, {112000000.000000, 0.000000}, {122000000.000000, 0.000000}, {129000000.000000, 0.000000}, {132000000.000000, 0.000000}, {159000000.000000, 0.000000}, {162000000.000000, 0.000000}, {182000000.000000, 0.000000}, {192000000.000000, 0.000000}, {202000000.000000, 0.000000}, {205000000.000000, 0.000000}},
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
						Values: [][]float64{{0.000000, 0.000000}, {35000000.000000, 0.000000}, {60000000.000000, 0.000000}, {72000000.000000, 0.000000}, {88000000.000000, 0.000000}, {112000000.000000, 0.000000}, {122000000.000000, 0.000000}, {132000000.000000, 0.000000}, {162000000.000000, 0.000000}, {182000000.000000, 0.000000}, {202000000.000000, 0.000000}},
					},
				},

				SeriesTests: seriesEqualityTestMap,
			},
		},
	},
}
