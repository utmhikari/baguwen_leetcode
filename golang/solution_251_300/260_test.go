package solution_251_300



//260. 只出现一次的数字 III
//
//给定一个整数数组 nums，其中恰好有两个元素只出现一次，其余所有元素均出现两次。
//找出只出现一次的那两个元素。你可以按 任意顺序 返回答案。


func singleNumber260(nums []int) []int {
	// a ^ b = n
	l := len(nums)
	if l == 2 {
		return nums
	}

	n := nums[0]
	for i := 1; i < l; i++ {
		n ^= nums[i]
	}

	div := 1
	for div & n == 0 {
		div <<= 1
	}
	a, b := 0, 0
	for _, x := range nums {
		if div & x != 0 {
			a ^= x
		} else {
			b ^= x
		}
	}
	return []int{a,b}
}
