package main

import (
	"bank/internal/config"
	"bank/internal/logger"

	"github.com/joho/godotenv"
)

func main() {
	// load environment variables from .env file
	err := godotenv.Load()
	if err != nil {
		panic(err)
	}

	// load all the configs into Config struct
	cfg, err := config.Load()
	if err != nil {
		panic(err)
	}

	// initialize global logger
	logger.Init(cfg.Env)
}
