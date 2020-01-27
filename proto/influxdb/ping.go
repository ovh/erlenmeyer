package influxdb

import (
	"net/http"
)

// Ping returns a simple response to let the client know the server is running.
func (i *InfluxDB) Ping(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusNoContent)
}
