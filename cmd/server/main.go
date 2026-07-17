package main

import (
	"fmt"
	"log/slog"
	"net/http"
	"os"

	"calculator-backend/internal/api"
	"calculator-backend/internal/config"
	"calculator-backend/internal/evaluator"
	"calculator-backend/internal/service"
)

var (
	listenAndServe = http.ListenAndServe
	osExit         = os.Exit
)

func main() {
	osExit(run())
}

func run() int {
	cfg := config.Load()
	logger := buildLogger(cfg)
	slog.SetDefault(logger)

	if cfg.GinMode == "release" {
		os.Setenv("GIN_MODE", "release")
	}

	router := buildRouter()
	address := fmt.Sprintf(":%s", cfg.Port)
	slog.Info("starting server", "address", address)
	if err := listenAndServe(address, router); err != nil {
		slog.Error("server failed", "error", err)
		return 1
	}
	return 0
}

func buildLogger(cfg config.Config) *slog.Logger {
	level := new(slog.LevelVar)
	switch cfg.LogLevel {
	case "debug":
		level.Set(slog.LevelDebug)
	case "warn":
		level.Set(slog.LevelWarn)
	case "error":
		level.Set(slog.LevelError)
	default:
		level.Set(slog.LevelInfo)
	}
	return slog.New(slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{Level: level}))
}

func buildRouter() http.Handler {
	eval := evaluator.NewEvaluator()
	calculator := service.NewCalculatorService(eval)
	return api.SetupRoutes(calculator)
}
