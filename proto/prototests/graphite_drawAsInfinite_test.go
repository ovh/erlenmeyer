package prototests

import (
	"testing"

	"github.com/ovh/erlenmeyer/proto/graphite"
)

//
// @code auto-generated
// Test can be started with
// go test proto/prototests/graphite_drawAsInfinite_test.go proto/prototests/exec_test.go -v
//

// TestGraphiteDrawAsInfinite process Invert Graphite Unit tests
func TestGraphiteDrawAsInfinite(t *testing.T) {
	RunTest(t, graphiteGraphiteDrawAsInfinite, "")
}

var graphiteGraphiteDrawAsInfinite = []unitTests{

	{
		Plan: []graphite.Function{
			{
				Name:       "drawAsInfinite",
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
						Values: [][]float64{{25000000.000000, 9223372036854775808.000000}, {85000000.000000, 9223372036854775808.000000}, {145000000.000000, 9223372036854775808.000000}, {205000000.000000, 9223372036854775808.000000}},
					},
					{
						Class:  "sample",
						Labels: map[string]string{"label": "42"},
						Attrs:  map[string]string{},
						Values: [][]float64{{25000000.000000, 9223372036854775808.000000}, {85000000.000000, 9223372036854775808.000000}, {145000000.000000, 9223372036854775808.000000}, {205000000.000000, 9223372036854775808.000000}},
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
						Values: [][]float64{{25000000.000000, 9223372036854775808.000000}, {85000000.000000, 9223372036854775808.000000}, {145000000.000000, 9223372036854775808.000000}, {205000000.000000, 9223372036854775808.000000}},
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
						Values: [][]float64{{25000000.000000, 9223372036854775808.000000}, {85000000.000000, 9223372036854775808.000000}, {145000000.000000, 9223372036854775808.000000}, {205000000.000000, 9223372036854775808.000000}},
					},
					{
						Class:  "sample",
						Labels: map[string]string{"label": "42"},
						Attrs:  map[string]string{},
						Values: [][]float64{{25000000.000000, 9223372036854775808.000000}, {85000000.000000, 9223372036854775808.000000}, {145000000.000000, 9223372036854775808.000000}, {205000000.000000, 9223372036854775808.000000}},
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
						Values: [][]float64{{25000000.000000, 9223372036854775808.000000}, {85000000.000000, 9223372036854775808.000000}, {145000000.000000, 9223372036854775808.000000}, {205000000.000000, 9223372036854775808.000000}},
					},
					{
						Class:  "sample",
						Labels: map[string]string{"label": "41", "other": "test"},
						Attrs:  map[string]string{},
						Values: [][]float64{{25000000.000000, 9223372036854775808.000000}, {85000000.000000, 9223372036854775808.000000}, {145000000.000000, 9223372036854775808.000000}, {205000000.000000, 9223372036854775808.000000}},
					},
					{
						Class:  "sample",
						Labels: map[string]string{"label": "42"},
						Attrs:  map[string]string{},
						Values: [][]float64{{25000000.000000, 9223372036854775808.000000}, {85000000.000000, 9223372036854775808.000000}, {145000000.000000, 9223372036854775808.000000}, {205000000.000000, 9223372036854775808.000000}},
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
						Labels: map[string]string{"label": "41", "other": "test"},
						Attrs:  map[string]string{},
						Values: [][]float64{{25000000.000000, 9223372036854775808.000000}, {85000000.000000, 9223372036854775808.000000}, {145000000.000000, 9223372036854775808.000000}, {205000000.000000, 9223372036854775808.000000}},
					},
					{
						Class:  "sample",
						Labels: map[string]string{"label": "42"},
						Attrs:  map[string]string{},
						Values: [][]float64{{25000000.000000, 9223372036854775808.000000}, {85000000.000000, 9223372036854775808.000000}, {145000000.000000, 9223372036854775808.000000}, {205000000.000000, 9223372036854775808.000000}},
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
						Values: [][]float64{{25000000.000000, 9223372036854775808.000000}, {85000000.000000, 9223372036854775808.000000}, {145000000.000000, 9223372036854775808.000000}, {205000000.000000, 9223372036854775808.000000}},
					},
					{
						Class:  "sample",
						Labels: map[string]string{"label": "42"},
						Attrs:  map[string]string{},
						Values: [][]float64{{25000000.000000, 9223372036854775808.000000}, {85000000.000000, 9223372036854775808.000000}, {145000000.000000, 9223372036854775808.000000}, {205000000.000000, 9223372036854775808.000000}},
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
						Values: [][]float64{{25000000.000000, 9223372036854775808.000000}, {85000000.000000, 9223372036854775808.000000}, {145000000.000000, 9223372036854775808.000000}, {205000000.000000, 9223372036854775808.000000}},
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
						Values: [][]float64{{0.000000, 9223372036854775808.000000}, {35000000.000000, 9223372036854775808.000000}, {60000000.000000, 9223372036854775808.000000}, {72000000.000000, 9223372036854775808.000000}, {88000000.000000, 9223372036854775808.000000}, {112000000.000000, 9223372036854775808.000000}, {122000000.000000, 9223372036854775808.000000}, {132000000.000000, 9223372036854775808.000000}, {162000000.000000, 9223372036854775808.000000}, {182000000.000000, 9223372036854775808.000000}, {202000000.000000, 9223372036854775808.000000}},
					},
					{
						Class:  "sample",
						Labels: map[string]string{"label": "42"},
						Attrs:  map[string]string{},
						Values: [][]float64{{0.000000, 9223372036854775808.000000}, {35000000.000000, 9223372036854775808.000000}, {60000000.000000, 9223372036854775808.000000}, {72000000.000000, 9223372036854775808.000000}, {88000000.000000, 9223372036854775808.000000}, {112000000.000000, 9223372036854775808.000000}, {122000000.000000, 9223372036854775808.000000}, {132000000.000000, 9223372036854775808.000000}, {162000000.000000, 9223372036854775808.000000}, {182000000.000000, 9223372036854775808.000000}, {202000000.000000, 9223372036854775808.000000}},
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
						Values: [][]float64{{0.000000, 9223372036854775808.000000}, {35000000.000000, 9223372036854775808.000000}, {60000000.000000, 9223372036854775808.000000}, {72000000.000000, 9223372036854775808.000000}, {88000000.000000, 9223372036854775808.000000}, {112000000.000000, 9223372036854775808.000000}, {122000000.000000, 9223372036854775808.000000}, {132000000.000000, 9223372036854775808.000000}, {162000000.000000, 9223372036854775808.000000}, {182000000.000000, 9223372036854775808.000000}, {202000000.000000, 9223372036854775808.000000}},
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
						Values: [][]float64{{0.000000, 9223372036854775808.000000}, {35000000.000000, 9223372036854775808.000000}, {60000000.000000, 9223372036854775808.000000}, {72000000.000000, 9223372036854775808.000000}, {88000000.000000, 9223372036854775808.000000}, {112000000.000000, 9223372036854775808.000000}, {122000000.000000, 9223372036854775808.000000}, {132000000.000000, 9223372036854775808.000000}, {162000000.000000, 9223372036854775808.000000}, {182000000.000000, 9223372036854775808.000000}, {202000000.000000, 9223372036854775808.000000}},
					},
					{
						Class:  "sample",
						Labels: map[string]string{"label": "42"},
						Attrs:  map[string]string{},
						Values: [][]float64{{0.000000, 9223372036854775808.000000}, {35000000.000000, 9223372036854775808.000000}, {60000000.000000, 9223372036854775808.000000}, {72000000.000000, 9223372036854775808.000000}, {88000000.000000, 9223372036854775808.000000}, {112000000.000000, 9223372036854775808.000000}, {122000000.000000, 9223372036854775808.000000}, {132000000.000000, 9223372036854775808.000000}, {162000000.000000, 9223372036854775808.000000}, {182000000.000000, 9223372036854775808.000000}, {202000000.000000, 9223372036854775808.000000}},
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
						Labels: map[string]string{"other": "test", "label": "41"},
						Attrs:  map[string]string{},
						Values: [][]float64{{0.000000, 9223372036854775808.000000}, {40000000.000000, 9223372036854775808.000000}, {82000000.000000, 9223372036854775808.000000}, {110000000.000000, 0.000000}, {129000000.000000, 9223372036854775808.000000}, {159000000.000000, 9223372036854775808.000000}, {192000000.000000, 0.000000}, {205000000.000000, 9223372036854775808.000000}},
					},
					{
						Class:  "sample",
						Labels: map[string]string{"other": "test", "label": "41"},
						Attrs:  map[string]string{},
						Values: [][]float64{{0.000000, 9223372036854775808.000000}, {35000000.000000, 9223372036854775808.000000}, {60000000.000000, 9223372036854775808.000000}, {72000000.000000, 9223372036854775808.000000}, {88000000.000000, 9223372036854775808.000000}, {112000000.000000, 9223372036854775808.000000}, {122000000.000000, 9223372036854775808.000000}, {132000000.000000, 9223372036854775808.000000}, {162000000.000000, 9223372036854775808.000000}, {182000000.000000, 9223372036854775808.000000}, {202000000.000000, 9223372036854775808.000000}},
					},
					{
						Class:  "sample",
						Labels: map[string]string{"label": "42"},
						Attrs:  map[string]string{},
						Values: [][]float64{{0.000000, 9223372036854775808.000000}, {35000000.000000, 9223372036854775808.000000}, {60000000.000000, 9223372036854775808.000000}, {72000000.000000, 9223372036854775808.000000}, {88000000.000000, 9223372036854775808.000000}, {112000000.000000, 9223372036854775808.000000}, {122000000.000000, 9223372036854775808.000000}, {132000000.000000, 9223372036854775808.000000}, {162000000.000000, 9223372036854775808.000000}, {182000000.000000, 9223372036854775808.000000}, {202000000.000000, 9223372036854775808.000000}},
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
						Labels: map[string]string{"other": "test", "label": "41"},
						Attrs:  map[string]string{},
						Values: [][]float64{{0.000000, 9223372036854775808.000000}, {35000000.000000, 9223372036854775808.000000}, {60000000.000000, 9223372036854775808.000000}, {72000000.000000, 9223372036854775808.000000}, {88000000.000000, 9223372036854775808.000000}, {112000000.000000, 9223372036854775808.000000}, {122000000.000000, 9223372036854775808.000000}, {132000000.000000, 9223372036854775808.000000}, {162000000.000000, 9223372036854775808.000000}, {182000000.000000, 9223372036854775808.000000}, {202000000.000000, 9223372036854775808.000000}},
					},
					{
						Class:  "sample",
						Labels: map[string]string{"label": "42"},
						Attrs:  map[string]string{},
						Values: [][]float64{{0.000000, 9223372036854775808.000000}, {35000000.000000, 9223372036854775808.000000}, {60000000.000000, 9223372036854775808.000000}, {72000000.000000, 9223372036854775808.000000}, {88000000.000000, 9223372036854775808.000000}, {112000000.000000, 9223372036854775808.000000}, {122000000.000000, 9223372036854775808.000000}, {132000000.000000, 9223372036854775808.000000}, {162000000.000000, 9223372036854775808.000000}, {182000000.000000, 9223372036854775808.000000}, {202000000.000000, 9223372036854775808.000000}},
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
						Values: [][]float64{{0.000000, 9223372036854775808.000000}, {40000000.000000, 9223372036854775808.000000}, {82000000.000000, 9223372036854775808.000000}, {110000000.000000, 0.000000}, {129000000.000000, 9223372036854775808.000000}, {159000000.000000, 9223372036854775808.000000}, {192000000.000000, 0.000000}, {205000000.000000, 9223372036854775808.000000}},
					},
					{
						Class:  "sample",
						Labels: map[string]string{"label": "42"},
						Attrs:  map[string]string{},
						Values: [][]float64{{0.000000, 9223372036854775808.000000}, {35000000.000000, 9223372036854775808.000000}, {60000000.000000, 9223372036854775808.000000}, {72000000.000000, 9223372036854775808.000000}, {88000000.000000, 9223372036854775808.000000}, {112000000.000000, 9223372036854775808.000000}, {122000000.000000, 9223372036854775808.000000}, {132000000.000000, 9223372036854775808.000000}, {162000000.000000, 9223372036854775808.000000}, {182000000.000000, 9223372036854775808.000000}, {202000000.000000, 9223372036854775808.000000}},
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
						Values: [][]float64{{0.000000, 9223372036854775808.000000}, {35000000.000000, 9223372036854775808.000000}, {60000000.000000, 9223372036854775808.000000}, {72000000.000000, 9223372036854775808.000000}, {88000000.000000, 9223372036854775808.000000}, {112000000.000000, 9223372036854775808.000000}, {122000000.000000, 9223372036854775808.000000}, {132000000.000000, 9223372036854775808.000000}, {162000000.000000, 9223372036854775808.000000}, {182000000.000000, 9223372036854775808.000000}, {202000000.000000, 9223372036854775808.000000}},
					},
				},

				SeriesTests: seriesEqualityTestMap,
			},
		},
	},
}
