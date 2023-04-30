from typing import List


class Solution:
    def removeElement(self, nums: List[int], val: int) -> int:
        l = len(nums)
        indices = []
        for i in range(l):
            if nums[i] != val:
                indices.append(i)
        a = len(indices)
        if a == l:
            return l
        for i in range(a):
            nums[i], nums[indices[i]] = nums[indices[i]], nums[i]
        # print(nums)
        return a

s = Solution()
print(s.removeElement([3,2,2,3], 3))
print(s.removeElement([0,1,2,2,3,0,4,2], 2))