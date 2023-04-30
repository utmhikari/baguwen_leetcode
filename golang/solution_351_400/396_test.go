package solution_351_400


//396. 旋转函数
//给定一个长度为 n 的整数数组 A 。
//
//假设 Bk 是数组 A 顺时针旋转 k 个位置后的数组，我们定义 A 的“旋转函数” F 为：
//
//F(k) = 0 * Bk[0] + 1 * Bk[1] + ... + (n-1) * Bk[n-1]。
//
//计算F(0), F(1), ..., F(n-1)中的最大值。
//
//注意:
//可以认为 n 的值小于 10^5。


func maxRotateFunction(nums []int) int {
	// 0 -> 0
	// 1 -> nums[0] + nums[1] + ... + nums[n - 2] - (n - 1) * nums[n - 1]
	//    = nums[0] + ... + nums[n - 1] - n * nums[n - 1]
	// 2 -> nums[0] + ... + nums[n - 3] + nums[n - 1] - (n - 1) * nums[2]
	//    = nums[0] + ... + nums[n - 1] - n * nums[n - 2]
	l := len(nums)
	if l <= 1 {
		return 0
	}

	tmp := 0
	sm := 0
	for i, n := range nums {
		sm += n
		tmp += i * n
	}

	mx := tmp
	for i := 0; i < l - 1; i++ {
		tmp += sm - l * nums[l - 1 - i]
		if tmp > mx {
			mx = tmp
		}
	}

	return mx
}