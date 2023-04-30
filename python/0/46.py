def swap(lst, i1, i2):
    tmp = lst[i1]
    lst[i1] = lst[i2]
    lst[i2] = tmp


def bt(tmp, l, ans, first):
    if first == l:
        ans.append(list(tmp))
    for i in range(first, l):
        swap(tmp, first, i)
        bt(tmp, l, ans, first + 1)
        swap(tmp, first, i)


class Solution:
    def permute(self, nums: [int]) -> [[int]]:
        ans = []
        tmp = []
        for n in nums:
            tmp.append(n)
        l = len(nums)
        bt(tmp, l, ans, 0)
        return ans


s = Solution()
print(s.permute([1,2,3,4]))
