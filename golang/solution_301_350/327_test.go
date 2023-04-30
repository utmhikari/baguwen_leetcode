package solution_301_350

//327. 区间和的个数
//给你一个整数数组 nums 以及两个整数 lower 和 upper 。
//求数组中，值位于范围 [lower, upper] （包含 lower 和 upper）之内的 区间和的个数 。
//
//区间和 S(i, j) 表示在 nums 中，位置从 i 到 j 的元素之和，
//包含 i 和 j (i ≤ j)。

var sums327 []int

func solve327(lower int, upper int, left int, right int) int {
	if left == right {
		return 0
	}

	mid := (left + right) / 2
	n1 := solve327(lower, upper, left, mid)
	n2 := solve327(lower, upper, mid + 1, right)
	ret := n1 + n2

	// get num of idx pairs
	i, l, r := left, mid + 1, mid + 1
	for i <= mid {
		for l <= right && sums327[l] - sums327[i] < lower {
			l++
		}
		for r <= right && sums327[r] - sums327[i] <= upper {
			r++
		}
		ret += r - l
		i++
	}

	// mergesort
	sorted := make([]int, right - left + 1)
	p1, p2 := left, mid + 1
	p := 0
	for p1 <= mid || p2 <= right {
		if p1 > mid {
			sorted[p] = sums327[p2]
			p++
			p2++
		} else if p2 > right {
			sorted[p] = sums327[p1]
			p++
			p1++
		} else {
			if sums327[p1] < sums327[p2] {
				sorted[p] = sums327[p1]
				p++
				p1++
			} else {
				sorted[p] = sums327[p2]
				p++
				p2++
			}
		}
	}
	for i := 0; i < len(sorted); i++ {
		sums327[left + i] = sorted[i]
	}
	return ret
}

func countRangeSum327(nums []int, lower int, upper int) int {
	l := len(nums)
	sums327 = make([]int, l + 1)
	for i := 0; i < l; i++ {
		sums327[i + 1] = sums327[i] + nums[i]
	}

	return solve327(lower, upper, 0, l)
}
