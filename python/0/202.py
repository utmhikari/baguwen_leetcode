d = [0,1,4,9,16,25,36,49,64,81]
s = {1,10,100,1000,10000,100000,1000000,10000000,100000000,1000000000,82,28,68,86,19,91}
t = {0}

def ih(n, v):
    if n in t:
        return False
    if n in s:
        return True
    sum = 0
    while n != 0:
        sum += d[n % 10]
        n //= 10
    if sum in v:
        t.add(sum)
        return False
    v.add(sum)
    if not ih(sum, v):
        t.add(sum)
        return False
    s.add(sum)
    return True

class Solution:
    def isHappy(self, n: int) -> bool:
        return ih(n, set())




so = Solution()
print(so.isHappy(199))
print(so.isHappy(19))