package graphite

import (
	"encoding/json"
	"errors"
	"fmt"
	"sort"
	"strconv"
	"strings"
	"time"
)

// TreeJSON graphite response format
type TreeJSON struct {
	Text          string `json:"text"`
	Expandable    int    `json:"expandable"`
	Leaf          int    `json:"leaf"`
	ID            string `json:"id"`
	AllowChildren int    `json:"allowChildren"`
}

// GTS is the format which is returned by warp10
type GTS struct {
	ClassName string            `json:"c"`
	Labels    map[string]string `json:"l"`
	Values    [][]float64       `json:"v"`
}

// JSON structure definition
type JSON struct {
	Target     string      `json:"target"`
	DataPoints [][]float64 `json:"datapoints"`
}

// Dygraph structure definition
type Dygraph struct {
	Labels []string    `json:"labels"`
	Data   [][]float64 `json:"data"`
}

// Ricksaw structure definition
type Ricksaw struct {
	Target     string           `json:"target"`
	DataPoints []RicksawElement `json:"datapoints"`
}

// RicksawElement structure definition
type RicksawElement struct {
	X int64   `json:"x"`
	Y float64 `json:"y"`
}

func (s *GTS) toGraphiteLabels() string {
	keys := make([]string, 0)
	labels := make([]string, 0)

	for k := range s.Labels {
		keys = append(keys, k)
	}

	sort.Strings(keys)
	for _, k := range keys {
		labels = append(labels, fmt.Sprintf("%s=%s", k, s.Labels[k]))
	}

	return strings.Join(labels, ";")
}

// ToRaw convert GTS to graphite raw format
// http://graphite-api.readthedocs.io/en/latest/api.html#raw
func (s *GTS) ToRaw() string {
	first := s.Values[0][0] / 1000000
	last := s.Values[len(s.Values)-1][0] / 1000000
	step := (last - first) / float64(len(s.Values))

	values := make([]string, 0)
	for _, value := range s.Values {
		values = append(values, strconv.FormatFloat(value[1], 'f', 6, 64))
	}

	return fmt.Sprintf("%s;%s,%d,%d,%f|%s", s.ClassName, s.toGraphiteLabels(), int(first), int(last), step, strings.Join(values, ","))
}

// ToCSV convert GTS to graphite csv format
// http://graphite-api.readthedocs.io/en/latest/api.html#csv
func (s *GTS) ToCSV() []string {
	metrics := make([]string, 0)

	for _, value := range s.Values {
		t := time.Unix(int64(value[0]/1000000), int64(value[0])%1000000)

		metrics = append(metrics, fmt.Sprintf("%s;%s,%s,%f", s.ClassName, s.toGraphiteLabels(), t.Format("2006-01-02 15:04:05"), value[1]))
	}

	return metrics
}

// ToJSON convert GTS to graphite JSON format
// http://graphite-api.readthedocs.io/en/latest/api.html#json
func (s *GTS) ToJSON() JSON {
	className := s.ClassName
	labels := s.toGraphiteLabels()
	if len(labels) > 0 {
		className = fmt.Sprintf("%s;%s", className, labels)
	}

	j := JSON{
		Target:     className,
		DataPoints: make([][]float64, 0),
	}

	for _, value := range s.Values {
		j.DataPoints = append(j.DataPoints, []float64{value[1], float64(int64(value[0] / 1000000))})
	}

	return j
}

// ToDygraph convert GTS to graphite dygraph format
// http://graphite-api.readthedocs.io/en/latest/api.html#dygraph
func (s *GTS) ToDygraph() Dygraph {
	d := Dygraph{
		Labels: []string{
			"Time",
			fmt.Sprintf("%s;%s", s.ClassName, s.toGraphiteLabels()),
		},

		Data: make([][]float64, 0),
	}

	for _, value := range s.Values {
		d.Data = append(d.Data, []float64{float64(int64(value[0] / 1000)), value[1]})
	}

	return d
}

// ToRicksaw convert GTS to graphite ricksaw format
// http://graphite-api.readthedocs.io/en/latest/api.html#rickshaw
func (s *GTS) ToRicksaw() Ricksaw {
	r := Ricksaw{
		Target:     fmt.Sprintf("%s;%s", s.ClassName, s.toGraphiteLabels()),
		DataPoints: []RicksawElement{},
	}

	for _, value := range s.Values {
		r.DataPoints = append(r.DataPoints, RicksawElement{
			X: int64(value[0] / 1000000),
			Y: value[1],
		})
	}

	return r
}

// Format GTS to desire format
// nolint: gocyclo
func Format(gtss []GTS, format string) ([]byte, error) {
	switch format {
	case "raw":
		raws := make([]string, 0)

		for _, gts := range gtss {
			raws = append(raws, gts.ToRaw())
		}

		return []byte(strings.Join(raws, "\n")), nil
	case "csv":
		csv := make([]string, 0)

		for _, gts := range gtss {
			csv = append(csv, gts.ToCSV()...)
		}

		return []byte(strings.Join(csv, "\n")), nil
	case "json":
		j := make([]JSON, 0)
		for _, gts := range gtss {
			j = append(j, gts.ToJSON())
		}

		return json.Marshal(j)
	case "dygraph":
		d := make([]Dygraph, 0)
		for _, gts := range gtss {
			d = append(d, gts.ToDygraph())
		}

		return json.Marshal(d)
	case "ricksaw":
		r := make([]Ricksaw, 0)
		for _, gts := range gtss {
			r = append(r, gts.ToRicksaw())
		}

		return json.Marshal(r)
	}

	return nil, errors.New("The output format is not supported")
}
