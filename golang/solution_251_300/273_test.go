package solution_251_300

//273. 整数转换英文表示


var numStrMap273 = map[int]string {
	0: "",
	1: "One",
	2: "Two",
	3: "Three",
	4: "Four",
	5: "Five",
	6: "Six",
	7: "Seven",
	8: "Eight",
	9: "Nine",
	10: "Ten",
	11: "Eleven",
	12: "Twelve",
	13: "Thirteen",
	14: "Fourteen",
	15: "Fifteen",
	16: "Sixteen",
	17: "Seventeen",
	18: "Eighteen",
	19: "Nineteen",
	20: "Twenty",
	30: "Thirty",
	40: "Forty",
	50: "Fifty",
	60: "Sixty",
	70: "Seventy",
	80: "Eighty",
	90: "Ninety",
}
var tens273 = []int{90,80,70,60,50,40,30,20}


func solve273(num int) string {
	s, ok := numStrMap273[num]
	if ok {
		return s
	}

	var left, right string

	// billion
	if num >= 1000000000 {
		left, right = solve273(num / 1000000000), solve273(num % 1000000000)
		if len(right) == 0 {
			return left + " Billion"
		} else {
			return left + " Billion " + right
		}
	}

	if num >= 1000000 {
		left, right = solve273(num / 1000000), solve273(num % 1000000)
		if len(right) == 0 {
			return left + " Million"
		} else {
			return left + " Million " + right
		}
	}

	if num >= 1000 {
		left, right = solve273(num / 1000), solve273(num % 1000)
		if len(right) == 0 {
			return left + " Thousand"
		} else {
			return left + " Thousand " + right
		}
	}

	if num >= 100 {
		left, right = solve273(num / 100), solve273(num % 100)
		if len(right) == 0 {
			return left + " Hundred"
		} else {
			return left + " Hundred " + right
		}
	}

	for _, n := range tens273 {
		if num > n {
			s, _ = numStrMap273[n]
			return s + " " + solve273(num - n)
		}
	}

	return ""
}

func numberToWords(num int) string {
	if num == 0 {
		return "Zero"
	}

	return solve273(num)
}


