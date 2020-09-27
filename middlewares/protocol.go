package middlewares

import (
	"strconv"

	"github.com/labstack/echo/v4"
	"github.com/prometheus/client_golang/prometheus"
)

var (
	requestProtocolCounter = prometheus.NewCounterVec(prometheus.CounterOpts{
		Namespace: "erlenmeyer",
		Subsystem: "protocol",
		Name:      "request",
		Help:      "Request by protocol",
	}, []string{"protocol"})

	requestProtocolErrorCounter = prometheus.NewCounterVec(prometheus.CounterOpts{
		Namespace: "erlenmeyer",
		Subsystem: "protocol",
		Name:      "error_request",
		Help:      "Request error by protocol",
	}, []string{"protocol", "status"})
)

func init() {
	prometheus.MustRegister(requestProtocolCounter)
	prometheus.MustRegister(requestProtocolErrorCounter)
}

// Protocol measure use of a protocol
func protocol(name string, ctx echo.Context, next echo.HandlerFunc) error {
	requestProtocolCounter.With(prometheus.Labels{
		"protocol": name,
	}).Inc()

	err := next(ctx)
	if ctx.Response().Status >= 300 {
		requestProtocolErrorCounter.With(prometheus.Labels{
			"protocol": name,
			"status":   strconv.Itoa(ctx.Response().Status),
		}).Inc()
	}

	return err
}

// Protocol middleware
func Protocol(name string) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(ctx echo.Context) error {
			return protocol(name, ctx, next)
		}
	}
}
