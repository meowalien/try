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
	fileAReaderWriter := fileA.NewReaderWriter()
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

//type SSS[T int | string] []T
//
//type MyStruct[T IIIII] struct {
//	Name T
//}
//
//type IIIII interface {
//	int | string
//	FFF()
//}
//
//func Add[T ~int | ~float32](a, b T) T {
//	return a + b
//}
