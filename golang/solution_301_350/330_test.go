package solution_301_350


//330. 按要求补齐数组
//
//给定一个已排序的正整数数组 nums，和一个正整数 n 。
//从 [1, n] 区间内选取任意个数字补充到 nums 中，
//使得 [1, n] 区间内的任何数字都可以用 nums 中某几个数字的和来表示。
//请输出满足上述要求的最少需要补充的数字个数。


// [1, 5, 10] 31
//



func minPatches330(nums []int, n int) int {
	l := len(nums)
	ans := 0
	x := 1
	i := 0
	for x <= n {
		if i < l && nums[i] <= x {
			x += nums[i]
			i++
		} else {
			// found x that cannot be summed
			// should be patched with x, then next x will be x * 2
			// as 1 ~ 2x - 1 is covered
			x *= 2
			ans++
		}
	}
	return ans
}