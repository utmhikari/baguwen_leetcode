package solution_201_250


//238. 除自身以外数组的乘积

//给你一个长度为n的整数数组nums，其中n > 1，
//返回输出数组output，其中 output[i]等于nums中除nums[i]之外
//其余各元素的乘积。

//请不要使用除法，且在 O(n) 时间复杂度内完成此题。

func productExceptSelf238(nums []int) []int {
	l := len(nums)
	pLeft, pRight := make([]int, l), make([]int, l)
	pLeft[0] = 1
	pRight[l - 1] = 1
	for i := 1; i < l; i++ {
		pLeft[i] = pLeft[i - 1] * nums[i - 1]
		pRight[l - 1 - i] = pRight[l - i] * nums[l - i]
	}
	output := make([]int, l)

	for i := 0; i < l; i++ {
		output[i] = pLeft[i] * pRight[i]
	}

	return output
}