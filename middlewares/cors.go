package middlewares

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func cors(ctx echo.Context, next echo.HandlerFunc) error {
	origin := ctx.Request().Header.Get("Origin")
	if len(origin) <= 0 {
		origin = "*"
	}

	ctx.Response().Header().Set("Access-Control-Allow-Origin", origin)
	ctx.Response().Header().Set("Access-Control-Allow-Credentials", "true")
	ctx.Response().Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
	ctx.Response().Header().Set("Access-Control-Allow-Headers", "Accept, content-type, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, Origin, X-Warp10-Elapsed, X-Warp10-Error-Line, X-Warp10-Error-Message, X-Warp10-Fetched, X-Warp10-Ops")
	ctx.Response().Header().Set("Access-Control-Expose-Headers", "X-Warp10-Elapsed, X-Warp10-Error-Line, X-Warp10-Error-Message, X-Warp10-Fetched, X-Warp10-Ops")

	if ctx.Request().Method == http.MethodOptions {
		return ctx.NoContent(http.StatusOK)
	}

	return next(ctx)
}

// CORS handle preflight request
// the default middleware is not used because we have to set dynamically
// the origin
func CORS() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(ctx echo.Context) error {
			return cors(ctx, next)
		}
	}
}
