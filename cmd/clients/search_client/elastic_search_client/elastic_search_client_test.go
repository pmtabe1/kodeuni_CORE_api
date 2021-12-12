package elastic_search_client

import (
	"testing"

	"github.com/stretchr/testify/require"
)



func TestNew(t *testing.T)  {
	got:=New("","","","")
	require.NotNilf(t,got,"Expected non Nil received %v instead",got)
}