class Solution:
    def increasingTriplet(self, nums: [int]) -> bool:
        l = len(nums)
        if l < 3:
            return False
        one, two = 999999999999999999, 99999999999999999
        for n in nums:
            if n <= one:
                one = n
            elif n <= two:
                two = n
            elif n > two:
                return True
        return False

s = Solution()
print(s.increasingTriplet([1,2,3,4,5]))
print(s.increasingTriplet([5,4,3,2,1]))
print(s.increasingTriplet([1,5,4,3,2]))
print(s.increasingTriplet([1,2,3,1,2,1]))
print(s.increasingTriplet([1,2,-10,-8,-7]))

