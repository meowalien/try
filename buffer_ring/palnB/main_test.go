package plainB

import (
	"bytes"
	"fmt"
	"io"
	"testing"
)

func TestBufferRing(t *testing.T) {
	br := NewBufferRing()
	fileA := br.NewFile(10)

	dataA := []byte("Hello")
	n, err := fileA.Write(dataA)
	if err != nil {
		panic(err)
	}
	fmt.Println("Write n: ", n)

	var dataB []byte
	n1, err := io.Copy(bytes.NewBuffer(dataB), fileA)
	if err != nil {
		panic(err)
	}
	fmt.Println("Read n: ", n1)

	fmt.Println("dataB: ", string(dataB))
}
