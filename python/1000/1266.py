

def solve(p1, p2):
    dx = abs(p1[0] - p2[0])
    dy = abs(p1[1] - p2[1])
    return max(dx, dy)

class Solution:
    def minTimeToVisitAllPoints(self, points) -> int:
        n = len(points)
        if n < 2:
            return 0
        d = 0
        for i in range(n - 1):
            d += solve(points[i], points[i + 1])
        return d