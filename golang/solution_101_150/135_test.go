package solution_101_150

import "fmt"

//135. 分发糖果
//老师想给孩子们分发糖果，有 N个孩子站成了一条直线，老师会根据每个孩子的表现，预先给他们评分。
//
//你需要按照以下要求，帮助老师给这些孩子分发糖果：
//
//每个孩子至少分配到 1 个糖果。
//评分更高的孩子必须比他两侧的邻位孩子获得更多的糖果。
//那么这样下来，老师至少需要准备多少颗糖果呢？




func candy135(ratings []int) int {
	l := len(ratings)
	if l == 0 {
		return 0
	}
	if l == 1 {
		return 1
	}

	tmp135 := make([]int, l)
	for i := 0; i < l; i++ {
		// seek left
		left, right := 1, 1
		for j := i - 1; j >= 0; j-- {
			if ratings[j] >= ratings[j + 1] {
				break
			}
			left++
		}
		for j := i + 1; j < l; j++ {
			if ratings[j] >= ratings[j - 1] {
				break
			}
			right++
		}
		fmt.Printf("%d, %d, %d\n", i, left, right)
		tmp135[i] = left
		if right > left {
			tmp135[i] = right
		}
	}

	fmt.Printf("%v\n", tmp135)
	sm := 0
	for _, i2 := range tmp135 {
		sm += i2
	}
	return sm
}
