package cleanup_utils

import (
	"fmt"
	"log"
	"regexp"
	"strings"

	"github.com/icza/gox/stringsx"
)

//

func RemoveAllSymbolWithRegex(inputData string) (data string) {
	//str1 := "how much for the maple syrup? $20.99? That's ridiculous!!!"

	re, err := regexp.Compile(`[^\w]`)
	if err != nil {
		log.Panicln(err)
	}
	data = re.ReplaceAllString(inputData, " ")
	data = stringsx.Clean(data)

	fmt.Println(data)

	return data
}

func CleanSpecifiedCharacter(dataToClean string, targetCharacters []string) (cleanedString string) {

	if len(targetCharacters) == 0 {
		targetCharacters = make([]string, 0)
		targetCharacter := " "
		targetCharacters = append(targetCharacters, targetCharacter)
	}

	// all is good get to work on cleaning

	for _, targetCharacter := range targetCharacters {

		cleanedString = RemoveAllSymbolWithRegex(dataToClean)
		cleanedString = stringsx.Clean(cleanedString)
		cleanedString = strings.ReplaceAll(cleanedString, targetCharacter, "")

	}

	return cleanedString
}

func CleanSpecifiedCharacterByReplacingWithAnother(dataToClean string, targetCharacters []string, replacerCharacter string) (cleanedString string) {

	if len(targetCharacters) == 0 {
		targetCharacters = make([]string, 0)
		targetCharacter := " "
		targetCharacters = append(targetCharacters, targetCharacter)
	}

	// all is good get to work on cleaning

	for _, targetCharacter := range targetCharacters {
		cleanedString = RemoveAllSymbolWithRegex(dataToClean)
		cleanedString = strings.ReplaceAll(cleanedString, targetCharacter, replacerCharacter)

	}

	return cleanedString
}
