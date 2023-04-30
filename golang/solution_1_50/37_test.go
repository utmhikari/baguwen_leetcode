package solution_1_50

import "fmt"

//37. 解数独


var sudokuChars37 = []byte{
	'1', '2', '3', '4', '5', '6', '7', '8', '9',
}

func solve37(board [][]byte, x int, y int) bool {
	if y == 9 && x == 0 {
		fmt.Printf("%+v\n", board)
		return true
	}

	var chars []byte
	formerVal := board[x][y]
	if formerVal == '.' {
		chars = sudokuChars37
	} else {
		chars = []byte{board[x][y]}
	}

	for _, c := range chars {
		board[x][y] = c

		flag := true
		for i := 0; i < 9; i++ {
			if i != x && board[i][y] == board[x][y] {
				flag = false
				break
			}
			if i != y && board[x][i] == board[x][y] {
				flag = false
				break
			}
		}

		sx, sy := x / 3 * 3, y / 3 * 3
		for i := sx; i < sx + 3; i++ {
			if !flag {
				break
			}
			for j := sy; j < sy + 3; j++ {
				if !(i == x && j == y) && board[i][j] == board[x][y] {
					flag = false
					break
				}
			}
		}

		if !flag {
			continue
		}

		nx, ny := x, y
		if x < 8 {
			nx++
		} else {
			nx = 0
			ny++
		}
		if solve37(board, nx, ny) {
			return true
		}
	}

	board[x][y] = formerVal
	return false
}

func solveSudoku37(board [][]byte) {
	solve37(board, 0, 0)
}

