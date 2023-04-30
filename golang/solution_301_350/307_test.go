package solution_301_350

//307. 区域和检索 - 数组可修改
//给你一个数组 nums ，请你完成两类查询，
//其中一类查询要求更新数组下标对应的值，
//另一类查询要求返回数组中某个范围内元素的总和。
//
//实现 NumArray 类：
//
//NumArray(int[] nums) 用整数数组 nums 初始化对象
//void update(int index, int val) 将 nums[index] 的值更新为 val
//int sumRange(int left, int right) 返回子数组 nums[left, right] 的总和
//（即，nums[left] + nums[left + 1], ..., nums[right]）


// 线段树解法

type NumArray struct {
	tree []int
	n int
}


func Constructor307(nums []int) NumArray {
	l := len(nums)
	if l > 0 {
		tree := make([]int, l * 2)
		i, j := l, 0
		for i < 2 * l {
			tree[i] = nums[j]
			i++
			j++
		}
		i = l - 1
		for i > 0 {
			tree[i] = tree[i * 2] + tree[i * 2 + 1]
			i--
		}
		return NumArray{
			tree: tree,
			n:    l,
		}
	} else {
		return NumArray{
			tree: nil,
			n:    0,
		}
	}
}


func (this *NumArray) Update(index int, val int)  {
	index += this.n
	this.tree[index] = val
	for index > 0 {
		left, right := index, index
		if index % 2 == 0 {
			right = index + 1
		} else {
			left = index - 1
		}

		nextIdx := index / 2
		this.tree[nextIdx] = this.tree[left] + this.tree[right]
		index = nextIdx
	}
}


func (this *NumArray) SumRange(left int, right int) int {
	left += this.n
	right += this.n
	sum := 0

	for left <= right {
		if left % 2 == 1 {
			sum += this.tree[left]
			left++
		}
		if right % 2 == 0 {
			sum += this.tree[right]
			right--
		}
		left /= 2
		right /= 2
	}

	return sum
}