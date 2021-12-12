package utils

import (

	"log"
	"os"
)

func FileWriter(file *os.File, fileContents string) (status bool) {

	defer file.Close()

	_, err2 := file.WriteString(fileContents)

	if err2 != nil {
		status = false
		log.Fatal(err2)
	}

	status = true

	log.Println("done")

	return status
}

func ReadFileAll(file *os.File) (fileContents string) {

	defer file.Close()


	return fileContents
}
