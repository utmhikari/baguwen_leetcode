package main



//867. 转置矩阵
//给你一个二维整数数组 matrix， 返回 matrix 的 转置矩阵 。
//
//矩阵的 转置 是指将矩阵的主对角线翻转，交换矩阵的行索引与列索引。





func transpose867(matrix [][]int) [][]int {
	h := len(matrix)
	if h == 0 {
		return make([][]int, 0)
	}

	w := len(matrix[0])
	if w == 0 {
		return make([][]int, 0)
	}

	ans := make([][]int, w)
	for i := 0; i < w; i++ {
		ans[i] = make([]int, h)
	}

	for i := 0; i < h; i++ {
		for j := 0; j < w; j++ {
			ans[j][i] = matrix[i][j]
		}
	}

	return ans
}