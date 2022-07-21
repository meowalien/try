package plainB

import "io"

type BufferRingFile interface {
	Abandonedable
	NewReaderWriter() io.ReadWriter
}

type Abandonedable interface {
	Abandoned()
}

type bufferRingFile struct {
	theBufferRing BufferRing
	pointerStart  BudderRingPointer
	pointerEnd    BudderRingPointer
}

func (b *bufferRingFile) Abandoned() {
	b.theBufferRing.FreeSpace(b.pointerStart, b.pointerEnd)
}

func (b *bufferRingFile) NewReaderWriter() io.ReadWriter {
	return newBufferRingFileReaderWriter(b)
}
