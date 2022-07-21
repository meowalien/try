package plainB

import (
	"github.com/meowalien/go-meowalien-lib/errs"
	"io"
)

func NewBufferRingFileReaderWriter(bf *bufferRingFile) io.Reader {
	return &bufferRingFileReaderWriter{theFile: bf}
}

type bufferRingFileReaderWriter struct {
	theFile *bufferRingFile
	cursor  BudderRingPointer
}

func (b *bufferRingFileReaderWriter) ResetCursor() {
	b.cursor = b.theFile.pointerStart
}

func (b *bufferRingFileReaderWriter) Read(buf []byte) (done int, err error) {
	total := len(buf)
	for {
		if b.theFile.outOfRange(b.theFile.pointerStart, b.theFile.pointerEnd, b.cursor) {
			return done, io.EOF
		}
		doneN, err1 := b.theFile.copyRange(b.cursor, b.theFile.pointerEnd, buf)
		if err1 != nil {
			err = errs.WithLine(err1)
			return
		}
		done += doneN
		b.theFile.theBufferRing.plusIndex(b.cursor, doneN+1)
		if total == done {
			return total, io.EOF
		}
	}
}

func (b *bufferRingFileReaderWriter) Write(p []byte) (n int, err error) {
	//TODO implement me
	panic("implement me")
}

func (b *bufferRingFileReaderWriter) outOfRange(start BudderRingPointer, end BudderRingPointer, cursor BudderRingPointer) bool {

}

func (b *bufferRingFileReaderWriter) copyRange(cursor BudderRingPointer, end BudderRingPointer, buf []byte) (n int, err error) {

}
