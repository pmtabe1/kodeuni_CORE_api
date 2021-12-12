package soap_client

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestNewSoapRequest(t *testing.T)  {
	
	req := NewSoapRequest()
	require.NotNilf(t, req, "Expecting a non nil result but got %v", req)
}
func TestPopulateRequest(t *testing.T) {
	req := NewSoapRequest()
	require.NotNilf(t, req, "Expecting a non nil result but got %v", req)
	got := req.PopulateRequest(req)
	require.NotNilf(t, got, "Expecting a non nil result but got %v", got)

}


func TestSoapCall(t *testing.T)  {
	
}


func TestGenerateSOAPRequest(t *testing.T)  {
	
}
