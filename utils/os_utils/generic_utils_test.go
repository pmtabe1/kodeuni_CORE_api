package utils

import "testing"

func TestCheckOSType(t *testing.T) {

	u:=&Utils{}
	got := u.CheckOSType()

	if got == "" {
		t.Errorf("Expected Non Empty %s  but found  Empty String", got)
	}
}

func TestOSIsWindows(t *testing.T) {
	u:=&Utils{}
	got := u.OSIsWindows()

	if got  {
		t.Errorf("Expected boolean  true  but found %v instead ", got)

	}
}
