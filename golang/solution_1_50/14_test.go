package solution_1_50

import "bytes"

//14. 最长公共前缀
//编写一个函数来查找字符串数组中的最长公共前缀。
//如果不存在公共前缀，返回空字符串 ""。

func longestCommonPrefix14(strs []string) string {
	if len(strs) == 0 {
		return ""
	}

	var b bytes.Buffer
	i := 0
	shouldBreak := false
	for !shouldBreak {
		tc := ""
		for _, str := range strs {
			if i >= len(str) {
				shouldBreak = true
				break
			}

			c := str[i:i+1]
			if tc == "" {
				tc = c
			} else if c != tc {
				shouldBreak = true
				break
			}
		}
		if !shouldBreak {
			b.WriteString(tc)
			i++
		}
	}
	return b.String()
}