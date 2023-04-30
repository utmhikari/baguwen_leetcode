package leetcode_001_100;

/**
 * Created by 36191 on 2017/10/8.
 * 024.每两个节点交换位置
 */
public class No_024_Swap_Nodes_In_Pairs {
    public class ListNode {
        int val;
        ListNode next;
        ListNode(int x) { val = x; }
    }

    public ListNode swapPairs(ListNode head) {
        if(head == null) {
            return null;
        }
        ListNode ans = new ListNode(-1);
        ans.next = head;
        ListNode first = ans;
        ListNode second = head;
        ListNode third = head;
        while(true) {
            third = second.next;
            if(third == null) {
                break;
            }
            second.next = third.next;
            first.next = third;
            third.next = second;
            first = second;
            second = second.next;
            third = second;
            if(third == null) {
                break;
            }
        }
        return ans.next;
    }

    public void test() {
        ListNode ln = new ListNode(1);
        ln.next = new ListNode(2);
        ln.next.next = new ListNode(3);
        ln.next.next.next = new ListNode(4);
        ListNode ans = swapPairs(ln);
        while(ans != null) {
            System.out.print(Integer.toString(ans.val) + ", ");
            ans = ans.next;
        }
        System.out.println();
    }

    public static void main(String[] args) {
        No_024_Swap_Nodes_In_Pairs solution = new No_024_Swap_Nodes_In_Pairs();
        solution.test();
    }
}
