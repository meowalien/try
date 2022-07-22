package planC

import (
	"io"
	"sync"
)

type Deletable interface {
	Delete()
}
type File interface {
	Deletable
	Reader() io.Reader
	Writer() io.Writer
	getCursorPair() CursorPair
}

type file[T any] struct {
	bufferRing BufferRing[T]
	cursorPair CursorPair
	fileReader io.Reader
	fileWriter io.Writer
	lock       sync.RWMutex
}

func (f *file[T]) getCursorPair() CursorPair {
	return f.cursorPair
}

func (f *file[T]) Delete() {
	f.bufferRing.DeleteFile(f)
}

// todo: if Reader been call first , make it as chan and replace fileWriter
func (f *file[T]) Reader() io.Reader {
	if f.fileReader != nil {
		return f.fileReader
	}
	f.fileReader = &fileReadWriter[T]{file: f, cursor: f.cursorPair.GetStartCursor()}
	return f.fileReader
}

func (f *file[T]) Writer() io.Writer {
	if f.fileWriter != nil {
		return f.fileWriter
	}
	f.fileWriter = &fileReadWriter[T]{file: f, cursor: f.cursorPair.GetStartCursor()}
	return f.fileWriter
}
