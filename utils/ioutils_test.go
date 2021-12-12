package utils

import (
	"os"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestFileWriter(t *testing.T) {
	file,err:=os.Create("test_file.json")
	require.NoError(t,err,"Expected error but received non error results")
	got:=FileWriter(file,`{"blabla":"blablabala...."}`)
	require.Truef(t,got,"Expected TRUe but received %v insted , TEST Failed",got)

}