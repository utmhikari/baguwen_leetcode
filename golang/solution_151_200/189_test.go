package solution_151_200



func reverse189(nums[]int, left int, right int) {
	for left < right {
		nums[left], nums[right] = nums[right], nums[left]
		left++
		right--
	}
}


func rotate(nums []int, k int)  {
	l := len(nums)
	if l <= 1 {
		return
	}
	m := k % l
	if m == 0 {
		return
	}

	reverse189(nums, 0, l - 1)
	reverse189(nums, 0, k - 1)
	reverse189(nums, k, l - 1)
}