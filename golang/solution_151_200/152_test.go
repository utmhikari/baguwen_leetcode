package solution_151_200

import "math"

//152. 乘积最大子数组
//给你一个整数数组 nums ，请你找出数组中乘积最大的连续子数组（该子数组中至少包含一个数字），并返回该子数组所对应的乘积。



func maxProduct(nums []int) int {
	l := len(nums)

	if l == 0 {
		return 0
	}

	if l == 1 {
		return nums[0]
	}

	// dp: [(left, product)]
	// dpBz: [(left, -product)]

	dp := make([]int, l)
	for i := 0; i < l; i++ {
		dp[i] = -1
	}
	dpBz := make([]int, l)
	for i := 0; i < l; i++ {
		dpBz[i] = 1
	}

	// 0 idx
	if nums[0] < 0 {
		dpBz[0] = nums[0]
	} else {
		dp[0] = nums[0]
	}

	for i := 1; i < l; i++ {
		if nums[i] > 0 {
			if dp[i - 1] > 0 {
				dp[i] = nums[i] * dp[i - 1]
			} else {
				dp[i] = nums[i]
			}
			if dpBz[i - 1] < 0 {
				dpBz[i] = nums[i] * dpBz[i - 1]
			} else {
				dpBz[i] = 1
			}
		} else if nums[i] == 0 {
			dp[i] = 0
			dpBz[i] = 0
		} else {
			if dp[i - 1] > 0 {
				dpBz[i] = nums[i] * dp[i - 1]
			} else {
				dpBz[i] = nums[i]
			}
			if dpBz[i - 1] < 0 {
				dp[i] = nums[i] * dpBz[i - 1]
			} else {
				dp[i] = -1
			}
		}
	}

	mx := math.MinInt32
	for _, n := range dp {
		if n >= 0 && n > mx {
			mx = n
		}
	}

	if mx >= 0 {
		return mx
	}

	for _, n := range dpBz {
		if n < 0 && n > mx {
			mx = n
		}
	}

	return mx
}