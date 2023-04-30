package solution_201_250

import "math"

//213. 打家劫舍 II
//你是一个专业的小偷，计划偷窃沿街的房屋，每间房内都藏有一定的现金。这个地方所有的房屋都 围成一圈 ，
//这意味着第一个房屋和最后一个房屋是紧挨着的。
//同时，相邻的房屋装有相互连通的防盗系统，如果两间相邻的房屋在同一晚上被小偷闯入，系统会自动报警 。
//
//给定一个代表每个房屋存放金额的非负整数数组，计算你 在不触动警报装置的情况下 ，
//今晚能够偷窃到的最高金额。


func rob213(nums []int) int {
	l := len(nums)
	switch l {
	case 0:
		return 0
	case 1:
		return nums[0]
	case 2:
		if nums[0] >= nums[1] {
			return nums[0]
		}
		return nums[1]
	case 3:
		if nums[0] >= nums[1] && nums[0] >= nums[2] {
			return nums[0]
		}
		if nums[1] >= nums[2] {
			return nums[1]
		}
		return nums[2]
	default:
		// with curNum && with 0/without 0?
		maxes := make([][]int, l)
		maxes[0] = []int{nums[0], 0}
		maxes[1] = []int{0, nums[1]}
		maxes[2] = []int{nums[0] + nums[2], nums[2]}

		for i := 3; i < l; i++ {
			maxesBeforeTwo := maxes[i - 2]
			maxesBeforeThree := maxes[i - 3]

			maxes[i] = []int{nums[i], nums[i]}
			if i + 1 < l {
				if maxesBeforeTwo[0] > maxesBeforeThree[0] {
					maxes[i][0] += maxesBeforeTwo[0]
				} else {
					maxes[i][0] += maxesBeforeThree[0]
				}
			}

			if maxesBeforeTwo[1] > maxesBeforeThree[1] {
				maxes[i][1] += maxesBeforeTwo[1]
			} else {
				maxes[i][1] += maxesBeforeThree[1]
			}
		}

		mx := math.MinInt32
		for i := l - 1; i >= l - 3; i-- {
			if maxes[i][0] > mx {
				mx = maxes[i][0]
			}
			if maxes[i][1] > mx {
				mx = maxes[i][1]
			}
		}
		return mx
	}
}