package solution_51_100



// 74. 搜索二维矩阵


var visited74 [][]bool

func search74(matrix [][]int, target int, i int, j int) bool {
	if i < 0 || i >= len(matrix) {
		return false
	}
	if j < 0 || j >= len(matrix[0]) {
		return false
	}
	if target == matrix[i][j] {
		return true
	}
	if visited74[i][j] {
		return false
	}
	visited74[i][j] = true
	if target < matrix[i][j] {
		if search74(matrix, target, i - 1, j) {
			return true
		}
		if search74(matrix, target, i, j - 1) {
			return true
		}
	} else {
		if search74(matrix, target, i + 1, j) {
			return true
		}
		if search74(matrix, target, i, j + 1) {
			return true
		}
	}

	return false
}


func searchMatrix74(matrix [][]int, target int) bool {
	h := len(matrix)
	if h == 0 {
		return false
	}
	w := len(matrix[0])
	if w == 0 {
		return false
	}
	if target < matrix[0][0] {
		return false
	}
	if target > matrix[h - 1][w - 1] {
		return false
	}

	visited74 = make([][]bool, h)
	for i := 0; i < h; i++ {
		visited74[i] = make([]bool, w)
	}

	return search74(matrix, target, 0, 0)
}