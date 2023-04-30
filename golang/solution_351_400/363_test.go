package solution_351_400

import (
	"fmt"
	"math"
)

//363. 矩形区域不超过 K 的最大数值和
//
//给你一个 m x n 的矩阵 matrix 和一个整数 k ，找出并返回矩阵内部矩形区域的不超过 k 的最大数值和。
//
//题目数据保证总会存在一个数值和不超过 k 的矩形区域。


var sums363 [][]int


func getSum363(h1, w1, h2, w2 int) int {
	return sums363[h2 + 1][w2 + 1] - sums363[h1][w2 + 1] - sums363[h2 + 1][w1] + sums363[h1][w1]
}

func maxSumSubmatrix(matrix [][]int, k int) int {
	h := len(matrix)
	if h == 0 {
		return 0
	}

	w := len(matrix[0])
	if w == 0 {
		return 0
	}

	sums363 = make([][]int, h + 1)
	for i := 0; i <= h; i++ {
		sums363[i] = make([]int, w + 1)
	}

	for i := 0; i < h; i++ {
		for j := 0; j < w; j++ {
			if matrix[i][j] == k {
				return k
			}
			sums363[i + 1][j + 1] = sums363[i][j + 1] + sums363[i + 1][j] - sums363[i][j] + matrix[i][j]
			if sums363[i + 1][j + 1] == k {
				return k
			}
		}
	}

	mx := math.MinInt32

	fmt.Printf("%v\n", sums363)

	for i1 := 0; i1 < h; i1++ {
		for i2 := i1; i2 < h; i2++ {
			for j1 := 0; j1 < w; j1++ {
				for j2 := j1; j2 < w; j2++ {
					sm := getSum363(i1, j1, i2, j2)
					if sm == k {
						return k
					}
					if sm > mx && sm < k {
						mx = sm
					}
				}
			}
		}
	}

	return mx
}
