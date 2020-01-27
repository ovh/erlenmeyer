package opentsdb

import (
	"encoding/json"
	"net/http"

	"github.com/ovh/erlenmeyer/core"
)

// Translation from OpenTSDB aggregator to WarpScript reducer
var aggregatorToReduce = map[string]string{
	"avg":    "reducer.mean.exclude-nulls",
	"sum":    "reducer.sum",
	"zimsum": "reducer.sum",
	"min":    "reducer.min",
	"max":    "reducer.max",
	"mimmin": "reducer.min",
	"mimmax": "reducer.max",
	"p50":    "50.0 reducer.percentile",
	"p75":    "75.0 reducer.percentile",
	"p90":    "90.0 reducer.percentile",
	"p95":    "95.0 reducer.percentile",
	"p99":    "99.0 reducer.percentile",
	"p999":   "99.9 reducer.percentile",
	"dev":    "true reducer.sd",
	none:     "",
	"count":  "reducer.count.exclude-nulls",
}

// HandleAggregators Handle OpenTSDB aggregators
func (c *OpenTSDB) HandleAggregators(responseWriter http.ResponseWriter, request *http.Request) {
	token := core.RetrieveToken(request)
	if len(token) == 0 {
		c.WarnCounter.Inc()
		http.Error(responseWriter, "Not authorized", 401)
		return
	}
	responseWriter.Header().Add("Content-Type", "application/json")
	responseWriter.WriteHeader(200)
	json.NewEncoder(responseWriter).Encode(GetAggregators())
}

// GetAggregators returns the list of supported aggregators
func GetAggregators() []string {
	aggregators := []string{}
	for aggregator := range aggregatorToReduce {
		aggregators = append(aggregators, aggregator)
	}

	for aggregator := range downsamplerToBucketizer {
		if _, exists := aggregatorToReduce[aggregator]; !exists {
			aggregators = append(aggregators, aggregator)
		}
	}
	return aggregators
}
