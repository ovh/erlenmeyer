package prototests

import (
	"testing"

	"github.com/ovh/erlenmeyer/proto/graphite"
)

//
// @code auto-generated
// Test can be started with
// go test proto/prototests/graphite_aggregateWithWildcards_test.go proto/prototests/exec_test.go -v
//

// TestGraphiteAggregateWithWildcards process Invert Graphite Unit tests
func TestGraphiteAggregateWithWildcards(t *testing.T) {
	RunTest(t, graphiteGraphiteAggregateWithWildcards, "")
}

var graphiteGraphiteAggregateWithWildcards = []unitTests{

	{
		Plan: []graphite.Function{
			graphite.Function{
				Name:       "aggregateWithWildcards",
				Arguments:  []string{"SWAP", "average", "2"},
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
	{
		Plan: []graphite.Function{
			graphite.Function{
				Name:       "aggregateWithWildcards",
				Arguments:  []string{"SWAP", "median", "2"},
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
						Values: [][]float64{[]float64{25000000.000000, 1.000000}, []float64{85000000.000000, 1.500000}, []float64{145000000.000000, 5.500000}, []float64{205000000.000000, 9.000000}},
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
		},
	},
	{
		Plan: []graphite.Function{
			graphite.Function{
				Name:       "aggregateWithWildcards",
				Arguments:  []string{"SWAP", "sum", "2"},
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
						Values: [][]float64{[]float64{25000000.000000, 2.000000}, []float64{85000000.000000, -0.666667}, []float64{145000000.000000, 20.000000}, []float64{205000000.000000, 65.333333}},
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
						Values: [][]float64{[]float64{25000000.000000, 2.000000}, []float64{85000000.000000, -0.666667}, []float64{145000000.000000, 20.000000}, []float64{205000000.000000, 65.333333}},
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
						Values: [][]float64{[]float64{25000000.000000, 3.000000}, []float64{85000000.000000, 3.166667}, []float64{145000000.000000, 16.000000}, []float64{205000000.000000, 42.666667}},
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
						Values: [][]float64{[]float64{25000000.000000, 2.000000}, []float64{85000000.000000, 1.666667}, []float64{145000000.000000, 15.500000}, []float64{205000000.000000, 41.666667}},
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
						Values: [][]float64{[]float64{25000000.000000, 2.000000}, []float64{85000000.000000, 1.166667}, []float64{145000000.000000, 10.500000}, []float64{205000000.000000, 33.666667}},
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
						Values: [][]float64{[]float64{0.000000, 2.000000}, []float64{35000000.000000, -2.000000}, []float64{60000000.000000, 4.000000}, []float64{72000000.000000, -4.000000}, []float64{88000000.000000, 8.000000}, []float64{112000000.000000, 16.000000}, []float64{122000000.000000, 24.000000}, []float64{132000000.000000, 32.000000}, []float64{162000000.000000, 36.000000}, []float64{182000000.000000, -40.000000}, []float64{202000000.000000, 200.000000}},
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
						Values: [][]float64{[]float64{0.000000, 2.000000}, []float64{35000000.000000, -2.000000}, []float64{60000000.000000, 4.000000}, []float64{72000000.000000, -4.000000}, []float64{88000000.000000, 8.000000}, []float64{112000000.000000, 16.000000}, []float64{122000000.000000, 24.000000}, []float64{132000000.000000, 32.000000}, []float64{162000000.000000, 36.000000}, []float64{182000000.000000, -40.000000}, []float64{202000000.000000, 200.000000}},
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
						Values: [][]float64{[]float64{0.000000, 3.000000}, []float64{35000000.000000, 0.000000}, []float64{40000000.000000, 2.000000}, []float64{60000000.000000, 4.000000}, []float64{72000000.000000, 1.000000}, []float64{82000000.000000, 1.000000}, []float64{88000000.000000, 8.000000}, []float64{110000000.000000, 0.000000}, []float64{112000000.000000, 13.000000}, []float64{122000000.000000, 18.000000}, []float64{129000000.000000, 1.000000}, []float64{132000000.000000, 23.000000}, []float64{159000000.000000, 2.000000}, []float64{162000000.000000, 26.000000}, []float64{182000000.000000, -11.000000}, []float64{192000000.000000, 0.000000}, []float64{202000000.000000, 110.000000}, []float64{205000000.000000, 1.000000}},
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
						Values: [][]float64{[]float64{0.000000, 2.000000}, []float64{35000000.000000, 0.000000}, []float64{60000000.000000, 4.000000}, []float64{72000000.000000, 1.000000}, []float64{88000000.000000, 8.000000}, []float64{112000000.000000, 13.000000}, []float64{122000000.000000, 18.000000}, []float64{132000000.000000, 23.000000}, []float64{162000000.000000, 26.000000}, []float64{182000000.000000, -11.000000}, []float64{202000000.000000, 110.000000}},
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
						Values: [][]float64{[]float64{0.000000, 2.000000}, []float64{35000000.000000, -1.000000}, []float64{40000000.000000, 2.000000}, []float64{60000000.000000, 2.000000}, []float64{72000000.000000, -2.000000}, []float64{82000000.000000, 1.000000}, []float64{88000000.000000, 4.000000}, []float64{110000000.000000, 0.000000}, []float64{112000000.000000, 8.000000}, []float64{122000000.000000, 12.000000}, []float64{129000000.000000, 1.000000}, []float64{132000000.000000, 16.000000}, []float64{159000000.000000, 2.000000}, []float64{162000000.000000, 18.000000}, []float64{182000000.000000, -20.000000}, []float64{192000000.000000, 0.000000}, []float64{202000000.000000, 100.000000}, []float64{205000000.000000, 1.000000}},
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
	{
		Plan: []graphite.Function{
			graphite.Function{
				Name:       "aggregateWithWildcards",
				Arguments:  []string{"SWAP", "min", "2"},
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
						Values: [][]float64{[]float64{25000000.000000, 1.000000}, []float64{85000000.000000, -0.333333}, []float64{145000000.000000, 0.500000}, []float64{205000000.000000, 1.000000}},
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
						Values: [][]float64{[]float64{25000000.000000, 1.000000}, []float64{85000000.000000, -0.333333}, []float64{145000000.000000, 5.500000}, []float64{205000000.000000, 9.000000}},
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
						Values: [][]float64{[]float64{25000000.000000, 1.000000}, []float64{85000000.000000, -0.333333}, []float64{145000000.000000, 0.500000}, []float64{205000000.000000, 1.000000}},
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
						Values: [][]float64{[]float64{0.000000, 1.000000}, []float64{35000000.000000, -1.000000}, []float64{40000000.000000, 2.000000}, []float64{60000000.000000, 2.000000}, []float64{72000000.000000, -2.000000}, []float64{82000000.000000, 1.000000}, []float64{88000000.000000, 4.000000}, []float64{110000000.000000, 0.000000}, []float64{112000000.000000, 5.000000}, []float64{122000000.000000, 6.000000}, []float64{129000000.000000, 1.000000}, []float64{132000000.000000, 7.000000}, []float64{159000000.000000, 2.000000}, []float64{162000000.000000, 8.000000}, []float64{182000000.000000, -20.000000}, []float64{192000000.000000, 0.000000}, []float64{202000000.000000, 10.000000}, []float64{205000000.000000, 1.000000}},
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
						Values: [][]float64{[]float64{0.000000, 1.000000}, []float64{35000000.000000, -1.000000}, []float64{60000000.000000, 2.000000}, []float64{72000000.000000, -2.000000}, []float64{88000000.000000, 4.000000}, []float64{112000000.000000, 5.000000}, []float64{122000000.000000, 6.000000}, []float64{132000000.000000, 7.000000}, []float64{162000000.000000, 8.000000}, []float64{182000000.000000, -20.000000}, []float64{202000000.000000, 10.000000}},
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
	{
		Plan: []graphite.Function{
			graphite.Function{
				Name:       "aggregateWithWildcards",
				Arguments:  []string{"SWAP", "max", "2"},
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
						Values: [][]float64{[]float64{25000000.000000, 1.000000}, []float64{85000000.000000, 2.000000}, []float64{145000000.000000, 10.000000}, []float64{205000000.000000, 32.666667}},
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
						Values: [][]float64{[]float64{25000000.000000, 1.000000}, []float64{85000000.000000, 2.000000}, []float64{145000000.000000, 10.000000}, []float64{205000000.000000, 32.666667}},
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
						Values: [][]float64{[]float64{25000000.000000, 1.000000}, []float64{85000000.000000, 1.500000}, []float64{145000000.000000, 10.000000}, []float64{205000000.000000, 32.666667}},
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
						Values: [][]float64{[]float64{0.000000, 1.000000}, []float64{35000000.000000, 1.000000}, []float64{40000000.000000, 2.000000}, []float64{60000000.000000, 2.000000}, []float64{72000000.000000, 3.000000}, []float64{82000000.000000, 1.000000}, []float64{88000000.000000, 4.000000}, []float64{110000000.000000, 0.000000}, []float64{112000000.000000, 8.000000}, []float64{122000000.000000, 12.000000}, []float64{129000000.000000, 1.000000}, []float64{132000000.000000, 16.000000}, []float64{159000000.000000, 2.000000}, []float64{162000000.000000, 18.000000}, []float64{182000000.000000, 9.000000}, []float64{192000000.000000, 0.000000}, []float64{202000000.000000, 100.000000}, []float64{205000000.000000, 1.000000}},
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
						Values: [][]float64{[]float64{0.000000, 1.000000}, []float64{35000000.000000, 1.000000}, []float64{60000000.000000, 2.000000}, []float64{72000000.000000, 3.000000}, []float64{88000000.000000, 4.000000}, []float64{112000000.000000, 8.000000}, []float64{122000000.000000, 12.000000}, []float64{132000000.000000, 16.000000}, []float64{162000000.000000, 18.000000}, []float64{182000000.000000, 9.000000}, []float64{202000000.000000, 100.000000}},
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
	{
		Plan: []graphite.Function{
			graphite.Function{
				Name:       "aggregateWithWildcards",
				Arguments:  []string{"SWAP", "diff", "2"},
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
						Values: [][]float64{[]float64{25000000.000000, -2.000000}, []float64{85000000.000000, 0.666667}, []float64{145000000.000000, -20.000000}, []float64{205000000.000000, -65.333333}},
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
						Values: [][]float64{[]float64{25000000.000000, -1.000000}, []float64{85000000.000000, 0.333333}, []float64{145000000.000000, -10.000000}, []float64{205000000.000000, -32.666667}},
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
						Values: [][]float64{[]float64{25000000.000000, -2.000000}, []float64{85000000.000000, 0.666667}, []float64{145000000.000000, -20.000000}, []float64{205000000.000000, -65.333333}},
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
						Values: [][]float64{[]float64{25000000.000000, -3.000000}, []float64{85000000.000000, -3.166667}, []float64{145000000.000000, -16.000000}, []float64{205000000.000000, -42.666667}},
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
						Values: [][]float64{[]float64{25000000.000000, -2.000000}, []float64{85000000.000000, -1.666667}, []float64{145000000.000000, -15.500000}, []float64{205000000.000000, -41.666667}},
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
						Values: [][]float64{[]float64{25000000.000000, -2.000000}, []float64{85000000.000000, -1.166667}, []float64{145000000.000000, -10.500000}, []float64{205000000.000000, -33.666667}},
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
						Values: [][]float64{[]float64{25000000.000000, -1.000000}, []float64{85000000.000000, 0.333333}, []float64{145000000.000000, -10.000000}, []float64{205000000.000000, -32.666667}},
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
						Values: [][]float64{[]float64{0.000000, -2.000000}, []float64{35000000.000000, 2.000000}, []float64{60000000.000000, -4.000000}, []float64{72000000.000000, 4.000000}, []float64{88000000.000000, -8.000000}, []float64{112000000.000000, -16.000000}, []float64{122000000.000000, -24.000000}, []float64{132000000.000000, -32.000000}, []float64{162000000.000000, -36.000000}, []float64{182000000.000000, 40.000000}, []float64{202000000.000000, -200.000000}},
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
						Values: [][]float64{[]float64{0.000000, -1.000000}, []float64{35000000.000000, 1.000000}, []float64{60000000.000000, -2.000000}, []float64{72000000.000000, 2.000000}, []float64{88000000.000000, -4.000000}, []float64{112000000.000000, -8.000000}, []float64{122000000.000000, -12.000000}, []float64{132000000.000000, -16.000000}, []float64{162000000.000000, -18.000000}, []float64{182000000.000000, 20.000000}, []float64{202000000.000000, -100.000000}},
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
						Values: [][]float64{[]float64{0.000000, -2.000000}, []float64{35000000.000000, 2.000000}, []float64{60000000.000000, -4.000000}, []float64{72000000.000000, 4.000000}, []float64{88000000.000000, -8.000000}, []float64{112000000.000000, -16.000000}, []float64{122000000.000000, -24.000000}, []float64{132000000.000000, -32.000000}, []float64{162000000.000000, -36.000000}, []float64{182000000.000000, 40.000000}, []float64{202000000.000000, -200.000000}},
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
						Values: [][]float64{[]float64{0.000000, -3.000000}, []float64{35000000.000000, -0.000000}, []float64{40000000.000000, -2.000000}, []float64{60000000.000000, -4.000000}, []float64{72000000.000000, -1.000000}, []float64{82000000.000000, -1.000000}, []float64{88000000.000000, -8.000000}, []float64{110000000.000000, -0.000000}, []float64{112000000.000000, -13.000000}, []float64{122000000.000000, -18.000000}, []float64{129000000.000000, -1.000000}, []float64{132000000.000000, -23.000000}, []float64{159000000.000000, -2.000000}, []float64{162000000.000000, -26.000000}, []float64{182000000.000000, 11.000000}, []float64{192000000.000000, -0.000000}, []float64{202000000.000000, -110.000000}, []float64{205000000.000000, -1.000000}},
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
						Values: [][]float64{[]float64{0.000000, -2.000000}, []float64{35000000.000000, -0.000000}, []float64{60000000.000000, -4.000000}, []float64{72000000.000000, -1.000000}, []float64{88000000.000000, -8.000000}, []float64{112000000.000000, -13.000000}, []float64{122000000.000000, -18.000000}, []float64{132000000.000000, -23.000000}, []float64{162000000.000000, -26.000000}, []float64{182000000.000000, 11.000000}, []float64{202000000.000000, -110.000000}},
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
						Values: [][]float64{[]float64{0.000000, -2.000000}, []float64{35000000.000000, 1.000000}, []float64{40000000.000000, -2.000000}, []float64{60000000.000000, -2.000000}, []float64{72000000.000000, 2.000000}, []float64{82000000.000000, -1.000000}, []float64{88000000.000000, -4.000000}, []float64{110000000.000000, -0.000000}, []float64{112000000.000000, -8.000000}, []float64{122000000.000000, -12.000000}, []float64{129000000.000000, -1.000000}, []float64{132000000.000000, -16.000000}, []float64{159000000.000000, -2.000000}, []float64{162000000.000000, -18.000000}, []float64{182000000.000000, 20.000000}, []float64{192000000.000000, -0.000000}, []float64{202000000.000000, -100.000000}, []float64{205000000.000000, -1.000000}},
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
						Values: [][]float64{[]float64{0.000000, -1.000000}, []float64{35000000.000000, 1.000000}, []float64{60000000.000000, -2.000000}, []float64{72000000.000000, 2.000000}, []float64{88000000.000000, -4.000000}, []float64{112000000.000000, -8.000000}, []float64{122000000.000000, -12.000000}, []float64{132000000.000000, -16.000000}, []float64{162000000.000000, -18.000000}, []float64{182000000.000000, 20.000000}, []float64{202000000.000000, -100.000000}},
					},
				},

				SeriesTests: seriesEqualityTestMap,
			},
		},
	},
	{
		Plan: []graphite.Function{
			graphite.Function{
				Name:       "aggregateWithWildcards",
				Arguments:  []string{"SWAP", "stddev", "2"},
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
						Values: [][]float64{[]float64{25000000.000000, 0.000000}, []float64{85000000.000000, 0.000000}, []float64{145000000.000000, 0.000000}, []float64{205000000.000000, 0.000000}},
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
						Values: [][]float64{[]float64{25000000.000000, 0.000000}, []float64{85000000.000000, 0.000000}, []float64{145000000.000000, 0.000000}, []float64{205000000.000000, 0.000000}},
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
						Values: [][]float64{[]float64{25000000.000000, 0.000000}, []float64{85000000.000000, 0.000000}, []float64{145000000.000000, 0.000000}, []float64{205000000.000000, 0.000000}},
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
						Values: [][]float64{[]float64{25000000.000000, 0.000000}, []float64{85000000.000000, 1.228519}, []float64{145000000.000000, 4.752192}, []float64{205000000.000000, 16.466577}},
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
						Values: [][]float64{[]float64{25000000.000000, 0.000000}, []float64{85000000.000000, 1.649916}, []float64{145000000.000000, 3.181981}, []float64{205000000.000000, 16.734860}},
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
						Values: [][]float64{[]float64{25000000.000000, 0.000000}, []float64{85000000.000000, 1.296362}, []float64{145000000.000000, 6.717514}, []float64{205000000.000000, 22.391715}},
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
						Values: [][]float64{[]float64{25000000.000000, 0.000000}, []float64{85000000.000000, 0.000000}, []float64{145000000.000000, 0.000000}, []float64{205000000.000000, 0.000000}},
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
						Values: [][]float64{[]float64{0.000000, 0.000000}, []float64{35000000.000000, 0.000000}, []float64{60000000.000000, 0.000000}, []float64{72000000.000000, 0.000000}, []float64{88000000.000000, 0.000000}, []float64{112000000.000000, 0.000000}, []float64{122000000.000000, 0.000000}, []float64{132000000.000000, 0.000000}, []float64{162000000.000000, 0.000000}, []float64{182000000.000000, 0.000000}, []float64{202000000.000000, 0.000000}},
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
						Values: [][]float64{[]float64{0.000000, 0.000000}, []float64{35000000.000000, 0.000000}, []float64{60000000.000000, 0.000000}, []float64{72000000.000000, 0.000000}, []float64{88000000.000000, 0.000000}, []float64{112000000.000000, 0.000000}, []float64{122000000.000000, 0.000000}, []float64{132000000.000000, 0.000000}, []float64{162000000.000000, 0.000000}, []float64{182000000.000000, 0.000000}, []float64{202000000.000000, 0.000000}},
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
						Values: [][]float64{[]float64{0.000000, 0.000000}, []float64{35000000.000000, 0.000000}, []float64{60000000.000000, 0.000000}, []float64{72000000.000000, 0.000000}, []float64{88000000.000000, 0.000000}, []float64{112000000.000000, 0.000000}, []float64{122000000.000000, 0.000000}, []float64{132000000.000000, 0.000000}, []float64{162000000.000000, 0.000000}, []float64{182000000.000000, 0.000000}, []float64{202000000.000000, 0.000000}},
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
						Values: [][]float64{[]float64{0.000000, 0.000000}, []float64{35000000.000000, 1.414214}, []float64{40000000.000000, 0.000000}, []float64{60000000.000000, 0.000000}, []float64{72000000.000000, 3.535534}, []float64{82000000.000000, 0.000000}, []float64{88000000.000000, 0.000000}, []float64{110000000.000000, 0.000000}, []float64{112000000.000000, 2.121320}, []float64{122000000.000000, 4.242641}, []float64{129000000.000000, 0.000000}, []float64{132000000.000000, 6.363961}, []float64{159000000.000000, 0.000000}, []float64{162000000.000000, 7.071068}, []float64{182000000.000000, 20.506097}, []float64{192000000.000000, 0.000000}, []float64{202000000.000000, 63.639610}, []float64{205000000.000000, 0.000000}},
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
						Values: [][]float64{[]float64{0.000000, 0.000000}, []float64{35000000.000000, 1.414214}, []float64{60000000.000000, 0.000000}, []float64{72000000.000000, 3.535534}, []float64{88000000.000000, 0.000000}, []float64{112000000.000000, 2.121320}, []float64{122000000.000000, 4.242641}, []float64{132000000.000000, 6.363961}, []float64{162000000.000000, 7.071068}, []float64{182000000.000000, 20.506097}, []float64{202000000.000000, 63.639610}},
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
						Values: [][]float64{[]float64{0.000000, 0.000000}, []float64{35000000.000000, 0.000000}, []float64{40000000.000000, 0.000000}, []float64{60000000.000000, 0.000000}, []float64{72000000.000000, 0.000000}, []float64{82000000.000000, 0.000000}, []float64{88000000.000000, 0.000000}, []float64{110000000.000000, 0.000000}, []float64{112000000.000000, 0.000000}, []float64{122000000.000000, 0.000000}, []float64{129000000.000000, 0.000000}, []float64{132000000.000000, 0.000000}, []float64{159000000.000000, 0.000000}, []float64{162000000.000000, 0.000000}, []float64{182000000.000000, 0.000000}, []float64{192000000.000000, 0.000000}, []float64{202000000.000000, 0.000000}, []float64{205000000.000000, 0.000000}},
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
						Values: [][]float64{[]float64{0.000000, 0.000000}, []float64{35000000.000000, 0.000000}, []float64{60000000.000000, 0.000000}, []float64{72000000.000000, 0.000000}, []float64{88000000.000000, 0.000000}, []float64{112000000.000000, 0.000000}, []float64{122000000.000000, 0.000000}, []float64{132000000.000000, 0.000000}, []float64{162000000.000000, 0.000000}, []float64{182000000.000000, 0.000000}, []float64{202000000.000000, 0.000000}},
					},
				},

				SeriesTests: seriesEqualityTestMap,
			},
		},
	},
	{
		Plan: []graphite.Function{
			graphite.Function{
				Name:       "aggregateWithWildcards",
				Arguments:  []string{"SWAP", "count", "2"},
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
						Values: [][]float64{[]float64{25000000.000000, 3.000000}, []float64{85000000.000000, 3.000000}, []float64{145000000.000000, 3.000000}, []float64{205000000.000000, 3.000000}},
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
						Values: [][]float64{[]float64{25000000.000000, 2.000000}, []float64{85000000.000000, 2.000000}, []float64{145000000.000000, 2.000000}, []float64{205000000.000000, 2.000000}},
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
						Values: [][]float64{[]float64{25000000.000000, 2.000000}, []float64{85000000.000000, 2.000000}, []float64{145000000.000000, 2.000000}, []float64{205000000.000000, 2.000000}},
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
						Values: [][]float64{[]float64{25000000.000000, 3.000000}, []float64{85000000.000000, 3.000000}, []float64{145000000.000000, 3.000000}, []float64{205000000.000000, 3.000000}},
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
						Values: [][]float64{[]float64{25000000.000000, 2.000000}, []float64{85000000.000000, 2.000000}, []float64{145000000.000000, 2.000000}, []float64{205000000.000000, 2.000000}},
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
						Values: [][]float64{[]float64{25000000.000000, 2.000000}, []float64{85000000.000000, 2.000000}, []float64{145000000.000000, 2.000000}, []float64{205000000.000000, 2.000000}},
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
						Values: [][]float64{[]float64{25000000.000000, 1.000000}, []float64{85000000.000000, 1.000000}, []float64{145000000.000000, 1.000000}, []float64{205000000.000000, 1.000000}},
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
						Values: [][]float64{[]float64{0.000000, 3.000000}, []float64{35000000.000000, 3.000000}, []float64{60000000.000000, 3.000000}, []float64{72000000.000000, 3.000000}, []float64{88000000.000000, 3.000000}, []float64{112000000.000000, 3.000000}, []float64{122000000.000000, 3.000000}, []float64{132000000.000000, 3.000000}, []float64{162000000.000000, 3.000000}, []float64{182000000.000000, 3.000000}, []float64{202000000.000000, 3.000000}},
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
						Values: [][]float64{[]float64{0.000000, 2.000000}, []float64{35000000.000000, 2.000000}, []float64{60000000.000000, 2.000000}, []float64{72000000.000000, 2.000000}, []float64{88000000.000000, 2.000000}, []float64{112000000.000000, 2.000000}, []float64{122000000.000000, 2.000000}, []float64{132000000.000000, 2.000000}, []float64{162000000.000000, 2.000000}, []float64{182000000.000000, 2.000000}, []float64{202000000.000000, 2.000000}},
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
						Values: [][]float64{[]float64{0.000000, 2.000000}, []float64{35000000.000000, 2.000000}, []float64{60000000.000000, 2.000000}, []float64{72000000.000000, 2.000000}, []float64{88000000.000000, 2.000000}, []float64{112000000.000000, 2.000000}, []float64{122000000.000000, 2.000000}, []float64{132000000.000000, 2.000000}, []float64{162000000.000000, 2.000000}, []float64{182000000.000000, 2.000000}, []float64{202000000.000000, 2.000000}},
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
						Values: [][]float64{[]float64{0.000000, 3.000000}, []float64{35000000.000000, 3.000000}, []float64{40000000.000000, 3.000000}, []float64{60000000.000000, 3.000000}, []float64{72000000.000000, 3.000000}, []float64{82000000.000000, 3.000000}, []float64{88000000.000000, 3.000000}, []float64{110000000.000000, 3.000000}, []float64{112000000.000000, 3.000000}, []float64{122000000.000000, 3.000000}, []float64{129000000.000000, 3.000000}, []float64{132000000.000000, 3.000000}, []float64{159000000.000000, 3.000000}, []float64{162000000.000000, 3.000000}, []float64{182000000.000000, 3.000000}, []float64{192000000.000000, 3.000000}, []float64{202000000.000000, 3.000000}, []float64{205000000.000000, 3.000000}},
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
						Values: [][]float64{[]float64{0.000000, 2.000000}, []float64{35000000.000000, 2.000000}, []float64{60000000.000000, 2.000000}, []float64{72000000.000000, 2.000000}, []float64{88000000.000000, 2.000000}, []float64{112000000.000000, 2.000000}, []float64{122000000.000000, 2.000000}, []float64{132000000.000000, 2.000000}, []float64{162000000.000000, 2.000000}, []float64{182000000.000000, 2.000000}, []float64{202000000.000000, 2.000000}},
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
						Values: [][]float64{[]float64{0.000000, 2.000000}, []float64{35000000.000000, 2.000000}, []float64{40000000.000000, 2.000000}, []float64{60000000.000000, 2.000000}, []float64{72000000.000000, 2.000000}, []float64{82000000.000000, 2.000000}, []float64{88000000.000000, 2.000000}, []float64{110000000.000000, 2.000000}, []float64{112000000.000000, 2.000000}, []float64{122000000.000000, 2.000000}, []float64{129000000.000000, 2.000000}, []float64{132000000.000000, 2.000000}, []float64{159000000.000000, 2.000000}, []float64{162000000.000000, 2.000000}, []float64{182000000.000000, 2.000000}, []float64{192000000.000000, 2.000000}, []float64{202000000.000000, 2.000000}, []float64{205000000.000000, 2.000000}},
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
						Values: [][]float64{[]float64{0.000000, 1.000000}, []float64{35000000.000000, 1.000000}, []float64{60000000.000000, 1.000000}, []float64{72000000.000000, 1.000000}, []float64{88000000.000000, 1.000000}, []float64{112000000.000000, 1.000000}, []float64{122000000.000000, 1.000000}, []float64{132000000.000000, 1.000000}, []float64{162000000.000000, 1.000000}, []float64{182000000.000000, 1.000000}, []float64{202000000.000000, 1.000000}},
					},
				},

				SeriesTests: seriesEqualityTestMap,
			},
		},
	},
	{
		Plan: []graphite.Function{
			graphite.Function{
				Name:       "aggregateWithWildcards",
				Arguments:  []string{"SWAP", "range", "2"},
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
						Values: [][]float64{[]float64{25000000.000000, -2.000000}, []float64{85000000.000000, 0.666667}, []float64{145000000.000000, -20.000000}, []float64{205000000.000000, -65.333333}},
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
						Values: [][]float64{[]float64{25000000.000000, -2.000000}, []float64{85000000.000000, 0.666667}, []float64{145000000.000000, -20.000000}, []float64{205000000.000000, -65.333333}},
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
						Values: [][]float64{[]float64{25000000.000000, -2.000000}, []float64{85000000.000000, 0.666667}, []float64{145000000.000000, -20.000000}, []float64{205000000.000000, -65.333333}},
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
						Values: [][]float64{[]float64{25000000.000000, -2.000000}, []float64{85000000.000000, -1.666667}, []float64{145000000.000000, -10.500000}, []float64{205000000.000000, -33.666667}},
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
						Values: [][]float64{[]float64{25000000.000000, -2.000000}, []float64{85000000.000000, -1.666667}, []float64{145000000.000000, -15.500000}, []float64{205000000.000000, -41.666667}},
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
						Values: [][]float64{[]float64{25000000.000000, -2.000000}, []float64{85000000.000000, -1.166667}, []float64{145000000.000000, -10.500000}, []float64{205000000.000000, -33.666667}},
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
						Values: [][]float64{[]float64{25000000.000000, -2.000000}, []float64{85000000.000000, 0.666667}, []float64{145000000.000000, -20.000000}, []float64{205000000.000000, -65.333333}},
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
						Values: [][]float64{[]float64{0.000000, -2.000000}, []float64{35000000.000000, 2.000000}, []float64{60000000.000000, -4.000000}, []float64{72000000.000000, 4.000000}, []float64{88000000.000000, -8.000000}, []float64{112000000.000000, -16.000000}, []float64{122000000.000000, -24.000000}, []float64{132000000.000000, -32.000000}, []float64{162000000.000000, -36.000000}, []float64{182000000.000000, 40.000000}, []float64{202000000.000000, -200.000000}},
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
						Values: [][]float64{[]float64{0.000000, -2.000000}, []float64{35000000.000000, 2.000000}, []float64{60000000.000000, -4.000000}, []float64{72000000.000000, 4.000000}, []float64{88000000.000000, -8.000000}, []float64{112000000.000000, -16.000000}, []float64{122000000.000000, -24.000000}, []float64{132000000.000000, -32.000000}, []float64{162000000.000000, -36.000000}, []float64{182000000.000000, 40.000000}, []float64{202000000.000000, -200.000000}},
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
						Values: [][]float64{[]float64{0.000000, -2.000000}, []float64{35000000.000000, 2.000000}, []float64{60000000.000000, -4.000000}, []float64{72000000.000000, 4.000000}, []float64{88000000.000000, -8.000000}, []float64{112000000.000000, -16.000000}, []float64{122000000.000000, -24.000000}, []float64{132000000.000000, -32.000000}, []float64{162000000.000000, -36.000000}, []float64{182000000.000000, 40.000000}, []float64{202000000.000000, -200.000000}},
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
						Values: [][]float64{[]float64{0.000000, -2.000000}, []float64{35000000.000000, 0.000000}, []float64{40000000.000000, -4.000000}, []float64{60000000.000000, -4.000000}, []float64{72000000.000000, -1.000000}, []float64{82000000.000000, -2.000000}, []float64{88000000.000000, -8.000000}, []float64{110000000.000000, -0.000000}, []float64{112000000.000000, -13.000000}, []float64{122000000.000000, -18.000000}, []float64{129000000.000000, -2.000000}, []float64{132000000.000000, -23.000000}, []float64{159000000.000000, -4.000000}, []float64{162000000.000000, -26.000000}, []float64{182000000.000000, 11.000000}, []float64{192000000.000000, -0.000000}, []float64{202000000.000000, -110.000000}, []float64{205000000.000000, -2.000000}},
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
						Values: [][]float64{[]float64{0.000000, -2.000000}, []float64{35000000.000000, 0.000000}, []float64{60000000.000000, -4.000000}, []float64{72000000.000000, -1.000000}, []float64{88000000.000000, -8.000000}, []float64{112000000.000000, -13.000000}, []float64{122000000.000000, -18.000000}, []float64{132000000.000000, -23.000000}, []float64{162000000.000000, -26.000000}, []float64{182000000.000000, 11.000000}, []float64{202000000.000000, -110.000000}},
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
						Values: [][]float64{[]float64{0.000000, -2.000000}, []float64{35000000.000000, 2.000000}, []float64{40000000.000000, -4.000000}, []float64{60000000.000000, -4.000000}, []float64{72000000.000000, 4.000000}, []float64{82000000.000000, -2.000000}, []float64{88000000.000000, -8.000000}, []float64{110000000.000000, -0.000000}, []float64{112000000.000000, -16.000000}, []float64{122000000.000000, -24.000000}, []float64{129000000.000000, -2.000000}, []float64{132000000.000000, -32.000000}, []float64{159000000.000000, -4.000000}, []float64{162000000.000000, -36.000000}, []float64{182000000.000000, 40.000000}, []float64{192000000.000000, -0.000000}, []float64{202000000.000000, -200.000000}, []float64{205000000.000000, -2.000000}},
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
						Values: [][]float64{[]float64{0.000000, -2.000000}, []float64{35000000.000000, 2.000000}, []float64{60000000.000000, -4.000000}, []float64{72000000.000000, 4.000000}, []float64{88000000.000000, -8.000000}, []float64{112000000.000000, -16.000000}, []float64{122000000.000000, -24.000000}, []float64{132000000.000000, -32.000000}, []float64{162000000.000000, -36.000000}, []float64{182000000.000000, 40.000000}, []float64{202000000.000000, -200.000000}},
					},
				},

				SeriesTests: seriesEqualityTestMap,
			},
		},
	},
	{
		Plan: []graphite.Function{
			graphite.Function{
				Name:       "aggregateWithWildcards",
				Arguments:  []string{"SWAP", "multiply", "2"},
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
						Values: [][]float64{[]float64{25000000.000000, 1.000000}, []float64{85000000.000000, 0.111111}, []float64{145000000.000000, 100.000000}, []float64{205000000.000000, 1067.111111}},
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
						Values: [][]float64{[]float64{25000000.000000, 1.000000}, []float64{85000000.000000, 0.111111}, []float64{145000000.000000, 100.000000}, []float64{205000000.000000, 1067.111111}},
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
						Values: [][]float64{[]float64{25000000.000000, 1.000000}, []float64{85000000.000000, -1.000000}, []float64{145000000.000000, 27.500000}, []float64{205000000.000000, 294.000000}},
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
						Values: [][]float64{[]float64{25000000.000000, 1.000000}, []float64{85000000.000000, -0.666667}, []float64{145000000.000000, 55.000000}, []float64{205000000.000000, 294.000000}},
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
						Values: [][]float64{[]float64{25000000.000000, 1.000000}, []float64{85000000.000000, -0.500000}, []float64{145000000.000000, 5.000000}, []float64{205000000.000000, 32.666667}},
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
						Values: [][]float64{[]float64{0.000000, 1.000000}, []float64{35000000.000000, 1.000000}, []float64{60000000.000000, 4.000000}, []float64{72000000.000000, 4.000000}, []float64{88000000.000000, 16.000000}, []float64{112000000.000000, 64.000000}, []float64{122000000.000000, 144.000000}, []float64{132000000.000000, 256.000000}, []float64{162000000.000000, 324.000000}, []float64{182000000.000000, 400.000000}, []float64{202000000.000000, 10000.000000}},
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
						Values: [][]float64{[]float64{0.000000, 1.000000}, []float64{35000000.000000, 1.000000}, []float64{60000000.000000, 4.000000}, []float64{72000000.000000, 4.000000}, []float64{88000000.000000, 16.000000}, []float64{112000000.000000, 64.000000}, []float64{122000000.000000, 144.000000}, []float64{132000000.000000, 256.000000}, []float64{162000000.000000, 324.000000}, []float64{182000000.000000, 400.000000}, []float64{202000000.000000, 10000.000000}},
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
						Values: [][]float64{[]float64{0.000000, 1.000000}, []float64{35000000.000000, -1.000000}, []float64{40000000.000000, 2.000000}, []float64{60000000.000000, 4.000000}, []float64{72000000.000000, -6.000000}, []float64{82000000.000000, 1.000000}, []float64{88000000.000000, 16.000000}, []float64{110000000.000000, 0.000000}, []float64{112000000.000000, 40.000000}, []float64{122000000.000000, 72.000000}, []float64{129000000.000000, 1.000000}, []float64{132000000.000000, 112.000000}, []float64{159000000.000000, 2.000000}, []float64{162000000.000000, 144.000000}, []float64{182000000.000000, -180.000000}, []float64{192000000.000000, 0.000000}, []float64{202000000.000000, 1000.000000}, []float64{205000000.000000, 1.000000}},
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
						Values: [][]float64{[]float64{0.000000, 1.000000}, []float64{35000000.000000, -1.000000}, []float64{60000000.000000, 4.000000}, []float64{72000000.000000, -6.000000}, []float64{88000000.000000, 16.000000}, []float64{112000000.000000, 40.000000}, []float64{122000000.000000, 72.000000}, []float64{132000000.000000, 112.000000}, []float64{162000000.000000, 144.000000}, []float64{182000000.000000, -180.000000}, []float64{202000000.000000, 1000.000000}},
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
	{
		Plan: []graphite.Function{
			graphite.Function{
				Name:       "aggregateWithWildcards",
				Arguments:  []string{"SWAP", "last", "2"},
				Parameters: make(map[string]string),
			},
		},
		Samples: []OperatorGTSTest{
			OperatorGTSTest{
				SamplePrefix: BucketizeGtsList(BuildGtsList([]FloatGeoTimeSeries{emptyGTS, gts1, gts1})),
				GTSResult: []FloatGeoTimeSeries{
					FloatGeoTimeSeries{
						Class:  "empty",
						Labels: map[string]string{"label": "42"},
						Attrs:  map[string]string{},
						Values: [][]float64{},
					},
					FloatGeoTimeSeries{
						Class:  "sample",
						Labels: map[string]string{"label": "42"},
						Attrs:  map[string]string{},
						Values: [][]float64{[]float64{205000000.000000, 32.666667}, []float64{145000000.000000, 10.000000}, []float64{85000000.000000, -0.333333}, []float64{25000000.000000, 1.000000}},
					},
					FloatGeoTimeSeries{
						Class:  "sample",
						Labels: map[string]string{"label": "42"},
						Attrs:  map[string]string{},
						Values: [][]float64{[]float64{205000000.000000, 32.666667}, []float64{145000000.000000, 10.000000}, []float64{85000000.000000, -0.333333}, []float64{25000000.000000, 1.000000}},
					},
				},

				SeriesTests: seriesEqualityTestMap,
			},
			OperatorGTSTest{
				SamplePrefix: BucketizeGtsList(BuildGtsList([]FloatGeoTimeSeries{emptyGTS, gts1})),
				GTSResult: []FloatGeoTimeSeries{
					FloatGeoTimeSeries{
						Class:  "empty",
						Labels: map[string]string{"label": "42"},
						Attrs:  map[string]string{},
						Values: [][]float64{},
					},
					FloatGeoTimeSeries{
						Class:  "sample",
						Labels: map[string]string{"label": "42"},
						Attrs:  map[string]string{},
						Values: [][]float64{[]float64{205000000.000000, 32.666667}, []float64{145000000.000000, 10.000000}, []float64{85000000.000000, -0.333333}, []float64{25000000.000000, 1.000000}},
					},
				},

				SeriesTests: seriesEqualityTestMap,
			},
			OperatorGTSTest{
				SamplePrefix: BucketizeGtsList(BuildGtsList([]FloatGeoTimeSeries{emptyGTS})),
				GTSResult: []FloatGeoTimeSeries{
					FloatGeoTimeSeries{
						Class:  "empty",
						Labels: map[string]string{"label": "42"},
						Attrs:  map[string]string{},
						Values: [][]float64{},
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
						Values: [][]float64{[]float64{205000000.000000, 32.666667}, []float64{145000000.000000, 10.000000}, []float64{85000000.000000, -0.333333}, []float64{25000000.000000, 1.000000}},
					},
					FloatGeoTimeSeries{
						Class:  "sample",
						Labels: map[string]string{"label": "42"},
						Attrs:  map[string]string{},
						Values: [][]float64{[]float64{205000000.000000, 32.666667}, []float64{145000000.000000, 10.000000}, []float64{85000000.000000, -0.333333}, []float64{25000000.000000, 1.000000}},
					},
				},

				SeriesTests: seriesEqualityTestMap,
			},
			OperatorGTSTest{
				SamplePrefix: BucketizeGtsList(BuildGtsList([]FloatGeoTimeSeries{gts1, gts2, gts3})),
				GTSResult: []FloatGeoTimeSeries{
					FloatGeoTimeSeries{
						Class:  "sample",
						Labels: map[string]string{"label": "42"},
						Attrs:  map[string]string{},
						Values: [][]float64{[]float64{205000000.000000, 32.666667}, []float64{145000000.000000, 10.000000}, []float64{85000000.000000, -0.333333}, []float64{25000000.000000, 1.000000}},
					},
					FloatGeoTimeSeries{
						Class:  "sample",
						Labels: map[string]string{"other": "test", "label": "41"},
						Attrs:  map[string]string{},
						Values: [][]float64{[]float64{205000000.000000, 9.000000}, []float64{145000000.000000, 5.500000}, []float64{85000000.000000, 2.000000}, []float64{25000000.000000, 1.000000}},
					},
					FloatGeoTimeSeries{
						Class:  "sample",
						Labels: map[string]string{"label": "41", "other": "test"},
						Attrs:  map[string]string{},
						Values: [][]float64{[]float64{205000000.000000, 1.000000}, []float64{145000000.000000, 0.500000}, []float64{85000000.000000, 1.500000}, []float64{25000000.000000, 1.000000}},
					},
				},

				SeriesTests: seriesEqualityTestMap,
			},
			OperatorGTSTest{
				SamplePrefix: BucketizeGtsList(BuildGtsList([]FloatGeoTimeSeries{gts1, gts2})),
				GTSResult: []FloatGeoTimeSeries{
					FloatGeoTimeSeries{
						Class:  "sample",
						Labels: map[string]string{"label": "42"},
						Attrs:  map[string]string{},
						Values: [][]float64{[]float64{205000000.000000, 32.666667}, []float64{145000000.000000, 10.000000}, []float64{85000000.000000, -0.333333}, []float64{25000000.000000, 1.000000}},
					},
					FloatGeoTimeSeries{
						Class:  "sample",
						Labels: map[string]string{"other": "test", "label": "41"},
						Attrs:  map[string]string{},
						Values: [][]float64{[]float64{205000000.000000, 9.000000}, []float64{145000000.000000, 5.500000}, []float64{85000000.000000, 2.000000}, []float64{25000000.000000, 1.000000}},
					},
				},

				SeriesTests: seriesEqualityTestMap,
			},
			OperatorGTSTest{
				SamplePrefix: BucketizeGtsList(BuildGtsList([]FloatGeoTimeSeries{gts1, gts3})),
				GTSResult: []FloatGeoTimeSeries{
					FloatGeoTimeSeries{
						Class:  "sample",
						Labels: map[string]string{"label": "42"},
						Attrs:  map[string]string{},
						Values: [][]float64{[]float64{205000000.000000, 32.666667}, []float64{145000000.000000, 10.000000}, []float64{85000000.000000, -0.333333}, []float64{25000000.000000, 1.000000}},
					},
					FloatGeoTimeSeries{
						Class:  "sample",
						Labels: map[string]string{"other": "test", "label": "41"},
						Attrs:  map[string]string{},
						Values: [][]float64{[]float64{205000000.000000, 1.000000}, []float64{145000000.000000, 0.500000}, []float64{85000000.000000, 1.500000}, []float64{25000000.000000, 1.000000}},
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
						Values: [][]float64{[]float64{205000000.000000, 32.666667}, []float64{145000000.000000, 10.000000}, []float64{85000000.000000, -0.333333}, []float64{25000000.000000, 1.000000}},
					},
				},

				SeriesTests: seriesEqualityTestMap,
			},
			OperatorGTSTest{
				SamplePrefix: BuildGtsList([]FloatGeoTimeSeries{emptyGTS, gts1, gts1}),
				GTSResult: []FloatGeoTimeSeries{
					FloatGeoTimeSeries{
						Class:  "empty",
						Labels: map[string]string{"label": "42"},
						Attrs:  map[string]string{},
						Values: [][]float64{},
					},
					FloatGeoTimeSeries{
						Class:  "sample",
						Labels: map[string]string{"label": "42"},
						Attrs:  map[string]string{},
						Values: [][]float64{[]float64{0.000000, 1.000000}, []float64{35000000.000000, -1.000000}, []float64{60000000.000000, 2.000000}, []float64{72000000.000000, -2.000000}, []float64{88000000.000000, 4.000000}, []float64{112000000.000000, 8.000000}, []float64{122000000.000000, 12.000000}, []float64{132000000.000000, 16.000000}, []float64{162000000.000000, 18.000000}, []float64{182000000.000000, -20.000000}, []float64{202000000.000000, 100.000000}},
					},
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
				SamplePrefix: BuildGtsList([]FloatGeoTimeSeries{emptyGTS, gts1}),
				GTSResult: []FloatGeoTimeSeries{
					FloatGeoTimeSeries{
						Class:  "empty",
						Labels: map[string]string{"label": "42"},
						Attrs:  map[string]string{},
						Values: [][]float64{},
					},
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
				SamplePrefix: BuildGtsList([]FloatGeoTimeSeries{emptyGTS}),
				GTSResult: []FloatGeoTimeSeries{
					FloatGeoTimeSeries{
						Class:  "empty",
						Labels: map[string]string{"label": "42"},
						Attrs:  map[string]string{},
						Values: [][]float64{},
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
						Labels: map[string]string{"label": "42"},
						Attrs:  map[string]string{},
						Values: [][]float64{[]float64{0.000000, 1.000000}, []float64{35000000.000000, -1.000000}, []float64{60000000.000000, 2.000000}, []float64{72000000.000000, -2.000000}, []float64{88000000.000000, 4.000000}, []float64{112000000.000000, 8.000000}, []float64{122000000.000000, 12.000000}, []float64{132000000.000000, 16.000000}, []float64{162000000.000000, 18.000000}, []float64{182000000.000000, -20.000000}, []float64{202000000.000000, 100.000000}},
					},
					FloatGeoTimeSeries{
						Class:  "sample",
						Labels: map[string]string{"other": "test", "label": "41"},
						Attrs:  map[string]string{},
						Values: [][]float64{[]float64{0.000000, 1.000000}, []float64{35000000.000000, 1.000000}, []float64{60000000.000000, 2.000000}, []float64{72000000.000000, 3.000000}, []float64{88000000.000000, 4.000000}, []float64{112000000.000000, 5.000000}, []float64{122000000.000000, 6.000000}, []float64{132000000.000000, 7.000000}, []float64{162000000.000000, 8.000000}, []float64{182000000.000000, 9.000000}, []float64{202000000.000000, 10.000000}},
					},
					FloatGeoTimeSeries{
						Class:  "sample",
						Labels: map[string]string{"other": "test", "label": "41"},
						Attrs:  map[string]string{},
						Values: [][]float64{[]float64{0.000000, 1.000000}, []float64{40000000.000000, 2.000000}, []float64{82000000.000000, 1.000000}, []float64{110000000.000000, 0.000000}, []float64{129000000.000000, 1.000000}, []float64{159000000.000000, 2.000000}, []float64{192000000.000000, 0.000000}, []float64{205000000.000000, 1.000000}},
					},
				},

				SeriesTests: seriesEqualityTestMap,
			},
			OperatorGTSTest{
				SamplePrefix: BuildGtsList([]FloatGeoTimeSeries{gts1, gts2}),
				GTSResult: []FloatGeoTimeSeries{
					FloatGeoTimeSeries{
						Class:  "sample",
						Labels: map[string]string{"label": "42"},
						Attrs:  map[string]string{},
						Values: [][]float64{[]float64{0.000000, 1.000000}, []float64{35000000.000000, -1.000000}, []float64{60000000.000000, 2.000000}, []float64{72000000.000000, -2.000000}, []float64{88000000.000000, 4.000000}, []float64{112000000.000000, 8.000000}, []float64{122000000.000000, 12.000000}, []float64{132000000.000000, 16.000000}, []float64{162000000.000000, 18.000000}, []float64{182000000.000000, -20.000000}, []float64{202000000.000000, 100.000000}},
					},
					FloatGeoTimeSeries{
						Class:  "sample",
						Labels: map[string]string{"other": "test", "label": "41"},
						Attrs:  map[string]string{},
						Values: [][]float64{[]float64{0.000000, 1.000000}, []float64{35000000.000000, 1.000000}, []float64{60000000.000000, 2.000000}, []float64{72000000.000000, 3.000000}, []float64{88000000.000000, 4.000000}, []float64{112000000.000000, 5.000000}, []float64{122000000.000000, 6.000000}, []float64{132000000.000000, 7.000000}, []float64{162000000.000000, 8.000000}, []float64{182000000.000000, 9.000000}, []float64{202000000.000000, 10.000000}},
					},
				},

				SeriesTests: seriesEqualityTestMap,
			},
			OperatorGTSTest{
				SamplePrefix: BuildGtsList([]FloatGeoTimeSeries{gts1, gts3}),
				GTSResult: []FloatGeoTimeSeries{
					FloatGeoTimeSeries{
						Class:  "sample",
						Labels: map[string]string{"label": "42"},
						Attrs:  map[string]string{},
						Values: [][]float64{[]float64{0.000000, 1.000000}, []float64{35000000.000000, -1.000000}, []float64{60000000.000000, 2.000000}, []float64{72000000.000000, -2.000000}, []float64{88000000.000000, 4.000000}, []float64{112000000.000000, 8.000000}, []float64{122000000.000000, 12.000000}, []float64{132000000.000000, 16.000000}, []float64{162000000.000000, 18.000000}, []float64{182000000.000000, -20.000000}, []float64{202000000.000000, 100.000000}},
					},
					FloatGeoTimeSeries{
						Class:  "sample",
						Labels: map[string]string{"other": "test", "label": "41"},
						Attrs:  map[string]string{},
						Values: [][]float64{[]float64{0.000000, 1.000000}, []float64{40000000.000000, 2.000000}, []float64{82000000.000000, 1.000000}, []float64{110000000.000000, 0.000000}, []float64{129000000.000000, 1.000000}, []float64{159000000.000000, 2.000000}, []float64{192000000.000000, 0.000000}, []float64{205000000.000000, 1.000000}},
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
