package solution_1_50

import "sort"

//39. 组合总和
//给定一个无重复元素的数组candidates
//和一个目标数target，
//找出candidates中所有可以使数字和为target的组合。
//
//candidates中的数字可以无限制重复被选取。


var ans39 [][]int
var tmp39 []int


func solve39(candidates []int, target int, cur int, sm int) {
	if cur < 0 {
		return
	}
	for i := cur; i >= 0; i-- {
		sm += candidates[i]
		if sm > target {
			sm -= candidates[i]
			continue
		}
		tmp39 = append(tmp39, candidates[i])
		if sm == target {
			newAns := make([]int, len(tmp39))
			copy(newAns, tmp39)
			ans39 = append(ans39, newAns)
		} else {
			solve39(candidates, target, i, sm)
		}
		tmp39 = tmp39[:len(tmp39) - 1]
		sm -= candidates[i]
	}
}


func combinationSum39(candidates []int, target int) [][]int {
	ans39 = make([][]int, 0)
	tmp39 = make([]int, 0)
	sort.Ints(candidates)

	l := len(candidates)
	solve39(candidates, target, l - 1, 0)

	return ans39
}