package solution_201_250



//210. 课程表 II
//现在你总共有 n 门课需要选，记为0到n-1。
//
//在选修某些课程之前需要一些先修课程。例如，想要学习课程 0 ，你需要先完成课程1 ，我们用一个匹配来表示他们: [0,1]
//
//给定课程总量以及它们的先决条件，返回你为了学完所有课程所安排的学习顺序。
//
//可能会有多个正确的顺序，你只要返回一种就可以了。如果不可能完成所有课程，返回一个空数组。


func findOrder210(numCourses int, prerequisites [][]int) []int {
	indegrees := make([]int, numCourses)
	adjacents := make(map[int][]int)
	orders := make([]int, 0)
	queue := make([]int, 0)

	for _, p := range prerequisites {
		indegrees[p[0]]++
		arr, ok := adjacents[p[1]]
		if !ok {
			adjacents[p[1]] = []int{p[0]}
		} else {
			adjacents[p[1]] = append(arr, p[0])
		}
	}

	for i, indegree := range indegrees {
		if indegree == 0 {
			orders = append(orders, i)
			queue = append(queue, i)
		}
	}

	if len(queue) == 0 {
		return []int{}
	}

	for len(queue) > 0 {
		pre := queue[0]
		queue = queue[1:]
		numCourses--

		curs, ok := adjacents[pre]
		if ok {
			for _, cur := range curs {
				indegrees[cur]--
				if indegrees[cur] == 0 {
					queue = append(queue, cur)
					orders = append(orders, cur)
				}
			}
		}
	}

	if numCourses > 0 {
		return []int{}
	} else {
		return orders
	}
}

