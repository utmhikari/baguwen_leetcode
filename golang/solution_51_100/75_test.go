package solution_51_100



//75. 颜色分类
//给定一个包含红色、白色和蓝色，一共n 个元素的数组，
//原地对它们进行排序，使得相同颜色的元素相邻，并按照红色、白色、蓝色顺序排列。
//
//此题中，我们使用整数 0、1 和 2 分别表示红色、白色和蓝色。



func sortColors75(nums []int)  {
	// move 0 to left & move 2 to right
	l := len(nums)
	if l <= 1 {
		return
	}

	left, right := 0, l - 1
	p := 0
	for p >= left && p <= right {
		switch nums[p] {
		case 0:
			nums[p], nums[left] = nums[left], nums[p]
			p++
			left++
		case 1:
			p++
		case 2:
			nums[p], nums[right] = nums[right], nums[p]
			right--
		default:
			break
		}
	}
}