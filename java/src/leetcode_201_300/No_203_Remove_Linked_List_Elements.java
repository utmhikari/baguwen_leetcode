package leetcode_201_300;

/**
 * Created by 36191 on 2018/10/1
 */

import utils.*;

public class No_203_Remove_Linked_List_Elements {
    public ListNode removeElements(ListNode head, int val) {
        ListNode start = head;
        while (start != null && start.val == val) { start = start.next; }
        if (start == null) { return null; }
        ListNode cur = start.next, pre = start;
        while (cur != null) {
            if (cur.val == val) { pre.next = cur.next; cur.next = null; cur = pre.next; }
            else { cur = cur.next; pre = pre.next; }
        }
        return start;
    }
}
