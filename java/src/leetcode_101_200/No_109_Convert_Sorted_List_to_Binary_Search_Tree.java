package leetcode_101_200;

import utils.TreeNode;
import utils.ListNode;
import java.util.ArrayList;

/**
 * Created by 36191 on 2017/12/8
 * 109.将排序的链表转化为二叉搜索树
 */

public class No_109_Convert_Sorted_List_to_Binary_Search_Tree {

    private TreeNode generateBST(ArrayList<Integer> nums, int start, int end) {
        if(start > end) {
            return null;
        }
        int half = start + (end - start) / 2;
        TreeNode node = new TreeNode(nums.get(half));
        node.left = generateBST(nums, start, half - 1);
        node.right = generateBST(nums, half + 1, end);
        return node;
    }

    public TreeNode sortedArrayToBST(ArrayList<Integer> nums) {
        return generateBST(nums, 0, nums.size() - 1);
    }

    public TreeNode sortedListToBST(ListNode head) {
        if(head == null) {
            return null;
        }
        ArrayList<Integer> sortedArray = new ArrayList<>();
        ListNode temp = head;
        while(temp != null) {
            sortedArray.add(temp.val);
            temp = temp.next;
        }
        return sortedArrayToBST(sortedArray);
    }
}
