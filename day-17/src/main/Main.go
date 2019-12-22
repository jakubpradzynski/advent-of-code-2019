package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"strconv"
	"strings"

	"./intcode"
)

const (
	HASH_NUMBER     int = 35
	DOT_NUMBER      int = 46
	NEW_LINE_NUMBER int = 10
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

func outputToGrid(output []int) [][]string {
	var grid [][]string
	row := 0
	grid = append(grid, []string{})
	for i := 0; i < len(output); i++ {
		switch output[i] {
		case HASH_NUMBER:
			grid[row] = append(grid[row], "#")
		case DOT_NUMBER:
			grid[row] = append(grid[row], ".")
		case NEW_LINE_NUMBER:
			grid = append(grid, []string{})
			row++
		default:
			log.Printf("Unknown output value! - %d\n", output[i])
			grid[row] = append(grid[row], "^")
		}
	}
	return grid[:len(grid)-2]
}

func printGrid(grid [][]string) {
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[i]); j++ {
			fmt.Print(grid[i][j])
		}
		fmt.Println()
	}
}

func getGridWithIntersections(originalGrid [][]string) [][]string {
	grid := make([][]string, len(originalGrid))
	copy(grid, originalGrid)
	for i := 1; i < len(grid)-1; i++ {
		for j := 1; j < len(grid[i])-1; j++ {
			if grid[i][j] == "#" && grid[i+1][j] == "#" && grid[i-1][j] == "#" && grid[i][j+1] == "#" && grid[i][j-1] == "#" {
				grid[i][j] = "O"
			}
		}
	}
	return grid
}

func sumAlignments(gridWithIntersections [][]string) int {
	grid := make([][]string, len(gridWithIntersections))
	copy(grid, gridWithIntersections)
	sum := 0
	for i := 1; i < len(grid)-1; i++ {
		for j := 1; j < len(grid[i])-1; j++ {
			if grid[i][j] == "O" {
				sum += i * j
			}
		}
	}
	return sum
}

func toASCII(input string) []int {
	result := []int{}
	split := strings.Split(input, "")
	for _, r := range split {
		result = append(result, int(r[0]))
	}
	return append(result, 10)
}

func main() {
	var filePath string
	if len(os.Args) != 2 {
		filePath = "day-17/src/resources/input.txt"
	} else {
		filePath = os.Args[1]
	}
	path, _ := filepath.Abs(filePath)
	fileContent, err := ioutil.ReadFile(path)
	if err != nil {
		log.Fatal(err)
	}
	originalCode := fileContentToIntArray(string(fileContent))
	partOneCode := make([]int, len(originalCode))
	copy(partOneCode, originalCode)
	computer := intcode.NewComputer(partOneCode)
	endStatus, output, err := computer.Process(nil)
	if err != nil {
		log.Fatal(nil)
	}
	if endStatus != intcode.END_OF_CODE {
		log.Fatalf("Wrong exit status! - %d\n", endStatus)
	}
	grid := outputToGrid(output)
	printGrid(grid)
	gridWithIntersections := getGridWithIntersections(grid)
	fmt.Println()
	printGrid(gridWithIntersections)
	sumOfAlignments := sumAlignments(gridWithIntersections)
	log.Printf("Part One - %d\n", sumOfAlignments)

	partTwoCode := make([]int, len(originalCode))
	copy(partTwoCode, originalCode)
	partTwoCode[0] = 2
	computerTwo := intcode.NewComputer(partTwoCode)
	computerTwo.Process(nil)
	mainRoutineInput, functionAInput, functionBInput, functionCInput := toASCII(MAIN_ROUTINE), toASCII(FUNCTION_A), toASCII(FUNCTION_B), toASCII(FUNCTION_C)
	for _, i := range mainRoutineInput {
		computerTwo.Process(&i)
	}
	for _, i := range functionAInput {
		computerTwo.Process(&i)
	}
	for _, i := range functionBInput {
		computerTwo.Process(&i)
	}
	for _, i := range functionCInput {
		computerTwo.Process(&i)
	}
	continuousVideoFeed := int('n')
	newLine := 10
	computerTwo.Process(&continuousVideoFeed)
	_, output, _ = computerTwo.Process(&newLine)
	log.Printf("Part Two - %d\n", output[len(output)-1])
}

const (
	MAIN_ROUTINE = "A,B,B,C,C,A,B,B,C,A"
	FUNCTION_A   = "R,4,R,12,R,10,L,12"
	FUNCTION_B   = "L,12,R,4,R,12"
	FUNCTION_C   = "L,12,L,8,R,10"
)

// R4 R12 R10 L12 L12 R4 R12 L12 R4 R12 L12 L8 R10 L12 L8 R10 R4 R12 R10 L12 L12 R4 R12 L12 R4 R12 L12 L8 R10 R4 R12 R10 L12
// A B B C C A B B C A
// A - R4 R12 R10 L12
// B - L12 R4 R12
// C - L12 L8 R10
