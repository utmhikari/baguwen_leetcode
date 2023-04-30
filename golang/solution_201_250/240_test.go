package solution_201_250

//240. 搜索二维矩阵 II
//编写一个高效的算法来搜索mxn矩阵 matrix 中的一个目标值 target 。该矩阵具有以下特性：
//
//每行的元素从左到右升序排列。
//每列的元素从上到下升序排列。


var visited240 map[int]bool
var ans240 bool

func search240(matrix [][]int, target int, x int, y int) {
	if ans240 {
		return
	}

	if x < 0 || y < 0 || x >= len(matrix) || y >= len(matrix[0]) {
		return
	}

	//fmt.Printf("target %d, x: %d, y: %d, value: %d\n", target, x, y, matrix[x][y])
	if matrix[x][y] == target {
		ans240 = true
		return
	}

	key := x * len(matrix[0]) + y
	_, ok := visited240[key]
	if ok {
		return
	}
	visited240[key] = true

	if matrix[x][y] > target {
		search240(matrix, target, x - 1, y)
		search240(matrix, target, x, y - 1)
	} else {
		search240(matrix, target, x + 1, y)
		search240(matrix, target, x, y + 1)
	}
}


func searchMatrix240(matrix [][]int, target int) bool {
	h := len(matrix)
	if h == 0 {
		return false
	}
	w := len(matrix[0])
	if w == 0 {
		return false
	}

	if target < matrix[0][0] || target > matrix[h - 1][w - 1] {
		return false
	}

	visited240 = make(map[int]bool)
	ans240 = false

	search240(matrix, target, 0, 0)
	return ans240
}
