package prototests

import (
	"testing"

	"github.com/ovh/erlenmeyer/proto/graphite"
)

//
// @code auto-generated
// Test can be started with
// go test proto/prototests/graphite_timeShift_test.go proto/prototests/exec_test.go -v
//

// TestGraphiteTimeShift process Invert Graphite Unit tests
func TestGraphiteTimeShift(t *testing.T) {
	RunTest(t, graphiteGraphiteTimeShift, "")
}

var graphiteGraphiteTimeShift = []unitTests{

	{
		Plan: []graphite.Function{
			{
				Name:       "timeShift",
				Arguments:  []string{"SWAP", "1h"},
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
						Values: [][]float64{{-3395000000.000000, 32.666667}, {-3455000000.000000, 10.000000}, {-3515000000.000000, -0.333333}, {-3575000000.000000, 1.000000}},
					},
					{
						Class:  "sample",
						Labels: map[string]string{"label": "42"},
						Attrs:  map[string]string{},
						Values: [][]float64{{-3395000000.000000, 32.666667}, {-3455000000.000000, 10.000000}, {-3515000000.000000, -0.333333}, {-3575000000.000000, 1.000000}},
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
						Values: [][]float64{{-3395000000.000000, 32.666667}, {-3455000000.000000, 10.000000}, {-3515000000.000000, -0.333333}, {-3575000000.000000, 1.000000}},
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
						Values: [][]float64{{-3395000000.000000, 32.666667}, {-3455000000.000000, 10.000000}, {-3515000000.000000, -0.333333}, {-3575000000.000000, 1.000000}},
					},
					{
						Class:  "sample",
						Labels: map[string]string{"label": "42"},
						Attrs:  map[string]string{},
						Values: [][]float64{{-3395000000.000000, 32.666667}, {-3455000000.000000, 10.000000}, {-3515000000.000000, -0.333333}, {-3575000000.000000, 1.000000}},
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
						Values: [][]float64{{-3395000000.000000, 32.666667}, {-3455000000.000000, 10.000000}, {-3515000000.000000, -0.333333}, {-3575000000.000000, 1.000000}},
					},
					{
						Class:  "sample",
						Labels: map[string]string{"other": "test", "label": "41"},
						Attrs:  map[string]string{},
						Values: [][]float64{{-3395000000.000000, 9.000000}, {-3455000000.000000, 5.500000}, {-3515000000.000000, 2.000000}, {-3575000000.000000, 1.000000}},
					},
					{
						Class:  "sample",
						Labels: map[string]string{"other": "test", "label": "41"},
						Attrs:  map[string]string{},
						Values: [][]float64{{-3395000000.000000, 1.000000}, {-3455000000.000000, 0.500000}, {-3515000000.000000, 1.500000}, {-3575000000.000000, 1.000000}},
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
						Values: [][]float64{{-3395000000.000000, 32.666667}, {-3455000000.000000, 10.000000}, {-3515000000.000000, -0.333333}, {-3575000000.000000, 1.000000}},
					},
					{
						Class:  "sample",
						Labels: map[string]string{"other": "test", "label": "41"},
						Attrs:  map[string]string{},
						Values: [][]float64{{-3395000000.000000, 9.000000}, {-3455000000.000000, 5.500000}, {-3515000000.000000, 2.000000}, {-3575000000.000000, 1.000000}},
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
						Values: [][]float64{{-3395000000.000000, 32.666667}, {-3455000000.000000, 10.000000}, {-3515000000.000000, -0.333333}, {-3575000000.000000, 1.000000}},
					},
					{
						Class:  "sample",
						Labels: map[string]string{"other": "test", "label": "41"},
						Attrs:  map[string]string{},
						Values: [][]float64{{-3395000000.000000, 1.000000}, {-3455000000.000000, 0.500000}, {-3515000000.000000, 1.500000}, {-3575000000.000000, 1.000000}},
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
						Values: [][]float64{{-3395000000.000000, 32.666667}, {-3455000000.000000, 10.000000}, {-3515000000.000000, -0.333333}, {-3575000000.000000, 1.000000}},
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
						Values: [][]float64{{-3600000000.000000, 1.000000}, {-3565000000.000000, -1.000000}, {-3540000000.000000, 2.000000}, {-3528000000.000000, -2.000000}, {-3512000000.000000, 4.000000}, {-3488000000.000000, 8.000000}, {-3478000000.000000, 12.000000}, {-3468000000.000000, 16.000000}, {-3438000000.000000, 18.000000}, {-3418000000.000000, -20.000000}, {-3398000000.000000, 100.000000}},
					},
					{
						Class:  "sample",
						Labels: map[string]string{"label": "42"},
						Attrs:  map[string]string{},
						Values: [][]float64{{-3600000000.000000, 1.000000}, {-3565000000.000000, -1.000000}, {-3540000000.000000, 2.000000}, {-3528000000.000000, -2.000000}, {-3512000000.000000, 4.000000}, {-3488000000.000000, 8.000000}, {-3478000000.000000, 12.000000}, {-3468000000.000000, 16.000000}, {-3438000000.000000, 18.000000}, {-3418000000.000000, -20.000000}, {-3398000000.000000, 100.000000}},
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
						Values: [][]float64{{-3600000000.000000, 1.000000}, {-3565000000.000000, -1.000000}, {-3540000000.000000, 2.000000}, {-3528000000.000000, -2.000000}, {-3512000000.000000, 4.000000}, {-3488000000.000000, 8.000000}, {-3478000000.000000, 12.000000}, {-3468000000.000000, 16.000000}, {-3438000000.000000, 18.000000}, {-3418000000.000000, -20.000000}, {-3398000000.000000, 100.000000}},
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
						Values: [][]float64{{-3600000000.000000, 1.000000}, {-3565000000.000000, -1.000000}, {-3540000000.000000, 2.000000}, {-3528000000.000000, -2.000000}, {-3512000000.000000, 4.000000}, {-3488000000.000000, 8.000000}, {-3478000000.000000, 12.000000}, {-3468000000.000000, 16.000000}, {-3438000000.000000, 18.000000}, {-3418000000.000000, -20.000000}, {-3398000000.000000, 100.000000}},
					},
					{
						Class:  "sample",
						Labels: map[string]string{"label": "42"},
						Attrs:  map[string]string{},
						Values: [][]float64{{-3600000000.000000, 1.000000}, {-3565000000.000000, -1.000000}, {-3540000000.000000, 2.000000}, {-3528000000.000000, -2.000000}, {-3512000000.000000, 4.000000}, {-3488000000.000000, 8.000000}, {-3478000000.000000, 12.000000}, {-3468000000.000000, 16.000000}, {-3438000000.000000, 18.000000}, {-3418000000.000000, -20.000000}, {-3398000000.000000, 100.000000}},
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
						Values: [][]float64{{-3600000000.000000, 1.000000}, {-3565000000.000000, -1.000000}, {-3540000000.000000, 2.000000}, {-3528000000.000000, -2.000000}, {-3512000000.000000, 4.000000}, {-3488000000.000000, 8.000000}, {-3478000000.000000, 12.000000}, {-3468000000.000000, 16.000000}, {-3438000000.000000, 18.000000}, {-3418000000.000000, -20.000000}, {-3398000000.000000, 100.000000}},
					},
					{
						Class:  "sample",
						Labels: map[string]string{"other": "test", "label": "41"},
						Attrs:  map[string]string{},
						Values: [][]float64{{-3600000000.000000, 1.000000}, {-3565000000.000000, 1.000000}, {-3540000000.000000, 2.000000}, {-3528000000.000000, 3.000000}, {-3512000000.000000, 4.000000}, {-3488000000.000000, 5.000000}, {-3478000000.000000, 6.000000}, {-3468000000.000000, 7.000000}, {-3438000000.000000, 8.000000}, {-3418000000.000000, 9.000000}, {-3398000000.000000, 10.000000}},
					},
					{
						Class:  "sample",
						Labels: map[string]string{"other": "test", "label": "41"},
						Attrs:  map[string]string{},
						Values: [][]float64{{-3600000000.000000, 1.000000}, {-3560000000.000000, 2.000000}, {-3518000000.000000, 1.000000}, {-3490000000.000000, 0.000000}, {-3471000000.000000, 1.000000}, {-3441000000.000000, 2.000000}, {-3408000000.000000, 0.000000}, {-3395000000.000000, 1.000000}},
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
						Values: [][]float64{{-3600000000.000000, 1.000000}, {-3565000000.000000, -1.000000}, {-3540000000.000000, 2.000000}, {-3528000000.000000, -2.000000}, {-3512000000.000000, 4.000000}, {-3488000000.000000, 8.000000}, {-3478000000.000000, 12.000000}, {-3468000000.000000, 16.000000}, {-3438000000.000000, 18.000000}, {-3418000000.000000, -20.000000}, {-3398000000.000000, 100.000000}},
					},
					{
						Class:  "sample",
						Labels: map[string]string{"other": "test", "label": "41"},
						Attrs:  map[string]string{},
						Values: [][]float64{{-3600000000.000000, 1.000000}, {-3565000000.000000, 1.000000}, {-3540000000.000000, 2.000000}, {-3528000000.000000, 3.000000}, {-3512000000.000000, 4.000000}, {-3488000000.000000, 5.000000}, {-3478000000.000000, 6.000000}, {-3468000000.000000, 7.000000}, {-3438000000.000000, 8.000000}, {-3418000000.000000, 9.000000}, {-3398000000.000000, 10.000000}},
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
						Values: [][]float64{{-3600000000.000000, 1.000000}, {-3565000000.000000, -1.000000}, {-3540000000.000000, 2.000000}, {-3528000000.000000, -2.000000}, {-3512000000.000000, 4.000000}, {-3488000000.000000, 8.000000}, {-3478000000.000000, 12.000000}, {-3468000000.000000, 16.000000}, {-3438000000.000000, 18.000000}, {-3418000000.000000, -20.000000}, {-3398000000.000000, 100.000000}},
					},
					{
						Class:  "sample",
						Labels: map[string]string{"label": "41", "other": "test"},
						Attrs:  map[string]string{},
						Values: [][]float64{{-3600000000.000000, 1.000000}, {-3560000000.000000, 2.000000}, {-3518000000.000000, 1.000000}, {-3490000000.000000, 0.000000}, {-3471000000.000000, 1.000000}, {-3441000000.000000, 2.000000}, {-3408000000.000000, 0.000000}, {-3395000000.000000, 1.000000}},
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
						Values: [][]float64{{-3600000000.000000, 1.000000}, {-3565000000.000000, -1.000000}, {-3540000000.000000, 2.000000}, {-3528000000.000000, -2.000000}, {-3512000000.000000, 4.000000}, {-3488000000.000000, 8.000000}, {-3478000000.000000, 12.000000}, {-3468000000.000000, 16.000000}, {-3438000000.000000, 18.000000}, {-3418000000.000000, -20.000000}, {-3398000000.000000, 100.000000}},
					},
				},

				SeriesTests: seriesEqualityTestMap,
			},
		},
	},
}
