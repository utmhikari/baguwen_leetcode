from typing import List


def v(ans, tmp, d, ks, idx):
    if idx >= len(ks):
        ans.append(list(tmp))
        return
    left = ks[idx]
    for i in range(d[left] + 1):
        for j in range(i):
            tmp.append(left)
        v(ans, tmp, d, ks, idx + 1)
        for j in range(i):
            tmp.pop()


class Solution:
    def subsetsWithDup(self, nums: List[int]) -> List[List[int]]:
        d = {}
        for n in nums:
            if n not in d.keys():
                d[n] = 0
            d[n] += 1
        tmp = []
        ans = []
        keys = sorted(d.keys())
        v(ans, tmp, d, keys, 0)
        return ans


s = Solution()
print(s.subsetsWithDup([1,2,2]))