from typing import List


class Solution:
    def trap(self, height: List[int]) -> int:
        l = len(height)
        if l <= 2:
            return 0
        lefts = [0] * l
        rights = [0] * l
        for i in range(1, l):
            lefts[i] = max(lefts[i - 1], height[i - 1])
            rights[l - i - 1] = max(rights[l - i], height[l - i])
        a = 0
        for i in range(l):
            h = min(lefts[i], rights[i]) - height[i]
            if h > 0:
                a += h
        return a

s = Solution()
print(s.trap([2,0,2]))
print(s.trap([0,1,0,2,1,0,1,3,2,1,2,1]))
