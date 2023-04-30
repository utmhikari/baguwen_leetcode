package solution_51_100

import (
	"fmt"
	"reflect"
	"testing"
)

//77. 组合
//给定两个整数 n 和 k，返回 1 ... n 中所有可能的 k 个数的组合。

var ans77 [][]int

func dfs77(depth int, arr []int, curDigitIdx int, digits []int) {
	// fmt.Printf("depth: %d, arr: %v, curDigit: %d\n", depth, arr, digits[curDigitIdx])
	leftDigits := len(arr) - depth
	if leftDigits == 0 {
		ansArr := make([]int, len(arr))
		copy(ansArr, arr)
		//fmt.Printf("ansArr: %v\n", ansArr)
		ans77 = append(ans77, ansArr)
		//fmt.Printf("ans77: %v\n", ans77)
		return
	}

	for i := curDigitIdx; i < len(digits) - leftDigits + 1; i++ {
		arr[depth] = digits[i]
		dfs77(depth + 1, arr, i + 1, digits)
	}

}



func combine77(n int, k int) [][]int {
	if k > n || n <= 0 || k <= 0 {
		return make([][]int, 0)
	}
	ans77 = make([][]int, 0)
	arr := make([]int, k)
	digits := make([]int, n)
	for i := 0; i < n; i++ {
		digits[i] = i + 1
	}

	dfs77(0, arr, 0, digits)
	return ans77
}


type test77 struct {
	n int
	k int
	expected [][]int
}


func Test_77(t *testing.T) {
	cases := append(make([]test77, 0),
		test77{
			n: 4,
			k: 2,
			expected: [][]int{
				{2,4},
				{3,4},
				{2,3},
				{1,2},
				{1,3},
				{1,4},
			},
		},
	)

	for i, c := range cases {
		a := combine77(c.n, c.k)
		if !reflect.DeepEqual(a, c.expected) {
			t.Errorf("failed at %d -> ans77: %v, case %v+\n", i, a, c)
		} else {
			fmt.Printf("success at %d -> ans77: %v\n", i, a)
		}
	}
}
