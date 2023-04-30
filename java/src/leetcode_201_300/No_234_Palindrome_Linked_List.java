package leetcode_201_300;

/**
 * Created by 36191 on 2018/10/9
 */

import utils.*;

public class No_234_Palindrome_Linked_List {

    private ListNode reverse(ListNode head) {
        ListNode prev = null;
        while (head != null) {
            ListNode next = head.next;
            head.next = prev;
            prev = head;
            head = next;
        }
        return prev;
    }

    public boolean isPalindrome(ListNode head) {
        ListNode slow = head, fast = head;
        if (head == null || head.next == null) { return true; }
        if (head.next.next == null) { return head.val == head.next.val; }
        while (fast != null && fast.next != null) {
            slow = slow.next; fast = fast.next.next;
        }
        fast = head; slow = reverse(slow);
        while (slow != null) {
            if (fast.val != slow.val) { return false; }
            fast = fast.next; slow = slow.next;
        }
        return true;
    }
}
