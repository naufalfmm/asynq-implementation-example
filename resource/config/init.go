package config

import (
	"os"

	"github.com/joho/godotenv"
	"github.com/kelseyhightower/envconfig"
	"github.com/pkg/errors"
)

type EnvConfig struct {
	Port int `envconfig:"PORT" default:"8090" required:"true"`

	RedisHost     string `envconfig:"REDIS_HOST" required:"true"`
	RedisPort     string `envconfig:"REDIS_PORT" required:"true"`
	RedisPassword string `envconfig:"REDIS_PASSWORD" required:"true"`
}

func NewConfig() (*EnvConfig, error) {
	var config EnvConfig

	filename := os.Getenv("CONFIG_FILE")

	if filename == "" {
		filename = ".env"
	}

	if _, err := os.Stat(filename); os.IsNotExist(err) {
		if err := envconfig.Process("", &config); err != nil {
			return nil, errors.Wrap(err, "failed to read from env variable")
		}
	}

	if err := godotenv.Load(filename); err != nil {
		return nil, errors.Wrap(err, "failed to read from .env file")
	}

	if err := envconfig.Process("", &config); err != nil {
		return nil, errors.Wrap(err, "failed to read from env variable")
	}

	return &config, nil
}
