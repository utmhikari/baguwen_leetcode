package solution_51_100


//51. N 皇后


var solutions51 [][]string
var grid51 [][]rune
var flag51 bool


func solve51(i int, n int) {
	if i == n {
		solutions51 = append(solutions51, make([]string, n))
		for x := 0; x < n; x++ {
			solutions51[len(solutions51) - 1][x] = string(grid51[x])
		}
	}
	for j := 0; j < n; j++ {  // only 1 in row
		flag51 = true
		// check prev rows
		for k := 0; k < i; k++ {
			// col
			if grid51[k][j] == 'Q' {
				flag51 = false
				break
			}
			// backslash (2,3) (1,2) (0,1)
			if k + j - i >= 0 && grid51[k][k + j - i] == 'Q' {
				flag51 = false
				break
			}
			// slash (3,1) (2,2) (1,3)
			if i + j - k < n && grid51[k][i + j - k] == 'Q' {
				flag51 = false
				break
			}
		}
		if flag51 {
			grid51[i][j] = 'Q'
			solve51(i + 1, n)
			grid51[i][j] = '.'
		}
	}
}

func solveNQueens51(n int) [][]string {
	// init
	solutions51 = make([][]string, 0)
	grid51 = make([][]rune, n)

	for i := 0; i < n; i++ {
		grid51[i] = make([]rune, n)
		for j := 0; j < n; j++ {
			grid51[i][j] = '.'
		}
	}

	solve51(0, n)

	return solutions51
}