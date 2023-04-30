# Definition for singly-linked list.
# class ListNode:
#     def __init__(self, x):
#         self.val = x
#         self.next = None

from node import ListNode

class Solution:
    def swapPairs(self, head: ListNode) -> ListNode:
        if not head or not head.next:
            return head
        p1 = head
        p2 = p1.next
        nxt = p2.next
        p2.next = None
        newnxt = self.swapPairs(nxt)
        p1.next = newnxt
        p2.next = p1
        return p2