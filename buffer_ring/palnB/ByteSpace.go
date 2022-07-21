package plainB

type byteSpace []byte

func (b byteSpace) Cap() int {
	//TODO implement me
	panic("implement me")
}

func (b byteSpace) Len() int {
	//TODO implement me
	panic("implement me")
}

func NewSpace(space int) Space {
	return byteSpace(make([]byte, 0, space))
}
