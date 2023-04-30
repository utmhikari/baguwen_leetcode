package solution_1_50

import "bytes"

func intToRoman12(num int) string {
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
	l := len(nums)
	var b bytes.Buffer
	for num > 0 {
		for i := 0; i < l; i++ {
			if num >= nums[i] {
				num -= nums[i]
				b.WriteString(chars[i])
				break
			}
		}
	}
	return b.String()
}