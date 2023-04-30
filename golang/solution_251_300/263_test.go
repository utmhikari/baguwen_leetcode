package solution_251_300



//263. 丑数
//给你一个整数 n ，请你判断 n 是否为 丑数 。如果是，返回 true ；否则，返回 false 。
//
//丑数 就是只包含质因数 2、3 和/或 5 的正整数。


func solve263(n int) bool {
	if n == 1 {
		return true
	}

	if n % 5 == 0 {
		return solve263(n / 5)
	}

	if n % 3 == 0 {
		return solve263(n / 3)
	}

	if n % 2 == 0 {
		return solve263(n / 2)
	}

	return false
}


func isUgly263(n int) bool {
	if n <= 0 {
		return false
	}
	if n == 1 {
		return true
	}
	return solve263(n)
}