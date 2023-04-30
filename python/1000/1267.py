

class Solution:
    def countServers(self, grid):
        m = len(grid)
        bools = []
        if m == 0:
            return 0
        n = len(grid[0])
        if n == 0:
            return 0
        for i in range(m):
            bools.append([False] * n)
        for i in range(m):
            col = -1
            for j in range(n):
                if grid[i][j] == 1:
                    if col != -1:
                        bools[i][col] = True
                        bools[i][j] = True
                    else:
                        col = j
        for j in range(n):
            row = -1
            for i in range(m):
                if grid[i][j] == 1:
                    if row != -1:
                        bools[row][j] = True
                        bools[i][j] = True
                    else:
                        row = i
        cnt = 0
        for i in range(m):
            for j in range(n):
                if bools[i][j]:
                    cnt += 1
        return cnt

s = Solution()

print(s.countServers([[1,0],[1,1]]))
print(s.countServers([[1,1,0,0],[0,0,1,0],[0,0,1,0],[0,0,0,1]]))