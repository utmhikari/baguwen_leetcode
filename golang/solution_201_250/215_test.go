package solution_201_250

import (
	"container/heap"
)

type IntHeap215 struct {
	IsDesc bool
	arr []int
}


func (h IntHeap215) Get(i int) int {
	return h.arr[i]
}

func (h IntHeap215) Len() int {
	return len(h.arr)
}

func (h IntHeap215) Less(i, j int) bool {
	if h.IsDesc {
		return h.arr[i] > h.arr[j]
	} else {
		return h.arr[i] < h.arr[j]
	}

}

func (h IntHeap215) Swap(i, j int) {
	h.arr[i], h.arr[j] = h.arr[j], h.arr[i]
}

func (h *IntHeap215) Pop() interface{} {
	n := h.Len()
	old := h.arr
	x := old[n - 1]
	h.arr = h.arr[0 : n - 1]
	return x
}

func (h *IntHeap215) Push(x interface{}) {
	h.arr = append(h.arr, x.(int))
}


func findKthLargest(nums []int, k int) int {
	h := &IntHeap215{IsDesc: true}
	heap.Init(h)
	for i := 0; i < len(nums); i++ {
		heap.Push(h, nums[i])
 	}
 	var x int
 	for i := 0; i < k; i++ {
		x = heap.Pop(h).(int)
	}
	return x
}