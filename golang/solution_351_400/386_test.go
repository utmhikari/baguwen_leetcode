package solution_351_400

//386. 字典序排数


var ans386 []int
var i386 int
var n386 int


func solve386() {
	if i386 <= n386 {
		ans386 = append(ans386, i386)
		for i := 0; i <= 9; i++ {
			cur := i386
			i386 = i386* 10 + i
			solve386()
			i386 = cur
		}
	}
}


func lexicalOrder(n int) []int {
	ans386 = make([]int, 0)
	if n < 1 {
		return ans386
	}
	n386 = n
	i386 = 0
	for i := 1; i <= 9; i++ {
		i386 = i
		solve386()
	}
	return ans386
}