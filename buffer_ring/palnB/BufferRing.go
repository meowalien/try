package plainB

type Space interface {
	Cap() int
	Len() int
}

type BufferRing interface {
	Space
	NewFile(needSpace int) BufferRingFile
}

type bufferRing struct {
	spaceArea          []Space
	globalPointerStart BudderRingPointer
	globalPointerEnd   BudderRingPointer
}

func (b *bufferRing) NewFile(needSpace int) BufferRingFile {
	if b.Len() < needSpace {
		b.scaleUP(b.calculateNewScaleNeedToBeScaleUP(needSpace))
	}
	occupiedPointerStart, occupiedPointerEnd := b.occupySpace(needSpace)
	return &bufferRingFile{
		pointerStart: occupiedPointerStart,
		pointerEnd:   occupiedPointerEnd,
	}
}

func (b *bufferRing) Len() (ans int) {
	if b.globalPointerStart.Layer(1) <= b.globalPointerEnd.Layer(1) {
		for i := b.globalPointerStart.Layer(1); i <= b.globalPointerEnd.Layer(1); i++ {
			ans += b.spaceArea[i].Len()
		}
	} else {
		for i := b.globalPointerStart.Layer(1); i < len(b.spaceArea); i++ {
			ans += b.spaceArea[i].Len()
		}
		for i := 0; i <= b.globalPointerEnd.Layer(1); i++ {
			ans += b.spaceArea[i].Len()
		}
	}
	return
}
func (b *bufferRing) Cap() (ans int) {
	for i := range b.spaceArea {
		ans += b.spaceArea[i].Cap()
	}
	return
}

func (b *bufferRing) calculateNewScaleNeedToBeScaleUP(space int) int {
	return DefaultBufferRingInitializationSpace
}

func (b *bufferRing) scaleUP(need int) {
	// todo: take from pool
	newSpace := make([]Space, 0, len(b.spaceArea)+need)
	for i := 0; i <= b.globalPointerStart.Layer(1); i++ {
		newSpace[i] = b.spaceArea[i]
	}
	newSpace[b.globalPointerStart.Layer(1)+1] = NewSpace(need)
	for i := b.globalPointerStart.Layer(1) + 1; i <= len(b.spaceArea); i++ {
		newSpace[i+1] = b.spaceArea[i]
	}
	// todo: reuse the old b.spaceArea
	b.spaceArea = newSpace
}

// no scaleUP
func (b *bufferRing) occupySpace(space int) (pointerST BudderRingPointer, pointerED BudderRingPointer) {
	pointerST = b.plusIndex(b.globalPointerEnd, 1)
	pointerED = b.plusIndex(pointerST, space)
	b.globalPointerEnd = pointerED
	return
}

// no scaleUP
func (b bufferRing) plusIndex(currentPointer BudderRingPointer, plusIndex int) BudderRingPointer {
	for {
		currentArea := b.spaceArea[currentPointer.Layer(1)]
		remainingSpace := currentArea.Cap() - currentPointer.Layer(2)
		if remainingSpace >= plusIndex {
			currentPointer.LayerSet(2, currentPointer.Layer(2)+plusIndex)
			return currentPointer
		}
		plusIndex -= remainingSpace
		if currentPointer.Layer(1)+1 >= len(b.spaceArea) {
			currentPointer.LayerSet(1, 0)
		} else {
			currentPointer.LayerAdd(1, 1)
		}
	}
}

const DefaultBufferRingInitializationSpace = 10

func NewBufferRing() BufferRing {
	return &bufferRing{
		spaceArea: []Space{NewSpace(DefaultBufferRingInitializationSpace)},
	}
}
