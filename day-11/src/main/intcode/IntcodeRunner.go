package main

import (
	"../utils"
	"./model"
	"flag"
	"fmt"
	"os"
	"strings"
)

func main() {
	inputFilePath := flag.String("inputFilePath", "",
		"Path to file with input list of int in one line separate by comma, for example: input.txt")
	inputArray := flag.String("inputArray", "",
		"Input array of ints separate by comma, for example: 1,2,3,4")
	flag.Parse()
	if *inputArray != "" && *inputFilePath != "" {
		fmt.Fprintln(os.Stderr, "Only one parameter should be pass")
		os.Exit(1)
	}
	var input []string
	var err error
	if *inputArray != "" {
		input = strings.Split(*inputArray, ",")
	}
	if *inputFilePath != "" {
		input, err = utils.ReadFirstLineInFile(*inputFilePath)
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(2)
		}
	}
	inputAsIntArray := utils.StringArrayToIntArray(input)
	computer := model.New(inputAsIntArray, nil, nil)
	output, err := computer.Process()
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(3)
	}
	fmt.Println(utils.IntArrayToString(output))
}
