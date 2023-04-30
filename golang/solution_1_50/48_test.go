package solution_1_50

//48. 旋转图像
//给定一个 n×n 的二维矩阵matrix 表示一个图像。
//请你将图像顺时针旋转 90 度。

var tmp48 int

func r48(matrix [][]int, left int, right int) {
	// 每组4个成员：
	//(left, left + i),
	//(left + i, right),
	//(right, right - i),
	//(right - i, left)
	for i := 0; i < right - left; i++ {
		tmp48 = matrix[left][left + i]
		matrix[left][left + i] = matrix[right - i][left]
		matrix[right - i][left] = matrix[right][right - i]
		matrix[right][right - i] = matrix[left + i][right]
		matrix[left + i][right] = tmp48
	}

}

func rotate48(matrix [][]int)  {
	l := len(matrix)
	if l <= 1 {
		return
	}

	left, right := 0, l - 1
	for left < right {
		r48(matrix, left, right)
		left++
		right--
	}
}