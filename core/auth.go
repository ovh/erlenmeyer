package core

import (
	"encoding/base64"
	"net/http"
	"strings"

	log "github.com/sirupsen/logrus"
)

func init() {
}

// RetrieveToken from multiple sources
func RetrieveToken(req *http.Request) string {
	// retrieve token from the basic auth
	token := retrieveTokenFromBasicAuth(req)
	if token != "" {
		return token
	}

	// support influx form authentication
	return retrieveTokenFromInfluxVariables(req)
}

// retrieveTokenFromBasicAuth is fetching the token for an HTTP Request
func retrieveTokenFromBasicAuth(request *http.Request) string {
	// Getting token from BasicAuth
	s := strings.SplitN(request.Header.Get("Authorization"), " ", 2)
	if len(s) != 2 {
		return ""
	}
	b, err := base64.StdEncoding.DecodeString(s[1])
	if err != nil {
		return ""
	}
	pair := strings.SplitN(string(b), ":", 2)
	if len(pair) != 2 {
		return ""
	}
	return pair[1]
}

// retrieve token from influx form
// https://docs.influxdata.com/influxdb/v1.7/administration/authentication_and_authorization/
func retrieveTokenFromInfluxVariables(req *http.Request) string {
	params := req.URL.Query()
	token := params.Get("p")
	if token != "" {
		return token
	}

	if err := req.ParseForm(); err != nil {
		log.WithError(err).WithFields(log.Fields{
			"protocol": "influxdb",
		}).Error("Could not parse form request")
	}

	return req.FormValue("p")
}
