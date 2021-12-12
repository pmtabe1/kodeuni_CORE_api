package reference_service

import (
	"os"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestNew(t *testing.T) {
	got := New()
	require.NotNilf(t, got, "Expected  non nil BUT got %v instead ", got)
}

func TestGenerateReferenceNumber(t *testing.T) {

	os.Setenv("REFERENCE_SIZE","100000000")

	got := New().GenerateReferenceNumber()
	require.Nilf(t, got, "Expected  non nil BUT got %v instead ", got)

}
