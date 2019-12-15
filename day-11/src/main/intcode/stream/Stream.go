package stream

import (
	"time"
)

type Stream struct {
	name   string
	data   []int
	IsOpen bool
}

func New(name string) *Stream {
	return &Stream{name, []int{}, true}
}

func (stream *Stream) SendNewData(value int) {
	stream.data = append(stream.data, value)
	// fmt.Printf("%s - Added %d\n", stream.name, value)
}

func (stream *Stream) WaitForNewData(newData chan int, lastValueIndex int) {
	go func() {
		currentDataLength := len(stream.data)
		// fmt.Printf("%s - current length %d\n", stream.name, currentDataLength)
		for currentDataLength <= (lastValueIndex + 1) {
			time.Sleep(time.Microsecond)
			currentDataLength = len(stream.data)
			// fmt.Printf("%s - current length %d and lastValueIndex %d\n", stream.name, currentDataLength, lastValueIndex)
		}
		newData <- stream.data[lastValueIndex+1]
		// fmt.Printf("%s - Read %d\n", stream.name, stream.data[lastValueIndex+1])
	}()
}

func (stream *Stream) Close() {
	stream.IsOpen = false
}
