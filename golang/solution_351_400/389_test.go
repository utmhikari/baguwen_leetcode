package solution_351_400


//389. 找不同

func findTheDifference(s string, t string) byte {
	mp := make(map[rune]int)
	for _, r := range s {
		c, ok := mp[r]
		if ok {
			mp[r] = c + 1
		} else {
			mp[r] = 1
		}
	}
	for _, r := range t {
		c, ok := mp[r]
		if !ok || c == 0 {
			return byte(r)
		} else {
			mp[r] = c - 1
		}
	}
	return 0
}