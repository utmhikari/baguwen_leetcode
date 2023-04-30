package solution_351_400

import "fmt"

//
//392. 判断子序列
//给定字符串 s 和 t ，判断 s 是否为 t 的子序列。
//
//字符串的一个子序列是原始字符串删除一些（也可以不删除）字符而不改变剩余字符相对位置形成的新字符串。
//（例如，"ace"是"abcde"的一个子序列，而"aec"不是）。
//
//进阶：
//
//如果有大量输入的 S，称作 S1, S2, ... , Sk
//其中 k >= 10亿，你需要依次检查它们是否为 T 的子序列。在这种情况下，你会怎样改变代码？


var mp392 map[int]map[string]bool

func solve392(t string, rS []rune, iS int) bool {
	fmt.Printf("%s, %v, %d\n", t, rS, iS)
	if len(t) == 0 {
		return true
	}
	if len(t) > (len(rS) - iS) {
		return false
	}

	strSet, strSetOk := mp392[iS]
	if strSetOk {
		tOk, hasT := strSet[t]
		if hasT {
			return tOk
		}
	}

	ok := false
	for i := iS; i < len(rS); i++ {
		if rune(t[0]) == rS[i] {
			ok = solve392(t[1:], rS, i + 1)
			if ok {
				break
			}
		}
	}

	if !strSetOk {
		mp392[iS] = make(map[string]bool)
	}
	mp392[iS][t] = ok
	return ok
}

func isSubsequence(s string, t string) bool {
	mp392 = make(map[int]map[string]bool)
	runes := []rune(t)
	return solve392(s, runes, 0)
}