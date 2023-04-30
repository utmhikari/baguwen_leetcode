from random import randint



class Solution:
    def findKthLargest(self, nums: [int], k: int) -> int:
        l = len(nums)
        if k > l:
            return -1
        p = randint(0, l - 1)
        pivot = nums[p]
        left = [i for i in nums[:p] + nums[p + 1:] if i < pivot]
        right = [i for i in nums[:p] + nums[p + 1:] if i > pivot]
        if len(right) == k - 1:
            return pivot
        if len(right) > k - 1:
            return self.findKthLargest(right, k)
        return self.findKthLargest(left, k - len(right) - 1)
