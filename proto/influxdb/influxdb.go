package influxdb

import "github.com/prometheus/client_golang/prometheus"

// InfluxDB endpoint
type InfluxDB struct {
	ReqCounter  prometheus.Counter
	ErrCounter  prometheus.Counter
	WarnCounter prometheus.Counter
}

// GetReqCounter satisfies the protocol interface
func (p *InfluxDB) GetReqCounter() prometheus.Counter {
	return p.ReqCounter
}

// NewInfluxDB is creating a new influxDB query handler
func NewInfluxDB() *InfluxDB {
	c := &InfluxDB{}

	// metrics
	c.ReqCounter = prometheus.NewCounter(prometheus.CounterOpts{
		Namespace: "erlenmeyer",
		Subsystem: "influxdb",
		Name:      "request",
		Help:      "Number of request handled.",
	})
	prometheus.MustRegister(c.ReqCounter)
	c.ErrCounter = prometheus.NewCounter(prometheus.CounterOpts{
		Namespace: "erlenmeyer",
		Subsystem: "influxdb",
		Name:      "errors",
		Help:      "Number of request in errors.",
	})
	prometheus.MustRegister(c.ErrCounter)
	c.WarnCounter = prometheus.NewCounter(prometheus.CounterOpts{
		Namespace: "erlenmeyer",
		Subsystem: "influxdb",
		Name:      "warning",
		Help:      "Number of errored client requests.",
	})
	prometheus.MustRegister(c.WarnCounter)

	return c
}
