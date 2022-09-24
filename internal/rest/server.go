package rest

import (
	"context"
	"fmt"

	"github.com/arxdsilva/desafios-api/internal/jwt"
	"github.com/arxdsilva/desafios-api/internal/service"
	"github.com/labstack/echo/v4"
)

type Server struct {
	service       service.Orders
	tokenProvider jwt.TokenProvider
	config        Config
	echo          *echo.Echo
}

func NewServer(svc service.Orders, tp jwt.TokenProvider, cfg Config) Server {
	return Server{
		service:       svc,
		tokenProvider: tp,
		config:        cfg,
	}
}

func (s Server) Run(ctx context.Context, cfg Config) error {
	go func() {
		<-ctx.Done()
		s.close()
	}()
	s.echo = echo.New()
	Register(s.echo)
	RouterRegister(s.echo)

	return s.echo.Start(portStr(cfg.Port))
}

func (s Server) close() {
	s.echo.Close()
}

func portStr(p int) string {
	return fmt.Sprintf(":%v", p)
}
