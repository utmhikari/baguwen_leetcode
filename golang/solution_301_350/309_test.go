package solution_301_350

//309. 最佳买卖股票时机含冷冻期
//
//给定一个整数数组，其中第 i 个元素代表了第 i 天的股票价格 。​
//
//设计一个算法计算出最大利润。在满足以下约束条件下，你可以尽可能地完成更多的交易（多次买卖一支股票）:
//
//你不能同时参与多笔交易（你必须在再次购买前出售掉之前的股票）。
//卖出股票后，你无法在第二天买入股票 (即冷冻期为 1 天)。


func min309(a int, b int) int {
	if a < b {
		return a
	}
	return b
}

func max309(a int, b int) int {
	if a > b {
		return a
	}
	return b
}


func maxProfit309(prices []int) int {
	l := len(prices)
	if l < 2 {
		return 0
	}

	maxBoughts, maxSoldAtCoolings, maxSoldWithoutCoolings :=
		make([]int, l), make([]int, l), make([]int, l)
	maxBoughts[0] = -prices[0]

	for i := 1; i < l; i++ {
		maxBoughts[i] = max309(maxBoughts[i - 1], maxSoldWithoutCoolings[i - 1] - prices[i])
		maxSoldAtCoolings[i] = maxBoughts[i - 1] + prices[i]
		maxSoldWithoutCoolings[i] = max309(maxSoldAtCoolings[i - 1], maxSoldWithoutCoolings[i - 1])
	}

	return max309(maxSoldAtCoolings[l - 1], maxSoldWithoutCoolings[l - 1])
}

