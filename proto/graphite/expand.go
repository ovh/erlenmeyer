package graphite

import (
	"encoding/json"
	"errors"
	"net/http"

	"github.com/ovh/erlenmeyer/middlewares"
)

// Expand graphite handler (/metrics/expand)
// http://graphite-api.readthedocs.io/en/latest/api.html#metrics-expand
// https://github.com/graphite-project/graphite-web/blob/master/webapp/graphite/metrics/views.py#L172
func Expand(w http.ResponseWriter, r *http.Request) {
	token, err := getToken(r)
	if err != nil {
		logWarn(r, http.StatusUnauthorized, err)
		respondWithError(w, http.StatusUnauthorized, err)
		return
	}

	q := new(ExpandQuery)
	if err = q.Parse(r); err != nil {
		logWarn(r, http.StatusBadRequest, err)
		respondWithError(w, http.StatusBadRequest, err)
		return
	}

	results := make([]string, 0)
	for _, query := range q.Query {
		ws, err := CreateExpandRequest(query)
		if err != nil {
			logWarn(r, http.StatusBadRequest, err)
			respondWithError(w, http.StatusBadRequest, err)
			return
		}

		resp, err := execute(token, w.Header().Get(middlewares.TxnHeader), ws)
		if err != nil {
			logErr(r, http.StatusInternalServerError, errors.New("WarpScript request failed"))
			respondWithError(w, http.StatusInternalServerError, err)
			return
		}

		stack := make([][]string, 0)
		if err := json.Unmarshal(resp, &stack); err != nil {
			logErr(r, http.StatusInternalServerError, err)
			respondWithError(w, http.StatusInternalServerError, err)
			return
		}

		results = append(results, stack[0]...)
	}

	response, err := json.Marshal(map[string][]string{"results": results})
	if err != nil {
		logErr(r, http.StatusInternalServerError, err)
		respondWithError(w, http.StatusInternalServerError, err)
		return
	}

	if len(q.JSONP) != 0 {
		respondWithJsonp(w, http.StatusOK, response, q.JSONP)
	} else {
		respond(w, http.StatusOK, response)
	}
}
