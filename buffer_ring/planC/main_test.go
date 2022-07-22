package planC

import (
	"bytes"
	"fmt"
	"io"
	"testing"
)

func TestRingBuffer(t *testing.T) {
	br := NewBufferRing[byte]()
	fileA := br.NewFile(10)

	dataA := []byte("Hello")
	fileAReaderWriter := fileA.Writer()
	n, err := fileAReaderWriter.Write(dataA)
	if err != nil {
		panic(err)
	}
	fmt.Println("Write n: ", n)

	var dataB []byte
	n1, err := io.Copy(bytes.NewBuffer(dataB), fileA.Reader())
	if err != nil {
		panic(err)
	}
	fmt.Println("Read n: ", n1)

	fmt.Println("dataB: ", string(dataB))

}
