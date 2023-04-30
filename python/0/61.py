# Definition for singly-linked list.
# class ListNode:
#     def __init__(self, x):
#         self.val = x
#         self.next = None
from node import ListNode


class Solution:
    def rotateRight(self, head: ListNode, k: int) -> ListNode:
        if not head or not head.next or k == 0:
            return head
        t = head
        l = 0
        while t:
            l += 1
            t = t.next
        k %= l
        if k == 0:
            return head
        step = l - k
        t = head
        ll = 1
        while ll < step:
            t = t.next
            ll += 1
        tmp = t.next
        t.next = None
        h = tmp
        while tmp.next:
            tmp = tmp.next
        tmp.next = head
        return h
