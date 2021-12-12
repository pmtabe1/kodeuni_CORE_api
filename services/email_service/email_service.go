package email_service

import (
	"context"
	"encoding/base64"
	"errors"
	"fmt"
	"log"
	"net/smtp"
	"os"
	"time"

	"github.com/joho/godotenv"
	"github.com/paulmsegeya/pos/services/services_utils"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
	"google.golang.org/api/gmail/v1"
	"google.golang.org/api/option"
)

type IEmailService interface {
}

type EmailService struct {
}

var emailAuth smtp.Auth

// GmailService : Gmail client for sending email
var GmailService *gmail.Service

func OAuthGmailService() {

	if len(os.Getenv("GMAIL_CLIENT_ID")) == 0 {

		log.Panicln("GMAIL_CLIENT_ID is not Set please set it")
	}

	if len(os.Getenv("GMAIL_CLIENT_SECRET")) == 0 {

		log.Panicln("GMAIL_CLIENT_SECRET is not Set please set it")
	}

	if len(os.Getenv("GMAIL_OAUTH_URL")) == 0 {
		os.Setenv("DEVMODE", "DEV")
		log.Println("GMAIL_OAUTH_URL is not Set please set it :: SETTING IT")

		if len(os.Getenv("DEVMODE")) > 0 {
			//os.Setenv("OAUTH_URL", "http://localhost:9000/redirect2Oauth")
			log.Println("Setting Devlopment CLIENT URL")
			os.Setenv("GMAIL_OAUTH_URL", "https://developers.google.com/oauthplayground/")

		} else {
			log.Println("Setting Production CLIENT URL")
			os.Setenv("GMAIL_OAUTH_URL", "https://developers.google.com/oauthplayground/")

		}
	}

	config := oauth2.Config{
		ClientID:     os.Getenv("GMAIL_CLIENT_ID"),
		ClientSecret: os.Getenv("GMAIL_CLIENT_SECRET"),
		Endpoint:     google.Endpoint,
		RedirectURL:  os.Getenv("GMAIL_CLIENT_SECRET"),
	}

	token := oauth2.Token{
		AccessToken:  os.Getenv("GMAIL_ACCESS_TOKEN"),
		RefreshToken: os.Getenv("GMAIL_REFRESH_TOKEN"),
		TokenType:    "Bearer",
		Expiry:       time.Now(),
	}

	var tokenSource = config.TokenSource(context.Background(), &token)

	srv, err := gmail.NewService(context.Background(), option.WithTokenSource(tokenSource))
	if err != nil {
		log.Printf("Unable to retrieve Gmail client: %v", err)
	}

	GmailService = srv
	if GmailService != nil {
		fmt.Println("Email service is initialized ")
	}
}

func SendEmailOAUTH2(to string, data interface{}, template string) (bool, error) {

	emailBody, err := services_utils.ParseTemplate(template, data)
	if err != nil {
		return false, errors.New("unable to parse email template")
	}

	var message gmail.Message

	emailTo := "To: " + to + "\r\n"
	subject := "Subject: " + "Test Email form Gmail API using OAuth" + "\n"
	mime := "MIME-version: 1.0;\nContent-Type: text/plain; charset=\"UTF-8\";\n\n"
	msg := []byte(emailTo + subject + mime + "\n" + emailBody)

	message.Raw = base64.URLEncoding.EncodeToString(msg)

	// Send the message
	_, err = GmailService.Users.Messages.Send("me", &message).Do()
	if err != nil {
		return false, err
	}
	return true, nil
}

func SendEmailSMTP(to []string, data interface{}, template string) (bool, error) {
	log.Println(">>>> sENDING EMAIL VIA SMTP >>>")
	os.Setenv("EMAIL_PASSWORD", "s@mb@mb@000")
	os.Setenv("EMAIL_HOST", "smtp.gmail.com")
	os.Setenv("EMAIL_FROM", "tiradevwork@gmail.com")
	os.Setenv("EMAIL_PORT", "587")
	os.Setenv("MAIL_TO", "paul@duxte.com")
	emailHost := os.Getenv("EMAIL_HOST")
	emailFrom := os.Getenv("EMAIL_FROM")
	emailPassword := os.Getenv("EMAIL_PASSWORD")
	emailPort := os.Getenv("EMAIL_PORT")

	emailAuth = smtp.PlainAuth("", emailFrom, emailPassword, emailHost)

	emailBody, err := services_utils.ParseTemplate(template, data)
	if err != nil {
		log.Panicln(err.Error())
		return false, errors.New("unable to parse email template")
	}

	mime := "MIME-version: 1.0;\nContent-Type: text/plain; charset=\"UTF-8\";\n\n"
	subject := "Subject: " + "Test Email" + "!\n"
	msg := []byte(subject + mime + "\n" + emailBody)
	addr := fmt.Sprintf("%s:%s", emailHost, emailPort)
	to = append(to, os.Getenv("MAIL_TO"))

	if err := smtp.SendMail(addr, emailAuth, emailFrom, to, msg); err != nil {
		log.Panicln(err)
		return false, err
	}
	log.Println("Successfully sent email...")
	return true, nil
}

func StartSendingEmails(sender string, receiver string, message string, inputMethod string) (status bool) {

	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error getting env, %v", err)
	}

	params := os.Args
	paramsLength := len(params)
	if paramsLength < 2 {
		log.Println("Please add SMTP or OAUTH along with go run main.go command")
		log.Println("Eg: go run main.go SMTP or go run main.go OAUTH")
		os.Exit(1)
	}

	if len(inputMethod) == 0 {
		inputMethod = os.Args[1]

	}

	valid := IsValidInputMethod(inputMethod)

	emailTo := os.Getenv("EMAIL_TO")

	if valid {

		if len(receiver) == 0 {

			receiver = "Paul"
		}

		if len(sender) == 0 {
			sender = "Doe"
		}

		data := struct {
			ReceiverName string
			SenderName   string
		}{
			ReceiverName: receiver,
			SenderName:   sender,
		}

		if inputMethod == "SMTP" {
			status, err := SendEmailSMTP([]string{emailTo}, data, "email_template.txt")
			if err != nil {
				log.Println(err)
			}
			if status {
				log.Println("Email sent successfully using SMTP")
			}
		}

		if inputMethod == "OAUTH" {
			OAuthGmailService()
			status, err := SendEmailOAUTH2(emailTo, data, "email_template.txt")
			if err != nil {
				log.Println(err)
			}
			if status {
				log.Println("Email sent successfully using OAUTH")
			}

		}
	} else {
		log.Println("Please add SMTP or OAUTH along with go run main.go command")
		log.Println("Eg: go run main.go SMTP or go run main.go OAUTH")
		os.Exit(1)
	}

	return status
}

func IsValidInputMethod(method string) bool {
	switch method {
	case
		"SMTP",
		"OAUTH":
		return true
	}
	return false
}
