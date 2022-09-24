package rest

import (
	"log"
	"net/http"
	"strings"

	"github.com/hack-btg/backend/internal/jwt"
	"github.com/hack-btg/backend/internal/rest/handlers"
	"github.com/labstack/echo/v4"
)

func RouterRegister(e *echo.Echo, tp jwt.TokenProvider) {
	adapter := handlers.NewHTTPPrimaryAdapter(tp)

	e.GET("/", HealthCheck)
	e.POST("/login", adapter.UserLogin)
	// todo
	users := e.Group("/users", JWT())
	users.POST("/caixinha", adapter.CreateCX)        //create
	users.GET("/caixinhas", adapter.ListCX)          //list
	users.PUT("/caixinhas/", adapter.UpdateCX)       //update
	users.GET("/caixinhas/:id", adapter.GetOneCX)    //getone
	users.DELETE("/caixinhas/:id", adapter.DeleteCX) //delete

	banks := e.Group("/banks", JWT())
	banks.GET("/", adapter.ListCX)
}

// todo: allow this to be configurable and to pass optional checks
// ex: db, services, connections ...
func HealthCheck(c echo.Context) (err error) {
	ok := struct {
		Service string `json:"service"`
	}{"ok"}
	return c.JSON(http.StatusOK, ok)
}

// JWT could parse the token, avoid unwanted access
//
// only log now as this func is not necessary
func JWT() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			log.Printf("[JWT] Middleware token: %v\n", extractToken(c.Request()))
			return next(c)
		}
	}
}

func extractToken(r *http.Request) string {
	bearToken := r.Header.Get("x-authentication-token")
	strArr := strings.Split(bearToken, " ")
	if len(strArr) == 2 {
		return strArr[1]
	}
	return ""
}
