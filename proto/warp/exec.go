package warp

import (
	"io/ioutil"
	"net/http"
	"strings"

	"github.com/labstack/echo"
	"github.com/ovh/erlenmeyer/core"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

// Exec call the /api/v0/exec in order to execute warpscript
func Exec(ctx echo.Context) error {

	body, err := ioutil.ReadAll(ctx.Request().Body)
	if err != nil {
		log.WithError(err).Error("Cannot read the request body")
		return ctx.NoContent(http.StatusInternalServerError)
	}

	defer func() {
		if err := ctx.Request().Body.Close(); err != nil {
			log.WithError(err).Error("Cannot close the request body")
		}
	}()

	txn, ok := ctx.Get("txn").(string)
	if !ok {
		txn = ""
	}

	server := core.NewWarpServer(viper.GetString("warp_endpoint"), "warp10-query")
	res, err := server.Query(string(body), txn)
	if err != nil {
		log.WithError(err).Error("Cannot execute the request on the warp endpoint")
		return ctx.JSON(http.StatusInternalServerError, echo.Map{
			"error": err.Error(),
		})
	}

	for k, v := range res.Header {
		if strings.HasPrefix(k, "X-Warp") {
			ctx.Response().Header().Set(k, v[0])
		}
	}

	return ctx.Stream(res.StatusCode, res.Header.Get("Content-Type"), res.Body)
}
