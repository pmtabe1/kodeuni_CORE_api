package string_utils

import (
	"bytes"
	"fmt"
	"log"
	"reflect"
	"regexp"
	"strconv"
	"strings"
)

func NextStringFiedNumber(stringWithNumbers string) (outNext string) {

	templatedSlice := ExtractNumberFromStringTemplated(stringWithNumbers)

	stringFiedSlice := ExtractNumberFromString(stringWithNumbers)

	intFound, _ := strconv.Atoi(stringFiedSlice[0])

	outNext = fmt.Sprintf(templatedSlice[0], intFound+1)

	return outNext
}

func ExtractNumberFromStringTemplated(stringWithNumbers string) (data []string) {
	re := regexp.MustCompile("[0-9]+")
	//fmt.Println(re.FindAllString("abc123def987asdf", -1))
	outCOme := re.FindAllString(stringWithNumbers, -1)

	outCOmeSlice := make([]string, 0)
	// outCOmeSplited := strings.Split(outCOme[0], " ")

	if len(outCOme) > 1 {
		outCOmeSlice = append(outCOmeSlice, strings.ReplaceAll(outCOme[0], "3", outCOme[1]))

	} else {
		outCOmeSlice = append(outCOmeSlice, outCOme...)
	}

	for _, v := range outCOmeSlice {

		vv, _ := strconv.Atoi(v)
		if vv > 0 {
			data = append(data, strings.ReplaceAll(v, fmt.Sprintf("%v", vv), "%v"))
		}

	}

	return data
}

//https://stackoverflow.com/questions/37334119/how-to-delete-an-element-from-a-slice-in-golang
func ExtractNumberFromString(stringWithNumbers string) []string {
	re := regexp.MustCompile("[0-9]+")
	//fmt.Println(re.FindAllString("abc123def987asdf", -1))
	outCOme := re.FindAllString(stringWithNumbers, -1)

	// check if the outcome contains space and if so take the first number a

	// Test

	outCOmeSlice := make([]string, 0)

	if len(outCOme) > 1 {
		extractedInt,_:=strconv.Atoi(outCOme[0])
		outCOmeSlice = append(outCOmeSlice, strings.ReplaceAll(outCOme[0],strconv.Itoa(int(extractedInt)), outCOme[1]))
	} else {
		outCOmeSlice = append(outCOmeSlice, outCOme...)
	}

	return outCOmeSlice
}

func FloatToString(floatInput interface{}) (t string) {

	t = fmt.Sprintf("%f", reflect.TypeOf(floatInput)) // s == "123.456000"

	log.Println(t)
	return t
}

func removeOrdered(slice []int, s int) []int {
	return append(slice[:s], slice[s+1:]...)
}

func removeNoOrdered(s []int, i int) []int {
	s[i] = s[len(s)-1]
	return s[:len(s)-1]
}

func GetCombinedString(data []string, combinnerCharacter string) (combinedString string) {

	var buffer bytes.Buffer

	// for i := 0; i < 1000; i++ {
	//     buffer.WriteString("a")
	// }

	//fmt.Println(buffer.String())
	if len(data) > 0 {

		for _, v := range data {
			buffer.WriteString(fmt.Sprintf("%v %v", v, combinnerCharacter))

		}

	}

	combinedString = buffer.String()
	log.Println(combinedString)

	return combinedString
}

func GetSplittedStringSlice(combinedSring string, splitterCharacter string) (splittedStringSlice []string) {

	if len(combinedSring) == 0 {

		log.Panicln("The provided string to split is empty please provide one   ")
	}

	if len(splitterCharacter) == 0 {

		log.Panicln("The provided Splitter Character is empty please provide one ")
	}

	splittedStringSlice = strings.Split(combinedSring, splitterCharacter)
	splittedStringSlice = CleanEmptySpacesInStringSlice(splittedStringSlice)

	return splittedStringSlice
}

func CleanEmpySpacesInStrings(subjectString string) (cleanedString string) {

	if len(subjectString) == 0 {
		log.Panicln("The Subject string for space character removal is empty")
	}

	sliceWithEmpties := strings.Split(subjectString, "")
	cleanSlice := CleanEmptySpacesInStringSlice(sliceWithEmpties)

	if len(cleanSlice) > 0 {
		cleanedString = cleanSlice[0]
	} else {
		log.Panicln("The String provided does not have any empty space")
		cleanedString = subjectString
	}

	return cleanedString
}

// removeEmptyStrings - Use this to remove empty string values inside an array.
// This happens when allocation is bigger and empty
func CleanEmptySpacesInStringSlice(s []string) []string {
	var r []string
	for _, str := range s {
		if str != "" {
			r = append(r, str)
		}
	}
	return r
}
