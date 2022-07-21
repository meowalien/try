package plainB

import "io"

type Abandonedable interface {
	Abandoned()
}

type BufferRingFile interface {
	io.ReadWriter
	Abandonedable
}

type bufferRingFile struct {
	theBufferRing BufferRing
	cursor        int64
	pointerStart  BudderRingPointer
	pointerEnd    BudderRingPointer
}

func (b bufferRingFile) Abandoned() {
	//TODO implement me
	panic("implement me")
}

func (b bufferRingFile) Read(p []byte) (n int, err error) {
	//TODO implement me
	panic("implement me")
}

func (b bufferRingFile) Write(p []byte) (n int, err error) {
	//TODO implement me
	panic("implement me")
}
