package leetcode_101_200;

/**
 * Created by 36191 on 2018/9/21
 */
import utils.*;
public class No_141_Linked_List_Cycle {
    public boolean hasCycle(ListNode head) {
        if(head == null || head.next == null) { return false; }
        ListNode a = head, b = head.next;
        while(b != null && b.next != null) {
            if(a == b) { return true; }
            a = a.next;
            b = b.next.next;
        }
        return false;
    }
}
