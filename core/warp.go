package core

import (
	"bufio"
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httputil"
	"regexp"
	"strconv"
	"strings"
	"time"

	"github.com/pkg/errors"
	"github.com/prometheus/client_golang/prometheus"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

var (
	// https://regex101.com/r/jlI5ad/1
	tokenRegex = regexp.MustCompile(`(?mU)['"][a-zA-Z_.0-9]{80,}['"]`)
	tokens     = make(map[string]string)

	requests = prometheus.NewCounterVec(prometheus.CounterOpts{
		Namespace: "erlenmeyer",
		Subsystem: "exec",
		Name:      "request",
		Help:      "Warp execution count",
	}, []string{"app", "token_id", "protocol"})
	times = prometheus.NewCounterVec(prometheus.CounterOpts{
		Namespace: "erlenmeyer",
		Subsystem: "exec",
		Name:      "time_ns",
		Help:      "Warp execution time",
	}, []string{"app", "token_id", "protocol"})
	fetched = prometheus.NewCounterVec(prometheus.CounterOpts{
		Namespace: "erlenmeyer",
		Subsystem: "exec",
		Name:      "fetched_datapoints",
		Help:      "Warp datapoint fetched",
	}, []string{"app", "token_id", "protocol"})
	operations = prometheus.NewCounterVec(prometheus.CounterOpts{
		Namespace: "erlenmeyer",
		Subsystem: "exec",
		Name:      "ops",
		Help:      "Warp ops",
	}, []string{"app", "token_id", "protocol"})
	requestAppErrorCounter = prometheus.NewCounterVec(prometheus.CounterOpts{
		Namespace: "erlenmeyer",
		Subsystem: "exec",
		Name:      "error_request",
		Help:      "Warp 10 error by user application",
	}, []string{"app", "token_id", "protocol"})
)

const (
	hugeNumberOfDatapoints = 100000
)

func init() {
	prometheus.MustRegister(requests)
	prometheus.MustRegister(times)
	prometheus.MustRegister(fetched)
	prometheus.MustRegister(operations)
	prometheus.MustRegister(requestAppErrorCounter)
}

// QueryResult The result of a Warp10 query
type QueryResult struct {
	Count   int             `json:"count"`
	Fetched int             `json:"fetched"`
	GTS     []GeoTimeSeries `json:"gts"`
}

// A GeoTimeSeries as returned by Warp10 (https://warp10.io/)
// WarpScript returns a JSON array: Format defined at
// https://www.warp10.io/content/03_Documentation/03_Interacting_with_Warp_10/04_Fetching_data/02_GTS_JSON_output_format
type GeoTimeSeries struct {
	// This is the class name of the Geo Time Serie. This may be present in only
	// one chunk of a given Geo Time Serie in the output.
	Class string `json:"c"`

	// This is an object containing the labels of the Geo Time Serie. This may be present in only
	// one chunk of a given Geo Time Serie in the output.

	Labels map[string]string `json:"l"`

	// This is an object containing the attributes (key/value) of the Geo Time Serie.
	// This may be present in only one chunk of a given Geo Time Serie in the output.

	Attrs map[string]string `json:"a"`

	// This is an id which is unique per Geo Time Serie in the output (but not across outputs).
	// All chunks of a given Geo Time Serie will have the same id and can therefore easily be identified and merged.

	ID string `json:"i"`

	// Array of Geo Time Series readings.
	// Each reading is itself an array containing 2, 3, 4 or 5 elements.
	// The first element of the array is the timestamp of the reading in microseconds since the Unix Epoch.
	// The last element of the array is the value of the reading, the type of this element varies with the type of the reading.
	// When the reading array has 3 elements, the second element is the elevation of the reading, in millimeters.
	// When the reading array has 4 elements, the second and third elements are the latitude and longitude of the reading.
	// When the reading array has 5 elements, the second and third elements are the latitude and longitude of the reading
	// and the fourth is its elevation.

	Values [][]interface{} `json:"v"`
}

// Warp10Server is the abstraction of Warp10
type Warp10Server interface {
	QueryGTS(body string) (*QueryResult, error)
	Query(body string) (*http.Response, error)
}

// HTTPWarp10Server Concrete implementation
type HTTPWarp10Server struct {
	Endpoint string
	Protocol string
	tokens   map[string]interface{}
	redirect map[string]string
}

// QueryGTS is a simple query, given the metric name and tags, and the start/end timestamps
func (server *HTTPWarp10Server) QueryGTS(body, txn string) (*QueryResult, error) {
	warp10Resp, err := server.Query(body, txn)
	if err != nil {
		return nil, err
	}
	defer warp10Resp.Body.Close()
	// Einstein dumps the current stack (an array) as the result of the program execution.
	// It contains a single element, which is an array of query results.
	result := []QueryResult{}
	err = json.NewDecoder(warp10Resp.Body).Decode(&result)
	if err != nil {
		return nil, err
	}
	return &(result[0]), nil
}

// QueryGTSs is multiple query, given the metric name and tags, and the start/end timestamps
func (server *HTTPWarp10Server) QueryGTSs(body, txn string) ([][]GeoTimeSeries, error) {
	warp10Resp, err := server.Query(body, txn)
	if err != nil {
		return nil, err
	}
	defer func() {
		if err := warp10Resp.Body.Close(); err != nil {
			log.
				WithError(err).
				Error("Cannot close Warp10 response body")
		}
	}()

	result := [][]GeoTimeSeries{}
	err = json.
		NewDecoder(warp10Resp.Body).
		Decode(&result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// Query is performing a simple query, given the metric name and tags, and the start/end timestamps
func (server *HTTPWarp10Server) Query(body string, txn string) (*http.Response, error) {
	log.WithFields(log.Fields{
		"type": "query",
		"txn":  txn,
	}).Debug("Debug Query")

	log.Debug(body)

	application := "unknown"
	tokenID := "unknown"
	for _, token := range tokenRegex.FindAllString(body, -1) {
		token = token[1 : len(token)-1]
		if _, ok := server.tokens[token]; ok {
			log.WithFields(log.Fields{
				"token": token[:6],
			}).Warn("Unautorized token")
			return nil, errors.New("Unauthorized")
		}

		app, ok := tokens[token]
		if !ok {
			tokeninfo := fmt.Sprintf("'%s' TOKENINFO", token)
			tokeninfo += " DUP <% 'type' GET ISNULL %> <% DROP 'notoken' STOP %> IFT"
			tokeninfo += " DUP <% 'type' GET 'READ' == %> <% 'application' GET %> <% DROP 'write' %> IFTE"

			res, err := http.Post(server.Endpoint+"/api/v0/exec", "text/plain", strings.NewReader(tokeninfo))
			if err != nil {
				log.WithFields(log.Fields{
					"txn": txn,
				}).WithError(err).Warn("Fail get token information")
				continue
			}
			defer func() {
				if err := res.Body.Close(); err != nil {
					log.
						WithFields(log.Fields{
							"txn": txn,
						}).
						WithError(err).
						Warn("Fail to close body")
				}
			}()

			if res.StatusCode != http.StatusOK {
				b, _ := ioutil.ReadAll(res.Body)
				log.WithFields(log.Fields{
					"txn":  txn,
					"body": string(b),
				}).Warn("Fail get token information")
				continue
			}

			var appRes []string
			err = json.NewDecoder(res.Body).Decode(&appRes)
			if err != nil {
				log.WithFields(log.Fields{
					"txn": txn,
				}).WithError(err).Warn("Fail to decode token information")
			}
			app = appRes[0]

			tokens[token] = app
		}

		// skip write tokens
		if app == "write" {
			continue
		}
		// skip invalid tokens
		if app == "notoken" {
			continue
		}

		application = app
		tokenID = token[:10]
	}

	url := server.Endpoint
	if u, ok := server.redirect[application]; ok {
		url = u
	}

	out := strings.NewReader(body)
	warp10Resp, err := http.Post(url+"/api/v0/exec", "text/plain", out)
	if err != nil {
		wErr := ""
		if warp10Resp.StatusCode == http.StatusInternalServerError {
			wErr = warp10Resp.Header.Get("X-Warp10-Error-Message")
		}
		if wErr == "" {
			dump, err := httputil.DumpResponse(warp10Resp, true)
			if err == nil {
				wErr = string(dump)
			} else {
				wErr = "Unparsable error"
			}
		}
		return nil, errors.Wrapf(err, "WarpScript error: %s", wErr)
	}

	if warp10Resp.StatusCode == http.StatusInternalServerError && !strings.HasPrefix(server.Protocol, "warp") {
		wErr := warp10Resp.Header.Get("X-Warp10-Error-Message")
		if wErr == "" {
			dump, err := httputil.DumpResponse(warp10Resp, true)
			if err == nil {
				wErr = string(dump)
			} else {
				wErr = "Unparsable error"
			}
		}
		warp10Resp.Body = ioutil.NopCloser(bytes.NewBufferString(wErr))
	}

	if warp10Resp.StatusCode >= 300 {
		requestAppErrorCounter.With(prometheus.Labels{
			"token_id": tokenID,
			"app":      application,
			"protocol": server.Protocol,
		}).Inc()
	}

	requests.With(prometheus.Labels{
		"token_id": tokenID,
		"app":      application,
		"protocol": server.Protocol,
	}).Inc()

	elapsed, err := strconv.ParseFloat(warp10Resp.Header.Get("X-Warp10-Elapsed"), 64)
	if err != nil {
		elapsed = 0
	}
	times.With(prometheus.Labels{
		"token_id": tokenID,
		"app":      application,
		"protocol": server.Protocol,
	}).Add(elapsed)

	datapoints, err := strconv.ParseFloat(warp10Resp.Header.Get("X-Warp10-Fetched"), 64)
	if err != nil {
		datapoints = 0
	}
	fetched.With(prometheus.Labels{
		"token_id": tokenID,
		"app":      application,
		"protocol": server.Protocol,
	}).Add(datapoints)
	if datapoints > hugeNumberOfDatapoints {
		log.WithFields(log.Fields{
			"token_id":   tokenID,
			"app":        application,
			"datapoints": datapoints,
			"elapsed":    elapsed,
			"txn":        txn,
		}).Warn("killroy")
	}

	ops, err := strconv.ParseFloat(warp10Resp.Header.Get("X-Warp10-Ops"), 64)
	if err != nil {
		ops = 0
	}
	operations.With(prometheus.Labels{
		"token_id": tokenID,
		"app":      application,
		"protocol": server.Protocol,
	}).Add(ops)

	return warp10Resp, nil
}

// NewWarpServer is returning a new Warp server
func NewWarpServer(endpoint string, protocol string) *HTTPWarp10Server {
	t := viper.GetStringSlice("deny.tokens")
	tokens := make(map[string]interface{})
	for _, token := range t {
		tokens[token] = nil
	}

	redirect := viper.GetStringMapString("redirect")
	if redirect == nil {
		redirect = make(map[string]string)
	}

	return &HTTPWarp10Server{
		Endpoint: endpoint,
		tokens:   tokens,
		redirect: redirect,
		Protocol: protocol,
	}
}

// IsoTime is returning the right format for Warp
func IsoTime(t time.Time) string {
	return t.UTC().Format(time.RFC3339Nano)
}

var noTime = time.Time{}

// TimeAndValue Get the time and value at a given position
func (series *GeoTimeSeries) TimeAndValue(i int) (bool, time.Time, float64) {
	vals := series.Values[i]
	val, found := vals[len(vals)-1].(float64)
	if !found {
		return false, noTime, 0.0
	}
	musecs := int64(vals[0].(float64))
	t := time.Unix(musecs/1000000, (musecs%1000000)%10000000)

	return true, t, val
}

// Delete is handling /api/v0/delete in Warp
func (server *HTTPWarp10Server) Delete(token string, query string) error {
	// Checking the trailing char
	if server.Endpoint[len(server.Endpoint)-1:] != "/" {
		server.Endpoint += "/"
	}

	resource := fmt.Sprintf("%sapi/v0/delete?%s", server.Endpoint, query)
	req, _ := http.NewRequest("GET", resource, nil) // nolint: gas
	req.Header.Set("Content-Type", "text/plain")
	req.Header.Set("X-Warp10-Token", token)
	req.Header.Set("X-CityzenData-Token", token)
	warpResp, err := http.DefaultClient.Do(req)

	if err != nil {
		return err
	}

	defer warpResp.Body.Close()

	if warpResp.StatusCode != 200 {
		var body []byte
		body, err = ioutil.ReadAll(warpResp.Body)
		if err != nil {
			return errors.Wrapf(err, "infos: %+v", warpResp.Header)
		}
		return errors.New(string(body))
	}
	return nil
}

// Find is Simple Find, given the metric name and tags, and the start/end timestamps
func (server *HTTPWarp10Server) Find(token string, selector string) (*http.Response, error) {
	req, _ := http.NewRequest("GET", server.Endpoint+"/api/v0/find?selector="+selector, nil) // nolint: gas
	req.Header.Set("Content-Type", "text/plain")
	req.Header.Set("X-Warp10-Token", token)
	req.Header.Set("X-CityzenData-Token", token)
	warpResp, err := http.DefaultClient.Do(req)

	if err != nil {
		return nil, err
	}

	if warpResp.StatusCode != 200 {
		var body []byte
		body, err = ioutil.ReadAll(warpResp.Body)
		if err != nil {
			return nil, err
		}
		return nil, errors.New(string(body))
	}
	return warpResp, nil
}

// FindGTS is find, given the metric name and tags, and the start/end timestamps
func (server *HTTPWarp10Server) FindGTS(token string, selector string) (*QueryResult, error) {
	warpResp, err := server.Find(token, selector)
	if err != nil {
		return nil, err
	}
	defer warpResp.Body.Close()
	result := QueryResult{}
	parseClassAndLabels := regexp.MustCompile(`[{}]`)

	var keyValArray []string

	reader := bufio.NewReader(warpResp.Body)
	for {
		line, isPrefix, err := reader.ReadLine()
		if line == nil || err != nil {
			// End of body
			break
		}
		if isPrefix {
			// a very long line is not usual, continue
			continue
		}
		labels := map[string]string{}
		parseOne := parseClassAndLabels.Split(string(line), -1) // => ["class.name", "labela,...", "", "tag=b,..."]
		if len(parseOne) > 1 {
			for _, keyValStr := range strings.Split(parseOne[1], ",") { // => ["key=val", "key1=val1"]
				keyValArray = strings.Split(keyValStr, "=") // => ["key", "val"]

				if len(keyValArray) > 1 {
					labels[keyValArray[0]] = keyValArray[1]
				}
			}
		}

		result.GTS = append(result.GTS, GeoTimeSeries{
			Class:  parseOne[0],
			Attrs:  make(map[string]string),
			Labels: labels,
		})
	}
	return &(result), nil
}
