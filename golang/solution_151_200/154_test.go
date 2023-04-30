package solution_151_200



//154. 寻找旋转排序数组中的最小值 II
//假设按照升序排序的数组在预先未知的某个点上进行了旋转。

func findMin254(nums []int) int {
	l := len(nums)
	if l == 1 || nums[0] < nums[l - 1] {
		return nums[0]
	}

	for i := l - 1; i > 0; i-- {
		if nums[i] < nums[i - 1] {
			return nums[i]
		}
	}

	return nums[0]
}