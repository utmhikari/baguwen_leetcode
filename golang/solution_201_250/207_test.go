package solution_201_250



//207. 课程表
//你这个学期必须选修 numCourses 门课程，记为0到numCourses - 1 。
//
//在选修某些课程之前需要一些先修课程。 先修课程按数组prerequisites 给出，
//其中prerequisites[i] = [ai, bi] ，表示如果要学习课程ai 则 必须 先学习课程 bi 。
//
//例如，先修课程对[0, 1] 表示：想要学习课程 0 ，你需要先完成课程 1 。
//请你判断是否可能完成所有课程的学习？如果可以，返回 true ；否则，返回 false 。


var mp207 map[int]map[int]bool

func canFinish207(numCourses int, prerequisites [][]int) bool {
	mp207 = make(map[int]map[int]bool)

	// fill all
	for _, prq := range prerequisites {
		st, ok := mp207[prq[0]]
		if !ok {
			mp207[prq[0]] = map[int]bool{
				prq[1]: true,
			}
		} else {
			st[prq[1]] = true
		}
	}



	for len(mp207) > 0 {
		prevLen := len(mp207)
		for _, pres := range mp207 {
			for pi, _ := range pres {
				_, ok := mp207[pi]
				if !ok {
					delete(pres, pi)
				}
			}
		}
		for i, pres := range mp207 {
			if len(pres) == 0 {
				delete(mp207, i)
			}
		}
		curLen := len(mp207)
		if curLen == prevLen {
			return false
		}
	}

	return true
}
