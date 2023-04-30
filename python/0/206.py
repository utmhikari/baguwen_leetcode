from node import ListNode

class Solution:
    def reverseList(self, head: ListNode) -> ListNode:
        if not head or not head.next:
            return head
        p = None
        while True:
            t = head.next
            head.next = p
            p = head
            if not t:
                return head
            else:
                head = t

