package solution_101_150


//128. 最长连续序列
//给定一个未排序的整数数组 nums ，找出数字连续的最长序列
//（不要求序列元素在原数组中连续）的长度。


func longestConsecutive128(nums []int) int {
	l := len(nums)
	if l == 0 {
		return 0
	}
	headMap := make(map[int]int)
	tailMap := make(map[int]int)
	visited := make(map[int]bool)

	for _, num := range nums {
		_, isVisited := visited[num]
		if isVisited {
			continue
		}
		visited[num] = true

		left, right := num - 1, num + 1
		head, tailOk := tailMap[left]
		tail, headOk := headMap[right]

		if !tailOk && !headOk {
			tailMap[num] = num
			headMap[num] = num
		} else if tailOk && !headOk {
			tailMap[num] = head
			headMap[head] = num
			delete(tailMap, left)
		} else if headOk && !tailOk {
			headMap[num] = tail
			tailMap[tail] = num
			delete(headMap, right)
		} else {
			headMap[head] = tail
			tailMap[tail] = head
			delete(headMap, right)
			delete(tailMap, left)
		}
	}

	maxLen := 1
	for k, v := range headMap {
		curLen := v - k + 1
		if curLen > maxLen {
			maxLen = curLen
		}
	}
	return maxLen
}