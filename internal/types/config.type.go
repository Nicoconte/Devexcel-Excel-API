package types

import (
	"fmt"

	"github.com/caarlos0/env"
)

type Config struct {
	Env     string `env:"ENV"`
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

	fmt.Println("Env ", cfg.Env)

	return cfg
}

func NewConfigLocal() *Config {
	return &Config{
		Env:     "LOCAL",
		Port:    "3000",
		Host:    "127.0.0.1",
		Storage: "storage",
	}
}
