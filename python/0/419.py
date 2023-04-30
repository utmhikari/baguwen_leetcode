class Solution:
    def countBattleships(self, board: [[str]]) -> int:
        c = 0
        for row in range(len(board)):
            for col in range(len(board[row])):
                if board[row][col] == 'X':
                    if (row == 0 or board[row - 1][col] != 'X') and (col == 0 or board[row][col - 1] != 'X'):
                        c += 1
        return c

