package prototests

import (
	"testing"

	"github.com/ovh/erlenmeyer/proto/graphite"
)

//
// @code auto-generated
// Test can be started with
// go test proto/prototests/graphite_averageSeriesWithWildcards_test.go proto/prototests/exec_test.go -v
//

// TestGraphiteAverageSeriesWithWildcards process Invert Graphite Unit tests
func TestGraphiteAverageSeriesWithWildcards(t *testing.T) {
	RunTest(t, graphiteGraphiteAverageSeriesWithWildcards, "")
}

var graphiteGraphiteAverageSeriesWithWildcards = []unitTests{

	{
		Plan: []graphite.Function{
			graphite.Function{
				Name:       "averageSeriesWithWildcards",
				Arguments:  []string{"SWAP", "2"},
				Parameters: make(map[string]string),
			},
		},
		Samples: []OperatorGTSTest{
			OperatorGTSTest{
				SamplePrefix: BucketizeGtsList(BuildGtsList([]FloatGeoTimeSeries{emptyGTS, gts1, gts1})),
				GTSResult: []FloatGeoTimeSeries{
					FloatGeoTimeSeries{
						Class:  "",
						Labels: map[string]string{"label": "42"},
						Attrs:  map[string]string{},
						Values: [][]float64{[]float64{25000000.000000, 1.000000}, []float64{85000000.000000, -0.333333}, []float64{145000000.000000, 10.000000}, []float64{205000000.000000, 32.666667}},
					},
				},

				SeriesTests: seriesEqualityTestMap,
			},
			OperatorGTSTest{
				SamplePrefix: BucketizeGtsList(BuildGtsList([]FloatGeoTimeSeries{emptyGTS, gts1})),
				GTSResult: []FloatGeoTimeSeries{
					FloatGeoTimeSeries{
						Class:  "",
						Labels: map[string]string{"label": "42"},
						Attrs:  map[string]string{},
						Values: [][]float64{[]float64{25000000.000000, 1.000000}, []float64{85000000.000000, -0.333333}, []float64{145000000.000000, 10.000000}, []float64{205000000.000000, 32.666667}},
					},
				},

				SeriesTests: seriesEqualityTestMap,
			},
			OperatorGTSTest{
				SamplePrefix: BucketizeGtsList(BuildGtsList([]FloatGeoTimeSeries{gts1, gts1})),
				GTSResult: []FloatGeoTimeSeries{
					FloatGeoTimeSeries{
						Class:  "sample",
						Labels: map[string]string{"label": "42"},
						Attrs:  map[string]string{},
						Values: [][]float64{[]float64{25000000.000000, 1.000000}, []float64{85000000.000000, -0.333333}, []float64{145000000.000000, 10.000000}, []float64{205000000.000000, 32.666667}},
					},
				},

				SeriesTests: seriesEqualityTestMap,
			},
			OperatorGTSTest{
				SamplePrefix: BucketizeGtsList(BuildGtsList([]FloatGeoTimeSeries{gts1, gts2, gts3})),
				GTSResult: []FloatGeoTimeSeries{
					FloatGeoTimeSeries{
						Class:  "sample",
						Labels: map[string]string{},
						Attrs:  map[string]string{},
						Values: [][]float64{[]float64{25000000.000000, 1.000000}, []float64{85000000.000000, 1.055556}, []float64{145000000.000000, 5.333333}, []float64{205000000.000000, 14.222222}},
					},
				},

				SeriesTests: seriesEqualityTestMap,
			},
			OperatorGTSTest{
				SamplePrefix: BucketizeGtsList(BuildGtsList([]FloatGeoTimeSeries{gts1, gts2})),
				GTSResult: []FloatGeoTimeSeries{
					FloatGeoTimeSeries{
						Class:  "sample",
						Labels: map[string]string{},
						Attrs:  map[string]string{},
						Values: [][]float64{[]float64{25000000.000000, 1.000000}, []float64{85000000.000000, 0.833333}, []float64{145000000.000000, 7.750000}, []float64{205000000.000000, 20.833333}},
					},
				},

				SeriesTests: seriesEqualityTestMap,
			},
			OperatorGTSTest{
				SamplePrefix: BucketizeGtsList(BuildGtsList([]FloatGeoTimeSeries{gts1, gts3})),
				GTSResult: []FloatGeoTimeSeries{
					FloatGeoTimeSeries{
						Class:  "sample",
						Labels: map[string]string{},
						Attrs:  map[string]string{},
						Values: [][]float64{[]float64{25000000.000000, 1.000000}, []float64{85000000.000000, 0.583333}, []float64{145000000.000000, 5.250000}, []float64{205000000.000000, 16.833333}},
					},
				},

				SeriesTests: seriesEqualityTestMap,
			},
			OperatorGTSTest{
				SamplePrefix: BucketizeGtsList(BuildGtsList([]FloatGeoTimeSeries{gts1})),
				GTSResult: []FloatGeoTimeSeries{
					FloatGeoTimeSeries{
						Class:  "sample",
						Labels: map[string]string{"label": "42"},
						Attrs:  map[string]string{},
						Values: [][]float64{[]float64{25000000.000000, 1.000000}, []float64{85000000.000000, -0.333333}, []float64{145000000.000000, 10.000000}, []float64{205000000.000000, 32.666667}},
					},
				},

				SeriesTests: seriesEqualityTestMap,
			},
			OperatorGTSTest{
				SamplePrefix: BuildGtsList([]FloatGeoTimeSeries{emptyGTS, gts1, gts1}),
				GTSResult: []FloatGeoTimeSeries{
					FloatGeoTimeSeries{
						Class:  "",
						Labels: map[string]string{"label": "42"},
						Attrs:  map[string]string{},
						Values: [][]float64{[]float64{0.000000, 1.000000}, []float64{35000000.000000, -1.000000}, []float64{60000000.000000, 2.000000}, []float64{72000000.000000, -2.000000}, []float64{88000000.000000, 4.000000}, []float64{112000000.000000, 8.000000}, []float64{122000000.000000, 12.000000}, []float64{132000000.000000, 16.000000}, []float64{162000000.000000, 18.000000}, []float64{182000000.000000, -20.000000}, []float64{202000000.000000, 100.000000}},
					},
				},

				SeriesTests: seriesEqualityTestMap,
			},
			OperatorGTSTest{
				SamplePrefix: BuildGtsList([]FloatGeoTimeSeries{emptyGTS, gts1}),
				GTSResult: []FloatGeoTimeSeries{
					FloatGeoTimeSeries{
						Class:  "",
						Labels: map[string]string{"label": "42"},
						Attrs:  map[string]string{},
						Values: [][]float64{[]float64{0.000000, 1.000000}, []float64{35000000.000000, -1.000000}, []float64{60000000.000000, 2.000000}, []float64{72000000.000000, -2.000000}, []float64{88000000.000000, 4.000000}, []float64{112000000.000000, 8.000000}, []float64{122000000.000000, 12.000000}, []float64{132000000.000000, 16.000000}, []float64{162000000.000000, 18.000000}, []float64{182000000.000000, -20.000000}, []float64{202000000.000000, 100.000000}},
					},
				},

				SeriesTests: seriesEqualityTestMap,
			},
			OperatorGTSTest{
				SamplePrefix: BuildGtsList([]FloatGeoTimeSeries{gts1, gts1}),
				GTSResult: []FloatGeoTimeSeries{
					FloatGeoTimeSeries{
						Class:  "sample",
						Labels: map[string]string{"label": "42"},
						Attrs:  map[string]string{},
						Values: [][]float64{[]float64{0.000000, 1.000000}, []float64{35000000.000000, -1.000000}, []float64{60000000.000000, 2.000000}, []float64{72000000.000000, -2.000000}, []float64{88000000.000000, 4.000000}, []float64{112000000.000000, 8.000000}, []float64{122000000.000000, 12.000000}, []float64{132000000.000000, 16.000000}, []float64{162000000.000000, 18.000000}, []float64{182000000.000000, -20.000000}, []float64{202000000.000000, 100.000000}},
					},
				},

				SeriesTests: seriesEqualityTestMap,
			},
			OperatorGTSTest{
				SamplePrefix: BuildGtsList([]FloatGeoTimeSeries{gts1, gts2, gts3}),
				GTSResult: []FloatGeoTimeSeries{
					FloatGeoTimeSeries{
						Class:  "sample",
						Labels: map[string]string{},
						Attrs:  map[string]string{},
						Values: [][]float64{[]float64{0.000000, 1.000000}, []float64{35000000.000000, 0.000000}, []float64{40000000.000000, 2.000000}, []float64{60000000.000000, 2.000000}, []float64{72000000.000000, 0.500000}, []float64{82000000.000000, 1.000000}, []float64{88000000.000000, 4.000000}, []float64{110000000.000000, 0.000000}, []float64{112000000.000000, 6.500000}, []float64{122000000.000000, 9.000000}, []float64{129000000.000000, 1.000000}, []float64{132000000.000000, 11.500000}, []float64{159000000.000000, 2.000000}, []float64{162000000.000000, 13.000000}, []float64{182000000.000000, -5.500000}, []float64{192000000.000000, 0.000000}, []float64{202000000.000000, 55.000000}, []float64{205000000.000000, 1.000000}},
					},
				},

				SeriesTests: seriesEqualityTestMap,
			},
			OperatorGTSTest{
				SamplePrefix: BuildGtsList([]FloatGeoTimeSeries{gts1, gts2}),
				GTSResult: []FloatGeoTimeSeries{
					FloatGeoTimeSeries{
						Class:  "sample",
						Labels: map[string]string{},
						Attrs:  map[string]string{},
						Values: [][]float64{[]float64{0.000000, 1.000000}, []float64{35000000.000000, 0.000000}, []float64{60000000.000000, 2.000000}, []float64{72000000.000000, 0.500000}, []float64{88000000.000000, 4.000000}, []float64{112000000.000000, 6.500000}, []float64{122000000.000000, 9.000000}, []float64{132000000.000000, 11.500000}, []float64{162000000.000000, 13.000000}, []float64{182000000.000000, -5.500000}, []float64{202000000.000000, 55.000000}},
					},
				},

				SeriesTests: seriesEqualityTestMap,
			},
			OperatorGTSTest{
				SamplePrefix: BuildGtsList([]FloatGeoTimeSeries{gts1, gts3}),
				GTSResult: []FloatGeoTimeSeries{
					FloatGeoTimeSeries{
						Class:  "sample",
						Labels: map[string]string{},
						Attrs:  map[string]string{},
						Values: [][]float64{[]float64{0.000000, 1.000000}, []float64{35000000.000000, -1.000000}, []float64{40000000.000000, 2.000000}, []float64{60000000.000000, 2.000000}, []float64{72000000.000000, -2.000000}, []float64{82000000.000000, 1.000000}, []float64{88000000.000000, 4.000000}, []float64{110000000.000000, 0.000000}, []float64{112000000.000000, 8.000000}, []float64{122000000.000000, 12.000000}, []float64{129000000.000000, 1.000000}, []float64{132000000.000000, 16.000000}, []float64{159000000.000000, 2.000000}, []float64{162000000.000000, 18.000000}, []float64{182000000.000000, -20.000000}, []float64{192000000.000000, 0.000000}, []float64{202000000.000000, 100.000000}, []float64{205000000.000000, 1.000000}},
					},
				},

				SeriesTests: seriesEqualityTestMap,
			},
			OperatorGTSTest{
				SamplePrefix: BuildGtsList([]FloatGeoTimeSeries{gts1}),
				GTSResult: []FloatGeoTimeSeries{
					FloatGeoTimeSeries{
						Class:  "sample",
						Labels: map[string]string{"label": "42"},
						Attrs:  map[string]string{},
						Values: [][]float64{[]float64{0.000000, 1.000000}, []float64{35000000.000000, -1.000000}, []float64{60000000.000000, 2.000000}, []float64{72000000.000000, -2.000000}, []float64{88000000.000000, 4.000000}, []float64{112000000.000000, 8.000000}, []float64{122000000.000000, 12.000000}, []float64{132000000.000000, 16.000000}, []float64{162000000.000000, 18.000000}, []float64{182000000.000000, -20.000000}, []float64{202000000.000000, 100.000000}},
					},
				},

				SeriesTests: seriesEqualityTestMap,
			},
		},
	},
}
