package solution_1_50

import (
	"bytes"
	"fmt"
)

// 49. 字母异位词分组
// 给定一个字符串数组，将字母异位词组合在一起。字母异位词指字母相同，但排列不同的字符串。


var ans49 [][]string
var ansMap49 map[string][]string
var hasher49 []int


func hash49(s string) string {
	for i := 0; i < 26; i++ {
		hasher49[i] = 0
	}

	l := len(s)
	for i := 0; i < l; i++ {
		hasher49[int(s[i]) - 'a']++
	}

	var b bytes.Buffer
	for i := 0; i < 26; i++ {
		if hasher49[i] != 0 {
			b.WriteString(fmt.Sprintf("%d_%d-", i, hasher49[i]))
		}
	}
	return b.String()
}


func groupAnagrams49(strs []string) [][]string {
	ans49 = make([][]string, 0)
	l := len(strs)
	if l == 0 {
		return ans49
	}

	ansMap49 = make(map[string][]string)
	hasher49 = make([]int, 26)

	for _, s := range strs {
		hs := hash49(s)
		lst, ok := ansMap49[hs]
		if !ok {
			ansMap49[hs] = []string{s}
		} else {
			ansMap49[hs] = append(lst, s)
		}
	}

	for _, v := range ansMap49 {
		ans49 = append(ans49, v)
	}
	return ans49
}