package solution_1_50

//20. 有效的括号
//给定一个只包括 '('，')'，'{'，'}'，'['，']'的字符串 s ，判断字符串是否有效。
//
//有效字符串需满足：
//
//左括号必须用相同类型的右括号闭合。
//左括号必须以正确的顺序闭合。


func isValid20(s string) bool {
	l := len(s)
	if l % 2 == 1 {
		return false
	}

	stk := make([]int, 0)

	for i := 0; i < l; i++ {
		switch s[i] {
		case '(':
			stk = append(stk, 1)
			break
		case '[':
			stk = append(stk, 2)
			break
		case '{':
			stk = append(stk, 3)
			break
		case ')':
			if len(stk) == 0 || stk[len(stk) - 1] != 1 {
				return false
			}
			stk = stk[:len(stk) - 1]
			break
		case ']':
			if len(stk) == 0 || stk[len(stk) - 1] != 2 {
				return false
			}
			stk = stk[:len(stk) - 1]
			break
		case '}':
			if len(stk) == 0 || stk[len(stk) - 1] != 3 {
				return false
			}
			stk = stk[:len(stk) - 1]
			break
		default:
			return false
		}
	}

	return len(stk) == 0
}
