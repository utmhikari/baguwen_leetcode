package solution_1_50



//42. 接雨水
//给定 n 个非负整数表示每个宽度为 1 的柱子的高度图，
//计算按此排列的柱子，下雨之后能接多少雨水。


func min42(a int, b int) int {
	if a < b {
		return a
	}
	return b
}

func max42(a int, b int) int {
	if a > b {
		return a
	}
	return b
}


func trap42(height []int) int {
	l := len(height)
	if l <= 2 {
		return 0
	}

	maxLefts, maxRights := make([]int, l), make([]int, l)
	maxLefts[0], maxRights[l - 1] = 0, 0
	for i := 1; i < l; i++ {
		maxLefts[i] = max42(maxLefts[i - 1], height[i - 1])
		maxRights[l - 1 - i] = max42(maxRights[l - i], height[l - i])
	}

	sm := 0
	for i := 0; i < l; i++ {
		am := min42(maxLefts[i], maxRights[i]) - height[i]
		if am > 0 {
			sm += am
		}
	}

	return sm
}