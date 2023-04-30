package solution_251_300



//268. 丢失的数字
//给定一个包含 [0, n] 中 n 个数的数组 nums ，
//找出 [0, n] 这个范围内没有出现在数组中的那个数。


func missingNumber268(nums []int) int {
	l := len(nums)
	i := 0
	for i < l {
		n := nums[i]
		if n < 0 || n >= l || nums[nums[i]] == nums[i] {
			i++
		} else {
			nums[i], nums[nums[i]] = nums[nums[i]], nums[i]
		}
	}
	i = 0
	for i < l {
		if i != nums[i] {
			return i
		}
		i++
	}
	return l
}