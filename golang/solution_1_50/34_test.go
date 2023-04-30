package solution_1_50


//34. 在排序数组中查找元素的第一个和最后一个位置
//给定一个按照升序排列的整数数组 nums，和一个目标值 target。
//找出给定目标值在数组中的开始位置和结束位置。
//
//如果数组中不存在目标值 target，返回[-1, -1]。




func searchRange34(nums []int, target int) []int {
	ans := []int{-1, -1}
	l := len(nums)
	if l == 0 {
		return ans
	}

	left, right := 0, l - 1
	for right >= left {
		mid := left + (right - left) / 2
		if nums[mid] == target {
			i, j := mid, mid
			for ; i >= 0; i-- {
				if nums[i] != target {
					break
				}
			}
			for ; j < l; j++ {
				if nums[j] != target {
					break
				}
			}
			ans[0], ans[1] = i + 1, j - 1
			break
		} else if nums[mid] > target {
			right = mid - 1
		} else {
			left = mid + 1
		}
	}
	return ans
}