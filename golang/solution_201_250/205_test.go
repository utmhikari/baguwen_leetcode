package solution_201_250

//205. 同构字符串
//给定两个字符串s和t，判断它们是否是同构的。
//
//如果s中的字符可以按某种映射关系替换得到t，
//那么这两个字符串是同构的。
//
//每个出现的字符都应当映射到另一个字符，同时不改变字符的顺序。
//不同字符不能映射到同一个字符上，相同字符只能映射到同一个字符上，字符可以映射到自己本身。




func isIsomorphic212(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}
	l := len(s)
	charss, charst := []rune(s), []rune(t)
	mps, mpt := make(map[rune]int), make(map[rune]int)
	for i := 0; i < l; i++ {
		cs, ct := charss[i], charst[i]
		prevs, oks := mps[cs]
		prevt, okt := mpt[ct]
		if (oks && !okt) || (!oks && okt) {
			return false
		}
		if oks && okt && prevs != prevt {
			return false
		}
		mps[cs] = i
		mpt[ct] = i
	}
	return true
}