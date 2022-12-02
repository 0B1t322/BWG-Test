package config

var GlobalConfig Config = Config{
	DB: DBConfig{
		PostgreSQLDSN: "postgres://user:password@localhost:5432/db?sslmode=disable",
	},
	App: AppConfig{
		Port: "8080",
	},
}

type Config struct {
	// DBConfig contains database configuration
	DB DBConfig

	App AppConfig
}

type (
	DBConfig struct {
		PostgreSQLDSN string `env:"BWG_APP_POSTGRESQL_DSN"`
	}

	AppConfig struct {
		// Port is the port on which the server will listen
		Port string `env:"BWG_APP_PORT"`
	}
)
