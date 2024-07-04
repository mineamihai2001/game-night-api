package helpers

import (
	env "github.com/Netflix/go-env"
	"log"
)

type Environment struct {
	App struct {
		Port int `env:"APP_PORT"`
	}
	Db struct {
		ConnectionString string `env:"DB_CONN_STRING"`
		Database         string `env:"DB_DATABASE"`
	}
	Log struct {
		ConsoleEnabled    bool   `env:"LOG_CONSOLE_ENABLED"`
		ConsoleLevel      string `env:"LOG_CONSOLE_LEVEL"`
		BackgroundEnabled bool   `env:"LOG_BACKGROUND_ENABLED"`
		Writer            struct {
			Protocol string `env:"LOG_WRITER_PROTOCOL"`
			Host     string `env:"LOG_WRITER_HOST"`
			User     string `env:"LOG_WRITER_USER"`
			Password string `env:"LOG_WRITER_PASSWORD"`
			Port     int    `env:"LOG_WRITER_PORT"`
			VHost    string `env:"LOG_WRITER_VHOST"`
			Queue    string `env:"LOG_WRITER_QUEUE"`
		}
	}
	Service struct {
		Name    string `env:"SERVICE_NAME"`
		Version string `env:"SERVICE_VERSION"`
	}
}

var instance *Environment = nil

func new() *Environment {
	var environment Environment
	_, err := env.UnmarshalFromEnviron(&environment)

	if err != nil {
		log.Fatal("[ENVIRONMENT ERROR] - ", err)
	}

	return &environment
}

func Env() *Environment {
	if instance == nil {
		instance = new()
	}

	return instance
}
