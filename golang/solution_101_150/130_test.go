package solution_101_150

//130. 被围绕的区域
//给你一个 m x n 的矩阵 board ，由若干字符 'X' 和 'O' ，
//找到所有被 'X' 围绕的区域，并将这些区域里所有的 'O' 用 'X' 填充。



var map130 map[int]bool
var width130 int


func getKey(x int, y int) int {
	return x *width130 + y
}

func fillMap130(board [][]byte, x int, y int) {
	if x < 0 || x >= len(board) || y < 0 || y >= len(board[0]) {
		return
	}
	if board[x][y] != 'O' {
		return
	}
	key := getKey(x, y)
	ok, _ := map130[key]
	if ok {
		return
	}
	map130[key] = true
	fillMap130(board, x - 1, y)
	fillMap130(board, x + 1, y)
	fillMap130(board, x, y - 1)
	fillMap130(board, x, y + 1)
}

func solve(board [][]byte)  {
	h := len(board)
	if h <= 2 {
		return
	}
	w := len(board[0])
	if w <= 2 {
		return
	}
	width130 = w

	map130 = make(map[int]bool)
	for i := 0; i < h; i++ {
		if board[i][0] == 'O' {
			fillMap130(board, i, 0)
		}
		if board[i][w - 1] == 'O' {
			fillMap130(board, i, w - 1)
		}
	}
	for j := 0; j < w; j++ {
		if board[0][j] == 'O' {
			fillMap130(board, 0, j)
		}
		if board[h - 1][j] == 'O' {
			fillMap130(board, h - 1, j)
		}
	}

	//fmt.Printf("%v\n", map130)

	for i := 1; i < h - 1; i++ {
		for j := 1; j < w - 1; j++ {
			if board[i][j] == 'O' {
				key := getKey(i, j)
				ok, _ := map130[key]
				if !ok {
					board[i][j] = 'X'
				}
			}
		}
	}
}
