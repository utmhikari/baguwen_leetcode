package leetcode_001_100;

/**
 * Created by 36191 on 2017/10/7.
 * 021.合并两个排好序的链表
 */
public class No_021_Merge_Two_Sorted_Lists {

    public class ListNode {
        int val;
        ListNode next;
        ListNode(int x) { val = x; }
    }

    public ListNode mergeTwoLists(ListNode l1, ListNode l2) {
        if(l1 == null && l2 == null) {
            return null;
        } else if(l1 == null) {
            return l2;
        } else if(l2 == null) {
            return l1;
        } else {
            ListNode first = new ListNode(Integer.MIN_VALUE);
            ListNode ans = first;
            while(l1 != null && l2 != null) {
                if(l1.val < l2.val) {
                    ans.next = new ListNode(l1.val);
                    l1 = l1.next;
                } else {
                    ans.next = new ListNode(l2.val);
                    l2 = l2.next;
                }
                ans = ans.next;
            }
            if(l1 == null) {
                while(l2 != null) {
                    ans.next = new ListNode(l2.val);
                    l2 = l2.next;
                    ans = ans.next;
                }
            } else {
                while(l1 != null) {
                    ans.next = new ListNode(l1.val);
                    l1 = l1.next;
                    ans = ans.next;
                }
            }
            return first.next;
        }
    }

    public void test() {
        ListNode l1 = new ListNode(2);
        l1.next = new ListNode(4);
        l1.next.next = new ListNode(3);
        ListNode l2 = new ListNode(5);
        l2.next = new ListNode(6);
        l2.next.next = new ListNode(4);
        ListNode merge = mergeTwoLists(l1, l2);
        while(merge != null) {
            System.out.print(String.valueOf(merge.val) + ", ");
            merge = merge.next;
        }
    }

    public static void main(String[] args) {
        No_021_Merge_Two_Sorted_Lists solution = new No_021_Merge_Two_Sorted_Lists();
        solution.test();
    }
}
