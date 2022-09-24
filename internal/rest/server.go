package rest

import (
	"context"
	"fmt"

	"github.com/hack-btg/backend/internal/jwt"
	"github.com/hack-btg/backend/internal/service"
	"github.com/labstack/echo/v4"
)

type Server struct {
	service service.Orders
	config  Config
	echo    *echo.Echo
}

func NewServer(svc service.Orders, cfg Config) Server {
	return Server{
		service: svc,
		config:  cfg,
	}
}

func (s Server) Run(ctx context.Context, cfg Config, tp jwt.TokenProvider) error {
	go func() {
		<-ctx.Done()
		s.close()
	}()
	s.echo = echo.New()
	Register(s.echo)
	RouterRegister(s.echo, tp)

	return s.echo.Start(portStr(cfg.Port))
}

func (s Server) close() {
	s.echo.Close()
}

func portStr(p int) string {
	return fmt.Sprintf(":%v", p)
}
