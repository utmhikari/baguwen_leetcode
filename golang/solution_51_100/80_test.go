package solution_51_100

import (
	"fmt"
	"testing"
)

// 80. 删除排序数组中的重复项 II
//给定一个增序排列数组 nums ，你需要在 原地 删除重复出现的元素，使得每个元素最多出现两次，返回移除后数组的新长度。
//不要使用额外的数组空间，你必须在 原地 修改输入数组 并在使用 O(1) 额外空间的条件下完成。


//说明：
//为什么返回数值是整数，但输出的答案是数组呢？
//请注意，输入数组是以“引用”方式传递的，这意味着在函数里修改输入数组对于调用者是可见的。
//你可以想象内部操作如下：
// nums 是以“引用”方式传递的。也就是说，不对实参做任何拷贝
// int len = removeDuplicates(nums)

// 在函数里修改输入数组对于调用者是可见的。
// 根据你的函数返回的长度, 它会打印出数组中该长度范围内的所有元素。
//for (int i = 0; i < len; i++) {
//	print(nums[i])
//}



func removeDuplicates(nums []int) int {
	l := len(nums)
	if l == 0 {
		return 0
	}

	offset := 0
	i := 0
	var j int
	var idx int
	var cur int
	var cnt int
	for {
		idx = i + offset
		//fmt.Printf("%d, %d, %d\n", i, offset, idx)
		if idx >= l {
			break
		}
		if idx == 0 || cur != nums[idx] {
			cur = nums[idx]
			cnt = 1
		} else {
			cnt++
			if cnt > 2 {
				j = idx + 1
				offset++
				for ; j < l; j++ {
					if nums[j] != cur {
						break
					}
					offset++
				}
				i = j - offset
				continue
			}
		}
		nums[i] = nums[idx]
		i++
	}
	return i
}


type test80 struct {
	nums []int
	expected []int
	expectedNum int
}


func Test_80(t *testing.T) {
	cases := append(make([]test80, 0),
		test80{
			nums: []int{1,1,1,2,2,3},
			expected: []int{1,1,2,2,3},
			expectedNum: 5,
		},
		test80{
			nums: []int{0,0,1,1,1,1,2,3,3},
			expected: []int{0,0,1,1,2,3,3},
			expectedNum: 7,
		})

	var ans int

	for _, c := range cases {
		ans = removeDuplicates(c.nums)
		if ans != c.expectedNum {
			t.Errorf("failed at %v+, ans: %d\n", c, ans)
			return
		}

		for i := 0; i < ans; i++ {
			if c.nums[i] != c.expected[i] {
				t.Errorf("failed at %v+\n", c)
				return
			}
		}

		fmt.Printf("success at %v+\n", c)
	}
}
