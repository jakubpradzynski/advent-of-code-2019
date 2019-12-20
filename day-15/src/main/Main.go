package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"sort"
	"strconv"
	"strings"

	"./intcode"
)

type Point struct {
	x int
	y int
}

type Neighbour struct {
	position Point
	contain  string
}

// Movement commands
const (
	NORTH int = 1
	SOUTH int = 2
	WEST  int = 3
	EAST  int = 4
)

// Status codes
const (
	HIT_A_WALL             int = 0
	MOVED                  int = 1
	MOVED_ON_OXYGEN_SYSTEM int = 2
)

const (
	DROID        string = "D"
	WALL         string = "#"
	CAN_TRAVERSE string = "."
	NOT_WISITED  string = " "
	OXYGEN       string = "O"
)

const MAZE_SIZE = 50

func getEmptyMaze() [MAZE_SIZE][MAZE_SIZE]string {
	maze := [MAZE_SIZE][MAZE_SIZE]string{}
	for i := 0; i < MAZE_SIZE; i++ {
		for j := 0; j < MAZE_SIZE; j++ {
			maze[i][j] = NOT_WISITED
		}
	}
	return maze
}

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

func getPosition(currentPosition Point, moveDirection int) (Point, error) {
	switch moveDirection {
	case NORTH:
		return Point{currentPosition.x, currentPosition.y - 1}, nil
	case SOUTH:
		return Point{currentPosition.x, currentPosition.y + 1}, nil
	case WEST:
		return Point{currentPosition.x - 1, currentPosition.y}, nil
	case EAST:
		return Point{currentPosition.x + 1, currentPosition.y}, nil
	}
	return Point{0, 0}, fmt.Errorf("Unknown direction! - %d\n", moveDirection)
}

func move(currentDroidPosition Point, moveDirection int, inputToCurrentPosition []int) {
	newPosition, err := getPosition(currentDroidPosition, moveDirection)
	if err != nil {
		log.Fatal(err)
	}

	if maze[newPosition.x][newPosition.y] != NOT_WISITED {
		return
	}

	code := make([]int, len(originalCode))
	copy(code, originalCode)
	computer := intcode.NewComputer(code)
	computer.Process(nil)
	for i := 0; i < len(inputToCurrentPosition); i++ {
		computer.Process(&inputToCurrentPosition[i])
	}

	endStatus, output, err := computer.Process(&moveDirection)
	if err != nil {
		log.Fatal(err)
	}
	if endStatus == intcode.END_OF_CODE {
		log.Fatal("End of code!")
	}
	if len(output) != 1 {
		log.Printf("Output different than 1! - %d\n", len(output))
	} else {
		switch output[0] {
		case HIT_A_WALL:
			maze[newPosition.x][newPosition.y] = WALL
		case MOVED:
			maze[newPosition.x][newPosition.y] = CAN_TRAVERSE
			move(newPosition, NORTH, append(inputToCurrentPosition, moveDirection))
			move(newPosition, EAST, append(inputToCurrentPosition, moveDirection))
			move(newPosition, SOUTH, append(inputToCurrentPosition, moveDirection))
			move(newPosition, WEST, append(inputToCurrentPosition, moveDirection))
		case MOVED_ON_OXYGEN_SYSTEM:
			maze[newPosition.x][newPosition.y] = OXYGEN
			pathSizes = append(pathSizes, len(inputToCurrentPosition)+1)
		}
	}
}

func partOne() {
	droidPosition := Point{MAZE_SIZE / 2, MAZE_SIZE / 2}
	maze[droidPosition.x][droidPosition.x] = DROID
	move(droidPosition, NORTH, []int{})
	move(droidPosition, WEST, []int{})
	move(droidPosition, SOUTH, []int{})
	move(droidPosition, EAST, []int{})
}

func printMaze() {
	for i := 0; i < MAZE_SIZE; i++ {
		for j := 0; j < MAZE_SIZE; j++ {
			fmt.Print(maze[i][j])
		}
		fmt.Println()
	}
}

func getOxygenPosition() Point {
	var x, y int
	for i := 0; i < MAZE_SIZE; i++ {
		for j := 0; j < MAZE_SIZE; j++ {
			if maze[i][j] == OXYGEN {
				x = i
				y = j
			}
		}
	}
	return Point{x, y}
}

func fillOxygen(currentPosition Point, moveDirection int, deep int) {
	times = append(times, deep)
	newPosition, err := getPosition(currentPosition, moveDirection)
	if err != nil {
		log.Fatal(err)
	}
	if maze[newPosition.x][newPosition.y] != CAN_TRAVERSE {
		return
	}
	maze[newPosition.x][newPosition.y] = OXYGEN
	fillOxygen(newPosition, NORTH, deep+1)
	fillOxygen(newPosition, EAST, deep+1)
	fillOxygen(newPosition, SOUTH, deep+1)
	fillOxygen(newPosition, WEST, deep+1)
}

func partTwo() {
	droidPosition := Point{MAZE_SIZE / 2, MAZE_SIZE / 2}
	maze[droidPosition.x][droidPosition.y] = CAN_TRAVERSE
	oxygenPosition := getOxygenPosition()
	fillOxygen(oxygenPosition, NORTH, 0)
	fillOxygen(oxygenPosition, EAST, 0)
	fillOxygen(oxygenPosition, SOUTH, 0)
	fillOxygen(oxygenPosition, WEST, 0)
}

var originalCode []int
var maze [MAZE_SIZE][MAZE_SIZE]string = getEmptyMaze()
var pathSizes []int = []int{}
var times []int = []int{}

func main() {
	var filePath string
	if len(os.Args) != 2 {
		filePath = "day-15/src/resources/input.txt"
	} else {
		filePath = os.Args[1]
	}
	path, _ := filepath.Abs(filePath)
	fileContent, err := ioutil.ReadFile(path)
	if err != nil {
		log.Fatal(err)
	}
	originalCode = fileContentToIntArray(string(fileContent))
	partOne()
	printMaze()
	sort.Ints(pathSizes)
	fmt.Printf("Part One - %d\n", pathSizes[0])
	partTwo()
	sort.Ints(times)
	fmt.Printf("Part Two - %d\n", times[len(times)-1])
}
