package leetcode_201_300;

/**
 * Created by 36191 on 2018/10/20
 */

import utils.*;

public class No_237_Delete_Node_in_a_Linked_List {
    public void deleteNode(ListNode node) {
        ListNode n = node.next;
        node.val = n.val;
        node.next = n.next;
        n = null;
    }
}
