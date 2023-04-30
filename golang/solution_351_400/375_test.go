package solution_351_400

import "math"

//375. 猜数字大小 II

func getMoneyAmount(n int) int {
	dp := make([][]int, n + 1)
	for i := 0; i <= n; i++ {
		dp[i] = make([]int, n + 1)
	}

	var res int
	for i := 2; i <= n; i++ {
		for start := 1; start <= n - i + 1; start++ {
			minRes := math.MaxInt32
			for p := start; p < start + i - 1; p++ {
				if dp[start][p - 1] > dp[p + 1][start + i - 1] {
					res = p + dp[start][p - 1]
				} else {
					res = p + dp[p + 1][start + i - 1]
				}
				if res < minRes {
					minRes = res
				}
				dp[start][start + i - 1] = minRes
			}
		}
	}
	return dp[1][n]
}