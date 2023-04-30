package solution_351_400

import "math/rand"

//398. 随机数索引
//给定一个可能含有重复元素的整数数组，要求随机输出给定的数字的索引。 您可以假设给定的数字一定存在于数组中。
//
//注意：
//数组大小可能非常大。 使用太多额外空间的解决方案将不会通过测试。




type Solution struct {
	mp map[int][]int
}


func Constructor(nums []int) Solution {
	s := Solution{
		mp: make(map[int][]int),
	}
	for i, n := range nums {
		list, ok := s.mp[n]
		if !ok {
			s.mp[n] = []int{i}
		} else {
			s.mp[n] = append(list, i)
		}
	}
	return s
}


func (s *Solution) Pick(target int) int {
	list, ok := s.mp[target]
	if !ok {
		return -1
	}
	return list[rand.Intn(len(list))]
}


/**
 * Your Solution object will be instantiated and called as such:
 * obj := Constructor(nums);
 * param_1 := obj.Pick(target);
 */