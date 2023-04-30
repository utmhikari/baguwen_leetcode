package solution_1_50


//3. 无重复字符的最长子串
//给定一个字符串，请你找出其中不含有重复字符的 最长子串 的长度。


func lengthOfLongestSubstring3(s string) int {
	l := len(s)
	if l <= 1 {
		return l
	}

	cache := make(map[uint8]int)
	cnt := 0
	mx := 0
	for i := 0; i < l; i++ {
		idx, ok := cache[s[i]]
		if !ok {
			cache[s[i]] = i
			cnt += 1
		} else {
			if cnt > mx {
				mx = cnt
			}
			left := i - cnt
			for j := left; j < idx; j++ {
				delete(cache, s[j])
			}
			cache[s[i]] = i
			cnt = i - idx
		}
	}
	if cnt > mx {
		return cnt
	}
	return mx
}