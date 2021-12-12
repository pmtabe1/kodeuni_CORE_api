package email_service

import (
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"net/mail"
	"os"
	"strconv"
	"time"

	"github.com/go-gomail/gomail"
	"github.com/matcornic/hermes/v2"
	"golang.org/x/crypto/ssh/terminal"
)

type example interface {
	Email() hermes.Email
	Name() string
}

type confirmationConfig interface {
	Email() hermes.Email
	Name() string
}

func SendSignupToken() {

}

type MailParam struct {
	ReceiverEmail           string
	SenderEmail             string
	MailIntent              string
	MailSubject             string
	ProductLink             string
	CompanyLogoLink         string
	Copyright               string
	TroubelText             string
	CompanyName             string
	VerificationLink        string
	Intent                  string
	CC                      string
	Company                 string
	Date                    string
	HasTableData            bool
	RecepientName           string
	OutrosMessages          []string
	IntrosMessages          []string
	Signature               string
	Title                   string
	Greeting                string
	IntroBody               string
	WebsiteLink             string
	ValiationLink           string
	ConfirmationLink        string
	OTPData                 string
	InstructionActionText   string
	InviteCode              string
	DashboardButtonLink     string
	DashboardButtonText     string
	DashboardButtonColor    string
	DashboardButtonTexColor string
	CallForAction           string
	EmailSubject            string
	SupportEmail            string
	CopyrightLink           string
	TermsAndConditionLink   string
	MaxTableRows            int
	MaxTableColumns         int
	OutrosMesmiddlewares    []string
	IntrosMesmiddlewares    []string
	KeyValue                map[string]string
	KeyValueTableEntry      []KeyValueTableEntry
	KeyValueTableEntryMap   map[string][]KeyValueTableEntry
}

func BuildKeyValueEntryWithReset(key string, mailParam MailParam, KeyValueTableEntryList []KeyValueTableEntry) (out map[string][]KeyValueTableEntry) {

	if len(KeyValueTableEntryList) == 0 {
		log.Println("KeyValueTableEntryList must be present")
	}

	mailParam.KeyValueTableEntryMap = make(map[string][]KeyValueTableEntry)
	mailParam.KeyValueTableEntryMap[key] = mailParam.KeyValueTableEntry
	//Reset after setting data to map
	mailParam.KeyValueTableEntry = make([]KeyValueTableEntry, 0)

	for _, v := range KeyValueTableEntryList {
		mailParam.KeyValueTableEntry = append(mailParam.KeyValueTableEntry, KeyValueTableEntry{
			Key: v.Key, Value: v.Value,
		})
	}
	out = make(map[string][]KeyValueTableEntry)
	out[key] = mailParam.KeyValueTableEntry
	//reset the list
	KeyValueTableEntryList = make([]KeyValueTableEntry, 0)

	return out

}

func GetMailerDate(mailParam MailParam) (mailerData confirmation) {

	mailerData.CC = mailParam.CC
	mailerData.RecepientName = mailParam.RecepientName
	mailerData.SupportEmail = mailParam.SupportEmail
	mailerData.SenderEmail = mailParam.SenderEmail
	mailerData.MailIntent = mailParam.MailIntent
	mailerData.MailSubject = mailParam.MailSubject
	mailerData.ProductLink = mailParam.ProductLink
	mailerData.CompanyLogoLink = mailParam.CompanyLogoLink
	mailerData.Copyright = mailParam.Copyright
	mailerData.TroubelText = mailParam.TroubelText
	mailerData.CompanyName = mailParam.CompanyName
	mailerData.VerificationLink = mailParam.ValiationLink
	mailerData.Intent = mailParam.Intent
	mailerData.CC = mailParam.CC
	mailerData.Company = mailParam.Company
	mailerData.Date = mailParam.Date
	mailerData.HasTableData = mailParam.HasTableData
	mailerData.RecepientName = mailParam.RecepientName
	mailerData.OutrosMessages = mailParam.OutrosMessages
	mailerData.IntrosMessages = mailParam.IntrosMessages
	mailerData.Signature = mailParam.Signature
	mailerData.Title = mailParam.Title
	mailerData.Greeting = mailParam.Greeting
	mailerData.IntroBody = mailParam.IntroBody
	mailerData.WebsiteLink = mailParam.WebsiteLink
	mailerData.ValiationLink = mailParam.ValiationLink
	mailerData.ConfirmationLink = mailParam.ConfirmationLink
	mailerData.OTPData = mailParam.OTPData
	mailerData.InstructionActionText = mailParam.InstructionActionText
	mailerData.InviteCode = mailParam.InviteCode
	mailerData.DashboardButtonLink = mailParam.DashboardButtonLink
	mailerData.DashboardButtonText = mailParam.DashboardButtonText
	mailerData.DashboardButtonColor = mailParam.DashboardButtonColor
	mailerData.DashboardButtonTexColor = mailParam.DashboardButtonTexColor
	mailerData.CallForAction = mailParam.CallForAction
	mailerData.EmailSubject = mailParam.EmailSubject
	mailerData.SupportEmail = mailParam.SupportEmail
	mailerData.CopyrightLink = mailParam.Copyright
	mailerData.TermsAndConditionLink = mailParam.TermsAndConditionLink
	mailerData.MaxTableRows = mailParam.MaxTableRows
	mailerData.MaxTableColumns = mailParam.MaxTableColumns
	mailerData.KeyValue = mailParam.KeyValue
	mailerData.KeyValueTableEntry = mailParam.KeyValueTableEntry
	//mailerData.KeyValueTableEntryMap = mailerData.KeyValueTableEntryMap

	return mailerData
}

func SendSignupVerificationEmail(receiverEmail string) (status bool) {

	h := hermes.Hermes{
		Product: hermes.Product{
			Name:        "Kodeuni Limited",
			Link:        "https://www.duxte.co.tz",
			Logo:        "https://github.com/pmtabe1/media/blob/main/duxte.png?raw=true",
			Copyright:   "Copyright © 2020 Kodeuni Ltd. All rights reserved.",
			TroubleText: "If you’re having trouble with the button '{ACTION}', copy and paste the URL below into your web browser.",
		},
	}

	sendEmails := os.Getenv("HERMES_SEND_EMAILS") == "true"
	var confirmationData confirmation
	confirmationData.Intent = ConfirmationEmail
	confirmationData.Company = "Duxte Limited"
	confirmationData.Signature = "Paul Msegeya"
	confirmationData.IntrosMessages = []string{"Welcome to Kodeuni we present the courses below for your Kids Future", "Computer Science for Kids", "Codding for Kids", "Critical thinking Courses for Kids"}
	confirmationData.OutrosMessages = []string{"Prepare your child future today ", "Subscribe today and get 10% off ", "From Kodeuni the Platform for Kids Future today.", "You are the change your kid need"}
	confirmationData.ValiationLink = "https://www.duxte.co.tz"
	confirmationData.ConfirmationLink = "https://www.duxte.co.tz"
	confirmationData.DashboardButtonLink = "https://www.duxte.co.tz"
	confirmationData.CopyrightLink = ""
	confirmationData.DashboardButtonText = "OTP"
	confirmationData.DashboardButtonText = "SUBSCRIBE TODAY"
	confirmationData.Date = time.Now().Local().String()
	confirmationData.HasTableData = true
	confirmationData.MaxTableRows = 3
	confirmationData.DashboardButtonColor = "#40b892"
	confirmationData.MaxTableColumns = 0 //
	confirmationData.SupportEmail = "support@duxte.com"
	confirmationData.RecepientName = "paul@duxte.com"
	confirmationData.KeyValue = make(map[string]string)
	confirmationData.KeyValue["price"] = "$10.00"
	confirmationData.KeyValue["item"] = "Some Item"
	confirmationData.KeyValueTableEntry = make([]KeyValueTableEntry, 0)

	confirmationData.KeyValueTableEntry = append(confirmationData.KeyValueTableEntry, KeyValueTableEntry{
		Key: "COURSE", Value: "Computer sciente for KIDS",
	})
	confirmationData.KeyValueTableEntry = append(confirmationData.KeyValueTableEntry, KeyValueTableEntry{
		Key: "Description", Value: "Kids will learn computer science techniques on how to solve day to day problems using computer as a tool.",
	})
	confirmationData.KeyValueTableEntry = append(confirmationData.KeyValueTableEntry, KeyValueTableEntry{
		Key: "Price", Value: "$10.99",
	})

	log.Println(confirmationData.KeyValueTableEntry)

	confirmationData.KeyValueTableEntryMap = make(map[string][]KeyValueTableEntry)
	confirmationData.KeyValueTableEntryMap["row1"] = confirmationData.KeyValueTableEntry
	//Reset after setting data to map
	confirmationData.KeyValueTableEntry = make([]KeyValueTableEntry, 0)

	confirmationData.KeyValueTableEntry = append(confirmationData.KeyValueTableEntry, KeyValueTableEntry{
		Key: "COURSE", Value: "Codding for Kids",
	})
	confirmationData.KeyValueTableEntry = append(confirmationData.KeyValueTableEntry, KeyValueTableEntry{
		Key: "Description", Value: "Kids will learn computer science techniques on how to solve day to day problems using computer as a tool.2",
	})
	confirmationData.KeyValueTableEntry = append(confirmationData.KeyValueTableEntry, KeyValueTableEntry{
		Key: "Price", Value: "$25.99",
	})

	confirmationData.KeyValueTableEntryMap["row2"] = confirmationData.KeyValueTableEntry
	//Reset after setting data to map
	confirmationData.KeyValueTableEntry = make([]KeyValueTableEntry, 0)

	confirmationData.KeyValueTableEntry = append(confirmationData.KeyValueTableEntry, KeyValueTableEntry{
		Key: "COURSE", Value: "Golang3",
	})
	confirmationData.KeyValueTableEntry = append(confirmationData.KeyValueTableEntry, KeyValueTableEntry{
		Key: "Description", Value: "Kids will learn computer science techniques on how to solve day to day problems using computer as a tool.3",
	})
	confirmationData.KeyValueTableEntry = append(confirmationData.KeyValueTableEntry, KeyValueTableEntry{
		Key: "Price", Value: "$85.99",
	})
	confirmationData.KeyValueTableEntryMap["row3"] = confirmationData.KeyValueTableEntry
	//Reset after setting data to map

	confirmation := confirmationData.CopyWith(confirmationData)

	configs := []confirmationConfig{
		&confirmation,
	}

	themes := []hermes.Theme{
		new(hermes.Default),
		//new(hermes.Flat),
	}

	// Generate emails
	for _, theme := range themes {
		h.Theme = theme
		for _, e := range configs {
			generateEmails(h, e.Email(), e.Name())
		}
	}

	// Send emails only when requested
	if sendEmails {
		port, _ := strconv.Atoi(os.Getenv("HERMES_SMTP_PORT"))
		password := os.Getenv("HERMES_SMTP_PASSWORD")
		SMTPUser := os.Getenv("HERMES_SMTP_USER")
		if password == "" {
			fmt.Printf("Enter SMTP password of '%s' account: ", SMTPUser)
			bytePassword, _ := terminal.ReadPassword(0)
			password = string(bytePassword)
		}
		smtpConfig := smtpAuthentication{
			Server:         os.Getenv("HERMES_SMTP_SERVER"),
			Port:           port,
			SenderEmail:    os.Getenv("HERMES_SENDER_EMAIL"),
			SenderIdentity: os.Getenv("HERMES_SENDER_IDENTITY"),
			SMTPPassword:   password,
			SMTPUser:       SMTPUser,
		}

		var recepient string

		if len(receiverEmail) == 0 {
			if len(os.Getenv("HERMES_TO")) == 0 {
				recepient = os.Getenv("HERMES_TO")

			}
		} else {
			recepient = receiverEmail
		}
		options := sendOptions{
			To: recepient,
		}
		for _, theme := range themes {
			h.Theme = theme
			for _, e := range configs {
				options.Subject = "App | " + h.Theme.Name() + " | " + e.Name()
				fmt.Printf("Sending email '%s'...\n", options.Subject)
				htmlBytes, err := ioutil.ReadFile(fmt.Sprintf("%v/%v.%v.html", h.Theme.Name(), h.Theme.Name(), e.Name()))
				if err != nil {
					panic(err)
				}
				txtBytes, err := ioutil.ReadFile(fmt.Sprintf("%v/%v.%v.txt", h.Theme.Name(), h.Theme.Name(), e.Name()))
				if err != nil {
					panic(err)
				}
				err = send(smtpConfig, options, string(htmlBytes), string(txtBytes))
				if err != nil {
					status = false
					panic(err)
				}

				status = true
			}
		}
	}

	return status
}

func SendEmail(mailParam MailParam) (status bool) {
	//"If you’re having trouble with the button '{ACTION}', copy and paste the URL below into your web browser."
	h := hermes.Hermes{
		Product: hermes.Product{
			Name:        mailParam.CompanyName,
			Link:        mailParam.ConfirmationLink,
			Logo:        mailParam.CompanyLogoLink,
			Copyright:   mailParam.Copyright,
			TroubleText: mailParam.TroubelText,
		},
	}

	sendEmails := os.Getenv("HERMES_SEND_EMAILS") == "true"
	var confirmationData confirmation
	confirmationData = GetMailerDate(mailParam)
	confirmationData.Intent = ConfirmationEmail
	confirmationData.Company = "Duxte Limited"
	confirmationData.Signature = "Paul Msegeya"
	confirmationData.IntrosMessages = []string{"Welcome to Kodeuni we present the courses below for your Kids Future", "Computer Science for Kids", "Codding for Kids", "Critical thinking Courses for Kids"}
	confirmationData.OutrosMessages = []string{"Prepare your child future today ", "Subscribe today and get 10% off ", "From Kodeuni the Platform for Kids Future today.", "You are the change your kid need"}
	confirmationData.ValiationLink = "https://www.duxte.co.tz"
	confirmationData.ConfirmationLink = "https://www.duxte.co.tz"
	confirmationData.DashboardButtonLink = "https://www.duxte.co.tz"
	confirmationData.CopyrightLink = ""
	confirmationData.DashboardButtonText = "OTP"
	confirmationData.DashboardButtonText = "SUBSCRIBE TODAY"
	confirmationData.Date = time.Now().Local().String()
	confirmationData.HasTableData = true
	confirmationData.MaxTableRows = 3
	confirmationData.DashboardButtonColor = "#40b892"
	confirmationData.MaxTableColumns = 0 //
	confirmationData.SupportEmail = "support@duxte.com"
	confirmationData.RecepientName = "paul@duxte.com"
	confirmationData.KeyValue = make(map[string]string)
	confirmationData.KeyValue["price"] = "$10.00"
	confirmationData.KeyValue["item"] = "Some Item"
	confirmationData.KeyValueTableEntry = make([]KeyValueTableEntry, 0)

	confirmationData.KeyValueTableEntry = append(confirmationData.KeyValueTableEntry, KeyValueTableEntry{
		Key: "COURSE", Value: "Computer sciente for KIDS",
	})
	confirmationData.KeyValueTableEntry = append(confirmationData.KeyValueTableEntry, KeyValueTableEntry{
		Key: "Description", Value: "Kids will learn computer science techniques on how to solve day to day problems using computer as a tool.",
	})
	confirmationData.KeyValueTableEntry = append(confirmationData.KeyValueTableEntry, KeyValueTableEntry{
		Key: "Price", Value: "$10.99",
	})

	log.Println(confirmationData.KeyValueTableEntry)

	confirmationData.KeyValueTableEntryMap = make(map[string][]KeyValueTableEntry)
	confirmationData.KeyValueTableEntryMap["row1"] = confirmationData.KeyValueTableEntry
	//Reset after setting data to map
	confirmationData.KeyValueTableEntry = make([]KeyValueTableEntry, 0)

	confirmationData.KeyValueTableEntry = append(confirmationData.KeyValueTableEntry, KeyValueTableEntry{
		Key: "COURSE", Value: "Codding for Kids",
	})
	confirmationData.KeyValueTableEntry = append(confirmationData.KeyValueTableEntry, KeyValueTableEntry{
		Key: "Description", Value: "Kids will learn computer science techniques on how to solve day to day problems using computer as a tool.2",
	})
	confirmationData.KeyValueTableEntry = append(confirmationData.KeyValueTableEntry, KeyValueTableEntry{
		Key: "Price", Value: "$25.99",
	})

	confirmationData.KeyValueTableEntryMap["row2"] = confirmationData.KeyValueTableEntry
	//Reset after setting data to map
	confirmationData.KeyValueTableEntry = make([]KeyValueTableEntry, 0)

	confirmationData.KeyValueTableEntry = append(confirmationData.KeyValueTableEntry, KeyValueTableEntry{
		Key: "COURSE", Value: "Golang3",
	})
	confirmationData.KeyValueTableEntry = append(confirmationData.KeyValueTableEntry, KeyValueTableEntry{
		Key: "Description", Value: "Kids will learn computer science techniques on how to solve day to day problems using computer as a tool.3",
	})
	confirmationData.KeyValueTableEntry = append(confirmationData.KeyValueTableEntry, KeyValueTableEntry{
		Key: "Price", Value: "$85.99",
	})
	confirmationData.KeyValueTableEntryMap["row3"] = confirmationData.KeyValueTableEntry
	//Reset after setting data to map

	confirmation := confirmationData.CopyWith(confirmationData)

	configs := []confirmationConfig{
		&confirmation,
	}

	themes := []hermes.Theme{
		new(hermes.Default),
		//new(hermes.Flat),
	}

	// Generate emails
	for _, theme := range themes {
		h.Theme = theme
		for _, e := range configs {
			generateEmails(h, e.Email(), e.Name())
		}
	}

	// Send emails only when requested
	if sendEmails {
		port, _ := strconv.Atoi(os.Getenv("HERMES_SMTP_PORT"))
		password := os.Getenv("HERMES_SMTP_PASSWORD")
		SMTPUser := os.Getenv("HERMES_SMTP_USER")
		if password == "" {
			fmt.Printf("Enter SMTP password of '%s' account: ", SMTPUser)
			bytePassword, _ := terminal.ReadPassword(0)
			password = string(bytePassword)
		}
		smtpConfig := smtpAuthentication{
			Server:         os.Getenv("HERMES_SMTP_SERVER"),
			Port:           port,
			SenderEmail:    os.Getenv("HERMES_SENDER_EMAIL"),
			SenderIdentity: os.Getenv("HERMES_SENDER_IDENTITY"),
			SMTPPassword:   password,
			SMTPUser:       SMTPUser,
		}

		var recepient string

		if len(mailParam.ReceiverEmail) == 0 {
			if len(os.Getenv("HERMES_TO")) == 0 {
				recepient = os.Getenv("HERMES_TO")

			}
		} else {
			recepient = mailParam.ReceiverEmail
		}
		options := sendOptions{
			To: recepient,
		}
		for _, theme := range themes {
			h.Theme = theme
			for _, e := range configs {
				options.Subject = "App | " + h.Theme.Name() + " | " + e.Name()
				fmt.Printf("Sending email '%s'...\n", options.Subject)
				htmlBytes, err := ioutil.ReadFile(fmt.Sprintf("%v/%v.%v.html", h.Theme.Name(), h.Theme.Name(), e.Name()))
				if err != nil {
					panic(err)
				}
				txtBytes, err := ioutil.ReadFile(fmt.Sprintf("%v/%v.%v.txt", h.Theme.Name(), h.Theme.Name(), e.Name()))
				if err != nil {
					panic(err)
				}
				err = send(smtpConfig, options, string(htmlBytes), string(txtBytes))
				if err != nil {
					status = false
					panic(err)
				}

				status = true
			}
		}
	}

	return status
}

func SendHermesEmail() (status bool) {

	h := hermes.Hermes{
		Product: hermes.Product{
			Name: "Duxte",
			Link: "https://www.duxte.co.tz",
			Logo: "https://github.com/pmtabe1/media/blob/main/duxte.png?raw=true",
		},
	}
	sendEmails := os.Getenv("HERMES_SEND_EMAILS") == "true"

	examples := []example{

		new(welcome),
		new(reset),
		new(receipt),
		new(maintenance),
		new(inviteCode),
	}

	themes := []hermes.Theme{
		new(hermes.Default),
		new(hermes.Flat),
	}

	// Generate emails
	for _, theme := range themes {
		h.Theme = theme
		for _, e := range examples {
			generateEmails(h, e.Email(), e.Name())
		}
	}

	// Send emails only when requested
	if sendEmails {
		port, _ := strconv.Atoi(os.Getenv("HERMES_SMTP_PORT"))
		password := os.Getenv("HERMES_SMTP_PASSWORD")
		SMTPUser := os.Getenv("HERMES_SMTP_USER")
		if password == "" {
			fmt.Printf("Enter SMTP password of '%s' account: ", SMTPUser)
			bytePassword, _ := terminal.ReadPassword(0)
			password = string(bytePassword)
		}
		smtpConfig := smtpAuthentication{
			Server:         os.Getenv("HERMES_SMTP_SERVER"),
			Port:           port,
			SenderEmail:    os.Getenv("HERMES_SENDER_EMAIL"),
			SenderIdentity: os.Getenv("HERMES_SENDER_IDENTITY"),
			SMTPPassword:   password,
			SMTPUser:       SMTPUser,
		}
		options := sendOptions{
			To: os.Getenv("HERMES_TO"),
		}
		for _, theme := range themes {
			h.Theme = theme
			for _, e := range examples {
				options.Subject = "App | " + h.Theme.Name() + " | " + e.Name()
				fmt.Printf("Sending email '%s'...\n", options.Subject)
				htmlBytes, err := ioutil.ReadFile(fmt.Sprintf("%v/%v.%v.html", h.Theme.Name(), h.Theme.Name(), e.Name()))
				if err != nil {
					panic(err)
				}
				txtBytes, err := ioutil.ReadFile(fmt.Sprintf("%v/%v.%v.txt", h.Theme.Name(), h.Theme.Name(), e.Name()))
				if err != nil {
					panic(err)
				}
				err = send(smtpConfig, options, string(htmlBytes), string(txtBytes))
				if err != nil {
					status = false
					panic(err)
				}

				status = true
			}
		}
	}

	return status
}

func generateEmails(h hermes.Hermes, email hermes.Email, example string) {
	// Generate the HTML template and save it
	res, err := h.GenerateHTML(email)
	if err != nil {
		panic(err)
	}
	err = os.MkdirAll(h.Theme.Name(), 0744)
	if err != nil {
		panic(err)
	}
	err = ioutil.WriteFile(fmt.Sprintf("%v/%v.%v.html", h.Theme.Name(), h.Theme.Name(), example), []byte(res), 0644)
	if err != nil {
		panic(err)
	}

	// Generate the plaintext template and save it
	res, err = h.GeneratePlainText(email)
	if err != nil {
		panic(err)
	}
	err = ioutil.WriteFile(fmt.Sprintf("%v/%v.%v.txt", h.Theme.Name(), h.Theme.Name(), example), []byte(res), 0644)
	if err != nil {
		panic(err)
	}
}

type smtpAuthentication struct {
	Server         string
	Port           int
	SenderEmail    string
	SenderIdentity string
	SMTPUser       string
	SMTPPassword   string
}

// sendOptions are options for sending an email
type sendOptions struct {
	To      string
	Subject string
}

// send sends the email
func send(smtpConfig smtpAuthentication, options sendOptions, htmlBody string, txtBody string) error {

	if smtpConfig.Server == "" {
		return errors.New("SMTP server config is empty")
	}
	if smtpConfig.Port == 0 {
		return errors.New("SMTP port config is empty")
	}

	if smtpConfig.SMTPUser == "" {
		return errors.New("SMTP user is empty")
	}

	if smtpConfig.SenderIdentity == "" {
		return errors.New("SMTP sender identity is empty")
	}

	if smtpConfig.SenderEmail == "" {
		return errors.New("SMTP sender email is empty")
	}

	if options.To == "" {
		return errors.New("no receiver emails configured")
	}

	from := mail.Address{
		Name:    smtpConfig.SenderIdentity,
		Address: smtpConfig.SenderEmail,
	}

	m := gomail.NewMessage()
	m.SetHeader("From", from.String())
	m.SetHeader("To", options.To)
	m.SetHeader("Subject", options.Subject)

	m.SetBody("text/plain", txtBody)
	m.AddAlternative("text/html", htmlBody)

	d := gomail.NewDialer(smtpConfig.Server, smtpConfig.Port, smtpConfig.SMTPUser, smtpConfig.SMTPPassword)

	return d.DialAndSend(m)
}
