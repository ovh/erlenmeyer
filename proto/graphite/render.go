package graphite

import (
	"encoding/json"
	"net/http"

	"github.com/pkg/errors"

	"github.com/ovh/erlenmeyer/middlewares"
)

// Render graphite handler (/render)
// http://graphite-api.readthedocs.io/en/latest/api.html#the-render-api-render
func Render(w http.ResponseWriter, r *http.Request) {
	token, err := getToken(r)
	if err != nil {
		logWarn(r, http.StatusUnauthorized, err)
		respondWithError(w, http.StatusUnauthorized, err)
		return
	}

	q := new(RenderQuery)
	if err = q.Parse(r); err != nil {
		logWarn(r, http.StatusBadRequest, err)
		respondWithError(w, http.StatusBadRequest, err)
		return
	}

	gts := make([]GTS, 0)
	for _, target := range q.Target {
		ws, err := CreateRenderRequest(target, q.From, q.Until)
		if err != nil {
			logWarn(r, http.StatusBadRequest, err)
			respondWithError(w, http.StatusBadRequest, err)
			return
		}

		resp, err := execute(token, w.Header().Get(middlewares.TxnHeader), ws)
		if err != nil {
			logErr(r, http.StatusInternalServerError, errors.Wrap(err, "WarpScript request failed"))
			respondWithError(w, http.StatusInternalServerError, err)
			return
		}

		gtss := make([][]GTS, 0)
		if err = json.Unmarshal(resp, &gtss); err != nil {
			logErr(r, http.StatusInternalServerError, errors.Wrap(err, "Query result parsing error"))
			respondWithError(w, http.StatusInternalServerError, errors.New("Query response is invalid - something went wrong with this request"))
			return
		}

		gts = append(gts, gtss[0]...)
	}

	result, err := Format(gts, q.Format)
	if err != nil {
		logErr(r, http.StatusInternalServerError, err)
		respondWithError(w, http.StatusInternalServerError, err)
		return
	}

	if len(q.JSONP) != 0 {
		respondWithJsonp(w, http.StatusOK, result, q.JSONP)
	} else {
		respond(w, http.StatusOK, result)
	}
}
