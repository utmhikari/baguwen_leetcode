class Solution:
    def removeDuplicates(self, nums) -> int:
        l = len(nums)
        if l == 0:
            return 0
        c = nums[0]
        d = 0
        i = 1
        while i < l - d:
            if nums[i] == c:
                d += 1
                nums.pop(i)
            else:
                c = nums[i]
                i += 1
        return i



s = Solution()
nums = [1,1,2]
nl = s.removeDuplicates(nums)
print(nl, nums)

nums = [0,0,1,1,1,2,2,3,3,4]
nl = s.removeDuplicates(nums)
print(nl, nums)



