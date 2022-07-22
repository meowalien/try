package planC

type Space interface {
	RemainingSpace() int
	RemainingSpaceAfter(index int) int
	Write(s []byte) (int, error)
	WriteToBuff(s []byte) (int, error)
	ReadToBuff(s []byte) (int, error)
	CleanUp() (empty bool)
	AreaID() bool
	// will not really clean up, just mark as empty
	CleanUpRange(index int, index2 int) (isEmpty bool)
	Free()
}

// use pool
func newSpace(space int) Space {

}
