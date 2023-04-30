package solution_1_50


//5.最长回文子串
//给你一个字符串 s，找到 s 中最长的回文子串。



func longestPalindrome5(s string) string {
	l := len(s)
	if l <= 1 {
		return s
	}

	mxLen, mxLeft, mxRight := 1, 0, 0
	for i := 0; i < l; i++ {
		// xxxxxaxxxxx
		left, right := i, i
		for left >= 0 && right < l && s[left] == s[right] {
			left--
			right++
		}
		left++
		right--
		lp := right - left + 1
		if lp > mxLen {
			mxLen = lp
			mxLeft = left
			mxRight = right
		}
		// xxxxxaaxxxxx
		if i < l - 1 && s[i] == s[i + 1] {
			left, right = i, i + 1
			for left >= 0 && right < l && s[left] == s[right] {
				left--
				right++
			}
			left++
			right--
			lp = right - left + 1
			if lp > mxLen {
				mxLen = lp
				mxLeft = left
				mxRight = right
			}
		}
	}

	return s[mxLeft:mxRight + 1]
}