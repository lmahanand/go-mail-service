package dto

//Content ...
type Content struct {
	Type  string `json:"type"`
	Value string `json:"value"`
}

// EmailDTO Model Structure
type EmailDTO struct {
	From          string    `json:"from"`
	To            []string  `json:"to"`
	Cc            []string  `json:"cc"`
	Bcc           []string  `json:"bcc"`
	Subject       string    `json:"subject"`
	Content       []Content `json:"content"`
	Status        string    `json:"status"`
	ScheduledTime string    `json:"scheduledTime"`
}
