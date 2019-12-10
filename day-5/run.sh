#!/bin/bash

echo -e "Getting dependencies..."
go get "github.com/fatih/color"
echo -e "Dependencies downloaded\n"

go build -o target/IntcodeRunner src/main/intcode/IntcodeRunner.go
go build -o target/IntcodeTests src/tests/IntcodeTests.go
go build -o target/ThermalEnvironmentSupervisionTerminal src/main/ThermalEnvironmentSupervisionTerminal.go

echo -e "Performing some Intcode tests...\n"
./target/IntcodeTests
echo -e "\nEnd of tests\n\n"

echo -e "Running Thermal Environment Supervision Terminal...\n"
./target/ThermalEnvironmentSupervisionTerminal
echo -e "\nEnd of Thermal Environment Supervision Terminal"
