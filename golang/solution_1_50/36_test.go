package solution_1_50

import "fmt"

//36. 有效的数独
//请你判断一个9x9 的数独是否有效。只需要 根据以下规则 ，验证已经填入的数字是否有效即可。
//
//数字1-9在每一行只能出现一次。
//数字1-9在每一列只能出现一次。
//数字1-9在每一个以粗实线分隔的3x3宫内只能出现一次。（请参考示例图）
//数独部分空格内已填入了数字，空白格用'.'表示。

//已经填入的数字 ==

var sudokuChars36 = []byte{
	'1', '2', '3', '4', '5', '6', '7', '8', '9',
}

func solve36(board [][]byte, x int, y int) bool {
	if y == 9 && x == 0 {
		fmt.Printf("%+v\n", board)
		return true
	}

	//var chars []byte
	//formerVal := board[x][y]
	//if formerVal == '.' {
	//	chars = sudokuChars36
	//} else {
	//	chars = []byte{board[x][y]}
	//}
	if board[x][y] != '.' {
		for i := 0; i < 9; i++ {
			if i != x && board[i][y] == board[x][y] {
				return false
			}
			if i != y && board[x][i] == board[x][y] {
				return false
			}
		}

		sx, sy := x / 3 * 3, y / 3 * 3
		for i := sx; i < sx + 3; i++ {
			for j := sy; j < sy + 3; j++ {
				if !(i == x && j == y) && board[i][j] == board[x][y] {
					return false
				}
			}
		}
	}

	nx, ny := x, y
	if x < 8 {
		nx++
	} else {
		nx = 0
		ny++
	}

	return solve36(board, nx, ny)

	//for _, c := range chars {
	//	board[x][y] = c
	//
	//	flag := true
	//	for i := 0; i < 9; i++ {
	//		if i != x && board[i][y] == board[x][y] {
	//			flag = false
	//			break
	//		}
	//		if i != y && board[x][i] == board[x][y] {
	//			flag = false
	//			break
	//		}
	//	}
	//
	//	sx, sy := x / 3 * 3, y / 3 * 3
	//	for i := sx; i < sx + 3; i++ {
	//		if !flag {
	//			break
	//		}
	//		for j := sy; j < sy + 3; j++ {
	//			if !(i == x && j == y) && board[i][j] == board[x][y] {
	//				flag = false
	//				break
	//			}
	//		}
	//	}
	//
	//	if !flag {
	//		continue
	//	}
	//
	//	nx, ny := x, y
	//	if x < 8 {
	//		nx++
	//	} else {
	//		nx = 0
	//		ny++
	//	}
	//	if solve36(board, nx, ny) {
	//		return true
	//	}
	//}

	//board[x][y] = formerVal
	//return false
}

func isValidSudoku36(board [][]byte) bool {
	return solve36(board, 0, 0)
}