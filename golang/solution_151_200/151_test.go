package solution_151_200

import "bytes"

//151. 翻转字符串里的单词
//给定一个字符串，逐个翻转字符串中的每个单词。

//说明：
//
//无空格字符构成一个 单词 。
//输入字符串可以在前面或者后面包含多余的空格，但是反转后的字符不能包括。
//如果两个单词间有多余的空格，将反转后单词间的空格减少到只含一个。


func reverseWords(s string) string {
	cache := make([][]int, 0)
	flag := 0
	i := 0
	l := len(s)

	for {
		if i == l || s[i:i+1] == " " {
			if flag == 1 {
				cache[len(cache)-1][1] = i - 1
				flag = 0
			}
			if i == l {
				break
			}
		} else {
			if flag == 0 {
				cache = append(cache, []int{i,-1})
				flag = 1
			}
		}
		i++
	}

	var b bytes.Buffer

	for j := len(cache) - 1; j >= 0; j-- {
		if j != len(cache)-1 {
			b.WriteString(" ")
		}
		b.WriteString(s[cache[j][0]:cache[j][1] + 1])
	}

	return b.String()
}