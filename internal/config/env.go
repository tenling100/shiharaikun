package config

import (
	"fmt"

	"github.com/kelseyhightower/envconfig"
)

type Env struct {
	// Env is the environment of the application.
	GRPCPort string `env:"GRPC_PORT" envDefault:"50051"`
	DBHost   string `env:"DB_HOST" envDefault:"localhost"`
	DBPort   string `env:"DB_PORT" envDefault:"3306"`
	DBUser   string `env:"DB_USER" envDefault:"user"`
	DBPass   string `env:"DB_PASS"`
	DBName   string `env:"DB_NAME" envDefault:"shiharaikun"`
}

// SetupEnv loads environment variables into the Env struct
func SetupEnv() (*Env, error) {
	var env Env

	// Load environment variables into the Env struct using envconfig
	if err := envconfig.Process("", &env); err != nil {
		return nil, fmt.Errorf("failed to load environment variables: %w", err)
	}

	return &env, nil
}
