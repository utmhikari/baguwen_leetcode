package solution_1_50



//11. 盛最多水的容器
//给你 n 个非负整数 a1，a2，...，an，每个数代表坐标中的一个点(i,ai) 。
//在坐标内画 n 条垂直线，垂直线 i的两个端点分别为(i,ai) 和 (i, 0) 。
//找出其中的两条线，使得它们与x轴共同构成的容器可以容纳最多的水。

func min11(a int, b int) int {
	if a < b {
		return a
	}
	return b
}

func max11(a int, b int) int {
	if a > b {
		return a
	}
	return b
}

func maxArea11(height []int) int {
	l := len(height)
	left, right := 0, l - 1

	mx := 0
	for left < right {
		mi := min11(height[left], height[right])
		sz := mi * (right - left)
		mx = max11(sz, mx)

		if height[left] == mi {
			for left < right && height[left] <= mi {
				left++
			}
		}

		if height[right] == mi {
			for left < right && height[right] <= mi {
				right--
			}
		}
	}

	return mx
}
