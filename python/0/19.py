from node import ListNode


class Solution:
    def removeNthFromEnd(self, head: ListNode, n: int) -> ListNode:
        if not head:
            return head
        p1, p2 = head, head
        for i in range(n):
            p2 = p2.next
            if not p2 and i < n - 1:
                return head
        if not p2:
            hd = p1.next
            p1.next = None
            return hd
        while p2:
            p2 = p2.next
            if not p2:
                tmp = p1.next
                p1.next = tmp.next
                tmp.next = None
                return head
            p1 = p1.next
