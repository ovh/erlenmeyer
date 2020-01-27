package prototests

import (
	"testing"

	"github.com/ovh/erlenmeyer/proto/graphite"
)

//
// @code auto-generated
// Test can be started with
// go test proto/prototests/graphite_logarithm_test.go proto/prototests/exec_test.go -v
//

// TestGraphiteLogarithm process Invert Graphite Unit tests
func TestGraphiteLogarithm(t *testing.T) {
	RunTest(t, graphiteGraphiteLogarithm, "")
}

var graphiteGraphiteLogarithm = []unitTests{

	{
		Plan: []graphite.Function{
			{
				Name:       "logarithm",
				Arguments:  []string{"SWAP", "2"},
				Parameters: make(map[string]string),
			},
		},
		Samples: []OperatorGTSTest{
			{
				SamplePrefix: BucketizeGtsList(BuildGtsList([]FloatGeoTimeSeries{emptyGTS, gts1, gts1})),
				SampleSuffix: `[ SWAP mapper.finite 0 0 0 ] MAP
		`,
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
						Values: [][]float64{{25000000.000000, 0.000000}, {145000000.000000, 1.000000}, {205000000.000000, 1.514105}},
					},
					{
						Class:  "sample",
						Labels: map[string]string{"label": "42"},
						Attrs:  map[string]string{},
						Values: [][]float64{{25000000.000000, 0.000000}, {145000000.000000, 1.000000}, {205000000.000000, 1.514105}},
					},
				},

				SeriesTests: seriesEqualityTestMap,
			},
			{
				SamplePrefix: BucketizeGtsList(BuildGtsList([]FloatGeoTimeSeries{emptyGTS, gts1})),
				SampleSuffix: `[ SWAP mapper.finite 0 0 0 ] MAP
		`,
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
						Values: [][]float64{{25000000.000000, 0.000000}, {145000000.000000, 1.000000}, {205000000.000000, 1.514105}},
					},
				},

				SeriesTests: seriesEqualityTestMap,
			},
			{
				SamplePrefix: BucketizeGtsList(BuildGtsList([]FloatGeoTimeSeries{emptyGTS})),
				SampleSuffix: `[ SWAP mapper.finite 0 0 0 ] MAP
		`,
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
				SampleSuffix: `[ SWAP mapper.finite 0 0 0 ] MAP
		`,
				GTSResult: []FloatGeoTimeSeries{
					{
						Class:  "sample",
						Labels: map[string]string{"label": "42"},
						Attrs:  map[string]string{},
						Values: [][]float64{{25000000.000000, 0.000000}, {145000000.000000, 1.000000}, {205000000.000000, 1.514105}},
					},
					{
						Class:  "sample",
						Labels: map[string]string{"label": "42"},
						Attrs:  map[string]string{},
						Values: [][]float64{{25000000.000000, 0.000000}, {145000000.000000, 1.000000}, {205000000.000000, 1.514105}},
					},
				},

				SeriesTests: seriesEqualityTestMap,
			},
			{
				SamplePrefix: BucketizeGtsList(BuildGtsList([]FloatGeoTimeSeries{gts1, gts2, gts3})),
				SampleSuffix: `[ SWAP mapper.finite 0 0 0 ] MAP
		`,
				GTSResult: []FloatGeoTimeSeries{
					{
						Class:  "sample",
						Labels: map[string]string{"label": "42"},
						Attrs:  map[string]string{},
						Values: [][]float64{{25000000.000000, 0.000000}, {145000000.000000, 1.000000}, {205000000.000000, 1.514105}},
					},
					{
						Class:  "sample",
						Labels: map[string]string{"other": "test", "label": "41"},
						Attrs:  map[string]string{},
						Values: [][]float64{{25000000.000000, 0.000000}, {85000000.000000, 0.301030}, {145000000.000000, 0.740363}, {205000000.000000, 0.954243}},
					},
					{
						Class:  "sample",
						Labels: map[string]string{"other": "test", "label": "41"},
						Attrs:  map[string]string{},
						Values: [][]float64{{25000000.000000, 0.000000}, {85000000.000000, 0.176091}, {145000000.000000, -0.301030}, {205000000.000000, 0.000000}},
					},
				},

				SeriesTests: seriesEqualityTestMap,
			},
			{
				SamplePrefix: BucketizeGtsList(BuildGtsList([]FloatGeoTimeSeries{gts1, gts2})),
				SampleSuffix: `[ SWAP mapper.finite 0 0 0 ] MAP
		`,
				GTSResult: []FloatGeoTimeSeries{
					{
						Class:  "sample",
						Labels: map[string]string{"label": "42"},
						Attrs:  map[string]string{},
						Values: [][]float64{{25000000.000000, 0.000000}, {145000000.000000, 1.000000}, {205000000.000000, 1.514105}},
					},
					{
						Class:  "sample",
						Labels: map[string]string{"label": "41", "other": "test"},
						Attrs:  map[string]string{},
						Values: [][]float64{{25000000.000000, 0.000000}, {85000000.000000, 0.301030}, {145000000.000000, 0.740363}, {205000000.000000, 0.954243}},
					},
				},

				SeriesTests: seriesEqualityTestMap,
			},
			{
				SamplePrefix: BucketizeGtsList(BuildGtsList([]FloatGeoTimeSeries{gts1, gts3})),
				SampleSuffix: `[ SWAP mapper.finite 0 0 0 ] MAP
		`,
				GTSResult: []FloatGeoTimeSeries{
					{
						Class:  "sample",
						Labels: map[string]string{"label": "42"},
						Attrs:  map[string]string{},
						Values: [][]float64{{25000000.000000, 0.000000}, {145000000.000000, 1.000000}, {205000000.000000, 1.514105}},
					},
					{
						Class:  "sample",
						Labels: map[string]string{"other": "test", "label": "41"},
						Attrs:  map[string]string{},
						Values: [][]float64{{25000000.000000, 0.000000}, {85000000.000000, 0.176091}, {145000000.000000, -0.301030}, {205000000.000000, 0.000000}},
					},
				},

				SeriesTests: seriesEqualityTestMap,
			},
			{
				SamplePrefix: BucketizeGtsList(BuildGtsList([]FloatGeoTimeSeries{gts1})),
				SampleSuffix: `[ SWAP mapper.finite 0 0 0 ] MAP
		`,
				GTSResult: []FloatGeoTimeSeries{
					{
						Class:  "sample",
						Labels: map[string]string{"label": "42"},
						Attrs:  map[string]string{},
						Values: [][]float64{{25000000.000000, 0.000000}, {145000000.000000, 1.000000}, {205000000.000000, 1.514105}},
					},
				},

				SeriesTests: seriesEqualityTestMap,
			},
			{
				SamplePrefix: BuildGtsList([]FloatGeoTimeSeries{emptyGTS, gts1, gts1}),
				SampleSuffix: `[ SWAP mapper.finite 0 0 0 ] MAP
		`,
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
						Values: [][]float64{{0.000000, 0.000000}, {60000000.000000, 0.301030}, {88000000.000000, 0.602060}, {112000000.000000, 0.903090}, {122000000.000000, 1.079181}, {132000000.000000, 1.204120}, {162000000.000000, 1.255273}, {202000000.000000, 2.000000}},
					},
					{
						Class:  "sample",
						Labels: map[string]string{"label": "42"},
						Attrs:  map[string]string{},
						Values: [][]float64{{0.000000, 0.000000}, {60000000.000000, 0.301030}, {88000000.000000, 0.602060}, {112000000.000000, 0.903090}, {122000000.000000, 1.079181}, {132000000.000000, 1.204120}, {162000000.000000, 1.255273}, {202000000.000000, 2.000000}},
					},
				},

				SeriesTests: seriesEqualityTestMap,
			},
			{
				SamplePrefix: BuildGtsList([]FloatGeoTimeSeries{emptyGTS, gts1}),
				SampleSuffix: `[ SWAP mapper.finite 0 0 0 ] MAP
		`,
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
						Values: [][]float64{{0.000000, 0.000000}, {60000000.000000, 0.301030}, {88000000.000000, 0.602060}, {112000000.000000, 0.903090}, {122000000.000000, 1.079181}, {132000000.000000, 1.204120}, {162000000.000000, 1.255273}, {202000000.000000, 2.000000}},
					},
				},

				SeriesTests: seriesEqualityTestMap,
			},
			{
				SamplePrefix: BuildGtsList([]FloatGeoTimeSeries{emptyGTS}),
				SampleSuffix: `[ SWAP mapper.finite 0 0 0 ] MAP
		`,
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
				SampleSuffix: `[ SWAP mapper.finite 0 0 0 ] MAP
		`,
				GTSResult: []FloatGeoTimeSeries{
					{
						Class:  "sample",
						Labels: map[string]string{"label": "42"},
						Attrs:  map[string]string{},
						Values: [][]float64{{0.000000, 0.000000}, {60000000.000000, 0.301030}, {88000000.000000, 0.602060}, {112000000.000000, 0.903090}, {122000000.000000, 1.079181}, {132000000.000000, 1.204120}, {162000000.000000, 1.255273}, {202000000.000000, 2.000000}},
					},
					{
						Class:  "sample",
						Labels: map[string]string{"label": "42"},
						Attrs:  map[string]string{},
						Values: [][]float64{{0.000000, 0.000000}, {60000000.000000, 0.301030}, {88000000.000000, 0.602060}, {112000000.000000, 0.903090}, {122000000.000000, 1.079181}, {132000000.000000, 1.204120}, {162000000.000000, 1.255273}, {202000000.000000, 2.000000}},
					},
				},

				SeriesTests: seriesEqualityTestMap,
			},
			{
				SamplePrefix: BuildGtsList([]FloatGeoTimeSeries{gts1, gts2, gts3}),
				SampleSuffix: `[ SWAP mapper.finite 0 0 0 ] MAP
		`,
				GTSResult: []FloatGeoTimeSeries{
					{
						Class:  "sample",
						Labels: map[string]string{"label": "42"},
						Attrs:  map[string]string{},
						Values: [][]float64{{0.000000, 0.000000}, {60000000.000000, 0.301030}, {88000000.000000, 0.602060}, {112000000.000000, 0.903090}, {122000000.000000, 1.079181}, {132000000.000000, 1.204120}, {162000000.000000, 1.255273}, {202000000.000000, 2.000000}},
					},
					{
						Class:  "sample",
						Labels: map[string]string{"other": "test", "label": "41"},
						Attrs:  map[string]string{},
						Values: [][]float64{{0.000000, 0.000000}, {35000000.000000, 0.000000}, {60000000.000000, 0.301030}, {72000000.000000, 0.477121}, {88000000.000000, 0.602060}, {112000000.000000, 0.698970}, {122000000.000000, 0.778151}, {132000000.000000, 0.845098}, {162000000.000000, 0.903090}, {182000000.000000, 0.954243}, {202000000.000000, 1.000000}},
					},
					{
						Class:  "sample",
						Labels: map[string]string{"label": "41", "other": "test"},
						Attrs:  map[string]string{},
						Values: [][]float64{{0.000000, 0.000000}, {40000000.000000, 0.301030}, {82000000.000000, 0.000000}, {129000000.000000, 0.000000}, {159000000.000000, 0.301030}, {205000000.000000, 0.000000}},
					},
				},

				SeriesTests: seriesEqualityTestMap,
			},
			{
				SamplePrefix: BuildGtsList([]FloatGeoTimeSeries{gts1, gts2}),
				SampleSuffix: `[ SWAP mapper.finite 0 0 0 ] MAP
		`,
				GTSResult: []FloatGeoTimeSeries{
					{
						Class:  "sample",
						Labels: map[string]string{"label": "42"},
						Attrs:  map[string]string{},
						Values: [][]float64{{0.000000, 0.000000}, {60000000.000000, 0.301030}, {88000000.000000, 0.602060}, {112000000.000000, 0.903090}, {122000000.000000, 1.079181}, {132000000.000000, 1.204120}, {162000000.000000, 1.255273}, {202000000.000000, 2.000000}},
					},
					{
						Class:  "sample",
						Labels: map[string]string{"label": "41", "other": "test"},
						Attrs:  map[string]string{},
						Values: [][]float64{{0.000000, 0.000000}, {35000000.000000, 0.000000}, {60000000.000000, 0.301030}, {72000000.000000, 0.477121}, {88000000.000000, 0.602060}, {112000000.000000, 0.698970}, {122000000.000000, 0.778151}, {132000000.000000, 0.845098}, {162000000.000000, 0.903090}, {182000000.000000, 0.954243}, {202000000.000000, 1.000000}},
					},
				},

				SeriesTests: seriesEqualityTestMap,
			},
			{
				SamplePrefix: BuildGtsList([]FloatGeoTimeSeries{gts1, gts3}),
				SampleSuffix: `[ SWAP mapper.finite 0 0 0 ] MAP
		`,
				GTSResult: []FloatGeoTimeSeries{
					{
						Class:  "sample",
						Labels: map[string]string{"label": "42"},
						Attrs:  map[string]string{},
						Values: [][]float64{{0.000000, 0.000000}, {60000000.000000, 0.301030}, {88000000.000000, 0.602060}, {112000000.000000, 0.903090}, {122000000.000000, 1.079181}, {132000000.000000, 1.204120}, {162000000.000000, 1.255273}, {202000000.000000, 2.000000}},
					},
					{
						Class:  "sample",
						Labels: map[string]string{"label": "41", "other": "test"},
						Attrs:  map[string]string{},
						Values: [][]float64{{0.000000, 0.000000}, {40000000.000000, 0.301030}, {82000000.000000, 0.000000}, {129000000.000000, 0.000000}, {159000000.000000, 0.301030}, {205000000.000000, 0.000000}},
					},
				},

				SeriesTests: seriesEqualityTestMap,
			},
			{
				SamplePrefix: BuildGtsList([]FloatGeoTimeSeries{gts1}),
				SampleSuffix: `[ SWAP mapper.finite 0 0 0 ] MAP
		`,
				GTSResult: []FloatGeoTimeSeries{
					{
						Class:  "sample",
						Labels: map[string]string{"label": "42"},
						Attrs:  map[string]string{},
						Values: [][]float64{{0.000000, 0.000000}, {60000000.000000, 0.301030}, {88000000.000000, 0.602060}, {112000000.000000, 0.903090}, {122000000.000000, 1.079181}, {132000000.000000, 1.204120}, {162000000.000000, 1.255273}, {202000000.000000, 2.000000}},
					},
				},

				SeriesTests: seriesEqualityTestMap,
			},
		},
	},
}
