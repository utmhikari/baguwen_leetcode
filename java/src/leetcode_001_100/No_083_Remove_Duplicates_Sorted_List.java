package leetcode_001_100;

/**
 * Created by 36191 on 2017/10/23
 * 083. 去除链表中重复元素
 */

public class No_083_Remove_Duplicates_Sorted_List {
    public class ListNode {
        int val;
        ListNode next;
        ListNode(int x) { val = x; }
    }

    public ListNode deleteDuplicates(ListNode head) {
        if(head == null || head.next == null) {
            return head;
        }
        ListNode first = new ListNode(Integer.MIN_VALUE);
        first.next = head;
        ListNode linkNode = first;
        ListNode current = head;
        while(true) {
            while(current.next != null && current.val == current.next.val) {
                current = current.next;
            }
            linkNode.next = current;
            linkNode = linkNode.next;
            current = current.next;
            if(current == null) {
                break;
            }
        }
        linkNode.next = null;
        return first.next;
    }

    public void test() {
        ListNode l1 = new ListNode(1);
        l1.next = new ListNode(2);
        l1.next.next = new ListNode(3);
        l1.next.next.next = new ListNode(3);
        l1.next.next.next.next = new ListNode(4);
        l1.next.next.next.next.next = new ListNode(4);
        l1.next.next.next.next.next.next = new ListNode(5);
        ListNode ans1 = deleteDuplicates(l1);
        while(ans1 != null) {
            System.out.print(ans1.val + ", ");
            ans1 = ans1.next;
        }
        System.out.println();
        ListNode l2 = new ListNode(1);
        l2.next = new ListNode(1);
        l2.next.next = new ListNode(1);
        l2.next.next.next = new ListNode(2);
        l2.next.next.next.next = new ListNode(3);
        l2.next.next.next.next.next = new ListNode(4);
        l2.next.next.next.next.next.next = new ListNode(4);
        ListNode ans2 = deleteDuplicates(l2);
        while(ans2 != null) {
            System.out.print(ans2.val + ", ");
            ans2 = ans2.next;
        }
        System.out.println();
    }

    public static void main(String[] args) {
        No_083_Remove_Duplicates_Sorted_List solution = new No_083_Remove_Duplicates_Sorted_List();
        solution.test();
    }
}
