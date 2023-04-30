class Solution:
    def intersect(self, nums1: [int], nums2: [int]) -> [int]:
        d1 = dict()
        for i in nums1:
            if i not in d1.keys():
                d1[i] = 0
            d1[i] += 1
        a = []
        for i in nums2:
            if i in d1.keys() and d1[i] > 0:
                d1[i] -= 1
                a.append(i)
        return a