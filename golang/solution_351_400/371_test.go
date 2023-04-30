package solution_351_400

// 371. 两整数之和


func getSum371(a int, b int) int {
	for b != 0 {
		c := (int)((uint)(a & b) << 1)
		a ^= b
		b = c
	}
	return a
}