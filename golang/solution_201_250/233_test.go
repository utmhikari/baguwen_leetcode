package solution_201_250

//233. 数字 1 的个数
//给定一个整数 n，计算所有小于等于 n 的非负整数中数字 1 出现的个数。


func min233(a int, b int) int {
	if a < b {
		return a
	}
	return b
}

func max233(a int, b int) int {
	if a > b {
		return a
	}
	return b
}

func countDigitOne233(n int) int {
	cnt := 0
	for i := 1; i <= n; i *= 10 {
		d := i * 10
		cnt += (n / d) * i + min233(max233(n % d - i + 1, 0), i)
	}
	return cnt
}