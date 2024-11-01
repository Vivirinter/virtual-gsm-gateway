package main

import (
	"github.com/Vivirinter/virtual-gsm-gateway/internal/gateway"
	"github.com/Vivirinter/virtual-gsm-gateway/internal/middleware"
	"github.com/go-chi/chi/v5"
	"log/slog"
	"net/http"
	"os"
)

func main() {
	logger := slog.New(slog.NewTextHandler(os.Stdout, nil))
	gw := gateway.NewGateway(logger)

	r := chi.NewRouter()

	r.Use(middleware.ErrorLogger(logger))

	// SMS routes
	r.Method(http.MethodPost, "/send-sms", http.HandlerFunc(gw.SendSMS))
	r.Method(http.MethodGet, "/messages", http.HandlerFunc(gw.GetSMS))
	r.Method(http.MethodDelete, "/delete-messages", http.HandlerFunc(gw.DeleteSMS))
	r.Method(http.MethodPut, "/update-sms", http.HandlerFunc(gw.UpdateSMS))

	// USSD routes
	r.Method(http.MethodPost, "/send-ussd", http.HandlerFunc(gw.SendUSSD))
	r.Method(http.MethodGet, "/ussd-requests", http.HandlerFunc(gw.GetUSSD))
	r.Method(http.MethodDelete, "/delete-ussd-requests", http.HandlerFunc(gw.DeleteUSSD))
	r.Method(http.MethodPut, "/update-ussd", http.HandlerFunc(gw.UpdateUSSD))

	// MMS routes
	r.Method(http.MethodPost, "/send-mms", http.HandlerFunc(gw.SendMMS))
	r.Method(http.MethodGet, "/mms-messages", http.HandlerFunc(gw.GetMMS))
	r.Method(http.MethodDelete, "/delete-mms-messages", http.HandlerFunc(gw.DeleteMMS))
	r.Method(http.MethodPut, "/update-mms", http.HandlerFunc(gw.UpdateMMS))

	// Contact routes
	r.Method(http.MethodPost, "/add-contact", http.HandlerFunc(gw.AddContact))
	r.Method(http.MethodGet, "/contacts", http.HandlerFunc(gw.GetContacts))
	r.Method(http.MethodDelete, "/delete-contacts", http.HandlerFunc(gw.DeleteContacts))
	r.Method(http.MethodPut, "/update-contact", http.HandlerFunc(gw.UpdateContact))

	// Start the server
	logger.Info("Starting server on :8080")
	if err := http.ListenAndServe(":8080", r); err != nil {
		logger.Error("could not start server", slog.Any("error", err))
	}
}
