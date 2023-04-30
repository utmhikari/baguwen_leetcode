package solution_351_400



// 374. 猜数字大小


func guess(n int) int {
	return 0
}

func solve374(n int) int {
	left, right := 1, n
	for left < right {
		mid := (left + right) / 2
		g := guess(mid)
		if g == 0 {
			return mid
		}
		if g > 0 {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}
	return left
}


func guessNumber(n int) int {
	return solve374(n)
}