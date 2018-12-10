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

	emailID := Sender
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

	resp := u.Message(true, m.SCHEDULED)

	isEmailToBeSentUsingSendGrid := true

	if isEmailToBeSentUsingSendGrid {
		res, err := SendEmailUsingSendGridServer(email)

		if err != nil || res == 400 {
			log.Printf("Could not use Send Grid server hence using Amazon Email Service %v", err)
			isEmailToBeSentUsingSendGrid = false
		}

		if res == 202 {
			email.Status = m.SENT
			resp = u.Message(true, m.SENT)
		}
	}

	// Send email using Amazon SES if Send Grid has failed to deliver
	if !isEmailToBeSentUsingSendGrid {
		awsRes, awsErr := SendEmailUsingAmazonSES(email)
		if awsErr != nil {
			resp = u.Message(true, m.FAILED)
			email.Status = m.FAILED

		} else if awsRes != nil {
			resp = u.Message(true, m.SENT)
			email.Status = m.SENT
		}
	}

	emails := Emails[emailID]

	emails = append(emails, email)

	Emails[emailID] = emails
	return resp

}

//GetEmails service method
func (emailService *EmailService) GetEmails() []m.Email {

	emails := Emails[Sender]

	return emails
}

//Validate method to validate all required fields
func (emailService *EmailService) Validate(emailDTO dto.EmailDTO) (map[string]interface{}, bool) {
	if len(emailDTO.To) == 0 {
		return u.Message(false, "Recepient email id is empty in the payload"), false
	}

	//All the required parameters are present
	return u.Message(true, "success"), true
}
