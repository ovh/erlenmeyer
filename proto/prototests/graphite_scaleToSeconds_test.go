package prototests

import (
	"testing"

	"github.com/ovh/erlenmeyer/proto/graphite"
)

//
// @code auto-generated
// Test can be started with
// go test proto/prototests/graphite_scaleToSeconds_test.go proto/prototests/exec_test.go -v
//

// TestGraphiteScaleToSeconds process Invert Graphite Unit tests
func TestGraphiteScaleToSeconds(t *testing.T) {
	RunTest(t, graphiteGraphiteScaleToSeconds, "")
}

var graphiteGraphiteScaleToSeconds = []unitTests{

	{
		Plan: []graphite.Function{
			{
				Name:       "scaleToSeconds",
				Arguments:  []string{"SWAP", "60"},
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
						Values: [][]float64{{25000000.000000, -0.000006}, {85000000.000000, 0.000021}, {145000000.000000, 0.000076}, {205000000.000000, 0.000105}},
					},
					{
						Class:  "sample",
						Labels: map[string]string{"label": "42"},
						Attrs:  map[string]string{},
						Values: [][]float64{{25000000.000000, -0.000006}, {85000000.000000, 0.000021}, {145000000.000000, 0.000076}, {205000000.000000, 0.000105}},
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
						Values: [][]float64{{25000000.000000, -0.000006}, {85000000.000000, 0.000021}, {145000000.000000, 0.000076}, {205000000.000000, 0.000105}},
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
						Values: [][]float64{{25000000.000000, -0.000006}, {85000000.000000, 0.000021}, {145000000.000000, 0.000076}, {205000000.000000, 0.000105}},
					},
					{
						Class:  "sample",
						Labels: map[string]string{"label": "42"},
						Attrs:  map[string]string{},
						Values: [][]float64{{25000000.000000, -0.000006}, {85000000.000000, 0.000021}, {145000000.000000, 0.000076}, {205000000.000000, 0.000105}},
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
						Values: [][]float64{{25000000.000000, -0.000006}, {85000000.000000, 0.000021}, {145000000.000000, 0.000076}, {205000000.000000, 0.000105}},
					},
					{
						Class:  "sample",
						Labels: map[string]string{"other": "test", "label": "41"},
						Attrs:  map[string]string{},
						Values: [][]float64{{25000000.000000, 0.000005}, {85000000.000000, 0.000010}, {145000000.000000, 0.000016}, {205000000.000000, 0.000016}},
					},
					{
						Class:  "sample",
						Labels: map[string]string{"other": "test", "label": "41"},
						Attrs:  map[string]string{},
						Values: [][]float64{{25000000.000000, 0.000002}, {85000000.000000, -0.000001}, {145000000.000000, -0.000001}, {205000000.000000, 0.000002}},
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
						Values: [][]float64{{25000000.000000, -0.000006}, {85000000.000000, 0.000021}, {145000000.000000, 0.000076}, {205000000.000000, 0.000105}},
					},
					{
						Class:  "sample",
						Labels: map[string]string{"label": "41", "other": "test"},
						Attrs:  map[string]string{},
						Values: [][]float64{{25000000.000000, 0.000005}, {85000000.000000, 0.000010}, {145000000.000000, 0.000016}, {205000000.000000, 0.000016}},
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
						Values: [][]float64{{25000000.000000, -0.000006}, {85000000.000000, 0.000021}, {145000000.000000, 0.000076}, {205000000.000000, 0.000105}},
					},
					{
						Class:  "sample",
						Labels: map[string]string{"other": "test", "label": "41"},
						Attrs:  map[string]string{},
						Values: [][]float64{{25000000.000000, 0.000002}, {85000000.000000, -0.000001}, {145000000.000000, -0.000001}, {205000000.000000, 0.000002}},
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
						Values: [][]float64{{25000000.000000, -0.000006}, {85000000.000000, 0.000021}, {145000000.000000, 0.000076}, {205000000.000000, 0.000105}},
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
						Values: [][]float64{{0.000000, -0.000016}, {35000000.000000, 0.000005}, {60000000.000000, -0.000008}, {72000000.000000, 0.000020}, {88000000.000000, 0.000069}, {112000000.000000, 0.000065}, {122000000.000000, 0.000111}, {132000000.000000, 0.000042}, {162000000.000000, -0.000200}, {182000000.000000, 0.000569}, {202000000.000000, 0.001667}},
					},
					{
						Class:  "sample",
						Labels: map[string]string{"label": "42"},
						Attrs:  map[string]string{},
						Values: [][]float64{{0.000000, -0.000016}, {35000000.000000, 0.000005}, {60000000.000000, -0.000008}, {72000000.000000, 0.000020}, {88000000.000000, 0.000069}, {112000000.000000, 0.000065}, {122000000.000000, 0.000111}, {132000000.000000, 0.000042}, {162000000.000000, -0.000200}, {182000000.000000, 0.000569}, {202000000.000000, 0.001667}},
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
						Values: [][]float64{{0.000000, -0.000016}, {35000000.000000, 0.000005}, {60000000.000000, -0.000008}, {72000000.000000, 0.000020}, {88000000.000000, 0.000069}, {112000000.000000, 0.000065}, {122000000.000000, 0.000111}, {132000000.000000, 0.000042}, {162000000.000000, -0.000200}, {182000000.000000, 0.000569}, {202000000.000000, 0.001667}},
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
						Values: [][]float64{{0.000000, -0.000016}, {35000000.000000, 0.000005}, {60000000.000000, -0.000008}, {72000000.000000, 0.000020}, {88000000.000000, 0.000069}, {112000000.000000, 0.000065}, {122000000.000000, 0.000111}, {132000000.000000, 0.000042}, {162000000.000000, -0.000200}, {182000000.000000, 0.000569}, {202000000.000000, 0.001667}},
					},
					{
						Class:  "sample",
						Labels: map[string]string{"label": "42"},
						Attrs:  map[string]string{},
						Values: [][]float64{{0.000000, -0.000016}, {35000000.000000, 0.000005}, {60000000.000000, -0.000008}, {72000000.000000, 0.000020}, {88000000.000000, 0.000069}, {112000000.000000, 0.000065}, {122000000.000000, 0.000111}, {132000000.000000, 0.000042}, {162000000.000000, -0.000200}, {182000000.000000, 0.000569}, {202000000.000000, 0.001667}},
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
						Values: [][]float64{{0.000000, -0.000016}, {35000000.000000, 0.000005}, {60000000.000000, -0.000008}, {72000000.000000, 0.000020}, {88000000.000000, 0.000069}, {112000000.000000, 0.000065}, {122000000.000000, 0.000111}, {132000000.000000, 0.000042}, {162000000.000000, -0.000200}, {182000000.000000, 0.000569}, {202000000.000000, 0.001667}},
					},
					{
						Class:  "sample",
						Labels: map[string]string{"other": "test", "label": "41"},
						Attrs:  map[string]string{},
						Values: [][]float64{{0.000000, 0.000000}, {35000000.000000, 0.000005}, {60000000.000000, 0.000015}, {72000000.000000, 0.000020}, {88000000.000000, 0.000014}, {112000000.000000, 0.000016}, {122000000.000000, 0.000028}, {132000000.000000, 0.000014}, {162000000.000000, 0.000011}, {182000000.000000, 0.000014}, {202000000.000000, 0.000014}},
					},
					{
						Class:  "sample",
						Labels: map[string]string{"other": "test", "label": "41"},
						Attrs:  map[string]string{},
						Values: [][]float64{{0.000000, 0.000007}, {40000000.000000, 0.000000}, {82000000.000000, -0.000008}, {110000000.000000, 0.000000}, {129000000.000000, 0.000011}, {159000000.000000, -0.000004}, {192000000.000000, -0.000006}, {205000000.000000, 0.000021}},
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
						Values: [][]float64{{0.000000, -0.000016}, {35000000.000000, 0.000005}, {60000000.000000, -0.000008}, {72000000.000000, 0.000020}, {88000000.000000, 0.000069}, {112000000.000000, 0.000065}, {122000000.000000, 0.000111}, {132000000.000000, 0.000042}, {162000000.000000, -0.000200}, {182000000.000000, 0.000569}, {202000000.000000, 0.001667}},
					},
					{
						Class:  "sample",
						Labels: map[string]string{"other": "test", "label": "41"},
						Attrs:  map[string]string{},
						Values: [][]float64{{0.000000, 0.000000}, {35000000.000000, 0.000005}, {60000000.000000, 0.000015}, {72000000.000000, 0.000020}, {88000000.000000, 0.000014}, {112000000.000000, 0.000016}, {122000000.000000, 0.000028}, {132000000.000000, 0.000014}, {162000000.000000, 0.000011}, {182000000.000000, 0.000014}, {202000000.000000, 0.000014}},
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
						Values: [][]float64{{0.000000, -0.000016}, {35000000.000000, 0.000005}, {60000000.000000, -0.000008}, {72000000.000000, 0.000020}, {88000000.000000, 0.000069}, {112000000.000000, 0.000065}, {122000000.000000, 0.000111}, {132000000.000000, 0.000042}, {162000000.000000, -0.000200}, {182000000.000000, 0.000569}, {202000000.000000, 0.001667}},
					},
					{
						Class:  "sample",
						Labels: map[string]string{"other": "test", "label": "41"},
						Attrs:  map[string]string{},
						Values: [][]float64{{0.000000, 0.000007}, {40000000.000000, 0.000000}, {82000000.000000, -0.000008}, {110000000.000000, 0.000000}, {129000000.000000, 0.000011}, {159000000.000000, -0.000004}, {192000000.000000, -0.000006}, {205000000.000000, 0.000021}},
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
						Values: [][]float64{{0.000000, -0.000016}, {35000000.000000, 0.000005}, {60000000.000000, -0.000008}, {72000000.000000, 0.000020}, {88000000.000000, 0.000069}, {112000000.000000, 0.000065}, {122000000.000000, 0.000111}, {132000000.000000, 0.000042}, {162000000.000000, -0.000200}, {182000000.000000, 0.000569}, {202000000.000000, 0.001667}},
					},
				},

				SeriesTests: seriesEqualityTestMap,
			},
		},
	},
}
