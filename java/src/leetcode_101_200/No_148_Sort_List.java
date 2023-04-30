package leetcode_101_200;

/**
 * Created by 36191 on 2018/8/16
 * No.148 链表排序
 */

import utils.ListNode;
import utils.TreeNode;

public class No_148_Sort_List {

    private ListNode merge(ListNode l1, ListNode l2) {
        if(l1 == null) return l2;
        if(l2 == null) return l1;
        if(l1.val < l2.val){
            l1.next = merge(l1.next, l2);
            return l1;
        } else{
            l2.next = merge(l1, l2.next);
            return l2;
        }
    }

    public ListNode sortList(ListNode head) {
        if(head == null || head.next == null) { return head; }
        ListNode s = head, f = head;
        while(f.next != null && f.next.next!= null){
            f = f.next.next;
            s = s.next;
        }
        f = s.next;
        s.next = null;
        ListNode left = sortList(head);
        ListNode right = sortList(f);
        return merge(left,right);
    }

    public static void main(String[] args) {

    }
}
