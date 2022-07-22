package planC

import (
	"github.com/meowalien/go-meowalien-lib/errs"
	"github.com/pkg/errors"
	"io"
	"sync"
)

type BufferRing[T any] interface {
	NewFile(i int) File[T]
	DeleteFile(f File[T])
	readBufferFrom(startAt Cursor, p []T) (n int, nextCursor Cursor, err error)
	writeBufferTo(startAt Cursor, p []T) (n int, nextCursor Cursor, err error)
}

type bufferRing[T any] struct {
	spaceRing        SpaceRing[T]
	globalCursorPair CursorPair
	spaceRingLock    sync.RWMutex
}

func (b *bufferRing[T]) writeBufferTo(startAt Cursor, p []T) (n int, nextCursor Cursor, err error) {
	total := len(p)
	b.spaceRing.forRangeSpace(startAt, nextCursor, func(space Space[T], isEnd bool) bool {
		var s []T
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

func (b *bufferRing[T]) readBufferFrom(startAt Cursor, p []T) (n int, nextCursor Cursor, err error) {
	total := len(p)
	b.spaceRing.forRangeSpace(startAt, nextCursor, func(space Space[T], isEnd bool) bool {
		var s []T
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
func (b *bufferRing[T]) NewFile(needSpace int) File[T] {
	b.spaceRingLock.Lock()
	defer b.spaceRingLock.Unlock()

	if b.spaceRing.TotalRemainingSpace() < needSpace {
		b.scaleUpRing(b.calculateNewScaleSizeToBeScaleUP(needSpace))
	}
	pair := b.occupySpace(needSpace)
	return &file[T]{cursorPair: pair, bufferRing: b}
}

// lock
func (b *bufferRing[T]) DeleteFile(f File[T]) {
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
func (b *bufferRing[T]) occupySpace(space int) (cursorPair CursorPair) {
	currentEnd := b.globalCursorPair.GetEndCursor()
	newEnd := currentEnd.Plus(space)
	b.globalCursorPair.SetEndCursor(newEnd)

	cursorPair = newCursorPair(b.spaceRing)
	cursorPair.SetStartCursor(currentEnd)
	cursorPair.SetEndCursor(newEnd)
	return
}

// no lock
func (b *bufferRing[T]) scaleUpRing(needSpace int) {
	b.spaceRing.insertSpaceBeforeCursor(b.globalCursorPair.GetStartCursor(), newSpace[T](needSpace))
}

func (b *bufferRing[T]) calculateNewScaleSizeToBeScaleUP(space int) int {
	// todo: not implemented
	return space
}

// no lock
func (b *bufferRing[T]) moveStartCursorToNotEmptySpace() {
	notEmptySpaceCursor := b.spaceRing.findNotEmptySpaceAfter(b.globalCursorPair.GetStartCursor())
	b.globalCursorPair.SetStartCursor(notEmptySpaceCursor)
}

const DefaultSpaceSize int = 10

func NewBufferRing[T any]() BufferRing[T] {
	spaceRing := newSpaceRing[T](newSpace[T](DefaultSpaceSize))
	pair := newCursorPair(spaceRing)
	return &bufferRing[T]{
		spaceRing:        spaceRing,
		globalCursorPair: pair,
	}
}
