#!/bin/bash

echo -e "Getting dependencies..."
go get "github.com/fatih/color"
echo -e "Dependencies downloaded\n"

echo -e "Performing some Intcode tests...\n"
go run src/main/intcode/tests/IntcodeTests.go
echo -e "\nEnd of tests\n\n"

echo -e "Running Basic Operation Of System Test...\n"
echo "Part One: "
echo 1 | go run src/main/intcode/IntcodeRunner.go --inputFilePath src/resources/input.txt
echo "Part Two: "
echo 2 | go run src/main/intcode/IntcodeRunner.go --inputFilePath src/resources/input.txt
echo -e "\nEnd"