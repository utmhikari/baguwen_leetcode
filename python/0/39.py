from typing import List


def solve(candidates, ans, buffer, left, right, target):
    if target == 0:
        ans.append(list(buffer))
        return True
    for i in range(left, right + 1):
        if candidates[i] <= target:
            buffer.append(candidates[i])
            solve(candidates, ans, buffer, i, right, target - candidates[i])
            buffer.pop()


class Solution:
    def combinationSum(self, candidates: List[int], target: int) -> List[List[int]]:
        candidates.sort(reverse=True)
        ans = []
        buffer = []
        l = len(candidates)
        if l == 0:
            return ans
        solve(candidates, ans, buffer, 0, l - 1, target)
        return ans


s = Solution()
print(s.combinationSum([2,3,6,7], 7))
print(s.combinationSum([2,3,5], 8))
