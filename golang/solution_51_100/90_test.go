package solution_51_100

//90. 子集 II
//给你一个整数数组 nums ，其中可能包含重复元素，请你返回该数组所有可能的子集（幂集）。
//
//解集 不能 包含重复的子集。返回的解集中，子集可以按 任意顺序 排列。


var ans90 [][]int
var tmp90 []int
var map90 map[int]int
var keys90 []int


func solve90(cur int) {
	if cur >= len(keys90) {
		subAns := make([]int, len(tmp90))
		if len(tmp90) > 0 {
			copy(subAns, tmp90)
		}
		ans90 = append(ans90, subAns)
		return
	}

	// cur is curIdx for key
	n := keys90[cur]
	c := map90[n]
	for i := 0; i <= c; i++ {
		for j := 0; j < i; j++ {
			tmp90 = append(tmp90, n)
		}
		solve90(cur + 1)
		tmp90 = tmp90[:len(tmp90) - i]
	}
}

func subsetsWithDup90(nums []int) [][]int {
	ans90 = make([][]int, 0)
	l := len(nums)
	if l == 0 {
		return ans90
	}
	tmp90 = make([]int, 0)
	keys90 = make([]int, 0)
	map90 = make(map[int]int)
	for _, n := range nums {
		cnt, ok := map90[n]
		if !ok {
			map90[n] = 1
			keys90 = append(keys90, n)
		} else {
			map90[n] = cnt + 1
		}
	}
	solve90(0)
	return ans90
}