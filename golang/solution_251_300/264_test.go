package solution_251_300

//264. 丑数 II
//
//给你一个整数 n ，请你找出并返回第 n 个 丑数 。
//
//丑数 就是只包含质因数 2、3 和/或 5 的正整数。



func nthUglyNumber264(n int) int {
	dp := make([]int, n + 1)
	dp[1] = 1
	p2, p3, p5 := 1, 1, 1
	for i := 2; i <= n; i++ {
		n2, n3, n5 := dp[p2] * 2, dp[p3] * 3, dp[p5] * 5
		dp[i] = n2
		if n3 < dp[i] {
			dp[i] = n3
		}
		if n5 < dp[i] {
			dp[i] = n5
		}

		if dp[i] == n2 {
			p2++
		}
		if dp[i] == n3 {
			p3++
		}
		if dp[i] == n5 {
			p5++
		}
	}

	return dp[n]
}