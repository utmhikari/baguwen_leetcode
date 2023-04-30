package solution_201_250

//229. 求众数 II
//给定一个大小为n的整数数组，找出其中所有出现超过⌊ n/3 ⌋次的元素。
//
//进阶：尝试设计时间复杂度为 O(n)、空间复杂度为 O(1)的算法解决此问题。


func majorityElement229(nums []int) []int {
	l := len(nums)
	ans := make([]int, 0)
	if l == 0 {
		return ans
	}

	cand1, cand2 := nums[0], nums[0]
	cnt1, cnt2 := 0, 0
	for _, num := range nums {
		if cand1 == num {
			cnt1++
			continue
		}
		if cand2 == num {
			cnt2++
			continue
		}
		if cnt1 == 0 {
			cand1 = num
			cnt1++
			continue
		}
		if cnt2 == 0 {
			cand2 = num
			cnt2++
			continue
		}

		cnt1--
		cnt2--
	}

	cnt1, cnt2 = 0, 0
	for _, num := range nums {
		if cand1 == num {
			cnt1++
		} else if cand2 == num {
			cnt2++
		}
	}

	if cnt1 > l / 3 {
		ans = append(ans, cand1)
	}
	if cnt2 > l / 3 {
		ans = append(ans, cand2)
	}
	return ans
}