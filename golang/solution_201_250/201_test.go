package solution_201_250

//201. 数字范围按位与
//给你两个整数 left 和 right ，表示区间 [left, right] ，
//返回此区间内所有数字 按位与 的结果（包含 left 、right 端点）。

func rangeBitwiseAnd201(left int, right int) int {
	if left == right {
		return left
	}
	if left <= 1 {
		return 0
	}
	sm := 0
	for i := 30; i >= 1; i-- {
		base := 1 << i
		if (base & left) == (base & right) {
			sm += base & left
		} else if (base & left) != (base & right) {
			break
		}
	}
	return sm
}
