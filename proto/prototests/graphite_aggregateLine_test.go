package prototests

import (
	"testing"

	"github.com/ovh/erlenmeyer/proto/graphite"
)

//
// @code auto-generated
// Test can be started with
// go test proto/prototests/graphite_aggregateLine_test.go proto/prototests/exec_test.go -v
//

// TestGraphiteAggregateLine process Invert Graphite Unit tests
func TestGraphiteAggregateLine(t *testing.T) {
	RunTest(t, graphiteGraphiteAggregateLine, "")
}

var graphiteGraphiteAggregateLine = []unitTests{

	{
		Plan: []graphite.Function{
			{
				Name:       "aggregateLine",
				Arguments:  []string{"SWAP", "average"},
				Parameters: make(map[string]string),
			},
		},
		Samples: []OperatorGTSTest{
			{
				SamplePrefix: BucketizeGtsList(BuildGtsList([]FloatGeoTimeSeries{emptyGTS, gts1, gts1})),
				GTSResult: []FloatGeoTimeSeries{
					{
						Class:  "sample",
						Labels: map[string]string{"label": "42"},
						Attrs:  map[string]string{},
						Values: [][]float64{{205000000.000000, 10.833333}, {145000000.000000, 10.833333}, {85000000.000000, 10.833333}, {25000000.000000, 10.833333}},
					},
					{
						Class:  "sample",
						Labels: map[string]string{"label": "42"},
						Attrs:  map[string]string{},
						Values: [][]float64{{205000000.000000, 10.833333}, {145000000.000000, 10.833333}, {85000000.000000, 10.833333}, {25000000.000000, 10.833333}},
					},
				},

				SeriesTests: seriesEqualityTestMap,
			},
			{
				SamplePrefix: BucketizeGtsList(BuildGtsList([]FloatGeoTimeSeries{emptyGTS, gts1})),
				GTSResult: []FloatGeoTimeSeries{
					{
						Class:  "sample",
						Labels: map[string]string{"label": "42"},
						Attrs:  map[string]string{},
						Values: [][]float64{{205000000.000000, 10.833333}, {145000000.000000, 10.833333}, {85000000.000000, 10.833333}, {25000000.000000, 10.833333}},
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
						Values: [][]float64{{205000000.000000, 10.833333}, {145000000.000000, 10.833333}, {85000000.000000, 10.833333}, {25000000.000000, 10.833333}},
					},
					{
						Class:  "sample",
						Labels: map[string]string{"label": "42"},
						Attrs:  map[string]string{},
						Values: [][]float64{{205000000.000000, 10.833333}, {145000000.000000, 10.833333}, {85000000.000000, 10.833333}, {25000000.000000, 10.833333}},
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
						Values: [][]float64{{205000000.000000, 10.833333}, {145000000.000000, 10.833333}, {85000000.000000, 10.833333}, {25000000.000000, 10.833333}},
					},
					{
						Class:  "sample",
						Labels: map[string]string{"other": "test", "label": "41"},
						Attrs:  map[string]string{},
						Values: [][]float64{{205000000.000000, 4.375000}, {145000000.000000, 4.375000}, {85000000.000000, 4.375000}, {25000000.000000, 4.375000}},
					},
					{
						Class:  "sample",
						Labels: map[string]string{"other": "test", "label": "41"},
						Attrs:  map[string]string{},
						Values: [][]float64{{205000000.000000, 1.000000}, {145000000.000000, 1.000000}, {85000000.000000, 1.000000}, {25000000.000000, 1.000000}},
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
						Values: [][]float64{{205000000.000000, 10.833333}, {145000000.000000, 10.833333}, {85000000.000000, 10.833333}, {25000000.000000, 10.833333}},
					},
					{
						Class:  "sample",
						Labels: map[string]string{"other": "test", "label": "41"},
						Attrs:  map[string]string{},
						Values: [][]float64{{205000000.000000, 4.375000}, {145000000.000000, 4.375000}, {85000000.000000, 4.375000}, {25000000.000000, 4.375000}},
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
						Values: [][]float64{{205000000.000000, 10.833333}, {145000000.000000, 10.833333}, {85000000.000000, 10.833333}, {25000000.000000, 10.833333}},
					},
					{
						Class:  "sample",
						Labels: map[string]string{"other": "test", "label": "41"},
						Attrs:  map[string]string{},
						Values: [][]float64{{205000000.000000, 1.000000}, {145000000.000000, 1.000000}, {85000000.000000, 1.000000}, {25000000.000000, 1.000000}},
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
						Values: [][]float64{{205000000.000000, 10.833333}, {145000000.000000, 10.833333}, {85000000.000000, 10.833333}, {25000000.000000, 10.833333}},
					},
				},

				SeriesTests: seriesEqualityTestMap,
			},
		},
	},
	{
		Plan: []graphite.Function{
			{
				Name:       "aggregateLine",
				Arguments:  []string{"SWAP", "median"},
				Parameters: make(map[string]string),
			},
		},
		Samples: []OperatorGTSTest{
			{
				SamplePrefix: BucketizeGtsList(BuildGtsList([]FloatGeoTimeSeries{emptyGTS, gts1, gts1})),
				SampleSuffix: `
		<%
			DUP TOSELECTOR SWAP true MUSIGMA DROP TOSTRING + 
		%>
		SORTBY
		`,
				GTSResult: []FloatGeoTimeSeries{
					{
						Class:  "sample",
						Labels: map[string]string{"label": "42"},
						Attrs:  map[string]string{},
						Values: [][]float64{{205000000.000000, 5.500000}, {145000000.000000, 5.500000}, {85000000.000000, 5.500000}, {25000000.000000, 5.500000}},
					},
					{
						Class:  "sample",
						Labels: map[string]string{"label": "42"},
						Attrs:  map[string]string{},
						Values: [][]float64{{205000000.000000, 5.500000}, {145000000.000000, 5.500000}, {85000000.000000, 5.500000}, {25000000.000000, 5.500000}},
					},
				},

				SeriesTests: seriesEqualityTestMap,
			},
			{
				SamplePrefix: BucketizeGtsList(BuildGtsList([]FloatGeoTimeSeries{emptyGTS, gts1})),
				SampleSuffix: `
		<%
			DUP TOSELECTOR SWAP true MUSIGMA DROP TOSTRING + 
		%>
		SORTBY
		`,
				GTSResult: []FloatGeoTimeSeries{
					{
						Class:  "sample",
						Labels: map[string]string{"label": "42"},
						Attrs:  map[string]string{},
						Values: [][]float64{{205000000.000000, 5.500000}, {145000000.000000, 5.500000}, {85000000.000000, 5.500000}, {25000000.000000, 5.500000}},
					},
				},

				SeriesTests: seriesEqualityTestMap,
			},
			{
				SamplePrefix: BucketizeGtsList(BuildGtsList([]FloatGeoTimeSeries{gts1, gts1})),
				SampleSuffix: `
		<%
			DUP TOSELECTOR SWAP true MUSIGMA DROP TOSTRING + 
		%>
		SORTBY
		`,
				GTSResult: []FloatGeoTimeSeries{
					{
						Class:  "sample",
						Labels: map[string]string{"label": "42"},
						Attrs:  map[string]string{},
						Values: [][]float64{{205000000.000000, 5.500000}, {145000000.000000, 5.500000}, {85000000.000000, 5.500000}, {25000000.000000, 5.500000}},
					},
					{
						Class:  "sample",
						Labels: map[string]string{"label": "42"},
						Attrs:  map[string]string{},
						Values: [][]float64{{205000000.000000, 5.500000}, {145000000.000000, 5.500000}, {85000000.000000, 5.500000}, {25000000.000000, 5.500000}},
					},
				},

				SeriesTests: seriesEqualityTestMap,
			},
			{
				SamplePrefix: BucketizeGtsList(BuildGtsList([]FloatGeoTimeSeries{gts1, gts2, gts3})),
				SampleSuffix: `
		<%
			DUP TOSELECTOR SWAP true MUSIGMA DROP TOSTRING + 
		%>
		SORTBY
		`,
				GTSResult: []FloatGeoTimeSeries{
					{
						Class:  "sample",
						Labels: map[string]string{"other": "test", "label": "41"},
						Attrs:  map[string]string{},
						Values: [][]float64{{205000000.000000, 1.000000}, {145000000.000000, 1.000000}, {85000000.000000, 1.000000}, {25000000.000000, 1.000000}},
					},
					{
						Class:  "sample",
						Labels: map[string]string{"other": "test", "label": "41"},
						Attrs:  map[string]string{},
						Values: [][]float64{{205000000.000000, 3.750000}, {145000000.000000, 3.750000}, {85000000.000000, 3.750000}, {25000000.000000, 3.750000}},
					},
					{
						Class:  "sample",
						Labels: map[string]string{"label": "42"},
						Attrs:  map[string]string{},
						Values: [][]float64{{205000000.000000, 5.500000}, {145000000.000000, 5.500000}, {85000000.000000, 5.500000}, {25000000.000000, 5.500000}},
					},
				},

				SeriesTests: seriesEqualityTestMap,
			},
			{
				SamplePrefix: BucketizeGtsList(BuildGtsList([]FloatGeoTimeSeries{gts1, gts2})),
				SampleSuffix: `
		<%
			DUP TOSELECTOR SWAP true MUSIGMA DROP TOSTRING + 
		%>
		SORTBY
		`,
				GTSResult: []FloatGeoTimeSeries{
					{
						Class:  "sample",
						Labels: map[string]string{"other": "test", "label": "41"},
						Attrs:  map[string]string{},
						Values: [][]float64{{205000000.000000, 3.750000}, {145000000.000000, 3.750000}, {85000000.000000, 3.750000}, {25000000.000000, 3.750000}},
					},
					{
						Class:  "sample",
						Labels: map[string]string{"label": "42"},
						Attrs:  map[string]string{},
						Values: [][]float64{{205000000.000000, 5.500000}, {145000000.000000, 5.500000}, {85000000.000000, 5.500000}, {25000000.000000, 5.500000}},
					},
				},

				SeriesTests: seriesEqualityTestMap,
			},
			{
				SamplePrefix: BucketizeGtsList(BuildGtsList([]FloatGeoTimeSeries{gts1, gts3})),
				SampleSuffix: `
		<%
			DUP TOSELECTOR SWAP true MUSIGMA DROP TOSTRING + 
		%>
		SORTBY
		`,
				GTSResult: []FloatGeoTimeSeries{
					{
						Class:  "sample",
						Labels: map[string]string{"other": "test", "label": "41"},
						Attrs:  map[string]string{},
						Values: [][]float64{{205000000.000000, 1.000000}, {145000000.000000, 1.000000}, {85000000.000000, 1.000000}, {25000000.000000, 1.000000}},
					},
					{
						Class:  "sample",
						Labels: map[string]string{"label": "42"},
						Attrs:  map[string]string{},
						Values: [][]float64{{205000000.000000, 5.500000}, {145000000.000000, 5.500000}, {85000000.000000, 5.500000}, {25000000.000000, 5.500000}},
					},
				},

				SeriesTests: seriesEqualityTestMap,
			},
			{
				SamplePrefix: BucketizeGtsList(BuildGtsList([]FloatGeoTimeSeries{gts1})),
				SampleSuffix: `
		<%
			DUP TOSELECTOR SWAP true MUSIGMA DROP TOSTRING + 
		%>
		SORTBY
		`,
				GTSResult: []FloatGeoTimeSeries{
					{
						Class:  "sample",
						Labels: map[string]string{"label": "42"},
						Attrs:  map[string]string{},
						Values: [][]float64{{205000000.000000, 5.500000}, {145000000.000000, 5.500000}, {85000000.000000, 5.500000}, {25000000.000000, 5.500000}},
					},
				},

				SeriesTests: seriesEqualityTestMap,
			},
		},
	},
	{
		Plan: []graphite.Function{
			{
				Name:       "aggregateLine",
				Arguments:  []string{"SWAP", "sum"},
				Parameters: make(map[string]string),
			},
		},
		Samples: []OperatorGTSTest{
			{
				SamplePrefix: BucketizeGtsList(BuildGtsList([]FloatGeoTimeSeries{emptyGTS, gts1, gts1})),
				SampleSuffix: `
		<%
			DUP TOSELECTOR SWAP true MUSIGMA DROP TOSTRING + 
		%>
		SORTBY
		`,
				GTSResult: []FloatGeoTimeSeries{
					{
						Class:  "sample",
						Labels: map[string]string{"label": "42"},
						Attrs:  map[string]string{},
						Values: [][]float64{{205000000.000000, 43.333333}, {145000000.000000, 43.333333}, {85000000.000000, 43.333333}, {25000000.000000, 43.333333}},
					},
					{
						Class:  "sample",
						Labels: map[string]string{"label": "42"},
						Attrs:  map[string]string{},
						Values: [][]float64{{205000000.000000, 43.333333}, {145000000.000000, 43.333333}, {85000000.000000, 43.333333}, {25000000.000000, 43.333333}},
					},
				},

				SeriesTests: seriesEqualityTestMap,
			},
			{
				SamplePrefix: BucketizeGtsList(BuildGtsList([]FloatGeoTimeSeries{emptyGTS, gts1})),
				SampleSuffix: `
		<%
			DUP TOSELECTOR SWAP true MUSIGMA DROP TOSTRING + 
		%>
		SORTBY
		`,
				GTSResult: []FloatGeoTimeSeries{
					{
						Class:  "sample",
						Labels: map[string]string{"label": "42"},
						Attrs:  map[string]string{},
						Values: [][]float64{{205000000.000000, 43.333333}, {145000000.000000, 43.333333}, {85000000.000000, 43.333333}, {25000000.000000, 43.333333}},
					},
				},

				SeriesTests: seriesEqualityTestMap,
			},
			{
				SamplePrefix: BucketizeGtsList(BuildGtsList([]FloatGeoTimeSeries{gts1, gts1})),
				SampleSuffix: `
		<%
			DUP TOSELECTOR SWAP true MUSIGMA DROP TOSTRING + 
		%>
		SORTBY
		`,
				GTSResult: []FloatGeoTimeSeries{
					{
						Class:  "sample",
						Labels: map[string]string{"label": "42"},
						Attrs:  map[string]string{},
						Values: [][]float64{{205000000.000000, 43.333333}, {145000000.000000, 43.333333}, {85000000.000000, 43.333333}, {25000000.000000, 43.333333}},
					},
					{
						Class:  "sample",
						Labels: map[string]string{"label": "42"},
						Attrs:  map[string]string{},
						Values: [][]float64{{205000000.000000, 43.333333}, {145000000.000000, 43.333333}, {85000000.000000, 43.333333}, {25000000.000000, 43.333333}},
					},
				},

				SeriesTests: seriesEqualityTestMap,
			},
			{
				SamplePrefix: BucketizeGtsList(BuildGtsList([]FloatGeoTimeSeries{gts1, gts2, gts3})),
				SampleSuffix: `
		<%
			DUP TOSELECTOR SWAP true MUSIGMA DROP TOSTRING + 
		%>
		SORTBY
		`,
				GTSResult: []FloatGeoTimeSeries{
					{
						Class:  "sample",
						Labels: map[string]string{"other": "test", "label": "41"},
						Attrs:  map[string]string{},
						Values: [][]float64{{205000000.000000, 17.500000}, {145000000.000000, 17.500000}, {85000000.000000, 17.500000}, {25000000.000000, 17.500000}},
					},
					{
						Class:  "sample",
						Labels: map[string]string{"other": "test", "label": "41"},
						Attrs:  map[string]string{},
						Values: [][]float64{{205000000.000000, 4.000000}, {145000000.000000, 4.000000}, {85000000.000000, 4.000000}, {25000000.000000, 4.000000}},
					},
					{
						Class:  "sample",
						Labels: map[string]string{"label": "42"},
						Attrs:  map[string]string{},
						Values: [][]float64{{205000000.000000, 43.333333}, {145000000.000000, 43.333333}, {85000000.000000, 43.333333}, {25000000.000000, 43.333333}},
					},
				},

				SeriesTests: seriesEqualityTestMap,
			},
			{
				SamplePrefix: BucketizeGtsList(BuildGtsList([]FloatGeoTimeSeries{gts1, gts2})),
				SampleSuffix: `
		<%
			DUP TOSELECTOR SWAP true MUSIGMA DROP TOSTRING + 
		%>
		SORTBY
		`,
				GTSResult: []FloatGeoTimeSeries{
					{
						Class:  "sample",
						Labels: map[string]string{"other": "test", "label": "41"},
						Attrs:  map[string]string{},
						Values: [][]float64{{205000000.000000, 17.500000}, {145000000.000000, 17.500000}, {85000000.000000, 17.500000}, {25000000.000000, 17.500000}},
					},
					{
						Class:  "sample",
						Labels: map[string]string{"label": "42"},
						Attrs:  map[string]string{},
						Values: [][]float64{{205000000.000000, 43.333333}, {145000000.000000, 43.333333}, {85000000.000000, 43.333333}, {25000000.000000, 43.333333}},
					},
				},

				SeriesTests: seriesEqualityTestMap,
			},
			{
				SamplePrefix: BucketizeGtsList(BuildGtsList([]FloatGeoTimeSeries{gts1, gts3})),
				SampleSuffix: `
		<%
			DUP TOSELECTOR SWAP true MUSIGMA DROP TOSTRING + 
		%>
		SORTBY
		`,
				GTSResult: []FloatGeoTimeSeries{
					{
						Class:  "sample",
						Labels: map[string]string{"other": "test", "label": "41"},
						Attrs:  map[string]string{},
						Values: [][]float64{{205000000.000000, 4.000000}, {145000000.000000, 4.000000}, {85000000.000000, 4.000000}, {25000000.000000, 4.000000}},
					},
					{
						Class:  "sample",
						Labels: map[string]string{"label": "42"},
						Attrs:  map[string]string{},
						Values: [][]float64{{205000000.000000, 43.333333}, {145000000.000000, 43.333333}, {85000000.000000, 43.333333}, {25000000.000000, 43.333333}},
					},
				},

				SeriesTests: seriesEqualityTestMap,
			},
			{
				SamplePrefix: BucketizeGtsList(BuildGtsList([]FloatGeoTimeSeries{gts1})),
				SampleSuffix: `
		<%
			DUP TOSELECTOR SWAP true MUSIGMA DROP TOSTRING + 
		%>
		SORTBY
		`,
				GTSResult: []FloatGeoTimeSeries{
					{
						Class:  "sample",
						Labels: map[string]string{"label": "42"},
						Attrs:  map[string]string{},
						Values: [][]float64{{205000000.000000, 43.333333}, {145000000.000000, 43.333333}, {85000000.000000, 43.333333}, {25000000.000000, 43.333333}},
					},
				},

				SeriesTests: seriesEqualityTestMap,
			},
		},
	},
	{
		Plan: []graphite.Function{
			{
				Name:       "aggregateLine",
				Arguments:  []string{"SWAP", "min"},
				Parameters: make(map[string]string),
			},
		},
		Samples: []OperatorGTSTest{
			{
				SamplePrefix: BucketizeGtsList(BuildGtsList([]FloatGeoTimeSeries{emptyGTS, gts1, gts1})),
				SampleSuffix: `
		<%
			DUP TOSELECTOR SWAP true MUSIGMA DROP TOSTRING + 
		%>
		SORTBY
		`,
				GTSResult: []FloatGeoTimeSeries{
					{
						Class:  "sample",
						Labels: map[string]string{"label": "42"},
						Attrs:  map[string]string{},
						Values: [][]float64{{205000000.000000, -0.333333}, {145000000.000000, -0.333333}, {85000000.000000, -0.333333}, {25000000.000000, -0.333333}},
					},
					{
						Class:  "sample",
						Labels: map[string]string{"label": "42"},
						Attrs:  map[string]string{},
						Values: [][]float64{{205000000.000000, -0.333333}, {145000000.000000, -0.333333}, {85000000.000000, -0.333333}, {25000000.000000, -0.333333}},
					},
				},

				SeriesTests: seriesEqualityTestMap,
			},
			{
				SamplePrefix: BucketizeGtsList(BuildGtsList([]FloatGeoTimeSeries{emptyGTS, gts1})),
				SampleSuffix: `
		<%
			DUP TOSELECTOR SWAP true MUSIGMA DROP TOSTRING + 
		%>
		SORTBY
		`,
				GTSResult: []FloatGeoTimeSeries{
					{
						Class:  "sample",
						Labels: map[string]string{"label": "42"},
						Attrs:  map[string]string{},
						Values: [][]float64{{205000000.000000, -0.333333}, {145000000.000000, -0.333333}, {85000000.000000, -0.333333}, {25000000.000000, -0.333333}},
					},
				},

				SeriesTests: seriesEqualityTestMap,
			},
			{
				SamplePrefix: BucketizeGtsList(BuildGtsList([]FloatGeoTimeSeries{gts1, gts1})),
				SampleSuffix: `
		<%
			DUP TOSELECTOR SWAP true MUSIGMA DROP TOSTRING + 
		%>
		SORTBY
		`,
				GTSResult: []FloatGeoTimeSeries{
					{
						Class:  "sample",
						Labels: map[string]string{"label": "42"},
						Attrs:  map[string]string{},
						Values: [][]float64{{205000000.000000, -0.333333}, {145000000.000000, -0.333333}, {85000000.000000, -0.333333}, {25000000.000000, -0.333333}},
					},
					{
						Class:  "sample",
						Labels: map[string]string{"label": "42"},
						Attrs:  map[string]string{},
						Values: [][]float64{{205000000.000000, -0.333333}, {145000000.000000, -0.333333}, {85000000.000000, -0.333333}, {25000000.000000, -0.333333}},
					},
				},

				SeriesTests: seriesEqualityTestMap,
			},
			{
				SamplePrefix: BucketizeGtsList(BuildGtsList([]FloatGeoTimeSeries{gts1, gts2, gts3})),
				SampleSuffix: `
		<%
			DUP TOSELECTOR SWAP true MUSIGMA DROP TOSTRING + 
		%>
		SORTBY
		`,
				GTSResult: []FloatGeoTimeSeries{
					{
						Class:  "sample",
						Labels: map[string]string{"other": "test", "label": "41"},
						Attrs:  map[string]string{},
						Values: [][]float64{{205000000.000000, 0.500000}, {145000000.000000, 0.500000}, {85000000.000000, 0.500000}, {25000000.000000, 0.500000}},
					},
					{
						Class:  "sample",
						Labels: map[string]string{"other": "test", "label": "41"},
						Attrs:  map[string]string{},
						Values: [][]float64{{205000000.000000, 1.000000}, {145000000.000000, 1.000000}, {85000000.000000, 1.000000}, {25000000.000000, 1.000000}},
					},
					{
						Class:  "sample",
						Labels: map[string]string{"label": "42"},
						Attrs:  map[string]string{},
						Values: [][]float64{{205000000.000000, -0.333333}, {145000000.000000, -0.333333}, {85000000.000000, -0.333333}, {25000000.000000, -0.333333}},
					},
				},

				SeriesTests: seriesEqualityTestMap,
			},
			{
				SamplePrefix: BucketizeGtsList(BuildGtsList([]FloatGeoTimeSeries{gts1, gts2})),
				SampleSuffix: `
		<%
			DUP TOSELECTOR SWAP true MUSIGMA DROP TOSTRING + 
		%>
		SORTBY
		`,
				GTSResult: []FloatGeoTimeSeries{
					{
						Class:  "sample",
						Labels: map[string]string{"other": "test", "label": "41"},
						Attrs:  map[string]string{},
						Values: [][]float64{{205000000.000000, 1.000000}, {145000000.000000, 1.000000}, {85000000.000000, 1.000000}, {25000000.000000, 1.000000}},
					},
					{
						Class:  "sample",
						Labels: map[string]string{"label": "42"},
						Attrs:  map[string]string{},
						Values: [][]float64{{205000000.000000, -0.333333}, {145000000.000000, -0.333333}, {85000000.000000, -0.333333}, {25000000.000000, -0.333333}},
					},
				},

				SeriesTests: seriesEqualityTestMap,
			},
			{
				SamplePrefix: BucketizeGtsList(BuildGtsList([]FloatGeoTimeSeries{gts1, gts3})),
				SampleSuffix: `
		<%
			DUP TOSELECTOR SWAP true MUSIGMA DROP TOSTRING + 
		%>
		SORTBY
		`,
				GTSResult: []FloatGeoTimeSeries{
					{
						Class:  "sample",
						Labels: map[string]string{"other": "test", "label": "41"},
						Attrs:  map[string]string{},
						Values: [][]float64{{205000000.000000, 0.500000}, {145000000.000000, 0.500000}, {85000000.000000, 0.500000}, {25000000.000000, 0.500000}},
					},
					{
						Class:  "sample",
						Labels: map[string]string{"label": "42"},
						Attrs:  map[string]string{},
						Values: [][]float64{{205000000.000000, -0.333333}, {145000000.000000, -0.333333}, {85000000.000000, -0.333333}, {25000000.000000, -0.333333}},
					},
				},

				SeriesTests: seriesEqualityTestMap,
			},
			{
				SamplePrefix: BucketizeGtsList(BuildGtsList([]FloatGeoTimeSeries{gts1})),
				SampleSuffix: `
		<%
			DUP TOSELECTOR SWAP true MUSIGMA DROP TOSTRING + 
		%>
		SORTBY
		`,
				GTSResult: []FloatGeoTimeSeries{
					{
						Class:  "sample",
						Labels: map[string]string{"label": "42"},
						Attrs:  map[string]string{},
						Values: [][]float64{{205000000.000000, -0.333333}, {145000000.000000, -0.333333}, {85000000.000000, -0.333333}, {25000000.000000, -0.333333}},
					},
				},

				SeriesTests: seriesEqualityTestMap,
			},
		},
	},
	{
		Plan: []graphite.Function{
			{
				Name:       "aggregateLine",
				Arguments:  []string{"SWAP", "max"},
				Parameters: make(map[string]string),
			},
		},
		Samples: []OperatorGTSTest{
			{
				SamplePrefix: BucketizeGtsList(BuildGtsList([]FloatGeoTimeSeries{emptyGTS, gts1, gts1})),
				SampleSuffix: `
		<%
			DUP TOSELECTOR SWAP true MUSIGMA DROP TOSTRING + 
		%>
		SORTBY
		`,
				GTSResult: []FloatGeoTimeSeries{
					{
						Class:  "sample",
						Labels: map[string]string{"label": "42"},
						Attrs:  map[string]string{},
						Values: [][]float64{{205000000.000000, 32.666667}, {145000000.000000, 32.666667}, {85000000.000000, 32.666667}, {25000000.000000, 32.666667}},
					},
					{
						Class:  "sample",
						Labels: map[string]string{"label": "42"},
						Attrs:  map[string]string{},
						Values: [][]float64{{205000000.000000, 32.666667}, {145000000.000000, 32.666667}, {85000000.000000, 32.666667}, {25000000.000000, 32.666667}},
					},
				},

				SeriesTests: seriesEqualityTestMap,
			},
			{
				SamplePrefix: BucketizeGtsList(BuildGtsList([]FloatGeoTimeSeries{emptyGTS, gts1})),
				SampleSuffix: `
		<%
			DUP TOSELECTOR SWAP true MUSIGMA DROP TOSTRING + 
		%>
		SORTBY
		`,
				GTSResult: []FloatGeoTimeSeries{
					{
						Class:  "sample",
						Labels: map[string]string{"label": "42"},
						Attrs:  map[string]string{},
						Values: [][]float64{{205000000.000000, 32.666667}, {145000000.000000, 32.666667}, {85000000.000000, 32.666667}, {25000000.000000, 32.666667}},
					},
				},

				SeriesTests: seriesEqualityTestMap,
			},
			{
				SamplePrefix: BucketizeGtsList(BuildGtsList([]FloatGeoTimeSeries{gts1, gts1})),
				SampleSuffix: `
		<%
			DUP TOSELECTOR SWAP true MUSIGMA DROP TOSTRING + 
		%>
		SORTBY
		`,
				GTSResult: []FloatGeoTimeSeries{
					{
						Class:  "sample",
						Labels: map[string]string{"label": "42"},
						Attrs:  map[string]string{},
						Values: [][]float64{{205000000.000000, 32.666667}, {145000000.000000, 32.666667}, {85000000.000000, 32.666667}, {25000000.000000, 32.666667}},
					},
					{
						Class:  "sample",
						Labels: map[string]string{"label": "42"},
						Attrs:  map[string]string{},
						Values: [][]float64{{205000000.000000, 32.666667}, {145000000.000000, 32.666667}, {85000000.000000, 32.666667}, {25000000.000000, 32.666667}},
					},
				},

				SeriesTests: seriesEqualityTestMap,
			},
			{
				SamplePrefix: BucketizeGtsList(BuildGtsList([]FloatGeoTimeSeries{gts1, gts2, gts3})),
				SampleSuffix: `
		<%
			DUP TOSELECTOR SWAP true MUSIGMA DROP TOSTRING + 
		%>
		SORTBY
		`,
				GTSResult: []FloatGeoTimeSeries{
					{
						Class:  "sample",
						Labels: map[string]string{"other": "test", "label": "41"},
						Attrs:  map[string]string{},
						Values: [][]float64{{205000000.000000, 1.500000}, {145000000.000000, 1.500000}, {85000000.000000, 1.500000}, {25000000.000000, 1.500000}},
					},
					{
						Class:  "sample",
						Labels: map[string]string{"other": "test", "label": "41"},
						Attrs:  map[string]string{},
						Values: [][]float64{{205000000.000000, 9.000000}, {145000000.000000, 9.000000}, {85000000.000000, 9.000000}, {25000000.000000, 9.000000}},
					},
					{
						Class:  "sample",
						Labels: map[string]string{"label": "42"},
						Attrs:  map[string]string{},
						Values: [][]float64{{205000000.000000, 32.666667}, {145000000.000000, 32.666667}, {85000000.000000, 32.666667}, {25000000.000000, 32.666667}},
					},
				},

				SeriesTests: seriesEqualityTestMap,
			},
			{
				SamplePrefix: BucketizeGtsList(BuildGtsList([]FloatGeoTimeSeries{gts1, gts2})),
				SampleSuffix: `
		<%
			DUP TOSELECTOR SWAP true MUSIGMA DROP TOSTRING + 
		%>
		SORTBY
		`,
				GTSResult: []FloatGeoTimeSeries{
					{
						Class:  "sample",
						Labels: map[string]string{"other": "test", "label": "41"},
						Attrs:  map[string]string{},
						Values: [][]float64{{205000000.000000, 9.000000}, {145000000.000000, 9.000000}, {85000000.000000, 9.000000}, {25000000.000000, 9.000000}},
					},
					{
						Class:  "sample",
						Labels: map[string]string{"label": "42"},
						Attrs:  map[string]string{},
						Values: [][]float64{{205000000.000000, 32.666667}, {145000000.000000, 32.666667}, {85000000.000000, 32.666667}, {25000000.000000, 32.666667}},
					},
				},

				SeriesTests: seriesEqualityTestMap,
			},
			{
				SamplePrefix: BucketizeGtsList(BuildGtsList([]FloatGeoTimeSeries{gts1, gts3})),
				SampleSuffix: `
		<%
			DUP TOSELECTOR SWAP true MUSIGMA DROP TOSTRING + 
		%>
		SORTBY
		`,
				GTSResult: []FloatGeoTimeSeries{
					{
						Class:  "sample",
						Labels: map[string]string{"other": "test", "label": "41"},
						Attrs:  map[string]string{},
						Values: [][]float64{{205000000.000000, 1.500000}, {145000000.000000, 1.500000}, {85000000.000000, 1.500000}, {25000000.000000, 1.500000}},
					},
					{
						Class:  "sample",
						Labels: map[string]string{"label": "42"},
						Attrs:  map[string]string{},
						Values: [][]float64{{205000000.000000, 32.666667}, {145000000.000000, 32.666667}, {85000000.000000, 32.666667}, {25000000.000000, 32.666667}},
					},
				},

				SeriesTests: seriesEqualityTestMap,
			},
			{
				SamplePrefix: BucketizeGtsList(BuildGtsList([]FloatGeoTimeSeries{gts1})),
				SampleSuffix: `
		<%
			DUP TOSELECTOR SWAP true MUSIGMA DROP TOSTRING + 
		%>
		SORTBY
		`,
				GTSResult: []FloatGeoTimeSeries{
					{
						Class:  "sample",
						Labels: map[string]string{"label": "42"},
						Attrs:  map[string]string{},
						Values: [][]float64{{205000000.000000, 32.666667}, {145000000.000000, 32.666667}, {85000000.000000, 32.666667}, {25000000.000000, 32.666667}},
					},
				},

				SeriesTests: seriesEqualityTestMap,
			},
		},
	},
	{
		Plan: []graphite.Function{
			{
				Name:       "aggregateLine",
				Arguments:  []string{"SWAP", "stddev"},
				Parameters: make(map[string]string),
			},
		},
		Samples: []OperatorGTSTest{
			{
				SamplePrefix: BucketizeGtsList(BuildGtsList([]FloatGeoTimeSeries{emptyGTS, gts1, gts1})),
				SampleSuffix: `
		<%
			DUP TOSELECTOR SWAP true MUSIGMA DROP TOSTRING + 
		%>
		SORTBY
		`,
				GTSResult: []FloatGeoTimeSeries{
					{
						Class:  "sample",
						Labels: map[string]string{"label": "42"},
						Attrs:  map[string]string{},
						Values: [][]float64{{25000000.000000, 15.261911}, {85000000.000000, 15.261911}, {145000000.000000, 15.261911}, {205000000.000000, 15.261911}},
					},
					{
						Class:  "sample",
						Labels: map[string]string{"label": "42"},
						Attrs:  map[string]string{},
						Values: [][]float64{{25000000.000000, 15.261911}, {85000000.000000, 15.261911}, {145000000.000000, 15.261911}, {205000000.000000, 15.261911}},
					},
				},

				SeriesTests: seriesEqualityTestMap,
			},
			{
				SamplePrefix: BucketizeGtsList(BuildGtsList([]FloatGeoTimeSeries{emptyGTS, gts1})),
				SampleSuffix: `
		<%
			DUP TOSELECTOR SWAP true MUSIGMA DROP TOSTRING + 
		%>
		SORTBY
		`,
				GTSResult: []FloatGeoTimeSeries{
					{
						Class:  "sample",
						Labels: map[string]string{"label": "42"},
						Attrs:  map[string]string{},
						Values: [][]float64{{25000000.000000, 15.261911}, {85000000.000000, 15.261911}, {145000000.000000, 15.261911}, {205000000.000000, 15.261911}},
					},
				},

				SeriesTests: seriesEqualityTestMap,
			},
			{
				SamplePrefix: BucketizeGtsList(BuildGtsList([]FloatGeoTimeSeries{gts1, gts1})),
				SampleSuffix: `
		<%
			DUP TOSELECTOR SWAP true MUSIGMA DROP TOSTRING + 
		%>
		SORTBY
		`,
				GTSResult: []FloatGeoTimeSeries{
					{
						Class:  "sample",
						Labels: map[string]string{"label": "42"},
						Attrs:  map[string]string{},
						Values: [][]float64{{25000000.000000, 15.261911}, {85000000.000000, 15.261911}, {145000000.000000, 15.261911}, {205000000.000000, 15.261911}},
					},
					{
						Class:  "sample",
						Labels: map[string]string{"label": "42"},
						Attrs:  map[string]string{},
						Values: [][]float64{{25000000.000000, 15.261911}, {85000000.000000, 15.261911}, {145000000.000000, 15.261911}, {205000000.000000, 15.261911}},
					},
				},

				SeriesTests: seriesEqualityTestMap,
			},
			{
				SamplePrefix: BucketizeGtsList(BuildGtsList([]FloatGeoTimeSeries{gts1, gts2, gts3})),
				SampleSuffix: `
		<%
			DUP TOSELECTOR SWAP true MUSIGMA DROP TOSTRING + 
		%>
		SORTBY
		`,
				GTSResult: []FloatGeoTimeSeries{
					{
						Class:  "sample",
						Labels: map[string]string{"other": "test", "label": "41"},
						Attrs:  map[string]string{},
						Values: [][]float64{{25000000.000000, 0.408248}, {85000000.000000, 0.408248}, {145000000.000000, 0.408248}, {205000000.000000, 0.408248}},
					},
					{
						Class:  "sample",
						Labels: map[string]string{"other": "test", "label": "41"},
						Attrs:  map[string]string{},
						Values: [][]float64{{25000000.000000, 3.637192}, {85000000.000000, 3.637192}, {145000000.000000, 3.637192}, {205000000.000000, 3.637192}},
					},
					{
						Class:  "sample",
						Labels: map[string]string{"label": "42"},
						Attrs:  map[string]string{},
						Values: [][]float64{{25000000.000000, 15.261911}, {85000000.000000, 15.261911}, {145000000.000000, 15.261911}, {205000000.000000, 15.261911}},
					},
				},

				SeriesTests: seriesEqualityTestMap,
			},
			{
				SamplePrefix: BucketizeGtsList(BuildGtsList([]FloatGeoTimeSeries{gts1, gts2})),
				SampleSuffix: `
		<%
			DUP TOSELECTOR SWAP true MUSIGMA DROP TOSTRING + 
		%>
		SORTBY
		`,
				GTSResult: []FloatGeoTimeSeries{
					{
						Class:  "sample",
						Labels: map[string]string{"other": "test", "label": "41"},
						Attrs:  map[string]string{},
						Values: [][]float64{{25000000.000000, 3.637192}, {85000000.000000, 3.637192}, {145000000.000000, 3.637192}, {205000000.000000, 3.637192}},
					},
					{
						Class:  "sample",
						Labels: map[string]string{"label": "42"},
						Attrs:  map[string]string{},
						Values: [][]float64{{25000000.000000, 15.261911}, {85000000.000000, 15.261911}, {145000000.000000, 15.261911}, {205000000.000000, 15.261911}},
					},
				},

				SeriesTests: seriesEqualityTestMap,
			},
			{
				SamplePrefix: BucketizeGtsList(BuildGtsList([]FloatGeoTimeSeries{gts1, gts3})),
				SampleSuffix: `
		<%
			DUP TOSELECTOR SWAP true MUSIGMA DROP TOSTRING + 
		%>
		SORTBY
		`,
				GTSResult: []FloatGeoTimeSeries{
					{
						Class:  "sample",
						Labels: map[string]string{"other": "test", "label": "41"},
						Attrs:  map[string]string{},
						Values: [][]float64{{25000000.000000, 0.408248}, {85000000.000000, 0.408248}, {145000000.000000, 0.408248}, {205000000.000000, 0.408248}},
					},
					{
						Class:  "sample",
						Labels: map[string]string{"label": "42"},
						Attrs:  map[string]string{},
						Values: [][]float64{{25000000.000000, 15.261911}, {85000000.000000, 15.261911}, {145000000.000000, 15.261911}, {205000000.000000, 15.261911}},
					},
				},

				SeriesTests: seriesEqualityTestMap,
			},
			{
				SamplePrefix: BucketizeGtsList(BuildGtsList([]FloatGeoTimeSeries{gts1})),
				SampleSuffix: `
		<%
			DUP TOSELECTOR SWAP true MUSIGMA DROP TOSTRING + 
		%>
		SORTBY
		`,
				GTSResult: []FloatGeoTimeSeries{
					{
						Class:  "sample",
						Labels: map[string]string{"label": "42"},
						Attrs:  map[string]string{},
						Values: [][]float64{{25000000.000000, 15.261911}, {85000000.000000, 15.261911}, {145000000.000000, 15.261911}, {205000000.000000, 15.261911}},
					},
				},

				SeriesTests: seriesEqualityTestMap,
			},
		},
	},
	{
		Plan: []graphite.Function{
			{
				Name:       "aggregateLine",
				Arguments:  []string{"SWAP", "last"},
				Parameters: make(map[string]string),
			},
		},
		Samples: []OperatorGTSTest{
			{
				SamplePrefix: BucketizeGtsList(BuildGtsList([]FloatGeoTimeSeries{emptyGTS, gts1, gts1})),
				GTSResult: []FloatGeoTimeSeries{
					{
						Class:  "sample",
						Labels: map[string]string{"label": "42"},
						Attrs:  map[string]string{},
						Values: [][]float64{{205000000.000000, 32.666667}, {145000000.000000, 32.666667}, {85000000.000000, 32.666667}, {25000000.000000, 32.666667}},
					},
					{
						Class:  "sample",
						Labels: map[string]string{"label": "42"},
						Attrs:  map[string]string{},
						Values: [][]float64{{205000000.000000, 32.666667}, {145000000.000000, 32.666667}, {85000000.000000, 32.666667}, {25000000.000000, 32.666667}},
					},
				},

				SeriesTests: seriesEqualityTestMap,
			},
			{
				SamplePrefix: BucketizeGtsList(BuildGtsList([]FloatGeoTimeSeries{emptyGTS, gts1})),
				GTSResult: []FloatGeoTimeSeries{
					{
						Class:  "sample",
						Labels: map[string]string{"label": "42"},
						Attrs:  map[string]string{},
						Values: [][]float64{{205000000.000000, 32.666667}, {145000000.000000, 32.666667}, {85000000.000000, 32.666667}, {25000000.000000, 32.666667}},
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
						Values: [][]float64{{205000000.000000, 32.666667}, {145000000.000000, 32.666667}, {85000000.000000, 32.666667}, {25000000.000000, 32.666667}},
					},
					{
						Class:  "sample",
						Labels: map[string]string{"label": "42"},
						Attrs:  map[string]string{},
						Values: [][]float64{{205000000.000000, 32.666667}, {145000000.000000, 32.666667}, {85000000.000000, 32.666667}, {25000000.000000, 32.666667}},
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
						Values: [][]float64{{205000000.000000, 32.666667}, {145000000.000000, 32.666667}, {85000000.000000, 32.666667}, {25000000.000000, 32.666667}},
					},
					{
						Class:  "sample",
						Labels: map[string]string{"other": "test", "label": "41"},
						Attrs:  map[string]string{},
						Values: [][]float64{{205000000.000000, 9.000000}, {145000000.000000, 9.000000}, {85000000.000000, 9.000000}, {25000000.000000, 9.000000}},
					},
					{
						Class:  "sample",
						Labels: map[string]string{"other": "test", "label": "41"},
						Attrs:  map[string]string{},
						Values: [][]float64{{205000000.000000, 1.000000}, {145000000.000000, 1.000000}, {85000000.000000, 1.000000}, {25000000.000000, 1.000000}},
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
						Values: [][]float64{{205000000.000000, 32.666667}, {145000000.000000, 32.666667}, {85000000.000000, 32.666667}, {25000000.000000, 32.666667}},
					},
					{
						Class:  "sample",
						Labels: map[string]string{"other": "test", "label": "41"},
						Attrs:  map[string]string{},
						Values: [][]float64{{205000000.000000, 9.000000}, {145000000.000000, 9.000000}, {85000000.000000, 9.000000}, {25000000.000000, 9.000000}},
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
						Values: [][]float64{{205000000.000000, 32.666667}, {145000000.000000, 32.666667}, {85000000.000000, 32.666667}, {25000000.000000, 32.666667}},
					},
					{
						Class:  "sample",
						Labels: map[string]string{"other": "test", "label": "41"},
						Attrs:  map[string]string{},
						Values: [][]float64{{205000000.000000, 1.000000}, {145000000.000000, 1.000000}, {85000000.000000, 1.000000}, {25000000.000000, 1.000000}},
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
						Values: [][]float64{{205000000.000000, 32.666667}, {145000000.000000, 32.666667}, {85000000.000000, 32.666667}, {25000000.000000, 32.666667}},
					},
				},

				SeriesTests: seriesEqualityTestMap,
			},
		},
	},
}
