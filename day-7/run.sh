#!/bin/bash

echo -e "Getting dependencies..."
go get "github.com/fatih/color"
echo -e "Dependencies downloaded\n"

echo -e "Performing some Intcode tests...\n"
go run src/main/intcode/tests/IntcodeTests.go
echo -e "\nEnd of tests\n\n"

echo -e "Running Amplifier Controller Software...\n"
go run src/main/AmplifierControllerSoftware.go
echo -e "\nEnd"