package sage300_utils

import (
	"testing"

	"github.com/stretchr/testify/require"
)


func TestGetTodaysDateWithISOFormatString(t *testing.T){
	got:=GetTodaysDateWithISOFormatString()
	require.NotEmptyf(t,got, "Expected non empty but got %v  INSTEAD",got)
}


func TestGetTodaysDateWithSag300FormatDatetime(t *testing.T)  {
	got:=GetTodaysDateWithSag300FormatDatetime()
	require.NotNilf(t,got, "Expected non nil but got %v  INSTEAD",got)

}

func TestGetTodaysDateWithSag300FormatString(t *testing.T)  {
	got:=GetTodaysDateWithSage300FormatString()
	require.Emptyf(t,got, "Expected non empty but got %v  INSTEAD",got)
}