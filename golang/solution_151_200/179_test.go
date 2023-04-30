package solution_151_200

import (
	"bytes"
	"fmt"
	"sort"
)

//179. 最大数
//给定一组非负整数 nums，重新排列每个数的顺序（每个数不可拆分）使之组成一个最大的整数。
//
//注意：输出结果可能非常大，所以你需要返回一个字符串而不是整数。


func cmp179(s1 string, s2 string, i1 int, i2 int) bool {
	l1, l2 := len(s1) - i1, len(s2) - i2
	if l1 == 0 {
		return true
	}
	if l2 == 0 {
		return false
	}
	for i := 0; i < l1 && i < l2; i++ {
		if s1[i1 + i] > s2[i2 + i] {
			return true
		}
		if s1[i1 + i] < s2[i2 + i] {
			return false
		}
	}
	if l1 < l2 {
		return cmp179(s1, s2, 0, i2 + l1)
	} else if l1 > l2 {
		return cmp179(s1, s2, i1 + l2, 0)
	} else {
		return true
	}
}


func largestNumber(nums []int) string {
	if len(nums) == 0 {
		return ""
	}
	numStrs := make([]string, len(nums))
	for i, n := range nums {
		numStrs[i] = fmt.Sprintf("%d", n)
	}
	sort.Slice(numStrs, func(i int, j int) bool {
		return cmp179(numStrs[i], numStrs[j], 0, 0)
	})
	if numStrs[0] == "0" {
		return "0"
	}
	var b bytes.Buffer
	for _, s := range numStrs {
		b.WriteString(s)
	}
	return b.String()
}