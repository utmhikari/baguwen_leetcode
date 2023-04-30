package solution_251_300

import (
	"bytes"
	"fmt"
	"reflect"
	"strconv"
	"testing"
)

//282. 给表达式添加运算符
//给定一个仅包含数字 0-9 的字符串和一个目标值，
//在数字之间添加 二元 运算符（不是一元）+、- 或 * ，
//返回所有能够得到目标值的表达式。



//var cache282 map[string][]string
//
//
//func solve282(num string, target int, canPlusOrMinus bool) []string {
//	key := fmt.Sprintf("%s_%d", num, target)
//	cached, ok := cache282[key]
//	if ok {
//		return cached
//	}
//
//	ans := make([]string, 0)
//
//	// self
//	l := len(num)
//	if !(l > 1 && num[0] == '0') {
//		numInt, _ := strconv.Atoi(num)
//		if numInt == target {
//			ans = append(ans, num)
//		}
//	}
//
//	// each idx
//
//	left := 0
//	var rightStr string
//	for i := 1; i < l; i++ {
//		if num[0] == '0' && i > 1 {
//			break
//		}
//
//		left = left * 10 + int(num[i - 1] - '0')
//		rightStr = num[i:]
//		if canPlusOrMinus {
//			// plus
//			for _, rightAns := range solve282(rightStr, target - left, true) {
//				ans = append(ans, fmt.Sprintf("%d+%s", left, rightAns))
//			}
//			// minus
//			for _, rightAns := range solve282(rightStr, left - target, true) {
//				ans = append(ans, fmt.Sprintf("%d-%s", left, rightAns))
//			}
//		}
//		// multiply
//		j, p := i, 0
//		for {
//			p = p * 10 + int(num[j] - '0')
//			leftP := left * p
//			if j == l - 1 {
//				if leftP == target && !(i < l - 1 && num[i] == '0') {
//					ans = append(ans, fmt.Sprintf("%d*%d", left, p))
//				}
//				break
//			} else {
//				rightStr = num[j + 1:]
//				if canPlusOrMinus {
//					// plus
//					for _, rightAns := range solve282(rightStr, target - leftP, true) {
//						ans = append(ans, fmt.Sprintf("%d*%d+%s", left, p, rightAns))
//					}
//					// minus
//					for _, rightAns := range solve282(rightStr, leftP - target, true) {
//						ans = append(ans, fmt.Sprintf("%d*%d-%s", left, p, rightAns))
//					}
//				}
//				// multiply
//				if leftP != 0 && target % leftP == 0 {
//					for _, rightAns := range solve282(rightStr, target / leftP, false) {
//						ans = append(ans, fmt.Sprintf("%d*%d*%s", left, p, rightAns))
//					}
//				}
//			}
//			j++
//		}
//	}
//
//	cache282[key] = ans
//	return ans
//}
//
//func addOperators282(num string, target int) []string {
//	cache282 = make(map[string][]string)
//	return solve282(num, target, true)
//}
//
//


var ans282 []string
var ops282 []string
var digits282 string
var target282 int


func recurse282(idx int, prevOperand int, curOperand int, value int) {
	nums := digits282

	// processed all digits
	if idx == len(nums) {
		// add all ops and nums
		if value == target282 && curOperand == 0 {
			var b bytes.Buffer
			for i := 1; i < len(ops282); i++ {
				b.WriteString(ops282[i])
			}
			ans282 = append(ans282, b.String())
		}
		return
	}

	// extend current operand by 1 digit
	curOperand = curOperand * 10 + int(nums[idx] - '0')
	if curOperand > 0 {
		recurse282(idx + 1, prevOperand, curOperand, value)
	}

	curOperandStr := strconv.Itoa(curOperand)

	// add
	// if len(ops282) == 0, meaning as 0 + xxxxxxxxxxx
	ops282 = append(ops282, "+")
	ops282 = append(ops282, curOperandStr)
	recurse282(idx + 1, curOperand, 0, value + curOperand)
	ops282 = ops282[:len(ops282) - 2]

	if len(ops282) > 0 {
		// minus
		ops282 = append(ops282, "-")
		ops282 = append(ops282, curOperandStr)
		recurse282(idx + 1, -curOperand, 0, value - curOperand)
		ops282 = ops282[:len(ops282) - 2]

		// multiple
		ops282 = append(ops282, "*")
		ops282 = append(ops282, curOperandStr)
		recurse282(idx + 1, prevOperand * curOperand, 0,
			value - prevOperand + (prevOperand * curOperand))
		ops282 = ops282[:len(ops282) - 2]
	}
}


func addOperators282(num string, target int) []string {
	if len(num) == 0 {
		return []string{}
	}

	target282 = target
	digits282 = num
	ans282 = make([]string, 0)
	ops282 = make([]string, 0)
	recurse282(0, 0, 0, 0)
	return ans282
}


type test282 struct {
	n string
	t int
	expected []string
}


func Test_282(t *testing.T) {
	cases := []test282{
		{
			"123", 6, []string{"1+2+3", "1*2*3"},
		},
		{
			"232", 8, []string{"2*3+2", "2+3*2"},
		},
		{
			"105", 5, []string{"1*0+5", "10-5"},
		},
		{
			"00", 0, []string{"0+0", "0-0", "0*0"},
		},
		{
			"3456237490", 9191, []string{},
		},
	}

	for i, c := range cases {
		a := addOperators282(c.n, c.t)
		if !reflect.DeepEqual(a, c.expected) {
			t.Errorf("failed at %d -> ans282: %v, case %v+\n", i, a, c)
		} else {
			fmt.Printf("success at %d -> ans282: %v\n", i, a)
		}
	}
}