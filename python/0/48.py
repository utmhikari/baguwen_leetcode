from typing import List
import pprint


def rot(matrix, left, right):
    steps = right - left
    q = []
    q1 = []
    for i in range(steps):
        q.append((left, left + i, matrix[left][left + i]))
    for _ in range(4):
        for i in range(steps):
            x, y, v = q[i]
            x1, y1 = y, right - x + left
            q1.append((x1, y1, matrix[x1][y1]))
            matrix[x1][y1] = v
        q, q1 = q1, q
        q1 = []


class Solution:
    def rotate(self, matrix: List[List[int]]) -> None:
        """
        Do not return anything, modify matrix in-place instead.
        """
        n = len(matrix)
        if n <= 1:
            return
        left = 0
        right = n - 1
        while left < right:
            rot(matrix, left, right)
            left += 1
            right -= 1
        pprint.pprint(matrix, indent=2, width=50)


s = Solution()
s.rotate([
    [1,2,3],
    [4,5,6],
    [7,8,9]
])
s.rotate([
    [5,1,9,11],
    [2,4,8,10],
    [13,3,6,7],
    [15,14,12,16]
])