package solution_51_100

import (
	"fmt"
	"strconv"
	"testing"
)

//67. 二进制求和
//给你两个二进制字符串，返回它们的和（用二进制表示）。
//输入为 非空 字符串且只包含数字 1 和 0。


func addBinary(a string, b string) string {
	switch {
	case a == "0":
		return b
	case b == "0":
		return a
	default:
		break
	}

	i := 0
	addOne := false
	var sm int64
	var sa int64
	var sb int64
	ans := ""
	for {
		if i >= len(a) && i >= len(b) {
			break
		}
		if i < len(a) {
			sa, _ = strconv.ParseInt(string(a[len(a) - 1 - i]), 10, 64)
		} else {
			sa = 0
		}
		if i < len(b) {
			sb, _ = strconv.ParseInt(string(b[len(b) - 1 - i]), 10, 64)
		} else {
			sb = 0
		}

		sm = sa + sb
		if addOne {
			sm++
		}

		switch sm {
		case 3:
			addOne = true
			ans = "1" + ans
			break
		case 2:
			addOne = true
			ans = "0" + ans
			break
		case 1:
			addOne = false
			ans = "1" + ans
			break
		case 0:
			addOne = false
			ans = "0" + ans
			break
			default:
				break
		}

		i++
	}

	if addOne {
		ans = "1" + ans
	}

	return ans
}


type test67 struct {
	a string
	b string
	expected string
}


func Test_67(t *testing.T) {
	cases := append(make([]test67, 0),
		test67{
			a:        "11",
			b:        "1",
			expected: "100",
		},
		test67{
			a:        "1010",
			b:        "1011",
			expected: "10101",
		})

	var ans string

	for _, c := range cases {
		ans = addBinary(c.a, c.b)
		if c.expected != ans {
			t.Errorf("failed at %v+, ans: %s\n", c, ans)
		} else {
			fmt.Printf("success at %v+\n", c)
		}
	}
}