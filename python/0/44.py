class Solution:
    def isMatch(self, s: str, p: str) -> bool:
        ls = len(s)
        lp = len(p)
        if ls == 0 and lp == 0:
            return True
        dp = []
        for _ in range(ls + 1):
            dp.append([False] * (lp + 1))
        dp[0][0] = True
        for j in range(lp):
            if p[j] != '*':
                break
            dp[0][j + 1] = True
        for i in range(1, ls + 1):
            for j in range(1, lp + 1):
                if p[j - 1] == '*':
                    dp[i][j] = dp[i - 1][j - 1] or dp[i - 1][j] or dp[i][j - 1]
                elif p[j - 1] == '?':
                    dp[i][j] = dp[i - 1][j - 1]
                else:
                    if p[j - 1] == s[i - 1]:
                        dp[i][j] = dp[i - 1][j - 1]
        return dp[ls][lp]

s = Solution()
print(s.isMatch('aa', 'a'))
print(s.isMatch('aa', '*'))
print(s.isMatch('cb', '?a'))
print(s.isMatch('adceb', '*a*b'))
print(s.isMatch('acdcb', 'a*c?b'))

