package solution_1_50

import "fmt"

//4. 寻找两个正序数组的中位数
//给定两个大小分别为 m 和 n 的正序（从小到大）数组 nums1 和 nums2。
//请你找出并返回这两个正序数组的 中位数 。

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	l1, l2 := len(nums1), len(nums2)
	if l1 == 0 && l2 == 0 {
		return 0
	}
	if l1 == 0 {
		if l2 % 2 == 0 {
			return (float64(nums2[(l2 - 1) / 2]) + float64(nums2[1 + (l2 - 1) / 2])) / 2
		} else {
			return float64(nums2[(l2 - 1) / 2])
		}
	}
	if l2 == 0 {
		if l1 % 2 == 0 {
			return (float64(nums1[(l1 - 1) / 2]) + float64(nums1[1 + (l1 - 1) / 2])) / 2
		} else {
			return float64(nums1[(l1 - 1) / 2])
		}
	}

	p1, p2, p3, p4 := 0, l1 - 1, 0, l2 - 1
	cnt := 0
	var c1 int
	var c2 int
	for {
		if p1 <= p2 && p3 <= p4 {
			if nums1[p1] <= nums2[p3] {
				c1 = nums1[p1]
				p1++
			} else {
				c1 = nums2[p3]
				p3++
			}
		} else if p1 > p2 {
			c1 = nums2[p3]
			p3++
		} else if p3 > p4 {
			c1 = nums1[p1]
			p1++
		}
		cnt++
		if cnt == l1 + l2 {
			return float64(c1)
		}
		if p1 <= p2 && p3 <= p4 {
			if nums1[p2] >= nums2[p4] {
				c2 = nums1[p2]
				p2--
			} else {
				c2 = nums2[p4]
				p4--
			}
		} else if p1 > p2 {
			c2 = nums2[p4]
			p4--
		} else if p3 > p4 {
			c2 = nums1[p2]
			p2--
		}
		cnt++
		if cnt == l1 + l2 {
			return (float64(c1) + float64(c2)) / 2
		}
		fmt.Printf("cnt: %d, c1: %d, c2: %d\n", cnt, c1, c2)
	}
}