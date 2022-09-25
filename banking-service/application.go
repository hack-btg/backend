package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/hack-btg/backend/banking-service/internal/jwt"
	"github.com/hack-btg/backend/banking-service/internal/logger"
	"github.com/hack-btg/backend/banking-service/internal/option"
	"github.com/hack-btg/backend/banking-service/internal/rest"
	"golang.org/x/sync/errgroup"

	log "github.com/sirupsen/logrus"
)

func main() {
	if err := run(); err != nil {
		log.Error(err)
		os.Exit(1)
	}
}

func run() error {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	setupGracefulShutdown(cancel)

	cfg, err := option.FromEnv()
	if err != nil {
		return err
	}
	return startup(ctx, cfg)
}

func startup(ctx context.Context, cfg *option.Config) error {

	if err := logger.Setup(cfg.Log); err != nil {
		return fmt.Errorf(`log setup failed %w`, err)
	}

	if err := cfg.Validate(); err != nil {
		return fmt.Errorf(`config validation error %w`, err)
	}

	log.Info("starting the service")

	srv := rest.NewServer(nil, cfg.Rest)

	provider := jwt.NewProvider(cfg.JWT)

	errg, ctx := errgroup.WithContext(ctx)
	errg.Go(func() error {
		return srv.Run(ctx, cfg.Rest, provider)
	})

	log.Infof("service started on port %v", cfg.Rest.Port)
	if err := errg.Wait(); err != nil {
		log.Error(err)
	}

	log.Info("service stopped")
	return nil
}

func setupGracefulShutdown(stop func()) {
	signalChannel := make(chan os.Signal, 1)
	signal.Notify(signalChannel, os.Interrupt, syscall.SIGTERM)
	go func() {
		<-signalChannel
		log.Info("received interrupt signal. Gracefully shutting down the service.")
		stop()
	}()
}
