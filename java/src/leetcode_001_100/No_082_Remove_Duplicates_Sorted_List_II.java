package leetcode_001_100;

/**
 * Created by 36191 on 2017/10/22
 * 082. 去除链表中重复元素II（不保留重复）
 */

public class No_082_Remove_Duplicates_Sorted_List_II {
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
        ListNode linkHead = first;
        ListNode current = head;
        while(true) {
            boolean dup = false;
            while(current.next != null && current.next.val == current.val) {
                current = current.next;
                dup = true;
            }
            if(!dup) {
                linkHead.next = current;
                linkHead = linkHead.next;
            } else {
                if(current.next == null) {
                    linkHead.next = null;
                    break;
                }
            }
            current = current.next;
            if(current == null) {
                break;
            }
        }
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
        No_082_Remove_Duplicates_Sorted_List_II solution = new No_082_Remove_Duplicates_Sorted_List_II();
        solution.test();
    }
}
