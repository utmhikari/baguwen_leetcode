from typing import List

class Solution:
    def maxSubArray(self, nums: List[int]) -> int:
        l = len(nums)
        if l == 1:
            return nums[0]
        mx = nums[0]
        sm = nums[0]
        for i in range(1, l):
            n = nums[i]
            if sm < 0:
                sm = 0
            sm += n
            mx = max(sm, mx)
        return mx

s = Solution()
print(s.maxSubArray([-2,1,-3,4,-1,2,1,-5,4]))