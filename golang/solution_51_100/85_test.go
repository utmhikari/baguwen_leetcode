package solution_51_100



//85. 最大矩形
//给定一个仅包含 0 和 1 、大小为 rows x cols 的二维二进制矩阵，
//找出只包含 1 的最大矩形，并返回其面积。


func maximalRectangle(matrix [][]byte) int {
	h := len(matrix)
	if h == 0 {
		return 0
	}
	w := len(matrix[0])
	if w == 0 {
		return 0
	}

	leftLens := make([][]int, h)
	for i := 0; i < h; i++ {
		leftLens[i] = make([]int, w)
	}

	for i := 0; i < h; i++ {
		for j := 0; j < w; j++ {
			if matrix[i][j] == '1' {
				if j == 0 {
					leftLens[i][j] = 1
				} else {
					leftLens[i][j] = leftLens[i][j - 1] + 1
				}
			}
		}
	}

	mx := 0
	for i := 0; i < h; i++ {
		for j := 0; j < w; j++ {
			if matrix[i][j] == '0' {
				continue
			}
			width := leftLens[i][j]
			tmpMx := width
			for k := i - 1; k >= 0; k-- {
				if matrix[i][j] != '0' {
					if leftLens[k][j] < width {
						width = leftLens[k][j]
					}
					area := width * (i - k + 1)
					if area > tmpMx {
						tmpMx = area
					}
				}
			}
			if tmpMx > mx {
				mx = tmpMx
			}
		}
	}
	return mx
}
