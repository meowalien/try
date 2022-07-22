package planC

type SpaceRing[T any] interface {
	insertSpaceBeforeCursor(cursor Cursor, space Space[T])
	TotalRemainingSpace() int
	cleanUpSpaceInRange(pair CursorPair)
	findNotEmptySpaceAfter(i Cursor) Cursor
	getSpace(spaceID int) Space[T]
	nextArea(spaceID int) int
	forRangeSpace(writeCursor Cursor, nextCursor Cursor, f func(space Space[T], isEnd bool) bool)
}

func newSpaceRing[T any](initialiseSpace ...Space[T]) SpaceRing[T] {
	return newSpaceLinkList(initialiseSpace...)
}
