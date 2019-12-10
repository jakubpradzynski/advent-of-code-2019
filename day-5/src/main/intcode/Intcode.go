package intcode

import "strconv"

type Intcode struct {
	code []int
}

func New(code []int) Intcode {
	return Intcode{code}
}

func (intcode Intcode) Process() {
	pointer := 0
	for pointer < len(intcode.code) {
		instructionType, firstParamMode, secondParamMode, _ := interpretInstruction(intcode.code[pointer])
		switch instructionType {
		case 99:
			pointer = len(intcode.code)
		}
	}
}

func interpretInstruction(instruction int) (int, int, int, int) {
	return splitInstruction(fullLengthInstruction(instruction))
}

func fullLengthInstruction(instruction int) string {
	instructionAsString := strconv.Itoa(instruction)
	fullLengthOpcode := "00000"
	switch len(instructionAsString) {
	case 1:
		fullLengthOpcode = "0000" + instructionAsString
	case 2:
		fullLengthOpcode = "000" + instructionAsString
	case 3:
		fullLengthOpcode = "00" + instructionAsString
	case 4:
		fullLengthOpcode = "0" + instructionAsString
	case 5:
		fullLengthOpcode = instructionAsString
	}
	return fullLengthOpcode
}

func splitInstruction(instruction string) (int, int, int, int) {
	instructionType, _ := strconv.Atoi(string(instruction[3:]))
	firstParamMode, _ := strconv.Atoi(string(instruction[2]))
	secondParamMode, _ := strconv.Atoi(string(instruction[1]))
	thirdParamMode, _ := strconv.Atoi(string(instruction[0]))
	return instructionType, firstParamMode, secondParamMode, thirdParamMode
}
