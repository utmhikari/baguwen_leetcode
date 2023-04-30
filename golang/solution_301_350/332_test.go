package solution_301_350

import "sort"

//332. 重新安排行程
//给定一个机票的字符串二维数组 [from, to]，
//子数组中的两个成员分别表示飞机出发和降落的机场地点，
//对该行程进行重新规划排序。
//所有这些机票都属于一个从 JFK（肯尼迪国际机场）出发的先生，所以该行程必须从 JFK 开始。


// location and indices
var numLocs332 int
var locToIdx332 map[string]int
var idxToLoc332 []string
var startIdx int

// destination map
var destMap332 map[int][]int

// answers for dfs
var ans332 []int

// dfs
func dfs332(curIdx int) {
	for {
		v, ok := destMap332[curIdx]
		if !ok || len(v) == 0 {
			break
		}

		tmp := v[0]
		destMap332[curIdx] = v[1:]
		dfs332(tmp)
	}
	ans332 = append(ans332, curIdx)
}

func findItinerary(tickets [][]string) []string {
	// generate caches
	numLocs332 = 0
	startIdx = -1
	locToIdx332 = make(map[string]int)
	destMap332 = make(map[int][]int)
	for _, ticket := range tickets {
		src, dest := ticket[0], ticket[1]

		// loc -> idx, idx -> loc
		srcIdx, ok := locToIdx332[src]
		if !ok {
			srcIdx = numLocs332
			numLocs332++
			locToIdx332[src] = srcIdx
		}
		destIdx, ok := locToIdx332[dest]
		if !ok {
			destIdx = numLocs332
			numLocs332++
			locToIdx332[dest] = destIdx
		}

		// destmap
		dests, ok := destMap332[srcIdx]
		if !ok {
			destMap332[srcIdx] = []int{destIdx}
		} else {
			destMap332[srcIdx] = append(dests, destIdx)
		}
	}
	idxToLoc332 = make([]string, numLocs332)
	for loc, idx := range locToIdx332 {
		idxToLoc332[idx] = loc
	}
	for _, dests := range destMap332 {
		// sort by alphabetical order
		sort.Slice(dests, func(i int, j int) bool{
			return idxToLoc332[dests[i]] < idxToLoc332[dests[j]]
		})
	}

	// ready for dfs
	startIdx, _ = locToIdx332["JFK"]
	ans332 = make([]int, 0)
	dfs332(startIdx)

	// generate ans
	ans := make([]string, len(ans332))
	for i := 0; i < len(ans332); i++ {
		loc := idxToLoc332[ans332[len(ans332) - 1 - i]]
		ans[i] = loc
	}
	return ans
}
