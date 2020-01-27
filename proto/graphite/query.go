package graphite

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"regexp"
	"strconv"
	"strings"
)

const (
	contentType = "Content-Type"
	mimeJSON    = "application/json"
	mimeForm    = "application/x-www-form-urlencoded"
)

// match template query parameters
var templateMatcher = regexp.MustCompile(`template\[(.+?)\]`)

// Parser interface in order to mutualize development
type Parser interface {
	Parse(*http.Request) error
}

// FindQuery which is used by the /metrics/find path
// http://graphite-api.readthedocs.io/en/latest/api.html#metrics-find
type FindQuery struct {
	Query     string `json:"query" description:"The query to search for."`
	Format    string `json:"format" description:"The output format to use."`
	Wildcards int    `json:"wildcards" description:"Whether to add a wildcard result at the end or no."`
	From      int64  `json:"from" description:"Epoch timestamp from which to consider metrics."`
	Until     int64  `json:"until" description:"Epoch timestamp until which to consider metrics."`
	JSONP     string `json:"jsonp" description:"Wraps the response in a JSONP callback"`
}

// Parse method is an implementation of Parser
// nolint: gocyclo
func (s *FindQuery) Parse(req *http.Request) error {
	var err error

	switch req.Header.Get(contentType) {
	case mimeJSON:
		body, err := ioutil.ReadAll(req.Body)
		if err != nil {
			return err
		}

		if err = json.Unmarshal(body, s); err != nil {
			return err
		}
	case mimeForm:
		if err = req.ParseForm(); err != nil {
			return err
		}

		for key, val := range req.Form {
			switch key {
			case "query":
				s.Query = val[0]
			case "format":
				s.Format = val[0]
			case "jsonp":
				s.JSONP = val[0]
			case "wildcards":
				s.Wildcards, err = strconv.Atoi(val[0])
				if err != nil {
					return err
				}
			case "from":
				s.From, err = strconv.ParseInt(val[0], 10, 64)
				if err != nil {
					return err
				}
			case "until":
				s.Until, err = strconv.ParseInt(val[0], 10, 64)
				if err != nil {
					return err
				}
			}
		}
	}

	params := req.URL.Query()
	if len(params.Get("query")) != 0 {
		s.Query = params.Get("query")
	}

	if len(s.Format) != 0 {
		s.Format = params.Get("format")
	}

	if len(params.Get("jsonp")) != 0 {
		s.JSONP = params.Get("jsonp")
	}

	if len(params.Get("wildcards")) != 0 {
		s.Wildcards, err = strconv.Atoi(params.Get("wildcards"))
		if err != nil {
			return err
		}
	}

	if len(params.Get("from")) != 0 {
		s.From, err = strconv.ParseInt(params.Get("from"), 10, 64)
		if err != nil {
			return err
		}
	}

	if len(params.Get("until")) != 0 {
		s.Until, err = strconv.ParseInt(params.Get("until"), 10, 64)
		if err != nil {
			return err
		}
	}

	if len(params.Get("until")) != 0 {
		s.Until, err = strconv.ParseInt(params.Get("until"), 10, 64)
		if err != nil {
			return err
		}
	}

	return nil
}

// ExpandQuery which is used by the /metrics/expand path
// http://graphite-api.readthedocs.io/en/latest/api.html#metrics-expand
type ExpandQuery struct {
	Query       []string `json:"query" description:"The metrics query. Can be specified multiple times."`
	GroupByExpr int      `json:"groupByExpr" description:"Whether to return a flat list of results or group them by query."`
	LeavesOnly  int      `json:"leavesOnly" description:"Whether to only return leaves or both branches and leaves."`
	JSONP       string   `json:"jsonp" description:"Wraps the response in a JSONP callback"`
}

// Parse method is an implementation of Parser
// nolint: gocyclo
func (s *ExpandQuery) Parse(req *http.Request) error {
	var err error

	switch req.Header.Get(contentType) {
	case mimeJSON:
		body, err := ioutil.ReadAll(req.Body)
		if err != nil {
			return err
		}

		if err = json.Unmarshal(body, s); err != nil {
			return err
		}
	case mimeForm:
		if err = req.ParseForm(); err != nil {
			return err
		}

		for key, val := range req.Form {
			switch key {
			case "query":
				s.Query = val
			case "jsonp":
				s.JSONP = val[0]
			case "groupByExpr":
				s.GroupByExpr, err = strconv.Atoi(val[0])
				if err != nil {
					return err
				}
			case "leavesOnly":
				s.LeavesOnly, err = strconv.Atoi(val[0])
				if err != nil {
					return err
				}
			}
		}
	}

	params := req.URL.Query()
	s.Query = append(s.Query, params["query"]...)
	if len(params.Get("jsonp")) != 0 {
		s.JSONP = params.Get("jsonp")
	}

	if len(params.Get("groupByExpr")) != 0 {
		s.GroupByExpr, err = strconv.Atoi(params.Get("groupByExpr"))
		if err != nil {
			return err
		}
	}

	if len(params.Get("leavesOnly")) != 0 {
		s.LeavesOnly, err = strconv.Atoi(params.Get("leavesOnly"))
		if err != nil {
			return err
		}
	}

	return nil
}

// IndexQuery which is used by the /metrics/index.json path
// http://graphite-api.readthedocs.io/en/latest/api.html#metrics-index-json
type IndexQuery struct {
	JSONP string `json:"jsonp" description:"Wraps the response in a JSONP callback"`
}

// Parse method is an implementation of Parser
func (s *IndexQuery) Parse(req *http.Request) error {
	switch req.Header.Get(contentType) {
	case mimeJSON:
		body, err := ioutil.ReadAll(req.Body)
		if err != nil {
			return err
		}

		if err := json.Unmarshal(body, s); err != nil {
			return err
		}
	case mimeForm:
		if err := req.ParseForm(); err != nil {
			return err
		}

		for key, val := range req.Form {
			switch key {
			case "jsonp":
				s.JSONP = val[0]
			}
		}
	}

	params := req.URL.Query()
	if len(params.Get("jsonp")) != 0 {
		s.JSONP = params.Get("jsonp")
	}

	return nil
}

// RenderQuery which is used by the /render path
// http://graphite-api.readthedocs.io/en/latest/api.html#the-render-api-render
type RenderQuery struct {
	Target []string `json:"target" description:"The target parameter specifies a path identifying one or several metrics, optionally with functions acting on those metrics"`
	From   string   `json:"from" description:"time period to graph from specifies the beginning, ..."`
	Until  string   `json:"until" description:"... until specifies the end"`
	Format string   `json:"format" description:"Controls the format of data returned Affects all targets passed in the URL."`
	JSONP  string   `json:"jsonp" description:"Wraps the response in a JSONP callback"`
}

// Parse method is an implementation of Parser
// nolint: gocyclo
func (s *RenderQuery) Parse(req *http.Request) error {
	template := make(map[string]string)
	switch req.Header.Get(contentType) {
	case mimeJSON:
		body, err := ioutil.ReadAll(req.Body)
		if err != nil {
			return err
		}

		if err = json.Unmarshal(body, s); err != nil {
			return err
		}
	case mimeForm:
		if err := req.ParseForm(); err != nil {
			return err
		}

		for key, val := range req.Form {
			switch key {
			case "target":
				s.Target = val
			case "format":
				s.Format = val[0]
			case "jsonp":
				s.JSONP = val[0]
			case "from":
				s.From = val[0]
			case "until":
				s.Until = val[0]
			}

			if templateMatcher.MatchString(key) {
				name := templateMatcher.FindStringSubmatch(key)

				template[name[1]] = val[0]
			}
		}
	}

	params := req.URL.Query()
	s.Target = append(s.Target, params["target"]...)
	if len(s.Format) != 0 {
		s.Format = params.Get("format")
	}

	if len(params.Get("jsonp")) != 0 {
		s.JSONP = params.Get("jsonp")
	}

	if len(params.Get("from")) != 0 {
		s.From = params.Get("from")
	}

	if len(params.Get("until")) != 0 {
		s.Until = params.Get("until")
	}

	for k, v := range params {
		if templateMatcher.MatchString(k) {
			name := templateMatcher.FindStringSubmatch(k)

			template[name[1]] = v[0]
		}
	}

	// replace template
	for i, target := range s.Target {
		for name, value := range template {
			if strings.Contains(target, "$"+name) {
				target = strings.Replace(target, "$"+name, value, -1)
			}
		}

		s.Target[i] = target
	}

	// set default values
	if len(s.Format) == 0 {
		s.Format = "json"
	}

	if len(s.From) == 0 {
		s.From = "-24h"
	}

	if len(s.Until) == 0 {
		s.Until = "now" // nolint: goconst
	}

	return nil
}
