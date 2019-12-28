package main

import (
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"strconv"
	"strings"

	"./intcode"
)

func fileContentToIntArray(fileContent string) []int {
	code := []int{}
	for _, digit := range strings.Split(strings.Trim(fileContent, "\n"), ",") {
		val, err := strconv.Atoi(digit)
		if err != nil {
			log.Fatal(err)
		}
		code = append(code, val)
	}
	return code
}

func partOne(code []int) {

}

func partTwo(code []int) {

}

func main() {
	var filePath string
	if len(os.Args) != 2 {
		filePath = "day-21/src/resources/input.txt"
	} else {
		filePath = os.Args[1]
	}
	path, _ := filepath.Abs(filePath)
	fileContent, err := ioutil.ReadFile(path)
	if err != nil {
		log.Fatal(err)
	}
	originalCode := fileContentToIntArray(string(fileContent))
	partOne(originalCode)
	partTwo(originalCode)
}
