package solution_101_150



//137.只出现一次的数字 II
//其余每个元素均出现了三次。找出那个只出现了一次的元素


func singleNumber137(nums []int) int {
	f1, f2 := 0, 0
	for _, n := range nums {
		f1 = ^f2 & (f1 ^ n)
		f2 = ^f1 & (f2 ^ n)
	}
	return f1
}
