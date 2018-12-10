package service

import (
	"encoding/json"
	"fmt"
	"log"
	"os"

	sendgrid "github.com/sendgrid/sendgrid-go"
)

//SendGridEmail ...
/*
	JSON format of SendGridEmail
	{
  		"personalizations": [
		{
		"to": [
			{
			"email": "<EMAIL_ID>"
			},
			{
			"email": "<EMAIL_ID>"
			}
		],
		"cc": [
			{
			"email": "<EMAIL_ID>"
			},
			{
			"email": "<EMAIL_ID>"
			}
		],
		"bcc": [],
		"subject": "Assignment Testing mail with CC"
		}
		],
		"from": {
			"email": "<EMAIL_ID>"
		},
		"content": [
			{
			"type": "text/plain",
			"value": "and easy to do anywhere, even with Go"
			}
		]
	}
*/
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
var SendEmail = func() {
	from := From{Email: "lmahanand@gmail.com"}
	to := [...]To{
		{Email: "lmahanand2010@gmail.com"},
	}

	ccc := [...]Cc{
		{Email: "raj.solidity@gmail.com"},
	}

	subject := "Assignment Testing mail with CC"

	content := [...]Content{
		{
			Type:  "text/plain",
			Value: "and easy to do anywhere, even with Go",
		},
	}

	personalizations := []Personalization{
		Personalization{
			To:      to[:],
			Cc:      ccc[:],
			Subject: subject,
		},
	}

	sge := SendGridEmail{
		Personalizations: personalizations,
		From:             from,
		Content:          content[:],
	}

	b, _ := json.MarshalIndent(sge, "", "")

	s := string(b)
	fmt.Println("\n", s)
	fmt.Printf("\n")

	request := sendgrid.GetRequest(os.Getenv("SENDGRID_API_KEY"), "/v3/mail/send", "https://api.sendgrid.com")
	request.Method = "POST"

	request.Body = []byte(s)
	response, err := sendgrid.API(request)
	if err != nil {
		log.Println(err)
	} else {
		fmt.Println(response.StatusCode)
		fmt.Println(response.Body)
		fmt.Println(response.Headers)
	}
}
