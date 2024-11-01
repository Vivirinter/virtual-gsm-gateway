package gateway

import (
	"encoding/json"
	"log/slog"
	"net/http"
)

type Gateway struct {
	messages     []SMS
	ussdRequests []USSD
	mmsMessages  []MMS
	contacts     []Contact
	logger       *slog.Logger
}

func NewGateway(logger *slog.Logger) *Gateway {
	return &Gateway{
		logger: logger,
	}
}

func logAndRespondError(w http.ResponseWriter, logger *slog.Logger, userMessage string, statusCode int, err error) {
	http.Error(w, userMessage, statusCode)
	logger.Error("Internal error", "error", err)
}

func (g *Gateway) SendSMS(w http.ResponseWriter, r *http.Request) {
	var sms SMS
	if err := json.NewDecoder(r.Body).Decode(&sms); err != nil {
		logAndRespondError(w, g.logger, "Invalid request payload", http.StatusBadRequest, err)
		return
	}
	g.messages = append(g.messages, sms)
	w.WriteHeader(http.StatusCreated)
	g.logger.Info("SMS received", "sms", sms)
}

func (g *Gateway) GetSMS(w http.ResponseWriter, r *http.Request) {
	if err := json.NewEncoder(w).Encode(g.messages); err != nil {
		logAndRespondError(w, g.logger, "Error encoding response", http.StatusInternalServerError, err)
	}
}

func (g *Gateway) DeleteSMS(w http.ResponseWriter, r *http.Request) {
	g.messages = []SMS{}
	w.WriteHeader(http.StatusOK)
	g.logger.Info("All SMS messages deleted")
}

func (g *Gateway) UpdateSMS(w http.ResponseWriter, r *http.Request) {
	var sms SMS
	if err := json.NewDecoder(r.Body).Decode(&sms); err != nil {
		logAndRespondError(w, g.logger, "Invalid request payload", http.StatusBadRequest, err)
		return
	}

	for i, msg := range g.messages {
		if msg.From == sms.From && msg.To == sms.To {
			g.messages[i] = sms
			w.WriteHeader(http.StatusOK)
			g.logger.Info("SMS updated", "sms", sms)
			return
		}
	}

	http.Error(w, "SMS not found", http.StatusNotFound)
	g.logger.Warn("SMS not found for update", "sms", sms)
}

func (g *Gateway) SendUSSD(w http.ResponseWriter, r *http.Request) {
	var ussd USSD
	if err := json.NewDecoder(r.Body).Decode(&ussd); err != nil {
		logAndRespondError(w, g.logger, "Invalid request payload", http.StatusBadRequest, err)
		return
	}
	ussd.Response = "USSD response for code: " + ussd.Code
	g.ussdRequests = append(g.ussdRequests, ussd)
	if err := json.NewEncoder(w).Encode(ussd); err != nil {
		logAndRespondError(w, g.logger, "Error encoding response", http.StatusInternalServerError, err)
		return
	}
	g.logger.Info("USSD request processed", "ussd", ussd)
}

func (g *Gateway) GetUSSD(w http.ResponseWriter, r *http.Request) {
	if err := json.NewEncoder(w).Encode(g.ussdRequests); err != nil {
		logAndRespondError(w, g.logger, "Error encoding response", http.StatusInternalServerError, err)
	}
}

func (g *Gateway) DeleteUSSD(w http.ResponseWriter, r *http.Request) {
	g.ussdRequests = []USSD{}
	w.WriteHeader(http.StatusOK)
	g.logger.Info("All USSD requests deleted")
}

func (g *Gateway) UpdateUSSD(w http.ResponseWriter, r *http.Request) {
	var ussd USSD
	if err := json.NewDecoder(r.Body).Decode(&ussd); err != nil {
		logAndRespondError(w, g.logger, "Invalid request payload", http.StatusBadRequest, err)
		return
	}

	for i, req := range g.ussdRequests {
		if req.Code == ussd.Code {
			g.ussdRequests[i] = ussd
			w.WriteHeader(http.StatusOK)
			g.logger.Info("USSD updated", "ussd", ussd)
			return
		}
	}

	http.Error(w, "USSD not found", http.StatusNotFound)
	g.logger.Warn("USSD not found for update", "ussd", ussd)
}

func (g *Gateway) SendMMS(w http.ResponseWriter, r *http.Request) {
	var mms MMS
	if err := json.NewDecoder(r.Body).Decode(&mms); err != nil {
		logAndRespondError(w, g.logger, "Invalid request payload", http.StatusBadRequest, err)
		return
	}
	g.mmsMessages = append(g.mmsMessages, mms)
	w.WriteHeader(http.StatusCreated)
	g.logger.Info("MMS received", "mms", mms)
}

func (g *Gateway) GetMMS(w http.ResponseWriter, r *http.Request) {
	if err := json.NewEncoder(w).Encode(g.mmsMessages); err != nil {
		logAndRespondError(w, g.logger, "Error encoding response", http.StatusInternalServerError, err)
	}
}

func (g *Gateway) DeleteMMS(w http.ResponseWriter, r *http.Request) {
	g.mmsMessages = []MMS{}
	w.WriteHeader(http.StatusOK)
	g.logger.Info("All MMS messages deleted")
}

func (g *Gateway) UpdateMMS(w http.ResponseWriter, r *http.Request) {
	var mms MMS
	if err := json.NewDecoder(r.Body).Decode(&mms); err != nil {
		logAndRespondError(w, g.logger, "Invalid request payload", http.StatusBadRequest, err)
		return
	}

	for i, msg := range g.mmsMessages {
		if msg.From == mms.From && msg.To == mms.To {
			g.mmsMessages[i] = mms
			w.WriteHeader(http.StatusOK)
			g.logger.Info("MMS updated", "mms", mms)
			return
		}
	}

	http.Error(w, "MMS not found", http.StatusNotFound)
	g.logger.Warn("MMS not found for update", "mms", mms)
}

func (g *Gateway) AddContact(w http.ResponseWriter, r *http.Request) {
	var contact Contact
	if err := json.NewDecoder(r.Body).Decode(&contact); err != nil {
		logAndRespondError(w, g.logger, "Invalid request payload", http.StatusBadRequest, err)
		return
	}
	g.contacts = append(g.contacts, contact)
	w.WriteHeader(http.StatusCreated)
	g.logger.Info("Contact added", "contact", contact)
}

func (g *Gateway) GetContacts(w http.ResponseWriter, r *http.Request) {
	if err := json.NewEncoder(w).Encode(g.contacts); err != nil {
		logAndRespondError(w, g.logger, "Error encoding response", http.StatusInternalServerError, err)
	}
}

func (g *Gateway) DeleteContacts(w http.ResponseWriter, r *http.Request) {
	g.contacts = []Contact{}
	w.WriteHeader(http.StatusOK)
	g.logger.Info("All contacts deleted")
}

func (g *Gateway) UpdateContact(w http.ResponseWriter, r *http.Request) {
	var contact Contact
	if err := json.NewDecoder(r.Body).Decode(&contact); err != nil {
		logAndRespondError(w, g.logger, "Invalid request payload", http.StatusBadRequest, err)
		return
	}

	for i, c := range g.contacts {
		if c.ID == contact.ID {
			g.contacts[i] = contact
			w.WriteHeader(http.StatusOK)
			g.logger.Info("Contact updated", "contact", contact)
			return
		}
	}

	http.Error(w, "Contact not found", http.StatusNotFound)
	g.logger.Warn("Contact not found for update", "contact", contact)
}
