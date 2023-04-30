from queue import PriorityQueue as PQ


class Solution:
    def kthSmallest(self, matrix: [[int]], k: int) -> int:
        if k == 1:
            return matrix[0][0]
        x = len(matrix)
        y = len(matrix[0])
        visited = []
        for _ in range(x):
            visited.append([False] * y)
        nexts = PQ()
        nexts.put((matrix[1][0], 1, 0))
        nexts.put((matrix[0][1], 0, 1))
        visited[0][0] = True
        visited[1][0] = True
        visited[0][1] = True
        c = 1
        while not nexts.empty():
            nxt = nexts.get()
            c += 1
            if c == k:
                return nxt[0]
            i, j = nxt[1], nxt[2]
            if i < x - 1 and not visited[i + 1][j]:
                visited[i + 1][j] = True
                nexts.put((matrix[i + 1][j], i + 1, j))
            if j < y - 1 and not visited[i][j + 1]:
                visited[i][j + 1] = True
                nexts.put((matrix[i][j + 1], i, j + 1))
        return -1



s = Solution()
print(s.kthSmallest([
    [ 1,  5,  9],
    [10, 11, 13],
    [12, 13, 15]
], 8))