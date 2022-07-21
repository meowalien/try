package queue

type Heap []SortableItem

func (h Heap) Len() int { return len(h) }
func (h Heap) Cap() int { return cap(h) }
func (h Heap) Less(i, j int) bool {
	return h[i].Less(h[j])
}
func (h Heap) Swap(i, j int) { h[i], h[j] = h[j], h[i] }

func (h *Heap) Push(x interface{}) {
	*h = append(*h, x.(SortableItem))
}

func (h *Heap) Peek() interface{} {
	return (*h)[0]
}
func (h *Heap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}
