package main

import (
	"log"
	"net/http"
	"os"

	"github.com/Vivirinter/virtual-gsm-gateway/internal/gateway"
)

func main() {
	logger := log.New(os.Stdout, "http: ", log.LstdFlags)
	gw := gateway.NewGateway(logger)

	mux := http.NewServeMux()

	// SMS routes
	mux.HandleFunc("/send-sms", gw.SendSMS)
	mux.HandleFunc("/messages", gw.GetSMS)
	mux.HandleFunc("/delete-messages", gw.DeleteSMS)
	mux.HandleFunc("/update-sms", gw.UpdateSMS)

	// USSD routes
	mux.HandleFunc("/send-ussd", gw.SendUSSD)
	mux.HandleFunc("/ussd-requests", gw.GetUSSD)
	mux.HandleFunc("/delete-ussd-requests", gw.DeleteUSSD)
	mux.HandleFunc("/update-ussd", gw.UpdateUSSD)

	// MMS routes
	mux.HandleFunc("/send-mms", gw.SendMMS)
	mux.HandleFunc("/mms-messages", gw.GetMMS)
	mux.HandleFunc("/delete-mms-messages", gw.DeleteMMS)
	mux.HandleFunc("/update-mms", gw.UpdateMMS)

	// Contact routes
	mux.HandleFunc("/add-contact", gw.AddContact)
	mux.HandleFunc("/contacts", gw.GetContacts)
	mux.HandleFunc("/delete-contacts", gw.DeleteContacts)
	mux.HandleFunc("/update-contact", gw.UpdateContact)

	// Start the server
	logger.Println("Starting server on :8080")
	if err := http.ListenAndServe(":8080", mux); err != nil {
		logger.Fatalf("could not start server: %v\n", err)
	}
}
