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

type file struct {
	bufferRing BufferRing
	cursorPair CursorPair
	fileReader io.Reader
	fileWriter io.Writer
	lock       sync.RWMutex
}

func (f *file) getCursorPair() CursorPair {
	return f.cursorPair
}

func (f *file) Delete() {
	f.bufferRing.DeleteFile(f)
}

// todo: if Reader been call first , make it as chan and replace fileWriter
func (f *file) Reader() io.Reader {
	if f.fileReader != nil {
		return f.fileReader
	}
	f.fileReader = &fileReadWriter{file: f, cursor: f.cursorPair.GetStartCursor()}
	return f.fileReader
}

func (f *file) Writer() io.Writer {
	if f.fileWriter != nil {
		return f.fileWriter
	}
	f.fileWriter = &fileReadWriter{file: f, cursor: f.cursorPair.GetStartCursor()}
	return f.fileWriter
}
