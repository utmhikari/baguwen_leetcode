package solution_51_100

import (
	"testing"
)

func minPathSum(grid [][]int) int {
	h := len(grid)
	if h == 0 {
		return 0
	}
	w := len(grid[0])
	if w == 0 {
		return 0
	}

	dp := make([][]int, h)
	for i := 0; i < h; i++ {
		dp[i] = make([]int, w)
	}

	for i := 0; i < h; i++ {
		for j := 0; j < w; j++ {
			if i == 0 && j == 0 {
				dp[i][j] = grid[i][j]
			} else if i == 0 {
				dp[i][j] = dp[i][j - 1] + grid[i][j]
			} else if j == 0 {
				dp[i][j] = dp[i - 1][j] + grid[i][j]
			} else {
				if dp[i - 1][j] < dp[i][j - 1] {
					dp[i][j] = dp[i - 1][j] + grid[i][j]
				} else {
					dp[i][j] = dp[i][j - 1] + grid[i][j]
				}
			}
		}
	}

	return dp[h - 1][w - 1]
}

func Test_64(t *testing.T) {
	var grid [][]int
	var ans int
	var expected int

	grid = [][]int{
		{1,3,1},
		{1,5,1},
		{4,2,1},
	}
	ans = minPathSum(grid)
	expected = 7
	if ans != expected {
		t.Errorf("failed at 1 -> %d (%d)", ans, expected)
	}

	grid = [][]int{
		{1,2,3},
		{4,5,6},
	}
	ans = minPathSum(grid)
	expected = 12
	if ans != expected {
		t.Errorf("failed at 2 -> %d (%d)", ans, expected)
	}
}