package solution_51_100


//72. 编辑距离


func minDistance(word1 string, word2 string) int {
	l1, l2 := len(word1), len(word2)
	if l1 == 0 {
		return l2
	}
	if l2 == 0 {
		return l1
	}

	dp := make([][]int, l1 + 1)
	for i := 0; i <= l1; i++ {
		dp[i] = make([]int, l2 + 1)
	}
	for i := 0; i <= l1; i++ {
		dp[i][0] = i
	}
	for i := 0; i <= l2; i++ {
		dp[0][i] = i
	}
	var minVal int
	for i := 1; i <= l1; i++ {
		for j := 1; j <= l2; j++ {
			if word1[i - 1] == word2[j - 1] {
				dp[i][j] = dp[i - 1][j - 1]
			} else {
				minVal = dp[i - 1][j]
				if minVal > dp[i][j - 1] {
					minVal = dp[i][j - 1]
				}
				if minVal > dp[i - 1][j - 1] {
					minVal = dp[i - 1][j - 1]
				}
				dp[i][j] = minVal + 1
			}
		}
	}
	return dp[l1][l2]
}