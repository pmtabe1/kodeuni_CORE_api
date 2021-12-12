package email_service

import (
	"testing"

	"github.com/stretchr/testify/require"
)



func TestSend(t *testing.T) {
	got := Send("Simple mail....")
	require.Truef(t, got, "Expects TRUE results but got %v Instead ", got)
}