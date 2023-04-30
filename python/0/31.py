from typing import List




class Solution:
    def nextPermutation(self, nums: List[int]) -> None:
        """
        Do not return anything, modify nums in-place instead.
        """
        l = len(nums)
        if l <= 1:
            return
        dec = -1
        for i in range(l - 1):
            if nums[l - i - 2] < nums[l - i - 1]:
                dec = l - i - 2
                break
        if dec == -1:
            nums.reverse()
            return
        large = l - 1
        for i in range(dec + 1, l):
            if nums[i] <= nums[dec]:
                large = i - 1
                break
        nums[dec], nums[large] = nums[large], nums[dec]
        left = dec + 1
        right = l - 1
        while left < right:
            nums[left], nums[right] = nums[right], nums[left]
            left += 1
            right -= 1


s = Solution()
lsts = [
    [2, 3, 1],
    [1, 2, 3],
    [3, 2, 1],
    [1, 1, 5],
    [1, 5, 1],
    [1, 2, 4, 3],
    [1, 2, 4, 3, 5]
]
for lst in lsts:
    s.nextPermutation(lst)
    print(lst)