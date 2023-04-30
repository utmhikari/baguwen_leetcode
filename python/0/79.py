# 给定一个二维网格和一个单词，找出该单词是否存在于网格中。
#
# 单词必须按照字母顺序，通过相邻的单元格内的字母构成，其中“相邻”单元格是那些水平相邻或垂直相邻的单元格。同一个单元格内的字母不允许被重复使用。
#
#
#
# 示例:
#
# board =
# [
#     ['A','B','C','E'],
#     ['S','F','C','S'],
#     ['A','D','E','E']
# ]
#
# 给定 word = "ABCCED", 返回 true
# 给定 word = "SEE", 返回 true
# 给定 word = "ABCB", 返回 false

from typing import List


class Solution:
    def __init__(self):
        self.__board = []
        self.__visited = []
        self.__word = ''

    def __find(self, w_idx: int, x: int, y: int):
        if self.__visited[x][y]:
            return False
        if w_idx == len(self.__word) - 1:
            return True
        self.__visited[x][y] = True
        if x > 0 and self.__board[x - 1][y] == self.__word[w_idx + 1]:
            if self.__find(w_idx + 1, x - 1, y):
                return True
        if x < len(self.__board) - 1 and self.__board[x + 1][y] == self.__word[w_idx + 1]:
            if self.__find(w_idx + 1, x + 1, y):
                return True
        if y > 0 and self.__board[x][y - 1] == self.__word[w_idx + 1]:
            if self.__find(w_idx + 1, x, y - 1):
                return True
        if y < len(self.__board[0]) - 1 and self.__board[x][y + 1] == self.__word[w_idx + 1]:
            if self.__find(w_idx + 1, x, y + 1):
                return True
        self.__visited[x][y] = False
        return False

    def exist(self, board: List[List[str]], word: str) -> bool:
        if not word:
            return False
        if len(board) == 0 or len(board[0]) == 0:
            return False
        self.__board = board
        self.__visited = []
        self.__word = word
        starts = []
        for i in range(len(self.__board)):
            self.__visited.append([])
            for j in range(len(self.__board[i])):
                self.__visited[-1].append(False)
                if self.__board[i][j] == self.__word[0]:
                    starts.append((i, j))
        print(starts)
        if len(starts) == 0:
            return False
        for start in starts:
            if self.__find(0, start[0], start[1]):
                return True
        return False


s = Solution()
board = [
    ['A','B','C','E'],
    ['S','F','C','S'],
    ['A','D','E','E']
]
print(s.exist(board, "ABCCED"))
print(s.exist(board, "SEE"))
print(s.exist(board, "ABCB"))
