# 给定两个大小为 m 和 n 的有序数组 nums1 和 nums2。
#
# 请你找出这两个有序数组的中位数，并且要求算法的时间复杂度为 O(log(m + n))。
#
# 你可以假设 nums1 和 nums2 不会同时为空。




class Solution:
    def findMedianSortedArrays(self, nums1, nums2) -> float:
        l1 = len(nums1)
        l2 = len(nums2)
        l = l1 + l2
        end = l // 2 + 1
        cur = 0
        prev = 0
        c = 0
        x1, x2 = 0, 0
        while c < end:
            prev = cur
            if x1 == l1:
                cur = nums2[x2]
                x2 += 1
            elif x2 == l2:
                cur = nums1[x1]
                x1 += 1
            elif nums1[x1] < nums2[x2]:
                cur = nums1[x1]
                x1 += 1
            else:
                cur = nums2[x2]
                x2 += 1
            c += 1
        if l % 2 == 0:
            return (cur + prev) / 2
        return float(cur)

s = Solution()

nums1 = [1, 3]
nums2 = [2]

print(s.findMedianSortedArrays(nums1, nums2))

nums1 = [1, 2]
nums2 = [3, 4]

print(s.findMedianSortedArrays(nums1, nums2))