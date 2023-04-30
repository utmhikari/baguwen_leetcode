from typing import List


class Solution:
    def jump(self, nums: List[int]) -> int:
        l = len(nums)
        if l == 1:
            return 0
        start = 0
        step = 0
        maxpos = 0
        for i in range(l - 1):
            maxpos = max(maxpos, i + nums[i])
            if i == start:
                start = maxpos
                step += 1
        return step


s = Solution()
print(s.jump([2,3,1,1,4]))
print(s.jump([0]))