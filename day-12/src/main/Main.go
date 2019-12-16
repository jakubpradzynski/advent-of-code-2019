package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"math"
	"os"
	"path/filepath"
	"strings"
)

type Point3D struct {
	x int
	y int
	z int
}

type Moon struct {
	position Point3D
	velocity Point3D
}

func getMoons(fileContent string) []Moon {
	moons := []Moon{}
	for _, line := range strings.Split(strings.Trim(string(fileContent), "\n"), "\n") {
		var x, y, z int
		fmt.Sscanf(line, "<x=%d, y=%d, z=%d>", &x, &y, &z)
		moons = append(moons, Moon{Point3D{x, y, z}, Point3D{0, 0, 0}})
	}
	return moons
}

func getVelocityChange(a int, b int) int {
	if a < b {
		return 1
	} else if a > b {
		return -1
	}
	return 0
}

func calculateMoonNewVelocity(moon1 Moon, moon2 Moon) Moon {
	xChange := getVelocityChange(moon1.position.x, moon2.position.x)
	yChange := getVelocityChange(moon1.position.y, moon2.position.y)
	zChange := getVelocityChange(moon1.position.z, moon2.position.z)
	moon1.velocity.x += xChange
	moon1.velocity.y += yChange
	moon1.velocity.z += zChange
	return moon1
}

func calculateMoonNewVelocityForDirection(moon1 Moon, moon2 Moon, direction string) Moon {
	if direction == "x" {
		xChange := getVelocityChange(moon1.position.x, moon2.position.x)
		moon1.velocity.x += xChange
	}
	if direction == "y" {
		yChange := getVelocityChange(moon1.position.y, moon2.position.y)
		moon1.velocity.y += yChange
	}
	if direction == "z" {
		zChange := getVelocityChange(moon1.position.z, moon2.position.z)
		moon1.velocity.z += zChange
	}
	return moon1
}

func applyVelocity(moons []Moon) []Moon {
	for i, moon := range moons {
		moon.position.x += moon.velocity.x
		moon.position.y += moon.velocity.y
		moon.position.z += moon.velocity.z
		moons[i] = moon
	}
	return moons
}

func simulateMotion(moons []Moon, iterations int) []Moon {
	for i := 0; i < iterations; i++ {
		for j, moon1 := range moons {
			for k, moon2 := range moons {
				if j != k {
					moon := calculateMoonNewVelocity(moon1, moon2)
					moons[j] = moon
					moon1 = moon
				}
			}
		}
		moons = applyVelocity(moons)
	}
	return moons
}

func hasTheSamePositions(moons []Moon, startPosition []Moon) bool {
	for i := 0; i < len(moons); i++ {
		if (moons[i].position.x != startPosition[i].position.x) || (moons[i].position.y != startPosition[i].position.y) || (moons[i].position.z != startPosition[i].position.z || moons[i].velocity.x != startPosition[i].velocity.x || moons[i].velocity.y != startPosition[i].velocity.y || moons[i].velocity.z != startPosition[i].velocity.z) {
			return false
		}
	}
	return true
}

func simulateMotionUntilStartPosition(moons []Moon, direction string) int {
	startPosition := make([]Moon, len(moons))
	copy(startPosition, moons)
	iterations := 0
	areAtStartPosition := false
	for !areAtStartPosition {
		for j, moon1 := range moons {
			for k, moon2 := range moons {
				if j != k {
					moon1 = calculateMoonNewVelocityForDirection(moon1, moon2, direction)
					moons[j] = moon1
				}
			}
		}
		moons = applyVelocity(moons)
		iterations++
		if hasTheSamePositions(moons, startPosition) {
			areAtStartPosition = true
		}
	}
	return iterations
}

func calculateTotalEnergyOfMoon(moon Moon) int {
	potentialEnergy := math.Abs(float64(moon.position.x)) + math.Abs(float64(moon.position.y)) + math.Abs(float64(moon.position.z))
	kineticEnergy := math.Abs(float64(moon.velocity.x)) + math.Abs(float64(moon.velocity.y)) + math.Abs(float64(moon.velocity.z))
	return int(potentialEnergy * kineticEnergy)
}

func calculateTotalEnergy(moons []Moon) int {
	totalEnergy := 0
	for _, moon := range moons {
		totalEnergy += calculateTotalEnergyOfMoon(moon)
	}
	return totalEnergy
}

func partOne(fileContent string) {
	moons := getMoons(string(fileContent))
	log.Println(moons)
	moonsAfterSimulation := simulateMotion(moons, 1000)
	log.Println(moonsAfterSimulation)
	log.Printf("Part One - %d\n", calculateTotalEnergy(moonsAfterSimulation))
}

func gcd(a int, b int) int {
	for b != 0 {
		t := b
		b = a % b
		a = t
	}
	return a
}

func lcm(a int, b int, integers ...int) int {
	result := a * b / gcd(a, b)
	for i := 0; i < len(integers); i++ {
		result = lcm(result, integers[i])
	}
	return result
}

func partTwo(fileContent string) {
	moons := getMoons(string(fileContent))
	xIterations := simulateMotionUntilStartPosition(moons, "x")
	yIterations := simulateMotionUntilStartPosition(moons, "y")
	zIterations := simulateMotionUntilStartPosition(moons, "z")
	log.Printf("Part Two - %d\n", lcm(xIterations, yIterations, zIterations))
}

func main() {
	var filePath string
	if len(os.Args) != 2 {
		filePath = "day-12/src/resources/input.txt"
	} else {
		filePath = os.Args[1]
	}
	path, _ := filepath.Abs(filePath)
	fileContent, err := ioutil.ReadFile(path)
	if err != nil {
		log.Fatal(err)
	}

	partOne(string(fileContent))
	partTwo(string(fileContent))
}
