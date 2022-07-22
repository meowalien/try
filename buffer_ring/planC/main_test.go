package planC

import (
	"bytes"
	"fmt"
	"io"
	"testing"
)

func TestRingBuffer(t *testing.T) {
	br := NewBufferRing()
	fileA := br.NewFile(10)

	dataA := []byte("Hello")
	fileAReaderWriter := fileA.NewReadWriter()
	n, err := fileAReaderWriter.Write(dataA)
	if err != nil {
		panic(err)
	}
	fmt.Println("Write n: ", n)

	var dataB []byte
	n1, err := io.Copy(bytes.NewBuffer(dataB), fileAReaderWriter)
	if err != nil {
		panic(err)
	}
	fmt.Println("Read n: ", n1)

	fmt.Println("dataB: ", string(dataB))

}
