package solution_1_50



//21. 直方图的水量
//给定一个直方图(也称柱状图)，假设有人从上面源源不断地倒水
//最后直方图能存多少水量?直方图的宽度为 1。


func max21(a int, b int) int {
	if a > b {
		return a
	}
	return b
}

func min21(a int, b int) int {
	if a < b {
		return a
	}
	return b
}


func trap21(height []int) int {
	l := len(height)
	if l <= 2 {
		return 0
	}

	leftMaxes, rightMaxes := make([]int, l), make([]int, l)
	leftMaxes[0] = height[0]
	for i := 1; i < l; i++ {
		leftMaxes[i] = max21(leftMaxes[i - 1], height[i])
	}
	rightMaxes[l - 1] = height[l - 1]
	for i := l - 2; i >= 0; i-- {
		rightMaxes[i] = max21(rightMaxes[i + 1], height[i])
	}

	sm := 0
	for i := 0; i < l; i++ {
		sm += min21(leftMaxes[i] - height[i], rightMaxes[i] - height[i])
	}

	return sm
}