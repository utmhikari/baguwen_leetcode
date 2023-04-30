package solution_151_200






//172. 阶乘后的零


func trailingZeroes172(n int) int {
	sm := 0
	base := 5
	for base <= n {
		sm += n / base
		base *= 5
	}
	return sm
}
