package solution_201_250



//219. 存在重复元素 II
//给定一个整数数组和一个整数k，判断数组中是否存在两个不同的索引i和j，
//使得nums [i] = nums [j]，并且 i 和 j的差的 绝对值 至多为 k。



func containsNearbyDuplicate219(nums []int, k int) bool {
	mp := make(map[int]int)
	for i := 0; i < len(nums); i++ {
		prev, ok := mp[nums[i]]
		if ok {
			if i - prev <= k {
				return true
			}
		}
		mp[nums[i]] = i
	}
	return false
}