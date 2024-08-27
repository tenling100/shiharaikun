package config

type Env struct {
	// Env is the environment of the application.
	GRPCPort string `env:"GRPC_PORT" envDefault:"50051"`
	DBHost   string `env:"DB_HOST" envDefault:"localhost"`
	DBPort   string `env:"DB_PORT" envDefault:"3306"`
	DBUser   string `env:"DB_USER" envDefault:"user"`
	DBPass   string `env:"DB_PASS"`
	DBName   string `env:"DB_NAME" envDefault:"shiharaikun"`
}
