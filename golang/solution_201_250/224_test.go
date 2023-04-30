package solution_201_250

import (
	"fmt"
	"strings"
)

//224. 基本计算器
//
//给你一个字符串表达式 s ，请你实现一个基本计算器来计算并返回它的值。


func calculate224(s string) int {
	s = strings.Replace(s, " ", "", -1)
	if s[0] != '(' {
		s = "(" + s + ")"
	}

	stk := make([][]int, 0)
	for _, r := range s {
		switch r {
		case '(':
			// prev, cur, operator (0: nothing, 1: +, 2: -)
			stk = append(stk, []int{0, 0, 0})
			break
		case ')':
			i := len(stk) - 1
			op := stk[i][2]
			if op == 1 {
				stk[i][0] += stk[i][1]
				stk[i][1] = 0
				stk[i][2] = 0
			} else if op == 2 {
				stk[i][0] -= stk[i][1]
				stk[i][1] = 0
				stk[i][2] = 0
			}
			if i > 0 {
				if stk[i - 1][2] == 0 {
					stk[i - 1][0] = stk[i][0]
				} else {
					stk[i - 1][1] = stk[i][0]
				}
				stk = stk[:i]
			}
		case '+':
			i := len(stk) - 1
			op := stk[i][2]
			if op == 1 {
				stk[i][0] += stk[i][1]
				stk[i][1] = 0
			} else if op == 2 {
				stk[i][0] -= stk[i][1]
				stk[i][1] = 0
			}
			stk[i][2] = 1
			break
		case '-':
			i := len(stk) - 1
			op := stk[i][2]
			if op == 1 {
				stk[i][0] += stk[i][1]
				stk[i][1] = 0
			} else if op == 2 {
				stk[i][0] -= stk[i][1]
				stk[i][1] = 0
			}
			stk[i][2] = 2
			break
		default:
			i := len(stk) - 1
			op := stk[i][2]
			if op == 0 {
				stk[i][0] = 10 * stk[i][0] + int(r - '0')
			} else {
				stk[i][1] = 10 * stk[i][1] + int(r - '0')
			}
			break
		}
		fmt.Printf("stk: %v\n", stk)
	}
	if stk[0][2] == 1 {
		return stk[0][0] + stk[0][1]
	} else if stk[0][2] == 2 {
		return stk[0][0] - stk[0][1]
	} else {
		return stk[0][0]
	}
}