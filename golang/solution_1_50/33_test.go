package solution_1_50

import "fmt"

//33. 搜索旋转排序数组
//整数数组 nums 按升序排列，数组中的值 互不相同 。
//
//在传递给函数之前，nums 在预先未知的某个下标 k（0 <= k < nums.length）上进行了 旋转，
//使数组变为 [nums[k], nums[k+1], ..., nums[n-1], nums[0], nums[1], ..., nums[k-1]]
//（下标 从 0 开始 计数）。例如， [0,1,2,4,5,6,7] 在下标 3 处经旋转后可能变为[4,5,6,7,0,1,2] 。
//
//给你 旋转后 的数组 nums 和一个整数 target ，
//如果 nums 中存在这个目标值 target ，则返回它的下标，否则返回-1。


func search33(nums []int, target int) int {
	l := len(nums)
	if l == 0 {
		return -1
	}
	if l == 1 {
		if target == nums[0] {
			return 0
		} else {
			return -1
		}
	}

	left, right := 0, l - 1
	leftVal, rightVal := nums[0], nums[l - 1]
	if target == rightVal {
		return l - 1
	}
	if target == leftVal {
		return 0
	}
	if leftVal < rightVal { // binary search
		for right >= left {
			mid := left + (right - left) / 2
			if nums[mid] == target {
				return mid
			}
			if nums[mid] < target {
				left = mid + 1
			} else {
				right = mid - 1
			}
		}
	} else {
		if target < leftVal && target > rightVal {
			return -1
		}
		if target > leftVal {
			for nums[left] < nums[left + 1] {
				next := nums[left + 1]
				if next == target {
					return left + 1
				}
				if next < target {
					left++
				} else {
					break
				}
			}
		} else {
			for nums[right] > nums[right - 1] {
				next := nums[right - 1]
				if next == target {
					return right - 1
				}
				if next > target {
					right--
				} else {
					break
				}
				fmt.Printf("%d\n", right)
			}
		}
	}

	return -1

}
