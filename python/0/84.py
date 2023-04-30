from typing import List

class Solution:
    def largestRectangleArea(self, heights: List[int]) -> int:
        l = len(heights)
        if l == 0:
            return 0
        if l == 1:
            return heights[0]
        rights = [0] * l
        lefts = [0] * l
        stk = list()
        # right
        stk.append(0)
        for i in range(1, l):
            while len(stk) > 0 and heights[i] < heights[stk[-1]]:
                previdx = stk.pop()
                rights[previdx] = i - previdx - 1
            stk.append(i)
        while len(stk) > 0:
            previdx = stk.pop()
            rights[previdx] = l - previdx - 1
        # left
        stk.append(l - 1)
        for j in range(1, l):
            i = l - 1 - j
            while len(stk) > 0 and heights[i] < heights[stk[-1]]:
                previdx = stk.pop()
                lefts[previdx] = previdx - i - 1
            stk.append(i)
        while len(stk) > 0:
            previdx = stk.pop()
            lefts[previdx] = previdx
        print(rights)
        print(lefts)
        mx = 0
        for i in range(l):
            sz = (rights[i] + lefts[i] + 1) * heights[i]
            if sz > mx:
                mx = sz
        return mx


s = Solution()
print(s.largestRectangleArea([2,1,5,6,2,3]))
print(s.largestRectangleArea([1,1]))