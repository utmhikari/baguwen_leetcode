

def sch(ns, l, r, t):
    if l > r:
        return -1
    if l == r:
        return l if ns[l] == t else -1
    mid = l + (r - l) // 2
    if ns[mid] == t:
        return mid
    if ns[mid] < t:
        return sch(ns, mid + 1, r, t)
    return sch(ns, l, mid - 1, t)


class Solution:
    def search(self, nums, target) -> int:
        l = len(nums)
        if l == 0:
            return -1
        if l == 1:
            return 0 if nums[0] == target else -1
        if nums[0] == target:
            return 0
        if nums[l - 1] == target:
            return l - 1
        if nums[0] < nums[l - 1]:
            return sch(nums, 0, l - 1, target)
        elif target >= nums[0]:
            i = 1
            while nums[i] > nums[i - 1]:
                if nums[i] == target:
                    return i
                i += 1
        else:
            i = l - 2
            while nums[i] < nums[i + 1]:
                if nums[i] == target:
                    return i
                i -= 1
        return -1




s = Solution()
nums = [4,5,6,7,0,1,2]
print(s.search(nums, 0))

nums = [2,4,5,6,7,0,1]
print(s.search(nums, 3))

