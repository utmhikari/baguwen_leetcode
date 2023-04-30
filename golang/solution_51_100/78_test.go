package solution_51_100

var ans78 [][]int
var flag78 []bool


func solve78(nums []int, cur int) {
	if cur == len(nums) {
		tmp := make([]int, 0)
		for i := 0; i < len(nums); i++ {
			if flag78[i] {
				tmp = append(tmp, nums[i])
			}
		}
		ans78 = append(ans78, tmp)
	} else {
		solve78(nums, cur + 1)
		flag78[cur] = true
		solve78(nums, cur + 1)
		flag78[cur] = false
	}
}

func subsets(nums []int) [][]int {
	ans78 = make([][]int, 0)
	l := len(nums)
	if l == 0 {
		return make([][]int, 0)
	}
	if l == 1 {
		return append(ans78, []int{}, []int{nums[0]})
	}
	flag78 = make([]bool, l)
	solve78(nums, 0)
	return ans78
}

