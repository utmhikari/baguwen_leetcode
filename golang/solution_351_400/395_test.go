package solution_351_400

import "sort"

//395. 至少有 K 个重复字符的最长子串
//给你一个字符串 s 和一个整数 k ，请你找出 s 中的最长子串， 要求该子串中的每一字符出现次数都不少于 k 。返回这一子串的长度。


func solve395(s string, k int, left int, right int) int {
	if right - left + 1 < k {
		return 0
	}
	counter := make(map[uint8]int)
	idxMap := make(map[uint8][]int)
	partIndices := []int{left - 1, right + 1}

	for i := left; i <= right; i++ {
		c := s[i]
		cnt, ok := counter[c]
		if !ok {
			counter[c] = 1
			idxMap[c] = []int{i}
		} else {
			counter[c] = cnt + 1
			idxMap[c] = append(idxMap[c], i)
		}
	}

	for c, cnt := range counter {
		if cnt < k {
			partIndices = append(partIndices, idxMap[c]...)
		}
	}
	if len(partIndices) == 2 {
		return right - left + 1
	}

	sort.Ints(partIndices)
	ans := 0
	for i := 0; i < len(partIndices) - 1; i++ {
		nextAns := solve395(s, k, partIndices[i] + 1, partIndices[i+1] - 1)
		if nextAns > ans {
			ans = nextAns
		}
	}
	return ans
}

func longestSubstring(s string, k int) int {
	return solve395(s, k, 0, len(s) - 1)
}