package solution_251_300

import "strings"

//290. 单词规律
//给定一种规律 pattern 和一个字符串 str ，
//判断 str 是否遵循相同的规律。
//
//这里的 遵循 指完全匹配，例如， pattern 里的
//每个字母和字符串 str 中的每个非空单词之间存在着双向连接的对应规律。



func wordPattern290(pattern string, s string) bool {
	splits := strings.Split(s, " ")
	l := len(pattern)
	if l != len(splits) {
		return false
	}

	mpStrByte := make(map[string]byte)
	mpByteStr := make(map[byte]string)

	for i := 0; i < l; i++ {
		p := pattern[i]
		str := splits[i]

		expectedP, ok := mpStrByte[str]
		if ok {
			if expectedP != p {
				return false
			}
		} else {
			mpStrByte[str] = p
		}

		expectedStr, ok := mpByteStr[p]
		if ok {
			if expectedStr != str {
				return false
			}
		} else {
			mpByteStr[p] = str
		}
	}

	return true
}