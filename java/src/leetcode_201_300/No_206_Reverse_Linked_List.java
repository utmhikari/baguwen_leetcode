package leetcode_201_300;

/**
 * Created by 36191 on 2018/10/1
 */

import utils.*;

public class No_206_Reverse_Linked_List {
    public ListNode reverseList(ListNode head) {
        if (head == null || head.next == null) { return head; }
        ListNode pre = head, cur = head.next, tmp;
        while (cur != null) {
            tmp = cur.next;
            cur.next = pre;
            head.next = tmp;
            pre = cur;
            cur = tmp;
        }
        return pre;
    }
}
