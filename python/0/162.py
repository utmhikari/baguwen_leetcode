


def f(nums, p, s):
    if p in s:
        return -1
    s.add(p)
    l = len(nums)
    if p > l - 1 or p < 0:
        return -1
    if (p == 0 or nums[p] > nums[p - 1]) and (p == l - 1 or nums[p] > nums[p + 1]):
        return p
    right = -1
    if p < l - 1 and nums[p] < nums[p + 1]:
        right = f(nums, p + 1, s)
    elif p < l - 2:
        right = f(nums, p + 2, s)
    if right != -1:
        return right
    left = -1
    if p > 0 and nums[p] < nums[p - 1]:
        left = f(nums, p - 1, s)
    elif p > 1:
        left = f(nums, p - 2, s)
    return left



class Solution:
    def findPeakElement(self, nums: [int]) -> int:
        s = set()
        l = len(nums)
        if l == 0:
            return -1
        p = (l - 1) // 2
        return f(nums, p, s)


s=  Solution()
print(s.findPeakElement([1,2,3,1]))
print(s.findPeakElement([1,2,1,3,5,6,4]))
print(s.findPeakElement([1]))
print(s.findPeakElement([4, 3, 2, 1]))