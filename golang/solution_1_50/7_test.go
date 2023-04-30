package solution_1_50

import (
	"fmt"
	"math"
)

//7. 整数反转
//给你一个 32 位的有符号整数 x ，返回将 x 中的数字部分反转后的结果。
//
//如果反转后整数超过 32 位的有符号整数的范围[−231, 231− 1] ，就返回 0。


func reverse7(x int) int {
	if x == 0 || x == math.MinInt32 {
		return 0
	}

	pos := x > 0
	if !pos {
		x = -x
	}

	sm := 0
	for x != 0 {
		m := x % 10
		nxt := sm * 10 + m
		fmt.Printf("x: %d, m: %d, nxt: %d\n", x, m, nxt)
		if nxt < sm || nxt > math.MaxInt32 {
			return 0
		}
		sm = nxt
		x /= 10
	}

	if !pos {
		sm = -sm
	}
	return sm
}