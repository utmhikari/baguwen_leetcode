
def cj(nums, pos, bs):
    if bs[pos] == 1:
        return True
    elif bs[pos] == -1:
        return False
    p = pos - 1
    while p >= 0:
        steps = nums[p]
        if steps + p >= pos:
            if cj(nums, p, bs):
                return True
        p -= 1
    bs[pos] = -1
    return False


class Solution:
    def canJump(self, nums: [int]) -> bool:
        bs = [0] * len(nums)
        bs[0] = 1
        l = len(nums)
        return cj(nums, l - 1, bs)



s = Solution()
print(s.canJump([2,3,1,1,4]))
print(s.canJump([3,2,1,0,4]))