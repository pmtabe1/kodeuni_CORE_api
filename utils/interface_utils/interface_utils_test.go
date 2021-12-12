package interface_utils

import (
	"testing"

	"github.com/stretchr/testify/require"
)



func TestGetReturnedDataFromInterface(t *testing.T)  {
	data:="Hello Spring Object"
	got:=GetReturnedDataFromInterface(data)
	require.NotEmptyf(t,got.(string),"Interface extraction for String should never be empty , but we got %v insted ",got.(string))
	dataInt:=700
	got=GetReturnedDataFromInterface(dataInt)
	require.NotEmptyf(t,got.(int),"Interface extraction for String should never be empty , but we got %v insted ",got.(int))
	dataFloat:=float32(67.00111)
	got=GetReturnedDataFromInterface(dataFloat)
	require.NotEmptyf(t,got.(float32),"Interface extraction for String should never be empty , but we got %v insted ",got.(float32))


	
}