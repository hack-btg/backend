package handlers

import (
	"net/http"
	"strconv"

	"github.com/hack-btg/backend/internal/domains/models"
	"github.com/labstack/echo/v4"
	log "github.com/sirupsen/logrus"
)

func (h *HTTPPrimaryAdapter) ListCX(c echo.Context) (err error) {
	l := log.WithField("cid", c.Response().Header().Get(echo.HeaderXRequestID))
	cxs := h.CaixinhaService.List(c.Request().Context())
	l.Info("[ListCX] listing")
	return c.JSON(http.StatusOK, cxs)
}

func (h *HTTPPrimaryAdapter) GetOneCX(c echo.Context) (err error) {
	l := log.WithField("cid", c.Response().Header().Get(echo.HeaderXRequestID))
	id := c.Param("id")
	idInt, err := strconv.Atoi(id)
	if err != nil {
		l.Error("[GetOneCX] error ", err.Error())
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	cx, err := h.CaixinhaService.GetOne(c.Request().Context(), idInt)
	if err != nil {
		l.Error("[GetOneCX] error ", err.Error())
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	l.Info("[GetOneCX] listing")
	return c.JSON(http.StatusOK, cx)
}

func (h *HTTPPrimaryAdapter) CreateCX(c echo.Context) (err error) {
	l := log.WithField("cid", c.Response().Header().Get(echo.HeaderXRequestID))
	cx := &models.Caixinha{}
	if err := c.Bind(cx); err != nil {
		l.Error("[CreateCX] Bind error ", err.Error())
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	ccx, err := h.CaixinhaService.Create(c.Request().Context(), *cx)
	if err != nil {
		l.Error("[CreateCX] svc create error ", err.Error())
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	l.Info("[CreateCX] listing")
	return c.JSON(http.StatusOK, ccx)
}
