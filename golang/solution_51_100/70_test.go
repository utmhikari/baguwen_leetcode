package solution_51_100



//70. 爬楼梯
//假设你正在爬楼梯。需要 n阶你才能到达楼顶。
//
//每次你可以爬 1 或 2 个台阶。你有多少种不同的方法可以爬到楼顶呢？
//
//注意：给定 n 是一个正整数。

var mp70 = map[int]int{
	1: 1,
	2: 2,
	3: 3,
	4: 5,
	5: 8,
	6: 13,
	7: 21,
	8: 34,
	9: 55,
	10: 89,
	11: 144,
}

func climbStairs(n int) int {
	v, ok := mp70[n]
	if ok {
		return v
	}
	v = climbStairs(n - 1) + climbStairs(n - 2)
	mp70[n] = v
	return v
}
