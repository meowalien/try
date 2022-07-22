package planC

import (
	"sync"
)

type Deletable interface {
	Delete()
}
type Reader[T any] interface {
	Read(p []T) (n int, err error)
}
type Writer[T any] interface {
	Write(p []T) (n int, err error)
}
type File[T any] interface {
	Deletable
	Reader() Reader[T]
	Writer() Writer[T]
	getCursorPair() CursorPair
}

type file[T any] struct {
	bufferRing BufferRing[T]
	cursorPair CursorPair
	fileReader Reader[T]
	fileWriter Writer[T]
	lock       sync.RWMutex
}

func (f *file[T]) getCursorPair() CursorPair {
	return f.cursorPair
}

func (f *file[T]) Delete() {
	f.bufferRing.DeleteFile(f)
}

// todo: if Reader been call first , make it as chan and replace fileWriter
func (f *file[T]) Reader() Reader[T] {
	if f.fileReader != nil {
		return f.fileReader
	}
	f.fileReader = &fileReadWriter[T]{file: f, cursor: f.cursorPair.GetStartCursor()}
	return f.fileReader
}

func (f *file[T]) Writer() Writer[T] {
	if f.fileWriter != nil {
		return f.fileWriter
	}
	f.fileWriter = &fileReadWriter[T]{file: f, cursor: f.cursorPair.GetStartCursor()}
	return f.fileWriter
}
