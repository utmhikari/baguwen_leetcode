# 给你一个字符串 s 和一个字符规律 p，请你来实现一个支持 '.' 和 '*' 的正则表达式匹配。
#
# '.' 匹配任意单个字符
# '*' 匹配零个或多个前面的那一个元素
# 所谓匹配，是要涵盖 整个 字符串 s的，而不是部分字符串。
#

# ppppppp
# s
# s
# s


def pt(dp, s, p):
    k = ''
    ls = len(s)
    lp = len(p)
    k += ' \t\"\"\t'
    for c in p:
        k += c
        k += '\t'
    k += '\n'
    for j in range(ls + 1):
        if j == 0:
            k += '\"\"'
        else:
            k += s[j - 1]
        k += '\t'
        for i in range(lp + 1):
            if dp[i][j]:
                k += 'T\t'
            else:
                k += 'F\t'
        k += '\n'
    return k

class Solution:
    def isMatch(self, s: str, p: str) -> bool:
        ls = len(s)
        lp = len(p)
        dp = []
        for _ in range(lp + 1):
            dp.append([])
            for _ in range(ls + 1):
                dp[-1].append(False)
        dp[0][0] = True
        for i in range(lp + 1):
            for j in range(ls + 1):
                if i == 0:
                    dp[i][j] = (j == 0)
                elif p[i - 1] == '*':
                    dp[i][j] = (i >= 2 and dp[i - 2][j]) or dp[i - 1][j] or (j >= 2 and (s[j - 1] == s[j - 2] and dp[i - 1][j - 1]) or (i >= 2 and p[i - 2] == '.' and (dp[i - 1][j - 1] or dp[i][j - 1])))
                elif j != 0 and (p[i - 1] == s[j - 1] or p[i - 1] == '.'):
                    dp[i][j] = dp[i - 1][j - 1] or (i >= 3 and p[i - 2] == '*' and dp[i - 3][j - 1])
        # print(pt(dp, s, p))
        return dp[lp][ls]




so = Solution()

s = "aa"
p = "a"
print(so.isMatch(s, p))

s = "aa"
p = "a*"
print(so.isMatch(s, p))

s = "ab"
p = ".*"
print(so.isMatch(s, p))

s = "aab"
p = "c*a*b"
print(so.isMatch(s, p))

s = "mississippi"
p = "mis*is*p*."
print(so.isMatch(s, p))

s = "aaba"
p = "ab*a*c*a"
print(so.isMatch(s, p))

s = "aaca"
p = "ab*a*c*a"
print(so.isMatch(s, p))

s = "aaa"
p = ".*"
print(so.isMatch(s, p))

s = "a"
p = "ab*"
print(so.isMatch(s, p))

s = "aaa"
p = "ab*a"
print(so.isMatch(s, p))

s = "a"
p = ".*"
print(so.isMatch(s, p))

s = "a"
p = ".*"
print(so.isMatch(s, p))

s = ""
p = ".*"
print(so.isMatch(s, p))
