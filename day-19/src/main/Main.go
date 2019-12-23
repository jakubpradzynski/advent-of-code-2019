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
	const size int = 50
	var grid [size][size]string
	for i := 0; i < size; i++ {
		for j := 0; j < size; j++ {
			computer := intcode.NewComputer(code)
			computer.Process(nil)
			computer.Process(&i)
			_, output, _ := computer.Process(&j)
			if output[0] == 0 {
				grid[i][j] = "."
			} else if output[0] == 1 {
				grid[i][j] = "#"
			}
		}
	}
	affectedPoints := 0
	for i := 0; i < size; i++ {
		for j := 0; j < size; j++ {
			if grid[i][j] == "#" {
				affectedPoints++
			}
		}
	}
	log.Printf("Part One - %d\n", affectedPoints)
}

const size int = 2000

func getClosestPoint(grid [size][size]string) (int, int) {
	for i := 0; i < size; i++ {
		for j := 0; j < size; j++ {
			if grid[i][j] == "#" {
				if grid[i+99][j] == "#" && grid[i][j+99] == "#" && grid[i+99][j+99] == "#" {
					return i, j
				}
			}
		}
	}
	return 0, 0
}

func partTwo(code []int) {
	var grid [size][size]string
	for i := 0; i < size; i++ {
		for j := 0; j < size; j++ {
			computer := intcode.NewComputer(code)
			computer.Process(nil)
			computer.Process(&i)
			_, output, _ := computer.Process(&j)
			if output[0] == 0 {
				grid[i][j] = "."
			} else if output[0] == 1 {
				grid[i][j] = "#"
			}
		}
	}
	x, y := getClosestPoint(grid)
	log.Printf("Part Two - %d\n", (x*10000 + y))
}

func main() {
	var filePath string
	if len(os.Args) != 2 {
		filePath = "day-19/src/resources/input.txt"
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
