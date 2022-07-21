package plainB

import (
	"github.com/meowalien/go-meowalien-lib/errs"
	"io"
)

func newBufferRingFileReaderWriter(bf *bufferRingFile) io.ReadWriter {
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
		if b.outOfRange(b.theFile.pointerStart, b.theFile.pointerEnd, b.cursor) {
			return done, io.EOF
		}
		doneN, err1 := b.copyRange(b.cursor, b.theFile.pointerEnd, buf)
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

func (b *bufferRingFileReaderWriter) Write(buf []byte) (done int, err error) {
	total := len(buf)
	for {
		if b.outOfRange(b.theFile.pointerStart, b.theFile.pointerEnd, b.cursor) {
			return done, io.EOF
		}
		doneN, err1 := b.writeRange(b.cursor, b.theFile.pointerEnd, buf)
		if err1 != nil {
			err = errs.WithLine(err1)
			return
		}
		done += doneN
		b.cursor = b.theFile.theBufferRing.plusIndex(b.cursor, doneN+1)
		if total == done {
			return total, io.EOF
		}
	}

}

func (b bufferRingFileReaderWriter) outOfRange(start BudderRingPointer, end BudderRingPointer, cursor BudderRingPointer) bool {
	if start.Layer(1) <= end.Layer(1) {
		return cursor.Layer(1) > end.Layer(1) || cursor.Layer(1) < start.Layer(1) || (cursor.Layer(2) > end.Layer(2) || cursor.Layer(2) < start.Layer(2))
	} else {
		return (end.Layer(1) < cursor.Layer(1) && start.Layer(1) > cursor.Layer(1)) || (cursor.Layer(2) > end.Layer(2) || cursor.Layer(2) < start.Layer(2))
	}
}

func (b bufferRingFileReaderWriter) copyRange(cursor BudderRingPointer, end BudderRingPointer, buf []byte) (n int, err error) {
	b.foreach(cursor, end, func(layer1Index int , layer2Index int ) (bool) {
		if n >= cap(buf){

		}
	}
}

func (b *bufferRingFileReaderWriter) writeRange(cursor BudderRingPointer, end BudderRingPointer, buf []byte) (n int, err error) {

}
