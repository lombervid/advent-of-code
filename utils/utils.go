package utils

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
)

const separator = string(os.PathSeparator)

func getRootDir(path string) string {
	regex := fmt.Sprintf(`\%v[\d]{4}\%v?(?:[\d]{2}\%v?)?$`, separator, separator, separator) // Remove "/{year}/{day}/" from path
	re, err := regexp.Compile(regex)

	if err != nil {
		log.Fatal(err)
	}

	return re.ReplaceAllString(path, "")
}

func getCurrentDir() string {
	path, err := os.Getwd()
	if err != nil {
		log.Println(err)
	}

	return path
}

func CheckErrNil(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func ReadFile(filepath string) []string {
	file, err := os.Open(filepath)

	if err != nil {
		log.Fatal(err)
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)
	var text []string

	for scanner.Scan() {
		text = append(text, scanner.Text())
	}

	return text
}

func GetInputFilePath(year string, day string) string {
	rootDir := getRootDir(getCurrentDir())

	return rootDir + fmt.Sprintf("%v%v%vinputs%v%v.txt", separator, year, separator, separator, day)
}
