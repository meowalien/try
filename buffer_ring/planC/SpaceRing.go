package planC

type Space interface {
	RemainingSpace() int
	RemainingSpaceAfter(index int) int
	Write(s []byte) (int, error)
	WriteToBuff(s []byte) (int, error)
	ReadToBuff(s []byte) (int, error)
}

// use pool
func newSpace(space int) Space {

}

type SpaceRing []Space

func (r SpaceRing) insertSpaceBeforeStart(space Space) {

}

func (r SpaceRing) RemainingSpace() int {

}

func (r SpaceRing) cleanUpSpaceInRange(pair CursorPair) {

}

func (r SpaceRing) findNotEmptySpace(i interface{}) Cursor {

}

// get Space at input index of area
func (r SpaceRing) getSpace(index int) Space {

}

func (r SpaceRing) nextArea(index int) int {

}

func (r SpaceRing) forRangeSpace(writeCursor Cursor, nextCursor Cursor, f func(space Space, isEnd bool) bool) {

}

func newSpaceRing(initialiseSpace ...Space) SpaceRing {

}
