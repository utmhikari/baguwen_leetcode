package solution_51_100

import (
	"reflect"
	"testing"
)

// 59. 螺旋矩阵 II
// 给你一个正整数 n ，生成一个包含 1 到 n2 所有元素，且元素按顺时针顺序螺旋排列的 n x n 正方形矩阵 matrix 。

func generateMatrix(n int) [][]int {
	var grid [][]int
	if n == 0 {
		return grid
	}

	grid = make([][]int, n)
	for i := 0; i < n; i++ {
		grid[i] = make([]int, n)
	}

	i := 0
	j := 0
	c := 0
	d := 4
	s := n * n
	for {
		c++
		grid[i][j] = c
		if c == s {
			break
		}
		switch d {
		case 1: // up
			if i == 0 || grid[i - 1][j] != 0 {
				d = 4
				j++
			} else {
				i--
			}
			break
		case 2: // down
			if i == n - 1 || grid[i + 1][j] != 0 {
				d = 3
				j--
			} else {
				i++
			}
			break
		case 3: // left
			if j == 0 || grid[i][j - 1] != 0 {
				d = 1
				i--
			} else {
				j--
			}
			break
		case 4: // right
			if j == n - 1 || grid[i][j + 1] != 0 {
				d = 2
				i++
			} else {
				j++
			}
			break
		default:
			break
		}
	}

	return grid
}


type test59 struct {
	n int
	expected [][]int
}

func (t *test59) getAns() [][]int {
	return generateMatrix(t.n)
}

func Test_59(t *testing.T) {
	cases := append(make([]test59, 0),
		test59{
			n: 3,
			expected: [][]int{
				{1,2,3},
				{8,9,4},
				{7,6,5},
			},
		},
		test59{
			n: 1,
			expected: [][]int{
				{1},
			},
		},
	)

	for _, cs := range cases {
		ans := cs.getAns()
		if !reflect.DeepEqual(ans, cs.expected) {
			t.Errorf("failed at %d -> %v (%v)", cs.n, ans, cs.expected)
		}
	}
}