package solution_201_250


//202. 快乐数
//编写一个算法来判断一个数 n 是不是快乐数。
//
//「快乐数」定义为：
//
//对于一个正整数，每一次将该数替换为它每个位置上的数字的平方和。
//然后重复这个过程直到这个数变为 1，也可能是 无限循环 但始终变不到 1。
//如果 可以变为 1，那么这个数就是快乐数。
//如果 n 是快乐数就返回 true ；不是，则返回 false 。


var visited202 = make(map[int]bool)


func solve202(n int) bool {
	if n == 0 {
		return false
	}
	if n == 1 {
		return true
	}

	_, ok := visited202[n]
	if ok {
		return false
	}
	visited202[n] = true

	sm := 0
	for n > 0 {
		md := n % 10
		n /= 10
		sm += md * md
	}

	return solve202(sm)
}

func isHappy202(n int) bool {
	visited202 = make(map[int]bool)
	return solve202(n)

}