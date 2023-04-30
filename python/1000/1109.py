from typing import List


class Solution:
    def corpFlightBookings(self, bookings: List[List[int]], n: int) -> List[int]:
        a = [0] * n
        l = len(bookings)
        for i in range(l):
            b = bookings[i]
            a[b[0] - 1] += b[2]
            if b[1] < n:
                a[b[1]] -= b[2]
        for i in range(1, n):
            a[i] += a[i - 1]
        return a


# def corpFlightBookings(self, bookings: List[List[int]], n: int) -> List[int]:
#     #     bigs = list()
#     #     smalls = list()
#     #     bigsm = 0
#     #     smallsms = [0] * n
#     #     l = len(bookings)
#     #     for i in range(l):
#     #         b = bookings[i]
#     #         addtimes = (b[1] - b[0] + 1)
#     #         if addtimes > l / 2:
#     #             bigs.append(i)
#     #             bigsm += b[2]
#     #         else:
#     #             smalls.append(i)
#     #     bigsms = [bigsm] * n
#     #     for i in smalls:
#     #         b = bookings[i]
#     #         for j in range(b[0] - 1, b[1]):
#     #             smallsms[j] += b[2]
#     #     for i in bigs:
#     #         b = bookings[i]
#     #         for j in range(0, b[0] - 1):
#     #             bigsms[j] -= b[2]
#     #         for j in range(b[1], n):
#     #             bigsms[j] -= b[2]
#     #     for i in range(n):
#     #         smallsms[i] += bigsms[i]
#     #     return smallsms


