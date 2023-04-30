package solution_51_100


//55. 跳跃游戏
//给定一个非负整数数组 nums ，你最初位于数组的 第一个下标 。
//
//数组中的每个元素代表你在该位置可以跳跃的最大长度。
//
//判断你是否能够到达最后一个下标。

func solve55(nums []int, cur int) bool {
	if cur == 0 {
		return true
	}

	for i := 0; i < cur; i++ {
		if nums[i] + i >= cur {
			return solve55(nums, i)
		}
	}

	return false
}

func canJump55(nums []int) bool {
	l := len(nums)
	if l <= 1 {
		return true
	}

	return solve55(nums, l - 1)
}