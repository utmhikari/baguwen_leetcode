package solution_201_250

import (
	"fmt"
	"strconv"
)

//228. 汇总区间
//给定一个无重复元素的有序整数数组 nums 。
//
//返回 恰好覆盖数组中所有数字 的 最小有序 区间范围列表。
//也就是说，nums 的每个元素都恰好被某个区间范围所覆盖，并且不存在属于某个范围但不属于 nums 的数字 x 。
//
//列表中的每个区间范围 [a,b] 应该按如下格式输出：


func summaryRanges(nums []int) []string {
	ans := make([]string, 0)
	l := len(nums)
	if l == 0 {
		return ans
	}
	if l == 1 {
		ans = append(ans, strconv.Itoa(nums[0]))
		return ans
	}
	leftIdx := 0
	for leftIdx < l {
		i := leftIdx + 1
		for i < l && nums[i] - nums[i - 1] == 1 {
			i++
		}
		if i == leftIdx + 1 {
			ans = append(ans, strconv.Itoa(nums[leftIdx]))
		} else {
			ans = append(ans, fmt.Sprintf("%d->%d", nums[leftIdx], nums[i - 1]))
		}
		leftIdx = i
	}
	return ans
}