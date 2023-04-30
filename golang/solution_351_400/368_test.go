package solution_351_400

import "sort"

func largestDivisibleSubset(nums []int) []int {
	sort.Ints(nums)
	l := len(nums)
	if l == 0 {
		return []int{}
	}
	if l == 1 {
		return []int{nums[0]}
	}
	mxCnt := 0
	mxIdx := -1
	cnts := make([]int, l)
	lastIndices := make([]int, l)
	for i := 0; i < l; i++ {
		lastIndices[i] = -1
	}
	for i := 1; i < l; i++ {
		for j := 0; j < i; j++ {
			if nums[i] % nums[j] == 0 {
				curCnt := cnts[j] + 1
				if curCnt > cnts[i] {
					cnts[i] = curCnt
					lastIndices[i] = j
				}
			}
		}
		if cnts[i] > mxCnt {
			mxCnt = cnts[i]
			mxIdx = i
		}
	}
	if mxIdx == -1 {
		return []int{nums[0]}
	}

	ans := make([]int, 0)
	idx := mxIdx
	for idx != -1 {
		ans = append(ans, nums[idx])
		idx = lastIndices[idx]
	}
	return ans

}
