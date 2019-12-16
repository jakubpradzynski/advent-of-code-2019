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

const gridSize = 50

func countBlockTiles(grid [gridSize][gridSize]int) int {
	blockTilesCount := 0
	for i := 0; i < gridSize; i++ {
		for j := 0; j < gridSize; j++ {
			if grid[i][j] == 2 {
				blockTilesCount++
			}
		}
	}
	return blockTilesCount
}

func partOne(code []int) [gridSize][gridSize]int {
	grid := [gridSize][gridSize]int{}
	computer := intcode.NewComputer(code)
	exitCode, output, err := computer.Process()
	if err != nil {
		log.Fatal(err)
	}
	if exitCode == intcode.END_OF_CODE {
		for i := 0; i <= len(output)-2; i += 3 {
			grid[output[i]][output[i+1]] = output[i+2]
		}
	}
	log.Printf("Part One - %d\n", countBlockTiles(grid))
	return grid
}

func partTwo(code []int) {
	grid := [gridSize][gridSize]int{}
	code[0] = 2
	var exitCode int
	var output []int
	var err error
	computer := intcode.NewComputer(code)
	ballX, paddleX, lastScore := 0, 0, 0
	isFirstRun := true
	for true {
		if isFirstRun {
			exitCode, output, err = computer.Process()
			isFirstRun = false
		} else {
			joistickCommand := 0
			if ballX > paddleX {
				joistickCommand = 1
			} else if ballX < paddleX {
				joistickCommand = -1
			}
			exitCode, output, err = computer.ContinueProcess(joistickCommand)
		}
		if err != nil {
			log.Fatal(err)
		}
		for i := 0; i <= len(output)-2; i += 3 {
			if output[i] == -1 && output[i+1] == 0 {
				lastScore = output[i+2]
			} else {
				grid[output[i]][output[i+1]] = output[i+2]
				if output[i+2] == 4 {
					ballX = output[i]
				} else if output[i+2] == 3 {
					paddleX = output[i]
				}
			}
		}
		// screen.Clear()
		// screen.MoveTopLeft()
		// fmt.Println()
		// for i := 0; i < gridSize; i++ {
		// 	for j := 0; j < gridSize; j++ {
		// 		switch grid[j][i] {
		// 		case 0:
		// 			fmt.Printf(" ")
		// 		case 1:
		// 			fmt.Printf("|")
		// 		case 2:
		// 			fmt.Printf("#")
		// 		case 3:
		// 			fmt.Printf("_")
		// 		case 4:
		// 			fmt.Printf("o")
		// 		}
		// 	}
		// 	fmt.Println()
		// }
		if exitCode == intcode.END_OF_CODE {
			break
		}
	}
	log.Printf("Part Two - %d\n", lastScore)
}

func main() {
	var filePath string
	if len(os.Args) != 2 {
		filePath = "day-13/src/resources/input.txt"
	} else {
		filePath = os.Args[1]
	}
	path, _ := filepath.Abs(filePath)
	fileContent, err := ioutil.ReadFile(path)
	if err != nil {
		log.Fatal(err)
	}
	code := fileContentToIntArray(string(fileContent))
	partOne(code)
	partTwo(code)
}
