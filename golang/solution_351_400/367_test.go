package solution_351_400




func isPerfectSquare(num int) bool {
	i := 1
	for {
		p := i * i
		if p == num {
			return true
		}
		if p > num {
			break
		}
		i++
	}
	return false
}