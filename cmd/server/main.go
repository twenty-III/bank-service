package main

import (
	"bank/internal/account"
	"bank/internal/client"
	"bank/internal/config"
	"bank/internal/database"
	"bank/internal/logger"
	"context"
	"log/slog"
	"net/http"

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
	logger.Init(cfg.ENV)

	// initialize database
	db, err := database.Open(context.Background(), cfg.DSN)
	if err != nil {
		panic(err)
	}

	// main router of our app
	mux := http.NewServeMux()

	// account feature setup
	accRepo := account.NewRepo(db)
	accService := account.NewService(accRepo)
	accHandler := account.NewHandler(accService)
	accHandler.Mount(mux)

	// client feature setup
	clientRepo := client.NewRepo(db)
	clientService := client.NewService(clientRepo, accService)
	clientHandler := client.NewHandler(clientService)
	clientHandler.Mount(mux)

	// main server
	server := &http.Server{
		Addr:     ":" + cfg.PORT,
		Handler:  mux,
		ErrorLog: slog.NewLogLogger(slog.Default().Handler(), slog.LevelError),
	}

	err = server.ListenAndServe()
	if err != nil {
		panic(err)
	}
}
