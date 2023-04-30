package solution_1_50

import (
	"math"
	"sort"
)

//
//16. 最接近的三数之和
//给定一个包括n 个整数的数组nums和 一个目标值target。
//找出nums中的三个整数，使得它们的和与target最接近。
//返回这三个数的和。假定每组输入只存在唯一答案。


func abs16(a int) int {
	if a >= 0 {
		return a
	}
	return 0 - a
}

func threeSumClosest16(nums []int, target int) int {
	l := len(nums)
	sort.Ints(nums)

	minAbs := math.MaxInt64
	minVal := 0

	p := 0
	for p < l - 2 {
		p1, p2 := p + 1, l - 1
		for p1 < p2 {
			sm := nums[p] + nums[p1] + nums[p2]
			if sm == target {
				return target
			}
			abs := abs16(target - sm)
			if abs < minAbs {
				minAbs = abs
				minVal = sm
			}
			if sm < target {
				p1++
			} else {
				p2--
			}
		}
		p++
	}
	return minVal
}

