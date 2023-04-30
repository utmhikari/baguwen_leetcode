package solution_1_50



//28. 实现 strStr()
//实现strStr()函数。
//
//给你两个字符串haystack 和 needle ，
//请你在 haystack 字符串中找出 needle 字符串出现的第一个位置（下标从 0 开始）。
//如果不存在，则返回 -1



func strStr28(haystack string, needle string) int {
	lh, ln := len(haystack), len(needle)
	if ln == 0 {
		return 0
	}
	if ln > lh {
		return -1
	}

	// find all starts
	starts := make([]int, 0)
	for i := 0; i <= lh - ln; i++ {
		if haystack[i] == needle[0] {
			if ln == 1 {
				return i
			}
			starts = append(starts, i)
		}
	}
	ls := len(starts)
	if ls == 0 {
		return -1
	}

	// traverse starts
	lens := make([]int, ls)
	for i := 1; i < ln; i++ {
		for j := 0; j < ls;	j++ {
			if lens[j] == -1 {
				continue
			}
			if needle[i] != haystack[starts[j] + i] {
				lens[j] = -1
			} else if i == ln - 1 {
				return starts[j]
			}
		}
	}
	return -1
}


func strStr28Kmp(haystack string, needle string) int {
	lh, ln := len(haystack), len(needle)
	if ln == 0 {
		return 0
	}
	if ln > lh {
		return -1
	}

	// init kmp
	prefix := make([]int, ln)
	prefix[0] = 0
	for i := 1; i < ln; i++ {
		if needle[i] == needle[prefix[i - 1]] {
			prefix[i] = prefix[i - 1] + 1
		} else {
			j := i - 1
			for prefix[j] > 0 && needle[prefix[j]] != needle[i] {
				j = prefix[j] - 1
			}
			if prefix[j] > 0 {
				prefix[i] = prefix[j] + 1
			} else if needle[i] == needle[0] {
				prefix[i] = 1
			} else {
				prefix[i] = 0
			}
		}
	}

	// traverse
	i, j := 0, 0
	for ; i < lh && j < ln; i++{
		if needle[j] == haystack[i] {
			j++
		} else if j > 0 {
			for prefix[j - 1] > 0 && needle[prefix[j - 1]] != haystack[i] {
				j = prefix[j - 1]
			}
			if prefix[j - 1] > 0 {
				j = prefix[j - 1] + 1
			} else if needle[0] == haystack[i] {
				j = 1
			} else {
				j = 0
			}
		}
	}

	if j < ln {
		return -1
	}
	return i - ln
}