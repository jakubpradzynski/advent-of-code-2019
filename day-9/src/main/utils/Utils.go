package utils

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func GeneratePermutations(array []int) [][]int {
	var helper func([]int, int)
	res := [][]int{}

	helper = func(arr []int, n int) {
		if n == 1 {
			tmp := make([]int, len(arr))
			copy(tmp, arr)
			res = append(res, tmp)
		} else {
			for i := 0; i < n; i++ {
				helper(arr, n-1)
				if n%2 == 1 {
					tmp := arr[i]
					arr[i] = arr[n-1]
					arr[n-1] = tmp
				} else {
					tmp := arr[0]
					arr[0] = arr[n-1]
					arr[n-1] = tmp
				}
			}
		}
	}
	helper(array, len(array))
	return res
}

func StringArrayToIntArray(stringArray []string) []int {
	intArray := []int{}
	for _, str := range stringArray {
		value, _ := strconv.Atoi(str)
		intArray = append(intArray, value)
	}
	return intArray
}

func IntArrayToString(array []int) string {
	return strings.Trim(strings.Join(strings.Fields(fmt.Sprint(array)), ","), "[]")
}

func ReadFirstLineInFile(inputFilePath string) ([]string, error) {
	file, err := os.Open(inputFilePath)
	if err != nil {
		return nil, errors.New("Can not read file " + inputFilePath)
	}
	scanner := bufio.NewScanner(file)
	scanner.Scan()
	input := scanner.Text()
	return strings.Split(input, ","), nil
}

func StringToIntArray(val string) []int {
	stringArray := strings.Split(val, ",")
	intArray := []int{}
	for _, str := range stringArray {
		value, _ := strconv.Atoi(str)
		intArray = append(intArray, value)
	}
	return intArray
}
