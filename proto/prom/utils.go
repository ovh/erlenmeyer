package prom

import (
	"bytes"
	"fmt"
	"regexp"
	"strconv"
	"strings"

	"github.com/ovh/erlenmeyer/core"
)

func contains(s []string, e string) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}

func unique(slice []string) []string {
	uslice := make([]string, 0)
	for _, v := range slice {
		if !contains(uslice, v) {
			uslice = append(uslice, v)
		}
	}

	return uslice
}

func warpMetricstoPrometheus(gts core.GeoTimeSeries) prometheusResultResponse {
	var p prometheusResultResponse
	var v string

	// skip Warp10 series name based on PromQL result
	keepName := true
	if removeName, ok := gts.Attrs[core.ShouldRemoveNameLabel]; ok {
		skipName, err := strconv.ParseBool(removeName)
		if err == nil && skipName {
			keepName = false
		}
	}

	p.Metric = gts.Labels

	// use Warp10 series name for PromQL name
	if keepName {
		p.Metric["__name__"] = gts.Class
	}

	// Looping over values
	for _, value := range gts.Values {
		ts := value[0].(float64) // Casting as gts is an interface
		ts /= 1000000.0          // Moving from us to ms

		// handling value case, when lat,long are specified string can occure from PromQL values
		switch val := value[len(value)-1].(type) {
		case float64:
			v = fmt.Sprintf("%f", val)
			p.Values = append(p.Values, []interface{}{ts, v})
		case string:
			floatVal, err := strconv.ParseFloat(val, 64)
			if err == nil {
				v = fmt.Sprintf("%f", floatVal)
				p.Values = append(p.Values, []interface{}{ts, v})
			}
		}
	}
	return p
}

func warpScalarToPrometheus(gts core.GeoTimeSeries) prometheusResultResponse {
	var p prometheusResultResponse
	var v string

	// Looping over values
	for _, value := range gts.Values {
		ts := value[0].(float64) // Casting as gts is an interface
		ts /= 1000000.0          // Moving from us to ms

		// handling value case, when lat,long are specified string can occure from PromQL values
		switch val := value[len(value)-1].(type) {
		case float64:
			v = fmt.Sprintf("%f", val)
			p.Values = append(p.Values, []interface{}{ts, v})
		case string:
			floatVal, err := strconv.ParseFloat(val, 64)
			if err == nil {
				v = fmt.Sprintf("%f", floatVal)
				p.Values = append(p.Values, []interface{}{ts, v})
			}
		}
	}
	return p
}

func warpToPrometheusResponseInstant(gtss []core.GeoTimeSeries, resultType string) (prometheusDataResponse, error) {

	var resp prometheusDataResponse
	resp.ResultType = resultType
	resp.Result = make([]interface{}, 0)
	var v string

	if len(gtss) == 1 && gtss[0].Class == "scalar" {
		resp.ResultType = "scalar"
		for _, value := range gtss[0].Values {
			ts := value[0].(float64) // Casting as gts is an interface
			ts /= 1000000.0          // Moving from us to ms
			v = fmt.Sprintf("%v", value[1])
			resp.Result = append(resp.Result, []interface{}{ts, v})
		}
		return resp, nil
	}

	for _, gts := range gtss {
		resp.Result = append(resp.Result, warpMetricstoPrometheus(gts))
	}

	return resp, nil
}

func warpToPrometheusResponseRange(gtss []core.GeoTimeSeries, resultType string) (prometheusDataResponse, error) {

	var resp prometheusDataResponse
	resp.ResultType = resultType
	resp.Result = make([]interface{}, 0)

	for _, gts := range gtss {
		resp.Result = append(resp.Result, warpMetricstoPrometheus(gts))
	}
	return resp, nil
}

func translateLabels(tags map[string]string, buffer *bytes.Buffer) {
	tagsLen := len(tags)
	buffer.WriteString("{")
	for tagKey, tagValue := range tags {
		tagsLen--
		if strings.HasPrefix(tagValue, "~") {
			fmt.Fprintf(buffer, "%s%s", tagKey, tagValue)
		} else if tagValue == "*" {
			fmt.Fprintf(buffer, "%s~.*", tagKey)
		} else if strings.Contains(tagValue, "|") {
			fmt.Fprintf(buffer, "%s~", tagKey) // open regexp
			values := strings.Split(tagValue, "|")
			quotedValues := make([]string, len(values))
			for i, value := range values {
				quotedValues[i] = regexp.QuoteMeta(value)
			}
			fmt.Fprintf(buffer, "%s", strings.Join(quotedValues, "|"))
		} else {
			fmt.Fprintf(buffer, "%s=%s", tagKey, tagValue)
		}
		if tagsLen > 0 {
			buffer.WriteString(",")
		}
	}
	buffer.WriteString("}")
}
