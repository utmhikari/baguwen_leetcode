# Definition for singly-linked list.
# class ListNode:
#     def __init__(self, x):
#         self.val = x
#         self.next = None
from node import ListNode


class Solution:
    def hasCycle(self, head: ListNode) -> bool:
        sl, fs = head, head
        while True:
            if not fs:
                return False
            fs = fs.next
            if fs == sl:
                return True
            if not fs:
                return False
            fs = fs.next
            sl = sl.next
            if fs == sl:
                return True
