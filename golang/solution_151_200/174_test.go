package solution_151_200

import (
	"math"
	"testing"
)

func min174(a int, b int) int {
	if a < b {
		return a
	}
	return b
}

func max174(a int, b int) int {
	if a > b {
		return a
	}
	return b
}


func calculateMinimumHP174(dungeon [][]int) int {
	h := len(dungeon)
	if h == 0 {
		return 0
	}

	w := len(dungeon[0])
	if w == 0 {
		return 0
	}

	dp := make([][]int, h + 1)
	for i := 0; i <= h; i++ {
		dp[i] = make([]int, w + 1)
		for j := 0; j <= w; j++ {
			dp[i][j] = math.MaxInt32
		}
	}

	dp[h][w - 1], dp[h - 1][w] = 1, 1

	for i := h - 1; i >= 0; i-- {
		for j := w - 1; j >= 0; j-- {
			minNum := min174(dp[i + 1][j], dp[i][j + 1])
			dp[i][j] = max174(minNum - dungeon[i][j], 1)
		}
	}
	return dp[0][0]
}


type test174 struct {
	d [][]int
	expected int
}


func Test_174(t *testing.T) {
	inputs := []test174{
		{[][]int{
			{-2, -3, 3},
			{-5, -10, 1},
			{10, 30, -5},
		}, 7},
		{[][]int{{0, -3}}, 4},
		{[][]int{
			{1, -3, 3},
			{0, -2, 0},
			{-3, -3, -3},
		}, 3},
	}
	for i := 0; i < len(inputs); i++ {
		ans := calculateMinimumHP174(inputs[i].d)
		if ans != inputs[i].expected {
			t.Errorf("failed at %d, expected %d, got %d\n", i, inputs[i].expected, ans)
		} else {
			t.Logf("success at %d -> %d\n", i, ans)
		}
	}
}