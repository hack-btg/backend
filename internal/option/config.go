package option

import (
	"fmt"

	"github.com/arxdsilva/desafios-api/internal/grpc"
	"github.com/arxdsilva/desafios-api/internal/jwt"
	"github.com/arxdsilva/desafios-api/internal/logger"
	"github.com/arxdsilva/desafios-api/internal/rest"
	"github.com/go-playground/validator/v10"
	"github.com/kelseyhightower/envconfig"
)

const prefix = "APP"

type Config struct {
	Rest rest.Config
	GRPC grpc.Config
	Log  logger.Config
	JWT  jwt.Config
}

func FromEnv() (*Config, error) {
	var config Config
	err := envconfig.Process(prefix, &config)
	if err != nil {
		return nil, fmt.Errorf("failed to load config: %w", err)
	}
	return &config, nil
}

func (c *Config) Validate() error {
	validate := validator.New()
	return validate.Struct(c)
}
