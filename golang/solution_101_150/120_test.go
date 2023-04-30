package solution_101_150

import (
	"math"
)


func minInt120(i1 int, i2 int) int {
	if i1 < i2 {
		return i1
	}
	return i2
}


func solve120(r int, c int, triangle [][]int, minValues [][]int) {
	if r == len(triangle) - 1 {
		minValues[r][c] = minInt120(minValues[r][c], triangle[r][c])
	} else {
		if c == 0 {
			solve120(r + 1, c, triangle, minValues)
		}
		solve120(r + 1, c + 1, triangle, minValues)
		minValues[r][c] = minInt120(minValues[r + 1][c], minValues[r + 1][c + 1]) + triangle[r][c]
	}
}

func minimumTotal120(triangle [][]int) int {
	l := len(triangle)
	if l == 0 {
		return 0
	}

	minValues := make([][]int, l)
	for i := 0; i < l; i++ {
		minValues[i] = make([]int, i + 1)
		for j := 0; j < i + 1; j++ {
			minValues[i][j] = math.MaxInt32
		}
	}
	solve120(0, 0, triangle, minValues)
	return minValues[0][0]
}