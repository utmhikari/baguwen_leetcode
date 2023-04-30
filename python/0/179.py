from typing import List, Any
import functools


def rawcmp(x: Any, y: Any):
    return (x > y) - (x < y)


def cmpf(s1: str, s2: str):
    l1, l2 = len(s1), len(s2)
    if l1 == l2:
        return rawcmp(s1, s2)
    if l1 > l2:
        rawret = rawcmp(s1[0:l2], s2)
        return rawret if rawret != 0 else cmpf(s1[l2:], s2)
    else:
        rawret = rawcmp(s1, s2[0:l1])
        return rawret if rawret != 0 else cmpf(s1, s2[l1:])


class Solution:
    def largestNumber(self, nums: List[int]) -> str:
        newnums = sorted([str(n) for n in nums],
                         key=functools.cmp_to_key(cmpf),
                         reverse=True)
        ans = ''.join(newnums)
        return '0' if ans[0] == '0' else ans


s = Solution()
print(s.largestNumber([10,2]))
print(s.largestNumber([3,30,34,5,9]))
print(s.largestNumber([0,0,0,0,0000]))