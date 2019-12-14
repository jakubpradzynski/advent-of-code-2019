package stream

import "time"

type Stream struct {
	data []int
}

func New() *Stream {
	return &Stream{[]int{}}
}

func (stream *Stream) SendNewData(value int) {
	stream.data = append(stream.data, value)
}

func (stream *Stream) WaitForNewData(newData chan int, lastValueIndex int) {
	go func() {
		currentDataLength := len(stream.data)
		for currentDataLength <= (lastValueIndex + 1) {
			time.Sleep(time.Microsecond)
			currentDataLength = len(stream.data)
		}
		newData <- stream.data[lastValueIndex+1]
	}()
}
