package promremote

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
	"text/template"

	"github.com/gogo/protobuf/proto"
	"github.com/golang/snappy"
	"github.com/labstack/echo"
	"github.com/ovh/erlenmeyer/core"
	"github.com/prometheus/prometheus/prompb"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

const prometheusClassNameLabel = "__name__"

var (
	log *logrus.Entry
	tpl *template.Template
	ws  = `
[ {{.token}} {{.selector}} {{.labels}} {{.start}} ISO8601 {{.end}} ISO8601 ] FETCHDOUBLE
{{.operation}}
[ SWAP mapper.finite 0 0 0 ] MAP
{ '.app' NULL } RELABEL
`
)

// HandlerBuilder prom remote read request
func HandlerBuilder() echo.HandlerFunc {
	log = logrus.WithField("proto", "prom_remote_read")

	var err error
	tpl, err = template.New("ws").Parse(ws)
	if err != nil {
		log.
			WithError(err).
			Fatal("Invalid WS template")
	}

	return handler
}

func handler(c echo.Context) error {
	token := core.RetrieveToken(c.Request())
	if len(token) == 0 {
		return c.NoContent(http.StatusUnauthorized)
	}

	compressed, err := ioutil.ReadAll(c.Request().Body)
	if err != nil {
		log.Warn("Cannot read body")
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	reqBuf, err := snappy.Decode(nil, compressed)
	if err != nil {
		log.Warn("Cannot snappy decode")
		return c.JSON(http.StatusBadRequest, err.Error())

	}

	var req prompb.ReadRequest
	if err := proto.Unmarshal(reqBuf, &req); err != nil {
		log.Warn("Cannot unmarshal")
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	txn, ok := c.Get("txn").(string)
	if !ok {
		txn = ""
	}

	resp, err := read(token, &req, txn)
	if err != nil {
		log.
			WithError(err).
			Warn("Cannot read queries")
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	data, err := proto.Marshal(resp)
	if err != nil {
		log.Warn("Cannot marshal response payload")
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	c.Response().Header().Set("Content-Type", "application/x-protobuf")
	c.Response().Header().Set("Content-Encoding", "snappy")

	compressed = snappy.Encode(nil, data)
	if _, err := c.Response().Write(compressed); err != nil {
		log.Warn("Cannot write response to client")
		return c.JSON(http.StatusBadRequest, err.Error())
	}
	return nil
}

func read(token string, req *prompb.ReadRequest, txn string) (*prompb.ReadResponse, error) {
	wServer := core.NewWarpServer(viper.GetString("warp_endpoint"), "prometheus-remote-read")
	reqBody := &bytes.Buffer{}

	for _, q := range req.Queries {
		log.Debugf("Q: %+v", q)

		selector, labels := formatLabel(q.GetMatchers())

		err := tpl.Execute(reqBody, map[string]string{
			"token":     fmt.Sprintf("%q", token),
			"selector":  fmt.Sprintf("%q", selector),
			"labels":    labels,
			"start":     strconv.FormatInt(q.GetStartTimestampMs()*1000, 10),
			"end":       strconv.FormatInt(q.GetEndTimestampMs()*1000, 10),
			"operation": "", // FIXME: getOperation(q.GetHints(), q.GetEndTimestampMs()),
		})
		if err != nil {
			return nil, err
		}
	}
	log.Debugf("WS: %s", reqBody.String())

	resp := prompb.ReadResponse{
		Results: make([]*prompb.QueryResult, len(req.Queries)),
	}

	results, err := wServer.QueryGTSs(reqBody.String(), txn)
	if err != nil {
		return nil, err
	}
	log.Debugf("RESULT %+v", results)

	for i, result := range results {
		resp.Results[i] = &prompb.QueryResult{
			Timeseries: make([]*prompb.TimeSeries, len(result)),
		}

		for j, gts := range result {
			resp.Results[i].Timeseries[j] = gtsToPromTS(gts)
			log.Debugf("Matched GTS: %d (%d datapoints)", len(result), len(resp.Results[i].Timeseries[j].Samples))
		}
	}

	return &resp, nil
}

func getOperation(rh *prompb.ReadHints, nowMs int64) string {
	if rh == nil || rh.GetFunc() == "" {
		return ""
	}

	switch rh.GetFunc() {
	}
	return ""
}

func gtsToPromTS(gts core.GeoTimeSeries) *prompb.TimeSeries {
	ts := prompb.TimeSeries{
		Labels: []*prompb.Label{{
			Name:  prometheusClassNameLabel,
			Value: gts.Class,
		}},
		Samples: make([]*prompb.Sample, len(gts.Values)),
	}

	for k, v := range gts.Labels {
		ts.Labels = append(ts.Labels, &prompb.Label{
			Name:  k,
			Value: v,
		})
	}

	for i, point := range gts.Values {
		t, okA := point[0].(float64)
		if !okA {
			log.Warnf("Failed to assert ts|(%+v)", point[0])
			continue
		}

		v, okB := point[len(point)-1].(float64)
		if !okB {
			log.Warnf("Failed to assert dp|(%+v)", point[len(point)-1])
			continue
		}

		ts.Samples[i] = &prompb.Sample{
			Timestamp: int64(t / 1000),
			Value:     v,
		}
	}
	return &ts
}
