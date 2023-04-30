# 给定 n 个非负整数 a1，a2，...，an，每个数代表坐标中的一个点 (i, ai) 。
# 在坐标内画 n 条垂直线，垂直线 i 的两个端点分别为 (i, ai) 和 (i, 0)。
# 找出其中的两条线，使得它们与 x 轴共同构成的容器可以容纳最多的水。


class Solution:
    def maxArea(self, h) -> int:
        lh = len(h)
        if lh < 2:
            return 0
        l = 0
        r = lh - 1
        ml = False
        if h[l] <= h[r]:
            ml = True
        ans = min(h[l], h[r]) * (r - l)
        while l < r:
            if ml:
                tl = l
                l += 1
                while l < r and h[l] <= h[tl]:
                    l += 1
            else:
                tr = r
                r -= 1
                while l < r and h[r] <= h[tr]:
                    r -= 1
            if l == r:
                break
            ans = max(ans, min(h[l], h[r]) * (r - l))
            if h[l] <= h[r]:
                ml = True
            else:
                ml = False
        return ans

s = Solution()

hs = [1,8,6,2,5,4,8,3,7]
print(s.maxArea(hs))