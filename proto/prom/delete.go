package prom

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"

	log "github.com/sirupsen/logrus"

	"github.com/ovh/erlenmeyer/core"
	"github.com/prometheus/prometheus/promql"
	"github.com/spf13/viper"
)

// Delete matched series entirely from a Prometheus server
//URL query parameters:
// - match[]=<series_selector>: Repeated label matcher argument that selects the series to delete. At least one match[] argument must be provided.
// FIXME: handle Delete
func (p *QL) Delete(w http.ResponseWriter, r *http.Request) {

	token := core.RetrieveToken(r)
	if len(token) == 0 {
		respondWithError(w, errors.New("Not authorized, please provide a READ token"), http.StatusForbidden)
		return
	}

	r.ParseForm()
	if len(r.Form["match[]"]) == 0 {
		respondWithError(w, errors.New("no match[] parameter provided"), http.StatusUnprocessableEntity)
	}

	deletedSeries := 0

	for _, s := range r.Form["match[]"] {
		var classname string
		labels := make(map[string]string)
		matchers, err := promql.ParseMetricSelector(s)
		if err != nil {
			respondWithError(w, err, http.StatusUnprocessableEntity)
		}
		for _, matcher := range matchers {
			if matcher.Name == "__name__" {
				classname = fmt.Sprintf("%v", matcher.Value)
			} else {
				labels[fmt.Sprintf("%v", matcher.Name)] = fmt.Sprintf("%v", matcher.Value)
			}
		}
		deleteQuery := buildWarp10Selector(classname, labels)
		fmt.Fprint(deleteQuery, "&deleteall")
		warpServer := core.NewWarpServer(viper.GetString("warp_endpoint"), "prometheus-delete")
		err = warpServer.Delete(token, "?selector="+deleteQuery.String())
		if err != nil {
			log.WithFields(log.Fields{
				"query": deleteQuery.String(),
				"error": err.Error(),
			}).Error("Error deleting some GTS")
			respondWithError(w, err, http.StatusInternalServerError)
			return
		}
		deletedSeries++
	}
	res := prometheusDeleteResponse{
		Status: "success",
		Data: prometheusDeleteDataResponse{
			NumDeleted: deletedSeries,
		},
	}

	b, _ := json.Marshal(res)
	w.Write(b)
}

type prometheusDeleteResponse struct {
	Status string                       `json:"status"`
	Data   prometheusDeleteDataResponse `json:"data"`
}

type prometheusDeleteDataResponse struct {
	NumDeleted int `json:"numDeleted"`
}

func buildWarp10Selector(metricName string, tags map[string]string) *bytes.Buffer {
	buffer := bytes.NewBufferString("")
	buffer.WriteString(metricName)
	translateLabels(tags, buffer)
	log.Debug("Selector", buffer.String())
	return buffer
}
