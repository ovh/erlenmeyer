package opentsdb

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"regexp"
	"strings"

	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"

	"github.com/ovh/erlenmeyer/core"
)

func translateTags(tags map[string]string, buffer *bytes.Buffer) []string {
	groupingTags := []string{}
	tagsLen := len(tags)
	buffer.WriteString("{")
	for tagKey, tagValue := range tags {
		tagsLen--
		if tagValue == "*" {
			fmt.Fprintf(buffer, "%s~.*", tagKey)
			groupingTags = append(groupingTags, tagKey)
		} else if strings.Contains(tagValue, "|") {
			fmt.Fprintf(buffer, "%s~", tagKey) // open regexp
			values := strings.Split(tagValue, "|")
			quotedValues := make([]string, len(values))
			for i, value := range values {
				quotedValues[i] = regexp.QuoteMeta(value)
			}
			fmt.Fprintf(buffer, "%s", strings.Join(quotedValues, "|"))
			groupingTags = append(groupingTags, tagKey)
		} else {
			fmt.Fprintf(buffer, "%s=%s", tagKey, tagValue)
		}
		if tagsLen > 0 {
			buffer.WriteString(",")
		}
	}
	buffer.WriteString("}")
	return groupingTags
}

func (query *QueryRequest) buildDeleteQueryString(metricName string, tags map[string]string) (*bytes.Buffer, []string) {
	buffer := bytes.NewBufferString("selector=")
	buffer.WriteString(metricName)
	groupingTags := translateTags(tags, buffer)
	if query.NoTimeRange {
		fmt.Fprint(buffer, "&deleteall")
	} else {
		fmt.Fprintf(buffer, "&start=%s", core.IsoTime(query.Start.Time))
		fmt.Fprintf(buffer, "&end=%s", core.IsoTime(query.End.Time))
	}
	return buffer, groupingTags
}

func executeDelete(w http.ResponseWriter, token string, query *QueryRequest) {

	responses := []*QueryResponse{}
	for _, subquery := range query.Queries {
		deleteQuery, groupingTags := query.buildDeleteQueryString(subquery.Metric, subquery.Tags)
		warpServer := core.NewWarpServer(viper.GetString("warp_endpoint"), "opentsdb-delete")
		err := warpServer.Delete(token, deleteQuery.String())
		if err != nil {
			log.WithFields(log.Fields{
				"query": deleteQuery.String(),
				"proto": "opentsdb",
				"error": err.Error(),
			}).Error("Error deleting some GTS")
		} else {
			rep := &QueryResponse{
				Metric:        subquery.Metric,
				Tags:          subquery.Tags,
				AggregateTags: groupingTags,
				DPs:           map[string]float64{},
			}
			responses = append(responses, rep)
		}
	}

	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(200)
	json.NewEncoder(w).Encode(responses)
}
