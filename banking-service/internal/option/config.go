package option

import (
	"fmt"

	"github.com/go-playground/validator/v10"
	"github.com/hack-btg/backend/banking-service/internal/jwt"
	"github.com/hack-btg/backend/banking-service/internal/logger"
	"github.com/hack-btg/backend/banking-service/internal/rest"
	"github.com/kelseyhightower/envconfig"
)

const prefix = "APP"

type Config struct {
	Rest rest.Config
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
