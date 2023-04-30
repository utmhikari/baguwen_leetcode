package solution_201_250


//239. 滑动窗口最大值
//给你一个整数数组 nums，有一个大小为k的滑动窗口
//从数组的最左侧移动到数组的最右侧
//。你只可以看到在滑动窗口内的 k个数字。滑动窗口每次只向右移动一位。


// 单调队列
// 如果i < j && nums[i] <= nums[j]，那么有索引i则没有nums[i]
func maxSlidingWindow239(nums []int, k int) []int {
	l := len(nums)
	q := make([]int, 0)

	// front k
	for i := 0; i < k; i++ {
		for len(q) > 0 && nums[i] >= nums[q[len(q) - 1]] {
			q = q[:len(q) - 1]
		}
		q = append(q, i)
	}

	// others
	ans := []int{nums[q[0]]}
	for i := k; i < l; i++ {
		for len(q) > 0 && nums[i] >= nums[q[len(q) - 1]] {
			q = q[:len(q) - 1]
		}
		q = append(q, i)

		for len(q) > 0 && q[0] <= i - k {
			q = q[1:]
		}
		ans = append(ans, nums[q[0]])
	}

	return ans
}