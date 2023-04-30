package solution_1_50

import "sort"

//47. 全排列 II
//给定一个可包含重复数字的序列 nums
//按任意顺序 返回所有不重复的全排列。


var ans47 [][]int
var tmp47 []int
var visited47 []bool


func solve47(nums []int, left int) {
	l := len(nums)
	if left == l {
		newAns := make([]int, l)
		copy(newAns, tmp47)
		ans47 = append(ans47, newAns)
	} else {
		for i, v := range nums {
			if visited47[i] || i > 0 && !visited47[i - 1] && v == nums[i - 1] {
				continue
			}
			tmp47 = append(tmp47, v)
			visited47[i] = true
			solve47(nums, left + 1)
			visited47[i] = false
			tmp47 = tmp47[:len(tmp47) - 1]
		}
	}
}


func permuteUnique(nums []int) [][]int {
	l := len(nums)
	ans47 = make([][]int, 0)

	if l == 0 {
		return ans47
	}

	tmp47 = make([]int, 0)
	visited47 = make([]bool, l)

	sort.Ints(nums)
	solve47(nums, 0)

	return ans47
}