package solution_251_300

import "fmt"

// 300. 最长递增子序列

//以输入序列 [0, 8, 4, 12, 2][0,8,4,12,2] 为例：
//
//第一步插入 0，d=[0]；
//
//第二步插入 8，d=[0,8]；
//
//第三步插入 4，d=[0,4]；
//
//第四步插入 12，d=[0,4,12]；
//
//第五步插入 2，d=[0,2,12]。


func lengthOfLIS300(nums []int) int {
	l := len(nums)
	if l <= 1 {
		return l
	}

	// 保证子序列第n个数的值尽量小，保存对应数字的索引
	minLastIndices := make([]int, l)
	// 求出第i个数的位置之后，用于回溯前一个数字索引的list
	prevIndices := make([]int, l)
	k := 0

	for i := 1; i < l; i++ {
		if nums[i] > nums[minLastIndices[k]] {
			minLastIndices[k + 1] = i
			prevIndices[i] = minLastIndices[k]
			k++
		} else {
			left, right := 0, k
			for left < right {
				mid := left + (right - left) / 2
				if nums[i] > nums[minLastIndices[mid]] {
					left = mid + 1
				} else {
					right = mid
				}
			}
			minLastIndices[left] = i
			if left > 0 {
				prevIndices[i] = minLastIndices[left - 1]
			}
		}
		fmt.Printf("minLastIndices: %v, prevIndices: %v\n", minLastIndices, prevIndices)
	}

	return k + 1
}