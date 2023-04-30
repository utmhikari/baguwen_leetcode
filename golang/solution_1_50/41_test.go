package solution_1_50

//41. 缺失的第一个正数
//给你一个未排序的整数数组 nums ，请你找出其中没有出现的最小的正整数。



func firstMissingPositive(nums []int) int {
	l := len(nums)
	i := 0
	for i < l {
		n := nums[i]
		if n == i + 1 {
			i++
			continue
		}
		if 0 < n && n <= l && nums[n - 1] != nums[i] {
			nums[n - 1], nums[i] = nums[i], nums[n - 1]
			continue
		}
		i++
	}

	//fmt.Printf("%v\n", nums)

	ans := 1
	for i, n := range nums {
		if n == i + 1 {
			ans++
		} else {
			break
		}
	}
	return ans
}