


class MedianFinder:

    def __init__(self):
        """
        initialize your data structure here.
        """
        self.nums = []

    def addNum(self, num: int) -> None:
        l = len(self.nums)
        if l == 0:
            self.nums.append(num)
            return
        
        mid = len(self.nums) // 2
        if num == self.nums[mid]:
            self.nums.insert(mid, num)
        elif num > self.nums[mid]:
            while mid < l and self.nums[mid] < num:
                mid += 1
            if mid == l:
                self.nums.append(num)
            else:
                self.nums.insert(mid, num)
        else:
            while mid >= 0 and self.nums[mid] > num:
                mid -= 1
            self.nums.insert(mid + 1, num)

    def findMedian(self) -> float:
        l = len(self.nums)
        mid = (l - 1) // 2
        if l % 2 == 0:
            return (self.nums[mid] + self.nums[mid + 1]) / 2
        return self.nums[mid]
