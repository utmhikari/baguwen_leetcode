package solution_1_50

import "strings"

func romanToInt13(s string) int {
	chars := []string{
		"M",
		"CM",
		"D",
		"CD",
		"C",
		"XC",
		"L",
		"XL",
		"X",
		"IX",
		"V",
		"IV",
		"I",
	}
	nums := []int{
		1000,
		900,
		500,
		400,
		100,
		90,
		50,
		40,
		10,
		9,
		5,
		4,
		1,
	}
	l := len(s)
	ln := len(nums)
	sm := 0
	i := 0
	for i < l {
		for j := 0; j < ln; j++ {
			if strings.HasPrefix(s[i:], chars[j]) {
				i += len(chars[j])
				sm += nums[j]
				break
			}
		}
	}
	return sm
}