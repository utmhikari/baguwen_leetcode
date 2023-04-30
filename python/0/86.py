# Definition for singly-linked list.
# class ListNode:
#     def __init__(self, x):
#         self.val = x
#         self.next = None

from node import ListNode


class Solution:
    def partition(self, head: ListNode, x: int) -> ListNode:
        if not head:
            return head
        leftd, rightd = ListNode(-1), ListNode(-1)
        left, right = leftd, rightd
        t = head
        while t:
            v = t.val
            tn = t.next
            t.next = None
            if v < x:
                left.next = t
                left = left.next
            else:
                right.next = t
                right = right.next
            t = tn
        lefts = leftd.next
        rights = rightd.next
        leftd.next = None
        rightd.next = None
        if not lefts:
            return rights
        left.next = rights
        return lefts
