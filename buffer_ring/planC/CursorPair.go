package planC

type Cursor interface {
	Plus(space int) Cursor
	AreaID() int
	SpaceIndex() int
}

type cursor[T any] struct {
	spaceRing  SpaceRing[T]
	areaIndex  int
	spaceIndex int
}

func (c cursor[T]) AreaID() int {
	return c.areaIndex
}

func (c cursor[T]) SpaceIndex() int {
	return c.spaceIndex
}

func (c cursor[T]) Plus(spaceNeed int) Cursor {
	remainingSpace := c.spaceRing.getSpace(c.areaIndex).RemainingSpaceAfter(c.spaceIndex)
	spaceNeed -= remainingSpace
	if spaceNeed <= 0 {
		return newCursor(c.spaceRing, c.areaIndex, c.spaceIndex+spaceNeed+remainingSpace)
	} else {
		c.areaIndex = c.spaceRing.nextArea(c.areaIndex)
		c.spaceIndex = 0
		return c.Plus(spaceNeed)
	}
}

func newCursor[T any](c SpaceRing[T], areaIndex int, spaceIndex int) Cursor {
	return &cursor[T]{spaceRing: c, areaIndex: areaIndex, spaceIndex: spaceIndex}
}

type CursorPair interface {
	GetStartCursor() Cursor
	GetEndCursor() Cursor
	SetStartCursor(cursor Cursor)
	SetEndCursor(cursor Cursor)
}

type cursorPair[T any] struct {
	ring        SpaceRing[T]
	startCursor Cursor
	endCursor   Cursor
}

func (c *cursorPair[T]) GetStartCursor() Cursor {
	return c.startCursor
}

func (c *cursorPair[T]) GetEndCursor() Cursor {
	return c.endCursor
}

func (c *cursorPair[T]) SetStartCursor(cursor Cursor) {
	c.startCursor = cursor
}

func (c *cursorPair[T]) SetEndCursor(cursor Cursor) {
	c.endCursor = cursor
}

func newCursorPair[T any](ring SpaceRing[T]) CursorPair {
	c := &cursorPair[T]{ring: ring}
	c.startCursor = newCursor(ring, 0, 0)
	c.endCursor = newCursor(ring, 0, 0)
	return c
}
