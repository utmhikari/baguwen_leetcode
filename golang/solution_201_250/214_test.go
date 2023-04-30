package solution_201_250

//var pdCache214 map[string]bool


func isPalindrome214(s string, left int, right int) bool {
	if left > right {
		return false
	}

	//key := fmt.Sprintf("%d_%d", left, right)
	//b, ok := pdCache214[key]
	//if ok {
	//	return b
	//}

	for left < right {
		if s[left] != s[right] {
			//pdCache214[key] = false
			return false
		}
		left++
		right--
	}
	//pdCache214[key] = true
	return true
}


func shortestPalindrome(s string) string {
	l := len(s)
	if l == 0 {
		return s
	}
	if isPalindrome214(s, 0, l - 1) {
		return s
	}
	// left
	mxLeft := 1
	//mxRight := 1
	for i := l - 2; i >= 0; i-- {
		if isPalindrome214(s, 0, i) {
			mxLeft = i + 1
			break
		}
	}
	//for i := 1; i < l; i++ {
	//	if isPalindrome214(s, i, l - 1) {
	//		mxRight = l - i
	//      break
	//	}
	//}

	//if mxLeft > mxRight {
		runes := make([]rune, l - mxLeft)
		for i := mxLeft; i < l; i++ {
			runes[len(runes) - 1 - i + mxLeft] = rune(s[i])
		}
		return string(runes) + s
	//} else {
	//	runes := make([]rune, l - mxRight)
	//	for i := 0; i < l - mxRight; i++ {
	//		runes[len(runes) - 1 - i] = rune(s[i])
	//	}
	//	return s + string(runes)
	//}
}