package solution_201_250



//221. 最大正方形
//在一个由 '0' 和 '1' 组成的二维矩阵内，找到只包含 '1' 的最大正方形，并返回其面积。


func maximalSquare221(matrix [][]byte) int {
	h := len(matrix)
	if h == 0 {
		return 0
	}
	w := len(matrix[0])
	if w == 0 {
		return 0
	}

	dp := make([][]int, h)
	for i := 0; i < h; i++ {
		dp[i] = make([]int, w)
	}

	mx := 0
	for i := 0; i < h; i++ {
		for j := 0; j < w; j++ {
			if matrix[i][j] == '1' {
				if i == 0 || j == 0 {
					dp[i][j] = 1
				} else {
					dp[i][j] = dp[i - 1][j - 1]
					if dp[i - 1][j] < dp[i][j] {
						dp[i][j] = dp[i - 1][j]
					}
					if dp[i][j - 1] < dp[i][j] {
						dp[i][j] = dp[i][j - 1]
					}
					dp[i][j]++
				}
				if dp[i][j] > mx {
					mx = dp[i][j]
				}
			}
		}
	}

	return mx * mx
}