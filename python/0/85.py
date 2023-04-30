from typing import List


class Solution:
    def maximalRectangle(self, matrix: List[List[str]]) -> int:
        h = len(matrix)
        if h == 0:
            return 0
        w = len(matrix[0])
        if w == 0:
            return 0
        mws = []
        for _ in range(h):
            mws.append([0] * w)
        mx = 0
        for i in range(h):
            for j in range(w):
                if j == 0:
                    mws[i][j] = int(matrix[i][j])
                elif matrix[i][j] == '1':
                    mws[i][j] = mws[i][j - 1] + 1
                wid = mws[i][j]
                if wid != 0:
                    for k in range(i, -1, -1):
                        wid = min(wid, mws[k][j])
                        if wid == 0:
                            break
                        mx = max(mx, wid * (i - k + 1))
        return mx


s = Solution()
print(s.maximalRectangle([
    ["1","0","1","0","0"],
    ["1","0","1","1","1"],
    ["1","1","1","1","1"],
    ["1","0","0","1","0"]
]))
