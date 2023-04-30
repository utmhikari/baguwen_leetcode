package solution_351_400

import "math"

//357. 计算各个位数不同的数字个数
//
//给定一个非负整数 n，计算各位数字都不同的数字 x 的个数，其中 0 ≤ x < 10n 。


func pow357(a int) int {
	return int(math.Pow10(a))
}

func countNumbersWithUniqueDigits(n int) int {
	dp := make([]int, n + 1)
	for i := 2; i <= n; i++ {
		dp[i] = dp[i - 1] * 10 + (9 * pow357(i - 2) - dp[i - 1]) * (i - 1)
	}
	sum := 0
	for _, a := range dp {
		sum += a
	}
	return pow357(n) - sum
}
