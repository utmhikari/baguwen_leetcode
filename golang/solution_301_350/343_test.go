package solution_301_350

//343. 整数拆分
//给定一个正整数 n，将其拆分为至少两个正整数的和，
//并使这些整数的乘积最大化。 返回你可以获得的最大乘积。


var ansMap343 = map[int]int{
	1: 1,
}

func solve343(n int) int {
	ans, ok := ansMap343[n]
	if ok {
		return ans
	}

	mx := 1
	for i := 2; i <= (n / 2 + 1); i++ {
		d, m := n / i, n % i
		product := 1
		for j := 0; j < i; j++ {
			digit := d
			if m > 0 {
				digit++
				m--
			}
			nextAns := solve343(digit)
			if nextAns > digit {
				product *= nextAns
			} else {
				product *= digit
			}
		}
		if product > mx {
			mx = product
		}
	}

	ansMap343[n] = mx
	return mx
}

func integerBreak343(n int) int {
	ans := solve343(n)
	return ans
}

