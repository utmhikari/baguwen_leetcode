package solution_51_100

import (
	"fmt"
	"strconv"
	"testing"
)

// 81. 搜索旋转排序数组 II
//假设按照升序排序的数组在预先未知的某个点上进行了旋转。
//
//( 例如，数组0,0,1,2,2,5,6]可能变为[2,5,6,0,0,1,2])。
//
//编写一个函数来判断给定的目标值是否存在于数组中。若存在返回true，否则返回false。



func search81(nums []int, target int) bool {
	l := len(nums)
	if l == 0 {
		return false
	}
	left, right := 0, l - 1
	if nums[left] >= nums[right] {
		// rotated
		if target == nums[left] || target == nums[right] {
			return true
		}
		if target > nums[left] {
			for {
				left++
				if left >= l - 1 || nums[left] < nums[left - 1] || nums[left] > target {
					return false
				}
				if nums[left] == target {
					return true
				}
			}
		}
		if target < nums[right] {
			for {
				right--
				if right <= 0 || nums[right] > nums[right + 1] || nums[right] < target {
					return false
				}
				if nums[right] == target {
					return true
				}
			}
		}
		return false
	} else {
		// not rotated
		var mid int
		for {
			if right < left || left >= l {
				return false
			}
			mid = left + (right - left) / 2
			if nums[mid] == target {
				return true
			}
			if nums[mid] < target {
				left = mid + 1
			} else {
				right = mid - 1
			}
		}
	}
}


type test81 struct {
	nums []int
	target int
	expected bool
}


func Test_81 (t *testing.T) {
	cases := append(make([]test81, 0),
		test81{
			nums:     []int{2,5,6,0,0,1,2},
			target:   0,
			expected: true,
		},
		test81{
			nums:     []int{2,5,6,0,0,1,2},
			target:   3,
			expected: false,
		},
		test81{
			nums:     []int{1,3},
			target:   1,
			expected: true,
		},
	)

	var ans bool

	for _, c := range cases {
		ans = search81(c.nums, c.target)
		if ans != c.expected {
			t.Errorf("failed at %v+, ans: %s\n", c, strconv.FormatBool(ans))
		} else {
			fmt.Printf("success at %v+\n", c)
		}
	}
}




