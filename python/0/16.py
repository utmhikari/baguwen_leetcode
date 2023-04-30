
def bs(n, l, r, t):
    if r < l:
        return 999999999, -1
    if r == l:
        return abs(t - n[l]), l
    m = l + (r - l) // 2
    tmp = abs(t - n[m])
    if tmp == 0:
        return tmp, m
    if t > n[m]:
        nxt, mnxt = bs(n, m + 1, r, t)
    else:
        nxt, mnxt = bs(n, l, m - 1, t)
    mn = min(nxt, tmp)
    if mn == tmp:
        return tmp, m
    return nxt, mnxt

class Solution:
    def threeSumClosest(self, nums, target: int) -> int:
        ns = sorted(nums)
        l = len(ns)
        i = 0
        m = 999999999
        a = 0
        while i < l - 2:
            j = i + 1
            while j < l - 1:
                t = target - ns[i] - ns[j]
                p, k = bs(ns, j + 1, l - 1, t)
                if p == 0:
                    return target
                if p < m:
                    m = p
                    a = ns[i] + ns[j] + ns[k]
                j += 1
            i += 1
        return a

s = Solution()
nums = [-1, 2, 1, -4]
print(s.threeSumClosest(nums, 1))

nums = [-73,-78,-47,65,22,21,-28,-87,86,-78,-29,97,30,98,-38,-68,76,-91,70,-48,-67,-26,-52,-17,76,-21,-25,15,-58,41,47,-40,86,44,-18,47,-94,-12,52,48,-90,65,52,-2,25,-39,-18,-60,41,59,95,10,44,-65,-17,-56,47,89,33,75,0,-62,-24,22,-23,-58,52,-71,-23,-91,-13,13,100,25,-25,22,-12,-77,-37,34,-45,10,-93,-92,49,88,61,89,69,25,46,-52,64,-60,-94,-61,18,34,18,24,-73,-30,13,27,65,75,-32,34,-54,-30,20,-85,-27,48,43,-13,-85,-82,94,11,-94,0,-78,54,-95,-11,-6,-86,-80,-80,73,-71,-68,20,94,52,-50,-78,87,25,-48,94,-65,89,46,33,26,-75,55,-20,58,0,-91,1,42,90,-62,11,-28,-4,33,58,-74,85,-93,42,9,-91,18,76]
print(s.threeSumClosest(nums, -283))

