package solution_51_100

import (
	"fmt"
	"reflect"
	"testing"
)

// 89. 格雷编码
//格雷编码是一个二进制数字系统，在该系统中，两个连续的数值仅有一个位数的差异。
//
//给定一个代表编码总位数的非负整数 n，打印其格雷编码序列。即使有多个不同答案，你也只需要返回其中一种。
//
//格雷编码序列必须以 0 开头。





func grayCode89(n int) []int {
	if n == 0 {
		return []int{0}
	}

	prevCodes := grayCode89(n - 1)
	base := 1 << (n - 1)

	newCodes := make([]int, base << 1)

	for i := 0; i < base; i++ {
		newCodes[i] = prevCodes[i]
		newCodes[len(newCodes) - 1 - i] = base + prevCodes[i]
	}

	return newCodes
}


type test89 struct {
	n int
	expected []int
}


func Test_89(t *testing.T) {
	cases := append(make([]test89, 0),
		test89{
			n: 3,
			expected: []int{0,1,3,2,6,7,5,4},
		},
		test89{
			n: 2,
			expected: []int{0,1,3,2},
		},
		test89{
			n: 1,
			expected: []int{0,1},
		},
		test89{
			n: 0,
			expected: []int{0},
		},
	)

	var ans []int

	for _, c := range cases {
		fmt.Printf("testing %v+\n", c)
		ans = grayCode89(c.n)
		if !reflect.DeepEqual(ans, c.expected) {
			t.Errorf("failed at %v+, ans: %v\n", c, ans)
		} else {
			fmt.Printf("success at %v+\n", c)
		}
	}
}
