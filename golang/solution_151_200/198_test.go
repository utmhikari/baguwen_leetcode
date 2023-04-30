package solution_151_200



//198. 打家劫舍
//你是一个专业的小偷，计划偷窃沿街的房屋。每间房内都藏有一定的现金，
//影响你偷窃的唯一制约因素就是相邻的房屋装有相互连通的防盗系统，
//如果两间相邻的房屋在同一晚上被小偷闯入，系统会自动报警。
//
//给定一个代表每个房屋存放金额的非负整数数组，计算你 不触动警报装置的情况下 ，
//一夜之内能够偷窃到的最高金额。



func rob(nums []int) int {
	l := len(nums)
	switch l {
	case 0:
		return 0
	case 1:
		return nums[0]
	case 2:
		if nums[0] > nums[1] {
			return nums[0]
		} else {
			return nums[1]
		}
	default:
		maxes := make([]int, l)
		for i := 0; i < l; i++ {
			if i < 2 {
				maxes[i] = nums[i]
			} else if i == 2 {
				maxes[i] = nums[i] + maxes[i - 2]
			} else {
				mx := maxes[i - 2]
				if mx < maxes[i - 3] {
					mx = maxes[i - 3]
				}
				maxes[i] = nums[i] + mx
			}
		}
		if maxes[l - 1] > maxes[l - 2] {
			return maxes[l - 1]
		} else {
			return maxes[l - 2]
		}
	}
}