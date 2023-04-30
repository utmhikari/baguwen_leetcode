package solution_251_300

import "math"

//279. 完全平方数
//
//给定正整数 n，找到若干个完全平方数（比如 1, 4, 9, 16, ...）
//使得它们的和等于 n。你需要让组成和的完全平方数的个数最少。
//
//给你一个整数 n ，返回和为 n 的完全平方数的 最少数量 。
//
//完全平方数 是一个整数，
//其值等于另一个整数的平方；换句话说，其值等于一个整数自乘的积。
//例如，1、4、9 和 16 都是完全平方数，而 3 和 11 不是。


var list279 []int

func ctor279() {
	n := 10000
	list279 = make([]int, n + 1)
	for i := 1; i <= n; i++ {
		minN := math.MaxInt32
		for j := 1; j * j <= i; j++ {
			left := i - j * j
			if list279[left] < minN {
				minN = list279[left]
			}
		}
		list279[i] = minN + 1
	}
}


func numSquares279(n int) int {
	if list279 == nil || len(list279) == 0 {
		ctor279()
	}
	return list279[n]
}