package leetcode_001_100;

/**
 * Created by 36191 on 2017/10/17
 * 061.旋转链表
 */

public class No_061_Rotate_List {
    public class ListNode {
       int val;
       ListNode next;
       ListNode(int x) { val = x; }
    }

    public ListNode rotateRight(ListNode head, int k) {
        if(head == null || head.next == null) {
            return head;
        }
        ListNode current = head;
        ListNode rear = current;
        int count = 1;
        while(rear.next != null) {
            rear = rear.next;
            count++;
        }
        while(k > count) {
            k -= count;
        }
        if(k == count) {
            return head;
        }
        int times = count - k;
        for(int i = 0; i < times - 1; i++) {
            current = current.next;
        }
        rear.next = head;
        head = current.next;
        current.next = null;
        return head;
    }

    public void test() {
        ListNode l1 = new ListNode(1);
        l1.next = new ListNode(2);
        l1.next.next = new ListNode(3);
        l1.next.next.next = new ListNode(4);
        l1.next.next.next.next = new ListNode(5);
        l1.next.next.next.next.next = null;
        ListNode a1 = rotateRight(l1, 2);
        while(a1 != null) {
            System.out.print(a1.val + ", ");
            a1 = a1.next;
        }
        System.out.println();
    }

    public static void main(String[] args) {
        No_061_Rotate_List solution = new No_061_Rotate_List();
        solution.test();
    }
}
