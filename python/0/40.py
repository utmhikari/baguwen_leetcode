from typing import List


def solve(candidates, ans, buffer, left, right, target):
    if left > right:
        return False
    i = left
    while i <= right:
        if candidates[i] == target:
            buffer.append(candidates[i])
            ans.append(list(buffer))
            buffer.pop()
        elif candidates[i] < target:
            buffer.append(candidates[i])
            solve(candidates, ans, buffer, i + 1, right, target - candidates[i])
            buffer.pop()
        j = i + 1
        while j <= right and candidates[j] == candidates[j - 1]:
            j += 1
        if j > right:
            break
        # print(candidates, candidates[j], candidates[j - 1])
        i = j


class Solution:
    def combinationSum2(self, candidates: List[int], target: int) -> List[List[int]]:
        candidates.sort(reverse=True)
        ans = []
        buffer = []
        l = len(candidates)
        if l == 0:
            return ans
        solve(candidates, ans, buffer, 0, l - 1, target)
        return ans


s = Solution()
print(s.combinationSum2([10,1,2,7,6,1,5], 8))
print(s.combinationSum2([2,5,2,1,2], 5))
