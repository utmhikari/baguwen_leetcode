package solution_301_350


//303. 区域和检索 - 数组不可变
//给定一个整数数组  nums，求出数组从索引 i 到 j（i ≤ j）范围内元素的总和，包含 i、j 两点。



//type NumArray struct {
//	sums []int
//}
//
//
//func Constructor303(nums []int) NumArray {
//	sums := make([]int, len(nums) + 1)
//	for i, n := range nums {
//		sums[i + 1] = sums[i] + n
//	}
//	return NumArray{sums: sums}
//}
//
//
//func (this *NumArray) SumRange(left int, right int) int {
//	return this.sums[right + 1] - this.sums[left]
//}
