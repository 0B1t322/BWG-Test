package config

import (
	"github.com/joho/godotenv"
	"github.com/kelseyhightower/envconfig"
	log "github.com/sirupsen/logrus"
)

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
		PostgreSQLDSN string `envconfig:"BWG_APP_POSTGRESQL_DSN"`
	}

	AppConfig struct {
		// Port is the port on which the server will listen
		Port string `envconfig:"BWG_APP_PORT"`
	}
)

func InitGlobalConfig() {
	if err := godotenv.Load("./.env"); err != nil {
		log.Info("Don't find .env file")
	}

	if err := envconfig.Process("bwg_app", &GlobalConfig); err != nil {
		log.WithFields(
			log.Fields{
				"function": "envconfig.Process",
				"error":    err,
			},
		).Info("Can't read env vars")
	}

	log.WithFields(
		log.Fields{
			"from": "InitGlobalConfig",
		},
	).Infof("config: %+v", GlobalConfig)
}
