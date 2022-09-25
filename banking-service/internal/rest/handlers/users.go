package handlers

import (
	"net/http"

	"github.com/labstack/echo/v4"
	log "github.com/sirupsen/logrus"
)

// UserInfo ideally would get a specific user's info
//
// as we're in a hackathon we're using the one user provided in the dump
// filtered in python
func (h *HTTPPrimaryAdapter) UserInfo(c echo.Context) (err error) {
	l := log.WithField("cid", c.Response().Header().Get(echo.HeaderXRequestID))
	infos := h.UserService.Info(c.Request().Context())
	l.Info("[UserInfo] ok")
	return c.JSON(http.StatusOK, infos)
}
