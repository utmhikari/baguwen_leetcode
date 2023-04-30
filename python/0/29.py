class Solution:
    def divide(self, dividend: int, divisor: int) -> int:
        if divisor == 0 or (divisor == -1 and dividend == -2147483648):
            return 2147483647
        if dividend == 0:
            return 0
        unsigned = (dividend > 0 and divisor > 0) or (dividend < 0 and divisor < 0)
        advd, advs = abs(dividend), abs(divisor)
        if advs == 1:
            return advd if unsigned else -advd
        if advd < advs:
            return 0
        mi, t = 1, advs
        mis, ts = [], []
        while True:
            ts.append(t)
            mis.append(mi)
            tmp = (t << 1)
            if tmp > advd:
                break
            else:
                t = tmp
                mi <<= 1
        mis.pop()
        ans = mi
        left = advd - ts.pop()
        while len(mis) > 0 and left > 0:
            tmpmi = mis.pop()
            tmpt = ts.pop()
            if left >= tmpt:
                left -= tmpt
                ans += tmpmi
        return ans if unsigned else -ans


s = Solution()
print(s.divide(10, 3))
print(s.divide(0, -1))
print(s.divide(-2147483648, -1))
print(s.divide(7, -3))
print(s.divide(-7, 3))
print(s.divide(2147483647, 2))
print(s.divide(-2147483648, 2))
