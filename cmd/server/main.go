package main

import (
	"context"
	"errors"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/AlexEgbuna/secure-business-platform/internal/api"
	"github.com/AlexEgbuna/secure-business-platform/internal/config"
)

const (
	readTimeout     = 10 * time.Second
	writeTimeout    = 15 * time.Second
	idleTimeout     = 60 * time.Second
	shutdownTimeout = 10 * time.Second
)

func main() {
	cfg, err := config.Load()
	if err != nil {
		log.Fatalf("failed to load configuration: %v", err)
	}

	if err := cfg.Validate(); err != nil {
		log.Fatalf("invalid configuration: %v", err)
	}

	router := api.NewRouter()

	server := &http.Server{
		Addr:         cfg.Application.Host + ":" + itoa(cfg.Application.Port),
		Handler:      router.Handler(),
		ReadTimeout:  readTimeout,
		WriteTimeout: writeTimeout,
		IdleTimeout:  idleTimeout,
	}

	serverErrors := make(chan error, 1)

	go func() {
		log.Printf(
			"server listening on %s",
			server.Addr,
		)

		serverErrors <- server.ListenAndServe()
	}()

	shutdownSignals := make(chan os.Signal, 1)
	signal.Notify(
		shutdownSignals,
		os.Interrupt,
		syscall.SIGTERM,
	)

	select {
	case err := <-serverErrors:
		if !errors.Is(err, http.ErrServerClosed) {
			log.Fatalf("HTTP server failed: %v", err)
		}

	case sig := <-shutdownSignals:
		log.Printf("received shutdown signal: %s", sig)

		ctx, cancel := context.WithTimeout(
			context.Background(),
			shutdownTimeout,
		)
		defer cancel()

		if err := server.Shutdown(ctx); err != nil {
			log.Printf("graceful shutdown failed: %v", err)
			_ = server.Close()
		}
	}
}

func itoa(value int) string {
	const digits = "0123456789"

	if value == 0 {
		return "0"
	}

	var result [20]byte
	position := len(result)

	for value > 0 {
		position--
		result[position] = digits[value%10]
		value /= 10
	}

	return string(result[position:])
}
