package solution_201_250



//204. 计数质数
//统计所有小于非负整数 n 的质数的数量。


func countPrimes204(n int) int {
	if n <= 2 {
		return 0
	}

	// non-primes
	nonPrimes := make([]bool, n)
	// 0+1 ~ n-1+1
	nonPrimes[0] = true
	nonPrimes[n - 1] = true

	cnt := 0
	for i := 1; i < n; i++ {
		if nonPrimes[i] {
			continue
		}
		cnt++
		for j := 2; (i + 1) * j <= n; j++ {
			nonPrimes[(i + 1) * j - 1] = true
		}
	}
	return cnt
}