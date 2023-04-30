package solution_201_250




func isPowerOfTwo231(n int) bool {
	if n <= 0 {
		return false
	}
	b := 1
	for i := 0; i < 31; i++ {
		if (n & b) == n {
			return true
		}
		b <<= 1
	}
	return false
}
