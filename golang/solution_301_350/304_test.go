package solution_301_350



//304. 二维区域和检索 - 矩阵不可变
//给定一个二维矩阵，计算其子矩形范围内元素的总和，
//该子矩阵的左上角为 (row1, col1) ，右下角为 (row2, col2) 。


type NumMatrix struct {
	sums [][]int
}


func Constructor304(matrix [][]int) NumMatrix {
	h := len(matrix)
	sums := make([][]int, h + 1)

	for i := 0; i < h; i++ {
		w := len(matrix[i])
		if i == 0 && w > 0 {
			sums[0] = make([]int, w + 1)
		}
		sums[i + 1] = make([]int, w + 1)
		for j := 0; j < w; j++ {
			sums[i + 1][j + 1] = sums[i][j + 1] + sums[i + 1][j] - sums[i][j] + matrix[i][j]
		}
	}

	return NumMatrix{
		sums: sums,
	}
}


func (this *NumMatrix) SumRegion(row1 int, col1 int, row2 int, col2 int) int {
	return this.sums[row2 + 1][col2 + 1] + this.sums[row1][col1] -
		this.sums[row1][col2 + 1] - this.sums[row2 + 1][col1]
}
