# 给定一个数组，它的第 i 个元素是一支给定股票第 i 天的价格。
#
# 如果你最多只允许完成一笔交易（即买入和卖出一支股票），设计一个算法来计算你所能获取的最大利润。
#
# 注意你不能在买入股票前卖出股票。

class Solution:
        def maxProfit(self, ps) -> int:
            l = len(ps)
            if l <= 1:
                return 0
            t = ps[0]
            m = 0
            i = 1
            while i < l:
                while i < l and ps[i] >= t:
                    m = max(m, ps[i] - t)
                    i += 1
                if i >= l:
                    break
                else:
                    t = ps[i]
                    i += 1
            return m


s = Solution()
ps = [7,1,5,3,6,4]
print(s.maxProfit(ps))

ps = [7,6,4,3,1]
print(s.maxProfit(ps))