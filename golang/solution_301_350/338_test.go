package solution_301_350


//338. 比特位计数
//给定一个非负整数 num。
//对于 0 ≤ i ≤ num 范围中的每个数字 i ，
//计算其二进制数中的 1 的数目并将它们作为数组返回。


func countBits338(n int) []int {
	base := 1
	ans := make([]int, n + 1)
	for i := 1; i <= n; i++ {
		if i == base {
			ans[i] = 1
			base <<= 1
		} else {
			ans[i] = 1 + ans[i - (base >> 1)]
		}

	}
	return ans
}