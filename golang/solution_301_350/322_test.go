package solution_301_350

import (
	"math"
	"sort"
)

//322. 零钱兑换
//给定不同面额的硬币 coins 和一个总金额 amount。
//编写一个函数来计算可以凑成总金额所需的最少的硬币个数。
//如果没有任何一种硬币组合能组成总金额，返回 -1。
//
//你可以认为每种硬币的数量是无限的。


var cache322 []int


func solve322(coins []int, amount int) int {
	if amount == 0 || cache322[amount] != 0 {
		return cache322[amount]
	}

	min := math.MaxInt32
	for _, n := range coins {
		next := amount - n
		if next >= 0 {
			nextAns := solve322(coins, next)
			if nextAns >= 0 && nextAns < min {
				min = nextAns + 1
			}
		}
	}

	if min == math.MaxInt32 {
		cache322[amount] = -1
	} else {
		cache322[amount] = min
	}
	return cache322[amount]
}


func coinChange322(coins []int, amount int) int {
	if amount == 0 {
		return 0
	}
	cache322 = make([]int, amount + 1)
	for _, n := range coins {
		cache322[n] = 1
	}
	sort.Ints(coins)
	solve322(coins, amount)

	return cache322[amount]
}