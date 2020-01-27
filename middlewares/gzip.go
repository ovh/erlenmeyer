package middlewares

import (
	"compress/gzip"
	"net/http"
	"strings"

	"github.com/labstack/echo"
	log "github.com/sirupsen/logrus"
)

const (
	contentEncodingHeader = "Content-Encoding"
	contentTypeHeader     = "Content-Type"
	gzipContentType       = "application/gzip"
)

// Gzip gunzip request before handlers
func Gzip() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(ctx echo.Context) error {
			return gunzip(ctx, next)
		}
	}
}

func gunzip(c echo.Context, next echo.HandlerFunc) error {

	contentEncoding := c.Request().Header.Get(contentEncodingHeader)
	contentType := c.Request().Header.Get(contentTypeHeader)

	if strings.Contains(strings.ToLower(contentEncoding), "gzip") || strings.ToLower(contentType) == gzipContentType {

		r, err := gzip.NewReader(c.Request().Body)
		if err != nil {
			return c.NoContent(http.StatusBadRequest)
		}
		defer func() {
			if err := r.Close(); err != nil {
				log.
					WithError(err).
					Warn("Cannot close GZIP body reader")
			}
		}()

		c.Request().Body = r
	}

	return next(c)
}
