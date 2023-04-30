class Solution:
    def maxProfit(self, ps) -> int:
        l = len(ps)
        if l <= 1:
            return 0
        t = ps[0]
        m = 0
        i = 1
        while i < l:
            while i < l and ps[i] >= ps[i - 1]:
                i += 1
            if i >= l:
                m += ps[l - 1] - t
                break
            else:
                m += ps[i - 1] - t
                t = ps[i]
                i += 1
        return m

s = Solution()
print(s.maxProfit([7,1,5,3,6,4]))
print(s.maxProfit([1,2,3,4,5]))