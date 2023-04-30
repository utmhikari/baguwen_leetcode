package solution_351_400


//376. 摆动序列
//如果连续数字之间的差严格地在正数和负数之间交替，则数字序列称为 摆动序列 。
//第一个差（如果存在的话）可能是正数或负数。仅有一个元素或者含两个不等元素的序列也视作摆动序列。


func max376(a, b int) int {
	if a > b {
		return a
	}
	return b
}


func wiggleMaxLength(nums []int) int {
	l := len(nums)
	if l <= 1 {
		return l
	}

	up, down := make([]int, l), make([]int, l)
	up[0] = 1
	down[0] = 1

	for i := 1; i < l; i++ {
		if nums[i] > nums[i - 1] {
			up[i] = max376(up[i - 1], down[i - 1] + 1)
			down[i] = down[i - 1]
		} else if nums[i] < nums[i - 1] {
			up[i] = up[i - 1]
			down[i] = max376(down[i - 1], up[i - 1] + 1)
		} else {
			up[i] = up[i - 1]
			down[i] = down[i - 1]
		}
	}

	return max376(up[l - 1], down[l - 1])
}