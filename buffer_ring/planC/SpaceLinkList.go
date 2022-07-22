package planC

import "container/list"

// todo: change firstElement if better

type LinkListRing[T any] interface {
	SpaceRing[T]
}

func newSpaceLinkList[T any](space ...Space[T]) LinkListRing[T] {
	if len(space) == 0 {
		panic("spaceLinkList must have at least one space")
	}
	l := list.New()
	firstElement := l.PushBack(space[0])
	for i := 1; i < len(space); i++ {
		l.PushBack(space[0])
	}

	return &spaceLinkList[T]{rawList: l, firstElement: firstElement}
}

type spaceLinkList[T any] struct {
	// todo: change to ring link list
	rawList      *list.List
	firstElement *list.Element
}

func (r *spaceLinkList[T]) insertSpaceBeforeCursor(cursor Cursor, space Space[T]) {
	element, ok := r.findElementByCursor(r.firstElement, cursor)
	if !ok {
		panic("cursor not found")
	}
	r.rawList.InsertBefore(space, element)
}

func (r *spaceLinkList[T]) findElementByCursor(startElement *list.Element, c Cursor) (elementFound *list.Element, ok bool) {
	ok = r.foreachElement(startElement, func(element *list.Element) bool {
		if elementFound.Value.(Space[T]).AreaID() == c.AreaID() {
			elementFound = element
			return false
		}
		return true
	})
	return
}

func (r *spaceLinkList[T]) TotalRemainingSpace() (total int) {
	_ = r.foreachElement(r.firstElement, func(element *list.Element) bool {
		total += element.Value.(Space[T]).RemainingSpace()
		return true
	})
	return
}

func (r *spaceLinkList[T]) cleanUpSpaceInRange(pair CursorPair) {
	cursorStart := pair.GetStartCursor()
	cursorEnd := pair.GetStartCursor()
	elementStart, ok := r.findElementByCursor(r.firstElement, cursorStart)
	if !ok {
		panic("cursor not found")
	}

	if cursorStart.AreaID() == cursorEnd.AreaID() {
		theSpace := elementStart.Value.(Space[T])
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
		_ = r.foreachElement(elementStart, func(element *list.Element) (next bool) {
			if element == elementEnd {
				return
			}
			theSpace := element.Value.(Space[T])
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

func (r *spaceLinkList[T]) findNotEmptySpaceAfter(i Cursor) Cursor {
	theElement, ok := r.findElementByCursor(r.firstElement, i)
	if !ok {
		panic("cursor not found")
	}
	var foundElement *list.Element
	_ = r.foreachElement(theElement, func(element *list.Element) bool {
		if !element.Value.(Space[T]).IsFree() {
			foundElement = element
			return false
		}
		return true
	})
	if foundElement == theElement {
		return i
	} else {
		// todo: move to space 0 will waste some memory
		return newCursor[T](r, foundElement.Value.(Space[T]).AreaID(), 0)
	}
}

// get Space at input index of area
func (r *spaceLinkList[T]) getSpace(index int) (theSpace Space[T]) {
	ok := r.foreachElement(r.firstElement, func(element *list.Element) bool {
		if theSpace = element.Value.(Space[T]); theSpace.AreaID() == index {
			return false
		}
		return true
	})
	if !ok {
		panic("area not found")
	}
	return
}

func (r *spaceLinkList[T]) nextArea(index int) int {
	var theElement *list.Element
	ok := r.foreachElement(r.firstElement, func(element *list.Element) bool {
		if theElement = element; theElement.Value.(Space[T]).AreaID() == index {
			return false
		}
		return true
	})
	if !ok {
		panic("area not found")
	}
	return theElement.Next().Value.(Space[T]).AreaID()
}

func (r *spaceLinkList[T]) forRangeSpace(start Cursor, end Cursor, f func(space Space[T], isEnd bool) bool) {
	elementStart, ok := r.findElementByCursor(r.firstElement, start)
	if !ok {
		panic("cursor not found")
	}
	elementEnd, ok := r.findElementByCursor(elementStart, end)
	if !ok {
		panic("cursor not found")
	}
	_ = r.foreachElement(elementStart, func(element *list.Element) (next bool) {
		next = f(element.Value.(Space[T]), element == elementEnd)
		return
	})
}

// return if find
func (r *spaceLinkList[T]) foreachElement(startElement *list.Element, f func(element *list.Element) bool) bool {
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
