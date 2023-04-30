package solution_301_350

import "bytes"

//301. 删除无效的括号
//给你一个由若干括号和字母组成的字符串 s ，
//删除最小数量的无效括号，使得输入的字符串有效。
//返回所有可能的结果。答案可以按 任意顺序 返回。


var len301 int
var runes301 []rune
var ansSet301 map[string]bool
var buf301 bytes.Buffer


func solve301(idx int, leftCnt int, rightCnt int, leftRemove int, rightRemove int) {
	if idx == len301 {
		if leftRemove == 0 && rightRemove == 0 {
			ansSet301[buf301.String()] = true
		}
		return
	}

	r := runes301[idx]
	if r == '(' && leftRemove > 0 {
		solve301(idx + 1, leftCnt, rightCnt, leftRemove - 1, rightRemove)
	}
	if r == ')' && rightRemove > 0 {
		solve301(idx + 1, leftCnt, rightCnt, leftRemove, rightRemove - 1)
	}

	buf301.WriteRune(r)
	if r != '(' && r != ')' {
		solve301(idx + 1, leftCnt, rightCnt, leftRemove, rightRemove)
	} else if r == '(' {
		solve301(idx + 1, leftCnt + 1, rightCnt, leftRemove, rightRemove)
	} else if rightCnt < leftCnt {
		solve301(idx + 1, leftCnt, rightCnt + 1, leftRemove, rightRemove)
	}
	buf301.Truncate(buf301.Len() - 1)

}


func removeInvalidParentheses301(s string) []string {
	len301 = len(s)
	runes301 = []rune(s)

	leftRemove, rightRemove := 0, 0
	for i := 0; i < len301; i++ {
		if runes301[i] == '(' {
			leftRemove++
		} else if runes301[i] == ')' {
			if leftRemove == 0 {
				rightRemove++
			} else if leftRemove > 0 {
				leftRemove--
			}
		}
	}

	buf301 = bytes.Buffer{}
	ansSet301 = make(map[string]bool)
	solve301(0, 0, 0, leftRemove, rightRemove)
	ans := make([]string, 0)
	for k, _ := range ansSet301 {
		ans = append(ans, k)
	}
	return ans
}