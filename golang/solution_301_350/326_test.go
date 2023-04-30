package solution_301_350






//326. 3的幂
//给定一个整数，写一个函数来判断它是否是 3 的幂次方。
//如果是，返回 true ；否则，返回 false 。
//
//整数 n 是 3 的幂次方需满足：存在整数 x 使得 n == 3x


func solve326(n int) bool {
	if n == 0 {
		return false
	}
	if n == 1 {
		return true
	}
	if n % 3 != 0 {
		return false
	}
	return solve326(n / 3)
}


func isPowerOfThree326(n int) bool {
	return solve326(n)
}