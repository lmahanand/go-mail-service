package model

// Email Model Structure
type Email struct {
	From          string
	To            []string
	Cc            []string
	Bcc           []string
	Subject       string
	TextBody      string
	HTMLBody      string
	Status        string
	ScheduledTime string
}
