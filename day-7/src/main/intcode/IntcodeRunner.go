package main

import (
	"bufio"
	"errors"
	"flag"
	"fmt"
	"os"
	"strconv"
	"strings"
	"./model"
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
		input, err = readInput(*inputFilePath)
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(2)
		}
	}
	inputAsIntArray := stringArrayToIntArray(input)
	computer := model.New(inputAsIntArray)
	output, err := computer.Process()
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(3)
	}
	fmt.Println(intArrayToString(output))
}

func readInput(inputFilePath string) ([]string, error) {
	file, err := os.Open(inputFilePath)
	if err != nil {
		return nil, errors.New("Can not read file " + inputFilePath)
	}
	scanner := bufio.NewScanner(file)
	scanner.Scan()
	input := scanner.Text()
	return strings.Split(input, ","), nil
}

func stringArrayToIntArray(stringArray []string) []int {
	intArray := []int{}
	for _, str := range stringArray {
		value, _ := strconv.Atoi(str)
		intArray = append(intArray, value)
	}
	return intArray
}

func intArrayToString(array []int) string {
	return strings.Trim(strings.Join(strings.Fields(fmt.Sprint(array)), ","), "[]")
}
