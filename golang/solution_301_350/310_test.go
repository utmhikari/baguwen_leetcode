package solution_301_350
//
//310. 最小高度树
//树是一个无向图，其中任何两个顶点只通过一条路径连接。
//换句话说，一个任何没有简单环路的连通图都是一棵树。
//
//给你一棵包含 n 个节点的树，标记为 0 到 n - 1 。
//给定数字 n 和一个有 n - 1 条无向边的 edges 列表（每一个边都是一对标签），
//其中 edges[i] = [ai, bi] 表示树中节点 ai 和 bi 之间存在一条无向边。
//
//可选择树中任何一个节点作为根。
//当选择节点 x 作为根节点时，设结果树的高度为 h 。
//在所有可能的树中，具有最小高度的树（即，min(h)）被称为 最小高度树 。
//
//请你找到所有的 最小高度树 并按 任意顺序 返回它们的根节点标签列表。
//
//树的 高度 是指根节点和叶子节点之间最长向下路径上边的数量。


func findMinHeightTrees310(n int, edges [][]int) []int {
	conns := make(map[int][]int)
	indegrees := make([]int, n)
	leftNodes := make(map[int]bool)

	for i := 0; i < n; i++ {
		conns[i] = make([]int, 0)
		leftNodes[i] = true
	}

	for _, e := range edges {
		indegrees[e[0]]++
		indegrees[e[1]]++
		arr0, _ := conns[e[0]]
		conns[e[0]] = append(arr0, e[1])
		arr1, _ := conns[e[1]]
		conns[e[1]] = append(arr1, e[0])
	}

	lastLevel := make([]int, 0)

	for len(leftNodes) > 0 {
		// find nodes already with zero degree (the top~)
		zeroIndegreesLvl := make([]int, 0)
		for i, _ := range leftNodes {
			if indegrees[i] == 0 {
				zeroIndegreesLvl = append(zeroIndegreesLvl, i)
			}
		}
		if len(zeroIndegreesLvl) > 0 {
			return zeroIndegreesLvl
		}

		// find nodes with only one degree
		oneIndegreesLvl := make([]int, 0)
		for i, _ := range leftNodes {
			if indegrees[i] == 1 {
				oneIndegreesLvl = append(oneIndegreesLvl, i)
				delete(leftNodes, i)
			}
		}
		for _, i := range oneIndegreesLvl {
			nodes, _ := conns[i]
			for _, j := range nodes {
				if indegrees[j] > 0 {
					indegrees[j]--
				}
			}
			indegrees[i] = 0
		}
		lastLevel = oneIndegreesLvl
	}

	return lastLevel
}

