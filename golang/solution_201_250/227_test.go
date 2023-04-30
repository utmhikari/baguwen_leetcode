package solution_201_250


//227. 基本计算器 II
//给你一个字符串表达式 s ，请你实现一个基本计算器来计算并返回它的值。
//
//整数除法仅保留整数部分。


func calculate227(s string) int {
	l := len(s)
	if l == 0 {
		return 0
	}

	pm := make([]int, 2)  // +, -
	md := make([]int, 3)  // *, /
	var mdSum int
	i := 0

	for i < l {
		// expression
		for i < l && s[i] != '+' && s[i] != '-' {
			switch s[i] {
			case ' ':
				break
			case '*':
				if md[2] == 1 {
					md[0] *= md[1]
					md[1] = 0
				} else if md[2] == 2 {
					md[0] /= md[1]
					md[1] = 0
				}
				md[2] = 1
			case '/':
				if md[2] == 1 {
					md[0] *= md[1]
					md[1] = 0
				} else if md[2] == 2 {
					md[0] /= md[1]
					md[1] = 0
				}
				md[2] = 2
			default:
				if md[2] == 0 {
					md[0] = 10 * md[0] + int(s[i] - '0')
				} else {
					md[1] = 10 * md[1] + int(s[i] - '0')
				}
			}
			i++
		}

		// get sum of md
		if md[2] == 0 {
			mdSum = md[0]
		} else if md[2] == 1 {
			mdSum = md[0] * md[1]
		} else {
			mdSum = md[0] / md[1]
		}
		md[0], md[1], md[2] = 0, 0, 0

		// update pm
		if pm[1] == 0 {
			pm[0] = mdSum
		} else if pm[1] == 1 {
			pm[0] += mdSum
		} else {
			pm[0] -= mdSum
		}

		if i < l {
			if s[i] == '+' {
				pm[1] = 1
			} else {
				pm[1] = 2
			}
			i++
		}
	}

	return pm[0]
}