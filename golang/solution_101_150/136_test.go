package solution_101_150


// 136. 只出现一次的数字
///给定一个非空整数数组，除了某个元素只出现一次以外，其余每个元素均出现两次。找出那个只出现了一次的元素。


func singleNumber136(nums []int) int {
	base := 0
	for _, num := range nums {
		base ^= num
	}
	return base
}
