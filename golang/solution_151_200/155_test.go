package solution_151_200

type MinStack struct {
	nums []int
}


/** initialize your data structure here. */
func Constructor155() MinStack {
	return MinStack{
		nums: make([]int, 0),
	}
}


func (ms *MinStack) Push(val int)  {
	l := len(ms.nums)

	if l == 0 {
		ms.nums = []int{val, val}
		return
	}

	curMin := ms.GetMin()

	ms.nums = append([]int{val, val}, ms.nums...)

	if curMin < val {
		ms.nums[0] = curMin
	}
}


func (ms *MinStack) Pop()  {
	if len(ms.nums) > 0 {
		ms.nums = ms.nums[2:]
	}
}


func (ms *MinStack) Top() int {
	if len(ms.nums) == 0 {
		return 0
	}

	return ms.nums[1]
}


func (ms *MinStack) GetMin() int {
	if len(ms.nums) == 0 {
		return 0
	}

	return ms.nums[0]
}

