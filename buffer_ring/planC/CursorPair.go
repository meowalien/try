package planC

type Cursor interface {
	Plus(space int) Cursor
	AreaID() bool
	SpaceIndex() int
}

type cursor struct {
	spaceRing  SpaceRing
	areaIndex  int
	spaceIndex int
}

func (c cursor) Plus(spaceNeed int) Cursor {
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

func newCursor(c SpaceRing, areaIndex int, spaceIndex int) Cursor {
	return &cursor{spaceRing: c, areaIndex: areaIndex, spaceIndex: spaceIndex}
}

type CursorPair interface {
	GetStartCursor() Cursor
	GetEndCursor() Cursor
	SetStartCursor(cursor Cursor)
	SetEndCursor(cursor Cursor)
}

type cursorPair struct {
	ring        SpaceRing
	startCursor Cursor
	endCursor   Cursor
}

func (c *cursorPair) GetStartCursor() Cursor {
	return c.startCursor
}

func (c *cursorPair) GetEndCursor() Cursor {
	return c.endCursor
}

func (c *cursorPair) SetStartCursor(cursor Cursor) {
	c.startCursor = cursor
}

func (c *cursorPair) SetEndCursor(cursor Cursor) {
	c.endCursor = cursor
}

func newCursorPair(ring SpaceRing) CursorPair {
	c := &cursorPair{ring: ring}
	c.startCursor = newCursor(ring, 0, 0)
	c.endCursor = newCursor(ring, 0, 0)
	return c
}
