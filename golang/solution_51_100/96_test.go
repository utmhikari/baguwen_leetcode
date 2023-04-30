package solution_51_100

//96. 不同的二叉搜索树

import "fmt"

var map96 map[string]int


func solve96(left int, right int) int{
	if left > right {
		return 0
	}
	if left == right {
		return 1
	}
	key := fmt.Sprintf("%d_%d", left, right)
	ans, ok := map96[key]
	if ok {
		return ans
	}

	ans = 0
	for i := left; i <= right; i++ {
		leftAns := solve96(left, i - 1)
		rightAns := solve96(i + 1, right)
		if leftAns == 0 && rightAns == 0 {
			ans += 1
		} else if leftAns == 0 {
			ans += rightAns
		} else if rightAns == 0 {
			ans += leftAns
		} else {
			ans += leftAns * rightAns
		}
	}
	map96[key] = ans
	return ans
}

func numTrees96(n int) int {
	if n == 0 {
		return 0
	}
	if n == 1 {
		return 1
	}
	map96 = make(map[string]int)
	return solve96(1, n)
}
