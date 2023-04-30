# Definition for singly-linked list.
# class ListNode:
#     def __init__(self, x):
#         self.val = x
#         self.next = None

from node import ListNode

class Solution:
    def reverseBetween(self, head: ListNode, m: int, n: int) -> ListNode:
        if n == m or not head:
            return head
        # m - 1, m, n, n + 1
        dummy = ListNode(-1)
        dummy.next = head
        p1, p2, p3, p4 = dummy, head, dummy, head
        for i in range(n):
            if i < m - 1:
                p1 = p1.next
                p2 = p2.next
            p3 = p3.next
            p4 = p4.next
        t, e = p2, p4
        while t != p3:
            tn = t.next
            t.next = e
            e = t
            t = tn
        t.next = e
        p1.next = t
        h = dummy.next
        dummy.next = None
        return h

