package solution_51_100


//76. 最小覆盖子串
//给你一个字符串 s 、一个字符串 t 。返回 s 中涵盖 t 所有字符的最小子串。
//如果 s 中不存在涵盖 t 所有字符的子串，则返回空字符串 "" 。
//
//注意：如果 s 中存在这样的子串，我们保证它是唯一的答案。


var srcCounter76 map[uint8]int
var targetCounter76 map[uint8]int

func minWindow76(s string, t string) string {
	ls, lt := len(s), len(t)
	if ls < lt {
		return ""
	}
	if ls == 0 || lt == 0 {
		return ""
	}

	srcCounter76 = make(map[uint8]int)
	targetCounter76 = make(map[uint8]int)
	for i := 0; i < lt; i++ {
		c := t[i]
		cnt, ok := targetCounter76[c]
		if !ok {
			targetCounter76[c] = 1
			srcCounter76[c] = 0
		} else {
			targetCounter76[c] = cnt + 1
		}
	}

	left, right := 0, 0
	minLeft, minRight := -1, -1
	for right < ls {
		// add cnt
		c := s[right]
		cnt, ok := srcCounter76[c]
		if !ok {
			right++
			continue
		}
		srcCounter76[c] = cnt + 1

		// cover all?
		isCover := true
		for k, srcCnt := range srcCounter76 {
			targetCnt := targetCounter76[k]
			if srcCnt < targetCnt {
				isCover = false
				break
			}
		}

		if isCover {
			// should be minimal choice if length equal
			if right - left + 1 == lt {
				return s[left:right + 1]
			}

			// find the minimal length
			for i := left; i < right; i++ {
				cnt, ok := srcCounter76[s[i]]
				if ok {
					// decrease cnt
					srcCounter76[s[i]] = cnt - 1
					targetCnt := targetCounter76[s[i]]

					// found the last one to be the left
					if targetCnt > cnt - 1 {
						if minLeft == -1 || right - i < minRight - minLeft {
							minRight, minLeft = right, i
						}
						left = i + 1
						break
					}
				}
			}
		}

		right++
	}

	if minLeft == -1 {
		return ""
	}
	return s[minLeft:minRight + 1]
}