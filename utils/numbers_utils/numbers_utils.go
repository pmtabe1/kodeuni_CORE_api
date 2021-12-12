package numbers_utils

import (
	"fmt"
	"strconv"
)


func StringToFloat64(number string)  (f float64)  {
	if s, err := strconv.ParseFloat(number, 64); err == nil {
		fmt.Println(s) // 3.14159265
		f=s
	}

	return f
}


func StringToFloat32(number string)  (f float64)  {
	if s, err := strconv.ParseFloat(number, 32); err == nil {
		fmt.Println(s) // 3.14159265
		f=s
	}

	return f
}