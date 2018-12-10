package service

import (
	"log"

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

	isEmailSentUsingSendGrid := true

	res, err := SendEmailUsingSendGridServer(email)
	println("err ", err)
	if err != nil || res == 400 {
		log.Printf("Could not use Send Grid server hence using Amazon Email Service %v", err)
		isEmailSentUsingSendGrid = false
	}

	// Send email using Amazon SES if Send Grid has failed to deliver
	if !isEmailSentUsingSendGrid {
		awsRes, awsErr := SendEmailUsingAmazonSES(email)
		if awsErr != nil {
			resp := u.Message(true, m.FAILED)
			return resp
		} else if awsRes != nil {
			resp := u.Message(true, m.SENT)
			return resp
		}
	}
	if res == 202 {
		resp := u.Message(true, m.SENT)
		return resp
	}

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
