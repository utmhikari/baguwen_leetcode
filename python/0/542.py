




class Solution:
    def updateMatrix(self, matrix: [[int]]) -> [[int]]:
        mx = -1
        h = len(matrix)
        w = len(matrix[0])
        ans = [[mx for _ in range(w)] for _ in range(h)]
        i, j = 0, 0
        dots = []
        while i < h:
            j = 0
            while j < w:
                if matrix[i][j] == 0:
                    dots.append((i, j))
                    ans[i][j] = 0
                j += 1
            i += 1
        nxt = 1
        while len(dots) > 0:
            new_dots = []
            for d in dots:
                i, j = d[0], d[1]
                if 0 <= j - 1 and ans[i][j - 1] == -1:
                    ans[i][j - 1] = nxt
                    new_dots.append((i, j - 1))
                if w > j + 1 and ans[i][j + 1] == -1:
                    ans[i][j + 1] = nxt
                    new_dots.append((i, j + 1))
                if 0 <= i - 1 and ans[i - 1][j] == -1:
                    ans[i - 1][j] = nxt
                    new_dots.append((i - 1, j))
                if h > i + 1 and ans[i + 1][j] == -1:
                    ans[i + 1][j] = nxt
                    new_dots.append((i + 1, j))
            dots = new_dots
            nxt += 1
        return ans

s = Solution()
print(s.updateMatrix([[0, 0, 0], [0, 1, 0], [0, 0, 0]]))
print(s.updateMatrix([[0, 0, 0], [0, 1, 0], [1, 1, 1]]))