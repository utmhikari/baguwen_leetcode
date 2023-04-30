package solution_301_350

import "container/heap"

//347. 前 K 个高频元素
//给你一个整数数组 nums 和一个整数 k ，
//请你返回其中出现频率前 k 高的元素。
//你可以按 任意顺序 返回答案。


var counter347 map[int]int
var tuples347 [][]int


type IntHeap347 struct {
	IsDesc bool
	arr [][]int
}

func (h IntHeap347) Get(i int) []int {
	return h.arr[i]
}

func (h IntHeap347) Len() int {
	return len(h.arr)
}

func (h IntHeap347) Less(i, j int) bool {
	if h.IsDesc {
		return h.arr[i][1] > h.arr[j][1]
	} else {
		return h.arr[i][1] < h.arr[j][1]
	}

}

func (h IntHeap347) Swap(i, j int) {
	h.arr[i], h.arr[j] = h.arr[j], h.arr[i]
}

func (h *IntHeap347) Pop() interface{} {
	n := h.Len()
	old := h.arr
	x := old[n - 1]
	h.arr = h.arr[0 : n - 1]
	return x
}

func (h *IntHeap347) Push(x interface{}) {
	h.arr = append(h.arr, x.([]int))
}


var heap347 *IntHeap347


func topKFrequent(nums []int, k int) []int {
	counter347 = make(map[int]int)
	for _, n := range nums {
		cnt, ok := counter347[n]
		if ok {
			counter347[n] = cnt + 1
		} else {
			counter347[n] = 1
		}
	}
	tuples347 = make([][]int, 0)
	for k, v := range counter347 {
		tuples347 = append(tuples347, []int{k, v})
	}

	heap347 = &IntHeap347{}

	for _, tuple := range tuples347 {
		heap.Push(heap347, tuple)
		if heap347.Len() > k {
			heap.Pop(heap347)
		}
	}

	ans := make([]int, 0)
	for heap347.Len() > 0 {
		tuple := heap.Pop(heap347).([]int)
		ans = append(ans, tuple[0])
	}
	return ans
}