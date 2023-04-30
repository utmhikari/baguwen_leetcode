class Solution:
    def gameOfLife(self, board: [[int]]) -> None:
        """
        Do not return anything, modify board in-place instead.
        """
        to_dead = 2
        to_live = -1
        h = len(board)
        w = len(board[0])
        for i in range(h):
            for j in range(w):
                c = 0
                for x in range(i - 1, i + 2):
                    for y in range(j - 1, j + 2):
                        if 0 <= x < h and 0 <= y < w and not (x == i and y == j):
                            if board[x][y] > 0:
                                c += 1
                if board[i][j] > 0:
                    if c < 2 or c > 3:
                        board[i][j] = to_dead
                else:
                    if c == 3:
                        board[i][j] = to_live
        for i in range(h):
            for j in range(w):
                if board[i][j] == to_live:
                    board[i][j] = 1
                elif board[i][j] == to_dead:
                    board[i][j] = 0


s = Solution()
board = [
    [0,1,0],
    [0,0,1],
    [1,1,1],
    [0,0,0]
]
s.gameOfLife(board)
print(board)