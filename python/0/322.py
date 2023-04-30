class Solution:
    def coinChange(self, coins: [int], amount: int) -> int:
        inf = 9999999999
        dp = [inf] * (amount + 1)
        dp[0] = 0
        for c in coins:
            for x in range(c, amount + 1):
                dp[x] = min(dp[x], dp[x-c] + 1)
        return dp[amount] if dp[amount] != inf else -1


s = Solution()

print(s.coinChange([1, 2, 5], amount=11))
print(s.coinChange([2], 3))
print(s.coinChange([1, 3, 5], 2))
print(s.coinChange([186,419,83,408], 6249))

print('-----------------\n\n')
print(s.coinChange([39,188,145,223,415,388], 1837))
