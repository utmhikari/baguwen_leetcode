package solution_301_350





//349. 两个数组的交集
//给定两个数组，编写一个函数来计算它们的交集。

func intersection349(nums1 []int, nums2 []int) []int {
	mp1, mp2 := make(map[int]bool), make(map[int]bool)
	for _, n := range nums1 {
		mp1[n] = true
	}
	for _, n := range nums2 {
		mp2[n] = true
	}
	ans := make([]int, 0)
	for n, _ := range mp1 {
		_, ok := mp2[n]
		if ok {
			ans = append(ans, n)
		}
	}
	return ans

}