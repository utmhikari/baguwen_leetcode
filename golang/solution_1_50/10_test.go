package solution_1_50



//10. 正则表达式匹配
//给你一个字符串s和一个字符规律p，
//请你来实现一个支持 '.'和'*'的正则表达式匹配。
//
//'.' 匹配任意单个字符
//'*' 匹配零个或多个前面的那一个元素
//所谓匹配，是要涵盖整个字符串s的，而不是部分字符串。


func isMatch10(s string, p string) bool {
	ls, lp := len(s), len(p)
	dp := make([][]bool, ls + 1)
	for i := 0; i < ls + 1; i++ {
		dp[i] = make([]bool, lp + 1)
	}
	dp[0][0] = true

	for j := 2; j < lp + 1; j++ {
		if p[j - 1] == '*' {
			dp[0][j] = dp[0][j - 2]
		}
	}

	for i := 1; i < ls + 1; i++ {
		for j := 1; j < lp + 1; j++ {
			switch p[j - 1] {
			case '.':
				dp[i][j] = dp[i - 1][j - 1]
			case '*':
				dp[i][j] = dp[i][j - 2] || dp[i][j - 1]
				if !dp[i][j] {
					dp[i][j] = dp[i - 1][j] && (p[j - 2] == s[i - 1] || p[j - 2] == '.')
				}
			default:
				if p[j - 1] == s[i - 1] {
					dp[i][j] = dp[i - 1][j - 1]
				}
				break
			}
		}
	}
	//fmt.Printf("%+v\n", dp)
	return dp[ls][lp]
}
