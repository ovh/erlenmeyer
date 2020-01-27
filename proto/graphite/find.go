package graphite

import (
	"encoding/json"
	"errors"
	"net/http"

	"github.com/ovh/erlenmeyer/middlewares"
)

// Find graphite handler (/metrics/find)
// http://graphite-api.readthedocs.io/en/latest/api.html#metrics-find
// https://github.com/graphite-project/graphite-web/blob/master/webapp/graphite/metrics/views.py#L55
func Find(w http.ResponseWriter, r *http.Request) {
	token, err := getToken(r)
	if err != nil {
		logWarn(r, http.StatusUnauthorized, err)
		respondWithError(w, http.StatusUnauthorized, err)
		return
	}

	q := new(FindQuery)
	if err = q.Parse(r); err != nil {
		logWarn(r, http.StatusBadRequest, err)
		respondWithError(w, http.StatusBadRequest, err)
		return
	}

	ws, err := CreateFindRequest(q.Query, q.Wildcards == 1)
	if err != nil {
		logWarn(r, http.StatusBadRequest, err)
		respondWithError(w, http.StatusBadRequest, err)
		return
	}

	resp, err := execute(token, w.Header().Get(middlewares.TxnHeader), ws)
	if err != nil {
		logErr(r, http.StatusInternalServerError, err)
		respondWithError(w, http.StatusInternalServerError, err)
		return
	}

	stack := make([][]TreeJSON, 0)
	if err = json.Unmarshal(resp, &stack); err != nil {
		logErr(r, http.StatusInternalServerError, err)
		respondWithError(w, http.StatusInternalServerError, errors.New(err.Error()+": "+string(resp)))
		return
	}

	result, err := json.Marshal(stack[0])
	if err != nil {
		logErr(r, http.StatusInternalServerError, err)
		respondWithError(w, http.StatusInternalServerError, err)
		return
	}

	if len(q.JSONP) != 0 {
		respondWithJsonp(w, http.StatusOK, result, q.JSONP)
	} else {
		w.Header().Set("Content-Type", "application/json;charset=utf-8")
		respond(w, http.StatusOK, result)
	}
}
