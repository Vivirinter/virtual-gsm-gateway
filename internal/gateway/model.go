package gateway

type SMS struct {
	From    string `json:"from"`
	To      string `json:"to"`
	Message string `json:"message"`
}

type USSD struct {
	From     string `json:"from"`
	Code     string `json:"code"`
	Response string `json:"response"`
}

type MMS struct {
	From    string   `json:"from"`
	To      string   `json:"to"`
	Message string   `json:"message"`
	Media   []string `json:"media"`
}

type Contact struct {
	Name  string `json:"name"`
	Phone string `json:"phone"`
}
