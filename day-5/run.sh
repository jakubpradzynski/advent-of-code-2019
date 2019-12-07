#!/bin/bash

go build -o target/Intcode src/main/intcode/Intcode.go
go build -o target/IntcodeTests src/tests/IntcodeTests.go
./target/IntcodeTests
