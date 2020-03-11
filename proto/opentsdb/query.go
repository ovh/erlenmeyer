package opentsdb

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"regexp"
	"strconv"
	"strings"
	"time"

	"github.com/ovh/erlenmeyer/core"
	"github.com/ovh/erlenmeyer/middlewares"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"

	"github.com/prometheus/client_golang/prometheus"
)

const (
	none                            = "none"
	errIllFormedTagText             = "Ill formed tag"
	timeseriesQueryStringKey        = "timeseries"
	errIllFormedQueryLastStringText = "Ill formed query last string"
	scriptHeader                    = `JSONSTRICT
  '%s' 'token' STORE
  []
  `
	subScriptTop = `[ $token '=%s'
  {
  `
	subScriptLabelLine = `'%s' '=%s'
  `
	subScriptBottom = `}
  NOW -1 ] FETCH
  <%
  DROP
  'ts' STORE
  $ts TICKLIST 0 GET 1000000 / 't' STORE
  $ts VALUES 0 GET TOSTRING 'v' STORE
  $ts LABELS '.app' REMOVE DROP 'l' STORE
  { 'metric' $ts NAME 'timestamp' $t 'value' $v 'tags' $l 'tsuid' 0 }
  %>
  LMAP APPEND
  `
	scriptFooter = `DUP SIZE 'pointsCount' STORE 'series' STORE
  { 'pointsCount' $pointsCount 'series' $series }
  `
	identifier                   = `[-a-zA-Z0-9_./]+`
	wildcard                     = "*"
	wildcardEinsteinREDefinition = "~.*"
	value                        = `\` + wildcard + `|` + identifier
	tagDefinition                = `(` + identifier + `)=(` + value + `)`
	tagDefinitionNoWildcards     = `(` + identifier + `)=(` + identifier + `)`

	// WildCard keyword for OpenTSDB filters
	WildCard = "wildcard"

	// LiteralOr keyword for OpenTSDB filters
	LiteralOr = "literal_or"

	// RegExp keyword for OpenTSDB filters
	RegExp = "regexp"
)

// Translator openTSDB translator to WarpScript
type Translator func(string) string

func rawRegularExpressionTranslator(in string) string {
	return "~" + in
}

func wildcardTranslator(in string) string {
	return rawRegularExpressionTranslator(strings.Replace(in, "*", ".*", -1))
}

func litteralOrTranslator(in string) string {
	if strings.Contains(in, "|") {
		return rawRegularExpressionTranslator(in)
	}
	return "=" + in
}

type queries struct {
	Queries []queryLast `json:"queries"`
}

type queryLast struct {
	Metric string            `json:"metric"`
	Tags   map[string]string `json:"tags"`
}
type queryLasts []queryLast

// Response OpenTSDB response struct
type Response struct {
	PointsCount int             `json:"pointsCount"`
	Buffer      json.RawMessage `json:"series"`
}

// Make sure we implement Validatable
var _ Validatable = &QueryRequest{}

// Validatable is the interface for the Validate method
type Validatable interface {
	Validate() error
}

// OpenTSDB endpoint
type OpenTSDB struct {
	ReqCounter  prometheus.Counter
	ErrCounter  prometheus.Counter
	WarnCounter prometheus.Counter
}

// GetReqCounter satisfies the protocol interface
func (o *OpenTSDB) GetReqCounter() prometheus.Counter {
	return o.ReqCounter
}

// QueryRequest an OpenTSDB valid request
type QueryRequest struct {
	// The start time for the query
	Start *TSDBTime `json:"start"`
	// An end time for the query. Default value = time.Now()
	End         *TSDBTime `json:"end"`
	NoTimeRange bool      `json:"-"`
	// One or more sub queries used to select the time series to return
	Queries []*Query `json:"queries"`
	// Whether or not to output data point timestamps in milliseconds or seconds
	MSResolution *bool `json:"msResolution"`
	Delete       bool  `json:"delete"`
}

// Query an OpenTSDB single query
type Query struct {
	// The name of an aggregation function to use
	Aggregator string `json:"aggregator"`
	// The name of a metric stored in the system
	Metric string `json:"metric"`
	// Whether or not the data should be converted into deltas before returning
	Rate *bool `json:"rate"`
	// Monotonically increasing counter handling options
	RateOptions *RateOptions `json:"rateOptions"`
	// An optional downsampling function to reduce the amount of data returned
	Downsample *string `json:"downsample"`
	// To drill down to specific timeseries or group results by tag
	Tags map[string]string `json:"tags"`
	// Filters the timeseries emitted in the results
	Filters []FilterSpec `json:"filters"`
	// Returns the series that include only the tag keys provided in the filters.
	ExplicitTags bool `json:"explicitTags"`
}

// QueryResponse an OpenTSDB Query response
type QueryResponse struct {
	Metric        string             `json:"metric"`
	Tags          map[string]string  `json:"tags"`
	Query         IndexResponse      `json:"query"`
	AggregateTags []string           `json:"aggregateTags"`
	DPs           map[string]float64 `json:"dps"`
}

// IndexResponse query struct containing the query index
type IndexResponse struct {
	Index int `json:"index"`
}

var (
	errIllFormedTag             = errors.New(errIllFormedTagText)
	identifierRE                = regexp.MustCompile("^" + identifier + "$")
	valueRE                     = regexp.MustCompile("^(" + value + ")$")
	tagDefinitionRE             = regexp.MustCompile("^" + tagDefinition + "$")
	tagDefinitionNoWildcardsRE  = regexp.MustCompile("^" + tagDefinitionNoWildcards + "$")
	errIllFormedQueryLastString = errors.New(errIllFormedQueryLastStringText)
	intervals                   = [16]int64{
		1000000, 5000000, 10000000, 15000000, 30000000, 60000000, 600000000,
		900000000, 1800000000, 3600000000, 7200000000, 14400000000, 21600000000,
		43200000000, 86400000000, 604800000000,
	}
)

// RateOptions monotonically increasing counter handling options
type RateOptions struct {
	// Whether or not the underlying data is a monotonically increasing counter that may roll over
	Counter bool `json:"counter"`
	// A positive integer representing the maximum value for the counter
	CounterMax *int64 `json:"counterMax"`
	// An optional value that, when exceeded, will cause the aggregator to return
	// a 0 instead of the calculated rate
	ResetValue *int64 `json:"resetValue"`
	// Optional: Whether or not to simply drop rolled-over or reset data points.
	DropResets bool `json:"dropResets"`
}

// FilterSpec the time series emitted in the results. Note that if no filters are specified, all time series for the given metric will be aggregated into the results.
type FilterSpec struct {
	Type       string `json:"type"`
	TagKey     string `json:"tagk"`
	Expression string `json:"filter"`
	GroupBy    bool   `json:"groupBy"`
}

var downsamplingRe, _ = regexp.Compile(`^` + durationPattern + `-([^-]+)(?:-(.+))?$`)

// Translation from OpenTSDB aggregator to WarpScript reducer
var downsamplerToBucketizer = map[string]string{
	"avg":    "bucketizer.mean",
	"sum":    "bucketizer.sum",
	"zimsum": "bucketizer.sum",
	"mimmin": "bucketizer.min",
	"mimmax": "bucketizer.max",
	"min":    "bucketizer.min",
	"max":    "bucketizer.max",
	"count":  "bucketizer.count",
	"first":  "bucketizer.first",
	"last":   "bucketizer.last",
	"p50":    "50.0 bucketizer.percentile",
	"p75":    "75.0 bucketizer.percentile",
	"p90":    "90.0 bucketizer.percentile",
	"p95":    "95.0 bucketizer.percentile",
	"p99":    "99.0 bucketizer.percentile",
	"p999":   "99.9 bucketizer.percentile",
	"dev":    "true reducer.sd",
}

// HandleQuery Entry point of an OpenTSDB request
func (o *OpenTSDB) HandleQuery(w http.ResponseWriter, r *http.Request) {

	token := core.RetrieveToken(r)
	if len(token) == 0 {
		o.WarnCounter.Inc()
		http.Error(w, "Not authorized", 401)
		return
	}

	query := &QueryRequest{}

	switch r.Method {
	case http.MethodGet:
		o.WarnCounter.Inc()
		http.Error(w, "Please use the POST version of queries", http.StatusNotImplemented)
		log.WithFields(log.Fields{
			"proto": "opentsdb",
		}).Error("GET Queries not implemented")
		return
	case http.MethodPost:
		if success, message, status := query.parseRequestBody(r.Body); !success {
			o.WarnCounter.Inc()
			http.Error(w, message, status)
			return
		}
		if query.Delete {
			executeDelete(w, token, query)
		} else {
			executeQuery(w, token, query)
		}
	case http.MethodDelete:
		if success, message, status := query.parseRequestBody(r.Body); !success {
			o.WarnCounter.Inc()
			http.Error(w, message, status)
			return
		}
		executeDelete(w, token, query)
	default:
		o.WarnCounter.Inc()
		w.Header().Add("Allow", "POST, DELETE")
		http.Error(w, fmt.Sprintf("Method '%s' is not allowed", r.Method), http.StatusMethodNotAllowed)
		return
	}
}

// HandleQueryLast is the handler for /api/query/last
func (o *OpenTSDB) HandleQueryLast(responseWriter http.ResponseWriter, request *http.Request) { // nolint: golint

	token := core.RetrieveToken(request)
	if len(token) == 0 {
		o.WarnCounter.Inc()
		log.WithFields(log.Fields{
			"proto":  "opentsdb",
			"source": request.RemoteAddr,
			"method": request.Method,
			"status": 401,
			"path":   request.URL.String(),
		}).Error("401 Not Authorized")
		http.Error(responseWriter, "Not authorized", 401)
		return
	}

	queries, err := extractQueryLastFromRequest(request)
	if err != nil {
		o.WarnCounter.Inc()
		log.WithFields(log.Fields{
			"proto":  "opentsdb",
			"source": request.RemoteAddr,
			"method": request.Method,
			"status": 400,
			"path":   request.URL.String(),
		}).Error("Bad request")
		http.Error(responseWriter, "Bad request", 400)
		return
	}
	out := &bytes.Buffer{}
	if err = buildWarpScriptFromQueryLast(queries, token, out); err != nil {
		o.WarnCounter.Inc()
		message := "Unable to build a backend query from the request"
		log.WithFields(log.Fields{
			"proto":  "opentsdb",
			"query":  queries.ToString(),
			"source": request.RemoteAddr,
			"method": request.Method,
			"status": 400,
			"path":   request.URL.String(),
		}).Error(message)
		http.Error(responseWriter, message, http.StatusBadRequest)
		return
	}

	outStr := out.String()

	warpServer := core.NewWarpServer(viper.GetString("warp_endpoint"), "opentsdb-query-last")
	response, err := warpServer.Query(outStr, responseWriter.Header().Get(middlewares.TxnHeader))
	if err != nil {
		o.ErrCounter.Inc()
		message := "Bad response from Egress"
		log.WithFields(log.Fields{
			"error": err.Error(),
			"proto": "opentsdb",
		}).Error(message)
		http.Error(responseWriter, message, http.StatusBadGateway)
		return
	}

	buffer, err := ioutil.ReadAll(response.Body)
	if err != nil {
		o.ErrCounter.Inc()
		message := "can't fully read Egress response"
		log.WithFields(log.Fields{
			"error": err.Error(),
			"proto": "opentsdb",
		}).Error(message)
		http.Error(responseWriter, message, http.StatusBadGateway)
		return
	}

	responses := []Response{}
	err = json.Unmarshal(buffer, &responses)
	if err != nil {
		o.ErrCounter.Inc()
		message := "Failed to unmarshal Egress response"
		log.WithFields(log.Fields{
			"error": err.Error(),
			"proto": "opentsdb",
		}).Error(message)
		http.Error(responseWriter, message, http.StatusBadGateway)
		return
	}

	if len(responses) != 1 {
		o.ErrCounter.Inc()
		message := "Wrong number of elements in Warp 10 response"
		log.WithFields(log.Fields{
			"error": err.Error(),
			"proto": "opentsdb",
		}).Error(message)
		http.Error(responseWriter, message, http.StatusBadGateway)
		return
	}
	responseWriter.Header().Add("Content-Type", "application/json")
	responseWriter.WriteHeader(200)
	responseWriter.Write(responses[0].Buffer)
}

func (handler *QueryRequest) parseRequestBody(requestBody io.ReadCloser) (bool, string, int) {
	defer requestBody.Close()
	body, err := ioutil.ReadAll(requestBody)
	if err != nil {
		log.WithFields(log.Fields{
			"error": err.Error(),
			"proto": "opentsdb",
		}).Error("can't fully read the request body")
		return false, "Bad request body", http.StatusInternalServerError
	}
	if err := json.Unmarshal(body, handler); err != nil {
		msg := "can't parse query"
		log.WithFields(log.Fields{
			"error":   err.Error(),
			"proto":   "opentsdb",
			"request": string(body),
		}).Warn(msg)
		return false, msg, http.StatusBadRequest
	}

	if err := handler.Validate(); err != nil {
		// Using err.Error() here as Validate returns user-facing messages.
		// Need to add a ValidationError though and check for it, to avoid potential mistakes
		log.WithFields(log.Fields{
			"error":   err.Error(),
			"proto":   "opentsdb",
			"request": string(body),
		}).Warn("Invalid query")
		return false, fmt.Sprintf("Invalid query: %s", err.Error()), http.StatusBadRequest
	}
	return true, "", 0
}

// NewOpenTSDB Instantiate an OpenTSDB struct
func NewOpenTSDB() *OpenTSDB {
	c := &OpenTSDB{}

	// metrics
	c.ReqCounter = prometheus.NewCounter(prometheus.CounterOpts{
		Namespace: "erlenmeyer",
		Subsystem: "opentsdb",
		Name:      "request",
		Help:      "Number of request handled.",
	})
	prometheus.MustRegister(c.ReqCounter)
	c.ErrCounter = prometheus.NewCounter(prometheus.CounterOpts{
		Namespace: "erlenmeyer",
		Subsystem: "opentsdb",
		Name:      "errors",
		Help:      "Number of request in errors.",
	})
	prometheus.MustRegister(c.ErrCounter)
	c.WarnCounter = prometheus.NewCounter(prometheus.CounterOpts{
		Namespace: "erlenmeyer",
		Subsystem: "opentsdb",
		Name:      "warning",
		Help:      "Number of errored client requests.",
	})
	prometheus.MustRegister(c.WarnCounter)

	return c
}

// GetRate is it a rate query
func (q *Query) GetRate() bool { // nolint: golint
	if q.Rate == nil {
		return false
	}
	return *q.Rate
}

// Validate query
func (q *Query) Validate() error { // nolint: golint
	if q.Metric == "" {
		log.WithFields(log.Fields{
			"proto": "opentsdb",
		}).Error("missing metric name")
		return errors.New("missing metric name")
	}

	if _, exists := aggregatorToReduce[q.Aggregator]; !exists {
		log.WithFields(log.Fields{
			"proto": "opentsdb",
		}).Errorf("invalid aggregator %s", q.Aggregator)
		return fmt.Errorf("invalid aggregator %s", q.Aggregator)
	}

	if q.Rate == nil {
		q.Rate = new(bool)
		*q.Rate = false
	} else if *q.Rate {
		if q.RateOptions == nil {
			q.RateOptions = &RateOptions{}
		}
	}

	return nil
}

// Apply a dropReset, as specify in OpenTSDB: dropResets can be apply only for monotonically increasing counter
// http://opentsdb.net/docs/build/html/api_http/query/index.html?#sub-queries (rateOptions)
func dropReset(out *bytes.Buffer) {
	// If non empty series
	out.WriteString("<% DUP SIZE 0 > %>\n")
	out.WriteString("<% false RESETS %> IFT\n")
}

// Filter value to limit series value to user counter specified max.
func counterReplaceMax(out *bytes.Buffer, max int64) {

	out.WriteString("UNBUCKETIZE [ SWAP ")

	out.WriteString(strconv.FormatInt(max, 10))
	out.WriteString(" mapper.min.x 0 0 0 ] MAP\n")
}

// Macromapper to apply resetValue.
// We need this for resetValue in RateOptions
func resetMacro(out *bytes.Buffer, max int64) {

	out.WriteString("<%\n")
	out.WriteString("\t'mapping_window' STORE\n")                   //  Storing macro input information
	out.WriteString("\t$mapping_window 7 GET 0 GET\n")              // Extract the current value
	fmt.Fprint(out, "\t<%  "+strconv.FormatInt(max, 10)+" >= %>\n") // Condition goes here

	out.WriteString("\t<%\n")
	out.WriteString("\t\t$mapping_window 0 GET\n")       // Tick
	out.WriteString("\t\t$mapping_window 4 GET 0 GET\n") // Latitude
	out.WriteString("\t\t$mapping_window 5 GET 0 GET\n") // Longitude
	out.WriteString("\t\t$mapping_window 6 GET 0 GET\n") // Elevation
	out.WriteString("\t\t0\n")                           // Value
	out.WriteString("\t%>\n")
	// Then
	out.WriteString("\t<%\n")
	out.WriteString("\t\t$mapping_window 0 GET\n")       // Tick
	out.WriteString("\t\t$mapping_window 4 GET 0 GET\n") // Latitude
	out.WriteString("\t\t$mapping_window 5 GET 0 GET\n") // Longitude
	out.WriteString("\t\t$mapping_window 6 GET 0 GET\n") // Elevation
	out.WriteString("\t\t$mapping_window 7 GET 0 GET\n") // Value
	out.WriteString("\t%>\n")
	out.WriteString("\tIFTE\n")
	out.WriteString("%>\n")
	// We are using this MACRO MAPPER like a classic Mapper
	out.WriteString("MACROMAPPER\n")
	out.WriteString("1 0 0\n")
	out.WriteString("5 ->LIST MAP\n")
}

func needsInterpolate(operator string) bool {
	switch operator {
	case "zimsum", "count", "mimmin", "mimmax", "first", "last":
		return false
	default:
		return true
	}
}

// nolint: gocyclo
// agregation: perform an aggregation on an OpenTSDB subquery
func (q *Query) agregation(out *bytes.Buffer, groupingTags []string) (success bool, message string, httpCode int) { // nolint: golint

	// If rate is asked for, add another reducer
	if q.Rate != nil && *q.Rate {
		if q.RateOptions != nil && q.RateOptions.Counter {

			if q.RateOptions.DropResets == true {
				dropReset(out)
			}
			if q.RateOptions.CounterMax != nil {
				counterReplaceMax(out, *q.RateOptions.CounterMax)
			}
		}
	}
	if q.Aggregator == none {

		fmt.Fprint(out, "'gts' STORE\n")
		success = true
		return
	}
	// Equivalence class
	if len(groupingTags) == 0 {
		fmt.Fprint(out, "[]\n")
	} else {
		for _, tagk := range groupingTags {
			fmt.Fprintf(out, "'%s' ", tagk)
		}
		fmt.Fprintf(out, "%d ->LIST\n", len(groupingTags))
	}

	fmt.Fprintf(out, "%s\n", aggregatorToReduce[q.Aggregator])
	fmt.Fprint(out, "3 ->LIST REDUCE\n\n")

	// If rate is asked for, add another reducer
	if q.Rate != nil && *q.Rate {

		if q.RateOptions != nil && q.RateOptions.Counter {

			if q.RateOptions.ResetValue != nil {
				// For some strange reason, bosun is putting 1 as ResetValue
				// As standard value
				if *q.RateOptions.ResetValue != 1 {
					resetMacro(out, *q.RateOptions.ResetValue)
				}
			}
		}

		// mapper pre post occurrences
		// OpenTSDB rate option correspond to mapper.rate
		// and not mapper.delta
		fmt.Fprint(out, "mapper.rate\n")
		fmt.Fprint(out, "1 0 $end TOTIMESTAMP $start TOTIMESTAMP - $bucketspan / TOLONG 1 - -1 * \n")
		fmt.Fprint(out, "5 ->LIST MAP\n")
		fmt.Fprint(out, " \n")
	}
	fmt.Fprint(out, "'gts' STORE\n")

	success = true
	return
}

func appendResponses(warp10Results []core.GeoTimeSeries, responses []*QueryResponse, groupingTags []string, queryResolution *bool, metric string, loopIndex int) ([]*QueryResponse, int) {
	byteCount := 0

	for _, result := range warp10Results {
		delete(result.Labels, ".app")
		rep := &QueryResponse{
			Metric:        metric,
			Tags:          result.Labels,
			AggregateTags: groupingTags,
			DPs:           map[string]float64{},
			Query:         IndexResponse{Index: loopIndex},
		}

		for i := range result.Values {
			if hasValue, ts, value := result.TimeAndValue(i); hasValue {
				var tsString string
				if queryResolution != nil && *queryResolution {
					tsString = strconv.FormatInt(ts.Unix()*int64(1000)+int64(ts.Nanosecond()/1000000), 10)
				} else {
					tsString = strconv.FormatInt(ts.Unix(), 10)
				}
				rep.DPs[tsString] = value
			}
		}
		if len(rep.DPs) != 0 {
			byteCount += WarpResultByteCount(rep)
			responses = append(responses, rep)
		}
	}
	return responses, byteCount
}

// Validate is the implementation of Validate for QueryRequest
func (handler *QueryRequest) Validate() error { // nolint: golint
	if handler.Start == nil && handler.End == nil {
		handler.NoTimeRange = true
	}
	if handler.Start == nil {
		handler.Start = &TSDBTime{}
	}
	if handler.End == nil {
		now := time.Now()
		handler.End = &TSDBTime{now}
	}

	if handler.Start.Time.After(handler.End.Time) {
		return errors.New("Start date is after end date")
	}

	for _, query := range handler.Queries {
		if err := query.Validate(); err != nil {
			return err
		}
	}
	return nil
}

// processTags process tags of an OpenTSDB sub-query
func (q *Query) processTags(out *bytes.Buffer, groupingTags []string) ([]string, int) { // nolint: golint, interfacer
	// Grouping tags based on native WarpScript "equivalence class"
	tags := q.Tags
	for tagk, tagv := range tags {
		if tagv == "*" {
			// Simple group by
			groupingTags = append(groupingTags, tagk)
			// Add a filter that makes sure the tag exists
			fmt.Fprintf(out, "'%s' '~.*'\n", tagk)

		} else if strings.Contains(tagv, "|") {
			// Filtered group by
			groupingTags = append(groupingTags, tagk)

			fmt.Fprintf(out, "'%s' '~", tagk) // open regexp

			values := strings.Split(tagv, "|")
			quotedValues := make([]string, len(values))
			for i, value := range values {
				quotedValues[i] = regexp.QuoteMeta(strings.TrimSpace(value))
			}
			fmt.Fprintf(out, "%s", strings.Join(quotedValues, "|"))
			fmt.Fprint(out, "'\n") // close regexp

		} else {
			// Regular filter
			fmt.Fprintf(out, "'%s' '=%s'\n", tagk, tagv)
		}
	}
	return groupingTags, len(tags)
}

// processFilters of an OpenTSDB subquery
func (q *Query) processFilters(out *bytes.Buffer, groupingTags []string) ([]string, int) { // nolint: golint, interfacer
	count := 0
	for _, filterSpec := range q.Filters {
		filter, ok := filters[filterSpec.Type]
		if ok {
			fmt.Fprintf(out, "'%s' '%s'\n", filterSpec.TagKey, filter.Translator(filterSpec.Expression))
			if filterSpec.GroupBy {
				groupingTags = append(groupingTags, filterSpec.TagKey)
			}
			count = count + 1
		}
	}
	return groupingTags, count
}

// selection of an OpenTSDB subquery
func (q *Query) selection(out *bytes.Buffer, start time.Time, end time.Time) []string { // nolint: golint
	// Store start and stop
	fmt.Fprintf(out, "'%s' 'start' STORE\n", core.IsoTime(start))
	fmt.Fprintf(out, "'%s' 'end' STORE\n", core.IsoTime(end))
	// Metric name
	fmt.Fprintf(out, "'%s'\n", q.Metric)

	// Grouping tags based on native WarpScript "equivalence class"
	groupingTags := make([]string, 0, len(q.Tags))
	groupingTags, tagsLabelCount := q.processTags(out, groupingTags)
	groupingTags, filterLabelCount := q.processFilters(out, groupingTags)
	numberOfLabels := filterLabelCount + tagsLabelCount
	fmt.Fprintf(out, "%d ->MAP\n", numberOfLabels*2)

	// Start time, end time
	fmt.Fprintf(out, "$start $end\n")

	// Fetch
	fmt.Fprint(out, "5 ->LIST FETCHDOUBLE\n\n")

	// Keep only series with filters tags when specified
	if q.ExplicitTags {
		// Add a new label containing all labels keys in JSON format
		out.WriteString("<% DROP DUP LABELS KEYLIST UNIQUE LSORT ->JSON { 'hash_945fa9bc3027d7025e3' ROT } RELABEL %> LMAP \n")

		// Filter series set to keep only series with filters tags keys
		out.WriteString("[ SWAP [] { 'hash_945fa9bc3027d7025e3' ")
		out.WriteString(q.queriesKeyTags())
		out.WriteString(" ->JSON } filter.bylabels ] FILTER \n")

		// Remove tmp label
		out.WriteString("<% DROP { 'hash_945fa9bc3027d7025e3' '' } RELABEL %> LMAP \n")
	}

	return groupingTags
}

// Get all filters query tag as WarpScript string List
func (q *Query) queriesKeyTags() string {
	var out bytes.Buffer

	out.WriteString("[ '.app'")
	for _, filter := range q.Filters {
		out.WriteString(" '" + filter.TagKey + "'")
	}

	for key := range q.Tags {
		out.WriteString(" '" + key + "'")
	}
	out.WriteString(" ] UNIQUE LSORT")

	return out.String()
}

func downSamplingInterval(durationNano int64) int64 {
	interval := durationNano / 1000000
	for _, v := range intervals {
		if interval <= v {
			return v
		}
	}
	return intervals[15]
}

// generateBucketizeScript of an OpenTSDB subquery
// nolint: interfacer
func (q *Query) generateBucketizeScript(out *bytes.Buffer, bucketSpan int64, bucketizer string, fillPolicy string) {
	fmt.Fprint(out, "'gts' STORE\n")
	fmt.Fprintf(out, "%d 'bucketspan' STORE\n", bucketSpan)

	if bucketSpan != 0 {
		fmt.Fprintf(out, "[ $gts %s $end TOTIMESTAMP $bucketspan $end TOTIMESTAMP $start TOTIMESTAMP - $bucketspan / TOLONG ]  BUCKETIZE\n", bucketizer)
	} else {
		fmt.Fprintf(out, "[ $gts %s $end TOTIMESTAMP $end TOTIMESTAMP $start TOTIMESTAMP - 1 ]  BUCKETIZE\n", bucketizer)
	}
	switch fillPolicy {
	case "", none:
		if needsInterpolate(q.Aggregator) {
			fmt.Fprint(out, "INTERPOLATE FILLPREVIOUS FILLNEXT\n")
		}
	case "nan":
		fmt.Fprint(out, "[ NaN NaN NaN NaN ] FILLVALUE\n")
	case "zero":
		fmt.Fprint(out, "[ NaN NaN NaN 0 ] FILLVALUE\n")
	}
}

// downSampling of an OpenTSDB subquery
// nolint: golint
func (q *Query) downSampling(out *bytes.Buffer, start time.Time, end time.Time) (success bool, message string, httpCode int) {
	if q.Aggregator == none {
		success = true
		return
	}
	if q.Downsample != nil {
		splits := downsamplingRe.FindStringSubmatch(*q.Downsample)
		if splits == nil {
			message = "Wrong downsampling value: " + *q.Downsample
			httpCode = http.StatusBadRequest
			return
		}

		period := DecodeDuration(splits[1], splits[2])
		bucketizer, exists := downsamplerToBucketizer[splits[3]]
		if !exists {
			message = fmt.Sprintf("Downsampling operator '%s' is not available", splits[3])
			httpCode = http.StatusBadRequest
			return
		}

		requestedFillPolicy := none
		if len(splits) == 5 {
			requestedFillPolicy = splits[4]
		}
		if !isValidFillPolicy(requestedFillPolicy) {
			message = fmt.Sprintf("Downsampling fill policy '%s' is not available", requestedFillPolicy)
			httpCode = http.StatusBadRequest
			return
		}
		q.generateBucketizeScript(out, period.Nanoseconds()/1000, bucketizer, requestedFillPolicy)
	} else {
		// Here we force a buketizer.mean by default as there is no down sampling specified.
		// We assume this is a sensible option between all available bucketizers!
		q.generateBucketizeScript(out, downSamplingInterval(int64(end.Sub(start))), "bucketizer.mean", none)
	}

	success = true
	return
}

// nolint: gocyclo
func executeQuery(w http.ResponseWriter, token string, query *QueryRequest) {

	responses := []*QueryResponse{}

	startTimeWithRetention := query.Start.Time

	for loopIndex, subquery := range query.Queries {

		out := &bytes.Buffer{}
		fmt.Fprint(out, "JSONSTRICT\n")
		fmt.Fprintf(out, "'%s'\n", token)
		fmt.Fprintf(out, "DUP AUTHENTICATE 20000000 LIMIT\n")

		// OpenTSDB order: Selection, Grouping, Downsampling, Aggregation, Interpolation, Rate conversion
		//---- Selection: FETCH

		groupingTags := subquery.selection(out, startTimeWithRetention, query.End.Time)

		//---- Downsample: BUCKETIZE
		if success, errorMsg, httpCode := subquery.downSampling(out, startTimeWithRetention, query.End.Time); !success {
			http.Error(w, errorMsg, httpCode)
			return
		}

		//---- Aggregation: REDUCE
		if success, errorMsg, httpCode := subquery.agregation(out, groupingTags); !success {
			http.Error(w, errorMsg, httpCode)
			return
		}

		//---- Build resulting structure on top of the stack with fields «count» and «gts»
		fmt.Fprint(out, "{ 'gts' $gts SORT }\n")

		//----- Send request
		body := out.String()
		warpServer := core.NewWarpServer(viper.GetString("warp_endpoint"), "opentsdb-query")
		warp10Results, err := warpServer.QueryGTS(body, w.Header().Get(middlewares.TxnHeader))

		if err != nil {
			errorText := err.Error()
			if strings.Contains(errorText, "URLDecoder") {
				re := regexp.MustCompile(`URLDecoder: ([^<]*)`)
				e := ""
				parts := re.FindStringSubmatch(errorText)
				if len(e) > 1 {
					e = parts[1]
				}
				re = regexp.MustCompile(`line #(\d+)`)
				l := "-1"
				parts = re.FindStringSubmatch(errorText)
				if len(e) > 1 {
					l = parts[1]
				}
				log.WithFields(log.Fields{
					"error":       errorText,
					"proto":       "opentsdb",
					"query-error": e,
					"line":        l,
				}).Warn("Malformed WarpScript")
			} else if strings.Contains(errorText, "FETCH exceeded limit of") {
				re := regexp.MustCompile(`(?:limit of )(\d+)(?: datapoints, current count is )(\d+)</pre>`)
				parts := re.FindStringSubmatch(errorText)
				current := "-1"
				max := "-1"
				if len(parts) > 2 {
					current = parts[1]
					max = parts[2]
				}
				log.WithFields(log.Fields{
					"error":   errorText,
					"proto":   "opentsdb",
					"current": current,
					"max":     max,
				}).Warn("Max datapoints fetching excedeed")
			} else if strings.Contains(errorText, "SocketTimeoutException") {
				re := regexp.MustCompile(`(?:hostname=)([^,]*)`)
				parts := re.FindStringSubmatch(errorText)
				host := ""
				if len(parts) > 1 {
					host = parts[1]
				}
				log.WithFields(log.Fields{
					"error": errorText,
					"proto": "opentsdb",
					"host":  host,
				}).Error("Socket timeout")
			} else {
				log.WithFields(log.Fields{
					"error": errorText,
					"proto": "opentsdb",
				}).Error("Egress error: uncategorized error")
			}
			http.Error(w, errorText, http.StatusInternalServerError)
			return
		}

		responses, _ = appendResponses(warp10Results.GTS, responses, groupingTags, query.MSResolution, subquery.Metric, loopIndex)
	}

	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(200)
	json.NewEncoder(w).Encode(responses)
}

func isIdentifier(identifier string) bool {
	return identifierRE.MatchString(identifier)
}

func isValue(value string) bool {
	return valueRE.MatchString(value)
}

func matchTagDefinition(definition string) (string, string, error) {
	matches := tagDefinitionRE.FindStringSubmatch(definition)
	if len(matches) != 3 {
		return "", "", errIllFormedTag
	}
	return matches[1], matches[2], nil
}

func matchTagDefinitionNoWildcards(definition string) (string, string, error) {
	matches := tagDefinitionNoWildcardsRE.FindStringSubmatch(definition)
	if len(matches) != 3 {
		return "", "", errIllFormedTag
	}
	return matches[1], matches[2], nil
}

func unMarshallQueryLastFromQueryString(queryString string) (*queryLast, error) {
	trimmedQuery := strings.TrimSpace(queryString)
	if len(trimmedQuery) == 0 {
		return nil, errIllFormedQueryLastString
	}
	leftBraceIndex := strings.Index(trimmedQuery, "{")
	rightBraceIndex := strings.Index(trimmedQuery, "}")
	if !checkBraceIndexes(leftBraceIndex, rightBraceIndex, len(trimmedQuery)-1) {
		return nil, errIllFormedQueryLastString
	}

	var query queryLast
	if noBraces(leftBraceIndex, rightBraceIndex) {
		query = queryLast{Metric: trimmedQuery, Tags: map[string]string{}}
	} else {
		query = queryLast{Metric: trimmedQuery[:leftBraceIndex]}
		tagDefinitions := trimmedQuery[leftBraceIndex+1 : rightBraceIndex]
		tagDefinitionList := strings.Split(tagDefinitions, ",")
		tags := map[string]string{}
		for _, definition := range tagDefinitionList {
			key, value, err := matchTagDefinitionNoWildcards(definition)
			if err != nil {
				return nil, errIllFormedQueryLastString
			}
			if _, yetPresent := tags[key]; yetPresent {
				return nil, errIllFormedQueryLastString
			}
			tags[key] = value
		}
		query.Tags = tags
	}
	if !isIdentifier(query.Metric) {
		return nil, errIllFormedQueryLastString
	}

	return &query, nil
}

func extractQueryLastFromRequest(request *http.Request) (queryLasts, error) {
	var err error
	switch request.Method {
	case http.MethodGet:
		queries := queryLasts{}
		queryParams, timeseriesParamOk := request.URL.Query()[timeseriesQueryStringKey]
		if !timeseriesParamOk {
			return nil, errors.New("No '" + timeseriesQueryStringKey + "' parameter available")
		}
		for _, queryParam := range queryParams {
			query, err := unMarshallQueryLastFromQueryString(queryParam)
			if err != nil {
				message := fmt.Sprintf("can't parse query param '%s=%s'", timeseriesQueryStringKey, queryParam)
				return nil, errors.New(message)
			}
			queries = append(queries, *query)
		}
		return queries, nil
	case http.MethodPost:
		defer request.Body.Close()
		var body []byte
		body, err = ioutil.ReadAll(request.Body)
		if err != nil {
			log.WithFields(log.Fields{
				"error":  err.Error(),
				"proto":  "opentsdb",
				"method": request.Method,
				"path":   request.URL.String(),
				"source": request.RemoteAddr,
				"ip":     request.Header.Get("X-Forwarded-For"),
			}).Error("Can't read the full request body")
			return nil, err
		}
		result := &queries{}
		if err := json.Unmarshal(body, result); err != nil {
			log.WithFields(log.Fields{
				"error":  err.Error(),
				"proto":  "opentsdb",
				"method": request.Method,
				"path":   request.URL.String(),
				"source": request.RemoteAddr,
				"ip":     request.Header.Get("X-Forwarded-For"),
			}).Error("Can't Decode the request body")
			return nil, errors.New("can't parse the request body")
		}
		return result.Queries, nil
	default:
		log.WithFields(log.Fields{
			"method": request.Method,
			"proto":  "opentsdb",
			"path":   request.URL.String(),
			"source": request.RemoteAddr,
			"ip":     request.Header.Get("X-Forwarded-For"),
		}).Error("Unsupported HTTP method")
		return nil, fmt.Errorf("Unsupported HTTP method: %s", request.Method)
	}
}

// nolint: interfacer
func buildWarpScriptFromQueryLast(queries []queryLast, token string, out *bytes.Buffer) error {

	fmt.Fprintf(out, scriptHeader, token)
	for _, query := range queries {
		if !isIdentifier(query.Metric) {
			return errIllFormedQueryLastString
		}
		fmt.Fprintf(out, subScriptTop, query.Metric)

		for key, value := range query.Tags {
			if !isIdentifier(key) || !isIdentifier(value) {
				return errIllFormedQueryLastString
			}
			fmt.Fprintf(out, subScriptLabelLine, key, value)
		}
		fmt.Fprint(out, subScriptBottom)
	}
	fmt.Fprint(out, scriptFooter)
	return nil
}

func (qls queryLasts) ToString() string {
	s := ""
	for _, ql := range qls {
		s += ql.Metric + ":"
		for k, v := range ql.Tags {
			s += k + "=" + v
		}
		s += "\n"
	}
	return s
}
