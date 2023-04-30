package solution_51_100

import (
	"fmt"
	"reflect"
	"testing"
)

func merge88(nums1 []int, m int, nums2 []int, n int)  {
	if n == 0 {
		return
	}

	var cnt1 int = 0
	var i1 int = 0
	var i2 int = 0

	for {
		if nums1[i1] > nums2[i2] {
			j := i2 + 1
			for ; j < n; j++ {
				if nums1[i1] <= nums2[j] {
					break
				}
			}
			offset := j - i2
			// fmt.Printf("%d, %d, %d, %d, %d, %d\n", m, n, i1, offset, j, cnt1)
			for x := m - cnt1 - 1; x >= 0; x-- {
				nums1[i1 + x + offset] = nums1[i1 + x]
			}
			for x := 0; x < offset; x++ {
				nums1[i1 + x] = nums2[i2 + x]
			}
			i2 = j
			i1 += offset

			if i2 >= n {
				break
			}
		}

		i1 += 1
		cnt1 += 1

		if cnt1 >= m {
			y := 1
			for x := len(nums2) - 1; x >= i2; x-- {
				nums1[len(nums1) - y] = nums2[x]
				y++
			}
			break
		}
	}
}


func Test_88(t *testing.T) {
	var nums1 []int
	var m int
	var nums2 []int
	var n int

	nums1 = []int{1,2,3,0,0,0}
	m = 3
	nums2 = []int{2,5,6}
	n = 3
	merge88(nums1, m, nums2, n)
	if !reflect.DeepEqual(nums1, []int{1,2,2,3,5,6}) {
		t.Error("failed at 1")
	} else {
		fmt.Println("success at 1")
	}

	nums1 = []int{1}
	m = 1
	nums2 = []int{}
	n = 0
	merge88(nums1, m, nums2, n)
	if !reflect.DeepEqual(nums1, []int{1}) {
		t.Error("failed at 2")
	} else {
		fmt.Println("success at 2")
	}

	nums1 = []int{1,7,8,0,0,0,0}
	m = 3
	nums2 = []int{2,5,6,10}
	n = 4
	merge88(nums1, m, nums2, n)
	if !reflect.DeepEqual(nums1, []int{1,2,5,6,7,8,10}) {
		t.Error("failed at 3")
	} else {
		fmt.Println("success at 3")
	}

	nums1 = []int{1,2,4,5,6,0}
	m = 5
	nums2 = []int{3}
	n = 1
	merge88(nums1, m, nums2, n)
	if !reflect.DeepEqual(nums1, []int{1,2,3,4,5,6}) {
		t.Error("failed at 4")
	} else {
		fmt.Println("success at 4")
	}
}