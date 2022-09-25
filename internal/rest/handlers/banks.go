package handlers

import (
	"net/http"

	"github.com/labstack/echo/v4"
	log "github.com/sirupsen/logrus"
)

func (h *HTTPPrimaryAdapter) ListBanks(c echo.Context) (err error) {
	l := log.WithField("cid", c.Response().Header().Get(echo.HeaderXRequestID))
	banks, err := h.BankService.List(c.Request().Context())
	if err != nil {
		l.Error("[ListBanks] error ", err.Error())
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	l.Info("[ListBanks] listing")
	return c.JSON(http.StatusOK, banks)
}
