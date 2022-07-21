package plainB

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
