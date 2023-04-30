package solution_151_200



func findMin153(nums []int) int {
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