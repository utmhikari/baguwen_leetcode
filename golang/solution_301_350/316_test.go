package solution_301_350

//316. 去除重复字母


//给你一个字符串 s ，请你去除字符串中重复的字母，使得每个字母只出现一次。
//需保证 返回结果的字典序最小（要求不能打乱其他字符的相对位置）。






func removeDuplicateLetters316(s string) string {
	vis, num := make([]bool, 26), make([]int, 26)
	l := len(s)

	for i := 0; i < l; i++ {
		num[s[i] - 'a']++
	}

	bytes := make([]byte, 0)

	for i := 0; i < l; i++ {
		if !vis[s[i] - 'a'] {
			// 每个字符只出现一次
			// 因此如果有重复的字符出现，就尝试把前面的字符一个个的去掉，尽量让小的字符在前面
			for len(bytes) > 0 && s[i] < bytes[len(bytes) - 1] {
				last := bytes[len(bytes) - 1] - 'a'
				if num[last] == 0 {
					break
				}
				bytes = bytes[:len(bytes) - 1]
				vis[last] = false
			}
			bytes = append(bytes, s[i])
			vis[s[i] - 'a'] = true
		}
		num[s[i] - 'a']--
	}

	return string(bytes)
}