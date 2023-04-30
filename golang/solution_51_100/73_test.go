package solution_51_100

import (
	"fmt"
	"reflect"
	"testing"
)

//73. 矩阵置零
//给定一个 m x n 的矩阵，如果一个元素为 0，则将其所在行和列的所有元素都设为 0。请使用原地算法。


func setZeroes(matrix [][]int)  {
	h := len(matrix)
	if h == 0 {
		return
	}
	w := len(matrix[0])
	if w == 0 {
		return
	}

	sp := 47386249
	zCols := make(map[int]bool)

	for i := 0; i < h; i++ {
		hasZ := false
		for j := 0; j < w; j++ {
			if matrix[i][j] == sp {
				continue
			}
			if matrix[i][j] == 0 {
				hasZ = true
				_, ok := zCols[j]
				if !ok {
					zCols[j] = true
					for k := 0; k < h; k++ {
						if matrix[k][j] != 0 {
							matrix[k][j] = sp
						}
					}
				}
			}
		}
		if hasZ {
			for j := 0; j < w; j++ {
				if matrix[i][j] != 0 {
					matrix[i][j] = sp
				}
			}
		}
	}

	for i := 0; i < h; i++ {
		for j := 0; j < w; j++ {
			if matrix[i][j] == sp {
				matrix[i][j] = 0
			}
		}
	}
}


type test73 struct {
	matrix [][]int
	expected [][]int
}


func Test_73(t *testing.T) {
	cases := append(make([]test73, 0),
		test73{
			matrix: [][]int{
				{1,1,1},
				{1,0,1},
				{1,1,1},
			},
			expected: [][]int{
				{1,0,1},
				{0,0,0},
				{1,0,1},
			},
		},
		test73{
			matrix: [][]int{
				{0,1,2,0},
				{3,4,5,2},
				{1,3,1,5},
			},
			expected: [][]int{
				{0,0,0,0},
				{0,4,5,0},
				{0,3,1,0},
			},
		},
		test73{
			matrix: [][]int{
				{1,2,3,4},
				{5,0,7,8},
				{0,10,11,12},
				{13,14,15,0},
			},
			expected: [][]int{
				{0,0,3,0},
				{0,0,0,0},
				{0,0,0,0},
				{0,0,0,0},
			},
		},
		)

	for i, c := range cases {
		setZeroes(c.matrix)
		if !reflect.DeepEqual(c.matrix, c.expected) {
			t.Errorf("failed at %d -> ans77: %v, expected %v\n", i, c.matrix, c.expected)
		} else {
			fmt.Printf("success at %d -> ans77: %v\n", i, c.matrix)
		}
	}
}