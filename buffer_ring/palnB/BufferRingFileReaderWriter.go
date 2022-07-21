package plainB

import (
	"github.com/meowalien/go-meowalien-lib/errs"
	"github.com/pkg/errors"
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
			if errors.Is(err1, io.EOF) {
				err = err1
			} else {
				err = errs.WithLine(err1)
			}
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

func (b bufferRingFileReaderWriter) copyRange(start BudderRingPointer, end BudderRingPointer, buf []byte) (n int, err error) {
	if cap(buf) == 0 {
		err = io.EOF
		return
	}
	b.foreach(start, end, func(layer1Index int, layer2Index int) bool {
		buf[n], err = b.theFile.theBufferRing.getByte(layer1Index, layer2Index)
		if err != nil {
			if !errors.Is(err, io.EOF) {
				err = errs.WithLine(err)
			}
			return false
		}
		if n == cap(buf)-1 {
			err = io.EOF
			return false
		}
		n++
		return true
	}
}

func (b *bufferRingFileReaderWriter) writeRange(cursor BudderRingPointer, end BudderRingPointer, buf []byte) (n int, err error) {

}

func (b *bufferRingFileReaderWriter) foreach(start BudderRingPointer, end BudderRingPointer, f func(layer1Index int, layer2Index int) bool) {
	
}
