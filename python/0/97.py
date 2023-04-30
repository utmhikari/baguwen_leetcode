"""
97
给定三个字符串 s1, s2, s3, 验证 s3 是否是由 s1 和 s2 交错组成的。

示例 1:

输入: s1 = "aabcc", s2 = "dbbca", s3 = "aadbbcbcac"
输出: true
示例 2:

输入: s1 = "aabcc", s2 = "dbbca", s3 = "aadbbbaccc"
输出: false

"""


def solve(s1, s2, s3, p1, p2, p3):
    if p1 == len(s1) and p2 == len(s2):
        return True
    if p1 != len(s1) and s1[p1] == s3[p3]:
        if solve(s1, s2, s3, p1 + 1, p2, p3 + 1):
            return True
    if p2 != len(s2) and s2[p2] == s3[p3]:
        if solve(s1, s2, s3, p1, p2 + 1, p3 + 1):
            return True
    return False


def check(s1, s2, s3):
    l1, l2, l3 = len(s1), len(s2), len(s3)
    if l1 + l2 != l3:
        return False
    if l1 == 0:
        return s2 == s3
    if l2 == 0:
        return s1 == s3
    cs1, cs2 = [0] * 128, [0] * 128
    for c in s1:
        cs1[ord(c)] += 1
    for c in s2:
        cs1[ord(c)] += 1
    for c in s3:
        cs2[ord(c)] += 1
    for i in range(128):
        if cs1[i] != cs2[i]:
            return False
    return True


class Solution:
    def isInterleave(self, s1: str, s2: str, s3: str) -> bool:
        if not check(s1, s2, s3):
            return False
        l1, l2, l3 = len(s1), len(s2), len(s3)
        dp = []
        for _ in range(l1 + 1):
            dp.append([])
            for _ in range(l2 + 1):
                dp[-1].append(False)
        dp[0][0] = True
        for i in range(l1 + 1):
            for j in range(l2 + 1):
                if i == 0 and j == 0:
                    continue
                if i >= 1 and s3[i + j - 1] == s1[i - 1] and dp[i - 1][j]:
                    dp[i][j] = True
                elif j >= 1 and s3[i + j - 1] == s2[j - 1] and dp[i][j - 1]:
                    dp[i][j] = True
        return dp[l1][l2]



if __name__ == '__main__':
    s = Solution()
    print(s.isInterleave("aabcc", "dbbca", "aadbbcbcac"))
    print(s.isInterleave("aabcc", "dbbca", "aadbbbaccc"))

