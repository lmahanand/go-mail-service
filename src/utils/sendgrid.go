package utils

import (
	"fmt"
	"log"
	"os"

	"github.com/sendgrid/sendgrid-go"
)

//SendGridEmail structure to send an email
type SendGridEmail struct {
	Personalizations []struct {
		To []struct {
			Email string `json:"email"`
		} `json:"to"`
		Cc []struct {
			Email string `json:"email"`
		} `json:"cc"`
		Subject string `json:"subject"`
	} `json:"personalizations"`
	From struct {
		Email string `json:"email"`
	} `json:"from"`
	Content []struct {
		Type  string `json:"type"`
		Value string `json:"value"`
	} `json:"content"`
}

//SendEmail service
var SendEmail = func() {
	request := sendgrid.GetRequest(os.Getenv("SENDGRID_API_KEY"), "/v3/mail/send", "https://api.sendgrid.com")
	request.Method = "POST"
	request.Body = []byte(` {
	"personalizations": [
		{
			"to": [
				{
					"email": "lmahanand2010@gmail.com"
				}
			],
			"cc": [
				{
					"email": "raj.solidity@gmail.com"
				}
			],
			"subject": "Assignment Testing mail with CC"
		}
	],
	"from": {
		"email": "lmahanand@gmail.com"
	},
	"content": [
		{
			"type": "text/plain",
			"value": "and easy to do anywhere, even with Go"
		}
	]
}`)
	response, err := sendgrid.API(request)
	if err != nil {
		log.Println(err)
	} else {
		fmt.Println(response.StatusCode)
		fmt.Println(response.Body)
		fmt.Println(response.Headers)
	}
}
