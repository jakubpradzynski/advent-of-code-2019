package main

import (
	"bufio"
	"errors"
	"flag"
	"fmt"
	"os"
	"strconv"
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
		input, err = readInput(*inputFilePath)
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(2)
		}
	}
	inputAsIntArray := stringArrayToIntArray(input)
	output, err := intcode(inputAsIntArray)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(3)
	}
	outputAsString := intArrayToString(output)
	fmt.Println(outputAsString)
	os.Exit(0)
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

func intcode(input []int) ([]int, error) {
	index := 0
	for index < len(input) {
		opcode := fullLengthOpcode(input[index])
		instructionType, firstParamMode, secondParamMode, _ := interpretFullLengthOpcode(opcode)
		switch instructionType {
		case 99:
			fmt.Println()
			index = len(input)
		case 1:
			var firstSummand int
			var secondSummand int
			if firstParamMode == 1 {
				firstSummand = input[index+1]
			} else {
				firstSummand = input[input[index+1]]
			}
			if secondParamMode == 1 {
				secondSummand = input[index+2]
			} else {
				secondSummand = input[input[index+2]]
			}
			sum := firstSummand + secondSummand
			input[input[index+3]] = sum
			index += 4
		case 2:
			var firstFactor int
			var secondFactor int
			if firstParamMode == 1 {
				firstFactor = input[index+1]
			} else {
				firstFactor = input[input[index+1]]
			}
			if secondParamMode == 1 {
				secondFactor = input[index+2]
			} else {
				secondFactor = input[input[index+2]]
			}
			product := firstFactor * secondFactor
			input[input[index+3]] = product
			index += 4
		case 3:
			reader := bufio.NewReader(os.Stdin)
			read, err := reader.ReadString('\n')
			if err != nil {
				return []int{}, err
			}
			value, err := strconv.Atoi(strings.TrimSuffix(read, "\n"))
			if err != nil {
				return []int{}, err
			}
			input[input[index+1]] = value
			index += 2
		case 4:
			var value int
			if firstParamMode == 1 {
				value = input[index+1]
			} else {
				value = input[input[index+1]]
			}
			fmt.Printf("%d ", value)
			index += 2
		default:
			fmt.Fprintln(os.Stderr, "Index = "+strconv.Itoa(index))
			fmt.Fprintln(os.Stderr, "Input state - "+intArrayToString(input))
			return []int{}, errors.New("Unknown opcode " + opcode)
		}
	}
	return input, nil
}

func fullLengthOpcode(opcode int) string {
	opcodeAsString := strconv.Itoa(opcode)
	fullLengthOpcode := "00000"
	switch len(opcodeAsString) {
	case 1:
		fullLengthOpcode = "0000" + opcodeAsString
	case 2:
		fullLengthOpcode = "000" + opcodeAsString
	case 3:
		fullLengthOpcode = "00" + opcodeAsString
	case 4:
		fullLengthOpcode = "0" + opcodeAsString
	case 5:
		fullLengthOpcode = opcodeAsString
	}
	return fullLengthOpcode
}

func interpretFullLengthOpcode(opcode string) (int, int, int, int) {
	instructionType, _ := strconv.Atoi(string(opcode[3:]))
	firstParamMode, _ := strconv.Atoi(string(opcode[2]))
	secondParamMode, _ := strconv.Atoi(string(opcode[1]))
	thirdParamMode, _ := strconv.Atoi(string(opcode[0]))
	return instructionType, firstParamMode, secondParamMode, thirdParamMode
}

func intArrayToString(array []int) string {
	return strings.Trim(strings.Join(strings.Fields(fmt.Sprint(array)), ","), "[]")
}
