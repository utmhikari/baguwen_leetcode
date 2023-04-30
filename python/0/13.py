d = {
    1: 'I',
    4: 'IV',
    5: 'V',
    9: 'IX',
    10: 'X',
    40: 'XL',
    50: 'L',
    90: 'XC',
    100: 'C',
    400: 'CD',
    500: 'D',
    900: 'CM',
    1000: 'M'
}
l = [1000, 900, 500, 400, 100, 90, 50, 40, 10, 9, 5, 4, 1]

def intToRoman(num: int) -> str:
    s = ''
    n = num
    i = 0
    while i < len(l) and n > 0:
        if n >= l[i]:
            n -= l[i]
            s += d[l[i]]
        else:
            i += 1
    return s

tb = {}
for i in range(3999):
    s = intToRoman(i + 1)
    tb[s] = i + 1

class Solution:
    def romanToInt(self, s: str) -> int:
        return tb[s]