package solution_301_350

//350. 两个数组的交集 II
//给定两个数组，编写一个函数来计算它们的交集。
//带计数

func intersect350(nums1 []int, nums2 []int) []int {
	mp1 := make(map[int]int)
	for _, n := range nums1 {
		cnt, ok := mp1[n]
		if ok {
			mp1[n] = cnt + 1
		} else {
			mp1[n] = 1
		}
	}
	ans := make([]int, 0)
	for _, n := range nums2 {
		cnt, ok := mp1[n]
		if ok && cnt > 0 {
			ans = append(ans, n)
			mp1[n] = cnt - 1
		}
	}
	return ans
}