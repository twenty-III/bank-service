package config

import (
	"fmt"
	"log/slog"
	"os"
)

type Config struct {
	ENV  string
	DSN  string
	PORT string
}

func Load() (*Config, error) {
	env := os.Getenv("ENV")

	dsn := os.Getenv("DSN")
	if dsn == "" {
		return nil, fmt.Errorf("failed to read 'DSN' env var")
	}

	port := getEnvOrDefault("PORT", "8080")

	return &Config{
		ENV:  env,
		DSN:  dsn,
		PORT: port,
	}, nil
}

func getEnvOrDefault(env, def string) string {
	res := os.Getenv(env)
	if res == "" {
		slog.Warn("'PORT' not set, falling back to '8080'")
		return def
	}
	return res
}
