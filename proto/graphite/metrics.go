package graphite

import (
	"github.com/prometheus/client_golang/prometheus"
)

var (
	requestFunctionCounter = prometheus.NewCounterVec(prometheus.CounterOpts{
		Namespace: "erlenmeyer",
		Subsystem: "graphite",
		Name:      "function",
		Help:      "Function used by user of graphite",
	}, []string{"function"})
)

func init() {
	prometheus.MustRegister(requestFunctionCounter)
}
