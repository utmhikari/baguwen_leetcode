package solution_1_50

import "math"

//29. 两数相除
//给定两个整数，被除数dividend和除数divisor。
//将两数相除，要求不使用乘法、除法和 mod 运算符。
//
//返回被除数dividend除以除数divisor得到的商。
//
//整数除法的结果应当截去（truncate）其小数部分，
//例如：truncate(8.345) = 8 以及 truncate(-2.7335) = -2


func divide29(dividend int, divisor int) int {
	if dividend == 0 {
		return 0
	}
	if dividend == math.MinInt32 && divisor == -1 {
		return math.MaxInt32
	}

	isBelowZero := (dividend < 0 && divisor > 0) || (dividend > 0 && divisor < 0)
	if divisor < 0 {
		divisor = -divisor
	}
	if dividend < 0 {
		dividend = -dividend
	}

	pow2s := make([][]int, 0)
	tmp, base := divisor, 1
	for tmp <= dividend {
		pow2s = append(pow2s, []int{tmp, base})
		tmp <<= 1
		base <<= 1
	}

	d := dividend
	sm := 0
	for i := len(pow2s) - 1; i >= 0; i-- {
		if d >= pow2s[i][0] {
			d -= pow2s[i][0]
			sm += pow2s[i][1]
			if d == 0 {
				break
			}
		}
	}

	if isBelowZero {
		sm = -sm
	}
	return sm
}