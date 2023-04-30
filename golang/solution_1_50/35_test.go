package solution_1_50

//35. 搜索插入位置
//给定一个排序数组和一个目标值，在数组中找到目标值，并返回其索引。
//如果目标值不存在于数组中，返回它将会被按顺序插入的位置。
//你可以假设数组中无重复元素。


func searchInsert35(nums []int, target int) int {
	l := len(nums)
	if l == 0 {
		return 1
	}
	if target <= nums[0] {
		return 0
	}
	if target > nums[l - 1] {
		return l
	}

	left, right := 0, l - 1
	ans := l
	for left <= right {
		mid := left + (right - left) / 2
		if nums[mid] == target {
			return mid
		}
		if nums[mid] > target {
			ans = mid
			right = mid - 1
		} else {
			left = mid + 1
		}
	}
	return ans
}
