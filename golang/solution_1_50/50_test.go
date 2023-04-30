package solution_1_50

import (
	"fmt"
	"math"
	"testing"
)

// 50. Pow(x, n)
// 实现 pow(x, n) ，即计算 x 的 n 次幂函数（即，xn）。



func myPow(x float64, n int) float64 {
	if x == 0 {
		return 0
	}
	if n == 0 {
		return 1
	}
	isMinInt := n == math.MinInt32
	bzFlag := n < 0
	if isMinInt {
		n = -(n / 2)
	} else if bzFlag {
		n = -n
	}

	if n == 1 {
		if bzFlag {
			return 1 / x
		} else {
			return x
		}
	}

	bitFlag := 1 << 31
	maxBit := 32
	for {
		bitFlag >>= 1
		maxBit--
		if n & bitFlag == bitFlag {
			break
		}
	}

	pows := make([]float64, maxBit)
	ans := 1.0
	for i := 0; i < len(pows); i++ {
		if i == 0 {
			pows[i] = x
		} else {
			pows[i] = pows[i - 1] * pows[i - 1]
		}
		flag := 1 << i
		if flag & n == flag {
			ans *= pows[i]
		}
	}

	if isMinInt {
		ans *= pows[len(pows)-1]
		ans *= pows[len(pows)-1]
	}

	if bzFlag {
		return 1 / ans
	} else {
		return ans
	}
}


type test50 struct {
	x float64
	n int
	expected float64
}


func Test_50 (t *testing.T) {
	cases := append(make([]test50, 0),
		test50{
			x: 2.0,
			n: 10,
			expected: 1024.0,
		},
		test50{
			x: 2.1,
			n: 3,
			expected: 9.26100,
		},
		test50{
			x: 2.0,
			n: -2,
			expected: 0.250,
		},
		test50{
			x: 2.0,
			n: -2147483648,
			expected: 0.0,
		},
	)

	var ans float64

	for _, c := range cases {
		ans = myPow(c.x, c.n)
		if math.Dim(ans, c.expected) > 0.000001 {
			t.Errorf("failed at %v+, ans: %f\n", c, ans)
		} else {
			fmt.Printf("success at %v+\n", c)
		}
	}
}



