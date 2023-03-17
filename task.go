package main

import (
	"flag"
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
	var fFileName, sFileName, resFileName string
	flag.StringVar(&fFileName, "fFileName", "", "Set first file name")
	flag.StringVar(&sFileName, "sFileName", "", "Set second file name")
	flag.StringVar(&resFileName, "resFileName", "", "Set result file name")
	flag.Parse()

	data := "You have not specified files to read"

	if fFileName != "" {
		data = getFileData(fFileName)
	}

	if sFileName != "" {
		data = strings.Join([]string{data, getFileData(sFileName)}, "\n")
	}

	if resFileName != "" {
		err := os.WriteFile(resFileName, []byte(data), 0777)
		if err != nil {
			log.Fatalf("Error writing to file: %v", err)
		}
	} else {
		fmt.Println(data)
	}

}
