package plainB

import "io"

type BufferRingFile interface {
	Abandonedable
	NewReadWriter() io.ReadWriter
}

type Abandonedable interface {
	Delete()
}

type bufferRingFile struct {
	theBufferRing *bufferRing
	pointerStart  BudderRingPointer
	pointerEnd    BudderRingPointer
}

func (b *bufferRingFile) Delete() {
	b.theBufferRing.FreeSpace(b.pointerStart, b.pointerEnd)
}

func (b *bufferRingFile) NewReadWriter() io.ReadWriter {
	return newBufferRingFileReaderWriter(b)
}
