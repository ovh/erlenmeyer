package middlewares

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/ovh/erlenmeyer/core"
	"github.com/ovh/erlenmeyer/factories"
)

// Deny is a middleware that check tokens and deny if there are in configuration
func Deny(tokens []string) echo.MiddlewareFunc {
	mapper := make(map[string]interface{})
	for _, token := range tokens {
		mapper[token] = nil
	}

	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(ctx echo.Context) error {
			logger := factories.Logger(ctx)

			token := core.RetrieveToken(ctx.Request())
			if token == "" {
				logger.Warn("could not retrieve token")
				return ctx.NoContent(http.StatusUnauthorized)
			}

			if _, ok := mapper[token]; ok {
				logger.Warn("Token is forbidden")
				return ctx.NoContent(http.StatusForbidden)
			}

			return next(ctx)
		}
	}
}
