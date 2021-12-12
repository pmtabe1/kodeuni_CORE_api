package conversion_utils

import (
	"testing"

	"github.com/stretchr/testify/require"
)


func TestGetDecimalPositionFromFloat64(t *testing.T)  {
	got:=GetDecimalPositionFromFloat64(9.67979,2)
	require.Nilf(t,got,"Expected non nil but got %v",got)
}


func TestConvertFloat64ToString(t *testing.T)  {
	
	got:=ConvertFloat64ToString(9.679797878233)
	require.NotEmptyf(t,got,"Expected non nil but got %v",got)
}