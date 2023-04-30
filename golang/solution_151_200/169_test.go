package solution_151_200



func majorityElement169(nums []int) int {
	l := len(nums)
	n := nums[0]
	c := 1
	i := 1
	for ; i < l; i++ {
		if nums[i] == n {
			c++
		} else {
			c--
			if c == 0 {
				n = nums[i + 1]
				i++
				c = 1
			}
		}
	}
	return n
}