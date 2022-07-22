package planC

type Space[T any] interface {
	RemainingSpace() int
	RemainingSpaceAfter(index int) int
	Write(s []T) (int, error)
	WriteToBuff(s []T) (int, error)
	ReadToBuff(s []T) (int, error)
	CleanUp() (empty bool)
	AreaID() int
	// will not really clean up, just mark as empty
	CleanUpRange(index int, index2 int) (isEmpty bool)
	// mark as recyclable
	Free()
	IsFree() bool
}

type space[T any] struct {
	theSpace []T
}

func (s space[T]) RemainingSpace() int {
	//TODO implement me
	panic("implement me")
}

func (s space[T]) RemainingSpaceAfter(index int) int {
	//TODO implement me
	panic("implement me")
}

func (s space[T]) Write(buff []T) (int, error) {
	//TODO implement me
	panic("implement me")
}

func (s space[T]) WriteToBuff(buff []T) (int, error) {
	//TODO implement me
	panic("implement me")
}

func (s space[T]) ReadToBuff(buff []T) (int, error) {
	//TODO implement me
	panic("implement me")
}

func (s space[T]) CleanUp() (empty bool) {
	//TODO implement me
	panic("implement me")
}

func (s space[T]) AreaID() int {
	//TODO implement me
	panic("implement me")
}

func (s space[T]) CleanUpRange(index int, index2 int) (isEmpty bool) {
	//TODO implement me
	panic("implement me")
}

func (s space[T]) Free() {
	//TODO implement me
	panic("implement me")
}

func (s space[T]) IsFree() bool {
	//TODO implement me
	panic("implement me")
}

// use pool
func newSpace[T any](s int) Space[T] {
	return &space[T]{}
}
