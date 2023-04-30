from typing import List


def pm(d, x, l, tmp, ans):
    if x == l:
        ans.append(list(tmp))
        return
    for k in d.keys():
        if d[k] == 0:
            continue
        tmp.append(k)
        d[k] -= 1
        pm(d, x + 1, l, tmp, ans)
        d[k] += 1
        tmp.pop()


class Solution:
    def permuteUnique(self, nums: List[int]) -> List[List[int]]:
        l = len(nums)
        if l <= 1:
            return [nums]
        d = dict()
        for n in nums:
            if not n in d.keys():
                d[n] = 0
            d[n] += 1
        ans = []
        tmp = []
        pm(d, 0, l, tmp, ans)
        return ans



s = Solution()
print(s.permuteUnique([1,1,2]))
print(s.permuteUnique([1,1,2,2]))
print(s.permuteUnique([1,2,3]))