package middlewares

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

// Native wrap native handler to echo handler
// @see echo.WrapHandler
func Native(h func(http.ResponseWriter, *http.Request)) echo.HandlerFunc {
	return func(c echo.Context) error {
		h(c.Response(), c.Request())
		return nil
	}
}
