package solution_51_100

//91. 解码方法

//一条包含字母A-Z 的消息通过以下映射进行了 编码 ：
//
//'A' -> 1
//'B' -> 2
//...
//'Z' -> 26
//要 解码 已编码的消息，所有数字必须基于上述映射的方法，反向映射回字母（可能有多种方法）。例如，"11106" 可以映射为：
//
//"AAJF" ，将消息分组为 (1 1 10 6)
//"KJF" ，将消息分组为 (11 10 6)
//注意，消息不能分组为 (1 11 06) ，因为 "06" 不能映射为 "F" ，这是由于 "6" 和 "06" 在映射中并不等价。
//
//给你一个只含数字的 非空 字符串 s ，请计算并返回 解码 方法的 总数 。
//
//题目数据保证答案肯定是一个 32 位 的整数。


var map91 []int


func solve91(s string, cur int) int {
	if cur < 0 {
		return 1
	}

	if map91[cur] >= 0 {
		return map91[cur]
	}

	sm := 0

	if s[cur] != '0' {
		sm += solve91(s, cur - 1)
	}

	if cur - 1 >= 0 {
		if (s[cur - 1] == '1' && s[cur] <= '9') || (s[cur - 1] == '2' && s[cur] <= '6') {
			sm += solve91(s, cur - 2)
		}
	}

	map91[cur] = sm
	return sm
}

func numDecodings91(s string) int {
	l := len(s)
	if l == 0 || s == "0" {
		return 0
	}
	if l == 1 {
		return 1
	}
	map91 = make([]int, l)
	for i := 0; i < l; i++ {
		map91[i] = -1
	}
	solve91(s, l - 1)
	return map91[l - 1]
}