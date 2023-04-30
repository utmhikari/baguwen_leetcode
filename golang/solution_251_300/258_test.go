package solution_251_300


//258. 各位相加


func addDigits(num int) int {
	if num >= 0 && num <= 9 {
		return num
	}

	sm := 0
	for num != 0 {
		sm += num % 10
		num /= 10
	}
	return addDigits(sm)
}