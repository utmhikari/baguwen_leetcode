package solution_101_150

import "strings"

//125. 验证回文串
//给定一个字符串，验证它是否是回文串，只考虑字母和数字字符，可以忽略字母的大小写。

func isValidChar125(s string) bool {
	return (s >= "a" && s <= "z") || (s >= "0" && s <= "9")
}

func isPalindrome125(s string) bool {
	if len(s) <= 1 {
		return true
	}
	s = strings.ToLower(s)

	i, j := 0, len(s) - 1

	for {
		for i < j && !isValidChar125(s[i:i+1]) {
			i++
		}

		if j <= i {
			return true
		}

		for i < j && !isValidChar125(s[j:j+1]) {
			j--
		}
		//fmt.Printf("%d, %d, %s, %s\n", i, j, s[i:i+1], s[j:j+1])
		if s[i:i+1] != s[j:j+1] {
			return false
		}

		i++
		j--
	}

}