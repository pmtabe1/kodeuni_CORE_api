package email_service

import (
	"os"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestSendHermesEmail(t *testing.T) {

	os.Setenv("HERMES_SMTP_PORT", "587")
	os.Setenv("HERMES_SMTP_PASSWORD", "s@mb@mb@000")
	os.Setenv("HERMES_SMTP_USER", "tiradevwork@gmail.com")
	os.Setenv("HERMES_TO", "paul@duxte.com")
	os.Setenv("HERMES_SEND_EMAILS", "true")
	os.Setenv("HERMES_SMTP_SERVER", "smtp.gmail.com")
	os.Setenv("HERMES_SENDER_IDENTITY", "Paul Mtabe")
	os.Setenv("HERMES_SENDER_EMAIL", "paul.msegeya.db2@gmail.com") //tiradevwork@gmail.com
	//os.Setenv("HERMES_SENDER_EMAIL", "tiradevwork@gmail.com")
	got := SendHermesEmail()
	require.Truef(t, got, "Expects TRUE results but got %v Instead ", got)

}

func TestSendSignupVerificationEmail(t *testing.T) {

	os.Setenv("HERMES_SMTP_PORT", "587")
	os.Setenv("HERMES_SMTP_PASSWORD", "s@mb@mb@000")
	os.Setenv("HERMES_SMTP_USER", "tiradevwork@gmail.com")
	os.Setenv("HERMES_TO", "paul@duxte.com")
	os.Setenv("HERMES_SEND_EMAILS", "true")
	os.Setenv("HERMES_SMTP_SERVER", "smtp.gmail.com")
	os.Setenv("HERMES_SENDER_IDENTITY", "Paul Mtabe")
	os.Setenv("HERMES_SENDER_EMAIL", "paul.msegeya.db2@gmail.com") //tiradevwork@gmail.com
	//os.Setenv("HERMES_SENDER_EMAIL", "tiradevwork@gmail.com")
	got := SendSignupVerificationEmail("paul@duxte.com")
	//got := SendSignupVerificationEmail("undule@duxte.com")

	//got := SendSignupVerificationEmail("")

	require.Truef(t, got, "Expects TRUE results but got %v Instead ", got)

}

func TestBuildKeyValueEntryWithReset(t *testing.T) {

	KeyValueTableEntryList := make([]KeyValueTableEntry, 0)
	KeyValueTableEntryList = append(KeyValueTableEntryList, KeyValueTableEntry{
		Key: "COURSE", Value: "Codding for Kids",
	})
	KeyValueTableEntryList = append(KeyValueTableEntryList, KeyValueTableEntry{
		Key: "Description", Value: "Kids will learn computer science techniques on how to solve day to day problems using computer as a tool.2",
	})
	KeyValueTableEntryList = append(KeyValueTableEntryList, KeyValueTableEntry{
		Key: "Price", Value: "$25.99",
	})
	got := BuildKeyValueEntryWithReset("row1", MailParam{}, KeyValueTableEntryList)
	require.NotEmptyf(t, got, "Expects NON EMPTY results but got %v Instead ", got)

}
