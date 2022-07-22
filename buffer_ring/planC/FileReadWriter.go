package planC

import (
	"github.com/meowalien/go-meowalien-lib/errs"
	"github.com/pkg/errors"
	"io"
)

type fileReadWriter[T any] struct {
	*file[T]
	cursor Cursor
}

func (f *fileReadWriter[T]) Read(p []byte) (n int, err error) {
	f.lock.RLock()
	defer f.lock.RUnlock()
	n, nextCursor, err := f.bufferRing.readBufferFrom(f.cursor, p)
	f.cursor = nextCursor
	if err != nil {
		if errors.Is(err, io.EOF) {
			return
		}
		err = errs.WithLine(err)
		return
	}
	return
}

func (f *fileReadWriter[T]) Write(p []byte) (n int, err error) {
	f.lock.Lock()
	defer f.lock.Unlock()
	n, nextCursor, err := f.bufferRing.writeBufferTo(f.cursor, p)
	f.cursor = nextCursor
	if err != nil {
		if errors.Is(err, io.EOF) {
			return
		}
		err = errs.WithLine(err)
		return
	}
	return
}
