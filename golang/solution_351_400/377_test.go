package solution_351_400

//377. 组合总和 Ⅳ
//给你一个由 不同 整数组成的数组 nums ，和一个目标整数 target 。
//请你从 nums 中找出并返回总和为 target 的元素组合的个数。
//
//题目数据保证答案符合 32 位整数范围。




func combinationSum4(nums []int, target int) int {
	dp := make([]int, target + 1)
	dp[0] = 1
	for i := 1; i <= target; i++ {
		for _, n := range nums {
			if n <= i {
				dp[i] += dp[i - n]
			}
		}
	}
	return dp[target]
}
