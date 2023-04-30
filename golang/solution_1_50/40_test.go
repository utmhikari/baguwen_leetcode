package solution_1_50

import "sort"

//40. 组合总和 II
//给定一个数组candidates和一个目标数target，
//找出candidates中所有可以使数字和为target的组合。
//
//candidates中的每个数字在每个组合中只能使用一次。

var ans40 [][]int
var tmp40 []int


func solve40(candidates []int, target int, cur int, sm int) {
	if cur < 0 {
		return
	}
	for i := cur; i >= 0; i-- {
		sm += candidates[i]
		if sm > target {
			sm -= candidates[i]
			continue
		}
		tmp40 = append(tmp40, candidates[i])
		if sm == target {
			newAns := make([]int, len(tmp40))
			copy(newAns, tmp40)
			ans40 = append(ans40, newAns)
		} else {
			solve40(candidates, target, i - 1, sm)
		}
		tmp40 = tmp40[:len(tmp40) - 1]
		sm -= candidates[i]

		for i >= 1 && candidates[i] == candidates[i - 1] {
			i--
		}
	}
}

func combinationSum240(candidates []int, target int) [][]int {
	ans40 = make([][]int, 0)
	tmp40 = make([]int, 0)
	sort.Ints(candidates)

	l := len(candidates)
	solve40(candidates, target, l - 1, 0)

	return ans40
}