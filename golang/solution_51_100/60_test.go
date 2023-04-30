package solution_51_100

import "fmt"

//60. 排列序列
//按大小顺序列出所有排列情况，并一一标记，当 n = 3 时, 所有排列如下：


var chars60 []rune
var products60 []int

func getPerm60(n int, k int, curIdx int) {
	fmt.Printf("chars: %v, k: %d, curIdx: %d\n", chars60, k, curIdx)

	if curIdx == n - 1 {  // reaches end
		return
	}

	// is last one? reverse
	// n = 4, k = 24, curIdx = 0
	if products60[n - curIdx] == k {
		left, right := curIdx, n - 1
		for left < right {
			chars60[left], chars60[right] = chars60[right], chars60[left]
			left++
			right--
		}
		return
	}

	k--
	if k == 0 {  // no change
		return
	}

	// find char to swap
	prevLadder := products60[n - curIdx - 1]
	d := k / prevLadder
	m := k % prevLadder

	// swap and push
	if d != 0 {
		tmp := chars60[curIdx]
		chars60[curIdx] = chars60[curIdx + d]
		for i := d - 1; i > 0; i-- {
			chars60[curIdx + i + 1] = chars60[curIdx + i]
		}
		chars60[curIdx + 1] = tmp
	}

	// solve next
	getPerm60(n, m + 1, curIdx + 1)
}

func getPermutation60(n int, k int) string {
	if n == 0 {
		return ""
	}
	if n == 1 {
		return "1"
	}

	// make arrays
	chars60 = make([]rune, n)
	for i := 1; i <= n; i++ {
		chars60[i - 1] = rune(int('1') + i - 1)
	}
	if k == 1 {
		return string(chars60)
	}
	products60 = make([]int, 10)
	products60[0] = 1
	for i := 1; i <= 9; i++ {
		products60[i] = i * products60[i - 1]
	}

	getPerm60(n, k, 0)
	return string(chars60)
}
