package opentsdb

import (
	"encoding/json"
	"net/http"

	"github.com/ovh/erlenmeyer/core"
)

// Filter OpenTSDB filters
type Filter struct {
	Examples    string     `json:"examples"`
	Description string     `json:"description"`
	Translator  Translator `json:"-"`
}

var filters = map[string]Filter{
	WildCard: {
		Examples:    "host=wildcard(web*),  host=wildcard(web*.tsdb.net)  {\"type\":\"wildcard\",\"tagk\":\"host\",\"filter\":\"web*.tsdb.net\",\"groupBy\":false}",
		Description: "Performs pre, post and in-fix glob matching of values. The globs are case sensitive and multiple wildcards can be used. The wildcard character is the * (asterisk). At least one wildcard must be present in the filter value. A wildcard by itself can be used as well to match on any value for the tag key.",
		Translator:  wildcardTranslator,
	},
	LiteralOr: {
		Examples:    "host=literal_or(web01),  host=literal_or(web01|web02|web03)  {\"type\":\"literal_or\",\"tagk\":\"host\",\"filter\":\"web01|web02|web03\",\"groupBy\":false}",
		Description: "Accepts one or more exact values and matches if the series contains any of them. Multiple values can be included and must be separated by the | (pipe) character. The filter is case sensitive and will not allow characters that TSDB does not allow at write time.",
		Translator:  litteralOrTranslator,
	},
	RegExp: {
		Examples:    "host=regexp(.*)  {\"type\":\"regexp\",\"tagk\":\"host\",\"filter\":\".*\",\"groupBy\":false}",
		Description: "Provides full, POSIX compliant regular expression using the built in Java Pattern class. Note that an expression containing curly braces {} will not parse properly in URLs. If the pattern is not a valid regular expression then an exception will be raised.",
		Translator:  rawRegularExpressionTranslator,
	},
}

// HandleConfigFilters Filters handling
func (c *OpenTSDB) HandleConfigFilters(responseWriter http.ResponseWriter, request *http.Request) {

	token := core.RetrieveToken(request)
	if len(token) == 0 {
		c.WarnCounter.Inc()
		http.Error(responseWriter, "Not authorized", 401)
		return
	}
	responseWriter.Header().Add("Content-Type", "application/json")
	responseWriter.WriteHeader(200)
	json.NewEncoder(responseWriter).Encode(filters)
}
