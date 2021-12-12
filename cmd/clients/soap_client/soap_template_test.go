package soap_client

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestCreateSoapConfigurationTemplate(t *testing.T) {
	got := CreateSoapConfigurationTemplate(getDefaultSoapTemplate)
	require.Truef(t, got, "Expected True result but got %v Instead ", got)
}
