class Solution:
    def firstMissingPositive(self, nums: [int]) -> int:
        i = 0
        l = len(nums)
        while i < l:
            if 0 < nums[i] < l and nums[i] != i + 1:
                t = nums[i]
                idx = nums[i] - 1
                if nums[idx] == nums[i]:
                    i += 1
                else:
                    nums[i] = nums[idx]
                    nums[idx] = t
            else:
                i += 1
        i = 0
        while i < l:
            if nums[i] != i + 1:
                return i + 1
            i += 1
        return l + 1


s = Solution()
print(s.firstMissingPositive([1,2,0]))
print(s.firstMissingPositive([3,4,-1,1]))
print(s.firstMissingPositive([7,8,9,11,12]))
print(s.firstMissingPositive([]))
print(s.firstMissingPositive([1]))
print(s.firstMissingPositive([1,1]))