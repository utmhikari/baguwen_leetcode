package leetcode_001_100;

/**
 * Created by 36191 on 2017/10/26
 * 086. 按值大小分割列表
 */

public class No_086_Partition_List {
    public class ListNode {
        int val;
        ListNode next;
        ListNode(int x) { val = x; }
    }

    public ListNode partition(ListNode head, int x) {
        if(head == null || head.next == null) {
            return head;
        }
        ListNode smaller = new ListNode(Integer.MIN_VALUE);
        ListNode larger = new ListNode(Integer.MIN_VALUE);
        ListNode small = smaller;
        ListNode large = larger;
        while(head != null) {
            if (head.val < x) {
                smaller.next = head;
                smaller = smaller.next;
            } else {
                larger.next = head;
                larger = larger.next;
            }
            head = head.next;
        }
        larger.next = null;
        smaller.next = large.next;
        return small.next;
    }

    public void test() {
        ListNode l1 = new ListNode(1);
        l1.next = new ListNode(4);
        l1.next.next = new ListNode(3);
        l1.next.next.next = new ListNode(2);
        l1.next.next.next.next = new ListNode(5);
        l1.next.next.next.next.next = new ListNode(2);
        ListNode ans1 = partition(l1, 3);
        while(ans1 != null) {
            System.out.print(ans1.val + ", ");
            ans1 = ans1.next;
        }
        System.out.println();

        ListNode l2 = new ListNode(1);
        l2.next = new ListNode(2);
        l2.next.next = new ListNode(3);
        l2.next.next.next = new ListNode(7);
        l2.next.next.next.next = new ListNode(4);
        l2.next.next.next.next.next = new ListNode(5);
        l2.next.next.next.next.next.next = new ListNode(6);
        ListNode ans2 = partition(l2, 7);
        while(ans2 != null) {
            System.out.print(ans2.val + ", ");
            ans2 = ans2.next;
        }
        System.out.println();
    }

    public static void main(String[] args) {
        No_086_Partition_List solution = new No_086_Partition_List();
        solution.test();
    }
}
