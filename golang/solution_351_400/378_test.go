package solution_351_400

import "container/heap"

//378. 有序矩阵中第 K 小的元素


type v378 struct {
	i int
	j int
	v int
}

type heap378 struct {
	IsDesc bool
	arr []v378
}

func (h heap378) Get(i int) v378 {
	return h.arr[i]
}

func (h heap378) Len() int {
	return len(h.arr)
}

func (h heap378) Less(i, j int) bool {
	if h.IsDesc {
		return h.arr[i].v > h.arr[j].v
	} else {
		return h.arr[i].v < h.arr[j].v
	}
}

func (h heap378) Swap(i, j int) {
	h.arr[i], h.arr[j] = h.arr[j], h.arr[i]
}

func (h *heap378) Pop() interface{} {
	n := h.Len()
	x := h.arr[n - 1]
	h.arr = h.arr[0 : n - 1]
	return x
}

func (h *heap378) Push(x interface{}) {
	h.arr = append(h.arr, x.(v378))
}


func (h heap378) Peek() v378 {
	if h.Len() == 0 {
		return v378{}
	}
	return h.arr[0]
}


var visited378 map[int]bool

func getKey378(i, j, w int) int {
	return i * w + j
}


func kthSmallest(matrix [][]int, k int) int {
	h := len(matrix)
	if h == 0 {
		return 0
	}
	w := len(matrix[0])
	if w == 0 {
		return 0
	}
	if k > h * w {
		return 0
	}

	visited378 = make(map[int]bool)
	cnt := 0
	cur := 0
	hp := &heap378{}

	for cnt < k {
		if hp.Len() == 0 {
			visited378[0] = true
			cur = matrix[0][0]
			if h > 1 {
				heap.Push(hp, v378{
					1, 0, matrix[1][0],
				})
			}
			if w > 1 {
				heap.Push(hp, v378{
					0, 1, matrix[0][1],
				})
			}
		} else {
			v := heap.Pop(hp).(v378)
			cur = v.v
			if v.i < h - 1 {
				key := getKey378(v.i + 1, v.j, w)
				if !visited378[key] {
					visited378[key] = true
					heap.Push(hp, v378{
						v.i + 1, v.j, matrix[v.i + 1][v.j],
					})
				}
			}
			if v.j < w - 1 {
				key := getKey378(v.i, v.j + 1, w)
				if !visited378[key] {
					visited378[key] = true
					heap.Push(hp, v378{
						v.i, v.j + 1, matrix[v.i][v.j + 1],
					})
				}
			}
		}
		cnt++
	}

	return cur
}
