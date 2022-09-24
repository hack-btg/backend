package rest

import (
	"context"

	"github.com/arxdsilva/desafios-api/internal/jwt"
	"github.com/arxdsilva/desafios-api/internal/service"
)

type Server struct {
	service       service.Orders
	tokenProvider jwt.TokenProvider
	config        Config
}

func NewServer(svc service.Orders, tp jwt.TokenProvider, cfg Config) Server {
	return Server{
		service:       svc,
		tokenProvider: tp,
		config:        cfg,
	}
}

func (s Server) Run(ctx context.Context) {
	// register routes
	//start http server
	go func() {
		<-ctx.Done()
		s.stop()
	}()
}

func (s Server) stop() {

}
