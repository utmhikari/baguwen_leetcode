package solution_1_50


//1. 两数之和
//给定一个整数数组 nums和一个整数目标值 target，请你在该数组中找出 和为目标值 的那两个整数，并返回它们的数组下标。
//
//你可以假设每种输入只会对应一个答案。但是，数组中同一个元素在答案里不能重复出现。
//
//你可以按任意顺序返回答案。


func twoSum1(nums []int, target int) []int {
	l := len(nums)
	if l <= 1 {
		return []int{}
	}

	set := make(map[int]int)
	for i := 0; i < l; i++ {
		v := target - nums[i]
		idx, ok := set[v]
		if ok {
			return []int{idx, i}
		}
		set[nums[i]] = i
	}
	return []int{}
}
