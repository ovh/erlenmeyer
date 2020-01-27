package graphite

import (
	"encoding/json"
	"errors"
	"net/http"

	"github.com/ovh/erlenmeyer/core"
	"github.com/ovh/erlenmeyer/middlewares"
)

// Index graphite handler (/metrics/index.json)
// http://graphite-api.readthedocs.io/en/latest/api.html#metrics-index-json
// https://github.com/graphite-project/graphite-web/blob/master/webapp/graphite/metrics/views.py#L35
func Index(w http.ResponseWriter, r *http.Request) {
	token, err := getToken(r)
	if err != nil {
		logWarn(r, http.StatusUnauthorized, err)
		respondWithError(w, http.StatusUnauthorized, err)
		return
	}

	query := new(IndexQuery)
	if err = query.Parse(r); err != nil {
		logWarn(r, http.StatusBadRequest, err)
		respondWithError(w, http.StatusBadRequest, err)
		return
	}

	ws := core.NewEmptyNode()
	_, err = Parse("name(find(*))", "", "", ws)
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
	if err = json.Unmarshal(resp, &stack); err != nil {
		logErr(r, http.StatusInternalServerError, err)
		respondWithError(w, http.StatusInternalServerError, err)
		return
	}

	result, err := json.Marshal(stack[0])
	if err != nil {
		logErr(r, http.StatusInternalServerError, err)
		respondWithError(w, http.StatusInternalServerError, err)
		return
	}

	if len(query.JSONP) != 0 {
		respondWithJsonp(w, http.StatusOK, result, query.JSONP)
	} else {
		respond(w, http.StatusOK, result)
	}
}
