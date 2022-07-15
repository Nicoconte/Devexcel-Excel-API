package types

import "github.com/caarlos0/env"

type Config struct {
	Port    string `env:"PORT"`
	Host    string `env:"HOST"`
	Storage string `env:"STORAGE"`
}

func NewConfig() *Config {
	cfg := &Config{}

	err := env.Parse(cfg)

	if err != nil {
		panic(1)
	}

	return cfg
}

func NewConfigLocal() *Config {
	return &Config{
		Port:    "3000",
		Host:    "127.0.0.1",
		Storage: "storage",
	}
}
