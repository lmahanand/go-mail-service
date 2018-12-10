package service

import (
	"fmt"

	m "../model"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/awserr"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/ses"
)

const (
	//Sender : This address must be verified with Amazon SES.
	Sender = "lmahanand@gmail.com"

	//CharSet ...
	// The character encoding for the email.
	CharSet = "UTF-8"
)

//SendEmailUsingAmazonSES ...
var SendEmailUsingAmazonSES = func(email m.Email) (*string, error) {
	sess, err := session.NewSession(&aws.Config{
		Region: aws.String("us-east-1")},
	)

	// Create an SES session.
	svc := ses.New(sess)

	//Extraction of Model Email

	lenOfTo := len(email.To)
	to := make([]string, lenOfTo)

	for i, t := range email.To {
		to[i] = t
	}

	lenOfCc := len(email.Cc)
	cc := make([]string, lenOfCc)
	for i, c := range email.Cc {
		cc[i] = c
	}

	subject := email.Subject
	var HTMLBody, TextBody string
	for _, con := range email.Content {
		if con.Type == "text/html" {
			HTMLBody = con.Value
		} else if con.Type == "text/plain" {
			TextBody = con.Value
		}
	}

	// Assemble the email.
	input := &ses.SendEmailInput{
		Destination: &ses.Destination{
			CcAddresses: []*string{aws.String(cc[0])},
			ToAddresses: []*string{
				aws.String(to[0]),
			},
		},
		Message: &ses.Message{
			Body: &ses.Body{
				Html: &ses.Content{
					Charset: aws.String(CharSet),
					Data:    aws.String(HTMLBody),
				},
				Text: &ses.Content{
					Charset: aws.String(CharSet),
					Data:    aws.String(TextBody),
				},
			},
			Subject: &ses.Content{
				Charset: aws.String(CharSet),
				Data:    aws.String(subject + " : Using Amazon SES"),
			},
		},
		Source: aws.String(Sender),
		// Uncomment to use a configuration set
		//ConfigurationSetName: aws.String(ConfigurationSet),
	}

	// Attempt to send the email.
	result, err := svc.SendEmail(input)

	// Display error messages if they occur.
	if err != nil {
		if aerr, ok := err.(awserr.Error); ok {
			switch aerr.Code() {
			case ses.ErrCodeMessageRejected:
				fmt.Println(ses.ErrCodeMessageRejected, aerr.Error())
			case ses.ErrCodeMailFromDomainNotVerifiedException:
				fmt.Println(ses.ErrCodeMailFromDomainNotVerifiedException, aerr.Error())
			case ses.ErrCodeConfigurationSetDoesNotExistException:
				fmt.Println(ses.ErrCodeConfigurationSetDoesNotExistException, aerr.Error())
			default:
				fmt.Println(aerr.Error())
			}
		} else {
			// Print the error, cast err to awserr.Error to get the Code and
			// Message from an error.
			fmt.Println(err.Error())
		}

		return result.MessageId, err
	}

	fmt.Printf("Email Sent to addresses: %v", email.To)
	fmt.Println(result)
	return result.MessageId, err
}
