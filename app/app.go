package app

import (
	"context"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

const (
	logLevelDebug = "DEBUG"
	logLevelInfo  = "INFO"
	logLevelWarn  = "WARN"
	logLevelError = "ERROR"
)

type Config struct {
	LogLevel              string `env:"LOG_LEVEL" envDefault:"INFO"`
	ServerAddress         string `env:"SERVER_ADDRESS" envDefault:"127.0.0.1:8080"`
	ReadTimeoutInSeconds  int    `env:"READ_TIMEOUT_IN_SEC" envDefault:"5"`
	WriteTimeoutInSeconds int    `env:"WRITE_TIMEOUT_IN_SEC" envDefault:"5"`
}

type App struct {
	logger *slog.Logger
	server *http.Server
	done   chan os.Signal
}

func New(cfg Config) *App {
	var logLevel slog.Level
	switch cfg.LogLevel {
	case logLevelDebug:
		logLevel = slog.LevelDebug
	case logLevelInfo:
		logLevel = slog.LevelInfo
	case logLevelWarn:
		logLevel = slog.LevelWarn
	case logLevelError:
		logLevel = slog.LevelError
	default:
		logLevel = slog.LevelInfo
	}
	logger := slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{
		AddSource: true,
		Level:     logLevel,
	}))

	mux := http.NewServeMux()
	server := &http.Server{
		Addr:         cfg.ServerAddress,
		Handler:      mux,
		ReadTimeout:  time.Duration(cfg.ReadTimeoutInSeconds) * time.Second,
		WriteTimeout: time.Duration(cfg.WriteTimeoutInSeconds) * time.Second,
	}

	done := make(chan os.Signal, 1)
	signal.Notify(done, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)

	app := App{
		logger: logger,
		server: server,
		done:   done,
	}
	app.registerRoutes(mux)

	return &app
}

func (app *App) Start() {
	go func() {
		if err := app.server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			app.logger.Error("failed to listen", err)
			os.Exit(1)
		}
	}()
	app.logger.Info("server started", "address", app.server.Addr)
}

func (app *App) Stop() {
	<-app.done
	app.logger.Info("server stopped")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := app.server.Shutdown(ctx); err != nil {
		app.logger.Error("failed to shutdown the http server", err)
		os.Exit(2)
	}
	app.logger.Info("server shutdown successfully")
}
