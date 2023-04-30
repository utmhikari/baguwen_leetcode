# 给定两个字符串 text1 和 text2，返回这两个字符串的最长公共子序列。
#
# 一个字符串的 子序列 是指这样一个新的字符串：它是由原字符串在不改变字符的相对顺序的情况下删除某些字符（也可以不删除任何字符）后组成的新字符串。
# 例如，"ace" 是 "abcde" 的子序列，但 "aec" 不是 "abcde" 的子序列。两个字符串的「公共子序列」是这两个字符串所共同拥有的子序列。
#
# 若这两个字符串没有公共子序列，则返回 0。
import pprint


class Solution:
    def longestCommonSubsequence(self, t1: str, t2: str) -> int:
        l1 = len(t1)
        l2 = len(t2)
        dp = []
        for _ in range(l1):
            dp.append([])
            for _ in range(l2):
                dp[-1].append(0)
        k = 0
        while k < l1:
            if t1[k] == t2[0]:
                while k < l1:
                    dp[k][0] = 1
                    k += 1
            else:
                k += 1
        k = 0
        while k < l2:
            if t1[0] == t2[k]:
                while k < l2:
                    dp[0][k] = 1
                    k += 1
            else:
                k += 1
        i = 1
        while i < l1:
            j = 1
            while j < l2:
                if t1[i] == t2[j]:
                    dp[i][j] = max(dp[i - 1][j - 1] + 1, max(dp[i - 1][j], dp[i][j - 1]))
                else:
                    dp[i][j] = max(dp[i - 1][j], dp[i][j - 1])
                j += 1
            i += 1
        pprint.pprint(dp)
        return dp[l1 - 1][l2 - 1]

s = Solution()
text1 = "ace"
text2 = "abcde"
print(s.longestCommonSubsequence(text1, text2))

text1 = "abc"
text2 = "def"
print(s.longestCommonSubsequence(text1, text2))

text1 = "bsbininm"
text2 = "jmjkbkjkv"
print(s.longestCommonSubsequence(text1, text2))

t1 = 'ABCABBA'
t2 = 'CBABAC'
print(s.longestCommonSubsequence(t1, t2))


