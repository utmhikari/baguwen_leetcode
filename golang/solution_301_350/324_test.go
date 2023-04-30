package solution_301_350

import "sort"

//324. 摆动排序 II
//给你一个整数数组nums，将它重新排列成nums[0] < nums[1] > nums[2] < nums[3]...
//的顺序。
//
//你可以假设所有输入数组都可以得到满足题目要求的结果。



// ln, rn, ... l2, r2, l1, r1


func wiggleSort324(nums []int)  {
	if len(nums) < 2 {
		return
	}

	sort.Ints(nums)
	l := len(nums)
	mid := (l + 1) / 2

	left := make([]int, mid)
	for i := 0; i < mid; i++ {
		left[i] = nums[i]
	}
	right := make([]int, l - mid)
	for i := mid; i < l; i++ {
		right[i - mid] = nums[i]
	}

	if l % 2 == 0 {
		for i := 0; i < len(left); i++ {
			nums[l - 2 - 2 * i] = left[i]
		}
		for i := 0; i < len(right); i++ {
			nums[l - 1 - 2 * i] = right[i]
		}
	} else {
		for i := 0; i < len(left); i++ {
			nums[l - 1 - 2 * i] = left[i]
		}
		for i := 0; i < len(right); i++ {
			nums[l - 2 - 2 * i] = right[i]
		}
	}
}