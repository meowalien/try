package planC

type Cursor interface {
	Plus(space int) Cursor
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
		//c.areaIndex++ /////
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
	SetEndCursor(push Cursor)
}

type cursorPair struct {
	ring        SpaceRing
	startCursor Cursor
	endCursor   Cursor
}

func (c cursorPair) GetStartCursor() Cursor {
	//TODO implement me
	panic("implement me")
}

func (c cursorPair) GetEndCursor() Cursor {
	//TODO implement me
	panic("implement me")
}

func (c cursorPair) SetStartCursor(cursor Cursor) {
	//TODO implement me
	panic("implement me")
}

func (c cursorPair) SetEndCursor(push Cursor) {
	//TODO implement me
	panic("implement me")
}

func newCursorPair(ring SpaceRing) CursorPair {
	c := &cursorPair{ring: ring}
	c.startCursor = newCursor(ring, 0, 0)
	c.endCursor = newCursor(ring, 0, 0)
	return c
}
