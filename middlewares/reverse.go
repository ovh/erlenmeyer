package middlewares

import (
	"net/http"
	"strings"

	"github.com/labstack/echo/v4"
	log "github.com/sirupsen/logrus"
)

// ReverseConfig is the configuration to describe
// a reverse proxy
type ReverseConfig struct {
	Path string
	URL  string
}

func reverse(config ReverseConfig, ctx echo.Context) error {
	req := ctx.Request()

	uri := config.URL + "/" + ctx.Param("*")
	if config.Path != "" {
		uri = config.URL + config.Path
	}
	if strings.Contains(req.RequestURI, "?") {
		interogation := strings.Index(req.RequestURI, "?")
		uri += req.RequestURI[interogation:]
	}

	log.WithFields(log.Fields{
		"reverse": uri,
		"remote":  ctx.RealIP(),
		"host":    req.Host,
		"uri":     req.RequestURI,
		"method":  req.Method,
		"path":    req.URL.Path,
		"referer": req.Referer(),
	}).Debug("Execute reverse proxy")

	req, err := http.NewRequest(ctx.Request().Method, uri, ctx.Request().Body)
	if err != nil {
		return ctx.String(http.StatusInternalServerError, err.Error())
	}

	req.Header = ctx.Request().Header
	res, err := http.DefaultClient.Do(req)

	if err != nil {
		log.WithError(err).Error("Cannot execute the request on the endpoint")
		return ctx.NoContent(http.StatusBadGateway)
	}

	if res.Header != nil {
		// Copy Warp10 headers
		for k, v := range res.Header {
			log.Debugf("Header %v: %v", k, v)
			if strings.HasPrefix(k, "X-Warp") {
				ctx.Response().Header().Set(k, v[0])
			}
		}
		// copy HTTP related headers
		copyHeaders(res.Header, ctx.Response(), []string{
			"Content-Type",
			"Content-Length",
			"Content-Encoding",
			"Vary",
		})
	}

	ct := res.Header.Get("Content-Type")
	if ct == "" {
		ct = "text/plain"
	}

	return ctx.Stream(res.StatusCode, ct, res.Body)
}

// ReverseWithConfig execute a reverse proxy using
// the configuration given in parameters
func ReverseWithConfig(config ReverseConfig) echo.HandlerFunc {
	return func(ctx echo.Context) error {
		return reverse(config, ctx)
	}
}

func copyHeaders(reqheaders http.Header, res *echo.Response, headers []string) {
	for _, header := range headers {
		if reqheaders.Get(header) != "" {
			res.Header().Set(header, reqheaders.Get(header))
		}
	}
}
