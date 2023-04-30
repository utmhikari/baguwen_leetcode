

def f(nums, left, right, target):
    if left > right:
        return -1
    mid = left + (right - left) // 2
    if nums[mid] == target:
        return mid
    if nums[mid] < target:
        return f(nums, mid + 1, right, target)
    return f(nums, left, mid - 1, target)


class Solution:
    def searchRange(self, nums: [int], target: int) -> [int]:
        l = len(nums)
        if l == 0:
            return [-1, -1]
        idx = f(nums, 0, l - 1, target)
        if idx == -1:
            return [-1, -1]
        left, right = idx, idx
        while True:
            if left == 0 or nums[left - 1] != target:
                break
            left -= 1
        while True:
            if right == l - 1 or nums[right + 1] != target:
                break
            right += 1
        return [left, right]


s = Solution()
p