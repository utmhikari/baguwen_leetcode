package solution_351_400

import (
	"container/heap"
	"fmt"
)

//373. 查找和最小的K对数字

type data373 struct {
	i int
	j int
	v int
}

type heap373 struct {
	IsDesc bool
	arr []data373
}

func (h heap373) Get(i int) data373 {
	return h.arr[i]
}

func (h heap373) Len() int {
	return len(h.arr)
}

func (h heap373) Less(i, j int) bool {
	if h.IsDesc {
		return h.arr[i].v > h.arr[j].v
	} else {
		return h.arr[i].v < h.arr[j].v
	}
}

func (h heap373) Swap(i, j int) {
	h.arr[i], h.arr[j] = h.arr[j], h.arr[i]
}

func (h *heap373) Pop() interface{} {
	n := h.Len()
	x := h.arr[n - 1]
	h.arr = h.arr[0 : n - 1]
	return x
}

func (h *heap373) Push(x interface{}) {
	h.arr = append(h.arr, x.(data373))
}


func (h heap373) Peek() data373 {
	if h.Len() == 0 {
		return data373{}
	}
	return h.arr[0]
}

func getKey373(i, j int) string {
	return fmt.Sprintf("%d_%d", i, j)
}

func kSmallestPairs(nums1 []int, nums2 []int, k int) [][]int {
	l1, l2 := len(nums1), len(nums2)
	if l1 == 0 || l2 == 0 || k == 0{
		return [][]int{}
	}
	if k == 1 {
		return [][]int{{nums1[0], nums2[0]}}
	}
	visited := map[string]bool{
		getKey373(0, 0): true,
	}
	cnt := 0
	mx := k
	if mx > l1 * l2 {
		mx = l1 * l2
	}
	queue := &heap373{
		IsDesc: false,
		arr: []data373{
			{
				i: 0,
				j: 0,
				v: nums1[0] + nums2[0],
			},
		},
	}
	ans := make([][]int, 0)
	for cnt < mx {
		o := heap.Pop(queue).(data373)
		ans = append(ans, []int{nums1[o.i], nums2[o.j]})
		cnt++
		if cnt == mx {
			break
		}
		if o.j != l2 - 1 {
			nextRightKey := getKey373(o.i, o.j + 1)
			if !visited[nextRightKey] {
				heap.Push(queue, data373{
					i: o.i,
					j: o.j + 1,
					v: nums1[o.i] + nums2[o.j + 1],
				})
				visited[nextRightKey] = true
			}
		}
		if o.i != l1 - 1 {
			nextUpKey := getKey373(o.i + 1, o.j)
			if !visited[nextUpKey] {
				heap.Push(queue, data373{
					i: o.i + 1,
					j: o.j,
					v: nums1[o.i + 1] + nums2[o.j],
				})
				visited[nextUpKey] = true
			}
		}
	}
	return ans
}