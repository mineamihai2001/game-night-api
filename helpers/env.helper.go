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
}

var instance *Environment = nil

func Create() *Environment {
	var environment Environment
	_, err := env.UnmarshalFromEnviron(&environment)

	if err != nil {
		log.Fatal("[ENVIRONMENT ERROR] - ", err)
	}

	return &environment
}

func Env() *Environment {
	if instance == nil {
		instance = Create()
	}

	return instance
}