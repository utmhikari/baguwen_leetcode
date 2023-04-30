package solution_101_150

import (
	"fmt"
	"testing"
)

//115. 不同的子序列
//给定一个字符串 s 和一个字符串 t ，计算在 s 的子序列中 t 出现的个数。
//字符串的一个 子序列 是指，通过删除一些（也可以不删除）字符且不干扰剩余字符相对位置所组成的新字符串。（例如，"ACE"是"ABCDE"的一个子序列，而"AEC"不是）
//
//题目数据保证答案符合 32 位带符号整数范围。


func numDistinct115(s string, t string) int {
	ls := len(s)
	if ls == 0 {
		return 0
	}
	lt := len(t)
	if lt == 0 {
		return 0
	}

	dp := make([][]int, ls + 1)
	for i := 0; i < ls + 1; i++ {
		dp[i] = make([]int, lt + 1)
		dp[i][0] = 1
	}

	for i := 1; i < ls + 1; i++ {
		for j := 1; j < lt + 1; j++ {
			if s[i-1:i] == t[j-1:j] {
				dp[i][j] = dp[i-1][j] + dp[i-1][j-1]
			} else {
				dp[i][j] = dp[i-1][j]
			}
		}
	}

	fmt.Printf("%+v\n", dp)

	return dp[ls][lt]
}

type test115 struct {
	s string
	t string
	expected int
}

func Test_115(t *testing.T) {
	inputs := []test115{
		{"rabbbit", "rabbit", 3},
		{"babgbag", "bag", 5},
	}

	for i, input := range inputs {
		ans := numDistinct115(input.s, input.t)
		if ans != input.expected {
			t.Errorf("failed at %d (%d) -> %+v\n", i, ans, input)
		} else {
			t.Logf("success at %d -> %+v\n", i, input)
		}
	}

}






