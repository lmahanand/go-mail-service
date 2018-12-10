package service

import (
	"encoding/json"
	"log"
	"os"

	m "../model"
	sendgrid "github.com/sendgrid/sendgrid-go"
)

//SendGridEmail ...
type SendGridEmail struct {
	Personalizations []Personalization `json:"personalizations"`
	From             From              `json:"from"`
	Content          []Content         `json:"content"`
}

//Personalization ...
type Personalization struct {
	To      []To   `json:"to"`
	Cc      []Cc   `json:"cc"`
	Subject string `json:"subject"`
}

//To ...
type To struct {
	Email string `json:"email"`
}

//Cc ...
type Cc struct {
	Email string `json:"email"`
}

//Bcc ...
type Bcc struct {
	Email string `json:"email"`
}

//From ...
type From struct {
	Email string `json:"email"`
}

//Content ...
type Content struct {
	Type  string `json:"type"`
	Value string `json:"value"`
}

//SendEmail service
var SendEmailUsingSendGridServer = func(email m.Email) (int, error) {
	from := From{Email: Sender}

	lenOfTo := len(email.To)
	to := make([]To, lenOfTo)
	for i, t := range email.To {
		to[i] = To{Email: t}
	}

	lenOfCc := len(email.Cc)
	cc := make([]Cc, lenOfCc)
	for i, c := range email.Cc {
		cc[i] = Cc{Email: c}
	}

	subject := email.Subject

	lenOfContent := len(email.Content)
	content := make([]Content, lenOfContent)

	for i, c := range email.Content {
		content[i] = Content{Type: c.Type, Value: c.Value}
	}

	personalizations := []Personalization{
		Personalization{
			To:      to[:],
			Cc:      cc[:],
			Subject: subject + " : Using Send GRID",
		},
	}

	sge := SendGridEmail{
		Personalizations: personalizations,
		From:             from,
		Content:          content[:],
	}
	b, _ := json.Marshal(sge)
	s := string(b)

	request := sendgrid.GetRequest(os.Getenv("SENDGRID_API_KEY"), "/v3/mail/send", "https://api.sendgrid.com")
	request.Method = "POST"

	request.Body = []byte(s)
	response, err := sendgrid.API(request)
	if err != nil {
		log.Println(err)
		return response.StatusCode, err
	}

	return response.StatusCode, err
}
