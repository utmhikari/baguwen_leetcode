package solution_51_100


//53. 最大子序和
//给定一个整数数组 nums ，找到一个具有最大和的连续子数组（子数组最少包含一个元素），返回其最大和。

import "math"

func maxSubArray53(nums []int) int {
	mx := math.MinInt32
	sm := 0
	for _, n := range nums {
		sm += n
		if sm > mx {
			mx = sm
		}
		if sm < 0 {
			sm = 0
		}
	}
	return mx
}