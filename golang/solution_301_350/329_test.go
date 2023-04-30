package solution_301_350




//329. 矩阵中的最长递增路径
//给定一个 m x n 整数矩阵 matrix ，找出其中 最长递增路径 的长度。
//
//对于每个单元格，你可以往上，下，左，右四个方向移动。 你
//不能 在 对角线 方向上移动或移动到 边界外（即不允许环绕）。


var matrix329 [][]int
var longests329 [][]int
var h329 int
var w329 int


func search329(i int, j int) {
	if longests329[i][j] != 0 {
		return
	}

	mxToAdd := 0

	// left
	if j > 0 && matrix329[i][j] < matrix329[i][j - 1] {
		search329(i, j - 1)
		if longests329[i][j - 1] > mxToAdd {
			mxToAdd = longests329[i][j - 1]
		}
	}

	// right
	if j < w329- 1 && matrix329[i][j] < matrix329[i][j + 1] {
		search329(i, j + 1)
		if longests329[i][j + 1] > mxToAdd {
			mxToAdd = longests329[i][j + 1]
		}
	}

	// up
	if i > 0 && matrix329[i][j] < matrix329[i - 1][j] {
		search329(i - 1, j)
		if longests329[i - 1][j] > mxToAdd {
			mxToAdd = longests329[i - 1][j]
		}
	}

	// down
	if i < h329- 1 && matrix329[i][j] < matrix329[i + 1][j] {
		search329(i + 1, j)
		if longests329[i + 1][j] > mxToAdd {
			mxToAdd = longests329[i + 1][j]
		}
	}

	// self
	longests329[i][j] = mxToAdd + 1
}

func longestIncreasingPath329(matrix [][]int) int {
	h := len(matrix)
	if h == 0 {
		return 0
	}
	w := len(matrix[0])
	if w == 0 {
		return 0
	}

	matrix329 = matrix
	longests329 = make([][]int, h)
	for i := 0; i < h; i++ {
		longests329[i] = make([]int, w)
	}
	h329, w329 = h, w

	for i := 0; i < h; i++ {
		for j := 0; j < w; j++ {
			search329(i, j)
		}
	}

	mx := 1
	for i := 0; i < h; i++ {
		for j := 0; j < w; j++ {
			if longests329[i][j] > mx {
				mx = longests329[i][j]
			}
		}
	}

	return mx
}