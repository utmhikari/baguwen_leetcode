# 给定一个整数数组 nums 和一个目标值 target，请你在该数组中找出和为目标值的那 两个 整数，并返回他们的数组下标。
#
# 你可以假设每种输入只会对应一个答案。但是，你不能重复利用这个数组中同样的元素。

class Solution:
    def twoSum(self, nums, target: int):
        i = 0
        l = len(nums)
        f = False
        x = -1
        y = -1
        while not f and i < l - 1:
            j = i + 1
            while j < l:
                if nums[i] + nums[j] == target:
                    f = True
                    x = i
                    y = j
                    break
                j += 1
            i += 1
        return [x, y]

s = Solution()

nums = [2, 7, 11, 15]
target = 9

print(s.twoSum(nums, target))