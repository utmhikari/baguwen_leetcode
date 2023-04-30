package solution_201_250


//242. 有效的字母异位词



func isAnagram242(s string, t string) bool {
	l := len(s)
	if len(t) != l {
		return false
	}
	ctr := make([]int, 26)
	for i := 0; i < l; i++ {
		ctr[s[i] - 'a']++
		ctr[t[i] - 'a']--
	}
	for i := 0; i < 26; i++ {
		if ctr[i] != 0 {
			return false
		}
	}
	return true
}