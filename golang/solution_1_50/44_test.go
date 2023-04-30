package solution_1_50

import "fmt"

//44. 通配符匹配
//
//给定一个字符串(s) 和一个字符模式(p) ，
//实现一个支持'?'和'*'的通配符匹配。
//
//'?' 可以匹配任何单个字符。
//'*' 可以匹配任意字符串（包括空字符串）。
//两个字符串完全匹配才算匹配成功。

// H, _, a, d, c, e, b
// _, 1, 0, 0, 0, 0, 0
// *, 1, 1, 1, 1, 1, 1
// a, 0, 1, 0, 0, 0, 0
// *, 0, 1, 1, 1, 1, 1
// b, 0, 0, 0, 0, 0, 1


// H, _, a, a
// _, 1, 0, 0
// a, 0, 1, 0


func isMatch44(s string, p string) bool {
	ls, lp := len(s), len(p)

	dp := make([][]bool, ls + 1)
	for i := 0; i < ls + 1; i++ {
		dp[i] = make([]bool, lp + 1)
	}
	dp[0][0] = true

	for i := 1; i < lp + 1; i++ {
		for j := 0; j < ls + 1; j++ {
			switch p[i - 1] {
			case '?':
				dp[j][i] = j != 0 && dp[j - 1][i - 1]
			case '*':
				dp[j][i] = dp[j][i - 1] || (j != 0 && (dp[j - 1][i] || dp[j - 1][i - 1]))
			default:
				dp[j][i] = j != 0 && dp[j - 1][i - 1] && p[i - 1] == s[j - 1]
			}
		}
	}

	fmt.Printf("%s, %s\n%v\n", s, p, dp)

	return dp[ls][lp]
}