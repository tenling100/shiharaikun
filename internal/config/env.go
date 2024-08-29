package config

import (
	"fmt"

	"github.com/kelseyhightower/envconfig"
)

type Env struct {
	// Env is the environment of the application.
	GRPCPort string `envconfig:"GRPC_PORT" default:"50051"`
	DBHost   string `envconfig:"DB_HOST" default:"localhost"`
	DBPort   string `envconfig:"DB_PORT" default:"3306"`
	DBUser   string `envconfig:"DB_USER" default:"testuser"`
	DBPass   string `envconfig:"DB_PASS" default:"testpassword"`
	DBName   string `envconfig:"DB_NAME" default:"testdb"`
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
