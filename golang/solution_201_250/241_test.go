package solution_201_250

import "strconv"

//241. 为运算表达式设计优先级
//
//
//给定一个含有数字和运算符的字符串，为表达式添加括号，
//改变其运算优先级以求出不同的结果。
//你需要给出所有可能的组合的结果。有效的运算符号包含 +,-以及*。


func solve241(expr string) []int {
	n, err := strconv.Atoi(expr)
	if err == nil {
		return []int{n}
	}

	ans := make([]int, 0)
	l := len(expr)
	for i := 0; i < l; i++ {
		r := expr[i]
		if r == '+' || r == '-' || r == '*' {
			left, right := solve241(expr[:i]), solve241(expr[i+1:])
			for _, leftSum := range left {
				for _, rightSum := range right {
					if r == '+' {
						ans = append(ans, leftSum + rightSum)
					} else if r == '-' {
						ans = append(ans, leftSum - rightSum)
					} else if r == '*' {
						ans = append(ans, leftSum * rightSum)
					}
				}
			}
		}
	}
	return ans
}

func diffWaysToCompute(expression string) []int {
	return solve241(expression)
}