package solution_351_400

import "strconv"

//400. 第 N 位数字
//在无限的整数序列 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, ...中找到第 n 位数字。
//
//注意：n 是正数且在 32 位整数范围内（n < 231）。



func findNthDigit(n int) int {
	digitCnt := 1
	bottom, top := 0, 10
	for n > (top - bottom) * digitCnt {
		n -= (top - bottom) * digitCnt
		digitCnt++
		bottom, top = top, top * 10
	}

	num := n / digitCnt
	idx := n % digitCnt
	s := strconv.Itoa(num + bottom)
	return int(s[idx] - '0')
}
