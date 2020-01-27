package opentsdb

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"regexp"
	"sort"
	"strings"

	"github.com/ovh/erlenmeyer/core"
	"github.com/spf13/viper"
)

// SuggestQuery is the OpenTSDB suggest query input
type SuggestQuery struct {
	Type string `schema:"type" json:"type"` // "metric", "tagk" or "tagv"
	Q    string `schema:"q"    json:"q,omitempty"`
	Max  int    `schema:"max"  json:"max,omitempty"`
}

func buildWarpScript(suggest *SuggestQuery) (string, error) {

	switch suggest.Type {
	case "metrics":
		// Metric name regexp
		return "~.*" + regexp.QuoteMeta(suggest.Q) + ".*{}", nil
	case "tagk":
		return "~.*{}", nil

	case "tagv":
		return "~.*{}", nil

	default:
		return "", errors.New("Invalid type value (must be 'metrics', 'tagk' or 'tagv')")
	}
}

// HandleSuggest main /suggest handler
// nolint: gocyclo
func (c *OpenTSDB) HandleSuggest(w http.ResponseWriter, r *http.Request) {

	token := core.RetrieveToken(r)
	if len(token) == 0 {
		c.WarnCounter.Inc()
		http.Error(w, "Not authorized", 401)
		return
	}

	var suggest = &SuggestQuery{}
	err := DecodeRequestParams(w, r, suggest)
	if err != nil {
		c.WarnCounter.Inc()
		return // Response was sent
	}

	// Default value (according to OpenTSDB API)
	if suggest.Max == 0 {
		suggest.Max = 25
	}

	selector, err := buildWarpScript(suggest)
	if err != nil {
		c.ErrCounter.Inc()
		http.Error(w, err.Error(), 500)
		return
	}

	warpServer := core.NewWarpServer(viper.GetString("warp_endpoint"), "opentsdb-suggest")
	result, err := warpServer.FindGTS(token, selector)

	if err != nil {
		c.ErrCounter.Inc()
		http.Error(w, err.Error(), 500)
		return
	}

	// Collect all metric names in a Map
	names := make(map[string]bool)

	switch suggest.Type {
	case "metrics":
		for _, gts := range result.GTS {
			names[gts.Class] = true
		}

	case "tagk":
		for _, gts := range result.GTS {
			delete(gts.Labels, ".app")
			for name := range gts.Labels {
				if strings.HasPrefix(name, suggest.Q) {
					names[name] = true
				}
			}
		}

	case "tagv":
		for _, gts := range result.GTS {
			delete(gts.Labels, ".app")
			for _, value := range gts.Labels {
				if strings.HasPrefix(value, suggest.Q) {
					names[value] = true
				}
			}
		}
	}

	names2 := make([]string, len(names))
	i := 0
	for name := range names {
		names2[i] = name
		i++
	}
	sort.Strings(names2)

	if len(names2) > suggest.Max {
		names2 = names2[:suggest.Max]
	}

	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(200)
	json.NewEncoder(w).Encode(names2)
}

// DecodeRequestParams parse a request into a query structure. This handles the fact that OpenTSDB
// accepts request parameters both at GET query string parameters and as a
// POSTed JSON structure
func DecodeRequestParams(w http.ResponseWriter, r *http.Request, dst interface{}) error {
	var err error
	var reason string

	if r.Method == "GET" {
		err = OpenTsdbDecoder.Decode(dst, r.URL.Query())
	} else {
		defer r.Body.Close()
		body, err := ioutil.ReadAll(r.Body)
		if err != nil {
			reason = fmt.Sprintf("can't fully read the request body: %s", err)
		}

		if err := json.Unmarshal(body, dst); err != nil {
			reason = fmt.Sprintf("can't parse query: %s", err)
		}
	}

	if err != nil {

		w.WriteHeader(http.StatusBadRequest)
		w.Header().Add("Content-Type", "text/plain")
		w.Write([]byte(reason))
		return err
	}

	validatable, found := dst.(Validatable)
	if found {
		if err := validatable.Validate(); err != nil {
			// OpenTSDB returns a 400 for validation errors even if
			// although 422 (Unprocessable entity) would be more suitable
			w.WriteHeader(http.StatusBadRequest)
			w.Header().Add("Content-Type", "text/plain")
			w.Write([]byte(err.Error()))

			return err
		}
	}

	return nil
}
