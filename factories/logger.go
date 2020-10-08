package factories

import (
	"github.com/labstack/echo/v4"
	log "github.com/sirupsen/logrus"
)

func Logger(ctx echo.Context) *log.Entry {
	entry := log.WithFields(log.Fields{})
	if txn, ok := ctx.Get("txn").(string); ok {
		entry = entry.WithField("txn", txn)
	}

	return entry
}
