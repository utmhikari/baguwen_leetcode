package solution_351_400

//352. 将数据流变为多个不相交区间
//给定一个非负整数的数据流输入 a1，a2，…，an，…，将到目前为止看到的数字总结为不相交的区间列表。
//
//例如，假设数据流中的整数为 1，3，7，2，6，…，每次的总结为：




type SummaryRanges struct {
	nums map[int]bool
	intervals [][]int
}


/** Initialize your data structure here. */
func Constructor352() SummaryRanges {
	return SummaryRanges{
		nums: make(map[int]bool),
		intervals: make([][]int, 0),
	}
}

// Len of intervals
func (this *SummaryRanges) Len() int {
	return len(this.intervals)
}


// find the first left idx against val
func (this *SummaryRanges) findFirstLeft(val int, left int, right int) int {
	if left > right {
		return right
	}

	if val > this.intervals[right][1] {
		return right
	}

	mid := (left + right) / 2
	if val > this.intervals[mid][1] {
		return this.findFirstLeft(val, mid + 1, right)
	} else {
		return this.findFirstLeft(val, left, mid - 1)
	}
}


// find the first right idx against val
func (this *SummaryRanges) findFirstRight(val int, left int, right int) int {
	if left > right {
		return left
	}

	if val < this.intervals[left][0] {
		return left
	}

	mid := (left + right) / 2
	if val < this.intervals[mid][0] {
		return this.findFirstRight(val, left, mid - 1)
	} else {
		return this.findFirstRight(val, mid + 1, right)
	}
}

// merge val with interval at left & right indices
func (this *SummaryRanges) merge(val int, left int, right int) {
	l := this.Len()

	if right == 0 {
		if this.intervals[0][0] - 1 == val {
			this.intervals[0][0] = val
		} else {
			this.intervals = append([][]int{{val, val}}, this.intervals...)
		}
		return
	}

	if left == this.Len() - 1 {
		if this.intervals[l - 1][1] + 1 == val {
			this.intervals[l - 1][1] = val
		} else {
			this.intervals = append(this.intervals, []int{val, val})
		}
		return
	}

	newIntervals := make([][]int, 0)
	for i := 0; i < left; i++ {
		newIntervals = append(newIntervals, this.intervals[i])
	}

	isConnLeft := this.intervals[left][1] + 1 == val
	isConnRight := this.intervals[right][0] - 1 == val
	if isConnLeft && isConnRight {
		newIntervals = append(newIntervals, []int{
			this.intervals[left][0], this.intervals[right][1],
		})
	} else if isConnLeft {
		newIntervals = append(newIntervals, []int{
			this.intervals[left][0], val,
		})
		newIntervals = append(newIntervals, []int{
			this.intervals[right][0], this.intervals[right][1],
		})
	} else if isConnRight {
		newIntervals = append(newIntervals, []int{
			this.intervals[left][0], this.intervals[left][1],
		})
		newIntervals = append(newIntervals, []int{
			val, this.intervals[right][1],
		})
	} else {
		newIntervals = append(newIntervals, []int{
			this.intervals[left][0], this.intervals[left][1],
		})
		newIntervals = append(newIntervals, []int{
			val, val,
		})
		newIntervals = append(newIntervals, []int{
			this.intervals[right][0], this.intervals[right][1],
		})
	}

	for i := right + 1; i < this.Len(); i++ {
		newIntervals = append(newIntervals, this.intervals[i])
	}

	this.intervals = newIntervals
}


func (this *SummaryRanges) AddNum(val int)  {
	_, added := this.nums[val]
	if added {
		return
	}
	this.nums[val] = true

	// first interval
	if this.Len() == 0 {
		this.intervals = append(this.intervals, []int{val, val})
		return
	}

	if val > this.intervals[this.Len() - 1][1] {
		this.merge(val, this.Len() - 1, -1)
	} else if val < this.intervals[0][0] {
		this.merge(val, -1, 0)
	} else {
		firstLeft := this.findFirstLeft(val, 0, this.Len() - 1)
		firstRight := this.findFirstRight(val, 0, this.Len() - 1)
		//fmt.Printf("intervals: %v, val: %d, left: %d, right: %d\n", this.intervals, val, firstLeft, firstRight)
		this.merge(val, firstLeft, firstRight)
	}
}


func (this *SummaryRanges) GetIntervals() [][]int {
	return this.intervals
}