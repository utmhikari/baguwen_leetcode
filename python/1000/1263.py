# 推箱子 是一款风靡全球的益智小游戏，玩家需要将箱子推到仓库中的目标位置。
#
# 游戏地图用大小为 n * m 的网格 grid 表示，其中每个元素可以是墙、地板或者是箱子。
#
# 现在你将作为玩家参与游戏，按规则将箱子 'B' 移动到目标位置 'T' ：
#
# 玩家用字符 'S' 表示，只要他在地板上，就可以在网格中向上、下、左、右四个方向移动。
# 地板用字符 '.' 表示，意味着可以自由行走。
# 墙用字符 '#' 表示，意味着障碍物，不能通行。 
# 箱子仅有一个，用字符 'B' 表示。相应地，网格上有一个目标位置 'T'。
# 玩家需要站在箱子旁边，然后沿着箱子的方向进行移动，此时箱子会被移动到相邻的地板单元格。记作一次「推动」。
# 玩家无法越过箱子。
# 返回将箱子推到目标位置的最小 推动 次数，如果无法做到，请返回 -1。

class Solution(object):
    def minPushBox(self, grid):
        """
        :type grid: List[List[str]]
        :rtype: int
        """
        self.move = [[1,0],[-1,0],[0,1],[0,-1]]
        self.N,self.M =len(grid),len(grid[0])
        for i in range(self.N):
            for j in range(self.M):
                if grid[i][j] == 'B':
                    grid[i][j] = '.'
                    bx,by = i,j
                if grid[i][j] == 'S':
                    grid[i][j] = '.'
                    sx,sy = i,j
        s = set(str([bx,by,sx,sy]))
        t = [[bx,by,sx,sy,0]]
        while t:
            bx,by,sx,sy,step = t.pop(0)
            if grid[bx][by]=='T':
                return step
            grid[bx][by] = 'B'
            for dx,dy in self.move:
                if dx+bx>=0 and dx+bx<self.N and dy+by>=0 and dy+by<self.M and grid[dx+bx][dy+by] in ".T":
                    if -dx+bx>=0 and -dx+bx<self.N and -dy+by>=0 and -dy+by<self.M and grid[-dx+bx][-dy+by] in ".T":
                        if self.bfs(sx,sy,-dx+bx,-dy+by,grid):
                            if str([dx+bx,dy+by,bx,by]) not in s:
                                t.append([dx+bx,dy+by,bx,by,step+1])
                                s.add(str([dx+bx,dy+by,bx,by]))
            grid[bx][by] = '.'
        return -1
    def bfs (self,sx,sy,ex,ey,road):
        t = [[sx,sy]]
        s = set(str([sx,sy]))
        while t:
            x,y = t.pop(0)
            if [x,y] == [ex,ey]:
                return True
            for dx,dy in self.move:
                if dx+x>=0 and dx+x<self.N and dy+y>=0 and dy+y<self.M and road[dx+x][dy+y] in ".T":
                    if str([dx+x,dy+y]) not in s:
                        t.append([dx+x,dy+y])
                        s.add(str([dx+x,dy+y]))
        return False



s = Solution()


grid = [["#","#","#","#","#","#"],
        ["#","T","#","#","#","#"],
        ["#",".",".","B",".","#"],
        ["#",".","#","#",".","#"],
        ["#",".",".",".","S","#"],
        ["#","#","#","#","#","#"]]

# 3
print(s.minPushBox(grid))

grid = [["#","#","#","#","#","#"],
        ["#","T","#","#","#","#"],
        ["#",".",".","B",".","#"],
        ["#","#","#","#",".","#"],
        ["#",".",".",".","S","#"],
        ["#","#","#","#","#","#"]]

# -1
print(s.minPushBox(grid))


grid = [["#","#","#","#","#","#"],
        ["#","T",".",".","#","#"],
        ["#",".","#","B",".","#"],
        ["#",".",".",".",".","#"],
        ["#",".",".",".","S","#"],
        ["#","#","#","#","#","#"]]

# 5
print(s.minPushBox(grid))

grid = [["#","#","#","#","#","#","#"],
        ["#","S","#",".","B","T","#"],
        ["#","#","#","#","#","#","#"]]

# -1
print(s.minPushBox(grid))

grid = [["#",".",".","#","#","#","#","#"],["#",".",".","T","#",".",".","#"],["#",".",".",".","#","B",".","#"],["#",".",".",".",".",".",".","#"],["#",".",".",".","#",".","S","#"],["#",".",".","#","#","#","#","#"]]

print(s.minPushBox(grid))

grid = [["#",".",".",".",".",".",".",".",".","."],[".",".",".",".",".","#",".",".",".","#"],["#",".","#",".",".","T",".",".",".","."],[".","#",".",".",".",".",".",".",".","."],[".",".",".",".",".",".","#",".",".","."],[".",".",".","#","#","S",".","B",".","."],["#",".",".",".",".",".",".","#",".","."],[".","#",".",".",".",".",".",".",".","."],[".",".",".",".",".",".",".",".",".","."],[".",".",".",".",".","#",".",".",".","."]]

print(s.minPushBox(grid))




