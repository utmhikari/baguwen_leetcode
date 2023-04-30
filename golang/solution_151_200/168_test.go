package solution_151_200

import (
	"fmt"
	"testing"
)

//168. Excel表列名称
//
//给定一个正整数，返回它在 Excel 表中相对应的列名称
//
//1 -> A
//2 -> B
//3 -> C
//...
//26 -> Z
//27 -> AA
//28 -> AB




func convertToTitle168(columnNumber int) string {
	n := columnNumber
	var s = ""
	for {
		n--
		md := n % 26
		n = n / 26
		s = string(rune(int('A') + md)) + s
		if n == 0 {
			break
		}
	}
	return s
}

type test168 struct {
	colNum int
	expected string
}


func Test_168(t *testing.T) {
	inputs := []test168{
		{colNum: 1, expected: "A"},
		{colNum: 26, expected: "Z"},
		{colNum: 27, expected: "AA"},
		{colNum: 28, expected: "AB"},
		{colNum: 701, expected: "ZY"},
	}

	for _, input := range inputs {
		ans := convertToTitle168(input.colNum)
		if ans != input.expected {
			t.Errorf("failed at %d, expected %s, got %s\n", input.colNum, input.expected, ans)
		} else {
			fmt.Printf("success at %+v\n", input)
		}
	}
}