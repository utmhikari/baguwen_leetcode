from typing import List
import pprint


def solve(board: List[List[str]], x: int, y: int) -> bool:
    if board[x][y] != '.':
        if x == 8 and y == 8:
            return True
        elif y == 8:
            return solve(board, x + 1, 0)
        else:
            return solve(board, x, y + 1)
    for i in range(1, 10):
        si = str(i)
        is_corrupt = False
        # x/y axis
        for a in range(9):
            if board[x][a] == si:
                is_corrupt = True
                break
        if is_corrupt:
            continue
        for b in range(9):
            if board[b][y] == si:
                is_corrupt = True
                break
        if is_corrupt:
            continue
        # 3x3 grid
        sx, sy = (x // 3) * 3, (y // 3) * 3
        for j in range(3):
            if is_corrupt:
                break
            for k in range(3):
                if board[sx + j][sy + k] == si:
                    is_corrupt = True
                    break
        if is_corrupt:
            continue
        board[x][y] = si
        if x == 8 and y == 8:
            return True
        elif y == 8:
            if solve(board, x + 1, 0):
                return True
        else:
            if solve(board, x, y + 1):
                return True
        board[x][y] = '.'


class Solution:
    def solveSudoku(self, board: List[List[str]]) -> None:
        """
        Do not return anything, modify board in-place instead.
        """
        print(pprint.pformat(board, indent=2, width=55))
        print(solve(board, 0, 0))
        print(pprint.pformat(board, indent=2, width=55))


s = Solution()
s.solveSudoku([["5","3",".",".","7",".",".",".","."],["6",".",".","1","9","5",".",".","."],[".","9","8",".",".",".",".","6","."],["8",".",".",".","6",".",".",".","3"],["4",".",".","8",".","3",".",".","1"],["7",".",".",".","2",".",".",".","6"],[".","6",".",".",".",".","2","8","."],[".",".",".","4","1","9",".",".","5"],[".",".",".",".","8",".",".","7","9"]])
