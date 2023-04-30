package solution_301_350




//321. 拼接最大数
//给定长度分别为 m 和 n 的两个数组，
//其元素由 0-9 构成，表示两个自然数各位上的数字。
//现在从这两个数组中选出 k (k <= m + n) 个数字拼接成一个新的数，
//要求从同一个数组中取出的数字保持其在原数组中的相对顺序。
//
//求满足该条件的最大数。结果返回一个表示该最大数的长度为 k 的数组。
//
//说明: 请尽可能地优化你算法的时间和空间复杂度。


// 单调栈，选择k个数组成最大的
func maxSubsequence321(a []int, k int) []int {
	seq := make([]int, 0)
	for i, v := range a {
		for len(seq) > 0 && len(seq) - 1 + len(a) - i >= k && v > seq[len(seq) - 1] {
			seq = seq[:len(seq) - 1]
		}
		if len(seq) < k {
			seq = append(seq, v)
		}
	}
	return seq
}



func arrMin321(a, b []int) bool {
	for i := 0; i < len(a) && i < len(b); i++ {
		if a[i] != b[i] {
			return a[i] < b[i]
		}
	}
	return len(a) < len(b)
}


func merge321(a, b []int) []int {
	merged := make([]int, len(a) + len(b))
	for i := range merged {
		if arrMin321(a, b) {
			merged[i], b = b[0], b[1:]
		} else {
			merged[i], a = a[0], a[1:]
		}
	}
	return merged
}


func maxNumber321(nums1 []int, nums2 []int, k int) []int {
	start := 0
	if k > len(nums2) {
		start = k - len(nums2)
	}
	ans := make([]int, 0)
	for i := start; i <= k && i <= len(nums1); i++ {
		s1 := maxSubsequence321(nums1, i)
		s2 := maxSubsequence321(nums2, k - i)
		merged := merge321(s1, s2)
		if arrMin321(ans, merged) {
			ans = merged
		}
	}
	return ans
}

