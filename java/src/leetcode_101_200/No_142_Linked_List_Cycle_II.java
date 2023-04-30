package leetcode_101_200;

/**
 * Created by 36191 on 2018/9/18
 */

import utils.*;

public class No_142_Linked_List_Cycle_II {

    public ListNode detectCycle(ListNode head) {
        if (head == null || head.next == null) { return null; }
        ListNode p = head, q = head;
        while (q.next != null) {
            p = p.next;
            q = q.next.next;
            if (q == null) { return null; }
            if (q == p) {
                ListNode r = head;
                while (r != p) {
                    r = r.next;
                    p = p.next;
                }
                return r;
            }
        }
        return null;
    }
}
