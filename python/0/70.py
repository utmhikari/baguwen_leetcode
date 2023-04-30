
d = {
    1: 1,
    2: 2
}



class Solution:
    def climbStairs(self, n: int) -> int:
        if n in d.keys():
            return d[n]
        c = self.climbStairs(n - 1) + self.climbStairs(n - 2)
        d[n] = c
        return c
