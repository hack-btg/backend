package rest

import (
	"net/http"

	"github.com/hack-btg/backend/internal/jwt"
	"github.com/hack-btg/backend/internal/rest/handlers"
	"github.com/labstack/echo/v4"
)

func RouterRegister(e *echo.Echo, tp jwt.TokenProvider) {
	adapter := handlers.NewHTTPPrimaryAdapter(nil, tp)

	e.GET("/", HealthCheck)
	e.POST("/login", adapter.UserLogin)
}

// todo: allow this to be configurable and to pass optional checks
// ex: db, services, connections ...
func HealthCheck(c echo.Context) (err error) {
	ok := struct {
		Service string `json:"service"`
	}{"ok"}
	return c.JSON(http.StatusOK, ok)
}
