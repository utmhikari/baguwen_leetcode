package solution_101_150

import (
	"math"
	"testing"
)

func maxInt123(a int, b int) int {
	if a > b {
		return a
	}
	return b
}

func maxProfit123(prices []int) int {
	l := len(prices)
	if l <= 1 {
		return 0
	}

	b1, b2 := math.MinInt32, math.MinInt32
	s1, s2 := 0, 0

	for _, p := range prices {
		b1 = maxInt123(b1, -p)
		s1 = maxInt123(s1, b1 + p)
		b2 = maxInt123(b2, s1 - p)
		s2 = maxInt123(s2, b2 + p)
	}

	return s2
}



type test123 struct {
	prices []int
	expected int
}

func Test_123(t *testing.T) {
	inputs := []test123{
		{[]int{3,3,5,0,0,3,1,4}, 6},
		{[]int{1,2,3,4,5}, 4},
		{[]int{7,6,4,3,1}, 0},
		{[]int{1}, 0},
	}

	for i, input := range inputs {
		ans := maxProfit123(input.prices)
		if ans != input.expected {
			t.Errorf("failed at %d (%d) -> %+v\n", i, ans, input)
		} else {
			t.Logf("success at %d -> %+v\n", i, input)
		}
	}

}