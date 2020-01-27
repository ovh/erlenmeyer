package opentsdb

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
	"time"

	"github.com/ovh/erlenmeyer/core"
	"github.com/ovh/erlenmeyer/middlewares"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

const (
	errIllFormedQueryLookupStringText = "Ill formed query lookup string"
)

var (
	errIllFormedQueryLookupString = errors.New(errIllFormedQueryLookupStringText)
)

type tag struct {
	Key   string `json:"key"`
	Value string `json:"value"`
}

type queryLookup struct {
	Metric string `json:"metric"`
	Tags   []tag  `json:"tags"`
}

type result struct {
	Metric string            `json:"metric"`
	TSUID  int               `json:"tsuid"`
	Tags   map[string]string `json:"tags"`
}

type lookupResponse struct {
	Type         string   `json:"type"`
	Query        string   `json:"query,omitempty"`
	Limit        int      `json:"limit"`
	StartIndex   int      `json:"startIndex"`
	Metric       string   `json:"metric"`
	Tags         []tag    `json:"tags"`
	Time         int64    `json:"time"`
	TotalResults int      `json:"totalResults"`
	Results      []result `json:"results"`
}

func unMarshallQueryLookupFromQueryString(queryString string) (*queryLookup, error) {
	trimmedQuery := strings.TrimSpace(queryString)
	if len(trimmedQuery) == 0 {
		return nil, errIllFormedQueryLookupString
	}
	leftBraceIndex := strings.Index(trimmedQuery, "{")
	rightBraceIndex := strings.Index(trimmedQuery, "}")
	if !checkBraceIndexes(leftBraceIndex, rightBraceIndex, len(trimmedQuery)-1) {
		return nil, errIllFormedQueryLookupString
	}

	var query queryLookup
	if noBraces(leftBraceIndex, rightBraceIndex) {
		query = queryLookup{Metric: trimmedQuery, Tags: []tag{}}
	} else {
		query = queryLookup{Metric: trimmedQuery[:leftBraceIndex]}
		tagDefinitions := trimmedQuery[leftBraceIndex+1 : rightBraceIndex]
		tagDefinitionList := strings.Split(tagDefinitions, ",")
		tags := make([]tag, 0, len(tagDefinitionList))
		for _, definition := range tagDefinitionList {
			key, value, err := matchTagDefinition(definition)
			if err != nil {
				return nil, errIllFormedQueryLookupString
			}
			tags = append(tags, tag{Key: key, Value: value})
		}
		query.Tags = tags
	}

	return &query, nil
}

func processMetric(metric string) (string, error) {
	if metric == "" || metric == wildcard {
		return wildcardEinsteinREDefinition, nil
	}
	if !isIdentifier(metric) {
		return "", errIllFormedQueryLookupString
	}
	return "=" + metric, nil
}

func processTags(tagList []tag) (map[string]string, error) {
	processedTags := make(map[string]string)
	for _, tag := range tagList {
		if !isIdentifier(tag.Key) || !isValue(tag.Value) {
			return make(map[string]string), errIllFormedQueryLookupString
		}
		if tag.Value == wildcard {
			processedTags[tag.Key] = wildcardEinsteinREDefinition
		}
		value, ok := processedTags[tag.Key]
		if ok {
			if value != wildcardEinsteinREDefinition {
				processedTags[tag.Key] = fmt.Sprintf("~%s|%s", strings.TrimPrefix(value, "~"), tag.Value)
			}
		} else {
			processedTags[tag.Key] = fmt.Sprintf("%s", tag.Value)
		}
	}
	return processedTags, nil
}

// nolint: interfacer
func buildWarpScriptFromQueryLookup(query *queryLookup, out *bytes.Buffer) error {
	metric, err := processMetric(query.Metric)
	if err != nil {
		return err
	}
	tags, err := processTags(query.Tags)
	if err != nil {
		return err
	}

	fmt.Fprintf(out, "'%s'\n", metric)
	for key, value := range tags {
		fmt.Fprintf(out, "'%s'\n", key)
		fmt.Fprintf(out, "'%s'\n", value)
	}
	fmt.Fprintf(out, "%d ->MAP\n", 2*len(tags))
	fmt.Fprint(out, " 3 ->LIST FIND\n")
	fmt.Fprint(out, "'gts' STORE { 'fetched' 0 'count' 0 'gts' $gts }\n")

	return nil
}

func extractQueryLookupFromRequest(request *http.Request) (*queryLookup, error) {
	var err error
	query := &queryLookup{}
	switch request.Method {
	case "GET":
		queryParam := request.URL.Query().Get("m")
		if query, err = unMarshallQueryLookupFromQueryString(queryParam); err != nil {
			log.WithFields(log.Fields{
				"m":     queryParam,
				"proto": "opentsdb",
				"error": err,
			}).Warn("can't parse query param")

			return nil, fmt.Errorf("can't parse query param 'm=%s'", queryParam)
		}
	case "POST":
		defer request.Body.Close()
		var body []byte
		body, err = ioutil.ReadAll(request.Body)
		if err != nil {
			log.WithFields(log.Fields{
				"error":  err.Error(),
				"proto":  "opentsdb",
				"source": request.RemoteAddr,
				"method": request.Method,
				"path":   request.URL.String(),
				"ip":     request.Header.Get("X-Forwarded-For"),
			}).Error("can't fully read the request body")
			return nil, err
		}
		if err := json.Unmarshal(body, query); err != nil {
			log.WithFields(log.Fields{
				"error":  err.Error(),
				"proto":  "opentsdb",
				"source": request.RemoteAddr,
				"method": request.Method,
				"path":   request.URL.String(),
				"ip":     request.Header.Get("X-Forwarded-For"),
				"query":  query.ToString(),
			}).Warn("can't parse the request body")
			return nil, err
		}
	default:
		log.WithFields(log.Fields{
			"error":  err.Error(),
			"proto":  "opentsdb",
			"source": request.RemoteAddr,
			"method": request.Method,
			"path":   request.URL.String(),
			"ip":     request.Header.Get("X-Forwarded-For"),
		}).Warn("Unsupported HTTP method")
		return nil, fmt.Errorf("Unsupported HTTP method: %s", request.Method)
	}
	return query, nil
}

func buildLookupResponse(metric string, tags []tag, elapsedTime int64, gtsList []core.GeoTimeSeries) *lookupResponse {
	results := make([]result, 0, len(gtsList))
	for _, gts := range gtsList {
		delete(gts.Labels, ".app")
		result := result{
			Metric: gts.Class,
			Tags:   gts.Labels,
			TSUID:  0,
		}
		results = append(results, result)
	}
	return &lookupResponse{
		Type:         "LOOKUP",
		Metric:       metric,
		Tags:         tags,
		Limit:        25,
		Time:         elapsedTime,
		Results:      results,
		StartIndex:   0,
		TotalResults: len(results),
	}
}

// HandleLookup handle OpenTSDB response
// nolint: golint
func (c *OpenTSDB) HandleLookup(responseWriter http.ResponseWriter, request *http.Request) {
	var err error
	startTime := time.Now()

	token := core.RetrieveToken(request)
	if len(token) == 0 {
		c.WarnCounter.Inc()
		http.Error(responseWriter, "Not authorized", 401)
		return
	}

	var query *queryLookup

	query, err = extractQueryLookupFromRequest(request)
	if err != nil {
		c.WarnCounter.Inc()
		http.Error(responseWriter, err.Error(), 400)
		return
	}

	out := &bytes.Buffer{}

	fmt.Fprint(out, "JSONSTRICT\n")
	fmt.Fprintf(out, "'%s'\n", token)

	if err = buildWarpScriptFromQueryLookup(query, out); err != nil {
		c.ErrCounter.Inc()
		message := "Unable to build a backend query from the request"
		log.WithFields(log.Fields{
			"error":  err.Error(),
			"proto":  "opentsdb",
			"source": request.RemoteAddr,
			"method": request.Method,
			"path":   request.URL.String(),
			"ip":     request.Header.Get("X-Forwarded-For"),
		}).Warn(message)
		http.Error(responseWriter, message, 400)
		return
	}

	outStr := out.String()
	var queryResult *core.QueryResult
	warpServer := core.NewWarpServer(viper.GetString("warp_endpoint"), "opentsdb-lookup")

	queryResult, err = warpServer.QueryGTS(outStr, responseWriter.Header().Get(middlewares.TxnHeader))
	if err != nil {
		c.ErrCounter.Inc()
		log.WithFields(log.Fields{
			"error":  err.Error(),
			"proto":  "opentsdb",
			"source": request.RemoteAddr,
			"method": request.Method,
			"path":   request.URL.String(),
			"ip":     request.Header.Get("X-Forwarded-For"),
		}).Warn("Bad response from Egress")
		http.Error(responseWriter, err.Error(), 500)
		return
	}

	elapsedTime := time.Since(startTime).Nanoseconds() / 1000
	response := buildLookupResponse(query.Metric, query.Tags, elapsedTime, queryResult.GTS)

	responseWriter.Header().Add("Content-Type", "application/json")
	responseWriter.WriteHeader(200)
	json.NewEncoder(responseWriter).Encode(response)
}

func (ql *queryLookup) ToString() string {
	queryStr := ql.Metric + ": "
	for _, tag := range ql.Tags {
		queryStr += tag.Key + "=" + tag.Value + " "
	}
	return queryStr
}
