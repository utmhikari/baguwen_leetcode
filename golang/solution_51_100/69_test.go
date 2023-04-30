package solution_51_100




func mySqrt(x int) int {
	switch x {
	case 0:
		return 0
	case 1:
		return 1
	default:
		i := 1
		for {
			p := i * i
			if p == x {
				return i
			}
			if p > x {
				return i - 1
			}
			i++
		}
	}
}
