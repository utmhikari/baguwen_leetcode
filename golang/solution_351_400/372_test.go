package solution_351_400



//372. 超级次方
//你的任务是计算 ab 对 1337 取模，a 是一个正整数，b 是一个非常大的正整数且会以数组形式给出。

var base372 = 1337

func pow372(a, b int) int {
	if b == 0 {
		return 1
	}
	a %= base372

	if b % 2 == 1 {
		return (a * pow372(a, b - 1)) % base372
	} else {
		sub := pow372(a, b / 2)
		return (sub * sub) % base372
	}
}


func solve372(a int, b []int) int {
	if len(b) == 0 {
		return 1
	}

	last := b[len(b) - 1]
	b = b[:len(b) - 1]

	return (pow372(a, last) * pow372(solve372(a, b), 10)) % base372
}


func superPow372(a int, b []int) int {
	return solve372(a, b)
}