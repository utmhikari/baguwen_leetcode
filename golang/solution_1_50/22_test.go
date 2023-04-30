package solution_1_50

import "bytes"

//22. 括号生成
//数字 n 代表生成括号的对数，请你设计一个函数，用于能够生成所有可能的并且 有效的 括号组合。

// 1. ["()"]
// 2. ["(())", "()()"]
// 3. ["((()))", "()(())", "(())()", "(()())", "()()()"]

var ans22 []string
var visited22 map[string]bool
var b22 bytes.Buffer

func solve22(n int, cl int, cr int) {
	s := b22.String()
	_, ok := visited22[s]
	if ok {
		return
	}
	visited22[s] = true

	if cl + cr == n * 2 {
		ans22 = append(ans22, s)
	} else {
		l := len(s)
		if cl < n {
			b22.WriteString("(")
			solve22(n, cl + 1, cr)
			b22.Truncate(l)
		}
		if cr < cl {
			b22.WriteString(")")
			solve22(n, cl, cr + 1)
			b22.Truncate(l)
		}
	}
}


func generateParenthesis22(n int) []string {
	ans22 = make([]string, 0)
	visited22 = make(map[string]bool)
	b22.Truncate(0)
	solve22(n, 0, 0)
	return ans22
}
