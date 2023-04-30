package solution_1_50

//全排列
//给定一个 没有重复 数字的序列，返回其所有可能的全排列。



var ans46 [][]int


func solve46(nums []int, left int) {
	if left == len(nums) {
		newAns := make([]int, len(nums))
		for i := 0; i < len(nums); i++ {
			newAns[i] = nums[i]
		}
		ans46 = append(ans46, newAns)
	} else {
		solve46(nums, left + 1)
		for i := left + 1; i < len(nums); i++ {
			nums[left], nums[i] = nums[i], nums[left]
			solve46(nums, left + 1)
			nums[left], nums[i] = nums[i], nums[left]
		}
	}
}


func permute46(nums []int) [][]int {
	ans46 = make([][]int, 0)
	solve46(nums, 0)
	return ans46
}