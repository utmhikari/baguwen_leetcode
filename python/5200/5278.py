# 给你一个由小写字母组成的字符串 s，和一个整数 k。
#
# 请你按下面的要求分割字符串：
#
# 首先，你可以将 s 中的部分字符修改为其他的小写英文字母。
# 接着，你需要把 s 分割成 k 个非空且不相交的子串，并且每个子串都是回文串。
# 请返回以这种方式分割字符串所需修改的最少字符数。
import json

ans = dict()
edits = dict()


def edittimes(s):
    if s in edits.keys():
        return edits[s]
    c = 0
    l = len(s)
    if l == 1:
        edits[s] = 0
        return 0
    k = 0
    while True:
        left = k
        right = l - 1 - k
        if left >= right:
            break
        if s[left] != s[right]:
            c += 1
        k += 1
    edits[s] = c
    return c


class Solution:
    def palindromePartition(self, s: str, k: int) -> int:
        if s not in ans.keys():
            ans[s] = dict()
        if k in ans[s].keys():
            return ans[s][k]
        if k == 1:
            ans[s][k] = edittimes(s)
            return ans[s][k]
        l = len(s) - k + 1
        if l < 0:
            return -99999999
        ans[s][k] = 999999999999999
        for i in range(l):
            m = edittimes(s[:i + 1]) + self.palindromePartition(s[i + 1:], k - 1)
            if ans[s][k] > m:
                ans[s][k] = m
        return ans[s][k]





so = Solution()

print(so.palindromePartition("abc", 2))  # 1
print(json.dumps(edits, indent=2))

print(so.palindromePartition("aabbc", 3))  # 0
print(json.dumps(edits, indent=2))

print(so.palindromePartition("leetcode", 8))  # 0
print(json.dumps(edits, indent=2))
