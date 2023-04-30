from typing import List


class Solution:
    def fourSum(self, nums: List[int], target: int) -> List[List[int]]:
        ans = []
        l = len(nums)
        if l < 4:
            return ans
        nums.sort()
        for a in range(0, l - 3):
            if a > 0 and nums[a] == nums[a - 1]:
                continue
            for b in range(a + 1, l - 2):
                if b > a + 1 and nums[b] == nums[b - 1]:
                    continue
                c, d = b + 1, l - 1
                while c < d:
                    sm = nums[a] + nums[b] + nums[c] + nums[d]
                    if sm < target:
                        c += 1
                    elif sm > target:
                        d -= 1
                    else:
                        ans.append([nums[a], nums[b], nums[c], nums[d]])
                        while c < d and nums[c + 1] == nums[c]:
                            c += 1
                        while c < d and nums[d - 1] == nums[d]:
                            d -= 1
                        c += 1
                        d -= 1
        return ans

