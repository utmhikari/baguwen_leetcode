package solution_151_200

import "math"

//188. 买卖股票的最佳时机 IV
//给定一个整数数组prices ，它的第 i 个元素prices[i] 是一支给定的股票在第 i 天的价格。
//
//设计一个算法来计算你所能获取的最大利润。你最多可以完成 k 笔交易。
//
//注意：你不能同时参与多笔交易（你必须在再次购买前出售掉之前的股票）。



func max188(a int, b int) int {
	if a > b {
		return a
	}
	return b
}

type stat188 struct {
	buy int
	sell int
}


func maxProfit188(k int, prices []int) int {
	if k == 0 {
		return 0
	}
	l := len(prices)
	if l == 0 {
		return 0
	}

	buySells := make([]stat188, k)
	for i := 0; i < k; i++ {
		buySells[i] = stat188{
			buy: math.MinInt32,
			sell: 0,
		}
	}

	for i := 0; i < l; i++ {
		p := prices[i]
		for j := 0; j < k; j++ {
			if j == 0 {
				buySells[j].buy = max188(buySells[j].buy, -p)
				buySells[j].sell = max188(buySells[j].sell, buySells[j].buy + p)
			} else {
				buySells[j].buy = max188(buySells[j].buy, buySells[j-1].sell - p)
				buySells[j].sell = max188(buySells[j].sell, buySells[j].buy + p)
			}
		}
	}

	return buySells[k - 1].sell
}


