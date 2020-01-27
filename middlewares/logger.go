package middlewares

import (
	"crypto/sha256"
	"fmt"
	"strconv"
	"time"

	"github.com/labstack/echo"
	"github.com/prometheus/client_golang/prometheus"
	log "github.com/sirupsen/logrus"
)

const (
	// TxnHeader allow to trace logs of a request
	TxnHeader = "X-App-Txn"
)

var (
	requestCounter = prometheus.NewCounter(prometheus.CounterOpts{
		Namespace: "erlenmeyer",
		Subsystem: "http",
		Name:      "request",
		Help:      "Number of http request handled.",
	})

	requestErrorCounter = prometheus.NewCounter(prometheus.CounterOpts{
		Namespace: "erlenmeyer",
		Subsystem: "http",
		Name:      "error_request",
		Help:      "Number of http request in error.",
	})

	requestHTTPStatusCode = prometheus.NewCounterVec(prometheus.CounterOpts{
		Namespace: "erlenmeyer",
		Subsystem: "http",
		Name:      "status_code",
		Help:      "Get status code from request",
	}, []string{"status"})

	requestResponseTimeCounter = prometheus.NewCounterVec(prometheus.CounterOpts{
		Namespace: "erlenmeyer",
		Subsystem: "http",
		Name:      "response_time",
		Help:      "Response time of the request in nanoseconds",
	}, []string{"path"})
)

func init() {
	prometheus.MustRegister(requestCounter)
	prometheus.MustRegister(requestErrorCounter)
	prometheus.MustRegister(requestHTTPStatusCode)
	prometheus.MustRegister(requestResponseTimeCounter)
}

func logger(c echo.Context, next echo.HandlerFunc) error {
	req := c.Request()
	res := c.Response()
	start := time.Now()
	hash := sha256.New()
	txn := ""
	_, err := hash.Write([]byte(fmt.Sprintf("%s%x", req.Header.Get("X-Forwarded-For"), start.UnixNano())))
	if err == nil {
		txn = fmt.Sprintf("%x", hash.Sum(nil))

		c.Set("txn", txn)
		c.Response().Header().Add(TxnHeader, txn)
	} else {
		c.Set("txn", "")
		log.WithError(err).Warn("Failed generating TXN for this request")
	}

	if err := next(c); err != nil {
		c.Error(err)
	}

	stop := time.Now()
	duration := stop.Sub(start)
	path := req.URL.Path
	if path == "" {
		path = "/"
	}

	log.WithFields(log.Fields{
		"remote":     c.RealIP(),
		"host":       req.Host,
		"uri":        req.RequestURI,
		"method":     req.Method,
		"path":       path,
		"referer":    req.Referer(),
		"user_agent": req.UserAgent(),
		"status":     res.Status,
		"latency":    duration.String(),
		"date":       stop.UTC().String(),
		"txn":        txn,
	}).Info("Access")

	requestCounter.Inc()
	requestResponseTimeCounter.With(prometheus.Labels{
		"path": path,
	}).Add(float64(duration.Nanoseconds()))
	requestHTTPStatusCode.With(prometheus.Labels{
		"status": strconv.Itoa(res.Status),
	}).Inc()

	if res.Status >= 300 {
		requestErrorCounter.Inc()
	}

	return nil
}

// Logger middleware
func Logger() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(ctx echo.Context) error {
			return logger(ctx, next)
		}
	}
}
