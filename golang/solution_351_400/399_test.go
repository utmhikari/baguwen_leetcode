package solution_351_400


//399. 除法求值
//给你一个变量对数组 equations 和一个实数值数组 values 作为已知条件，
//其中 equations[i] = [Ai, Bi] 和 values[i] 共同表示等式 Ai / Bi = values[i] 。
//每个 Ai 或 Bi 是一个表示单个变量的字符串。
//
//另有一些以数组 queries 表示的问题，
//其中 queries[j] = [Cj, Dj] 表示第 j 个问题，
//请你根据已知条件找出 Cj / Dj = ? 的结果作为答案。
//
//返回 所有问题的答案 。
//如果存在某个无法确定的答案，则用 -1.0 替代这个答案。
//如果问题中出现了给定的已知条件中没有出现的字符串，也需要用 -1.0 替代这个答案。
//
//注意：输入总是有效的。你可以假设除法运算中不会出现除数为 0 的情况，且不存在任何矛盾的结果。

var mp399 map[string]map[string]float64
var ans399 []float64


func solve399(cur float64, left string, right string, visited map[string]bool) float64 {
	if visited[left] {
		return -1.0
	}
	visited[left] = true

	value, ok := mp399[left][right]
	if ok {
		return cur * value
	}

	for s, v := range mp399[left] {
		ans := solve399(cur * v, s, right, visited)
		if ans != -1 {
			return ans
		}
	}

	visited[left] = false
	return -1.0
}

func calcEquation(equations [][]string, values []float64, queries [][]string) []float64 {
	mp399 = make(map[string]map[string]float64)

	for i, equation := range equations {
		left, right := equation[0], equation[1]
		leftValue, rightValue := values[i], 1.0 / values[i]
		_, ok := mp399[left]
		if !ok {
			mp399[left] = map[string]float64{right: leftValue}
		} else {
			mp399[left][right] = leftValue
		}
		_, ok = mp399[right]
		if !ok {
			mp399[right] = map[string]float64{left: rightValue}
		} else {
			mp399[right][left] = rightValue
		}
	}

	ans399 = make([]float64, len(queries))
	for i, query := range queries {
		left, right := query[0], query[1]
		curAns := -1.0
		_, okLeft := mp399[left]
		_, okRight := mp399[right]
		if okLeft && okRight {
			if left == right {
				curAns = 1.0
			} else {
				curAns = solve399(1.0, left, right, make(map[string]bool))
			}
		}
		ans399[i] = curAns
	}
	return ans399
}