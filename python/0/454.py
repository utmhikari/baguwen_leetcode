class Solution:
    def fourSumCount(self, A, B, C, D) -> int:
        ab = dict()
        for i in A:
            for j in B:
                if i + j not in ab.keys():
                    ab[i + j] = 1
                else:
                    ab[i + j] += 1
        ans = 0
        for i in C:
            for j in D:
                if -(i + j) in ab.keys():
                    ans += ab[-(i + j)]
        return ans

s = Solution()
print(s.fourSumCount([ 1, 2], [-2,-1], [-1, 2], [ 0, 2]))