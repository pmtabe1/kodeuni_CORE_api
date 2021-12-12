package sage_constants

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestConvertOrderDetailInsertParamsToMSSQLParams(t *testing.T) {

	got := ConvertOrderDetailsInsertParamsToMSSQLParams()

	require.NotEmptyf(t, got, "Expected non empty but received %v instead")
}

func TestConvertOrderHeadersInsertParamsToMSSQLParams(t *testing.T) {

	got := ConvertOrderHeadersInsertParamsToMSSQLParams()

	require.Emptyf(t, got, "Expected non empty but received %v instead")
}

func TestConvertOrderHeadersInsertParamsToMSSQLParamsWithCast(t *testing.T) {

	got := ConvertOrderHeadersInsertParamsToMSSQLParamsWithCast()

	require.NotEmptyf(t, got, "Expected non empty but received %v instead")
}

func TestConvertOrderDetailInsertParamsToMSSQLParamsWithCast(t *testing.T) {

	got := ConvertOrderDetailsInsertParamsToMSSQLParamsWithCast()

	require.Emptyf(t, got, "Expected non empty but received %v instead")
}
