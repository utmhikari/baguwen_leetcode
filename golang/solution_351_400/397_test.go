package solution_351_400


//397. 整数替换
//给定一个正整数 n ，你可以做如下操作：
//
//如果 n 是偶数，则用 n / 2替换 n 。
//如果 n 是奇数，则可以用 n + 1或n - 1替换 n 。
//n 变为 1 所需的最小替换次数是多少？



var ansMap397 = map[int]int{
	1: 0,
	2: 1,
	3: 2,
	4: 2,
}


func integerReplacement(n int) int {
	cached, ok := ansMap397[n]
	if ok {
		return cached
	}

	var ans int
	if n % 2 == 0 {
		ans = integerReplacement(n / 2)
	} else {
		left, right := integerReplacement(n - 1), integerReplacement(n + 1)
		if left < right {
			ans = left
		} else {
			ans = right
		}
	}
	ans++
	ansMap397[n] = ans
	return ans
}