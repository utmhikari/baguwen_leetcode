package solution_351_400


//387. 字符串中的第一个唯一字符


func firstUniqChar(s string) int {
	indices, mp := make([]int, 0), make(map[rune]int)
	for i, r := range s {
		cnt, ok := mp[r]
		if ok {
			mp[r] = cnt + 1
		} else {
			indices = append(indices, i)
			mp[r] = 1
		}
	}
	for _, idx := range indices {
		if mp[rune(s[idx])] == 1 {
			return idx
		}
	}
	return -1
}