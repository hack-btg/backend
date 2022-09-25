package handlers

import (
	"net/http"

	"github.com/labstack/echo/v4"
	log "github.com/sirupsen/logrus"
)

func (h *HTTPPrimaryAdapter) UsersSocialGet(c echo.Context) (err error) {
	l := log.WithField("cid", c.Response().Header().Get(echo.HeaderXRequestID))
	l.Info("[Get] ok")
	return c.JSON(http.StatusOK, h.socialService.Get(c.Request().Context()))
}
