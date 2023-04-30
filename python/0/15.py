from typing import List


class Solution:
    def threeSum(self, nums: List[int]) -> List[List[int]]:
        l = len(nums)
        if l < 3:
            return []
        nums.sort()
        ans = []
        for i in range(l - 2):
            if i > 0 and nums[i] == nums[i - 1]:
                continue
            j, k = i + 1, l - 1
            while j < k:
                sm = nums[i] + nums[j] + nums[k]
                if sm < 0:
                    j += 1
                elif sm > 0:
                    k -= 1
                else:
                    ans.append([nums[i], nums[j], nums[k]])
                    while j < k and nums[j] == nums[j + 1]:
                        j += 1
                    while j < k and nums[k] == nums[k - 1]:
                        k -= 1
                    j += 1
                    k -= 1
        return ans


s = Solution()

nums = [-1, 0, 1, 2, -1, -4]
print(s.threeSum(nums))

nums = [0, 0, 0]
print(s.threeSum(nums))
