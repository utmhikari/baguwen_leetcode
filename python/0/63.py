from typing import List


class Solution:
    def uniquePathsWithObstacles(self, obstacleGrid: List[List[int]]) -> int:
        w = len(obstacleGrid)
        if w == 0:
            return 0
        h = len(obstacleGrid[0])
        if h == 0:
            return 0
        if obstacleGrid[0][0] == 1:
            return 0
        dp = []
        for _ in range(w + 1):
            dp.append([0] * (h + 1))
        dp[1][1] = 1
        for i in range(1, w + 1):
            for j in range(1, h + 1):
                if i == 1 and j == 1:
                    continue
                if obstacleGrid[i - 1][j - 1] != 1:
                    dp[i][j] = dp[i - 1][j] + dp[i][j - 1]
        return dp[w][h]

s = Solution()
print(s.uniquePathsWithObstacles([
    [0,0,0],
    [0,1,0],
    [0,0,0]
]))