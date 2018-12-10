package model

//Content ...
type Content struct {
	Type  string `json:"type"`
	Value string `json:"value"`
}

// Email Model Structure
type Email struct {
	From          string
	To            []string
	Cc            []string
	Bcc           []string
	Subject       string
	Content       []Content
	Status        string
	ScheduledTime string
}
