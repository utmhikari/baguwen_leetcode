package solution_351_400

import "math/rand"

//384. 打乱数组
//给你一个整数数组 nums ，设计算法来打乱一个没有重复元素的数组。
//
//实现 Solution384 class:
//
//Solution384(int[] nums) 使用整数数组 nums 初始化对象
//int[] reset() 重设数组到它的初始状态并返回
//int[] shuffle() 返回数组随机打乱后的结果



type Solution384 struct {
	arr []int
	shuf []int
}


func Constructor384(nums []int) Solution384 {
	l := len(nums)
	s := Solution384{
		arr:     make([]int, l),
		shuf: make([]int, l),
	}
	copy(s.arr, nums)
	copy(s.shuf, nums)

	return s
}


/** Resets the array to its original configuration and return it. */
func (s *Solution384) Reset() []int {
	return s.arr
}


/** Returns a random shuffling of the array. */
func (s *Solution384) Shuffle() []int {
	l := len(s.arr)
	for i := 0; i < l; i++ {
		ri := rand.Intn(l)
		s.shuf[i], s.shuf[ri] = s.shuf[ri], s.shuf[i]
	}
	return s.shuf
}
