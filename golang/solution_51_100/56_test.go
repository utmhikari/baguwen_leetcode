package solution_51_100

import (
	"sort"
)

//56. 合并区间
//以数组 intervals 表示若干个区间的集合，其中单个区间为 intervals[i] = [starti, endi] 。
//请你合并所有重叠的区间，并返回一个不重叠的区间数组，该数组需恰好覆盖输入中的所有区间。


func merge56(intervals [][]int) [][]int {
	l := len(intervals)
	if l <= 1 {
		return intervals
	}

	sort.Slice(intervals, func(i, j int) bool {
		if intervals[i][0] != intervals[j][0] {
			return intervals[i][0] < intervals[j][0]
		}
		return intervals[i][1] < intervals[j][1]
	})
	//fmt.Printf("%v\n", intervals)

	// <1> merge left
	// <2> merge right
	// <3> include
	// <4> exclude

	ans := [][]int{intervals[0]}
	for i := 1; i < l; i++ {
		interval := intervals[i]
		if ans[len(ans) - 1][1] < interval[0] {
			ans = append(ans, interval)
		} else if interval[1] > ans[len(ans) - 1][1] {
			ans[len(ans) - 1][1] = interval[1]
		}
	}

	return ans
}