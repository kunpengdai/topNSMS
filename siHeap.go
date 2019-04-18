package generatemsg

type SiHeap []si

func (h SiHeap) Len() int           { return len(h) }
func (h SiHeap) Less(i, j int) bool { return h[i].times < h[j].times }
func (h SiHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *SiHeap) Push(x interface{}) {
	// Push and Pop use pointer receivers because they modify the slice's length,
	// not just its contents.
	*h = append(*h, x.(si))
}

func (h *SiHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

func (h *SiHeap) Top() si {
	return (*h)[0]
}
