package solution_301_350

//318. 最大单词长度乘积
//给定一个字符串数组 words，找到 length(word[i]) * length(word[j]) 的最大值，
//并且这两个单词不含有公共字母。你可以认为每个单词只包含小写字母。
//如果不存在这样的两个单词，返回 0。


func maxProduct318(words []string) int {
	l := len(words)

	lens := make([]int, l)
	codes := make([]int, l)  // 26 chars exist?
	for i, w := range words {
		lens[i] = len(w)
		for j := 0; j < lens[i]; j++ {
			mask := 1 << (w[j] - 'a' + 1)
			if (mask & codes[i]) == 0 {
				codes[i] |= mask
			}
		}
	}

	mx := 0
	for i := 0; i < l - 1; i++ {
		for j := i + 1; j < l; j++ {
			if (codes[i] & codes[j]) == 0 {
				p := lens[i] * lens[j]
				if p > mx {
					mx = p
				}
			}
		}
	}

	return mx
}
