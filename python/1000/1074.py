from typing import List


def encode(x, y, w):
    return y * w + x


def decode(p, w):
    y = p // w
    x = p % w
    return x, y


class Solution:
    def numSubmatrixSumTarget(self, matrix: List[List[int]], target: int) -> int:
        h = len(matrix)
        if h == 0:
            return 0
        w = len(matrix[0])
        if w == 0:
            return 0
        prefixes = []
        for _ in range(h + 1):
            prefixes.append([0] * (w + 1))
        for i in range(h):
            for j in range(w):
                prefixes[i + 1][j + 1] = matrix[i][j] + prefixes[i + 1][j] + prefixes[i][j + 1] - prefixes[i][j]
        a = 0
        for i in range(h):
            for j in range(i, h):
                d = {0: 1}
                for k in range(w):
                    s = prefixes[j + 1][k + 1] - prefixes[i][k + 1]
                    if s - target in d.keys():
                        a += d[s - target]
                    if s not in d.keys():
                        d[s] = 0
                    d[s] += 1
        return a


s = Solution()
print(s.numSubmatrixSumTarget([[0,1,0],[1,1,1],[0,1,0]], 0))
print(s.numSubmatrixSumTarget([[1,-1],[-1,1]], 0))