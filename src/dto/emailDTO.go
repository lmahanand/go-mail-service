package dto

// EmailDTO Model Structure
type EmailDTO struct {
	From          string   `json:"from"`
	To            []string `json:"to"`
	Cc            []string `json:"cc"`
	Bcc           []string `json:"bcc"`
	Subject       string   `json:"subject"`
	Body          string   `json:"body"`
	Status        string   `json:"status"`
	ScheduledTime string   `json:"scheduledTime"`
}
