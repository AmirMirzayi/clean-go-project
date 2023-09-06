package config

import (
	"github.com/caarlos0/env/v9"
)

type HttpServer struct {
	Address string `env:"HTTP_SERVER_ADDRESS"`
	Port    int    `env:"HTTP_SERVER_PORT"`
}

type Config struct {
	HttpServer
}

func LoadConfig() Config {
	cfg := Config{}
	if err := env.Parse(&cfg); err != nil {
		panic(err)
	}
	return cfg
}
