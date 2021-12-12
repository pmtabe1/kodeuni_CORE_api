package sage300_utils

import (
	"fmt"
	"log"
	"strings"
	"time"
)

func GetTodaysDateWithSag300FormatDatetime()  time.Time {
	sageDateTime, err := time.Parse("2006-01-02T00:00:00Z", GetTodaysDateWithSage300FormatString())


	if err!=nil {
		log.Panicln("ERROR WHILE PARSING SAGE300 String date to datetime aka time.Time")
	}


	return sageDateTime
}

func GetTodaysDateWithSage300FormatString() (date string) {
	//2021-11-08T00:00:00Z
    //2021-11-25T16:25:13Z
	//2021-11-25T16:03:04+03:00
	d := GetTodaysDateWithISOFormatString()
	dSlicedWithPlus := strings.Split(d, "+")

	if len(dSlicedWithPlus) == 0 {
		log.Panicln("ERROR the splitted string is EMPTY and it must never be EMPTY FIX your code , this is a very bad code smell")
	}

	//YOU ARE GOOD HERE

	date = dSlicedWithPlus[0] + "Z"

	return date

}

func GetTodaysDateWithISOFormatString() string {

	return fmt.Sprintf("%v", time.Now().Format(time.RFC3339))
}
