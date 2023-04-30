import math

class Solution:
    def countPrimes(self, n: int) -> int:
        if n <= 2:
            return 0
        i = 2
        j = i
        s = set()
        c = 0
        while True:
            while j in s:
                j += 1
                if j >= n:
                    return c
            c += 1
            k = j
            while k < n:
                k += j
                s.add(k)
            j += 1
            if j >= n:
                return c


s = Solution()
print(s.countPrimes(10))

print(s.countPrimes(3))
