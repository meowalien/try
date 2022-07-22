package plainB

type Space interface {
	Cap() int
	Len() int
}

type BufferRing interface {
	Space
	NewFile(needSpace int) *bufferRingFile
	FreeSpace(start BudderRingPointer, end BudderRingPointer)
	plusIndex(currentPointer BudderRingPointer, plusIndex int) BudderRingPointer
}

type bufferRing struct {
	spaceArea          []Space
	globalPointerStart BudderRingPointer
	globalPointerEnd   BudderRingPointer
}

func (b *bufferRing) FreeSpace(start BudderRingPointer, end BudderRingPointer) {
	freeStart := start.Layer(1)
	freeEnd := end.Layer(1)
	spaceDiffer := freeEnd - freeStart

	if spaceDiffer == 0 {
		// don't need to do wipe datas
		return
	} else if spaceDiffer >= 2 {
		st := freeStart + 1
		ed := freeEnd - 1
		b.cutOffSpace(st, ed)
	} else {
		st := freeStart + 1
		ed := len(b.spaceArea) - 1
		b.cutOffSpace(st, ed)
		b.cutOffSpace(0, freeEnd-1)
		//for i := freeStart + 1; i < len(b.spaceArea); i++ {
		//	b.takeOffSpace(i)
		//}
		//for i := 0; i < freeEnd; i++ {
		//	b.takeOffSpace(i)
		//}
	}
}

func (b *bufferRing) NewFile(needSpace int) *bufferRingFile {
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

func (b *bufferRing) cutOffSpace(st int, ed int) {
	//cutOffCount := ed - st + 1
	//newSpace := make([]Space, 0, len(b.spaceArea)-cutOffCount)
	panic("implement me")
}

// no scaleUP
func (b *bufferRing) occupySpace(space int) (pointerST BudderRingPointer, pointerED BudderRingPointer) {
	pointerST = b.plusIndex(b.globalPointerEnd, 1)
	pointerED = b.plusIndex(pointerST, space)
	b.globalPointerEnd = pointerED
	return
}

// no scaleUP
func (b bufferRing) plusIndex(oldPointer BudderRingPointer, plusIndex int) BudderRingPointer {
	newPointer := oldPointer.Copy()
	for {
		currentArea := b.spaceArea[newPointer.Layer(1)]
		remainingSpace := currentArea.Len() - newPointer.Layer(2)
		if remainingSpace >= plusIndex {
			newPointer.LayerSet(2, newPointer.Layer(2)+plusIndex)
			return newPointer
		}
		plusIndex -= remainingSpace
		if newPointer.Layer(1)+1 >= len(b.spaceArea) {
			newPointer.LayerSet(1, 0)
		} else {
			newPointer.LayerAdd(1, 1)
		}
	}
}

func (b *bufferRing) getByte(index int, index2 int) (byte, error) {
	panic("implement me")
}

func (b *bufferRing) setByte(index int, index2 int, b2 byte) error {
	panic("implement me")
}

const DefaultBufferRingInitializationSpace = 10

func NewBufferRing() BufferRing {
	return &bufferRing{
		spaceArea: []Space{NewSpace(DefaultBufferRingInitializationSpace)},
	}
}
