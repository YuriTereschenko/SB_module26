package main

import (
	"fmt"
	"log"
	"os"
	"strings"
)

func getFileData(fileName string) string {
	data, err := os.ReadFile(fileName)
	if err != nil {
		log.Fatalf("File %s opening error: %v", fileName, err)
	}
	return fmt.Sprintf("%s", data)
}

func main() {
	filesName := os.Args

	data := "You have not specified files to read"

	if len(filesName) > 1 {
		data = getFileData(filesName[1])
	}

	if len(filesName) > 2 {
		data = strings.Join([]string{data, getFileData(filesName[2])}, "\n")
	}

	if len(filesName) > 3 {
		err := os.WriteFile(filesName[3], []byte(data), 0777)
		if err != nil {
			log.Fatalf("Error writing to file: %v", err)
		}
	} else {
		fmt.Println(data)
	}

}
