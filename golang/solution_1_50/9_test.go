package solution_1_50

import "strconv"

//9. 回文数
//给你一个整数 x ，如果 x 是一个回文整数，返回 true ；否则，返回 false 。
//
//回文数是指正序（从左向右）和倒序（从右向左）读都是一样的整数。例如，121 是回文，而 123 不是。

func isPalindrome9(x int) bool {
	if x == 0 {
		return true
	}
	if x < 0 {
		return false
	}

	s := strconv.Itoa(x)
	l := len(s)
	p1, p2 := 0, l - 1
	for p2 > p1 {
		if s[p1] != s[p2] {
			return false
		}
		p1++
		p2--
	}

	return true
}