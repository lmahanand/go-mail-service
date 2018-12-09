package utils

import (
	"fmt"
	"log"
	"os"

	"github.com/sendgrid/sendgrid-go"
)

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
