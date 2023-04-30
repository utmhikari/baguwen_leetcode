package solution_1_50

import (
	"bytes"
)

//43. 字符串相乘
//给定两个以字符串形式表示的非负整数 num1 和 num2，返回 num1 和 num2 的乘积，
//它们的乘积也表示为字符串形式。


func multiply(num1 string, num2 string) string {
	l1, l2 := len(num1), len(num2)
	if l1 == 0 || l2 == 0 {
		return ""
	}

	values43 := make(map[int]int)

	for i := l1 - 1; i >= 0; i-- {
		for j := l2 - 1; j >= 0; j-- {
			pos := l1 - 1 - i + l2 - 1 - j
			product := int(num1[i] - '0') * int(num2[j] - '0')
			v, ok := values43[pos]
			if !ok {
				values43[pos] = product
			} else {
				values43[pos] = v + product
			}
		}
	}

	var b bytes.Buffer
	p := 0
	for {
		v, ok := values43[p]
		if !ok {
			break
		}
		b.WriteRune(rune(v % 10 + '0'))
		v /= 10
		for np := p + 1; v != 0; np, v = np + 1, v / 10 {
			_, ok := values43[np]
			if !ok {
				values43[np] = 0
			}
			nv, _ := values43[np]
			values43[np] = nv + v % 10
		}
		p++
	}
	runes := []rune(b.String())
	for left, right := 0, len(runes) - 1; left < right; left, right = left + 1, right - 1 {
		runes[left], runes[right] = runes[right], runes[left]
	}
	i := 0
	for i < len(runes) - 1 {
		if runes[i] == '0' {
			i++
		} else {
			break
		}
	}
	return string(runes[i:])
}