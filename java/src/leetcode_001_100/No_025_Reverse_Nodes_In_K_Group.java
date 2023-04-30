package leetcode_001_100;

/**
 * Created by 36191 on 2017/10/8.
 * 025.每k个节点调换位置
 */
public class No_025_Reverse_Nodes_In_K_Group {

    public class ListNode {
        int val;
        ListNode next;
        ListNode(int x) { val = x; }
    }

    /*
    public ListNode reverseKGroup(ListNode head, int k) {
        ListNode curr = head;
        int count = 0;
        while (curr != null && count != k) { // find the k+1 node
            curr = curr.next;
            count++;
        }
        if (count == k) { // if k+1 node is found
            curr = reverseKGroup(curr, k);  // reverse list with k+1 node as head
            // head - head-pointer to direct part,
            // curr - head-pointer to reversed part;
            while (count-- > 0) { // reverse current k-group:
                ListNode tmp = head.next;  // tmp - next head in direct part
                head.next = curr;  // preappending "direct" head to the reversed list
                curr = head;  // move head of reversed part to a new node
                head = tmp;  // move "direct" head to the next node in direct part
            }
            head = curr;
        }
        return head;
    }
    */

    public ListNode reverseKGroup(ListNode head, int k) {
        if(head == null) {
            return null;
        }
        if(k == 1) {
            return head;
        }
        ListNode ans = new ListNode(-1);
        ans.next = head;
        ListNode first = ans;
        ListNode[] kNodes = new ListNode[k];
        for(int i = 0; i < k; i++) {
            kNodes[i] = head;
        }
        while(true) {
            if(kNodes[0] == null) {
                break;
            }
            boolean nullFlag = false;
            for(int i = 1; i < k; i++) {
                kNodes[i] = kNodes[i-1].next;
                if(kNodes[i] == null) {
                    nullFlag = true;
                    break;
                }
            }
            if(nullFlag) {
                break;
            }
            ListNode temp = kNodes[k - 1].next;
            kNodes[0].next = temp;
            first.next = kNodes[k - 1];
            for(int i = k - 1; i > 0; i--) {
                kNodes[i].next = kNodes[i - 1];
            }
            first = kNodes[0];
            for(int i = 0; i < k; i++) {
                kNodes[i] = temp;
            }
        }
        return ans.next;
    }

    public void test() {
        ListNode ln = new ListNode(1);
        ln.next = new ListNode(2);
        ln.next.next = new ListNode(3);
        ln.next.next.next = new ListNode(4);
        ln.next.next.next.next = new ListNode(5);
        ListNode ans = reverseKGroup(ln, 2);
        while(ans != null) {
            System.out.print(Integer.toString(ans.val) + ", ");
            ans = ans.next;
        }
        System.out.println();

        ListNode ln1 = new ListNode(1);
        ln1.next = new ListNode(2);
        ln1.next.next = new ListNode(3);
        ln1.next.next.next = new ListNode(4);
        ln1.next.next.next.next = new ListNode(5);
        ListNode ans1 = reverseKGroup(ln1, 3);
        while(ans1 != null) {
            System.out.print(Integer.toString(ans1.val) + ", ");
            ans1 = ans1.next;
        }
        System.out.println();
    }

    public static void main(String[] args) {
        No_025_Reverse_Nodes_In_K_Group solution = new No_025_Reverse_Nodes_In_K_Group();
        solution.test();
    }
}
