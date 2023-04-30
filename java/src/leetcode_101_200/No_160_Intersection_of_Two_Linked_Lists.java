package leetcode_101_200;

/**
 * Created by 36191 on 2018/9/24
 */
import utils.*;
import java.util.*;

public class No_160_Intersection_of_Two_Linked_Lists {

    public ListNode getIntersectionNode(ListNode headA, ListNode headB) {
        ListNode ha = headA, hb = headB;
        int la = 0, lb = 0, st;
        while (ha != null) { la += 1; ha = ha.next; }
        while (hb != null) { lb += 1; hb = hb.next; }
        ha = headA; hb = headB;
        if (la > lb) { st = la - lb; while (st-- > 0) { ha = ha.next; } }
        else { st = lb - la; while (st-- > 0) { hb = hb.next; } }
        while (ha != null) {
            System.out.println("ha: " + ha.val + ", hb: " + hb.val);
            if (ha.equals(hb)) { return ha; }
            ha = ha.next; hb = hb.next;
        }
        return null;
    }

    public static void main(String[] args) {
        No_160_Intersection_of_Two_Linked_Lists solution = new No_160_Intersection_of_Two_Linked_Lists();
        ListNode n1 = new ListNode(1);
        n1.next = new ListNode(2); n1.next.next = new ListNode(3);
        n1.next.next.next = new ListNode(4); n1.next.next.next.next = new ListNode(5);
        ListNode n2 = new ListNode(-3);
        n2.next = new ListNode(-2); n2.next.next = new ListNode(-1);
        n2.next.next.next = n1.next.next;
        ListNode is = solution.getIntersectionNode(n1, n2);
        System.out.println(is.val);
    }
}
