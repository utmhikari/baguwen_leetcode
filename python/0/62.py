class Solution:
    def uniquePaths(self, m: int, n: int) -> int:
        if m == 0 or n == 0:
            return 0
        if m == 1 or n == 1:
            return 1
        dp = []
        for i in range(m):
            if i == 0:
                dp.append([1] * n)
            else:
                dp.append([0] * n)
                dp[-1][0] = 1
        for i in range(1, m):
            for j in range(1, n):
                dp[i][j] = dp[i - 1][j] + dp[i][j - 1]
        return dp[m - 1][n - 1]



s = Solution()
print(s.uniquePaths(3,2))
print(s.uniquePaths(7,3))