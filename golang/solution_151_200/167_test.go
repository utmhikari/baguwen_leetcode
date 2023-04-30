package solution_151_200


//167. 两数之和 II - 输入有序数组
// https://leetcode-cn.com/problems/two-sum-ii-input-array-is-sorted/
//给定一个已按照 升序排列 的整数数组numbers ，请你从数组中找出两个数满足相加之和等于目标数target 。
//
//函数应该以长度为 2 的整数数组的形式返回这两个数的下标值。numbers 的下标 从 1 开始计数 ，所以答案数组应当满足 1 <= answer[0] < answer[1] <= numbers.length 。
//
//你可以假设每个输入只对应唯一的答案，而且你不可以重复使用相同的元素。



func twoSum167(numbers []int, target int) []int {
	l := len(numbers)
	p1, p2 := 0, l - 1
	for p1 < p2 {
		sm := numbers[p1] + numbers[p2]
		if sm == target {
			return []int{p1 + 1, p2 + 1}
		}
		if sm < target {
			p1++
		} else {
			p2--
		}
	}

	return []int{0, 0}
}