package config

import (
	"log"

	"github.com/kelseyhightower/envconfig"
)

type Config struct {
	Env  string `envconfig:"ENV" default:"development"`
	Port int    `envconfig:"HTTP_PORT" default:"8000"`
	DB   PostgresDBConfig
}

type PostgresDBConfig struct {
	DSN              string `envconfig:"DB_DSN"`
	MigrationsSource string `envconfig:"MIGRATIONS_SOURCE"`
}

func NewConfig() Config {
	cfg := Config{}

	err := envconfig.Process("", &cfg)
	if err != nil {
		log.Fatal(err.Error())
	}

	return cfg
}
