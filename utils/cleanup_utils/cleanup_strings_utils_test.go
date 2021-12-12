package cleanup_utils

import (
	"testing"

	"github.com/stretchr/testify/require"
)


func TestCleanSpecifiedCharacter(t *testing.T)  {
	got:=CleanSpecifiedCharacter("+255782Z 531 360",[]string{" ","+","0X20"})
	require.Equalf(t,got,"255782531360","Should be equal but got %v insted ",got)
}