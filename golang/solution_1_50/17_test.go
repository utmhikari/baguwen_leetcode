package solution_1_50

import "bytes"

//17. 电话号码的字母组合
//给定一个仅包含数字2-9的字符串，返回所有它能表示的字母组合。答案可以按 任意顺序 返回。
//给出数字到字母的映射如下（与电话按键相同）。注意 1 不对应任何字母。


var ans17 []string
var mp17 = map[string][]string{
	"2": {"a", "b", "c"},
	"3": {"d", "e", "f"},
	"4": {"g", "h", "i"},
	"5": {"j", "k", "l"},
	"6": {"m", "n", "o"},
	"7": {"p", "q", "r", "s"},
	"8": {"t", "u", "v"},
	"9": {"w", "x", "y", "z"},
}
var b17 bytes.Buffer


func dfs17(digits string, idx int) {
	if idx == len(digits) {
		ans17 = append(ans17, b17.String())
	} else {
		chars := mp17[digits[idx:idx+1]]
		for _, char := range chars {
			b17.WriteString(char)
			dfs17(digits, idx + 1)
			b17.Truncate(idx)
		}
	}
}


func letterCombinations17(digits string) []string {
	ans17 = make([]string, 0)
	if len(digits) > 0 {
		dfs17(digits, 0)
	}
	return ans17
}