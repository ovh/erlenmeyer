package graphite

import (
	"bytes"
	"errors"
	"io/ioutil"
	"net/http"
	"net/http/httputil"

	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"

	"github.com/ovh/erlenmeyer/core"
)

func execute(token, txn string, tree *core.Node) ([]byte, error) {
	server := core.NewWarpServer(viper.GetString("warp_endpoint"), "graphite-query")
	resp, err := server.Query(tree.ToWarpScript(token, "", ""), txn)
	if err != nil {
		wErr := resp.Header.Get("X-Warp10-Error-Message")
		if wErr == "" {
			dump, err := httputil.DumpResponse(resp, true)
			if err == nil {
				wErr = string(dump)
			} else {
				wErr = "Unparsable error"
			}
		}
		return nil, errors.New("WarpScript request failed: " + wErr + " - " + err.Error())
	}

	body, err := getRequestBody(resp)
	if err != nil {
		return nil, err
	}

	return body, nil
}

func getToken(r *http.Request) (string, error) {
	token := core.RetrieveToken(r)
	if len(token) == 0 {
		return "", errors.New("no token provided")
	}

	return token, nil
}

func getRequestBody(resp *http.Response) ([]byte, error) {
	defer resp.Body.Close()

	return ioutil.ReadAll(resp.Body)
}

func logWarn(r *http.Request, code int, err error) {
	log.WithFields(log.Fields{
		"proto":  "graphite",
		"source": r.RemoteAddr,
		"method": r.Method,
		"status": code,
		"path":   r.URL.String(),
	}).Warn(err.Error())
}

func logErr(r *http.Request, code int, err error) {
	log.WithFields(log.Fields{
		"proto":  "graphite",
		"source": r.RemoteAddr,
		"method": r.Method,
		"status": code,
		"path":   r.URL.String(),
	}).Error(err.Error())
}

func respond(w http.ResponseWriter, code int, body []byte) {
	w.WriteHeader(code)
	w.Write(body) // nolint: gas
}

func respondWithJsonp(w http.ResponseWriter, code int, body []byte, callback string) {
	jsonp := bytes.NewBuffer([]byte{})

	jsonp.WriteString("<script>")   // nolint: gas
	jsonp.WriteString(callback)     // nolint: gas
	jsonp.WriteString("(")          // nolint: gas
	jsonp.Write(body)               // nolint: gas
	jsonp.WriteString(")</script>") // nolint: gas

	w.WriteHeader(code)
	w.Write(jsonp.Bytes()) // nolint: gas
}

func respondWithError(w http.ResponseWriter, code int, err error) {
	w.WriteHeader(code)
	w.Write([]byte(err.Error())) // nolint: gas
}
