package routes

import (
	"github.com/Vivirinter/virtual-gsm-gateway/internal/gateway"
	"github.com/go-chi/chi/v5"
	"net/http"
)

func RegisterRoutes(r chi.Router, gw *gateway.Gateway) {
	registerSMSRoutes(r, gw)
	registerUSSDRoutes(r, gw)
	registerMMSRoutes(r, gw)
	registerContactRoutes(r, gw)
}

func registerSMSRoutes(r chi.Router, gw *gateway.Gateway) {
	r.Route("/sms", func(r chi.Router) {
		r.Method(http.MethodPost, "/send", http.HandlerFunc(gw.SendSMS))
		r.Method(http.MethodGet, "/messages", http.HandlerFunc(gw.GetSMS))
		r.Method(http.MethodDelete, "/delete", http.HandlerFunc(gw.DeleteSMS))
		r.Method(http.MethodPut, "/update", http.HandlerFunc(gw.UpdateSMS))
	})
}

func registerUSSDRoutes(r chi.Router, gw *gateway.Gateway) {
	r.Route("/ussd", func(r chi.Router) {
		r.Method(http.MethodPost, "/send", http.HandlerFunc(gw.SendUSSD))
		r.Method(http.MethodGet, "/requests", http.HandlerFunc(gw.GetUSSD))
		r.Method(http.MethodDelete, "/delete", http.HandlerFunc(gw.DeleteUSSD))
		r.Method(http.MethodPut, "/update", http.HandlerFunc(gw.UpdateUSSD))
	})
}

func registerMMSRoutes(r chi.Router, gw *gateway.Gateway) {
	r.Route("/mms", func(r chi.Router) {
		r.Method(http.MethodPost, "/send", http.HandlerFunc(gw.SendMMS))
		r.Method(http.MethodGet, "/messages", http.HandlerFunc(gw.GetMMS))
		r.Method(http.MethodDelete, "/delete", http.HandlerFunc(gw.DeleteMMS))
		r.Method(http.MethodPut, "/update", http.HandlerFunc(gw.UpdateMMS))
	})
}

func registerContactRoutes(r chi.Router, gw *gateway.Gateway) {
	r.Route("/contacts", func(r chi.Router) {
		r.Method(http.MethodPost, "/add", http.HandlerFunc(gw.AddContact))
		r.Method(http.MethodGet, "/", http.HandlerFunc(gw.GetContacts))
		r.Method(http.MethodDelete, "/delete", http.HandlerFunc(gw.DeleteContacts))
		r.Method(http.MethodPut, "/update", http.HandlerFunc(gw.UpdateContact))
	})
}
