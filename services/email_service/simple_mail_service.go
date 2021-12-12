package email_service

import (
	"log"
	"net/smtp"
	//"reflect"
)

// func SendMail() {
// 	send("hello there")
// }

func Send(body string) (status bool){
	from := "tiradevwork@gmail.com"
	pass := "s@mb@mb@000"
	to := "paul@duxte.com"

	msg := "From: " + from + "\n" +
		"To: " + to + "\n" +
		"Subject: Hello there\n\n" +
		body

	err := smtp.SendMail("smtp.gmail.com:587",
		smtp.PlainAuth("", from, pass, "smtp.gmail.com"),
		from, []string{to}, []byte(msg))

	if err != nil {
		log.Printf("smtp error: %s", err)
		status=false
		return status
	}

	status=true
	
	log.Print("sent, visit http://paul.duxte.com")

	return status
}