package main

import (
	"fmt"
	queue2 "try/buffer_ring/planA/queue"
)

type Cap interface {
	Cap() (ans int)
}
type Len interface {
	Len() (ans int)
}

type Space interface {
	Cap
	Len
}

type Space5 [5]int

func (s Space5) Less(other queue2.SortableItem) bool {
	return s.Cap() > other.(Space).Cap()
}

func (s Space5) Cap() (ans int) {
	return 5
}

func (s *Space5) Len() (ans int) {
	return len(*s)
}

type Space10 [10]int

func (s Space10) Less(other queue2.SortableItem) bool {
	return s.Cap() > other.(Space).Cap()
}
func (s Space10) Cap() (ans int) {
	return 10
}

func (s *Space10) Len() (ans int) {
	return len(*s)
}

type BufferRing interface {
	queue2.Queue
	Space
}

type bufferRing struct {
	queue2.PriorityQueue
}

func (p *bufferRing) Len() (ans int) {
	for i := range p.Heap {
		ans += p.Heap[i].(Space).Len()
	}
	return
}
func (p *bufferRing) Cap() (ans int) {
	for i := range p.Heap {
		ans += p.Heap[i].(Space).Cap()
	}
	return
}

func NewBufferRing() BufferRing {
	return &bufferRing{}
}

func main() {
	a := Space5{1, 2, 3, 4, 5}
	b := Space10{6, 7, 8, 9, 10, 11, 12, 13, 14, 15}

	c := NewBufferRing()
	c.Push(&a)
	c.Push(&b)
	fmt.Println(c.Cap())
	fmt.Println(c.Len())
}
