package planC

type SpaceRing interface {
	insertSpaceBeforeCursor(cursor Cursor, space Space)
	TotalRemainingSpace() int
	cleanUpSpaceInRange(pair CursorPair)
	findNotEmptySpaceAfter(i Cursor) Cursor
	getSpace(spaceID int) Space
	nextArea(spaceID int) int
	forRangeSpace(writeCursor Cursor, nextCursor Cursor, f func(space Space, isEnd bool) bool)
}

func newSpaceRing(initialiseSpace ...Space) SpaceRing {
	return newSpaceLinkList(initialiseSpace...)
}
