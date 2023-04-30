package solution_51_100

//54. 螺旋矩阵
//给你一个 m 行 n 列的矩阵 matrix ，
//请按照 顺时针螺旋顺序 ，返回矩阵中的所有元素。



func spiralOrder54(matrix [][]int) []int {
	h := len(matrix)
	if h == 0 {
		return []int{}
	}

	w := len(matrix[0])
	if w == 0 {
		return []int{}
	}

	ans := make([]int, h * w)
	idx := 0
	x, lx, y, ly := 0, w - 1, 0, h - 1
	for x <= lx && y <= ly {
		if x == lx && y == ly {
			ans[idx] = matrix[x][y]
			break
		}
		if x == lx {
			for i := y; i <= ly; i++ {
				ans[idx] = matrix[i][x]
				idx++
			}
			break
		}
		if y == ly {
			for i := x; i <= lx; i++ {
				ans[idx] = matrix[y][i]
				idx++
			}
			break
		}
		for i := x; i < lx; i++ {
			ans[idx] = matrix[y][i]
			idx++
		}
		for i := y; i < ly; i++ {
			ans[idx] = matrix[i][lx]
			idx++
		}
		for i := lx; i > x; i-- {
			ans[idx] = matrix[ly][i]
			idx++
		}
		for i := ly; i > y; i-- {
			ans[idx] = matrix[i][x]
			idx++
		}
		x++
		lx--
		y++
		ly--
	}

	return ans
}

