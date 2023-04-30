package solution_51_100

//84. 柱状图中最大的矩形
//给定 n 个非负整数，用来表示柱状图中各个柱子的高度。每个柱子彼此相邻，且宽度为 1 。
//求在该柱状图中，能够勾勒出来的矩形的最大面积。


// 2, 1, 5, 6, 2, 3
// 6 -> 1
// 5 -> 1
// 4 -> 2
// 3 -> 2
// 2 -> 5
// 1 -> 6


func largestRectangleArea84(heights []int) int {
	l := len(heights)
	if l == 0 {
		return 0
	}
	if l == 1 {
		return heights[0]
	}

	stk := make([]int, 0)
	leftMinIndices, rightMinIndices := make([]int, l), make([]int, l)
	for i := 0; i < l; i++ {
		for len(stk) > 0 && heights[stk[len(stk) - 1]] >= heights[i] {
			stk = stk[:len(stk) - 1]
		}
		if len(stk) == 0 {
			leftMinIndices[i] = -1
		} else {
			leftMinIndices[i] = stk[len(stk) - 1]
		}
		stk = append(stk, i)
	}
	stk = make([]int, 0)
	for i := l - 1; i >= 0; i-- {
		for len(stk) > 0 && heights[stk[len(stk) - 1]] >= heights[i] {
			stk = stk[:len(stk) - 1]
		}
		if len(stk) == 0 {
			rightMinIndices[i] = l
		} else {
			rightMinIndices[i] = stk[len(stk) - 1]
		}
		stk = append(stk, i)
	}

	mx := 0
	for i := 0; i < l; i++ {
		area := heights[i] * (rightMinIndices[i] - leftMinIndices[i] - 1)
		if area > mx {
			mx = area
		}
	}
	return mx
}