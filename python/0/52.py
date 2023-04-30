from typing import List


def genans(grid):
    a = []
    for i in range(len(grid)):
        a.append(''.join(grid[i]))
    return a


def solve(grid, x, n, ans):
    for y in range(n):
        canput = True
        # col
        for i in range(x):
            if grid[i][y] == 'Q':
                canput = False
                break
        if not canput:
            continue
        # /
        for i in range(min(n - 1 - y, x)):
            x1, y1 = x - 1 - i, y + 1 + i
            if grid[x1][y1] == 'Q':
                canput = False
                break
        if not canput:
            continue
        # \
        for i in range(min(x, y)):
            x1, y1 = x - 1 - i, y - 1 - i
            if grid[x1][y1] == 'Q':
                canput = False
                break
        if not canput:
            continue
        # put it
        grid[x][y] = 'Q'
        if x == n - 1:
            ans.append(genans(grid))
        else:
            solve(grid, x + 1, n, ans)
        grid[x][y] = '.'



class Solution:
    def totalNQueens(self, n: int) -> int:
        ans = []
        grid = []
        for _ in range(n):
            grid.append(['.'] * n)
        solve(grid, 0, n, ans)
        return len(ans)