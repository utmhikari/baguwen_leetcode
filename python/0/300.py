class Solution:
    def lengthOfLIS(self, nums: [int]) -> int:
        l = len(nums)
        if l == 0:
            return 0
        min_lis_last_indices = [None] * l
        min_lis_last_indices[0] = 0
        prev_indices = [None] * l
        k = 0
        for i in range(l):
            ni = nums[i]
            if ni > nums[min_lis_last_indices[k]]:
                min_lis_last_indices[k + 1] = i
                prev_indices[i] = min_lis_last_indices[k]
                k += 1
            else:
                left, right = 0, k
                while right > left:
                    mid = (right + left) // 2
                    if ni > nums[min_lis_last_indices[mid]]:
                        left = mid + 1
                    else:
                        right = mid
                min_lis_last_indices[left] = i
                prev_indices[i] = min_lis_last_indices[left - 1]
        # a = []
        # kk = min_lis_last_indices[k]
        # print(prev_indices)
        # print(min_lis_last_indices)
        # while kk:
        #     a.append(str(nums[kk]))
        #     kk = prev_indices[kk]
        # a.reverse()
        # print(', '.join(a))
        return k + 1

s = Solution()
print(s.lengthOfLIS([10, 9, 2, 5,3,4]))
print(s.lengthOfLIS([0,8,4,12,2,10,6,14,1,9,5,13,3,11,7,15]))
print(s.lengthOfLIS([10,9,2,5,3,7,101,18]))