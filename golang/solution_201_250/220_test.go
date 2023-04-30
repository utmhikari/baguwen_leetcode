package solution_201_250

import "sort"

//220. 存在重复元素 III
//给你一个整数数组 nums 和两个整数k 和 t 。
//请你判断是否存在 两个不同下标 i 和 j，使得abs(nums[i] - nums[j]) <= t ，
//同时又满足 abs(i - j) <= k 。
//
//如果存在则返回 true，不存在返回 false。


func abs220(a int, b int) int {
	if a < b {
		return b - a
	}
	return a - b
}


func containsNearbyAlmostDuplicate(nums []int, k int, t int) bool {
	l := len(nums)
	if l < 2 {
		return false
	}
	numIndices := make([][]int, l)
	for i, n := range nums {
		numIndices[i] = []int{n, i}
	}
	sort.Slice(numIndices, func(i int, j int) bool {
		if numIndices[i][0] == numIndices[j][0] {
			return numIndices[i][1] < numIndices[j][1]
		} else {
			return numIndices[i][0] < numIndices[j][0]
		}
	})
	for i := 0; i < l - 1; i++ {
		for j := i + 1; j < l; j++ {
			if abs220(numIndices[j][0], numIndices[i][0]) > t {
				break
			}
			if abs220(numIndices[j][1], numIndices[i][1]) <= k {
				return true
			}
		}
	}
	return false
}