package planC

import "container/list"

// todo: change firstElement if better

type LinkListRing interface {
	SpaceRing
}

func newSpaceLinkList(space ...Space) LinkListRing {
	if len(space) == 0 {
		panic("spaceLinkList must have at least one space")
	}
	l := list.New()
	firstElement := l.PushBack(space[0])
	for i := 1; i < len(space); i++ {
		l.PushBack(space[0])
	}

	return &spaceLinkList{rawList: l, firstElement: firstElement}
}

type spaceLinkList struct {
	rawList      *list.List
	firstElement *list.Element
}

func (r *spaceLinkList) insertSpaceBeforeCursor(cursor Cursor, space Space) {
	element, ok := r.findElementByCursor(r.firstElement, cursor)
	if !ok {
		panic("cursor not found")
	}
	r.rawList.InsertBefore(space, element)
}

func (r *spaceLinkList) findElementByCursor(startElement *list.Element, c Cursor) (elementFound *list.Element, ok bool) {
	ok = r.foreachElementStartAt(startElement, func(element *list.Element) bool {
		if elementFound.Value.(Space).AreaID() == c.AreaID() {
			elementFound = element
			return false
		}
		return true
	})
	return
}

func (r *spaceLinkList) TotalRemainingSpace() (total int) {
	_ = r.foreachElementStartAt(r.firstElement, func(element *list.Element) bool {
		total += element.Value.(Space).RemainingSpace()
		return true
	})
	return
}

func (r *spaceLinkList) cleanUpSpaceInRange(pair CursorPair) {
	cursorStart := pair.GetStartCursor()
	cursorEnd := pair.GetStartCursor()
	elementStart, ok := r.findElementByCursor(r.firstElement, cursorStart)
	if !ok {
		panic("cursor not found")
	}

	if cursorStart.AreaID() == cursorEnd.AreaID() {
		theSpace := elementStart.Value.(Space)
		isEmpty := theSpace.CleanUpRange(cursorStart.SpaceIndex(), cursorEnd.SpaceIndex())
		if isEmpty {
			r.rawList.Remove(elementStart)
			theSpace.Free()
		}
	} else {
		var elementEnd *list.Element
		elementEnd, ok = r.findElementByCursor(elementStart, cursorStart)
		if !ok {
			panic("cursor not found")
		}
		_ = r.foreachElementStartAt(elementStart, func(element *list.Element) (next bool) {
			if element == elementEnd {
				return
			}
			theSpace := element.Value.(Space)
			var cleanUpStart int
			var cleanUpEnd int
			next = true
			if element == elementStart {
				cleanUpStart = cursorStart.SpaceIndex()
			} else if element == elementEnd {
				cleanUpEnd = cursorEnd.SpaceIndex()
				next = false
			}
			isEmpty := theSpace.CleanUpRange(cleanUpStart, cleanUpEnd)
			if isEmpty {
				r.rawList.Remove(element)
				theSpace.Free()
			}
			return
		})
	}
	return
}

func (r *spaceLinkList) findNotEmptySpaceAfter(i Cursor) Cursor {

}

// get Space at input index of area
func (r *spaceLinkList) getSpace(index int) Space {

}

func (r *spaceLinkList) nextArea(index int) int {

}

func (r *spaceLinkList) forRangeSpace(writeCursor Cursor, nextCursor Cursor, f func(space Space, isEnd bool) bool) {

}

// return if find
func (r *spaceLinkList) foreachElementStartAt(startElement *list.Element, f func(element *list.Element) bool) bool {
	ele := startElement
	if ele == nil {
		return false
	}
	if !f(ele) {
		return true
	}
	for {
		ele = ele.Next()
		if ele == nil || ele == startElement {
			return false
		}
		if !f(ele) {
			return true
		}
	}
}

func (r *spaceLinkList) forRangeElement(start Cursor, end Cursor, f func(space *list.Element, isEnd bool) bool) {

}
