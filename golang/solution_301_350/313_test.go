package solution_301_350

import (
	"fmt"
	"math"
)

//313. 超级丑数
//编写一段程序来查找第 n 个超级丑数。
//
//超级丑数是指其所有质因数都是长度为 k 的质数列表 primes 中的正整数。



func nthSuperUglyNumber313(n int, primes []int) int {
	if n == 0 {
		return 1
	}

	l := len(primes)
	nums := make([]int, n)
	nums[0] = 1
	indices := make([]int, l)

	for i := 1; i < n; i++ {
		minIndices, minNum := make([]int, 0), math.MaxInt32
		for j := 0; j < l; j++ {
			curNum := nums[indices[j]] * primes[j]
			if curNum < minNum {
				minNum = curNum
				minIndices = []int{j}
			} else if curNum == minNum {
				minIndices = append(minIndices, j)
			}
		}

		nums[i] = minNum
		for _, minIdx := range minIndices {
			indices[minIdx]++
		}
	}

	fmt.Printf("%v\n", nums)
	return nums[n - 1]
}