package types

import "github.com/caarlos0/env"

type Config struct {
	Env  string `env:"APP_ENV"`
	Port string `env:"PORT"`
	Host string `env:"HOST"`
}

func NewConfig() *Config {
	cfg := &Config{}

	err := env.Parse(cfg)

	if err != nil {
		panic(1)
	}

	return cfg
}
