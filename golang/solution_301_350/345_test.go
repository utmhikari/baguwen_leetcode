package solution_301_350


//345. 反转字符串中的元音字母


func isVowel345(b byte) bool {
	return b == 'a' || b == 'e' || b == 'i' || b == 'o' || b == 'u' ||
		b == 'A' || b == 'E' || b == 'I' || b == 'O' || b == 'U'
}

func reverseVowels345(s string) string {
	i, j := 0, len(s) - 1
	bs := []byte(s)
	for i < j {
		for i < j {
			if !isVowel345(s[i]) {
				i++
			} else {
				break
			}
		}
		for i < j {
			if !isVowel345(s[j]) {
				j--
			} else {
				break
			}
		}
		if i >= j {
			break
		}
		bs[i], bs[j] = bs[j], bs[i]
		i++
		j--
	}
	return string(bs)
}
