package solution_201_250



//223. 矩形面积
//给你 二维 平面上两个 由直线构成的 矩形，请你计算并返回两个矩形覆盖的总面积。
//
//每个矩形由其 左下 顶点和 右上 顶点坐标表示：
//
//第一个矩形由其左下顶点 (ax1, ay1) 和右上顶点 (ax2, ay2) 定义。
//第二个矩形由其左下顶点 (bx1, by1) 和右上顶点 (bx2, by2) 定义。


func min223(a int, b int) int {
	if a < b {
		return a
	}
	return b
}

func max223(a int, b int) int {
	if a > b {
		return a
	}
	return b
}

func computeArea(ax1 int, ay1 int, ax2 int, ay2 int, bx1 int, by1 int, bx2 int, by2 int) int {
	szA := (ax2 - ax1) * (ay2 - ay1)
	szB := (bx2 - bx1) * (by2 - by1)

	// separate?
	if bx1 >= ax2 || bx2 <= ax1 || by1 >= ay2 || by2 <= ay1 {
		return szA + szB
	}

	// a includes b?
	if ax1 <= bx1 && ax2 >= bx2 && ay1 <= by1 && ay2 >= by2 {
		return szA
	}

	// b includes a?
	if bx1 <= ax1 && bx2 >= ax2 && by1 <= ay1 && by2 >= ay2 {
		return szB
	}

	// get dx & dy
	dx := min223(bx2, ax2) - max223(bx1, ax1)
	dy := min223(by2, ay2) - max223(by1, ay1)
	return szA + szB - dx * dy
}