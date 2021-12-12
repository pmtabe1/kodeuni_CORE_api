package email_service

import (
	"testing"

	"github.com/stretchr/testify/require"
)

// MUST ALLOW LESS SECURE APP FOR THE GMAIL ACCOUNT TO WORK
// https://support.google.com/mail/thread/15048240/keep-getting-username-and-password-not-accepted-when-trying-to-download-gmail-to-my-ms-outlook?hl=en
//https://myaccount.google.com/lesssecureapps?pli=1&rapt=AEjHL4M2pUuIuk3m5APubhPfcRUMd9Z_WtJ-nVh-5IdSa93uNlG0aeiT5jZIhF0Td292jio1HF4Q8HsDzYErk0u1F9NRlSYsOw

func TestStartSendingEmails(t *testing.T) {
	got := StartSendingEmails("", "", "WEBSOCKET API FOR USER IS down please start it remotely while you can", "OAUTH")
	require.Truef(t, got, "Expects TRUE results but got %v Instead ", got)
}

func TestSendEmailSMTP(t *testing.T) {
	ReceiverName := "Doe"
	data := map[string]interface{}{"ReceiverName": string(ReceiverName)}
	data["to"] = "paul@duxte.com"
	data["SenderName"] = "tiradevworks@gmail.com"
	data["message"] = "Sample message to paul"
	got, _ := SendEmailSMTP([]string{"paul@duxte.com"}, data, "email_template.txt")
	require.Truef(t, got, "Expects TRUE results but got %v Instead ", got)
}
