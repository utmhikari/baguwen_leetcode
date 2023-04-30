# 编写一个高效的算法来判断m x n矩阵中，是否存在一个目标值。该矩阵具有如下特性：
#
# 每行中的整数从左到右按升序排列。
# 每行的第一个整数大于前一行的最后一个整数。
# 示例1:
#
# 输入:
# matrix = [
#     [1,   3,  5,  7],
#     [10, 11, 16, 20],
#     [23, 30, 34, 50]
# ]
# target = 3
# 输出: true
# 示例2:
#
# 输入:
# matrix = [
#     [1,   3,  5,  7],
#     [10, 11, 16, 20],
#     [23, 30, 34, 50]
# ]
# target = 13
# 输出: false

from typing import List


class Solution:
    def bs(self, nums: List[int], target: int, left: int, right: int) -> bool:
        if left > right:
            return False
        mid = (left + right) // 2
        if target == nums[mid]:
            return True
        if self.bs(nums, target, left, mid - 1):
            return True
        if self.bs(nums, target, mid + 1, right):
            return True
        return False

    def searchMatrix(self, matrix: List[List[int]], target: int) -> bool:
        if len(matrix) == 0 or len(matrix[0]) == 0:
            return False
        # find x yekeyierfen
        x = -1
        for i in range(len(matrix)):
            if target < matrix[i][0]:
                x = i - 1
                break
            if i == len(matrix) - 1:
                if target <= matrix[i][len(matrix[i]) - 1]:
                    x = i
        if x == -1:
            return False
        return self.bs(matrix[x], target, 0, len(matrix[x]) - 1)


s = Solution()
matrix = [
    [1,   3,  5,  7],
    [10, 11, 16, 20],
    [23, 30, 34, 50]
]
print(s.searchMatrix(matrix, 13))
print(s.searchMatrix(matrix, 31))
print(s.searchMatrix(matrix, -1))
print(s.searchMatrix(matrix, 34))
print(s.searchMatrix(matrix, 50))

