class Solution:
    def fractionToDecimal(self, numerator: int, denominator: int) -> str:
        if numerator == 0:
            return "0"
        if numerator == denominator:
            return "1"
        isabovezero = (numerator > 0 and denominator > 0) or (numerator < 0 and denominator < 0)
        numerator = abs(numerator)
        denominator = abs(denominator)
        z = 0
        if numerator > denominator:
            z = numerator // denominator
            numerator %= denominator
        if numerator == 0:
            return str(z) if isabovezero else '-' + str(z)
        divlist = list()
        i = 0
        moddict = dict()
        while numerator != 0 and numerator not in moddict.keys():
            moddict[numerator] = i
            i += 1
            numerator *= 10
            divlist.append(str(numerator // denominator))
            numerator %= denominator
        if numerator == 0:
            ans = '%d.%s' % (z, ''.join(divlist))
            return ans if isabovezero else '-' + ans
        sliceidx = moddict[numerator]
        ans = '%d.%s(%s)' % (z, ''.join(divlist[:sliceidx]), ''.join(divlist[sliceidx:]))
        return ans if isabovezero else '-' + ans


s = Solution()

print(s.fractionToDecimal(1, 2))
print(s.fractionToDecimal(2, 1))
print(s.fractionToDecimal(2, 3))
print(s.fractionToDecimal(1, 7))
print(s.fractionToDecimal(1, 725))
print(s.fractionToDecimal(1, 6))
print(s.fractionToDecimal(-50, 8))
print(s.fractionToDecimal(-2147483648, 1))