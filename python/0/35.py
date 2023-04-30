from typing import List

class Solution:
    def searchInsert(self, nums: List[int], target: int) -> int:
        l = len(nums)
        if l == 0:
            return 0
        for i in range(l):
            if nums[i] >= target:
                return i
        return l


s = Solution()
print(s.searchInsert([1,3,5,6], 5))
print(s.searchInsert([1,3,5,6], 2))
print(s.searchInsert([1,3,5,6],7))
print(s.searchInsert([1,3,5,6], 0))