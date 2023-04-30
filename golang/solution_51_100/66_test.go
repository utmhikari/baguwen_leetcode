package solution_51_100

import (
	"fmt"
	"reflect"
	"testing"
)

func plusOne(digits []int) []int {
	hasOne := true

	for i := len(digits) - 1; i >= 0; i-- {
		if digits[i] == 9 {
			digits[i] = 0
		} else {
			digits[i] += 1
			hasOne = false
			break
		}
	}

	if hasOne {
		return append([]int{1}, digits...)
	} else {
		return digits
	}
}


func Test_66(t *testing.T) {
	var digits []int
	var ans []int
	var expected []int

	digits = []int{1,2,3}
	expected = []int{1,2,4}
	ans = plusOne(digits)
	if !reflect.DeepEqual(ans, expected) {
		t.Error(fmt.Sprintf("failed at 1, %v", ans))
	} else {
		fmt.Println("success at 1")
	}

	digits = []int{4,3,2,1}
	expected = []int{4,3,2,2}
	ans = plusOne(digits)
	if !reflect.DeepEqual(ans, expected) {
		t.Error(fmt.Sprintf("failed at 2, %v", ans))
	} else {
		fmt.Println("success at 2")
	}

	digits = []int{0}
	expected = []int{1}
	ans = plusOne(digits)
	if !reflect.DeepEqual(ans, expected) {
		t.Error(fmt.Sprintf("failed at 3, %v", ans))
	} else {
		fmt.Println("success at 3")
	}
}