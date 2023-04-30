package solution_101_150



//122. 买卖股票的最佳时机 II
//给定一个数组 prices ，其中prices[i] 是一支给定股票第 i 天的价格。
//设计一个算法来计算你所能获取的最大利润。你可以尽可能地完成更多的交易（多次买卖一支股票）。
//注意：你不能同时参与多笔交易（你必须在再次购买前出售掉之前的股票）。



func maxProfit(prices []int) int {
	l := len(prices)
	if l < 2 {
		return 0
	}
	i := 0
	// first i
	for i + 1 < l && prices[i + 1] <= prices[i] {
		i++
	}
	if i == l - 1 {
		return 0
	}
	// start visit
	mx := 0
	curMin := i
	for i < l {
		for i + 1 < l && prices[i + 1] > prices[i] {
			i++
		}
		mx += prices[i] - prices[curMin]
		for i + 1 < l && prices[i + 1] <= prices[i] {
			i++
		}
		curMin = i
		if i == l - 1 {
			break
		}
	}
	return mx
}