package leetcode_201_300;

/**
 * Created by 36191 on 2018/10/15
 */

import utils.TreeNode;

public class No_230_Kth_Smallest_Element_in_a_BST {

    private static int sum, ans;

    private void traverse(TreeNode root, int k) {
        if (root != null) {
            traverse(root.left, k);
            if (++sum == k) { ans = root.val; }
            else if (sum < k) { traverse(root.right, k); }
        }
    }

    public int kthSmallest(TreeNode root, int k) {
        sum = 0;
        traverse(root, k);
        return ans;
    }
}
