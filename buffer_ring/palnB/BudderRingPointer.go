package plainB

type BudderRingPointer [2]int

func (b BudderRingPointer) Layer(i int) int {
	return b[i]
}

func (b BudderRingPointer) LayerAdd(i int, data int) {
	b[i] += data
}

func (b BudderRingPointer) LayerSet(i int, data int) {
	b[i] = data
}
