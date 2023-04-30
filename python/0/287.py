class Solution:
    def findDuplicate(self, nums: [int]) -> int:
        l = len(nums)
        i = 0
        while i < l:
            if nums[i] != i + 1:
                idx = nums[i] - 1
                t = nums[i]
                if nums[idx] == t:
                    return t
                nums[i] = nums[idx]
                nums[idx] = t
            else:
                i += 1
        return -1


s = Solution()
print(s.findDuplicate([1, 3,4,2,2]))
print(s.findDuplicate([3,1,3,4,2]))