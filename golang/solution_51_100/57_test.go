package solution_51_100

//57. 插入区间
//给你一个 无重叠的 ，按照区间起始端点排序的区间列表。
//
//在列表中插入一个新的区间，你需要确保列表中的区间仍然有序且不重叠（如果有必要的话，可以合并区间）。

func insert57(intervals [][]int, newInterval []int) [][]int {
	l := len(intervals)
	if l == 0 {
		return [][]int{newInterval}
	}

	if newInterval[1] < intervals[0][0] {
		return append([][]int{newInterval}, intervals...)
	}
	if newInterval[0] > intervals[l - 1][1] {
		return append(intervals, newInterval)
	}

	// [[1,3],[6,9], [13,15]]
	// <1, 1~3, 4~5, 6~9, 10~12, 13~15, >15
	// 0,   1,   2,   3,   4,     5,     6
	leftPos, rightPos := -1, -1
	for i := 0; i < l; i++ {
		if intervals[i][0] <= newInterval[0] && intervals[i][1] >= newInterval[0] {
			leftPos = 2 * i + 1
		}
		if intervals[i][0] <= newInterval[1] && intervals[i][1] >= newInterval[1] {
			rightPos = 2 * i + 1
		}
		if leftPos >= 0 && rightPos >= 0 {
			break
		}
		if i == 0 {
			if intervals[i][0] > newInterval[0] {
				leftPos = 0
			}
		} else {
			if intervals[i - 1][1] < newInterval[0] && intervals[i][0] > newInterval[0] {
				leftPos = 2 * i
			}
			if intervals[i - 1][1] < newInterval[1] && intervals[i][0] > newInterval[1] {
				rightPos = 2 * i
			}
		}
		if leftPos >= 0 && rightPos >= 0 {
			break
		}
	}
	if rightPos == -1 {
		rightPos = 2 * l
	}

	//fmt.Printf("intervals: %v, newInterval: %v, leftPos: %d, rightPos: %d\n", intervals, newInterval, leftPos, rightPos)

	if leftPos == rightPos {
		if leftPos % 2 == 1 {  // is included
			return intervals
		} else {
			rightIdx := leftPos / 2
			ans := make([][]int, l + 1)
			for i := 0; i < rightIdx; i++ {
				ans[i] = intervals[i]
			}
			ans[rightIdx] = newInterval
			for i := rightIdx + 1; i < l + 1; i++ {
				ans[i] = intervals[i - 1]
			}
			return ans
		}
	}

	if leftPos == 0 && rightPos == 2 * l {
		return [][]int{newInterval}
	}

	if leftPos == 0 {
		rightIdx := rightPos / 2
		if rightPos % 2 == 0 {  // no merge
			return append([][]int{newInterval}, intervals[rightIdx:]...)
		} else {  // merge with rightIdx
			return append([][]int{{newInterval[0], intervals[rightIdx][1]}}, intervals[rightIdx + 1:]...)
		}
	}

	if rightPos == 2 * l {
		leftIdx := leftPos / 2
		if leftPos % 2 == 0 {  // no merge
			return append(intervals[:leftIdx], newInterval)
		} else {  // merge with leftIdx
			return append(intervals[:leftIdx], []int{intervals[leftIdx][0], newInterval[1]})
		}
	}

	removeStart, removeEnd := leftPos / 2, (rightPos / 2) - 1
	if leftPos % 2 == 1 {
		newInterval[0] = intervals[removeStart][0]
	}
	if rightPos % 2 == 1 {
		removeEnd++
		newInterval[1] = intervals[removeEnd][1]
	}
	ans := make([][]int, 0)
	for i := 0; i < removeStart; i++ {
		ans = append(ans, intervals[i])
	}
	ans = append(ans, newInterval)
	for i := removeEnd + 1; i < l; i++ {
		ans = append(ans, intervals[i])
	}
	return ans
}