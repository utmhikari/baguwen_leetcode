package solution_1_50

//31. 下一个排列
//实现获取 下一个排列 的函数，算法需要将给定数字序列重新排列成字典序中下一个更大的排列。
//
//如果不存在下一个更大的排列，则将数字重新排列成最小的排列（即升序排列）。
//
//必须 原地 修改，只允许使用额外常数空间。


func reverse31(nums[] int, left int, right int) {
	for left < right {
		nums[left], nums[right] = nums[right], nums[left]
		left++
		right--
	}
}


func nextPermutation31(nums []int)  {
	l := len(nums)
	if l <= 1 {
		return
	}

	// find the first one that larger than prev
	i := l - 1
	for ; i > 0; i-- {
		if nums[i] > nums[i - 1] {
			break
		}
	}

	// last permutation
	if i == 0 {
		reverse31(nums, 0, l - 1)
		return
	}

	if i == l - 1 {
		nums[i], nums[i - 1] = nums[i - 1], nums[i]
		return
	}

	// compare left & right
	if nums[i - 1] >= nums[i + 1] {
		// swap & reverse
		nums[i], nums[i - 1] = nums[i - 1], nums[i]
	} else {
		// find the last num larger than left, swap & reverse
		j := i + 1
		for j < l - 1 {
			if nums[j + 1] <= nums[i - 1] {
				break
			}
			j++
		}
		nums[i - 1], nums[j] = nums[j], nums[i - 1]
	}

	reverse31(nums, i, l - 1)
}
