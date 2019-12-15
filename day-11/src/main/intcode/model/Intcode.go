package model

import (
	"errors"
	"fmt"
	"strconv"

	"../stream"
)

type Opcode int
type ParameterMode int

const (
	FINISH               Opcode = 99
	ADD                  Opcode = 1
	MULTIPLY             Opcode = 2
	INPUT                Opcode = 3
	OUTPUT               Opcode = 4
	JUMP_IF_TRUE         Opcode = 5
	JUMP_IF_FALSE        Opcode = 6
	LESS_THAN            Opcode = 7
	EQUALS               Opcode = 8
	RELATIVE_BASE_OFFSET Opcode = 9
)

const (
	POSITION_MODE  ParameterMode = 0
	IMMEDIATE_MODE ParameterMode = 1
	RELATIVE_MODE  ParameterMode = 2
)

type Intcode struct {
	code         []int
	inputStream  *stream.Stream
	outputStream *stream.Stream
}

func New(givenCode []int, inputStream *stream.Stream, outputStream *stream.Stream) Intcode {
	code := make([]int, 10*len(givenCode))
	copy(code, givenCode)
	return Intcode{code, inputStream, outputStream}
}

func (intcode Intcode) Process() ([]int, error) {
	pointer := 0
	results := []int{}
	lastValueIndex := -1
	relativeBase := 0
	for pointer < len(intcode.code) {
		opcode, firstParamMode, secondParamMode, thirdParamMode, err := interpretInstruction(intcode.code[pointer])
		if err != nil {
			return results, err
		}
		switch opcode {
		case FINISH:
			pointer = len(intcode.code)
		case ADD:
			firstParam, _ := getParam(intcode.code, pointer+1, firstParamMode, relativeBase)
			secondParam, _ := getParam(intcode.code, pointer+2, secondParamMode, relativeBase)
			intcode.code[getWritePosition(intcode.code, pointer+3, thirdParamMode, relativeBase)] = firstParam + secondParam
			pointer += 4
		case MULTIPLY:
			firstParam, _ := getParam(intcode.code, pointer+1, firstParamMode, relativeBase)
			secondParam, _ := getParam(intcode.code, pointer+2, secondParamMode, relativeBase)
			intcode.code[getWritePosition(intcode.code, pointer+3, thirdParamMode, relativeBase)] = firstParam * secondParam
			pointer += 4
		case INPUT:
			var value int
			if intcode.inputStream != nil {
				channel := make(chan int)
				intcode.inputStream.WaitForNewData(channel, lastValueIndex)
				val := <-channel
				value = val
				lastValueIndex++
			} else {
				if _, err := fmt.Scanf("%d", &value); err != nil {
					return results, err
				}
			}
			intcode.code[getWritePosition(intcode.code, pointer+1, firstParamMode, relativeBase)] = value
			pointer += 2
		case OUTPUT:
			firstParam, _ := getParam(intcode.code, pointer+1, firstParamMode, relativeBase)
			if intcode.outputStream != nil {
				intcode.outputStream.SendNewData(firstParam)
			}
			results = append(results, firstParam)
			pointer += 2
		case JUMP_IF_TRUE:
			firstParam, _ := getParam(intcode.code, pointer+1, firstParamMode, relativeBase)
			if firstParam != 0 {
				secondParam, _ := getParam(intcode.code, pointer+2, secondParamMode, relativeBase)
				pointer = secondParam
			} else {
				pointer += 3
			}
		case JUMP_IF_FALSE:
			firstParam, _ := getParam(intcode.code, pointer+1, firstParamMode, relativeBase)
			if firstParam == 0 {
				secondParam, _ := getParam(intcode.code, pointer+2, secondParamMode, relativeBase)
				pointer = secondParam
			} else {
				pointer += 3
			}
		case LESS_THAN:
			firstParam, _ := getParam(intcode.code, pointer+1, firstParamMode, relativeBase)
			secondParam, _ := getParam(intcode.code, pointer+2, secondParamMode, relativeBase)
			if firstParam < secondParam {
				intcode.code[getWritePosition(intcode.code, pointer+3, thirdParamMode, relativeBase)] = 1
			} else {
				intcode.code[getWritePosition(intcode.code, pointer+3, thirdParamMode, relativeBase)] = 0
			}
			pointer += 4
		case EQUALS:
			firstParam, _ := getParam(intcode.code, pointer+1, firstParamMode, relativeBase)
			secondParam, _ := getParam(intcode.code, pointer+2, secondParamMode, relativeBase)
			if firstParam == secondParam {
				intcode.code[getWritePosition(intcode.code, pointer+3, thirdParamMode, relativeBase)] = 1
			} else {
				intcode.code[getWritePosition(intcode.code, pointer+3, thirdParamMode, relativeBase)] = 0
			}
			pointer += 4
		case RELATIVE_BASE_OFFSET:
			firstParam, _ := getParam(intcode.code, pointer+1, firstParamMode, relativeBase)
			relativeBase += firstParam
			pointer += 2
		default:
			return results, errors.New("Unknown opcode: " + strconv.Itoa(int(opcode)))
		}
	}
	if intcode.inputStream != nil {
		intcode.inputStream.Close()
	}
	if intcode.outputStream != nil {
		intcode.outputStream.Close()
	}
	return results, nil
}

func interpretInstruction(instruction int) (Opcode, ParameterMode, ParameterMode, ParameterMode, error) {
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

func splitInstruction(instruction string) (Opcode, ParameterMode, ParameterMode, ParameterMode, error) {
	opcode, err := stringToOpcode(instruction[3:])
	if err != nil {
		return FINISH, POSITION_MODE, POSITION_MODE, POSITION_MODE, err
	}
	firstParamMode, err := stringToParameterMode(string(instruction[2]))
	if err != nil {
		return FINISH, POSITION_MODE, POSITION_MODE, POSITION_MODE, err
	}
	secondParamMode, err := stringToParameterMode(string(instruction[1]))
	if err != nil {
		return FINISH, POSITION_MODE, POSITION_MODE, POSITION_MODE, err
	}
	thirdParamMode, err := stringToParameterMode(string(instruction[0]))
	if err != nil {
		return FINISH, POSITION_MODE, POSITION_MODE, POSITION_MODE, err
	}
	return opcode, firstParamMode, secondParamMode, thirdParamMode, nil
}

func stringToOpcode(value string) (Opcode, error) {
	val, err := strconv.Atoi(value)
	if err != nil {
		return -1, err
	}
	switch val {
	case 1:
		return ADD, nil
	case 2:
		return MULTIPLY, nil
	case 3:
		return INPUT, nil
	case 4:
		return OUTPUT, nil
	case 5:
		return JUMP_IF_TRUE, nil
	case 6:
		return JUMP_IF_FALSE, nil
	case 7:
		return LESS_THAN, nil
	case 8:
		return EQUALS, nil
	case 9:
		return RELATIVE_BASE_OFFSET, nil
	case 99:
		return FINISH, nil
	}
	return -1, errors.New("Uknown opcode " + strconv.Itoa(val))
}

func stringToParameterMode(value string) (ParameterMode, error) {
	val, err := strconv.Atoi(value)
	if err != nil {
		return -1, err
	}
	switch val {
	case 0:
		return POSITION_MODE, nil
	case 1:
		return IMMEDIATE_MODE, nil
	case 2:
		return RELATIVE_MODE, nil
	}
	return -1, errors.New("Uknown parameter mode " + strconv.Itoa(val))
}

func getParam(code []int, pointer int, paramMode ParameterMode, relativeBase int) (int, error) {
	if paramMode == POSITION_MODE {
		return code[code[pointer]], nil
	} else if paramMode == IMMEDIATE_MODE {
		return code[pointer], nil
	} else if paramMode == RELATIVE_MODE {
		return code[code[pointer]+relativeBase], nil
	}
	return -1, errors.New("Unknown parameter mode " + string(paramMode))
}

func getWritePosition(code []int, pointer int, paramMode ParameterMode, relativeBase int) int {
	if paramMode == RELATIVE_MODE {
		return code[pointer] + relativeBase
	}
	return code[pointer]
}
