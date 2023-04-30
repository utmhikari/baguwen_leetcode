class Solution:
    def rob(self, nums: [int]) -> int:
        l = len(nums)
        if l == 0:
            return 0
        if l == 1:
            return nums[0]
        if l == 2:
            return max(nums[0], nums[1])
        ns = [0] * l
        ns[0] = nums[0]
        ns[1] = max(nums[0], nums[1])
        for i in range(2, l):
            ns[i] = max(ns[i - 1], ns[i - 2] + nums[i])
        return ns[l - 1]
