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

class Solution:
    def intToRoman(self, num: int) -> str:
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


s = Solution()

i = 3
print(s.intToRoman(i))

i = 4
print(s.intToRoman(i))

i = 9
print(s.intToRoman(i))

i = 58
print(s.intToRoman(i))

i = 1994
print(s.intToRoman(i))
