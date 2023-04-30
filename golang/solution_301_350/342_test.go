package solution_301_350

import "math"

func isPowerOfFour(n int) bool {
	b := 1
	for b <= math.MaxInt32 {
		if n == b {
			return true
		}
		b <<= 2
	}
	return false
}