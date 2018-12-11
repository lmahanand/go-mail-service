# go-mail-service

## Steps to run the app

### This app accepts the required information about an email and schedules and sends emails in the future. It has also implemented a scheduler that could be set to send emails in future. It accepts the scheudling time in UTC format. E.g. 09 Dec 18 4:36 UTC. It uses email services from SendGrid (https://sendgrid.com) and AWS SES (https://aws.amazon.com/ses/). In case one service is failed to deliver then the application failovers to another available service to deliver the emails.

#### For future message scheduling according to given time, library from github.com/robfig/cron is used.

1. clone the repository to local directory
2. go to src folder and run the below commands

    go get github.com/aws/aws-sdk-go/aws

    go get github.com/gorilla/mux

    go get github.com/robfig/cron

    go get github.com/sendgrid/sendgrid-go

3. run the command go run main.go to start the application. And application will start at 8081 port

4. To send an email hit the URL http://localhost:8081/email with POST method

```json
{
    "to": [
        "email1@gmail.com"
    ],
    "cc": [
        "email2@gmail.com"
    ],
    "bcc": [
        "email3@yourdomain.com"
    ],
    "subject": "Hello World - Email-Failover-Service",
  "content": [
    {
      "type": "text/plain",
      "value": "This app could be used to send emails using Send Grid and AWS SES."
    },
    {
      "type": "text/html",
      "value": "<h3>This app could be used to send emails using Send Grid and AWS SES.</h3>"
    }
  ],
    "scheduledTime":"11 Dec 18 4:36 UTC"
}
```
