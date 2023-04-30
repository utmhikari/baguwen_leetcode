package solution_101_150

import "strconv"

//150. 逆波兰表达式求值
//根据 逆波兰表示法，求表达式的值。
//有效的算符包括 +、-、*、/ 。每个运算对象可以是整数，也可以是另一个逆波兰表达式。



func evalRPN(tokens []string) int {
	stk := make([]int, 0)
	for _, token := range tokens {
		switch token {
		case "+":
			x, y := stk[len(stk) - 2], stk[len(stk) - 1]
			stk[len(stk) - 2] = x + y
			stk = stk[:len(stk) - 1]
			break
		case "-":
			x, y := stk[len(stk) - 2], stk[len(stk) - 1]
			stk[len(stk) - 2] = x - y
			stk = stk[:len(stk) - 1]
			break
		case "*":
			x, y := stk[len(stk) - 2], stk[len(stk) - 1]
			stk[len(stk) - 2] = x * y
			stk = stk[:len(stk) - 1]
			break
		case "/":
			x, y := stk[len(stk) - 2], stk[len(stk) - 1]
			stk[len(stk) - 2] = x / y
			stk = stk[:len(stk) - 1]
			break
		default:
			i, _ := strconv.Atoi(token)
			stk = append(stk, i)
			break
		}
	}
	return stk[0]
}