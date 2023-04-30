package solution_1_50

//32. 给你一个只包含 '(' 和 ')' 的字符串，找出最长有效（格式正确且连续）括号子串的长度。


func longestValidParentheses32(s string) int {
	l := len(s)
	if l == 0 {
		return 0
	}

	mx := 0
	stk := make([]int, 0)
	lastPairs := make([][]int, 0)

	// ((())()
	for i := 0; i < l; i++ {
		if s[i] == '(' {
			stk = append(stk, i)
			for len(lastPairs) < len(stk) {
				lastPairs = append(lastPairs, []int{i, i - 1})
			}
		} else {
			lStk := len(stk)
			if lStk > 0 {
				if stk[lStk - 1] - 1 == lastPairs[lStk - 1][1] {
					lastPairs[lStk - 1][1] = i
				} else {
					lastPairs[lStk - 1] = []int{stk[lStk - 1], i}
				}
				sm := lastPairs[lStk - 1][1] - lastPairs[lStk - 1][0] + 1
				if sm > mx {
					mx = sm
				}
				stk = stk[:lStk - 1]
			}
		}
		//fmt.Printf("%d -> sm: %d, lastRights: %v, stk: %v\n", i, sm, lastRights, stk)
	}

	return mx
}