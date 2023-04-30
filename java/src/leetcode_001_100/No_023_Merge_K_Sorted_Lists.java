package leetcode_001_100;

/**
 * Created by 36191 on 2017/10/7.
 * 023.合并k个排序的链表
 */
public class No_023_Merge_K_Sorted_Lists {

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

    public ListNode mergeKLists(ListNode[] lists) {
        int len = lists.length;
        if(len == 0) {
            return null;
        } else if(len == 1) {
            return lists[0];
        } else if(len == 2) {
            return mergeTwoLists(lists[0], lists[1]);
        } else {
            int half = len / 2;
            ListNode[] firstHalf = new ListNode[half];
            ListNode[] secondHalf = new ListNode[len - half];
            for(int i = 0; i < half; i++) {
                firstHalf[i] = lists[i];
            }
            for(int i = 0; i < len - half; i++) {
                secondHalf[i] = lists[half + i];
            }
            return mergeTwoLists(mergeKLists(firstHalf), mergeKLists(secondHalf));
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
        No_023_Merge_K_Sorted_Lists solution = new No_023_Merge_K_Sorted_Lists();
        solution.test();
    }
}
