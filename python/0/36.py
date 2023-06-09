from typing import List


class Solution:
    def isValidSudoku(self, board: List[List[str]]) -> bool:
        for i in range(9):
            for j in range(9):
                if board[i][j] == '.':
                    continue
                for k in range(9):
                    if k != j and board[i][k] == board[i][j]:
                        return False
                    if k != i and board[k][j] == board[i][j]:
                        return False
                sx, sy = (i // 3) * 3, (j // 3) * 3
                for a in range(3):
                    for b in range(3):
                        if not (sx + a == i and sy + b == j) and \
                                board[sx + a][sy + b] == board[i][j]:
                            return False
        return True


s = Solution()
print(s.isValidSudoku([
    ["5", "3", ".", ".", "7", ".", ".", ".", "."],
    ["6", ".", ".", "1", "9", "5", ".", ".", "."],
    [".", "9", "8", ".", ".", ".", ".", "6", "."],
    ["8", ".", ".", ".", "6", ".", ".", ".", "3"],
    ["4", ".", ".", "8", ".", "3", ".", ".", "1"],
    ["7", ".", ".", ".", "2", ".", ".", ".", "6"],
    [".", "6", ".", ".", ".", ".", "2", "8", "."],
    [".", ".", ".", "4", "1", "9", ".", ".", "5"],
    [".", ".", ".", ".", "8", ".", ".", "7", "9"]
]))
print(s.isValidSudoku([
    ["8", "3", ".", ".", "7", ".", ".", ".", "."],
    ["6", ".", ".", "1", "9", "5", ".", ".", "."],
    [".", "9", "8", ".", ".", ".", ".", "6", "."],
    ["8", ".", ".", ".", "6", ".", ".", ".", "3"],
    ["4", ".", ".", "8", ".", "3", ".", ".", "1"],
    ["7", ".", ".", ".", "2", ".", ".", ".", "6"],
    [".", "6", ".", ".", ".", ".", "2", "8", "."],
    [".", ".", ".", "4", "1", "9", ".", ".", "5"],
    [".", ".", ".", ".", "8", ".", ".", "7", "9"]
]))
print(s.isValidSudoku([
    [".", ".", ".", "9", ".", ".", ".", ".", "."],
    [".", ".", ".", ".", ".", ".", ".", ".", "."],
    [".", ".", "3", ".", ".", ".", ".", ".", "1"],
    [".", ".", ".", ".", ".", ".", ".", ".", "."],
    ["1", ".", ".", ".", ".", ".", "3", ".", "."],
    [".", ".", ".", ".", "2", ".", "6", ".", "."],
    [".", "9", ".", ".", ".", ".", ".", "7", "."],
    [".", ".", ".", ".", ".", ".", ".", ".", "."],
    ["8", ".", ".", "8", ".", ".", ".", ".", "."]
]))
