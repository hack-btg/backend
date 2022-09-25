package handlers

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/hack-btg/backend/banking-service/internal/domains/models"
	"github.com/labstack/echo/v4"

	log "github.com/sirupsen/logrus"
)

var (
	ErrUser = errors.New("could not login")
)

func (h *HTTPPrimaryAdapter) UserLogin(c echo.Context) (err error) {
	l := log.WithField("cid", c.Response().Header().Get(echo.HeaderXRequestID))
	u := &models.User{}
	if err = c.Bind(u); err != nil {
		l.Error(fmt.Sprintf("[UserLogin] (Bind) err: %s", err.Error()))
		return echo.NewHTTPError(http.StatusNotAcceptable, ErrUser.Error())
	}
	t, err := h.tokenProvider.New(u.Email, "uuid")
	if err != nil {
		l.Error(fmt.Sprintf("[UserLogin] (jwt.New) err: %s", err.Error()))
		return echo.NewHTTPError(http.StatusInternalServerError, ErrUser.Error())
	}
	ut := &models.API{
		Token:    t,
		UserUUID: "uuid",
		UserType: "user",
		Name:     u.Email,
	}
	l.Info(fmt.Sprintf("[UserLogin] success: %s", u.Email))
	return c.JSON(http.StatusOK, ut)
}
