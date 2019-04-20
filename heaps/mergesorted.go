package heaps

import (
	"container/heap"
)

func MergeSorted(arrs [][]int) []int {
	h := &minHeap{}
	heap.Init(h)

	for _, outer := range arrs {
		for _, inner := range outer {
			heap.Push(h, inner)
		}
	}

	result := []int{}
	for h.Len() > 0 {
		next := heap.Pop(h)

		v, _ := next.(int)
		result = append(result, v)
	}

	return result
}

type minHeap []int

func (h minHeap) Len() int           { return len(h) }
func (h minHeap) Less(i, j int) bool { return h[i] < h[j] }
func (h minHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *minHeap) Push(v interface{}) {
	i, _ := v.(int)
	*h = append(*h, i)
}

func (h *minHeap) Pop() interface{} {
	var x int
	x, *h = (*h)[h.Len()-1], (*h)[:h.Len()-1]

	return x
}
