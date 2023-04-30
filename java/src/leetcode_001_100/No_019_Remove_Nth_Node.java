package leetcode_001_100;

/**
 * Created by 36191 on 2017/10/7.
 * 019.移除从尾部数第n个节点
 */

import java.util.ArrayList;
public class No_019_Remove_Nth_Node {

    public class ListNode {
        int val;
        ListNode next;
        ListNode(int x) { val = x; }
    }

    /*
    public ListNode removeNthFromEnd(ListNode head, int n) {

        ListNode start = new ListNode(0);
        ListNode slow = start, fast = start;
        slow.next = head;

        // Move fast in front so that the gap between slow and fast becomes n
        for(int i=1; i<=n+1; i++)   {
            fast = fast.next;
        }
        // Move fast to the end, maintaining the gap
        while(fast != null) {
            slow = slow.next;
            fast = fast.next;
        }
        // Skip the desired node
        slow.next = slow.next.next;
        return start.next;
    }
    */

    public ListNode removeNthFromEnd(ListNode head, int n) {
        ArrayList<ListNode> nodes = new ArrayList<>();
        ListNode temp = head;
        while(temp != null) {
            nodes.add(temp);
            temp = temp.next;
        }
        int size = nodes.size();
        if(n > size || n < 1) {
            return head;
        } else if(n == size) {
            return head.next;
        } else if(n == 1) {
            nodes.get(size - 2).next = null;
            return head;
        } else {
            nodes.get(size - n - 1).next = nodes.get(size - n + 1);
            return head;
        }
    }

    public void test() {
        ListNode l = new ListNode(1);
        l.next = new ListNode(2);
        l.next.next = new ListNode(3);
        l.next.next.next = new ListNode(4);
        l.next.next.next.next = new ListNode(5);
        ListNode ans = removeNthFromEnd(l, 2);
        while(ans != null) {
            System.out.print(String.valueOf(ans.val) + ", ");
            ans = ans.next;
        }
    }

    public static void main(String[] args) {
        No_019_Remove_Nth_Node solution = new No_019_Remove_Nth_Node();
        solution.test();
    }
}
