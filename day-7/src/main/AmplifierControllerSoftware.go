package main

import (
	"fmt"
	"os"

	"./intcode/model"
	"./intcode/stream"
	"./utils"
)

func main() {
	input, err := utils.ReadFirstLineInFile("src/resources/input.txt")
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(2)
	}
	sequences := utils.GeneratePermutations([]int{0, 1, 2, 3, 4})
	inputAsIntArray := utils.StringArrayToIntArray(input)
	max := 0
	bestSequence := []int{}
	for _, seq := range sequences {
		a := runAmplifierTest(inputAsIntArray, seq[0], 0)
		b := runAmplifierTest(inputAsIntArray, seq[1], a)
		c := runAmplifierTest(inputAsIntArray, seq[2], b)
		d := runAmplifierTest(inputAsIntArray, seq[3], c)
		e := runAmplifierTest(inputAsIntArray, seq[4], d)
		if e > max {
			max = e
			bestSequence = seq
		}
	}
	fmt.Println("Part One")
	fmt.Printf("Max thruster signal: %d\n", max)
	fmt.Printf("Best sequence: %s\n\n", utils.IntArrayToString(bestSequence))

	sequences = utils.GeneratePermutations([]int{5, 6, 7, 8, 9})
	max = 0
	bestSequence = []int{}
	for _, seq := range sequences {
		aInputStream := stream.New()
		bInputStream := stream.New()
		cInputStream := stream.New()
		dInputStream := stream.New()
		eInputStream := stream.New()
		signal := make(chan int)
		go func() {
			a := model.New(inputAsIntArray, aInputStream, bInputStream)
			a.Process()
		}()
		go func() {
			b := model.New(inputAsIntArray, bInputStream, cInputStream)
			b.Process()
		}()
		go func() {
			c := model.New(inputAsIntArray, cInputStream, dInputStream)
			c.Process()
		}()
		go func() {
			d := model.New(inputAsIntArray, dInputStream, eInputStream)
			d.Process()
		}()
		go func() {
			e := model.New(inputAsIntArray, eInputStream, aInputStream)
			output, _ := e.Process()
			signal <- output[len(output)-1]
		}()
		aInputStream.SendNewData(seq[0])
		bInputStream.SendNewData(seq[1])
		cInputStream.SendNewData(seq[2])
		dInputStream.SendNewData(seq[3])
		eInputStream.SendNewData(seq[4])
		aInputStream.SendNewData(0)
		sig := <-signal
		if max < sig {
			max = sig
			bestSequence = seq
		}
	}

	fmt.Println("Part Two")
	fmt.Printf("Max thruster signal: %d\n", max)
	fmt.Printf("Best sequence: %s\n", utils.IntArrayToString(bestSequence))
}

func runAmplifierTest(input []int, phaseSetting int, inputSignal int) int {
	inputStream := stream.New()
	computer := model.New(input, inputStream, nil)
	inputStream.SendNewData(phaseSetting)
	inputStream.SendNewData(inputSignal)
	output, err := computer.Process()
	if err != nil {
		fmt.Fprintln(os.Stderr, err.Error())
		os.Exit(1)
	}
	return output[0]
}
