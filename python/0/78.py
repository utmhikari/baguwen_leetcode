from typing import List


def sb(ans, tmp, nums, left):
    ans.append(list(tmp))
    for i in range(left, len(nums)):
        tmp.append(nums[i])
        sb(ans, tmp, nums, i + 1)
        tmp.pop()



class Solution:
    def subsets(self, nums: List[int]) -> List[List[int]]:
        ans = []
        tmp = []
        sb(ans, tmp, nums, 0)
        return ans



s = Solution()
print(s.subsets([1,2,3]))