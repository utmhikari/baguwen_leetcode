package solution_1_50


//26. 删除有序数组中的重复项
//给你一个有序数组 nums ，请你 原地 删除重复出现的元素，使每个元素 只出现一次 ，
//返回删除后数组的新长度。
//
//不要使用额外的数组空间，你必须在 原地 修改输入数组
//并在使用 O(1) 额外空间的条件下完成。

func removeDuplicates26(nums []int) int {
	l := len(nums)
	if l <= 1 {
		return l
	}

	n := nums[0]
	offset := 0
	i := 1
	cnt := 0
	for i < l {
		for i < l && nums[i] == n {
			i++
			offset++
		}
		if i >= l {
			break
		}
		n = nums[i]
		nums[i], nums[i - offset] = nums[i - offset], nums[i]
		cnt++
		i++
	}

	return cnt + 1
}