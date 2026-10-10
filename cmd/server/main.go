package main

import (
	"bank/internal/logger"
	"log/slog"
)

func main() {
	logger.Init("production")

	slog.Info("hello from slog")
}
