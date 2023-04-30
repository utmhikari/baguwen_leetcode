package solution_101_150

import "math"

//121. 买卖股票的最佳时机
//给定一个数组 prices ，它的第i 个元素prices[i] 表示一支给定股票第 i 天的价格。
//
//你只能选择 某一天 买入这只股票，并选择在 未来的某一个不同的日子 卖出该股票。
//设计一个算法来计算你所能获取的最大利润。
//
//返回你可以从这笔交易中获取的最大利润。如果你不能获取任何利润，返回 0 。


func maxProfit121(prices []int) int {
	mx := 0
	l := len(prices)
	if l < 2 {
		return mx
	}
	minPrice := math.MaxInt32
	for i := 0; i < l; i++ {
		if prices[i] < minPrice {
			minPrice = prices[i]
		} else {
			profit := prices[i] - minPrice
			if profit > mx {
				mx = profit
			}
		}
	}
	return mx
}
