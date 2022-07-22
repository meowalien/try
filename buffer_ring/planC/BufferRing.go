package planC

import (
	"sync"
)

type BufferRing interface {
	NewFile(i int) File
	DeleteFile(f File)
}

type bufferRing struct {
	spaceRing        SpaceRing
	globalCursorPair CursorPair
	spaceRingLock    sync.RWMutex
}

// lock
func (b *bufferRing) NewFile(needSpace int) File {
	b.spaceRingLock.Lock()
	defer b.spaceRingLock.Unlock()

	if b.spaceRing.RemainingSpace() < needSpace {
		b.scaleUpRing(b.calculateNewScaleSizeToBeScaleUP(needSpace))
	}
	cursorPair := b.occupySpace(needSpace)
	return &file{cursorPair: cursorPair, bufferRing: b}
}

// lock
func (b *bufferRing) DeleteFile(f File) {
	fileCursorPair := f.getCursorPair()
	b.spaceRing.cleanUpSpaceInRange(fileCursorPair)
	if b.globalCursorPair.GetStartCursor() == fileCursorPair.GetStartCursor() {
		b.spaceRingLock.Lock()
		b.moveStartCursorToNotEmptySpace()
		b.spaceRingLock.Unlock()
	}
}

// no lock
// occupySpace will move the current end cursor to "currentEnd + needSpace"
// and return the CursorPair of oldEndCursor ~ newEndCursor
func (b *bufferRing) occupySpace(space int) (cursorPair CursorPair) {
	currentEnd := b.globalCursorPair.GetEndCursor()
	newEnd := currentEnd.Plus(space)
	b.globalCursorPair.SetEndCursor(newEnd)

	cursorPair = newCursorPair(b.spaceRing)
	cursorPair.SetStartCursor(currentEnd)
	cursorPair.SetEndCursor(newEnd)
	return
}

// no lock
func (b *bufferRing) scaleUpRing(needSpace int) {
	b.spaceRing.insertSpaceBeforeStart(newSpace(needSpace))
}

func (b *bufferRing) calculateNewScaleSizeToBeScaleUP(space int) int {
	// todo: not implemented
	return space
}

// no lock
func (b *bufferRing) moveStartCursorToNotEmptySpace() {
	notEmptySpaceCursor := b.spaceRing.findNotEmptySpace(b.globalCursorPair.GetStartCursor())
	b.globalCursorPair.SetStartCursor(notEmptySpaceCursor)
}

const DefaultSpaceSize int = 10

func NewBufferRing() BufferRing {
	spaceRing := newSpaceRing(newSpace(DefaultSpaceSize))
	cursorPair := newCursorPair(spaceRing)
	return &bufferRing{
		spaceRing:        spaceRing,
		globalCursorPair: cursorPair,
	}
}
