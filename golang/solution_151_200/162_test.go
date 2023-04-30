package solution_151_200


//162. 寻找峰值
//峰值元素是指其值大于左右相邻值的元素。


func solve162(nums []int, left int, right int) int {
	if left > right {
		return -1
	}
	l := len(nums)
	mid := left + (right - left) / 2
	if (mid == 0 || nums[mid - 1] < nums[mid]) &&
		(mid == l - 1 || nums[mid] > nums[mid + 1]) {
		return mid
	}
	switch mid {
	case 0:
		return solve162(nums, mid + 1, right)
	case l - 1:
		return solve162(nums, left, mid - 1)
	default:
		isRightSmaller := nums[mid + 1] < nums[mid]
		var firstRet int
		if !isRightSmaller {
			firstRet = solve162(nums, left, mid - 1)
			if firstRet != -1 {
				return firstRet
			}
			return solve162(nums, mid + 1, right)
		} else {
			firstRet = solve162(nums, mid + 1, right)
			if firstRet != -1 {
				return firstRet
			}
			return solve162(nums, left, mid - 1)
		}
	}
}

func findPeakElement162(nums []int) int {
	l := len(nums)
	if l == 0 {
		return -1
	}
	return solve162(nums, 0, l - 1)
}
