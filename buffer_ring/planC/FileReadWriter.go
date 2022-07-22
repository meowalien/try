package planC

import (
	"github.com/meowalien/go-meowalien-lib/errs"
	"github.com/pkg/errors"
	"io"
)

type fileReadWriter struct {
	*file
	cursor Cursor
}

func (f *fileReadWriter) Read(p []byte) (n int, err error) {
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

func (f *fileReadWriter) Write(p []byte) (n int, err error) {
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
