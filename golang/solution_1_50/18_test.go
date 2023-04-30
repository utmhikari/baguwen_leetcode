package solution_1_50

import "sort"

//18. 四数之和
//给定一个包含n 个整数的数组nums和一个目标值target，
//判断nums中是否存在四个元素 a，b，c和 d，
//使得a + b + c + d的值与target相等？找出所有满足条件且不重复的四元组。
//
//注意：答案中不可以包含重复的四元组。


var ans18 [][]int


func fourSum18(nums []int, target int) [][]int {
	ans18 = make([][]int, 0)
	l := len(nums)
	if l < 4 {
		return ans18
	}

	sort.Ints(nums)
	for i := 0; i < l - 3; i++ {
		if i > 0 && nums[i] == nums[i - 1] {
			continue
		}
		if nums[i] + nums[i + 1] + nums[i + 2] + nums[i + 3] > target {
			break
		}
		if nums[i] + nums[l - 3] + nums[l - 2] + nums[l - 1] < target {
			continue
		}
		for j := i + 1; j < l - 2; j++ {
			if j > i + 1 && nums[j] == nums[j - 1] {
				continue
			}
			if nums[i] + nums[j] + nums[j + 1] + nums[j + 2] > target {
				break
			}
			if nums[i] + nums[j] + nums[l - 2] + nums[l - 1] < target {
				continue
			}
			left, right := j + 1, l - 1
			for left < right {
				sm := nums[i] + nums[j] + nums[left] + nums[right]
				if sm == target {
					ans18 = append(ans18, []int{nums[i], nums[j], nums[left], nums[right]})
					for left < right && nums[left] == nums[left + 1] {
						left++
					}
					left++
					for left < right && nums[right] == nums[right - 1] {
						right--
					}
					right--
				} else if sm > target {
					right--
				} else {
					left++
				}
			}
		}
	}

	return ans18
}