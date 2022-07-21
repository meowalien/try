package queue

import (
	sheep "container/heap"
)

type PriorityQueue struct {
	Heap
}

func (p *PriorityQueue) Push(item SortableItem) {
	sheep.Push(&p.Heap, item)
}

func (p *PriorityQueue) Pop() SortableItem {
	return sheep.Pop(&p.Heap).(SortableItem)
}

func (p *PriorityQueue) Peek() SortableItem {
	return p.Heap.Peek().(SortableItem)
}

func (p *PriorityQueue) Len() int {
	return p.Heap.Len()
}
func (p *PriorityQueue) Cap() int {
	return p.Heap.Cap()
}
