package solution_201_250

import "math"

//209. 长度最小的子数组
//给定一个含有n个正整数的数组和一个正整数 target 。
//
//找出该数组中满足其和 ≥ target 的长度最小的 连续子数组
//[numsl, numsl+1, ..., numsr-1, numsr] ，并返回其长度。
//如果不存在符合条件的子数组，返回 0 。



func minSubArrayLen209(target int, nums []int) int {
	sm := 0
	minLen := math.MaxInt32
	l := len(nums)
	left, right := 0, -1
	for right < l {
		if sm < target {
			right++
			if right == l {
				break
			}
			sm += nums[right]
		} else {
			if sm - nums[left] >= target {
				sm -= nums[left]
				left++
			}
			//fmt.Printf("sm: %d, left: %d, right: %d\n", sm, left, right)
			curLen := right - left + 1
			if curLen < minLen {
				minLen = curLen
			}
			if left == right {
				return 1
			} else {
				sm -= nums[left]
				left++
			}
		}
	}
	if minLen == math.MaxInt32 {
		return 0
	} else {
		return minLen
	}
}
