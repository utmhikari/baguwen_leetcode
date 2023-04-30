from node import ListNode

class Solution:
    def mergeTwoLists(self, l1: ListNode, l2: ListNode) -> ListNode:
        p1 = l1
        p2 = l2
        first = ListNode(0)
        p = first
        while p1 or p2:
            if not p1:
                p.next = ListNode(p2.val)
                p2 = p2.next
            elif not p2:
                p.next = ListNode(p1.val)
                p1 = p1.next
            elif p1.val <= p2.val:
                p.next = ListNode(p1.val)
                p1 = p1.next
            else:
                p.next = ListNode(p2.val)
                p2 = p2.next
            p = p.next
        return first.next
