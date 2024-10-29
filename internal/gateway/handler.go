package gateway

import (
	"encoding/json"
	"log"
	"net/http"
)

type Gateway struct {
	messages     []SMS
	ussdRequests []USSD
	mmsMessages  []MMS
	contacts     []Contact
	logger       *log.Logger
}

func NewGateway(logger *log.Logger) *Gateway {
	return &Gateway{
		messages:     []SMS{},
		ussdRequests: []USSD{},
		mmsMessages:  []MMS{},
		contacts:     []Contact{},
		logger:       logger,
	}
}

func decodeJSON(w http.ResponseWriter, r *http.Request, v interface{}) bool {
	if err := json.NewDecoder(r.Body).Decode(v); err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return false
	}
	return true
}

func encodeJSON(w http.ResponseWriter, v interface{}, logger *log.Logger) bool {
	if err := json.NewEncoder(w).Encode(v); err != nil {
		http.Error(w, "Error encoding response", http.StatusInternalServerError)
		logger.Printf("Error encoding response: %v", err)
		return false
	}
	return true
}

// Helper function to log and send error response
func logAndRespondError(w http.ResponseWriter, logger *log.Logger, message string, statusCode int, err error) {
	http.Error(w, message, statusCode)
	logger.Printf("%s: %v", message, err)
}

func (g *Gateway) SendSMS(w http.ResponseWriter, r *http.Request) {
	var sms SMS
	if !decodeJSON(w, r, &sms) {
		g.logger.Printf("Error decoding SMS")
		return
	}
	g.messages = append(g.messages, sms)
	w.WriteHeader(http.StatusCreated)
	g.logger.Printf("SMS received: %+v", sms)
}

func (g *Gateway) GetSMS(w http.ResponseWriter, r *http.Request) {
	encodeJSON(w, g.messages, g.logger)
}

func (g *Gateway) DeleteSMS(w http.ResponseWriter, r *http.Request) {
	g.messages = []SMS{}
	w.WriteHeader(http.StatusOK)
	g.logger.Println("All SMS messages deleted")
}

func (g *Gateway) UpdateSMS(w http.ResponseWriter, r *http.Request) {
	var sms SMS
	if !decodeJSON(w, r, &sms) {
		g.logger.Printf("Error decoding SMS")
		return
	}

	for i, msg := range g.messages {
		if msg.From == sms.From && msg.To == sms.To {
			g.messages[i] = sms
			w.WriteHeader(http.StatusOK)
			g.logger.Printf("SMS updated: %+v", sms)
			return
		}
	}

	http.Error(w, "SMS not found", http.StatusNotFound)
	g.logger.Printf("SMS not found for update: %+v", sms)
}

func (g *Gateway) SendUSSD(w http.ResponseWriter, r *http.Request) {
	var ussd USSD
	if !decodeJSON(w, r, &ussd) {
		g.logger.Printf("Error decoding USSD")
		return
	}
	ussd.Response = "USSD response for code: " + ussd.Code
	g.ussdRequests = append(g.ussdRequests, ussd)
	if !encodeJSON(w, ussd, g.logger) {
		return
	}
	g.logger.Printf("USSD received: %+v", ussd)
}

func (g *Gateway) GetUSSD(w http.ResponseWriter, r *http.Request) {
	encodeJSON(w, g.ussdRequests, g.logger)
}

func (g *Gateway) DeleteUSSD(w http.ResponseWriter, r *http.Request) {
	g.ussdRequests = []USSD{}
	w.WriteHeader(http.StatusOK)
	g.logger.Println("All USSD requests deleted")
}

func (g *Gateway) UpdateUSSD(w http.ResponseWriter, r *http.Request) {
	var ussd USSD
	if !decodeJSON(w, r, &ussd) {
		g.logger.Printf("Error decoding USSD")
		return
	}

	for i, req := range g.ussdRequests {
		if req.From == ussd.From && req.Code == ussd.Code {
			g.ussdRequests[i] = ussd
			w.WriteHeader(http.StatusOK)
			g.logger.Printf("USSD updated: %+v", ussd)
			return
		}
	}

	http.Error(w, "USSD not found", http.StatusNotFound)
	g.logger.Printf("USSD not found for update: %+v", ussd)
}

func (g *Gateway) SendMMS(w http.ResponseWriter, r *http.Request) {
	var mms MMS
	if !decodeJSON(w, r, &mms) {
		g.logger.Printf("Error decoding MMS")
		return
	}
	g.mmsMessages = append(g.mmsMessages, mms)
	w.WriteHeader(http.StatusCreated)
	g.logger.Printf("MMS received: %+v", mms)
}

func (g *Gateway) GetMMS(w http.ResponseWriter, r *http.Request) {
	encodeJSON(w, g.mmsMessages, g.logger)
}

func (g *Gateway) DeleteMMS(w http.ResponseWriter, r *http.Request) {
	g.mmsMessages = []MMS{}
	w.WriteHeader(http.StatusOK)
	g.logger.Println("All MMS messages deleted")
}

func (g *Gateway) UpdateMMS(w http.ResponseWriter, r *http.Request) {
	var mms MMS
	if !decodeJSON(w, r, &mms) {
		g.logger.Printf("Error decoding MMS")
		return
	}

	for i, msg := range g.mmsMessages {
		if msg.From == mms.From && msg.To == mms.To {
			g.mmsMessages[i] = mms
			w.WriteHeader(http.StatusOK)
			g.logger.Printf("MMS updated: %+v", mms)
			return
		}
	}

	http.Error(w, "MMS not found", http.StatusNotFound)
	g.logger.Printf("MMS not found for update: %+v", mms)
}

func (g *Gateway) AddContact(w http.ResponseWriter, r *http.Request) {
	var contact Contact
	if !decodeJSON(w, r, &contact) {
		g.logger.Printf("Error decoding contact")
		return
	}
	g.contacts = append(g.contacts, contact)
	w.WriteHeader(http.StatusCreated)
	g.logger.Printf("Contact added: %+v", contact)
}

func (g *Gateway) GetContacts(w http.ResponseWriter, r *http.Request) {
	encodeJSON(w, g.contacts, g.logger)
}

func (g *Gateway) DeleteContacts(w http.ResponseWriter, r *http.Request) {
	g.contacts = []Contact{}
	w.WriteHeader(http.StatusOK)
	g.logger.Println("All contacts deleted")
}

func (g *Gateway) UpdateContact(w http.ResponseWriter, r *http.Request) {
	var contact Contact
	if !decodeJSON(w, r, &contact) {
		g.logger.Printf("Error decoding contact")
		return
	}

	for i, c := range g.contacts {
		if c.Phone == contact.Phone {
			g.contacts[i] = contact
			w.WriteHeader(http.StatusOK)
			g.logger.Printf("Contact updated: %+v", contact)
			return
		}
	}

	http.Error(w, "Contact not found", http.StatusNotFound)
	g.logger.Printf("Contact not found for update: %+v", contact)
}
