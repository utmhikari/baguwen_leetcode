package solution_301_350



//306. 累加数
//
//累加数是一个字符串，组成它的数字可以形成累加序列。
//
//一个有效的累加序列必须至少包含 3 个数。
//除了最开始的两个数以外，字符串中的其他数都等于它之前两个数相加的和。
//
//给定一个只包含数字 '0'-'9' 的字符串，
//编写一个算法来判断给定输入是否是累加数。
//
//说明: 累加序列里的数不会以 0 开头，所以不会出现 1, 2, 03 或者 1, 02, 3 的情况。


// > 0: true, == 0: false, < 0: no le value
func dfs306(num string, left int, first int, second int) int {
	l := len(num)
	if left == l {
		return 1
	}

	expected := first + second
	sm := 0
	for i := left; i < l; i++ {
		if num[left] == '0' && i > left {
			return 0
		}
		sm = sm * 10 + int(num[i] - '0')
		if sm == expected {
			ret := dfs306(num, i + 1, second, expected)
			if ret > 0 {
				return 1
			} else {
				return 0
			}
		} else if sm > expected {
			return 0
		}
	}
	return -1
}


func isAdditiveNumber(num string) bool {
	l := len(num)
	if l < 3 {
		return false
	}
	first := 0
	for i := 0; i < l - 2; i++ {
		if num[0] == '0' && i > 0 {
			break
		}
		first = first * 10 + int(num[i] - '0')
		second := 0
		for j := i + 1; j < l - 1; j++ {
			if num[i + 1] == '0' && j > i + 1 {
				break
			}
			second = second * 10 + int(num[j] - '0')
			ret := dfs306(num, j + 1, first, second)
			if ret > 0 {
				return true
			}
			if ret < 0 {
				break
			}
		}
	}
	return false
}
