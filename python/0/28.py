import json

class Solution:
    def strStr(self, haystack: str, needle: str) -> int:
        l1, l2 = len(haystack), len(needle)
        if l2 > l1:
            return -1
        if l2 == 0:
            return 0
        if l1 == 0:
            return -1
        mx = l1 - l2
        p = 0
        while p <= mx:
            if haystack[p] == needle[0]:
                i = 1
                while True:
                    if i == l2:
                        return p
                    if haystack[p + i] != needle[i]:
                        break
                    i += 1
            p += 1
        return -1

    def strstrkmp(self, haystack: str, needle: str) -> int:
        m, n = len(haystack), len(needle)
        if n == 0:
            return 0
        if m < n:
            return -1
        # construct fsm for needle
        dp = []
        for i in range(n):
            dp.append([0] * 256)
        x = 0
        dp[0][ord(needle[0])] = 1
        for i in range(1, n):
            for j in range(256):
                dp[i][j] = dp[x][j]
            dp[i][ord(needle[i])] = i + 1
            x = dp[x][ord(needle[i])]
        s = 0
        for i in range(m):
            s = dp[s][ord(haystack[i])]
            if s == n:
                return i - n + 1
        return -1




s = Solution()
# print(s.strStr('hello', 'll'))
# print(s.strStr('aaaaa', 'aba'))
print(s.strstrkmp('hello', 'll'))
print(s.strstrkmp('aaaaa', 'aba'))