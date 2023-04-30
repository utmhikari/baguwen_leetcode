package solution_101_150

import "testing"

//132. 分割回文串 II
//给你一个字符串 s，请你将 s 分割成一些子串，使每个子串都是回文。
//返回符合要求的 最少分割次数 。



var minCutCache132 []int
var visitedCache132 []bool


func isPalindrome132(start int, end int, s string) bool {
	p1, p2 := start, end
	for p2 > p1 {
		if s[p2] != s[p1] {
			return false
		}
		p2--
		p1++
	}
	return true
}


func findPalindromeEnds132(start int, s string) *[]int {
	ends := make([]int, 0)
	l := len(s)
	for i := l - 1; i >= start; i-- {
		if isPalindrome132(start, i, s) {
			ends = append(ends, i)
		}
	}
	return &ends
}

func visit132(start int, s string) {
	if visitedCache132[start] {
		return
	}
	visitedCache132[start] = true

	l := len(s)
	if start >= l {
		return
	}

	ends := findPalindromeEnds132(start, s)

	for _, end := range *ends {
		if end == l - 1 {
			minCutCache132[start] = 0
			return
		}
		visit132(end + 1, s)
		times := minCutCache132[end + 1] + 1
		if times < minCutCache132[start] {
			minCutCache132[start] = times
		}
	}
}

func minCut132(s string) int {
	l := len(s)
	if l <= 1 {
		return 0
	}

	minCutCache132 = make([]int, l)
	visitedCache132 = make([]bool, l)
	for i := 0; i < l; i++ {
		minCutCache132[i] = l - 1 - i
	}
	visit132(0, s)
	return minCutCache132[0]
}


type test132 struct {
	s string
	expected int
}


func Test_132(t *testing.T) {
	inputs := []test132{
		{
			"apjesgpsxoeiokmqmfgvjslcjukbqxpsobyhjpbgdfruqdkeiszrlmtwgfxyfostpqczidfljwfbbrflkgdvtytbgqalguewnhvvmcgxboycffopmtmhtfizxkmeftcucxpobxmelmjtuzigsxnncxpaibgpuijwhankxbplpyejxmrrjgeoevqozwdtgospohznkoyzocjlracchjqnggbfeebmuvbicbvmpuleywrpzwsihivnrwtxcukwplgtobhgxukwrdlszfaiqxwjvrgxnsveedxseeyeykarqnjrtlaliyudpacctzizcftjlunlgnfwcqqxcqikocqffsjyurzwysfjmswvhbrmshjuzsgpwyubtfbnwajuvrfhlccvfwhxfqthkcwhatktymgxostjlztwdxritygbrbibdgkezvzajizxasjnrcjwzdfvdnwwqeyumkamhzoqhnqjfzwzbixclcxqrtniznemxeahfozp",
			452,
		},
		{
			"aab",
			1,
		},
		{
			"a",
			0,
		},
		{
			"ab",
			1,
		},
	}

	for i, input := range inputs {
		ans := minCut132(input.s)
		if ans != input.expected {
			t.Errorf("failed at %d -> %+v, got %v\n", i, input, ans)
		} else {
			t.Logf("success at %d -> %+v\n", i, input)
		}
	}
}