package leetcode_001_100;

/**
 * Created by 36191 on 2017/11/6
 * 092. 反转链表2，指定反转范围
 */

public class No_092_Reverse_Linked_List_II {
    public class ListNode {
        int val;
        ListNode next;
        ListNode(int x) { val = x; }
    }

    public ListNode reverseBetween(ListNode head, int m, int n) {
        ListNode first = new ListNode(Integer.MIN_VALUE);
        first.next = head;
        ListNode start = first;
        ListNode end = first;
        int i = 0;
        while(true) {
            if(i < m - 1) {
                start = start.next;
            }
            if(i < n) {
                end = end.next;
                i++;
            } else {
                break;
            }
        }
        ListNode tempStart = start;
        start = start.next;
        tempStart.next = end;
        while(start != end) {
            ListNode temp = start.next;
            start.next = end.next;
            end.next = start;
            start = temp;
        }
        return first.next;
    }

    public void test() {
        ListNode l1 = new ListNode(1);
        l1.next = new ListNode(2);
        l1.next.next = new ListNode(3);
        l1.next.next.next = new ListNode(4);
        l1.next.next.next.next = new ListNode(5);
        l1.next.next.next.next.next = new ListNode(6);
        l1.next.next.next.next.next.next = new ListNode(7);
        ListNode ans1 = reverseBetween(l1, 1, 5);
        while(ans1 != null) {
            System.out.print(ans1.val + ", ");
            ans1 = ans1.next;
        }
        System.out.println();
        ListNode l2 = new ListNode(1);
        l2.next = new ListNode(2);
        l2.next.next = new ListNode(3);
        l2.next.next.next = new ListNode(4);
        l2.next.next.next.next = new ListNode(5);
        l2.next.next.next.next.next = new ListNode(6);
        l2.next.next.next.next.next.next = new ListNode(7);
        ListNode ans2 = reverseBetween(l2, 3, 7);
        while(ans2 != null) {
            System.out.print(ans2.val + ", ");
            ans2 = ans2.next;
        }
        System.out.println();
    }

    public static void main(String[] args) {
        No_092_Reverse_Linked_List_II solution = new No_092_Reverse_Linked_List_II();
        solution.test();
    }
}
