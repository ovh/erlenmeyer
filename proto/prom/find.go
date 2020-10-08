package prom

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"

	queryPromql "github.com/ovh/erlenmeyer/proto/prom/promql"

	"github.com/ovh/erlenmeyer/core"
)

// FindSeries returns the list of time series that match a certain label set.
func (p *QL) FindSeries(w http.ResponseWriter, r *http.Request) {
	token := core.RetrieveToken(r)
	if len(token) == 0 {
		respondWithError(w, errors.New("Not authorized, please provide a READ token"), http.StatusForbidden)
		return
	}

	r.ParseForm()
	if len(r.Form["match[]"]) == 0 {
		respondWithError(w, errors.New("no match[] parameter provided"), http.StatusUnprocessableEntity)
		return
	}

	resp := []map[string]string{}

	for _, s := range r.Form["match[]"] {
		classname := s
		labels := map[string]string{}

		matchers, err := queryPromql.ParseMetricSelector(classname)
		if err != nil {
			respondWithError(w, err, http.StatusUnprocessableEntity)
			return
		}

		for _, matcher := range matchers {
			if matcher.Name == "__name__" {
				classname = fmt.Sprintf("%v", matcher.Value)
			} else {
				labelsValue := matcher.Value

				if matcher.Type.String() == "=~" {
					labelsValue = "~" + labelsValue
				} else if matcher.Type.String() == "!=" || matcher.Type.String() == "!~" {
					labelsValue = fmt.Sprintf("~(?!%s).*", labelsValue)
				}
				labels[fmt.Sprintf("%v", matcher.Name)] = fmt.Sprintf("%v", labelsValue)
			}
		}

		findQuery := buildWarp10Selector(classname, labels)
		warpServer := core.NewWarpServer(viper.GetString("warp_endpoint"), "prometheus-find")
		gtss, err := warpServer.FindGTS(token, findQuery.String())

		if err != nil {
			log.WithFields(log.Fields{
				"query": findQuery.String(),
				"error": err.Error(),
			}).Error("Error finding some GTS")
			respondWithError(w, err, http.StatusInternalServerError)
			return
		}

		for _, gts := range gtss.GTS {
			data := make(map[string]string)
			data["__name__"] = gts.Class
			for key, value := range gts.Labels {
				if key == ".app" {
					continue
				}
				data[key] = value
			}
			for key, value := range gts.Attrs {
				data[key] = value
			}
			resp = append(resp, data)
		}
	}
	respondFind(w, resp)
}

type prometheusFindResponse struct {
	Status status              `json:"status"`
	Data   []map[string]string `json:"data"`
}

type prometheusFindLabelsResponse struct {
	Status string   `json:"status"`
	Data   []string `json:"data"`
}

// FindLabelsValues is handling finding labels values
func (p *QL) FindLabelsValues(ctx echo.Context) error {
	w := ctx.Response()
	r := ctx.Request()

	token := core.RetrieveToken(r)
	if len(token) == 0 {
		respondWithError(w, errors.New("Not authorized, please provide a READ token"), http.StatusForbidden)
		return nil
	}

	labelValue := ctx.Param("label")

	if len(labelValue) == 0 {
		log.WithFields(log.Fields{}).Error("missing label")
		respondWithError(w, errors.New("Unprocessable Entity: label"), http.StatusBadRequest)
		return nil
	}

	selector := "~.*{" + labelValue + "~.*}"

	warpServer := core.NewWarpServer(viper.GetString("warp_endpoint"), "prometheus-find-labels")
	gtss, err := warpServer.FindGTS(token, selector)
	if err != nil {
		log.WithFields(log.Fields{
			"query": selector,
			"error": err.Error(),
		}).Error("Error finding some GTS")
		respondWithError(w, err, http.StatusInternalServerError)
		return nil
	}

	var resp prometheusFindLabelsResponse

	resp.Status = "success"

	for _, gts := range gtss.GTS {
		for key, value := range gts.Labels {
			if key != labelValue {
				continue
			}
			resp.Data = append(resp.Data, value)
		}
	}

	resp.Data = unique(resp.Data)
	b, _ := json.Marshal(resp)
	w.Write(b)

	return nil
}

// FindAndDeleteSeries is handling /find and /delete for series
func (p *QL) FindAndDeleteSeries(w http.ResponseWriter, r *http.Request) {

	switch r.Method {
	case "DELETE":
		p.Delete(w, r)
	case "GET":
		p.FindSeries(w, r)
	}
}
