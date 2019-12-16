package intcode

import (
	"errors"
	"strconv"
)

const (
	FINISH               int = 99
	ADD                  int = 1
	MULTIPLY             int = 2
	INPUT                int = 3
	OUTPUT               int = 4
	JUMP_IF_TRUE         int = 5
	JUMP_IF_FALSE        int = 6
	LESS_THAN            int = 7
	EQUALS               int = 8
	RELATIVE_BASE_OFFSET int = 9
)

const (
	POSITION_MODE  int = 0
	IMMEDIATE_MODE int = 1
	RELATIVE_MODE  int = 2
)

const (
	END_OF_CODE    int = 0
	WAIT_FOR_INPUT int = 1
	END_WITH_ERROR int = -1
)

type Intcode struct {
	code         []int
	Pointer      int
	relativeBase int
}

func NewComputer(givenCode []int) *Intcode {
	code := make([]int, 10*len(givenCode))
	copy(code, givenCode)
	return &Intcode{code, 0, 0}
}

func (intcode *Intcode) Process() (int, []int, error) {
	results := []int{}
	for intcode.Pointer < len(intcode.code) {
		opcode, firstParamMode, secondParamMode, thirdParamMode, err := interpretInstruction(intcode.code[intcode.Pointer])
		if err != nil {
			return END_WITH_ERROR, results, err
		}
		switch opcode {
		case FINISH:
			intcode.Pointer = len(intcode.code)
		case ADD:
			firstParam, _ := intcode.getParam(1, firstParamMode)
			secondParam, _ := intcode.getParam(2, secondParamMode)
			intcode.code[intcode.getWritePosition(3, thirdParamMode)] = firstParam + secondParam
			intcode.Pointer += 4
		case MULTIPLY:
			firstParam, _ := intcode.getParam(1, firstParamMode)
			secondParam, _ := intcode.getParam(2, secondParamMode)
			intcode.code[intcode.getWritePosition(3, thirdParamMode)] = firstParam * secondParam
			intcode.Pointer += 4
		case INPUT:
			return WAIT_FOR_INPUT, results, nil
		case OUTPUT:
			firstParam, _ := intcode.getParam(1, firstParamMode)
			results = append(results, firstParam)
			intcode.Pointer += 2
		case JUMP_IF_TRUE:
			firstParam, _ := intcode.getParam(1, firstParamMode)
			if firstParam != 0 {
				secondParam, _ := intcode.getParam(2, secondParamMode)
				intcode.Pointer = secondParam
			} else {
				intcode.Pointer += 3
			}
		case JUMP_IF_FALSE:
			firstParam, _ := intcode.getParam(1, firstParamMode)
			if firstParam == 0 {
				secondParam, _ := intcode.getParam(2, secondParamMode)
				intcode.Pointer = secondParam
			} else {
				intcode.Pointer += 3
			}
		case LESS_THAN:
			firstParam, _ := intcode.getParam(1, firstParamMode)
			secondParam, _ := intcode.getParam(2, secondParamMode)
			if firstParam < secondParam {
				intcode.code[intcode.getWritePosition(3, thirdParamMode)] = 1
			} else {
				intcode.code[intcode.getWritePosition(3, thirdParamMode)] = 0
			}
			intcode.Pointer += 4
		case EQUALS:
			firstParam, _ := intcode.getParam(1, firstParamMode)
			secondParam, _ := intcode.getParam(2, secondParamMode)
			if firstParam == secondParam {
				intcode.code[intcode.getWritePosition(3, thirdParamMode)] = 1
			} else {
				intcode.code[intcode.getWritePosition(3, thirdParamMode)] = 0
			}
			intcode.Pointer += 4
		case RELATIVE_BASE_OFFSET:
			firstParam, _ := intcode.getParam(1, firstParamMode)
			intcode.relativeBase += firstParam
			intcode.Pointer += 2
		default:
			return END_WITH_ERROR, results, errors.New("Unknown opcode: " + strconv.Itoa(int(opcode)))
		}
	}
	return END_OF_CODE, results, nil
}

func (intcode *Intcode) ContinueProcess(input int) (int, []int, error) {
	_, firstParamMode, _, _, err := interpretInstruction(intcode.code[intcode.Pointer])
	if err != nil {
		return END_WITH_ERROR, []int{}, err
	}
	intcode.code[intcode.getWritePosition(1, firstParamMode)] = input
	intcode.Pointer += 2
	return intcode.Process()
}

func interpretInstruction(instruction int) (int, int, int, int, error) {
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

func splitInstruction(instruction string) (int, int, int, int, error) {
	opcode, err := strconv.Atoi(instruction[3:])
	if err != nil {
		return FINISH, POSITION_MODE, POSITION_MODE, POSITION_MODE, err
	}
	firstParamMode, err := strconv.Atoi(string(instruction[2]))
	if err != nil {
		return FINISH, POSITION_MODE, POSITION_MODE, POSITION_MODE, err
	}
	secondParamMode, err := strconv.Atoi(string(instruction[1]))
	if err != nil {
		return FINISH, POSITION_MODE, POSITION_MODE, POSITION_MODE, err
	}
	thirdParamMode, err := strconv.Atoi(string(instruction[0]))
	if err != nil {
		return FINISH, POSITION_MODE, POSITION_MODE, POSITION_MODE, err
	}
	return opcode, firstParamMode, secondParamMode, thirdParamMode, nil
}

func (intcode *Intcode) getParam(pointerDelta int, paramMode int) (int, error) {
	if paramMode == POSITION_MODE {
		return intcode.code[intcode.code[intcode.Pointer+pointerDelta]], nil
	} else if paramMode == IMMEDIATE_MODE {
		return intcode.code[intcode.Pointer+pointerDelta], nil
	} else if paramMode == RELATIVE_MODE {
		return intcode.code[intcode.code[intcode.Pointer+pointerDelta]+intcode.relativeBase], nil
	}
	return -1, errors.New("Unknown parameter mode " + string(paramMode))
}

func (intcode *Intcode) getWritePosition(pointerDelta int, paramMode int) int {
	if paramMode == RELATIVE_MODE {
		return intcode.code[intcode.Pointer+pointerDelta] + intcode.relativeBase
	}
	return intcode.code[intcode.Pointer+pointerDelta]
}
