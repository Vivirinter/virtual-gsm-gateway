package gateway

type Contact struct {
	ID    string `json:"id"`
	Name  string `json:"name"`
	Phone string `json:"phone"`
	Email string `json:"email"`
}

type SMS struct {
	From string `json:"from"`
	To   string `json:"to"`
	Body string `json:"body"`
}

type USSD struct {
	Code     string `json:"code"`
	Response string `json:"response"`
}

type MMS struct {
	From    string `json:"from"`
	To      string `json:"to"`
	Subject string `json:"subject"`
	Body    string `json:"body"`
}
