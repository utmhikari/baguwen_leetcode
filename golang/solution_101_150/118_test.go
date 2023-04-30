package solution_101_150



//118. 杨辉三角
//给定一个非负整数 numRows，生成杨辉三角的前 numRows 行。


func generate118(numRows int) [][]int {
	ans := make([][]int, numRows)

	if numRows == 0 {
		return ans
	}

	ans[0] = []int{1}
	for i := 1; i < numRows; i++ {
		ans[i] = make([]int, i + 1)
		for j := 0; j < i + 1; j++ {
			if j == 0 || j == i {
				ans[i][j] = 1
			} else {
				ans[i][j] = ans[i - 1][j - 1] + ans[i - 1][j]
			}
		}
	}
	return ans
}
