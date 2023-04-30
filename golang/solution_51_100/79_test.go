package solution_51_100


//79. 单词搜索
//给定一个m x n 二维字符网格board 和一个字符串单词word 。
//如果word 存在于网格中，返回 true ；否则，返回 false 。
//
//单词必须按照字母顺序，通过相邻的单元格内的字母构成，其中“相邻”单元格是那些水平相邻或垂直相邻的单元格。
//同一个单元格内的字母不允许被重复使用。



var visited79 [][]bool

func dfs79(board [][]byte, word string, i int, j int, cur int) bool {
	if cur == len(word) {
		return true
	}
	if i < 0 || i >= len(board) {
		return false
	}
	if j < 0 || j >= len(board[0]) {
		return false
	}
	if visited79[i][j] {
		return false
	}
	if word[cur] != board[i][j] {
		return false
	}

	visited79[i][j] = true
	if dfs79(board, word, i - 1, j, cur + 1) {
		return true
	}
	if dfs79(board, word, i + 1, j, cur + 1) {
		return true
	}
	if dfs79(board, word, i, j - 1, cur + 1) {
		return true
	}
	if dfs79(board, word, i, j + 1, cur + 1) {
		return true
	}
	visited79[i][j] = false
	return false
}


func exist79(board [][]byte, word string) bool {
	lw := len(word)
	if lw == 0 {
		return false
	}

	h := len(board)
	if h == 0 {
		return false
	}
	w := len(board[0])
	if w == 0 {
		return false
	}

	visited79 = make([][]bool, h)
	for i := 0; i < h; i++ {
		visited79[i] = make([]bool, w)
	}

	for i := 0; i < h; i++ {
		for j := 0; j < w; j++ {
			if dfs79(board, word, i, j, 0) {
				return true
			}
		}
	}

	return false
}