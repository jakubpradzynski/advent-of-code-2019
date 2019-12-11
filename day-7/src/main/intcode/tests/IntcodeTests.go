package main

import (
	"bytes"
	"fmt"
	"os/exec"
	"strings"

	"github.com/fatih/color"
)

func main() {
	fmt.Println("Tests for FINISH")
	testIntcode("99", "")
	testIntcode("1199", "")
	testIntcode("1099", "")
	testIntcode("199", "")

	fmt.Println("Tests for ADD")
	testIntcode("1,2,3,0,4,0,99", "3")
	testIntcode("1101,2,3,0,4,0,99", "5")
	testIntcode("1001,2,3,0,4,0,99", "6")
	testIntcode("101,2,3,0,4,0,99", "2")

	fmt.Println("Tests for MULTIPLY")
	testIntcode("2,2,3,0,4,0,99", "0")
	testIntcode("1102,2,3,0,4,0,99", "6")
	testIntcode("1002,2,3,0,4,0,99", "9")
	testIntcode("102,2,3,0,4,0,99", "0")

	fmt.Println("Tests for INPUT")
	testIntcode("3,1,4,1,99", "5", "5")
	testIntcode("1103,1,4,1,99", "5", "5")
	testIntcode("1003,1,4,1,99", "5", "5")
	testIntcode("103,1,4,1,99", "5", "5")

	fmt.Println("Tests for OUTPUT")
	testIntcode("4,0,99", "4")
	testIntcode("1104,0,99", "0")
	testIntcode("1004,0,99", "1004")
	testIntcode("104,0,99", "0")

	fmt.Println("Tests for JUMP IF TRUE")
	// testIntcode("5,0,5,4,0,99", "")
	// testIntcode("5,1,3,4,0,99", "5")

	fmt.Println("Tests for JUMP IF FALSE")

	fmt.Println("Tests for LESS THAN")

	fmt.Println("Tests for EQUALS")

	fmt.Println("Various tests from tasks")
	testIntcode("1,9,10,3,2,3,11,0,99,30,40,50", "")                                                                                                                                     // Without printing
	testIntcode("3,0,4,0,99", "5", "5")                                                                                                                                                  // Print anything given
	testIntcode("1002,4,3,4,33", "")                                                                                                                                                     // Just end
	testIntcode("3,9,8,9,10,9,4,9,99,-1,8", "1", "8")                                                                                                                                    // If input is 8, output is 1
	testIntcode("3,9,8,9,10,9,4,9,99,-1,8", "0", "10")                                                                                                                                   // If input is not 8, output is 0
	testIntcode("3,9,8,9,10,9,4,9,99,-1,8", "0", "5")                                                                                                                                    // If input is not 8, output is 0
	testIntcode("3,9,7,9,10,9,4,9,99,-1,8", "1", "7")                                                                                                                                    // If input is less then 8, output is 1
	testIntcode("3,9,7,9,10,9,4,9,99,-1,8", "0", "8")                                                                                                                                    // If input is above or is 8, output 0
	testIntcode("3,3,1108,-1,8,3,4,3,99", "1", "8")                                                                                                                                      // If input is 8, output is 1
	testIntcode("3,3,1108,-1,8,3,4,3,99", "0", "9")                                                                                                                                      // If input is not 8, output is 0
	testIntcode("3,3,1108,-1,8,3,4,3,99", "0", "7")                                                                                                                                      // If input is not 8, outpus is 0
	testIntcode("3,3,1107,-1,8,3,4,3,99", "1", "7")                                                                                                                                      // If input is less than 8, output is 1
	testIntcode("3,3,1107,-1,8,3,4,3,99", "0", "8")                                                                                                                                      // If input is above or is 8, output is 0
	testIntcode("3,12,6,12,15,1,13,14,13,4,13,99,-1,0,1,9", "0", "0")                                                                                                                    // If input is 0, output is 0
	testIntcode("3,12,6,12,15,1,13,14,13,4,13,99,-1,0,1,9", "1", "5")                                                                                                                    // If input is not 0, output is 1
	testIntcode("3,3,1105,-1,9,1101,0,0,12,4,12,99,1", "0", "0")                                                                                                                         // If input is 0, output is 0
	testIntcode("3,3,1105,-1,9,1101,0,0,12,4,12,99,1", "1", "3")                                                                                                                         // If input is not 0, output is 1
	testIntcode("3,21,1008,21,8,20,1005,20,22,107,8,21,20,1006,20,31,1106,0,36,98,0,0,1002,21,125,20,4,20,1105,1,46,104,999,1105,1,46,1101,1000,1,20,4,20,1105,1,46,98,99", "999", "7")  // If input is below 8, output is 999
	testIntcode("3,21,1008,21,8,20,1005,20,22,107,8,21,20,1006,20,31,1106,0,36,98,0,0,1002,21,125,20,4,20,1105,1,46,104,999,1105,1,46,1101,1000,1,20,4,20,1105,1,46,98,99", "1000", "8") // If input is 8, output is 1000
	testIntcode("3,21,1008,21,8,20,1005,20,22,107,8,21,20,1006,20,31,1106,0,36,98,0,0,1002,21,125,20,4,20,1105,1,46,104,999,1105,1,46,1101,1000,1,20,4,20,1105,1,46,98,99", "1001", "9") // If input is above 8, output is 1001
	testIntcode("3,3,1105,-1,9,1101,0,0,12,4,12,99,1", "1", "100")
}

func testIntcode(code string, expectedOutput string, input ...string) {
	cmd := exec.Command("go", "run", "src/main/intcode/IntcodeRunner.go", "-inputArray", code)
	if len(input) > 0 {
		buffer := bytes.Buffer{}
		buffer.Write([]byte(input[0]))
		cmd.Stdin = &buffer
	}
	combinedOutput, err := cmd.CombinedOutput()
	if err != nil {
		c := color.New(color.FgRed)
		c.Print("Test failed!")
		c.DisableColor()
		c.Printf(" - %s\n", err.Error())
		c.Printf("%s\n", string(combinedOutput))
	} else {
		output := string(combinedOutput)[:strings.IndexByte(string(combinedOutput), '\n')]
		if string(output) != expectedOutput {
			c := color.New(color.FgRed)
			c.Print("Test failed!")
			c.DisableColor()
			c.Printf(" - For input %s expeced output was %s and receive output %s\n", code, expectedOutput, string(output))
		} else {
			c := color.New(color.FgGreen)
			c.Print("Test passed!")
			c.DisableColor()
			c.Printf(" - For input %s excpected output was %s and receive output %s\n", code, expectedOutput, string(output))
		}
	}
}
