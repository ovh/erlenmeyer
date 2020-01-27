package prototests

import (
	"testing"

	"github.com/ovh/erlenmeyer/proto/graphite"
)

//
// @code auto-generated
// Test can be started with
// go test proto/prototests/graphite_invert_test.go proto/prototests/exec_test.go -v
//

// TestGraphiteInvert process Invert Graphite Unit tests
func TestGraphiteInvert(t *testing.T) {
	RunTest(t, graphiteGraphiteInvert, "")
}

var graphiteGraphiteInvert = []unitTests{

	{
		Plan: []graphite.Function{
			{
				Name:       "invert",
				Arguments:  []string{"SWAP"},
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
						Values: [][]float64{{25000000.000000, 1.000000}, {85000000.000000, -3.000000}, {145000000.000000, 0.100000}, {205000000.000000, 0.030612}},
					},
					{
						Class:  "sample",
						Labels: map[string]string{"label": "42"},
						Attrs:  map[string]string{},
						Values: [][]float64{{25000000.000000, 1.000000}, {85000000.000000, -3.000000}, {145000000.000000, 0.100000}, {205000000.000000, 0.030612}},
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
						Values: [][]float64{{25000000.000000, 1.000000}, {85000000.000000, -3.000000}, {145000000.000000, 0.100000}, {205000000.000000, 0.030612}},
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
						Values: [][]float64{{25000000.000000, 1.000000}, {85000000.000000, -3.000000}, {145000000.000000, 0.100000}, {205000000.000000, 0.030612}},
					},
					{
						Class:  "sample",
						Labels: map[string]string{"label": "42"},
						Attrs:  map[string]string{},
						Values: [][]float64{{25000000.000000, 1.000000}, {85000000.000000, -3.000000}, {145000000.000000, 0.100000}, {205000000.000000, 0.030612}},
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
						Labels: map[string]string{"label": "41", "other": "test"},
						Attrs:  map[string]string{},
						Values: [][]float64{{25000000.000000, 1.000000}, {85000000.000000, 0.500000}, {145000000.000000, 0.181818}, {205000000.000000, 0.111111}},
					},
					{
						Class:  "sample",
						Labels: map[string]string{"other": "test", "label": "41"},
						Attrs:  map[string]string{},
						Values: [][]float64{{25000000.000000, 1.000000}, {85000000.000000, 0.666667}, {145000000.000000, 2.000000}, {205000000.000000, 1.000000}},
					},
					{
						Class:  "sample",
						Labels: map[string]string{"label": "42"},
						Attrs:  map[string]string{},
						Values: [][]float64{{25000000.000000, 1.000000}, {85000000.000000, -3.000000}, {145000000.000000, 0.100000}, {205000000.000000, 0.030612}},
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
						Values: [][]float64{{25000000.000000, 1.000000}, {85000000.000000, 0.500000}, {145000000.000000, 0.181818}, {205000000.000000, 0.111111}},
					},
					{
						Class:  "sample",
						Labels: map[string]string{"label": "42"},
						Attrs:  map[string]string{},
						Values: [][]float64{{25000000.000000, 1.000000}, {85000000.000000, -3.000000}, {145000000.000000, 0.100000}, {205000000.000000, 0.030612}},
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
						Values: [][]float64{{25000000.000000, 1.000000}, {85000000.000000, 0.666667}, {145000000.000000, 2.000000}, {205000000.000000, 1.000000}},
					},
					{
						Class:  "sample",
						Labels: map[string]string{"label": "42"},
						Attrs:  map[string]string{},
						Values: [][]float64{{25000000.000000, 1.000000}, {85000000.000000, -3.000000}, {145000000.000000, 0.100000}, {205000000.000000, 0.030612}},
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
						Values: [][]float64{{25000000.000000, 1.000000}, {85000000.000000, -3.000000}, {145000000.000000, 0.100000}, {205000000.000000, 0.030612}},
					},
				},

				SeriesTests: seriesEqualityTestMap,
			},
			{
				SamplePrefix: BuildGtsList([]FloatGeoTimeSeries{emptyGTS, gts1, gts1}),
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
						Values: [][]float64{{0.000000, 1.000000}, {35000000.000000, -1.000000}, {60000000.000000, 0.500000}, {72000000.000000, -0.500000}, {88000000.000000, 0.250000}, {112000000.000000, 0.125000}, {122000000.000000, 0.083333}, {132000000.000000, 0.062500}, {162000000.000000, 0.055556}, {182000000.000000, -0.050000}, {202000000.000000, 0.010000}},
					},
					{
						Class:  "sample",
						Labels: map[string]string{"label": "42"},
						Attrs:  map[string]string{},
						Values: [][]float64{{0.000000, 1.000000}, {35000000.000000, -1.000000}, {60000000.000000, 0.500000}, {72000000.000000, -0.500000}, {88000000.000000, 0.250000}, {112000000.000000, 0.125000}, {122000000.000000, 0.083333}, {132000000.000000, 0.062500}, {162000000.000000, 0.055556}, {182000000.000000, -0.050000}, {202000000.000000, 0.010000}},
					},
				},

				SeriesTests: seriesEqualityTestMap,
			},
			{
				SamplePrefix: BuildGtsList([]FloatGeoTimeSeries{emptyGTS, gts1}),
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
						Values: [][]float64{{0.000000, 1.000000}, {35000000.000000, -1.000000}, {60000000.000000, 0.500000}, {72000000.000000, -0.500000}, {88000000.000000, 0.250000}, {112000000.000000, 0.125000}, {122000000.000000, 0.083333}, {132000000.000000, 0.062500}, {162000000.000000, 0.055556}, {182000000.000000, -0.050000}, {202000000.000000, 0.010000}},
					},
				},

				SeriesTests: seriesEqualityTestMap,
			},
			{
				SamplePrefix: BuildGtsList([]FloatGeoTimeSeries{gts1, gts1}),
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
						Values: [][]float64{{0.000000, 1.000000}, {35000000.000000, -1.000000}, {60000000.000000, 0.500000}, {72000000.000000, -0.500000}, {88000000.000000, 0.250000}, {112000000.000000, 0.125000}, {122000000.000000, 0.083333}, {132000000.000000, 0.062500}, {162000000.000000, 0.055556}, {182000000.000000, -0.050000}, {202000000.000000, 0.010000}},
					},
					{
						Class:  "sample",
						Labels: map[string]string{"label": "42"},
						Attrs:  map[string]string{},
						Values: [][]float64{{0.000000, 1.000000}, {35000000.000000, -1.000000}, {60000000.000000, 0.500000}, {72000000.000000, -0.500000}, {88000000.000000, 0.250000}, {112000000.000000, 0.125000}, {122000000.000000, 0.083333}, {132000000.000000, 0.062500}, {162000000.000000, 0.055556}, {182000000.000000, -0.050000}, {202000000.000000, 0.010000}},
					},
				},

				SeriesTests: seriesEqualityTestMap,
			},
			{
				SamplePrefix: BuildGtsList([]FloatGeoTimeSeries{gts1, gts2, gts3}),
				SampleSuffix: `
		<%
			DUP TOSELECTOR SWAP true MUSIGMA DROP TOSTRING + 
		%>
		SORTBY
		`,
				GTSResult: []FloatGeoTimeSeries{
					{
						Class:  "sample",
						Labels: map[string]string{"label": "41", "other": "test"},
						Attrs:  map[string]string{},
						Values: [][]float64{{0.000000, 1.000000}, {35000000.000000, 1.000000}, {60000000.000000, 0.500000}, {72000000.000000, 0.333333}, {88000000.000000, 0.250000}, {112000000.000000, 0.200000}, {122000000.000000, 0.166667}, {132000000.000000, 0.142857}, {162000000.000000, 0.125000}, {182000000.000000, 0.111111}, {202000000.000000, 0.100000}},
					},
					{
						Class:  "sample",
						Labels: map[string]string{"other": "test", "label": "41"},
						Attrs:  map[string]string{},
						Values: [][]float64{{0.000000, 1.000000}, {40000000.000000, 0.500000}, {82000000.000000, 1.000000}, {110000000.000000, 0.000000}, {129000000.000000, 1.000000}, {159000000.000000, 0.500000}, {192000000.000000, 0.000000}, {205000000.000000, 1.000000}},
					},
					{
						Class:  "sample",
						Labels: map[string]string{"label": "42"},
						Attrs:  map[string]string{},
						Values: [][]float64{{0.000000, 1.000000}, {35000000.000000, -1.000000}, {60000000.000000, 0.500000}, {72000000.000000, -0.500000}, {88000000.000000, 0.250000}, {112000000.000000, 0.125000}, {122000000.000000, 0.083333}, {132000000.000000, 0.062500}, {162000000.000000, 0.055556}, {182000000.000000, -0.050000}, {202000000.000000, 0.010000}},
					},
				},

				SeriesTests: seriesEqualityTestMap,
			},
			{
				SamplePrefix: BuildGtsList([]FloatGeoTimeSeries{gts1, gts2}),
				SampleSuffix: `
		<%
			DUP TOSELECTOR SWAP true MUSIGMA DROP TOSTRING + 
		%>
		SORTBY
		`,
				GTSResult: []FloatGeoTimeSeries{
					{
						Class:  "sample",
						Labels: map[string]string{"label": "41", "other": "test"},
						Attrs:  map[string]string{},
						Values: [][]float64{{0.000000, 1.000000}, {35000000.000000, 1.000000}, {60000000.000000, 0.500000}, {72000000.000000, 0.333333}, {88000000.000000, 0.250000}, {112000000.000000, 0.200000}, {122000000.000000, 0.166667}, {132000000.000000, 0.142857}, {162000000.000000, 0.125000}, {182000000.000000, 0.111111}, {202000000.000000, 0.100000}},
					},
					{
						Class:  "sample",
						Labels: map[string]string{"label": "42"},
						Attrs:  map[string]string{},
						Values: [][]float64{{0.000000, 1.000000}, {35000000.000000, -1.000000}, {60000000.000000, 0.500000}, {72000000.000000, -0.500000}, {88000000.000000, 0.250000}, {112000000.000000, 0.125000}, {122000000.000000, 0.083333}, {132000000.000000, 0.062500}, {162000000.000000, 0.055556}, {182000000.000000, -0.050000}, {202000000.000000, 0.010000}},
					},
				},

				SeriesTests: seriesEqualityTestMap,
			},
			{
				SamplePrefix: BuildGtsList([]FloatGeoTimeSeries{gts1, gts3}),
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
						Values: [][]float64{{0.000000, 1.000000}, {40000000.000000, 0.500000}, {82000000.000000, 1.000000}, {110000000.000000, 0.000000}, {129000000.000000, 1.000000}, {159000000.000000, 0.500000}, {192000000.000000, 0.000000}, {205000000.000000, 1.000000}},
					},
					{
						Class:  "sample",
						Labels: map[string]string{"label": "42"},
						Attrs:  map[string]string{},
						Values: [][]float64{{0.000000, 1.000000}, {35000000.000000, -1.000000}, {60000000.000000, 0.500000}, {72000000.000000, -0.500000}, {88000000.000000, 0.250000}, {112000000.000000, 0.125000}, {122000000.000000, 0.083333}, {132000000.000000, 0.062500}, {162000000.000000, 0.055556}, {182000000.000000, -0.050000}, {202000000.000000, 0.010000}},
					},
				},

				SeriesTests: seriesEqualityTestMap,
			},
			{
				SamplePrefix: BuildGtsList([]FloatGeoTimeSeries{gts1}),
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
						Values: [][]float64{{0.000000, 1.000000}, {35000000.000000, -1.000000}, {60000000.000000, 0.500000}, {72000000.000000, -0.500000}, {88000000.000000, 0.250000}, {112000000.000000, 0.125000}, {122000000.000000, 0.083333}, {132000000.000000, 0.062500}, {162000000.000000, 0.055556}, {182000000.000000, -0.050000}, {202000000.000000, 0.010000}},
					},
				},

				SeriesTests: seriesEqualityTestMap,
			},
		},
	},
	{
		Plan: []graphite.Function{
			{
				Name:       "invert",
				Arguments:  []string{"SWAP"},
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
						Values: [][]float64{{25000000.000000, 1.000000}, {85000000.000000, -3.000000}, {145000000.000000, 0.100000}, {205000000.000000, 0.030612}},
					},
					{
						Class:  "sample",
						Labels: map[string]string{"label": "42"},
						Attrs:  map[string]string{},
						Values: [][]float64{{25000000.000000, 1.000000}, {85000000.000000, -3.000000}, {145000000.000000, 0.100000}, {205000000.000000, 0.030612}},
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
						Values: [][]float64{{25000000.000000, 1.000000}, {85000000.000000, -3.000000}, {145000000.000000, 0.100000}, {205000000.000000, 0.030612}},
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
						Values: [][]float64{{25000000.000000, 1.000000}, {85000000.000000, -3.000000}, {145000000.000000, 0.100000}, {205000000.000000, 0.030612}},
					},
					{
						Class:  "sample",
						Labels: map[string]string{"label": "42"},
						Attrs:  map[string]string{},
						Values: [][]float64{{25000000.000000, 1.000000}, {85000000.000000, -3.000000}, {145000000.000000, 0.100000}, {205000000.000000, 0.030612}},
					},
				},

				SeriesTests: seriesEqualityTestMap,
			},
			{
				SamplePrefix: BucketizeGtsList(BuildGtsList([]FloatGeoTimeSeries{gts1, gts2, gts3})),
				GTSResult: []FloatGeoTimeSeries{
					{
						Class:  "sample",
						Labels: map[string]string{"other": "test", "label": "41"},
						Attrs:  map[string]string{},
						Values: [][]float64{{25000000.000000, 1.000000}, {85000000.000000, 0.666667}, {145000000.000000, 2.000000}, {205000000.000000, 1.000000}},
					},
					{
						Class:  "sample",
						Labels: map[string]string{"label": "42"},
						Attrs:  map[string]string{},
						Values: [][]float64{{25000000.000000, 1.000000}, {85000000.000000, -3.000000}, {145000000.000000, 0.100000}, {205000000.000000, 0.030612}},
					},
					{
						Class:  "sample",
						Labels: map[string]string{"other": "test", "label": "41"},
						Attrs:  map[string]string{},
						Values: [][]float64{{25000000.000000, 1.000000}, {85000000.000000, 0.500000}, {145000000.000000, 0.181818}, {205000000.000000, 0.111111}},
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
						Values: [][]float64{{25000000.000000, 1.000000}, {85000000.000000, -3.000000}, {145000000.000000, 0.100000}, {205000000.000000, 0.030612}},
					},
					{
						Class:  "sample",
						Labels: map[string]string{"other": "test", "label": "41"},
						Attrs:  map[string]string{},
						Values: [][]float64{{25000000.000000, 1.000000}, {85000000.000000, 0.500000}, {145000000.000000, 0.181818}, {205000000.000000, 0.111111}},
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
						Values: [][]float64{{25000000.000000, 1.000000}, {85000000.000000, -3.000000}, {145000000.000000, 0.100000}, {205000000.000000, 0.030612}},
					},
					{
						Class:  "sample",
						Labels: map[string]string{"other": "test", "label": "41"},
						Attrs:  map[string]string{},
						Values: [][]float64{{25000000.000000, 1.000000}, {85000000.000000, 0.666667}, {145000000.000000, 2.000000}, {205000000.000000, 1.000000}},
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
						Values: [][]float64{{25000000.000000, 1.000000}, {85000000.000000, -3.000000}, {145000000.000000, 0.100000}, {205000000.000000, 0.030612}},
					},
				},

				SeriesTests: seriesEqualityTestMap,
			},
			{
				SamplePrefix: BuildGtsList([]FloatGeoTimeSeries{emptyGTS, gts1, gts1}),
				GTSResult: []FloatGeoTimeSeries{
					{
						Class:  "sample",
						Labels: map[string]string{"label": "42"},
						Attrs:  map[string]string{},
						Values: [][]float64{{0.000000, 1.000000}, {35000000.000000, -1.000000}, {60000000.000000, 0.500000}, {72000000.000000, -0.500000}, {88000000.000000, 0.250000}, {112000000.000000, 0.125000}, {122000000.000000, 0.083333}, {132000000.000000, 0.062500}, {162000000.000000, 0.055556}, {182000000.000000, -0.050000}, {202000000.000000, 0.010000}},
					},
					{
						Class:  "sample",
						Labels: map[string]string{"label": "42"},
						Attrs:  map[string]string{},
						Values: [][]float64{{0.000000, 1.000000}, {35000000.000000, -1.000000}, {60000000.000000, 0.500000}, {72000000.000000, -0.500000}, {88000000.000000, 0.250000}, {112000000.000000, 0.125000}, {122000000.000000, 0.083333}, {132000000.000000, 0.062500}, {162000000.000000, 0.055556}, {182000000.000000, -0.050000}, {202000000.000000, 0.010000}},
					},
				},

				SeriesTests: seriesEqualityTestMap,
			},
			{
				SamplePrefix: BuildGtsList([]FloatGeoTimeSeries{emptyGTS, gts1}),
				GTSResult: []FloatGeoTimeSeries{
					{
						Class:  "sample",
						Labels: map[string]string{"label": "42"},
						Attrs:  map[string]string{},
						Values: [][]float64{{0.000000, 1.000000}, {35000000.000000, -1.000000}, {60000000.000000, 0.500000}, {72000000.000000, -0.500000}, {88000000.000000, 0.250000}, {112000000.000000, 0.125000}, {122000000.000000, 0.083333}, {132000000.000000, 0.062500}, {162000000.000000, 0.055556}, {182000000.000000, -0.050000}, {202000000.000000, 0.010000}},
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
						Values: [][]float64{{0.000000, 1.000000}, {35000000.000000, -1.000000}, {60000000.000000, 0.500000}, {72000000.000000, -0.500000}, {88000000.000000, 0.250000}, {112000000.000000, 0.125000}, {122000000.000000, 0.083333}, {132000000.000000, 0.062500}, {162000000.000000, 0.055556}, {182000000.000000, -0.050000}, {202000000.000000, 0.010000}},
					},
					{
						Class:  "sample",
						Labels: map[string]string{"label": "42"},
						Attrs:  map[string]string{},
						Values: [][]float64{{0.000000, 1.000000}, {35000000.000000, -1.000000}, {60000000.000000, 0.500000}, {72000000.000000, -0.500000}, {88000000.000000, 0.250000}, {112000000.000000, 0.125000}, {122000000.000000, 0.083333}, {132000000.000000, 0.062500}, {162000000.000000, 0.055556}, {182000000.000000, -0.050000}, {202000000.000000, 0.010000}},
					},
				},

				SeriesTests: seriesEqualityTestMap,
			},
			{
				SamplePrefix: BuildGtsList([]FloatGeoTimeSeries{gts1, gts2, gts3}),
				GTSResult: []FloatGeoTimeSeries{
					{
						Class:  "sample",
						Labels: map[string]string{"label": "41", "other": "test"},
						Attrs:  map[string]string{},
						Values: [][]float64{{0.000000, 1.000000}, {40000000.000000, 0.500000}, {82000000.000000, 1.000000}, {110000000.000000, 0.000000}, {129000000.000000, 1.000000}, {159000000.000000, 0.500000}, {192000000.000000, 0.000000}, {205000000.000000, 1.000000}},
					},
					{
						Class:  "sample",
						Labels: map[string]string{"label": "42"},
						Attrs:  map[string]string{},
						Values: [][]float64{{0.000000, 1.000000}, {35000000.000000, -1.000000}, {60000000.000000, 0.500000}, {72000000.000000, -0.500000}, {88000000.000000, 0.250000}, {112000000.000000, 0.125000}, {122000000.000000, 0.083333}, {132000000.000000, 0.062500}, {162000000.000000, 0.055556}, {182000000.000000, -0.050000}, {202000000.000000, 0.010000}},
					},
					{
						Class:  "sample",
						Labels: map[string]string{"other": "test", "label": "41"},
						Attrs:  map[string]string{},
						Values: [][]float64{{0.000000, 1.000000}, {35000000.000000, 1.000000}, {60000000.000000, 0.500000}, {72000000.000000, 0.333333}, {88000000.000000, 0.250000}, {112000000.000000, 0.200000}, {122000000.000000, 0.166667}, {132000000.000000, 0.142857}, {162000000.000000, 0.125000}, {182000000.000000, 0.111111}, {202000000.000000, 0.100000}},
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
						Values: [][]float64{{0.000000, 1.000000}, {35000000.000000, -1.000000}, {60000000.000000, 0.500000}, {72000000.000000, -0.500000}, {88000000.000000, 0.250000}, {112000000.000000, 0.125000}, {122000000.000000, 0.083333}, {132000000.000000, 0.062500}, {162000000.000000, 0.055556}, {182000000.000000, -0.050000}, {202000000.000000, 0.010000}},
					},
					{
						Class:  "sample",
						Labels: map[string]string{"other": "test", "label": "41"},
						Attrs:  map[string]string{},
						Values: [][]float64{{0.000000, 1.000000}, {35000000.000000, 1.000000}, {60000000.000000, 0.500000}, {72000000.000000, 0.333333}, {88000000.000000, 0.250000}, {112000000.000000, 0.200000}, {122000000.000000, 0.166667}, {132000000.000000, 0.142857}, {162000000.000000, 0.125000}, {182000000.000000, 0.111111}, {202000000.000000, 0.100000}},
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
						Values: [][]float64{{0.000000, 1.000000}, {35000000.000000, -1.000000}, {60000000.000000, 0.500000}, {72000000.000000, -0.500000}, {88000000.000000, 0.250000}, {112000000.000000, 0.125000}, {122000000.000000, 0.083333}, {132000000.000000, 0.062500}, {162000000.000000, 0.055556}, {182000000.000000, -0.050000}, {202000000.000000, 0.010000}},
					},
					{
						Class:  "sample",
						Labels: map[string]string{"other": "test", "label": "41"},
						Attrs:  map[string]string{},
						Values: [][]float64{{0.000000, 1.000000}, {40000000.000000, 0.500000}, {82000000.000000, 1.000000}, {110000000.000000, 0.000000}, {129000000.000000, 1.000000}, {159000000.000000, 0.500000}, {192000000.000000, 0.000000}, {205000000.000000, 1.000000}},
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
						Values: [][]float64{{0.000000, 1.000000}, {35000000.000000, -1.000000}, {60000000.000000, 0.500000}, {72000000.000000, -0.500000}, {88000000.000000, 0.250000}, {112000000.000000, 0.125000}, {122000000.000000, 0.083333}, {132000000.000000, 0.062500}, {162000000.000000, 0.055556}, {182000000.000000, -0.050000}, {202000000.000000, 0.010000}},
					},
				},

				SeriesTests: seriesEqualityTestMap,
			},
		},
	},
}
