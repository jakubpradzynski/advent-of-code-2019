package main

import (
	"github.com/fatih/color"
	"io"
	"os"
	"os/exec"
)

var intcodeAppPath = "./target/Intcode"
var inputFilePath = "./src/resources/input.txt"

func main() {
	err := execIntcode(inputFilePath)
	if err != nil {
		c := color.New(color.FgRed)
		c.Fprintln(os.Stderr, "Error occured:")
		c.DisableColor()
		c.Println(err.Error())
		os.Exit(1)
	}
}

func execIntcode(input string) error {
	cmd := exec.Command(intcodeAppPath, "-inputFilePath", input)
	var stdin io.WriteCloser
	var err error
	if stdin, err = cmd.StdinPipe(); err != nil {
		return err
	}
	defer stdin.Close()
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	if err = cmd.Start(); err != nil {
		return err
	}
	if _, err = io.WriteString(stdin, "1\n"); err != nil {
		return err
	}
	if err = cmd.Wait(); err != nil {
		return err
	}
	return nil
}
