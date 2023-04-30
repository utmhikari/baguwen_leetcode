package leetcode_101_200;

/**
 * Created by 36191 on 2018/9/20
 */

import utils.*;

public class No_143_Reorder_List {

    private ListNode reverse(ListNode head) {
        ListNode current = head, temp = head, pre = null;
        while (current != null) {
            temp = current.next;
            current.next = pre;
            pre = current;
            current = temp;
        }
        return pre;
    }

    public void reorderList(ListNode head) {
        if (head != null && head.next != null && head.next.next != null) {
            ListNode a = head, b = head, c = head;
            boolean status;
            while (true) {
                if (c == null || c.next == null) {
                    status = c == null;
                    break;
                }
                b = b.next;
                c = c.next.next;
            }
            ListNode temp = b.next;
            b.next = reverse(temp);
            b = b.next;
            ListNode t1, t2;
            while (b != null) {
                t1 = a.next;
                t2 = b.next;
                a.next = b;
                b.next = t1;
                a = t1;
                b = t2;
            }
            if (status) { a.next.next = null; }
            else { a.next = null; }
        }
    }

    public static void main(String[] args) {
        No_143_Reorder_List solution = new No_143_Reorder_List();
        ListNode l1 = new ListNode(new int[] {1, 2, 3, 4, 5});
        l1.output();
        ListNode l2 = new ListNode(new int[] {1, 2, 3, 4});
        l2.output();
        solution.reorderList(l1);
        solution.reorderList(l2);
        l1.output();
        l2.output();
    }
}
