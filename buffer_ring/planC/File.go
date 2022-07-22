package planC

import "io"

type Deletable interface {
	Delete()
}
type File interface {
	Deletable
	NewReadWriter() io.ReadWriter
	getCursorPair() CursorPair
}
type file struct {
	bufferRing *bufferRing
	cursorPair CursorPair
}

func (f *file) getCursorPair() CursorPair {
	return f.cursorPair
}

func (f *file) Delete() {
	f.bufferRing.DeleteFile(f)
}

func (f *file) NewReadWriter() io.ReadWriter {
	return &fileReadWriter{f: f}
}
