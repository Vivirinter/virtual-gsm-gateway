package main

import (
	"github.com/Vivirinter/virtual-gsm-gateway/internal/gateway"
	"github.com/Vivirinter/virtual-gsm-gateway/internal/routes"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"log/slog"
	"net/http"
	"os"
)

func main() {
	logger := slog.New(slog.NewTextHandler(os.Stdout, nil))
	gw := gateway.NewGateway(logger)

	r := chi.NewRouter()
	r.Use(middleware.RequestID)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	routes.RegisterRoutes(r, gw)

	// Start the server
	logger.Info("Starting server on :8080")
	if err := http.ListenAndServe(":8080", r); err != nil {
		logger.Error("could not start server", slog.Any("error", err))
	}
}
