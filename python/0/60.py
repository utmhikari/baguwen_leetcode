


class Solution:
    def getPermutation(self, n: int, k: int) -> str:
        facts = [1]
        for i in range(1, n):
            facts.append(facts[i - 1] * i)
        nums = []
        for i in range(n):
            nums.append(str(i + 1))
        k -= 1
        ans = []
        for i in range(n - 1, -1, -1):
            idx = k // facts[i]
            k -= idx * facts[i]
            ans.append(nums[idx])
            del nums[idx]
        return ''.join(ans)



s = Solution()
print(s.getPermutation(3, 3))
print(s.getPermutation(3, 5))
print(s.getPermutation(4, 9))