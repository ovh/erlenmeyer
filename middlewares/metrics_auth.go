package middlewares

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/ovh/erlenmeyer/factories"
	"github.com/spf13/viper"
)

// BasicAuth is a middleware that check Metrics endpoint basic auth and deny access when auth is invalid
func BasicAuth() echo.MiddlewareFunc {
	username := viper.GetString("metrics.basicauth.user")
	password := viper.GetString("metrics.basicauth.password")
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(ctx echo.Context) error {
			logger := factories.Logger(ctx)
			basic_user, basic_pwd, ok := ctx.Request().BasicAuth()
			if !ok {
				logger.Warn("Error parsing /metrics basic auth")
				return ctx.NoContent(http.StatusUnauthorized)
			}

			if basic_user != username || basic_pwd != password {
				logger.Warnf("Invalid authentication on /metrics for user: %s\n", basic_user)
				return ctx.NoContent(http.StatusUnauthorized)
			}

			return next(ctx)
		}
	}
}
