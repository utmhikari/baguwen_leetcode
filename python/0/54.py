# 给定一个包含 m x n 个元素的矩阵（m 行, n 列），请按照顺时针螺旋顺序，返回矩阵中的所有元素。
#
# 示例 1:
#
# 输入:
# [
#     [ 1, 2, 3 ],
#     [ 4, 5, 6 ],
#     [ 7, 8, 9 ]
# ]
# 输出: [1,2,3,6,9,8,7,4,5]
# 示例 2:
#
# 输入:
# [
#     [1, 2, 3, 4],
#     [5, 6, 7, 8],
#     [9,10,11,12]
# ]
# 输出: [1,2,3,4,8,12,11,10,9,5,6,7]
#
from typing import List


class Solution:
    def __init__(self):
        self.__orders = ['r', 'd', 'l', 'u']
        self.__cur_order_idx = 0

    def insertOrder(self, matrix: List[List[int]], order: List[int], x: int, y: int, mx: int, my: int):
        if x + y > mx + my:
            return False
        self.__cur_order_idx = 0
        is_first_visited = False
        cur_x, cur_y = x, y
        while x <= mx and y <= my:
            # first one
            if cur_x == x and cur_y == y:
                if is_first_visited:
                    return True
                is_first_visited = True
            order.append(matrix[cur_x][cur_y])
            last_fail = False
            while self.__cur_order_idx < len(self.__orders):
                cur_order = self.__orders[self.__cur_order_idx]
                if cur_order == 'r':
                    if cur_y + 1 <= my:
                        cur_y += 1
                        break
                elif cur_order == 'd':
                    if cur_x + 1 <= mx:
                        cur_x += 1
                        break
                    elif last_fail:
                        return True
                elif cur_order == 'l':
                    if cur_y - 1 >= y:
                        cur_y -= 1
                        break
                    elif last_fail:
                        return True
                elif cur_order == 'u':
                    if cur_x - 1 >= x:
                        cur_x -= 1
                        break
                    elif last_fail:
                        return True
                last_fail = True
                self.__cur_order_idx += 1
        return False


    def spiralOrder(self, matrix: List[List[int]]) -> List[int]:
        order: List[int] = []
        if len(matrix) == 0 or len(matrix[0]) == 0:
            return order
        x, y, mx, my = 0, 0, len(matrix) - 1, len(matrix[0]) - 1
        while 2 * (x + y) <= mx + my:
            ok = self.insertOrder(matrix, order, x, y, mx - x, my - y)
            if not ok:
                break
            x += 1
            y += 1
        return order




s = Solution()

m1 = [
    [ 1, 2, 3 ],
    [ 4, 5, 6 ],
    [ 7, 8, 9 ]
]
print(s.spiralOrder(m1))

m2 = [
    [1, 2, 3, 4],
    [5, 6, 7, 8],
    [9,10,11,12]
]
print(s.spiralOrder(m2))


