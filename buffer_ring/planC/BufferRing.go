package planC

import (
	"github.com/meowalien/go-meowalien-lib/errs"
	"github.com/pkg/errors"
	"io"
	"sync"
)

type BufferRing interface {
	NewFile(i int) File
	DeleteFile(f File)
	readBufferFrom(startAt Cursor, p []byte) (n int, nextCursor Cursor, err error)
	writeBufferTo(startAt Cursor, p []byte) (n int, nextCursor Cursor, err error)
}

type bufferRing struct {
	spaceRing        SpaceRing
	globalCursorPair CursorPair
	spaceRingLock    sync.RWMutex
}

func (b *bufferRing) writeBufferTo(startAt Cursor, p []byte) (n int, nextCursor Cursor, err error) {
	total := len(p)
	b.spaceRing.forRangeSpace(startAt, nextCursor, func(space Space, isEnd bool) bool {
		var s []byte
		if isEnd {
			s = p[:total+1]
		} else {
			s = p
		}
		var written int
		written, err = space.WriteToBuff(s)
		total -= written
		if err != nil {
			if errors.Is(err, io.EOF) {
				return false
			}
			err = errs.WithLine(err)
			return false
		}
		return false
	})
	nextCursor = startAt.Plus(total + 1)
	n = len(p) - total
	return
}

func (b *bufferRing) readBufferFrom(startAt Cursor, p []byte) (n int, nextCursor Cursor, err error) {
	total := len(p)
	b.spaceRing.forRangeSpace(startAt, nextCursor, func(space Space, isEnd bool) bool {
		var s []byte
		if isEnd {
			s = p[:total+1]
		} else {
			s = p
		}
		var read int
		read, err = space.ReadToBuff(s)
		total -= read
		if err != nil {
			if errors.Is(err, io.EOF) {
				return false
			}
			err = errs.WithLine(err)
			return false
		}
		return false
	})
	nextCursor = startAt.Plus(total + 1)
	n = len(p) - total
	return
}

// lock
func (b *bufferRing) NewFile(needSpace int) File {
	b.spaceRingLock.Lock()
	defer b.spaceRingLock.Unlock()

	if b.spaceRing.RemainingSpace() < needSpace {
		b.scaleUpRing(b.calculateNewScaleSizeToBeScaleUP(needSpace))
	}
	pair := b.occupySpace(needSpace)
	return &file{cursorPair: pair, bufferRing: b}
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
	pair := newCursorPair(spaceRing)
	return &bufferRing{
		spaceRing:        spaceRing,
		globalCursorPair: pair,
	}
}
