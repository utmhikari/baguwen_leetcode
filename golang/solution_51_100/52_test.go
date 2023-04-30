package solution_51_100



var grid52 [][]rune
var flag52 bool
var ans52 int


func solve52(i int, n int) {
	if i == n {
		ans52++
	}
	for j := 0; j < n; j++ {  // only 1 in row
		flag52 = true
		// check prev rows
		for k := 0; k < i; k++ {
			// col
			if grid52[k][j] == 'Q' {
				flag52 = false
				break
			}
			// backslash (2,3) (1,2) (0,1)
			if k + j - i >= 0 && grid52[k][k + j - i] == 'Q' {
				flag52 = false
				break
			}
			// slash (3,1) (2,2) (1,3)
			if i + j - k < n && grid52[k][i + j - k] == 'Q' {
				flag52 = false
				break
			}
		}
		if flag52 {
			grid52[i][j] = 'Q'
			solve52(i + 1, n)
			grid52[i][j] = '.'
		}
	}
}

func totalNQueens(n int) int {
	// init
	grid52 = make([][]rune, n)
	ans52 = 0

	for i := 0; i < n; i++ {
		grid52[i] = make([]rune, n)
		for j := 0; j < n; j++ {
			grid52[i][j] = '.'
		}
	}

	solve52(0, n)

	return ans52
}
