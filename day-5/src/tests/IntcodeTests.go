package main

import (
	"fmt"
	"github.com/fatih/color"
	"os"
	"os/exec"
	"strings"
)

var intcodeAppPath = "./target/Intcode"

func main() {
	// Tests for position mode instructions 1 and 2
	testIntcode("1,2,3,0,99", "3,2,3,0,99")
	testIntcode("2,2,3,0,99", "0,2,3,0,99")
	testIntcode("1,0,0,0,99", "2,0,0,0,99")
	testIntcode("2,3,0,3,99", "2,3,0,6,99")
	testIntcode("2,4,4,5,99,0", "2,4,4,5,99,9801")
	testIntcode("1,1,1,4,99,5,6,0,99", "30,1,1,4,2,5,6,0,99")
	testIntcode("1,9,10,3,2,3,11,0,99,30,40,50", "3500,9,10,70,2,3,11,0,99,30,40,50")

	// Tests for immediate mode instructions 1 and 2
	testIntcode("1001,2,3,0,99", "6,2,3,0,99")
	testIntcode("0101,2,3,0,99", "2,2,3,0,99")
	testIntcode("1101,2,3,0,99", "5,2,3,0,99")
	testIntcode("1002,4,3,4,33", "1002,4,3,4,99")
	testIntcode("1002,2,3,0,99", "9,2,3,0,99")
	testIntcode("0102,2,3,0,99", "0,2,3,0,99")
	testIntcode("1102,2,3,0,99", "6,2,3,0,99")

	os.Exit(0)
}

func testIntcode(input string, expectedOutput string) {
	output, _ := execIntcode(input)
	if output != expectedOutput {
		c := color.New(color.FgRed)
		c.Print("Test failed!")
		c.DisableColor()
		c.Printf(" - For input %s expeced output was %s and receive output %s\n", input, expectedOutput, output)
	} else {
		c := color.New(color.FgGreen)
		c.Print("Test passed!")
		c.DisableColor()
		c.Printf(" - For input %s excpected output was %s and receive output %s\n", input, expectedOutput, output)
	}
}

func execIntcode(input string) (string, error) {
	out, err := exec.Command(intcodeAppPath, "--inputArray", input).CombinedOutput()
	if err != nil {
		fmt.Fprintln(os.Stderr, "Intcode exit with code: "+err.Error())
		fmt.Fprintln(os.Stderr, "Intcode output:")
		fmt.Fprintln(os.Stderr, string(out))
		os.Exit(1)
	}
	return strings.TrimSuffix(string(out), "\n"), nil
}
