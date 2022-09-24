package grpc

import (
	"context"
	"errors"
	"fmt"
	"net"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

type (
	Config struct {
		Port int `envconfig:"APP_GRPC_PORT" default:"8080"`
	}

	Server struct {
		server *grpc.Server
		config ServerOptions
	}

	ServerOptions struct {
		address     string
		port        int
		grpcOptions []grpc.ServerOption
	}

	ServerOption func(*ServerOptions) error
)

const (
	DefaultServerAddress = "0.0.0.0"
	DefaultServerPort    = 8080
)

var (
	ErrEmptyRegisterFunc = errors.New("register function parameter can't be nil")
	ErrEmptyHost         = errors.New("host can't be empty")
	ErrEmptyPort         = errors.New("port can't be empty")
)

// NewServer creates instance of Server, it accepts registerService
// function in which you connect(register) your proto server implementation with handler
// also you can specify options to change default grpc server behaviour
func NewServer(resolver Resolver, options ...ServerOption) (*Server, error) {
	config := ServerOptions{
		address: DefaultServerAddress,
		port:    DefaultServerPort,
	}

	for _, option := range options {
		if err := option(&config); err != nil {
			return nil, err
		}
	}

	config.grpcOptions = append(
		config.grpcOptions,
		grpc.ChainUnaryInterceptor(
			CorrelationIDInterceptor,
		),
	)

	server := grpc.NewServer(config.grpcOptions...)

	// register resolver to server

	reflection.Register(server)

	return &Server{
		config: config,
		server: server,
	}, nil
}

func (s *Server) Run(ctx context.Context) error {
	address := fmt.Sprintf("%s:%d", s.config.address, s.config.port)
	listener, err := net.Listen("tcp", address)
	if err != nil {
		return fmt.Errorf("failed to create listener on address:port: %s:%d", s.config.address, s.config.port)
	}

	go func() {
		<-ctx.Done()
		s.stop()
	}()
	return s.server.Serve(listener)
}

func (s *Server) stop() {
	s.server.Stop()
}
