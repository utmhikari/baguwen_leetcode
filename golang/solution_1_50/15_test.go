package solution_1_50




//15. 三数之和
//给你一个包含 n 个整数的数组nums，判断nums中是否存在三个元素 a，b，c ，
//使得a + b + c = 0 ？请你找出所有和为 0 且不重复的三元组。
//
//注意：答案中不可以包含重复的三元组。


var ans15 [][]int

func threeSum15(nums []int) [][]int {
	ans15 = make([][]int, 0)

	l := len(nums)
	if l < 3 {
		return ans15
	}

	zeroCnt := 0
	belowZeroNumCnt := make(map[int]int)
	aboveZeroNumCnt := make(map[int]int)
	distinctAboveZeros := make([]int, 0)
	distinctBelowZeros := make([]int, 0)
	for _, num := range nums {
		if num == 0 {
			zeroCnt++
		} else if num > 0 {
			cnt, ok := aboveZeroNumCnt[num]
			if !ok {
				aboveZeroNumCnt[num] = 1
				distinctAboveZeros = append(distinctAboveZeros, num)
			} else {
				aboveZeroNumCnt[num] = cnt + 1
			}
		} else {
			cnt, ok := belowZeroNumCnt[num]
			if !ok {
				belowZeroNumCnt[num] = 1
				distinctBelowZeros = append(distinctBelowZeros, num)
			} else {
				belowZeroNumCnt[num] = cnt + 1
			}
		}
	}

	// 0 + 0 + 0
	if zeroCnt >= 3 {
		ans15 = append(ans15, []int{0, 0, 0})
	}

	// 0 + a + b
	if zeroCnt > 0 {
		for num, _ := range aboveZeroNumCnt {
			_, ok := belowZeroNumCnt[-num]
			if ok {
				ans15 = append(ans15, []int{0, num, -num})
			}
		}
	}

	// a + a + b
	for num, cnt := range aboveZeroNumCnt {
		if cnt >= 2 {
			other := 0 - 2 * num
			_, ok := belowZeroNumCnt[other]
			if ok {
				ans15 = append(ans15, []int{num, num, other})
			}
		}
	}
	for num, cnt := range belowZeroNumCnt {
		if cnt >= 2 {
			other := 0 - 2 * num
			_, ok := aboveZeroNumCnt[other]
			if ok {
				ans15 = append(ans15, []int{num, num, other})
			}
		}
	}

	// a + b + c
	lda, ldb := len(distinctAboveZeros), len(distinctBelowZeros)
	if lda >= 2 {
		for i := 0; i < lda - 1; i++ {
			for j := i + 1; j < lda; j++ {
				other := 0 - distinctAboveZeros[i] - distinctAboveZeros[j]
				_, ok := belowZeroNumCnt[other]
				if ok {
					ans15 = append(ans15, []int{distinctAboveZeros[i], distinctAboveZeros[j], other})
				}
			}
		}
	}
	if ldb >= 2 {
		for i := 0; i < ldb - 1; i++ {
			for j := i + 1; j < ldb; j++ {
				other := 0 - distinctBelowZeros[i] - distinctBelowZeros[j]
				_, ok := aboveZeroNumCnt[other]
				if ok {
					ans15 = append(ans15, []int{distinctBelowZeros[i], distinctBelowZeros[j], other})
				}
			}
		}
	}

	return ans15
}