package prom

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httputil"
	"strings"

	"github.com/ovh/erlenmeyer/core"
	"github.com/ovh/erlenmeyer/middlewares"

	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/common/model"

	"github.com/ovh/erlenmeyer/proto/prom/promql"
)

// QL is the underlying struct to handle PromQL
type QL struct {
	QueryEngine *promql.Engine

	ReqCounter prometheus.Counter
}

// GetReqCounter satisfies the protocol interface
func (p *QL) GetReqCounter() prometheus.Counter {
	return p.ReqCounter
}

// NewPromQL returns a new Prometheus Server
func NewPromQL() *QL {

	c := QL{}

	// metrics
	c.ReqCounter = prometheus.NewCounter(prometheus.CounterOpts{
		Namespace: "erlenmeyer",
		Subsystem: "promql",
		Name:      "request",
		Help:      "Number of request handled.",
	})

	return &c
}

type status string

const (
	statusSuccess status = "success"
	statusError   status = "error"
)

type errorType string

const (
	errorNone     errorType = ""
	errorTimeout  errorType = "timeout"
	errorCanceled errorType = "canceled"
	errorExec     errorType = "execution"
	errorBadData  errorType = "bad_data"
)

type prometheusResponse struct {
	Status    status                 `json:"status"`
	Data      prometheusDataResponse `json:"data,omitempty"`
	ErrorType errorType              `json:"errorType,omitempty"`
	Error     string                 `json:"error,omitempty"`
}

type prometheusDataResponse struct {
	ResultType string        `json:"resultType"`
	Result     []interface{} `json:"result"`
}

type prometheusResultResponse struct {
	Metric map[string]string `json:"metric"`

	// Values are heterogen, hence the interface. One datapoint is looking like this:
	//[ 1435781460.781, "1" ]
	Values [][]interface{} `json:"values"`
}

// Context is holding the informations like token, start, end, and so on
type Context struct {
	core.Context
	Expr         promql.Expr
	Bucketizer   string
	Mapper       string
	MapperValue  string
	HasMapper    bool
	HasFunction  bool
	FunctionName string
	Args         []string
	IsInstant    bool
	hasAbsent    bool
}

// QueryRange evaluates an expression query over a range of time:
// URL query parameters:
// - query=<string>: Prometheus expression query string.
// - start=<rfc3339 | unix_timestamp>: Start timestamp.
// - end=<rfc3339 | unix_timestamp>: End timestamp.
// - step=<duration>: Query resolution step width.
func (p *QL) QueryRange(w http.ResponseWriter, r *http.Request) {

	token := core.RetrieveToken(r)
	if len(token) == 0 {
		respondWithError(w, errors.New("Not authorized, please provide a READ token"), http.StatusForbidden)
		return
	}

	context := Context{}
	var err error

	context.Start, err = core.ParsePromTime(r.FormValue("start"))
	if err != nil {
		log.WithFields(log.Fields{
			"error":   err.Error(),
			"proto":   "promql",
			"entity":  "start",
			"context": fmt.Sprintf("%+v", context),
		}).Error("Unprocessable entity")
		respondWithError(w, errors.New("Unprocessable Entity: start"), http.StatusBadRequest)
		return
	}

	context.End, err = core.ParsePromTime(r.FormValue("end"))
	if err != nil {
		log.WithFields(log.Fields{
			"error":   err.Error(),
			"entity":  "end",
			"proto":   "promql",
			"context": fmt.Sprintf("%+v", context),
		}).Error("Unprocessable entity")
		respondWithError(w, errors.New("Unprocessable Entity: start"), http.StatusBadRequest)
		return
	}

	if context.End.Before(context.Start) {
		log.WithFields(log.Fields{
			"error":   errors.New("end is before start"),
			"context": fmt.Sprintf("%+v", context),
			"proto":   "promql",
			"entity":  "start",
		}).Error("Unprocessable entity")
		respondWithError(w, errors.New("End is before start"), http.StatusBadRequest)
		return
	}

	context.Step, err = core.ParsePromDuration(r.FormValue("step"))
	if err != nil {
		log.WithFields(log.Fields{
			"error":   err.Error(),
			"entity":  "step",
			"proto":   "promql",
			"context": fmt.Sprintf("%+v", context),
		}).Error("Unprocessable entity")
		respondWithError(w, errors.New("Unprocessable Entity: step"), http.StatusBadRequest)
		return
	}

	if context.Step == "0 s" {
		respondWithError(w, errors.New("zero or negative query resolution step widths are not accepted. Try a positive integer"), http.StatusBadRequest)
		return
	}

	if context.Step == "" {
		context.Step = "5 m"
	}

	context.Query = r.FormValue("query")

	log.WithFields(log.Fields{
		"query":   context.Query,
		"proto":   "promql",
		"context": fmt.Sprintf("%+v", context),
	}).Debug("Evaluating query")

	context.Expr, err = promql.ParseExpr(context.Query)
	if err != nil {
		log.WithFields(log.Fields{
			"query":   context.Query,
			"proto":   "promql",
			"context": fmt.Sprintf("%+v", context),
			"err":     err,
		}).Debug("Bad query")
		respondWithError(w, err, http.StatusUnprocessableEntity)
		return
	}

	log.WithFields(log.Fields{
		"query":   context.Query,
		"proto":   "promql",
		"context": fmt.Sprintf("%+v", context),
	}).Debug("Query is OK")

	evaluator := evaluator{}
	tree := evaluator.GenerateQueryTree(context)

	mc2 := tree.ToWarpScriptWithTime(token, context.Query, context.Step, context.Start, context.End)

	log.WithFields(log.Fields{
		"query":  context.Query,
		"source": r.RemoteAddr,
		"proto":  "promql",
		"method": r.Method,
		"path":   r.URL.String(),
	}).Debug("PromQL query")

	warpServer := core.NewWarpServer(viper.GetString("warp_endpoint"), "prometheus-query-range")
	response, err := warpServer.Query(mc2, w.Header().Get(middlewares.TxnHeader))
	if err != nil {
		wErr := response.Header.Get("X-Warp10-Error-Message")
		if wErr == "" {
			dump, err := httputil.DumpResponse(response, true)
			if err == nil {
				wErr = string(dump)
			} else {
				wErr = "Unparsable error"
			}
		}
		log.WithFields(log.Fields{
			"error": fmt.Errorf(wErr),
			"proto": "promql",
		}).Error("Bad response from Egress: " + err.Error())
		respondWithError(w, fmt.Errorf(wErr), http.StatusServiceUnavailable)
		return
	}
	buffer, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.WithFields(log.Fields{
			"error": err.Error(),
			"proto": "promql",
		}).Error("can't fully read Egress response")
		respondWithError(w, err, http.StatusServiceUnavailable)
		return
	}

	// HACK : replace Infinity values from Warp to Inf
	s := strings.Replace(string(buffer), "Infinity", "+Inf", -1)
	s = strings.Replace(s, "-+Inf", "-Inf", -1)
	buffer = []byte(s)

	responses := [][]core.GeoTimeSeries{}
	err = json.Unmarshal(buffer, &responses)
	if err != nil {
		wErr := response.Header.Get("X-Warp10-Error-Message")
		if wErr == "" {
			dump, err := httputil.DumpResponse(response, true)
			if err == nil {
				wErr = string(dump)
			} else {
				wErr = "Unparsable error"
			}
		}
		log.WithFields(log.Fields{
			"error": fmt.Errorf(wErr),
			"proto": "promql",
		}).Error("Cannot unmarshal egress response: " + err.Error())
		respondWithError(w, fmt.Errorf(wErr), http.StatusServiceUnavailable)
		return
	}
	// Since it's a range_query, we can enforce the matrix resultType
	prometheusResponse, err := warpToPrometheusResponseRange(responses[0], model.ValMatrix.String())
	if err != nil {
		w.Write([]byte(err.Error()))
		respondWithError(w, err, http.StatusServiceUnavailable)
	}
	respond(w, prometheusResponse)
}

func respondWithError(w http.ResponseWriter, err error, statusCode int) {
	var resp prometheusResponse
	resp.Status = "error"
	switch statusCode {
	case http.StatusUnprocessableEntity:
		resp.ErrorType = errorBadData
	case http.StatusServiceUnavailable:
		resp.ErrorType = errorExec
	case http.StatusBadRequest:
		resp.ErrorType = errorExec
	}

	if strings.Contains(err.Error(), "in section [TOP] (MSGFAIL") {
		errors := strings.Split(err.Error(), "in section [TOP] (MSGFAIL ")
		if len(errors) > 0 {
			resp.Error = strings.TrimSuffix(errors[1], ")")
			// when their is a second )
			resp.Error = strings.TrimSuffix(resp.Error, ")")
		}
	} else {
		resp.Error = err.Error()
	}

	b, err := json.Marshal(&resp)
	if err != nil {
		panic(err)
	}
	w.WriteHeader(statusCode)
	w.Write(b)
	return
}

func respondFind(w http.ResponseWriter, data []map[string]string) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	b, err := json.Marshal(&prometheusFindResponse{
		Status: statusSuccess,
		Data:   data,
	})
	if err != nil {
		log.WithError(err).Error("cannot marshal 'prometheusFindResponse'")
		return
	}
	_, err = w.Write(b)
	if err != nil {
		log.WithError(err).Error("Cannot write body")
	}
}

func respond(w http.ResponseWriter, data prometheusDataResponse) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	b, err := json.Marshal(&prometheusResponse{
		Status: statusSuccess,
		Data:   data,
	})
	if err != nil {
		return
	}
	w.Write(b)
}

// InstantQuery evaluates an instant query at a single point in time.
// URL query parameters:
// - query=<string>: Prometheus expression query string.
// - time=<rfc3339 | unix_timestamp>: Evaluation timestamp. Optional.
// - timeout=<duration>: Evaluation timeout. Optional. Defaults to and is capped by the value of the -query.timeout flag.
func (p *QL) InstantQuery(w http.ResponseWriter, r *http.Request) {

	token := core.RetrieveToken(r)
	if len(token) == 0 {
		respondWithError(w, errors.New("Not authorized, please provide a READ token"), http.StatusForbidden)
		return
	}

	context := Context{}
	var err error

	if r.FormValue("time") == "" {
		core.Now()
	} else {
		context.End, err = core.ParsePromTime(r.FormValue("time"))
		if err != nil {
			log.WithFields(log.Fields{
				"error":   err.Error(),
				"proto":   "promql",
				"entity":  "start",
				"context": fmt.Sprintf("%+v", context),
			}).Error("Unprocessable entity")
			respondWithError(w, errors.New("Unprocessable Entity: start"), http.StatusBadRequest)
			return
		}
	}

	context.Query = r.FormValue("query")

	log.WithFields(log.Fields{
		"query":   context.Query,
		"proto":   "promql",
		"context": fmt.Sprintf("%+v", context),
	}).Debug("Evaluating query")

	context.Expr, err = promql.ParseExpr(context.Query)
	if err != nil {
		log.WithFields(log.Fields{
			"query":   context.Query,
			"proto":   "promql",
			"context": fmt.Sprintf("%+v", context),
			"err":     err,
		}).Error("Bad query")
		respondWithError(w, err, http.StatusUnprocessableEntity)
		return
	}

	log.WithFields(log.Fields{
		"query":   context.Query,
		"proto":   "promql",
		"context": fmt.Sprintf("%+v", context),
	}).Debug("Query is OK")

	evaluator := evaluator{}
	tree := evaluator.GenerateInstantQueryTree(context)

	mc2 := tree.ToWarpScriptWithTime(token, context.Query, context.Step, context.Start, context.End)

	log.WithFields(log.Fields{
		"query":   context.Query,
		"proto":   "promql",
		"context": fmt.Sprintf("%+v", context),
	}).Debug("warpscript generated")

	warpServer := core.NewWarpServer(viper.GetString("warp_endpoint"), "prometheus-query-instant")
	response, err := warpServer.Query(mc2, w.Header().Get(middlewares.TxnHeader))
	if err != nil {
		wErr := response.Header.Get("X-Warp10-Error-Message")
		if wErr == "" {
			dump, err := httputil.DumpResponse(response, true)
			if err == nil {
				wErr = string(dump)
			} else {
				wErr = "Unparsable error"
			}
		}
		log.WithFields(log.Fields{
			"error": fmt.Errorf(wErr),
			"proto": "promql",
		}).Error("Bad response from Egress: " + err.Error())
		respondWithError(w, fmt.Errorf(wErr), http.StatusServiceUnavailable)
		return
	}
	buffer, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.WithFields(log.Fields{
			"error": err.Error(),
			"proto": "promql",
		}).Error("can't fully read Egress response")
		respondWithError(w, err, http.StatusServiceUnavailable)
		return
	}
	responses := [][]core.GeoTimeSeries{}
	err = json.Unmarshal(buffer, &responses)
	if err != nil {
		wErr := response.Header.Get("X-Warp10-Error-Message")
		if wErr == "" {
			dump, err := httputil.DumpResponse(response, true)
			if err == nil {
				wErr = string(dump)
			} else {
				wErr = "Unparsable error"
			}
		}
		log.WithFields(log.Fields{
			"error": fmt.Errorf(wErr),
			"proto": "promql",
		}).Error("Cannot unmarshal egress response: " + err.Error())
		respondWithError(w, fmt.Errorf(wErr), http.StatusServiceUnavailable)
		return
	}
	prometheusResponse, err := warpToPrometheusResponseInstant(responses[0], context.Expr.String())
	if err != nil {
		w.Write([]byte(err.Error()))
		respondWithError(w, err, http.StatusServiceUnavailable)
	}
	respond(w, prometheusResponse)
}
