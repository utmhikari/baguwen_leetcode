package solution_351_400


//383. 赎金信
//给定一个赎金信 (ransom) 字符串和一个杂志(magazine)字符串，
//判断第一个字符串 ransom 能不能由第二个字符串 magazines 里面的字符构成。
//如果可以构成，返回 true ；否则返回 false。


func canConstruct383(ransomNote string, magazine string) bool {
	charMap := make(map[rune]int)
	for _, r := range ransomNote {
		cnt, ok := charMap[r]
		if ok {
			charMap[r] = cnt + 1
		} else {
			charMap[r] = 1
		}
	}
	for _, r := range magazine {
		cnt, ok := charMap[r]
		if ok {
			if cnt == 1 {
				delete(charMap, r)
				if len(charMap) == 0 {
					break
				}
			} else {
				charMap[r] = cnt - 1
			}
		}
	}
	return len(charMap) == 0
}