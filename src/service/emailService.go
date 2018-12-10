package service

import (
	dto "../dto"
	m "../model"
	repo "../repository"
	u "../utils"
)

//EmailService service
type EmailService struct{}

//Emails Repository against each email id
var Emails = repo.Emails

var isSendGridActive = true

//SendEmail service
func (emailService *EmailService) SendEmail(emailDTO dto.EmailDTO) map[string]interface{} {

	if resp, ok := emailService.Validate(emailDTO); !ok {
		return resp
	}

	emailID := emailDTO.From
	lenOfContent := len(emailDTO.Content)
	content := make([]m.Content, lenOfContent)

	for i, c := range emailDTO.Content {
		content[i] = m.Content{Type: c.Type, Value: c.Value}
	}

	email := m.Email{
		From:          emailID,
		To:            emailDTO.To,
		Cc:            emailDTO.Cc,
		Bcc:           emailDTO.Bcc,
		Subject:       emailDTO.Subject,
		Content:       content[:],
		Status:        m.SCHEDULED,
		ScheduledTime: emailDTO.ScheduledTime,
	}
	emails := Emails[emailID]

	emails = append(emails, email)

	Emails[emailID] = emails

	SendEmailUsingSendGridServer(email)

	resp := u.Message(true, m.SCHEDULED)
	return resp
}

//GetEmails service method
func (emailService *EmailService) GetEmails(emaildID string) []m.Email {

	emails := Emails[emaildID]

	return emails
}

//Validate method to validate all required fields
func (emailService *EmailService) Validate(emailDTO dto.EmailDTO) (map[string]interface{}, bool) {
	if emailDTO.From == "" {
		return u.Message(false, "From email id is empty in the payload"), false
	}

	if len(emailDTO.To) == 0 {
		return u.Message(false, "Recepient email id is empty in the payload"), false
	}

	//All the required parameters are present
	return u.Message(true, "success"), true
}
