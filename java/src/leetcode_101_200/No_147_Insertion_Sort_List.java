package leetcode_101_200;

/**
 * Created by 36191 on 2018/9/21
 */
import utils.*;

public class No_147_Insertion_Sort_List {
    public ListNode insertionSortList(ListNode head) {
        if (head == null || head.next == null) { return head; }
        ListNode h = head;
        if (head.next.val < head.val) {
            ListNode t = head.next;
            head.next = t.next;
            t.next = head;
            h = t;
            if (h.next.next == null) { return h; }
        }
        ListNode l, p1, p2 = h.next, r = h.next.next;
        while(r != null) {
            l = h;
            p1 = l;
            while(l.val <= r.val && l != r) {
                if (l != h) { p1 = p1.next; }
                l = l.next;
            }
            if(l != r) {
                ListNode t = r;
                p2.next = r.next;
                r.next = l;
                if(l == h) {
                    h = r;
                } else {
                    p1.next = r;
                }
                r = p2.next;
            } else {
                p2 = r;
                r = r.next;
            }
        }
        return h;

    }
}
