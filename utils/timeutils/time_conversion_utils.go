package timeutils

import (
	"fmt"
	"log"
	"strings"
	"time"
)

func GetTokenRemainingValidity(timestamp interface{}) int {
	var expireOffset = 3600
	if validity, ok := timestamp.(int64); ok {
		tm := time.Unix(int64(validity), 0)
		remainder := tm.Sub(time.Now())

		if remainder > 0 {
			return int(remainder.Seconds() + float64(expireOffset))
		}
	}
	return expireOffset
}

func GetTRATokenRemainingValidity(timestamp interface{}, tin string) int {
	var expireOffset = 76399 //time.Now().Add(time.Hour * 24).Unix()

	//get expireOffset aka ExpiresIn from registration table

	//registationRepository := registration_repository.New()
	//registration := registationRepository.GetByTin(tin).Registration
	//expireOffset = registration.ExpiresIn
	var remainder time.Duration

	if validity, ok := timestamp.(int64); ok {
		tm := time.Unix(int64(validity), 0)
		remainder = tm.Sub(time.Now())

		if remainder > 0 {
			log.Printf("Remaining Valid Time : %v ", remainder)
			return int(remainder.Seconds() + float64(expireOffset))
		} else {
			log.Printf("Remaining Valid Time : %v ", remainder)

		}
	}
	log.Printf("Remaining Valid Time : %v ", remainder.Seconds())
	log.Printf("TRA Token ExpireIn offset %v", expireOffset)
	return int(expireOffset)
}

func GetJWTTokenExpireOffsetFromClaims() int {
	var claims map[string]int64
	claims = make(map[string]int64)
	claims["exp"] = time.Now().Add(time.Hour * time.Duration(1)).Unix()

	x := GetTokenRemainingValidity(claims["exp"])

	fmt.Println(x)

	return x
}

func FromJavaLongMilliseconds(milliseconds int64) time.Time {
	return time.Unix(0, milliseconds*int64(time.Millisecond))
}

func FromJavaLongMillisecondsGetDate(milliseconds int64) (data string) {
	longTime := time.Unix(0, milliseconds*int64(time.Millisecond))

	slicedWithSpace := strings.Split(longTime.Format(""), " ")

	if len(slicedWithSpace) > 0 {
		data = slicedWithSpace[0]

		return data

	} else {

		return "Error while converting Time"
	}
}

func FromJavaLongMillisecondsGetTime(milliseconds int64) (data string) {
	// longTime := time.Unix(0, milliseconds*int64(time.Millisecond))
	// slicedWithSpace := strings.Split(longTime.Format("","-"))

	// if len(slicedWithSpace) > 0 {
	// 	data = slicedWithSpace[1]

	// 	return data

	// } else {

	// 	return "Error while converting Time"
	// }

	return data
}

func GetFineractDate() string {

	var fineractDate string

	m := strings.Split(fmt.Sprintf("%v", time.Now().Format(time.RFC3339)), "T")
	log.Println(m)
	mm := strings.Split(m[0], "-")
	ddMMMYYYY:=`%v %v %v`

	if len(mm) > 0 {

		log.Println(";;;;;")

		switch mm[1] {

		case "1":
			//fineractDate = mm[2] + " January " + mm[0]
 			fineractDate =fmt.Sprintf(ddMMMYYYY,mm[2],"January",mm[0])

		case "2":
			//fineractDate = mm[2] + " February " + mm[0]
			fineractDate =fmt.Sprintf(ddMMMYYYY,mm[2],"February",mm[0])

		case "3":
			//fineractDate = mm[2] + " March  " + mm[0]
			fineractDate =fmt.Sprintf(ddMMMYYYY,mm[2],"March",mm[0])

		case "4":
			//fineractDate = mm[2] + " April  " + mm[0]
			fineractDate =fmt.Sprintf(ddMMMYYYY,mm[2],"April",mm[0])

		case "5":
			//fineractDate = mm[2] + " May  " + mm[0]
			fineractDate =fmt.Sprintf(ddMMMYYYY,mm[2],"May",mm[0])

		case "6":
			//fineractDate = mm[2] + " June  " + mm[0]
			fineractDate =fmt.Sprintf(ddMMMYYYY,mm[2],"June",mm[0])

		case "7":
			//fineractDate = mm[2] + " July  " + mm[0]
			fineractDate =fmt.Sprintf(ddMMMYYYY,mm[2],"July",mm[0])

		case "8":
			//fineractDate = mm[2] + " August  " + mm[0]
			fineractDate =fmt.Sprintf(ddMMMYYYY,mm[2],"August",mm[0])

		case "9":
			//fineractDate = mm[2] + " September  " + mm[0]
			fineractDate =fmt.Sprintf(ddMMMYYYY,mm[2],"September",mm[0])

		case "10":
			//fineractDate = mm[2] + " October  " + mm[0]
			fineractDate =fmt.Sprintf(ddMMMYYYY,mm[2],"October",mm[0])

		case "11":
			//fineractDate = mm[2] + " November  " + mm[0]
			fineractDate =fmt.Sprintf(ddMMMYYYY,mm[2],"November",mm[0])

		case "12":
			//fineractDate = mm[2] + " December  " + mm[0]
			fineractDate =fmt.Sprintf(ddMMMYYYY,mm[2],"December",mm[0])
		}

	}

	return fineractDate
}
