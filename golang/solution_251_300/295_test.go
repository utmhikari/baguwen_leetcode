package solution_251_300

import (
	"container/heap"
)

//295. 数据流的中位数
//
//中位数是有序列表中间的数。如果列表长度是偶数，中位数则是中间两个数的平均值。
//
//例如，
//
//[2,3,4] 的中位数是 3
//
//[2,3] 的中位数是 (2 + 3) / 2 = 2.5


type heap295 struct {
	IsDesc bool
	arr []int
}

func (h heap295) Get(i int) int {
	return h.arr[i]
}

func (h heap295) Len() int {
	return len(h.arr)
}

func (h heap295) Less(i, j int) bool {
	if h.IsDesc {
		return h.arr[i] > h.arr[j]
	} else {
		return h.arr[i] < h.arr[j]
	}
}

func (h heap295) Swap(i, j int) {
	h.arr[i], h.arr[j] = h.arr[j], h.arr[i]
}

func (h *heap295) Pop() interface{} {
	n := h.Len()
	x := h.arr[n - 1]
	h.arr = h.arr[0 : n - 1]
	return x
}

func (h *heap295) Push(x interface{}) {
	h.arr = append(h.arr, x.(int))
}


func (h heap295) Peek() int {
	if h.Len() == 0 {
		return 0
	}
	return h.arr[0]
}


type MedianFinder struct {
	hLeft  heap295
	hRight heap295
}


/** initialize your data structure here. */
func Constructor295() MedianFinder {
	return MedianFinder{
		hLeft:  heap295{IsDesc: true},
		hRight: heap295{},
	}
}


func (this *MedianFinder) AddNum(num int)  {
	ll, lr := this.hLeft.Len(), this.hRight.Len()
	if ll == 0 && lr == 0 {
		heap.Push(&this.hLeft, num)
		return
	}

	if ll == lr {
		rightMax := this.hRight.Peek()
		if num <= rightMax {
			heap.Push(&this.hLeft, num)
		} else {
			heap.Push(&this.hLeft, rightMax)
			heap.Pop(&this.hRight)
			heap.Push(&this.hRight, num)
		}
	} else {
		leftMax := this.hLeft.Peek()
		if num >= leftMax {
			heap.Push(&this.hRight, num)
		} else {
			heap.Pop(&this.hLeft)
			heap.Push(&this.hLeft, num)
			heap.Push(&this.hRight, leftMax)
		}
	}
}


func (this *MedianFinder) FindMedian() float64 {
	ll, lr := this.hLeft.Len(), this.hRight.Len()
	if ll == 0 && lr == 0 {
		return 0.0
	}

	if ll == lr {
		maxL, minR := float64(this.hLeft.Peek()), float64(this.hRight.Peek())
		return (maxL + minR) / 2.0
	} else {
		return float64(this.hLeft.Peek())
	}
}


/**
 * Your MedianFinder object will be instantiated and called as such:
 * obj := Constructor();
 * obj.AddNum(num);
 * param_2 := obj.FindMedian();
 */