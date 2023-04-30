package solution_1_50



//45. 跳跃游戏 II
//给定一个非负整数数组，你最初位于数组的第一个位置。
//
//数组中的每个元素代表你在该位置可以跳跃的最大长度。
//
//你的目标是使用最少的跳跃次数到达数组的最后一个位置。
//
//假设你总是可以到达数组的最后一个位置。



func jump45(nums []int) int {
	l := len(nums)
	if l <= 1 {
		return 0
	}

	p := l - 1
	step := 0

	for p != 0 {
		i := 0
		for ; i < p; i++ {
			if i + nums[i] >= p {
				step++
				p = i
				break
			}
		}
	}

	return step
}