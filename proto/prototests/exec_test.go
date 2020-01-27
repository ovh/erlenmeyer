package prototests

//
// Test can be started with
// go test proto/prototests/exec_test.go  proto/prototests/graphite_test.go  -v
// Will execute the test on a warp 10 instancte started at WARP_TEST_ENDPOINT or "http://127.0.0.1:8090/api/v0/exec"
//

import (
	"bytes"
	"crypto/sha256"
	"flag"
	"fmt"
	"math"
	"os"
	"sort"
	"strings"
	"testing"

	"github.com/ovh/erlenmeyer/core"
	"github.com/ovh/erlenmeyer/proto/graphite"
)

// unitTests main structuctre
type unitTests struct {
	Plan           []graphite.Function
	Samples        []OperatorGTSTest
	generate       bool
	generateSample map[string]string
	generateSuffix string
	generateChecks string
	PlanID         string
}

// singleTestOperator
type OperatorGTSTest struct {
	SamplePrefix string
	SampleSuffix string
	GTSResult    []FloatGeoTimeSeries
	SeriesTests  map[WarpTestTokens]WarpTest
}

// WarpTestTokens is an enum for test to apply
type WarpTestTokens int

const (
	SeriesListSizeTest WarpTestTokens = iota
	SeriesEqualityTest
	SeriesEqualitySkipValuesTest
)

func (wt WarpTestTokens) String() string {
	typeToStr := map[WarpTestTokens]string{
		SeriesListSizeTest:           "checkSeriesListSize",
		SeriesEqualityTest:           "checkEquality",
		SeriesEqualitySkipValuesTest: "checkEqualitySkipValues",
	}
	if str, ok := typeToStr[wt]; ok {
		return str
	}
	return ""
}

type WarpTest interface {
	TestType(wt WarpTestTokens) string
}

// SimpleTest is the default test case with no params
type SimpleTest struct {
}

//TestType return type of current Test
func (r SimpleTest) TestType(wt WarpTestTokens) string {
	return wt.String()
}

// FloatGeoTimeSeries is a GTS with float values
type FloatGeoTimeSeries struct {
	Class  string            `json:"c"`
	Labels map[string]string `json:"l"`
	Attrs  map[string]string `json:"a"`
	Values [][]float64       `json:"v"`
}

var (
	seriesEqualityTestMap = map[WarpTestTokens]WarpTest{
		SeriesListSizeTest: SimpleTest{},
		SeriesEqualityTest: SimpleTest{},
	}

	seriesEqualitySkipValuesTestMap = map[WarpTestTokens]WarpTest{
		SeriesListSizeTest:           SimpleTest{},
		SeriesEqualitySkipValuesTest: SimpleTest{},
	}

	redefFetch = `
	'' 'token' STORE
<% 
2 GET
'tags' STORE
[
    NEWGTS 'sample'  RENAME
    $tags RELABEL
    0.000000 NaN NaN NaN 1.000000 ADDVALUE
    35000000.000000 NaN NaN NaN -1.000000 ADDVALUE
    60000000.000000 NaN NaN NaN 2.000000 ADDVALUE
    72000000.000000 NaN NaN NaN -2.000000 ADDVALUE
    88000000.000000 NaN NaN NaN 4.000000 ADDVALUE
    112000000.000000 NaN NaN NaN 8.000000 ADDVALUE
    122000000.000000 NaN NaN NaN 12.000000 ADDVALUE
    132000000.000000 NaN NaN NaN 16.000000 ADDVALUE
    162000000.000000 NaN NaN NaN 18.000000 ADDVALUE
    182000000.000000 NaN NaN NaN -20.000000 ADDVALUE
    202000000.000000 NaN NaN NaN 100.000000 ADDVALUE
    
]
%> 'FETCH' DEF
	`

	emptySample       = map[string]string{"\"CLEAR\"": "CLEAR"}
	divideSeriesFetch = map[string]string{
		"BuildGtsList([]FloatGeoTimeSeries{gts1}) + " + "`" + redefFetch + "`":     BuildGtsList([]FloatGeoTimeSeries{gts1}) + redefFetch,
		"BuildGtsList([]FloatGeoTimeSeries{emptyGTS}) + " + "`" + redefFetch + "`": BuildGtsList([]FloatGeoTimeSeries{emptyGTS}) + redefFetch,
		"BuildGtsList([]FloatGeoTimeSeries{emptyGTS, gts1}) + " + "`" + redefFetch + "`": BuildGtsList([]FloatGeoTimeSeries{
			emptyGTS, gts1}) + redefFetch,
		"BuildGtsList([]FloatGeoTimeSeries{emptyGTS, gts1, gts1}) + " + "`" + redefFetch + "`": BuildGtsList([]FloatGeoTimeSeries{
			emptyGTS, gts1, gts1}) + redefFetch,
		"BuildGtsList([]FloatGeoTimeSeries{gts1, gts1}) + " + "`" + redefFetch + "`": BuildGtsList([]FloatGeoTimeSeries{
			gts1, gts1}) + redefFetch,
		"BuildGtsList([]FloatGeoTimeSeries{gts1, gts2}) + " + "`" + redefFetch + "`": BuildGtsList([]FloatGeoTimeSeries{
			gts1, gts2}) + redefFetch,
		"BuildGtsList([]FloatGeoTimeSeries{gts1, gts3}) + " + "`" + redefFetch + "`": BuildGtsList([]FloatGeoTimeSeries{
			gts1, gts3}) + redefFetch,
		"BuildGtsList([]FloatGeoTimeSeries{gts1, gts2, gts3}) + " + "`" + redefFetch + "`": BuildGtsList([]FloatGeoTimeSeries{
			gts1, gts2, gts3}) + redefFetch,
		"BucketizeGtsList(BuildGtsList([]FloatGeoTimeSeries{gts1})) + " + "`" + redefFetch + "`":     BucketizeGtsList(BuildGtsList([]FloatGeoTimeSeries{gts1})),
		"BucketizeGtsList(BuildGtsList([]FloatGeoTimeSeries{emptyGTS})) + " + "`" + redefFetch + "`": BucketizeGtsList(BuildGtsList([]FloatGeoTimeSeries{emptyGTS})),
		"BucketizeGtsList(BuildGtsList([]FloatGeoTimeSeries{emptyGTS, gts1})) + " + "`" + redefFetch + "`": BucketizeGtsList(BuildGtsList([]FloatGeoTimeSeries{
			emptyGTS, gts1}) + redefFetch),
		"BucketizeGtsList(BuildGtsList([]FloatGeoTimeSeries{emptyGTS, gts1, gts1})) + " + "`" + redefFetch + "`": BucketizeGtsList(BuildGtsList([]FloatGeoTimeSeries{
			emptyGTS, gts1, gts1}) + redefFetch),
		"BucketizeGtsList(BuildGtsList([]FloatGeoTimeSeries{gts1, gts1})) + " + "`" + redefFetch + "`": BucketizeGtsList(BuildGtsList([]FloatGeoTimeSeries{
			gts1, gts1}) + redefFetch),
		"BucketizeGtsList(BuildGtsList([]FloatGeoTimeSeries{gts1, gts2})) + " + "`" + redefFetch + "`": BucketizeGtsList(BuildGtsList([]FloatGeoTimeSeries{
			gts1, gts2}) + redefFetch),
		"BucketizeGtsList(BuildGtsList([]FloatGeoTimeSeries{gts1, gts3})) + " + "`" + redefFetch + "`": BucketizeGtsList(BuildGtsList([]FloatGeoTimeSeries{
			gts1, gts3}) + redefFetch),
		"BucketizeGtsList(BuildGtsList([]FloatGeoTimeSeries{gts1, gts2, gts3}))+ " + "`" + redefFetch + "`": BucketizeGtsList(BuildGtsList([]FloatGeoTimeSeries{
			gts1, gts2, gts3}) + redefFetch),
	}
	redefinedFetch = map[string]string{"`" + redefFetch + "`": redefFetch}

	generateSample = map[string]string{
		"BuildGtsList([]FloatGeoTimeSeries{gts1})":     BuildGtsList([]FloatGeoTimeSeries{gts1}),
		"BuildGtsList([]FloatGeoTimeSeries{emptyGTS})": BuildGtsList([]FloatGeoTimeSeries{emptyGTS}),
		"BuildGtsList([]FloatGeoTimeSeries{emptyGTS, gts1})": BuildGtsList([]FloatGeoTimeSeries{
			emptyGTS, gts1}),
		"BuildGtsList([]FloatGeoTimeSeries{emptyGTS, gts1, gts1})": BuildGtsList([]FloatGeoTimeSeries{
			emptyGTS, gts1, gts1}),
		"BuildGtsList([]FloatGeoTimeSeries{gts1, gts1})": BuildGtsList([]FloatGeoTimeSeries{
			gts1, gts1}),
		"BuildGtsList([]FloatGeoTimeSeries{gts1, gts2})": BuildGtsList([]FloatGeoTimeSeries{
			gts1, gts2}),
		"BuildGtsList([]FloatGeoTimeSeries{gts1, gts3})": BuildGtsList([]FloatGeoTimeSeries{
			gts1, gts3}),
		"BuildGtsList([]FloatGeoTimeSeries{gts1, gts2, gts3})": BuildGtsList([]FloatGeoTimeSeries{
			gts1, gts2, gts3}),
		"BucketizeGtsList(BuildGtsList([]FloatGeoTimeSeries{gts1}))":     BucketizeGtsList(BuildGtsList([]FloatGeoTimeSeries{gts1})),
		"BucketizeGtsList(BuildGtsList([]FloatGeoTimeSeries{emptyGTS}))": BucketizeGtsList(BuildGtsList([]FloatGeoTimeSeries{emptyGTS})),
		"BucketizeGtsList(BuildGtsList([]FloatGeoTimeSeries{emptyGTS, gts1}))": BucketizeGtsList(BuildGtsList([]FloatGeoTimeSeries{
			emptyGTS, gts1})),
		"BucketizeGtsList(BuildGtsList([]FloatGeoTimeSeries{emptyGTS, gts1, gts1}))": BucketizeGtsList(BuildGtsList([]FloatGeoTimeSeries{
			emptyGTS, gts1, gts1})),
		"BucketizeGtsList(BuildGtsList([]FloatGeoTimeSeries{gts1, gts1}))": BucketizeGtsList(BuildGtsList([]FloatGeoTimeSeries{
			gts1, gts1})),
		"BucketizeGtsList(BuildGtsList([]FloatGeoTimeSeries{gts1, gts2}))": BucketizeGtsList(BuildGtsList([]FloatGeoTimeSeries{
			gts1, gts2})),
		"BucketizeGtsList(BuildGtsList([]FloatGeoTimeSeries{gts1, gts3}))": BucketizeGtsList(BuildGtsList([]FloatGeoTimeSeries{
			gts1, gts3})),
		"BucketizeGtsList(BuildGtsList([]FloatGeoTimeSeries{gts1, gts2, gts3}))": BucketizeGtsList(BuildGtsList([]FloatGeoTimeSeries{
			gts1, gts2, gts3})),
	}
)

const (
	testSwap     = "SWAP"
	generateFile = "generate"
)

// TestUnit process Graphite Unit tests
func RunTest(t *testing.T, unitsTests []unitTests, doGenerate string) {
	flag.Parse()

	warpTestEndpoint := os.Getenv("WARP_TEST_ENDPOINT")
	if warpTestEndpoint == "" {
		warpTestEndpoint = "http://127.0.0.1:8090/api/v0/exec"
	}

	// Start a Warp 10 server
	server := core.NewWarpServer(warpTestEndpoint, "test")
	txn := fmt.Sprintf("%x", sha256.New().Sum(nil))

	seenFiles := make((map[string]bytes.Buffer), 0)

	proto := "graphite"

	// Execute all tests for each implemented methods
	for _, test := range unitsTests {

		id := test.PlanID

		root := core.NewEmptyNode()

		node := root

		idSeparator := "+ "

		for functionIndex := range test.Plan {
			function := test.Plan[len(test.Plan)-1-functionIndex]
			fn, err := graphite.GetFunction(function.Name)

			if test.PlanID == "" {
				if functionIndex != 0 {
					id += idSeparator
				}
				id += function.Name
			}

			if err != nil {
				t.Error(err)
				continue
			}

			_, err = fn(node, function.Arguments, function.Parameters)
			if err != nil {
				t.Errorf("%s: %s", function.Name, err.Error())
				continue
			}
			node = node.Left
		}

		filename := strings.Replace(id, idSeparator, "", -1) + "_test.go"

		var fileBuffer bytes.Buffer

		if _, ok := seenFiles[filename]; ok {
			fileBuffer = seenFiles[filename]
		}

		var buffer bytes.Buffer

		// Generate query string for current operator isolated
		ws := root.InternalToWarpScript("query")

		stackLength := 0
		for _, sample := range test.Samples {

			// Add specified query prefix
			buffer.WriteString(sample.SamplePrefix)
			buffer.WriteString("\n")
			buffer.WriteString(ws)
			buffer.WriteString("\n")
			buffer.WriteString(sample.SampleSuffix)

			// Build expected stack length
			if len(sample.GTSResult) > 0 {
				stackLength++
			}
		}

		if test.generate {

			if len(test.generateSample) == 0 {
				test.generateSample = generateSample
			}

			var keys []string
			for k := range test.generateSample {
				keys = append(keys, k)
			}
			sort.Strings(keys)
			for _, key := range keys {
				inputSeries := test.generateSample[key]
				buffer.WriteString(inputSeries)
				buffer.WriteString("\n")
				buffer.WriteString(ws)
				buffer.WriteString("\n")
				buffer.WriteString(test.generateSuffix)
			}

			//t.Logf("%s", buffer.String())
		} else {
			//t.Logf("%s", buffer.String())
		}

		// Execute the query on test backend
		gtsss, err := server.QueryGTSs(buffer.String(), txn)
		if err != nil {
			t.Errorf("Function %s: fail to execute WarpScript", id)
		}

		if test.generate && len(gtsss) > 0 {
			fileBuffer.WriteString(generateUnitTest(gtsss, test, t, id))
			seenFiles[filename] = fileBuffer
			continue
		}

		if len(gtsss) != stackLength {
			t.Errorf("Function %s: WarpScript result expected a stack length of %d, got %d", id, stackLength, len(gtsss))
		}

		// Check each stack level individually
		for stackIndex, gtss := range gtsss {

			sample := test.Samples[len(gtsss)-1-stackIndex]

			_, checkSeriesListLength := sample.SeriesTests[SeriesListSizeTest]
			if checkSeriesListLength && len(gtss) != len(sample.GTSResult) {
				t.Errorf("Function %s: WarpScript result expected a GTS list of %d, got %d", id, len(gtss), len(sample.GTSResult))
			}

			for index, gts := range gtss {
				testGts := sample.GTSResult[index]
				_, checkSeriesEquality := sample.SeriesTests[SeriesEqualityTest]

				_, checkSeriesEqualitySkipValues := sample.SeriesTests[SeriesEqualitySkipValuesTest]

				if checkSeriesEquality {
					CheckEquality(gts, testGts, t, index, id, false)
				}
				if checkSeriesEqualitySkipValues {
					CheckEquality(gts, testGts, t, index, id, true)
				}
			}
		}
		if !test.generate {
			t.Logf("%s all tests are completed", id)
		}
	}

	if len(seenFiles) > 0 {
		for filename, buffer := range seenFiles {
			writeGoFile(filename, buffer.String(), t, doGenerate, proto)
		}
	}
}

func writeGoFile(name string, buffer string, t *testing.T, doGenerate string, proto string) {
	file := &os.File{}

	filename := proto + "_" + name

	if doGenerate == generateFile {
		var err error

		file, err = os.Create(filename)
		if err != nil {
			t.Errorf("Couldn't create file %s ", filename)
		}
	}

	defer file.Close()

	if filename != "" {
		b, err := file.WriteString(getFileGeneratedPrefix(name, proto))
		if err != nil {
			t.Errorf("Function %s: Could not generate file", filename)
		} else {
			t.Logf("%d", b)
			file.Sync()
		}
		b, err = file.WriteString(buffer)
		if err != nil {
			t.Errorf("Function %s: Could not generate file", filename)
		} else {
			t.Logf("%d", b)
			file.Sync()
		}
		b, err = file.WriteString("\n}")
		if err != nil {
			t.Errorf("Function %s: Could not generate file", filename)
		} else {
			t.Logf("%d", b)
			file.Sync()
		}
	}
}

// Generate a Test for current function and push it in t.logs
// Care SeriesTests use seriesEqualityTestMap local variable
func generateUnitTest(stack [][]core.GeoTimeSeries, ut unitTests, t *testing.T, id string) string {
	var buffer bytes.Buffer

	buffer.WriteString("\n{ \n")
	buffer.WriteString("	Plan: " + getPlanAsGoString(ut.Plan))

	buffer.WriteString("	Samples: []OperatorGTSTest{\n")

	index := 0

	var keys []string
	for k := range ut.generateSample {
		keys = append(keys, k)
	}
	sort.Strings(keys)

	genChecks := "seriesEqualityTestMap"

	if ut.generateChecks != "" {
		genChecks = ut.generateChecks
	}

	for _, key := range keys {
		seriesList := stack[len(ut.generateSample)-1-index]

		if len(seriesList) != 0 {
			buffer.WriteString("	OperatorGTSTest{\n")
			buffer.WriteString("		SamplePrefix: " + key + ", \n")

			if ut.generateSuffix != "" {
				buffer.WriteString("		SampleSuffix: `" + ut.generateSuffix + "`, \n")
			}
			buffer.WriteString("		GTSResult: " + getCoreAsFloatGTSString(seriesList) + " \n")
			buffer.WriteString("		SeriesTests: " + genChecks + ",\n")
			buffer.WriteString("	},\n")
		} else {
			t.Logf("Function %s: skip %s key in generation: empty result", id, key)
		}
		index++
	}
	buffer.WriteString("},\n")
	buffer.WriteString("},")

	return buffer.String()
}

func getFileGeneratedPrefix(filename string, proto string) string {
	funName := strings.Title(proto) + strings.Title(strings.Replace(filename, "_test.go", "", -1))
	return `
package prototests

import (
	"testing"

	"github.com/ovh/erlenmeyer/proto/graphite"
)

//
// @code auto-generated
// Test can be started with
// go test proto/prototests/` + proto + "_" + filename + ` proto/prototests/exec_test.go -v
//

// Test` + funName + ` process Invert Graphite Unit tests
func Test` + funName + `(t *testing.T) {
	RunTest(t, graphite` + funName + `, "")
}

var graphite` + funName + ` = []unitTests{
	`
}

func getCoreAsFloatGTSString(seriesList []core.GeoTimeSeries) string {
	var buffer bytes.Buffer

	buffer.WriteString("[]FloatGeoTimeSeries{\n")

	for _, gts := range seriesList {

		buffer.WriteString("		FloatGeoTimeSeries{\n")
		buffer.WriteString("			Class: \"" + gts.Class + "\",\n")
		buffer.WriteString("			Labels: " + getFunctionParametersAsString(gts.Labels) + ",\n")
		buffer.WriteString("			Attrs: " + getFunctionParametersAsString(gts.Attrs) + ",\n")
		buffer.WriteString("			Values: [][]float64{")
		for index, v := range gts.Values {
			if index != 0 {
				buffer.WriteString(",")
			}
			buffer.WriteString(fmt.Sprintf("[]float64{ %f, %f}", v[0], v[1]))
		}
		buffer.WriteString("},\n")
		buffer.WriteString("			},\n")
	}
	buffer.WriteString("		},\n")

	return buffer.String()
}

func getPlanAsGoString(plan []graphite.Function) string {
	var buffer bytes.Buffer

	buffer.WriteString("	[]graphite.Function{\n")

	for _, function := range plan {
		buffer.WriteString("		graphite.Function{\n")
		buffer.WriteString("			Name: \"" + function.Name + "\", \n")
		buffer.WriteString("			Arguments:  []string{" + getFunctionArgumentsAsString(function.Arguments) + ", \n")
		if len(function.Parameters) == 0 {
			buffer.WriteString("			Parameters:  make(map[string]string), \n")
		} else {
			buffer.WriteString("			Parameters: " + getFunctionParametersAsString(function.Parameters) + ", \n")
		}
		buffer.WriteString("		},\n")
	}

	buffer.WriteString("	},\n")

	return buffer.String()
}

func getFunctionArgumentsAsString(arguments []string) string {
	var buffer bytes.Buffer
	for index, argument := range arguments {
		if index != 0 {
			buffer.WriteString(",")
		}
		buffer.WriteString("\"" + argument + "\"")
	}
	buffer.WriteString("}")
	return buffer.String()
}

func getFunctionParametersAsString(parameters map[string]string) string {
	var buffer bytes.Buffer
	start := 0
	buffer.WriteString("map[string]string{")
	for key, value := range parameters {
		if start != 0 {
			buffer.WriteString(",")
		}
		buffer.WriteString("\"" + key + "\"" + ":" + "\"" + value + "\"")
		start++
	}
	buffer.WriteString("}")
	return buffer.String()
}

// Check a WarpScript GTS and a Test GTS equality
func CheckEquality(gts core.GeoTimeSeries, testGts FloatGeoTimeSeries, t *testing.T, index int, id string, skipValues bool) {
	// Class names tests
	if gts.Class != testGts.Class {
		t.Errorf("Function %s: Gts at index %d gts name doesn't match expected value: %s, got %s", id, index, testGts.Class, gts.Class)
	}

	// Value length test
	if len(gts.Values) != len(testGts.Values) {
		t.Errorf("Function %s: Gts at index %d values size doesn't match expected size of %d", id, index, len(testGts.Values))
	}

	// Labels tests
	for key, value := range testGts.Labels {
		gtsLabelValue, ok := gts.Labels[key]

		if !ok {
			t.Errorf("Function %s: Gts at index %d miss label key: %s", id, index, key)
		} else if gtsLabelValue != value {
			t.Errorf("Function %s: Gts at index %d label value doesn't match for key %s and value %s, got %s", id, index, key, value, gtsLabelValue)
		}
	}

	// Individual values tests
	for testIndex, value := range testGts.Values {
		//t.Logf("index %d: %f == %f", testIndex, gts.Values[testIndex], value)

		gtsValue := gts.Values[testIndex]

		for itemIndex, item := range value {

			if skipValues && itemIndex >= 1 {
				continue
			}
			if Round(item, 0.001) != Round(gtsValue[itemIndex].(float64), 0.001) {
				t.Errorf("Function %s: Values at index %d and item %d doesn't match expected value: %f, got %f", id, testIndex, itemIndex, value, gts.Values[testIndex])
			}
		}
	}

}

// BuildGtsList generate a WarpScript List from a FloatGeoTimeSeries array
func BuildGtsList(seriesList []FloatGeoTimeSeries) string {
	var buffer bytes.Buffer

	buffer.WriteString("[ \n")

	for _, gts := range seriesList {
		buffer.WriteString(BuildGts(gts))
		buffer.WriteString("\n")
	}
	buffer.WriteString("]")
	return buffer.String()
}

// Add a bucketize to a Series List
func BucketizeGtsList(seriesList string) string {

	var buffer bytes.Buffer

	buffer.WriteString(seriesList)
	buffer.WriteString("\n")
	buffer.WriteString("[ SWAP bucketizer.mean 205000000 60 s 0 ] BUCKETIZE")

	return buffer.String()
}

// BuildGts generate a single GTS in WarpScript
func BuildGts(gts FloatGeoTimeSeries) string {
	var buffer bytes.Buffer

	// Set GTS Name
	buffer.WriteString("NEWGTS ")
	buffer.WriteString("'" + gts.Class + "'")
	buffer.WriteString("  RENAME \n")

	// Set Labels
	buffer.WriteString("{ ")
	for key, value := range gts.Labels {
		buffer.WriteString("'" + key + "' '" + value + "' ")
	}
	buffer.WriteString("} RELABEL \n")

	// Set Values
	for _, value := range gts.Values {
		for indexValue, data := range value {
			if indexValue == len(value)-1 && len(value) < 5 {
				i := len(value)
				for i < 5 {
					buffer.WriteString("NaN ")
					i++
				}
				buffer.WriteString(fmt.Sprintf("%f ", data))
			} else {
				buffer.WriteString(fmt.Sprintf("%f ", data))
			}
		}
		buffer.WriteString("ADDVALUE\n")
	}
	return buffer.String()
}

// Sample GTS used in Unit Test
var gts3 = FloatGeoTimeSeries{
	Class:  "sample",
	Labels: map[string]string{"label": "41", "other": "test"},
	Attrs:  map[string]string{},
	Values: [][]float64{{0, 1}, {40000000, 2}, {82000000, 1}, {110000000, 0}, {129000000, 1}, {159000000, 2}, {192000000, 0}, {205000000, 1}},
}

// Sample GTS used in Unit Test
var gts2 = FloatGeoTimeSeries{
	Class:  "sample",
	Labels: map[string]string{"label": "41", "other": "test"},
	Attrs:  map[string]string{},
	Values: [][]float64{{0, 1}, {35000000, 1}, {60000000, 2}, {72000000, 3}, {88000000, 4},
		{112000000, 5}, {122000000, 6}, {132000000, 7}, {162000000, 8}, {182000000, 9}, {202000000, 10}},
}

// Sample GTS used in Unit Test
var gts1 = FloatGeoTimeSeries{
	Class:  "sample",
	Labels: map[string]string{"label": "42"},
	Attrs:  map[string]string{},
	Values: [][]float64{{0, 1}, {35000000, -1}, {60000000, 2}, {72000000, -2}, {88000000, 4},
		{112000000, 8}, {122000000, 12}, {132000000, 16}, {162000000, 18}, {182000000, -20}, {202000000, 100}},
}

// Sample GTS used in Unit Test
var emptyGTS = FloatGeoTimeSeries{
	Class:  "empty",
	Labels: map[string]string{"label": "42"},
	Attrs:  map[string]string{},
	Values: [][]float64{},
}

// ToAbsoluteGts generate a GTS applying math.Abs to all values
func ToAbsoluteGts(gts FloatGeoTimeSeries) *FloatGeoTimeSeries {
	absoluteGts := &FloatGeoTimeSeries{
		Class:  gts.Class,
		Labels: gts.Labels,
		Attrs:  gts.Attrs,
		Values: gts.Values,
	}

	for index, value := range absoluteGts.Values {
		indexValue := len(value) - 1
		value[indexValue] = math.Abs(value[indexValue])
		absoluteGts.Values[index] = value
	}
	return absoluteGts
}

// ToAbsoluteGts generate a GTS applying math.Abs to all values
func renameGts(gts FloatGeoTimeSeries, newName string) *FloatGeoTimeSeries {
	newNameGTS := &FloatGeoTimeSeries{
		Class:  newName,
		Labels: gts.Labels,
		Attrs:  gts.Attrs,
		Values: gts.Values,
	}
	return newNameGTS
}

// Round to a custom unit
func Round(x, unit float64) float64 {
	return math.Round(x/unit) * unit
}
