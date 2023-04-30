from node import ListNode

class Solution:
    def reverseKGroup(self, head: ListNode, k: int) -> ListNode:
        checkpoint = head
        rev_cnt = 0
        while rev_cnt < k and checkpoint:
            rev_cnt += 1
            checkpoint = checkpoint.next
        if rev_cnt < k:
            return head
        cnt = 0
        prev = head
        tail = checkpoint
        while cnt < k:
            nxt = prev.next
            prev.next = tail
            tail = prev
            prev = nxt
            cnt += 1
        head.next = self.reverseKGroup(checkpoint, k)
        return tail
