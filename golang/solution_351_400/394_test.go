package solution_351_400

import "bytes"

//394. 字符串解码
//给定一个经过编码的字符串，返回它解码后的字符串。
//
//编码规则为: k[encoded_string]，表示其中方括号内部的 encoded_string 正好重复 k 次。注意 k 保证为正整数。
//
//你可以认为输入字符串总是有效的；输入字符串中没有额外的空格，且输入的方括号总是符合格式要求的。
//
//此外，你可以认为原始数据不包含数字，所有的数字只表示重复的次数 k ，例如不会出现像 3a 或 2[4] 的输入。

func decode394(s string, i int) (string, int) {
	// is character only?
	var b bytes.Buffer
	j := i
	for j < len(s) {
		if s[j] == ']' {
			j++
			break
		}
		// chars
		if s[j] < '0' || s[j] > '9' {
			b.WriteRune(rune(s[j]))
			j++
		} else {
			sm := 0
			for s[j] != '[' {
				sm = sm * 10 + int(s[j] - '0')
				j++
			}
			nextStr, nextIdx := decode394(s, j + 1)
			for x := 0; x < sm; x++ {
				b.WriteString(nextStr)
			}
			j = nextIdx
		}
	}
	return b.String(), j
}


func decodeString(s string) string {
	ans, _ := decode394(s, 0)
	return ans
}