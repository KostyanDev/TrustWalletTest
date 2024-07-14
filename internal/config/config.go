package config

import (
	"time"

	"github.com/caarlos0/env/v8"
)

func New() (Config, error) {
	var cfg Config
	err := env.Parse(&cfg)
	return cfg, err
}

type Config struct {
	App        App
	HTTPServer HTTPServer `envPrefix:"HTTP_"`
}

type App struct {
	Name string `env:"APP_NAME" envDefault:"app"`
}

type HTTPServer struct {
	Host         string        `env:"SERVER_HOST" envDefault:"localhost"`
	Port         int           `env:"SERVER_PORT" envDefault:"8181"`
	WriteTimeout time.Duration `env:"WRITE_TIMEOUT" envDefault:"15s"`
	ReadTimeout  time.Duration `env:"READ_TIMEOUT" envDefault:"15s"`
}
