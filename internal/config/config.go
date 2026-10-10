package config

import "os"

type Config struct {
	Env string
}

func Load() (*Config, error) {
	env := os.Getenv("env")

	return &Config{
		Env: env,
	}, nil
}
