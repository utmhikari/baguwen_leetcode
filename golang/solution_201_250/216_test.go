package solution_201_250


//216. 组合总和 III
//找出所有相加之和为 n 的 k 个数的组合。
//组合中只允许含有 1 - 9 的正整数，并且每种组合中不存在重复的数字。

var ans216 [][]int
var tmp216 []int

func solve216(cur int, n int, k int, sum int) {
	if len(tmp216) + (n - cur + 1) < k || len(tmp216) > k {
		return
	}
	if len(tmp216) == k {
		sm := 0
		for i := 0; i < k; i++ {
			sm += tmp216[i]
		}
		if sm == sum {
			ans := make([]int, k)
			copy(ans, tmp216)
			ans216 = append(ans216, ans)
		}
	} else {
		tmp216 = append(tmp216, cur)
		solve216(cur + 1, n, k, sum)
		tmp216 = tmp216[:len(tmp216) - 1]
		solve216(cur + 1, n, k, sum)
	}

}


func combinationSum3216(k int, n int) [][]int {
	ans216 = make([][]int, 0)

	if 9 * k < n || 1 * k > n {
		return ans216
	}

	tmp216 = make([]int, 0)
	solve216(1, 9, k, n)
	return ans216
}