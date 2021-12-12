package utils

import (
	//"fmt"
	"runtime"
)


type Utils struct {
	
}

func (u *Utils) CheckOSType() (m string) {

	if runtime.GOOS == "windows" {
		//fmt.Println("Hello from Windows")

		m = "windows"
	} else {
		//fmt.Println("Hello from Mac")

		m = "unix"
	}
	return m
}

func (u *Utils)  OSIsWindows() (status bool) {
	status = false
	status = u.CheckOSType() == "windows"

	if status {
		status = true
		//return true
	} else {
		status = false
		//return false
	}

	return status
}
